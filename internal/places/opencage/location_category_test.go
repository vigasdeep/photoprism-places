package opencage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocation_Category(t *testing.T) {
	t.Run("information", func(t *testing.T) {

		l := &Location{LocCategory: "information", LocName: "test"}
		assert.Equal(t, "visitor center", l.Category())
	})

	t.Run("water", func(t *testing.T) {

		l := &Location{LocCategory: "water", LocName: "test"}
		assert.Equal(t, "water", l.Category())
	})

	t.Run("shop", func(t *testing.T) {

		l := &Location{LocCategory: "shop", LocName: "test"}
		assert.Equal(t, "shop", l.Category())
	})

	t.Run("xxx", func(t *testing.T) {

		l := &Location{LocCategory: "xxx", LocName: "test"}
		assert.Equal(t, "", l.Category())
	})
}
