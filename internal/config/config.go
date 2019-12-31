package config

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	gc "github.com/patrickmn/go-cache"
	"github.com/photoprism/photoprism-places/internal/entity"
	"github.com/photoprism/photoprism-places/internal/event"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var log = event.Log

type Config struct {
	db     *gorm.DB
	cache  *gc.Cache
	config *Params
}

func initLogger(debug bool) {
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	if debug {
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}
}

func NewConfig(ctx *cli.Context) *Config {
	initLogger(ctx.GlobalBool("debug"))

	c := &Config{
		config: NewParams(ctx),
	}

	log.SetLevel(c.LogLevel())

	return c
}

// CreateDirectories creates all folders that places needs.
func (c *Config) CreateDirectories() error {
	if err := os.MkdirAll(filepath.Dir(c.LogFilename()), os.ModePerm); err != nil {
		return err
	}

	return nil
}

// connectToDatabase establishes a database connection.
// When used with the internal driver, it may create a new database server instance.
// It tries to do this 12 times with a 5 second sleep interval in between.
func (c *Config) connectToDatabase(ctx context.Context) error {
	dbDriver := c.DatabaseDriver()
	dbDsn := c.DatabaseDsn()

	if dbDriver == "" {
		return errors.New("can't connect: database driver not specified")
	}

	if dbDsn == "" {
		return errors.New("can't connect: database DSN not specified")
	}

	db, err := gorm.Open(dbDriver, dbDsn)
	if err != nil || db == nil {
		for i := 1; i <= 12; i++ {
			time.Sleep(5 * time.Second)

			db, err = gorm.Open(dbDriver, dbDsn)

			if db != nil && err == nil {
				break
			}
		}

		if err != nil || db == nil {
			log.Fatal(err)
		}
	}

	c.db = db
	return err
}

// Name returns the application name.
func (c *Config) Name() string {
	return c.config.Name
}

// Author returns the site author / copyright.
func (c *Config) Author() string {
	return c.config.Author
}

// Version returns the application version.
func (c *Config) Version() string {
	return c.config.Version
}

// Copyright returns the application copyright.
func (c *Config) Copyright() string {
	return c.config.Copyright
}

// Debug returns true if Debug mode is on.
func (c *Config) Debug() bool {
	return c.config.Debug
}

// Experimental returns true if experimental features should be enabled.
func (c *Config) Experimental() bool {
	return c.config.Experimental
}

// LogLevel returns the logrus log level.
func (c *Config) LogLevel() logrus.Level {
	if c.Debug() {
		c.config.LogLevel = "debug"
	}

	if logLevel, err := logrus.ParseLevel(c.config.LogLevel); err == nil {
		return logLevel
	} else {
		return logrus.InfoLevel
	}
}

// ConfigFile returns the config file name.
func (c *Config) ConfigFile() string {
	return c.config.ConfigFile
}

// ConfigPath returns the config path.
func (c *Config) ConfigPath() string {
	if c.config.ConfigPath == "" {
		return c.AssetsPath() + "/config"
	}

	return c.config.ConfigPath
}

// LogFilename returns the filename for storing server logs.
func (c *Config) LogFilename() string {
	if c.config.LogFilename == "" {
		return c.AssetsPath() + "/places.log"
	}

	return c.config.LogFilename
}

// HttpServerHost returns the built-in HTTP server host name or IP address (empty for all interfaces).
func (c *Config) HttpServerHost() string {
	if c.config.HttpServerHost == "" {
		return "0.0.0.0"
	}

	return c.config.HttpServerHost
}

// HttpServerPort returns the built-in HTTP server port.
func (c *Config) HttpServerPort() int {
	if c.config.HttpServerPort == 0 {
		return 8080
	}

	return c.config.HttpServerPort
}

// HttpServerMode returns the server mode.
func (c *Config) HttpServerMode() string {
	if c.config.HttpServerMode == "" {
		if c.Debug() {
			return "debug"
		}

		return "release"
	}

	return c.config.HttpServerMode
}

// HttpStaticPath returns the static server assets path.
func (c *Config) HttpStaticPath() string {
	return c.AssetsPath() + "/static"
}

// PIDFilename returns the filename for storing the server process id (pid).
func (c *Config) PIDFilename() string {
	if c.config.PIDFilename == "" {
		return c.AssetsPath() + "/photoprism.pid"
	}

	return c.config.PIDFilename
}

// DetachServer returns true if server should detach from console (daemon mode).
func (c *Config) DetachServer() bool {
	return c.config.DetachServer
}

// DatabaseDriver returns the database driver name.
func (c *Config) DatabaseDriver() string {
	if c.config.DatabaseDriver == "" {
		return DbMySQL
	}

	return c.config.DatabaseDriver
}

// DatabaseDsn returns the database data source name (DSN).
func (c *Config) DatabaseDsn() string {
	if c.config.DatabaseDsn == "" {
		return "places:places@tcp(places-db:3306)/places?parseTime=true"
	}

	return c.config.DatabaseDsn
}

// CachePath returns the path to the cache.
func (c *Config) CachePath() string {
	return c.config.CachePath
}

// AssetsPath returns the path to the assets.
func (c *Config) AssetsPath() string {
	return c.config.AssetsPath
}

// Cache returns the in-memory cache.
func (c *Config) Cache() *gc.Cache {
	if c.cache == nil {
		c.cache = gc.New(336*time.Hour, 30*time.Minute)
	}

	return c.cache
}

// Db returns the db connection.
func (c *Config) Db() *gorm.DB {
	if c.db == nil {
		log.Fatal("database not initialised.")
	}

	return c.db
}

// CloseDb closes the db connection (if any).
func (c *Config) CloseDb() error {
	if c.db != nil {
		if err := c.db.Close(); err == nil {
			c.db = nil
		} else {
			return err
		}
	}

	return nil
}

// MigrateDb will start a migration process.
func (c *Config) MigrateDb() {
	db := c.Db()

	// db.LogMode(true)

	db.AutoMigrate(
		&entity.Event{},
		&entity.Place{},
		&entity.Location{},
		&entity.Country{},
	)

	entity.CreateUnknownPlace(db)
	entity.CreateUnknownCountry(db)
}

// Init initialises the Database.
func (c *Config) Init(ctx context.Context) error {
	return c.connectToDatabase(ctx)
}

// Shutdown closes open database connections.
func (c *Config) Shutdown() {
	if err := c.CloseDb(); err != nil {
		log.Errorf("could not close database connection: %s", err)
	} else {
		log.Info("closed database connection")
	}
}
