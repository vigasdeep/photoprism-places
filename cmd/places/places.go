package main

import (
	"os"

	"github.com/photoprism/photoprism-places/internal/commands"
	"github.com/photoprism/photoprism-places/internal/config"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var version = "development"

func main() {
	app := cli.NewApp()
	app.Name = "PhotoPrism Places"
	app.Usage = "Reverse Geocoding API"
	app.Version = version
	app.Copyright = "(c) 2018-2020 The PhotoPrism contributors <hello@photoprism.org>"
	app.EnableBashCompletion = true
	app.Flags = config.GlobalFlags

	app.Commands = []cli.Command{
		commands.ConfigCommand,
		commands.StartCommand,
		commands.MigrateCommand,
		commands.VersionCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.Error(err)
	}
}
