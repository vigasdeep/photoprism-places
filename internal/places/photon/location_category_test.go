package photon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocation_Category(t *testing.T) {
	t.Run("information", func(t *testing.T) {

		l := &Location{OsmValue: "information", LocName: "test"}
		assert.Equal(t, "visitor center", l.Category())
	})

	t.Run("water", func(t *testing.T) {

		l := &Location{OsmValue: "water", LocName: "test"}
		assert.Equal(t, "water", l.Category())
	})

	t.Run("shop", func(t *testing.T) {

		l := &Location{OsmValue: "shop", LocName: "test"}
		assert.Equal(t, "shop", l.Category())
	})

	t.Run("xxx", func(t *testing.T) {

		l := &Location{OsmValue: "xxx", LocName: "test"}
		assert.Equal(t, "", l.Category())
	})
}
