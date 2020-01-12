package config

import (
	"strings"
	"testing"

	"github.com/photoprism/photoprism/pkg/fs"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	ctx := CliTestContext()

	assert.True(t, ctx.IsSet("assets-path"))
	assert.False(t, ctx.Bool("debug"))

	c := NewConfig(ctx)

	assert.IsType(t, new(Config), c)

	assert.Equal(t, fs.ExpandFilename("../../assets"), c.AssetsPath())
	assert.False(t, c.Debug())
}

func TestConfig_Name(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	name := c.Name()
	assert.Equal(t, "config.test", name)
}

func TestConfig_Version(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	version := c.Version()
	assert.Equal(t, "0.0.0", version)
}

func TestConfig_Copyright(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	copyright := c.Copyright()
	assert.Equal(t, "", copyright)
}

func TestConfig_ConfigFile(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	configFile := c.ConfigFile()
	assert.Equal(t, "", configFile)
}

func TestConfig_ConfigPath(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	configPath := c.ConfigPath()
	assert.True(t, strings.Contains(configPath, "assets/config"))
}

func TestConfig_LogFilename(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	filename := c.LogFilename()
	assert.True(t, strings.Contains(filename, "assets/places.log"))
}

func TestConfig_DatabaseDriver(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	driver := c.DatabaseDriver()
	assert.Equal(t, "mysql", driver)
}

func TestConfig_DatabaseDsn(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	dsn := c.DatabaseDriver()
	assert.Equal(t, "mysql", dsn)
}

func TestConfig_CachePath(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	path := c.CachePath()
	assert.Equal(t, "", path)
}

func TestConfig_Db(t *testing.T) {
	c := NewTestConfig()

	assert.NotNil(t, c.Db())
}

func TestConfig_CloseDb(t *testing.T) {
	c := NewTestConfig()

	assert.NotNil(t, c.Db())

	err := c.CloseDb()
	assert.Nil(t, err)
}

func TestConfig_Shutdown(t *testing.T) {
	c := NewTestConfig()
	c.Shutdown()
}
