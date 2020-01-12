package photon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocation_City(t *testing.T) {
	t.Run("Berlin", func(t *testing.T) {
		l := Location{LocCity: "Berlin"}

		assert.Equal(t, "Berlin", l.City())
	})
	t.Run("Bowen Island", func(t *testing.T) {
		l := Location{LocCity: "Bowen Island Municipality"}

		assert.Equal(t, "Bowen Island", l.City())
	})
	t.Run("KwaDukuza Local Municipality", func(t *testing.T) {
		l := Location{LocCity: "KwaDukuza Local Municipality"}

		assert.Equal(t, "KwaDukuza", l.City())
	})
	t.Run("Saint-Paul", func(t *testing.T) {
		l := Location{LocCity: "Saint-Paul"}

		assert.Equal(t, "Saint-Paul", l.City())
	})
	t.Run("Hlabisa Local Municipality", func(t *testing.T) {
		l := Location{LocCity: "Hlabisa Local Municipality"}

		assert.Equal(t, "Hlabisa", l.City())
	})
	t.Run("Mtubatuba Local Municipality", func(t *testing.T) {
		l := Location{LocCity: "Mtubatuba Local Municipality"}

		assert.Equal(t, "Mtubatuba", l.City())
	})
	t.Run("King Sabata Dalindyebo Local Municipality", func(t *testing.T) {
		l := Location{LocCity: "King Sabata Dalindyebo Local Municipality"}

		assert.Equal(t, "", l.City())
	})
	t.Run("Sundays River Valley Local Municipality", func(t *testing.T) {
		l := Location{LocCity: "Sundays River Valley Local Municipality"}

		assert.Equal(t, "Sundays River Valley", l.City())
	})
}
