// +build osm

package osm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocation_State(t *testing.T) {
	t.Run("Berlin", func(t *testing.T) {

		a := Address{CountryCode: "de", City: "Berlin", State: "Berlin", HouseNumber: "63", Suburb: "Neukölln"}
		l := &Location{LocCategory: "natural", LocName: "Nice title", LocType: "hill", LocDisplayName: "dipslay name", Address: a}
		assert.Equal(t, "Berlin", l.State())
	})
}

func TestLocation_Suburb(t *testing.T) {
	t.Run("Neukölln", func(t *testing.T) {

		a := Address{CountryCode: "de", City: "Berlin", State: "Berlin", HouseNumber: "63", Suburb: "Neukölln"}
		l := &Location{LocCategory: "natural", LocName: "Nice title", LocType: "hill", LocDisplayName: "dipslay name", Address: a}
		assert.Equal(t, "Neukölln", l.Suburb())
	})
}

func TestLocation_CountryCode(t *testing.T) {
	t.Run("de", func(t *testing.T) {

		a := Address{CountryCode: "de", City: "Berlin", State: "Berlin", HouseNumber: "63", Suburb: "Neukölln"}
		l := &Location{LocCategory: "natural", LocName: "Nice title", LocType: "hill", LocDisplayName: "dipslay name", Address: a}
		assert.Equal(t, "de", l.CountryCode())
	})
}

func TestLocation_Keywords(t *testing.T) {
	t.Run("cat", func(t *testing.T) {

		a := Address{CountryCode: "de", City: "Berlin", State: "Berlin", HouseNumber: "63", Suburb: "Neukölln"}
		l := &Location{LocCategory: "natural", LocName: "Nice title", LocType: "hill", LocDisplayName: "cat", Address: a}
		assert.Equal(t, []string{"cat"}, l.Keywords())
	})
}

func TestLocation_Source(t *testing.T) {

	a := Address{CountryCode: "de", City: "Berlin", State: "Berlin", HouseNumber: "63", Suburb: "Neukölln"}
	l := &Location{LocCategory: "natural", LocName: "Nice title", LocType: "hill", LocDisplayName: "cat", Address: a}
	assert.Equal(t, "osm", l.Source())
}
