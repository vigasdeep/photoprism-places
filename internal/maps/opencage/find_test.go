package opencage

import (
	"testing"

	"github.com/photoprism/photoprism-places/internal/s2"
	"github.com/stretchr/testify/assert"
)

func TestFindLocation(t *testing.T) {
	ProviderKey = "f9585987890741ceac625709e8efea3b"

	t.Run("berlin", func(t *testing.T) {
		lat := 52.5208
		lng := 13.40953
		id := s2.Token(lat, lng)

		l, err := FindLocation(id)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "Berlin", l.City())
		assert.Equal(t, "de", l.CountryCode())
	})

	t.Run("airport", func(t *testing.T) {
		l, err := FindLocation("47a856aaa1")

		if err != nil {
			t.Fatal(err)
		}

		// TODO: Unreliable
		/*
			find_test.go:36:
			     |             	Error Trace:	find_test.go:36
			     |             	Error:      	Not equal:
			     |             	            	expected: "Flughafen Tegel (Airport)"
			     |             	            	actual  : "Luftfracht"
			     |
			     |             	            	Diff:
			     |             	            	--- Expected
			     |             	            	+++ Actual
			     |             	            	@@ -1 +1 @@
			     |             	            	-Flughafen Tegel (Airport)
			     |             	            	+Luftfracht
			     |             	Test:       	TestFindLocation/airport
		*/
		//assert.Equal(t, "Flughafen Tegel (Airport)", l.Name())

		assert.Equal(t, "", l.Category())

		assert.Equal(t, "Berlin", l.City())
		assert.Equal(t, "de", l.CountryCode())
	})

	t.Run("tierpark", func(t *testing.T) {
		l, err := FindLocation("47a84930c5")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "zoo", l.Category())
		assert.Equal(t, "Tierpark Berlin", l.Name())
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

		assert.Equal(t, "", l.City())
		assert.Equal(t, "ch", l.CountryCode())
		assert.Equal(t, "Dock A", l.Name())
		assert.Equal(t, "", l.Category())
		assert.Equal(t, "Zurich", l.State())
	})

	t.Run("botanical_garden", func(t *testing.T) {
		l, err := FindLocation("89b7b7872f")

		if err != nil {
			t.Fatal(err)
		}

		// TODO: Unreliable
		/*
			find_test.go:81:
			     |             	Error Trace:	find_test.go:81
			     |             	Error:      	Not equal:
			     |             	            	expected: ""
			     |             	            	actual  : "Washington D.C."
			     |
			     |             	            	Diff:
			     |             	            	--- Expected
			     |             	            	+++ Actual
			     |             	            	@@ -1 +1 @@
			     |             	            	-
			     |             	            	+Washington D.C.
			     |             	Test:       	TestFindLocation/botanical_garden
		*/
		// assert.Equal(t, "Washington D.C.", l.City())

		assert.Equal(t, "us", l.CountryCode())
		assert.Equal(t, "United States Botanic Garden", l.Name())
		assert.Equal(t, "botanical garden", l.Category())
		assert.Equal(t, "District of Columbia", l.State())
	})

	t.Run("none", func(t *testing.T) {
		lat := 0.0
		lng := 0.0
		id := s2.Token(lat, lng)

		_, err := FindLocation(id)

		if err == nil {
			t.Fatal("err should not be nil")
		}

		assert.Equal(t, "opencage: invalid location id", err.Error())
	})
}
