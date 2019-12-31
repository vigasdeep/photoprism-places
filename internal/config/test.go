package config

import (
	"context"
	"flag"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/photoprism/photoprism-places/internal/util"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const (
	TestDataZip  = "/tmp/photoprism/testdata.zip"
	TestDataURL  = "https://dl.photoprism.org/fixtures/testdata.zip"
	TestDataHash = "a217ac5242de2189ffb414d819b628c7957c67d7"
)

var testConfig *Config

func testDataPath(assetsPath string) string {
	return assetsPath + "/testdata"
}

func NewTestParams() *Params {
	assetsPath := util.ExpandedFilename("../../assets")

	testDataPath := testDataPath(assetsPath)

	c := &Params{
		AssetsPath:     assetsPath,
		CachePath:      testDataPath + "/cache",
		DatabaseDriver: "mysql",
		DatabaseDsn:    "places:places@tcp(places-db:3306)/places?parseTime=true",
	}

	return c
}

func NewTestParamsError() *Params {
	assetsPath := util.ExpandedFilename("../..")

	testDataPath := testDataPath("../../assets")

	c := &Params{
		AssetsPath:     assetsPath,
		CachePath:      testDataPath + "/cache",
		DatabaseDriver: "mysql",
		DatabaseDsn:    "places:places@tcp(places-db:3306)/places?parseTime=true",
	}

	return c
}

func TestConfig() *Config {
	if testConfig == nil {
		testConfig = NewTestConfig()
	}

	return testConfig
}

func NewTestConfig() *Config {
	log.SetLevel(logrus.DebugLevel)

	c := &Config{config: NewTestParams()}
	err := c.Init(context.Background())
	if err != nil {
		log.Fatalf("failed init config: %v", err)
	}

	c.MigrateDb()
	return c
}

func NewTestErrorConfig() *Config {
	log.SetLevel(logrus.DebugLevel)

	c := &Config{config: NewTestParamsError()}
	err := c.Init(context.Background())
	if err != nil {
		log.Fatalf("failed init config: %v", err)
	}

	c.MigrateDb()
	return c
}

// Returns example cli config for testing
func CliTestContext() *cli.Context {
	config := NewTestParams()

	globalSet := flag.NewFlagSet("test", 0)
	globalSet.Bool("debug", false, "doc")
	globalSet.String("config-file", config.ConfigFile, "doc")
	globalSet.String("assets-path", config.AssetsPath, "doc")

	app := cli.NewApp()

	c := cli.NewContext(app, globalSet, nil)

	c.Set("config-file", config.ConfigFile)
	c.Set("assets-path", config.AssetsPath)

	return c
}
