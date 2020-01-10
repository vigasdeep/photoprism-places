// +build osm

package osm

import (
	"testing"

	"github.com/photoprism/photoprism-places/internal/s2"
	"github.com/stretchr/testify/assert"
)

func TestFindLocation(t *testing.T) {
	t.Run("Fernsehturm Berlin 1", func(t *testing.T) {
		lat := 52.5208
		lng := 13.40953
		id := s2.Token(lat, lng)

		l, err := FindLocation(id)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "Fernsehturm Berlin", l.LocName)
		assert.Equal(t, "10178", l.Address.Postcode)
		assert.Equal(t, "Berlin", l.Address.State)
		assert.Equal(t, "de", l.Address.CountryCode)
		assert.Equal(t, "Germany", l.Address.Country)
		assert.Equal(t, "tower", l.Category()) // Should be "tower", but type is "attraction"
	})

	t.Run("Fernsehturm Berlin 2", func(t *testing.T) {
		lat := 52.52057
		lng := 13.40889
		id := s2.Token(lat, lng)

		l, err := FindLocation(id)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "Menschen Museum", l.LocName)
		assert.Equal(t, "10178", l.Address.Postcode)
		assert.Equal(t, "Berlin", l.Address.State)
		assert.Equal(t, "de", l.Address.CountryCode)
		assert.Equal(t, "Germany", l.Address.Country)
		assert.Equal(t, "museum", l.Category())
	})

	t.Run("No Location", func(t *testing.T) {
		lat := 0.0
		lng := 0.0
		id := s2.Token(lat, lng)

		_, err := FindLocation(id)

		if err == nil {
			t.Fatal("err should not be nil")
		}

		assert.Equal(t, "osm: invalid location id", err.Error())
	})
}
