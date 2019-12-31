package commands

import (
	"testing"

	"github.com/photoprism/photoprism-places/internal/config"
	"github.com/photoprism/photoprism-places/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestConfigCommand(t *testing.T) {
	var err error

	ctx := config.CliTestContext()

	output := util.CaptureOutput(func() {
		err = ConfigCommand.Run(ctx)
	})

	assert.Contains(t, output, "NAME                  VALUE")
	assert.Contains(t, output, "config-file")
	assert.Contains(t, output, "cache-path")
	assert.Contains(t, output, "assets-path")

	assert.Equal(t, output, output)
	assert.Nil(t, err)
}
