package places

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIgnoreCategories(t *testing.T) {
	t.Run("hotel", func(t *testing.T) {
		s := StructSource{
			LocCellID:      "abc",
			LocCountryCode: "de",
			LocCategory:    "hotel",
			LocName:        "Hilton",
			LocCity:        "Berlin",
			LocState:       "Berlin",
			LocSource:      "test",
		}

		l := Location{}

		if err := l.Assign(s); err != nil {
			assert.FailNow(t, err.Error())
		}

		assert.Equal(t, "", l.LocName)
		assert.Equal(t, "", l.LocCategory)
		assert.Equal(t, "Berlin, Germany", l.LocLabel)
	})

	t.Run("tower", func(t *testing.T) {
		s := StructSource{
			LocCellID:      "abc",
			LocCountryCode: "de",
			LocCategory:    "tower",
			LocName:        "Fernsehturm",
			LocCity:        "Berlin",
			LocState:       "Berlin",
			LocSource:      "test",
		}

		l := Location{}

		if err := l.Assign(s); err != nil {
			assert.FailNow(t, err.Error())
		}

		assert.Equal(t, "Fernsehturm", l.LocName)
		assert.Equal(t, "tower", l.LocCategory)
		assert.Equal(t, "Berlin, Germany", l.LocLabel)
	})
}
