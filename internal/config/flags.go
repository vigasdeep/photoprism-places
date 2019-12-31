package config

import (
	"github.com/urfave/cli"
)

// Global CLI flags
var GlobalFlags = []cli.Flag{
	cli.BoolFlag{
		Name:   "debug",
		Usage:  "run in debug mode",
		EnvVar: "PLACES_DEBUG",
	},
	cli.BoolFlag{
		Name:   "experimental, e",
		Usage:  "enable experimental features",
		EnvVar: "PLACES_EXPERIMENTAL",
	},
	cli.StringFlag{
		Name:   "log-level, l",
		Usage:  "trace, debug, info, warning, error, fatal or panic",
		Value:  "info",
		EnvVar: "PLACES_LOG_LEVEL",
	},
	cli.StringFlag{
		Name:   "log-filename",
		Usage:  "filename for storing server logs",
		EnvVar: "PLACES_LOG_FILENAME",
		Value:  "~/.local/share/places/places.log",
	},
	cli.StringFlag{
		Name:   "pid-filename",
		Usage:  "filename for the server process id (pid)",
		EnvVar: "PLACES_PID_FILENAME",
		Value:  "~/.local/share/places/places.pid",
	},
	cli.StringFlag{
		Name:   "config-file, c",
		Usage:  "load configuration from `FILENAME`",
		Value:  "~/.config/places/places.yml",
		EnvVar: "PLACES_CONFIG_FILE",
	},
	cli.StringFlag{
		Name:   "config-path",
		Usage:  "config `PATH`",
		Value:  "~/.config/places",
		EnvVar: "PLACES_CONFIG_PATH",
	},
	cli.StringFlag{
		Name:   "resources-path",
		Usage:  "resources `PATH`",
		EnvVar: "PLACES_RESOURCES_PATH",
	},
	cli.StringFlag{
		Name:   "cache-path",
		Usage:  "cache `PATH`",
		Value:  "~/.cache/places",
		EnvVar: "PLACES_CACHE_PATH",
	},
	cli.StringFlag{
		Name:   "assets-path",
		Usage:  "assets `PATH`",
		Value:  "~/.local/share/places",
		EnvVar: "PLACES_ASSETS_PATH",
	},
	cli.StringFlag{
		Name:   "database-driver",
		Usage:  "database `DRIVER` (only mysql supported)",
		Value:  "mysql",
		EnvVar: "PLACES_DATABASE_DRIVER",
	},
	cli.StringFlag{
		Name:   "database-dsn",
		Usage:  "database data source name (`DSN`)",
		Value:  "places:places@tcp(places-db:3306)/places?parseTime=true",
		EnvVar: "PLACES_DATABASE_DSN",
	},
	cli.IntFlag{
		Name:   "http-port, p",
		Usage:  "HTTP server port",
		EnvVar: "PLACES_HTTP_PORT",
	},
	cli.StringFlag{
		Name:   "http-host, i",
		Usage:  "HTTP server host",
		EnvVar: "PLACES_HTTP_HOST",
	},
	cli.StringFlag{
		Name:   "http-mode, m",
		Usage:  "debug, release or test",
		EnvVar: "PLACES_HTTP_MODE",
	},
}
