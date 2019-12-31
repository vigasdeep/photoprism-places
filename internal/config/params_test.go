package config

import (
	"testing"

	"github.com/photoprism/photoprism-places/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestNewParams(t *testing.T) {
	ctx := CliTestContext()

	assert.True(t, ctx.IsSet("assets-path"))
	assert.False(t, ctx.Bool("debug"))

	c := NewParams(ctx)

	assert.IsType(t, new(Params), c)

	assert.Equal(t, util.ExpandedFilename("../../assets"), c.AssetsPath)
	assert.False(t, c.Debug)
}

func TestParams_SetValuesFromFile(t *testing.T) {
	c := NewParams(CliTestContext())

	err := c.SetValuesFromFile("testdata/config.yml")

	assert.Nil(t, err)

	assert.False(t, c.Debug)
	assert.Equal(t, "/srv/places", c.AssetsPath)
	assert.Equal(t, "/srv/places/cache", c.CachePath)
	assert.Equal(t, "mysql", c.DatabaseDriver)
	assert.Equal(t, "places:places@tcp(places-db:3306)/places?parseTime=true", c.DatabaseDsn)
	assert.Equal(t, 8080, c.HttpServerPort)
}
