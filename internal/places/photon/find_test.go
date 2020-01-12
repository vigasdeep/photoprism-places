package photon

import (
	"testing"

	"github.com/photoprism/photoprism/pkg/s2"
	"github.com/stretchr/testify/assert"
)

func TestFindLocation(t *testing.T) {
	t.Run("berlin", func(t *testing.T) {
		lat := 52.5208
		lng := 13.40953
		id := s2.Token(lat, lng)

		l, err := FindLocation(id)

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("%+v", l)

		assert.Equal(t, "Berlin", l.City())
		assert.Equal(t, "Fernsehturm Berlin", l.Name())
		assert.Equal(t, "tower", l.Category())
		assert.Equal(t, "de", l.CountryCode())
	})

	t.Run("museum", func(t *testing.T) {
		lat := 52.52057
		lng := 13.40889
		id := s2.Token(lat, lng)

		l, err := FindLocation(id)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "museum", l.Category())
		assert.Equal(t, "Menschen Museum", l.Name())
		assert.Equal(t, "Berlin", l.City())
		assert.Equal(t, "Berlin", l.State())
		assert.Equal(t, "de", l.CountryCode())

	})

	t.Run("airport", func(t *testing.T) {
		l, err := FindLocation("47a856aaa1")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "Airport", l.Name())
		assert.Equal(t, "airport", l.Category())
		assert.Equal(t, "Berlin", l.City())
		assert.Equal(t, "de", l.CountryCode())
	})

	t.Run("tierpark", func(t *testing.T) {
		l, err := FindLocation("47a84930c5")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "", l.Category())
		assert.Equal(t, "Berlin", l.City())
		assert.Equal(t, "de", l.CountryCode())
	})

	t.Run("zurich", func(t *testing.T) {
		lat := 47.45401666666667
		lng := 8.557494444444446
		id := s2.Token(lat, lng)

		l, err := FindLocation(id)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "Kloten", l.City())
		assert.Equal(t, "ch", l.CountryCode())
		assert.Equal(t, "Airport", l.Name())
		assert.Equal(t, "airport", l.Category())
		assert.Equal(t, "Zurich", l.State())
	})

	t.Run("botanical_garden", func(t *testing.T) {
		l, err := FindLocation("89b7b7872f")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "us", l.CountryCode())
		assert.Equal(t, "Maryland Avenue Southwest", l.Name())
		assert.Equal(t, "", l.Category())
		assert.Equal(t, "", l.City())
		assert.Equal(t, "Washington, D.C.", l.State())
	})

	t.Run("none", func(t *testing.T) {
		lat := 0.0
		lng := 0.0
		id := s2.Token(lat, lng)

		_, err := FindLocation(id)

		if err == nil {
			t.Fatal("err should not be nil")
		}

		assert.Equal(t, "photon: invalid location id", err.Error())
	})
}
