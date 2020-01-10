package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/photoprism/photoprism-places/internal/util"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

const (
	DbTiDB  = "internal"
	DbMySQL = "mysql"
)

// Params provides a struct in which application configuration is stored.
// Application code must use functions to get config values, for two reasons:
//
// 1. Some values are computed and we don't want to leak implementation details (aims at reducing refactoring overhead).
//
// 2. Paths might actually be dynamic later (if we build a multi-user version).
//
// See https://github.com/photoprism/photoprism/issues/50#issuecomment-433856358
type Params struct {
	Name           string
	Url            string `yaml:"url" flag:"url"`
	Author         string `yaml:"author" flag:"author"`
	Version        string
	Copyright      string
	Debug          bool   `yaml:"debug" flag:"debug"`
	Experimental   bool   `yaml:"experimental" flag:"experimental"`
	LogLevel       string `yaml:"log-level" flag:"log-level"`
	ConfigFile     string
	ConfigPath     string `yaml:"config-path" flag:"config-path"`
	AssetsPath     string `yaml:"assets-path" flag:"assets-path"`
	ResourcesPath  string `yaml:"resources-path" flag:"resources-path"`
	CachePath      string `yaml:"cache-path" flag:"cache-path"`
	HttpServerHost string `yaml:"http-host" flag:"http-host"`
	HttpServerPort int    `yaml:"http-port" flag:"http-port"`
	HttpServerMode string `yaml:"http-mode" flag:"http-mode"`
	DatabaseDriver string `yaml:"database-driver" flag:"database-driver"`
	DatabaseDsn    string `yaml:"database-dsn" flag:"database-dsn"`
	LogFilename    string `yaml:"log-filename" flag:"log-filename"`
	DetachServer   bool   `yaml:"detach-server" flag:"detach-server"`
	PIDFilename    string `yaml:"pid-filename" flag:"pid-filename"`
	OpenCageKey    string `yaml:"opencage-key" flag:"opencage-key"`
	PhotonUrl      string `yaml:"photon-url" flag:"photon-url"`
	NominatimUrl   string `yaml:"nominatim-url" flag:"nominatim-url"`
}

// NewParams() creates a new configuration entity by using two methods:
//
// 1. SetValuesFromFile: This will initialize values from a yaml config file.
//
// 2. SetValuesFromCliContext: Which comes after SetValuesFromFile and overrides
//    any previous values giving an option two override file configs through the CLI.
func NewParams(ctx *cli.Context) *Params {
	c := &Params{}

	c.Name = ctx.App.Name
	c.Copyright = ctx.App.Copyright
	c.Version = ctx.App.Version
	c.ConfigFile = util.ExpandedFilename(ctx.GlobalString("config-file"))

	if err := c.SetValuesFromFile(c.ConfigFile); err != nil {
		log.Debug(err)
	}

	if err := c.SetValuesFromCliContext(ctx); err != nil {
		log.Error(err)
	}

	c.expandFilenames()

	return c
}

func (c *Params) expandFilenames() {
	c.ConfigPath = util.ExpandedFilename(c.ConfigPath)
	c.AssetsPath = util.ExpandedFilename(c.AssetsPath)
	c.CachePath = util.ExpandedFilename(c.CachePath)
	c.LogFilename = util.ExpandedFilename(c.LogFilename)
}

// SetValuesFromFile uses a yaml config file to initiate the configuration entity.
func (c *Params) SetValuesFromFile(fileName string) error {
	if !util.Exists(fileName) {
		return errors.New(fmt.Sprintf("config file not found: \"%s\"", fileName))
	}

	yamlConfig, err := ioutil.ReadFile(fileName)

	if err != nil {
		return err
	}

	return yaml.Unmarshal(yamlConfig, c)
}

// SetValuesFromCliContext uses values from the CLI to setup configuration overrides
// for the entity.
func (c *Params) SetValuesFromCliContext(ctx *cli.Context) error {
	v := reflect.ValueOf(c).Elem()

	// Iterate through all config fields
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)

		tagValue := v.Type().Field(i).Tag.Get("flag")

		// Automatically assign values to fields with "flag" tag
		if tagValue != "" {
			switch t := fieldValue.Interface().(type) {
			case int, int64:
				// Only if explicitly set or current value is empty (use default)
				if ctx.IsSet(tagValue) {
					f := ctx.Int64(tagValue)
					fieldValue.SetInt(f)
				} else if ctx.GlobalIsSet(tagValue) || fieldValue.Int() == 0 {
					f := ctx.GlobalInt64(tagValue)
					fieldValue.SetInt(f)
				}
			case uint, uint64:
				// Only if explicitly set or current value is empty (use default)
				if ctx.IsSet(tagValue) {
					f := ctx.Uint64(tagValue)
					fieldValue.SetUint(f)
				} else if ctx.GlobalIsSet(tagValue) || fieldValue.Uint() == 0 {
					f := ctx.GlobalUint64(tagValue)
					fieldValue.SetUint(f)
				}
			case string:
				// Only if explicitly set or current value is empty (use default)
				if ctx.IsSet(tagValue) {
					f := ctx.String(tagValue)
					fieldValue.SetString(f)
				} else if ctx.GlobalIsSet(tagValue) || fieldValue.String() == "" {
					f := ctx.GlobalString(tagValue)
					fieldValue.SetString(f)
				}
			case bool:
				if ctx.IsSet(tagValue) {
					f := ctx.Bool(tagValue)
					fieldValue.SetBool(f)
				} else if ctx.GlobalIsSet(tagValue) {
					f := ctx.GlobalBool(tagValue)
					fieldValue.SetBool(f)
				}
			default:
				log.Warnf("can't assign value of type %s from cli flag %s", t, tagValue)
			}
		}
	}

	return nil
}
