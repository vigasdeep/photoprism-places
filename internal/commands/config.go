package commands

import (
	"fmt"

	"github.com/photoprism/photoprism-places/internal/config"
	"github.com/urfave/cli"
)

// Prints current configuration
var ConfigCommand = cli.Command{
	Name:   "config",
	Usage:  "Displays global configuration values",
	Action: configAction,
}

func configAction(ctx *cli.Context) error {
	conf := config.NewConfig(ctx)

	fmt.Printf("NAME                  VALUE\n")
	fmt.Printf("name                  %s\n", conf.Name())
	fmt.Printf("version               %s\n", conf.Version())
	fmt.Printf("copyright             %s\n", conf.Copyright())
	fmt.Printf("debug                 %t\n", conf.Debug())
	fmt.Printf("experimental          %t\n", conf.Experimental())
	fmt.Printf("log-level             %s\n", conf.LogLevel())
	fmt.Printf("log-filename          %s\n", conf.LogFilename())
	fmt.Printf("config-file           %s\n", conf.ConfigFile())
	fmt.Printf("config-path           %s\n", conf.ConfigPath())
	fmt.Printf("static-path           %s\n", conf.HttpStaticPath())
	fmt.Printf("assets-path           %s\n", conf.AssetsPath())

	fmt.Printf("http-host             %s\n", conf.HttpServerHost())
	fmt.Printf("http-port             %d\n", conf.HttpServerPort())
	fmt.Printf("http-mode             %s\n", conf.HttpServerMode())
	fmt.Printf("database-driver       %s\n", conf.DatabaseDriver())
	fmt.Printf("database-dsn          %s\n", conf.DatabaseDsn())

	fmt.Printf("cache-path            %s\n", conf.CachePath())

	return nil
}
