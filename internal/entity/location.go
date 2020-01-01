package entity

import (
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/photoprism/photoprism-places/internal/maps"
	"github.com/photoprism/photoprism-places/internal/s2"
	"github.com/photoprism/photoprism-places/internal/util"
	"gopkg.in/ugjka/go-tz.v2/tz"
)

// Photo location
type Location struct {
	ID          string    `gorm:"type:varbinary(16);primary_key;auto_increment:false;" json:"id"`
	PlaceID     string    `gorm:"type:varbinary(16);" json:"-"`
	Place       *Place    `json:"place"`
	LocName     string    `gorm:"type:varchar(100);" json:"name"`
	LocCategory string    `gorm:"type:varchar(50);" json:"category"`
	LocSuburb   string    `gorm:"type:varchar(100);" json:"suburb"`
	LocSource   string    `gorm:"type:varbinary(16);" json:"-"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

func NewLocation(lat, lng float64) *Location {
	result := &Location{}

	result.ID = s2.Token(lat, lng)

	return result
}

func FindLocation(token string, db *gorm.DB) (l *Location, err error) {
	l = &Location{ID: token}

	if err := l.Find(db); err != nil {
		return nil, fmt.Errorf("location: %s for token %s", err.Error(), token)
	}

	return l, nil
}

func (m *Location) Find(db *gorm.DB) error {
	if err := db.First(m, "id = ?", m.ID).Error; err == nil {
		m.Place = FindPlace(m.PlaceID, db)
		return nil
	}

	l := &maps.Location{
		ID: m.ID,
	}

	if err := l.QueryOSM(); err != nil {
		return err
	}

	m.Place = FindPlaceByLabel(l.ID, l.LocLabel, db)

	if m.Place.NoID() {
		m.Place.ID = l.ID
		m.Place.LocLabel = l.LocLabel
		m.Place.LocCity = l.LocCity
		m.Place.LocState = l.LocState
		m.Place.LocCountry = l.LocCountry
	}

	m.LocName = l.LocName
	m.LocCategory = l.LocCategory
	m.LocSuburb = l.LocSuburb
	m.LocSource = l.LocSource

	if err := db.Create(m).Error; err != nil {
		log.Errorf("location: %s", err)
		return err
	}

	return nil
}

func (m *Location) Keywords() []string {
	result := []string{
		strings.ToLower(m.City()),
		strings.ToLower(m.Suburb()),
		strings.ToLower(m.State()),
		strings.ToLower(m.CountryName()),
		strings.ToLower(m.Category()),
	}

	result = append(result, util.Keywords(m.Name())...)
	result = append(result, util.Keywords(m.Label())...)
	result = append(result, util.Keywords(m.Notes())...)

	return result
}

func (m *Location) Unknown() bool {
	return m.ID == ""
}

func (m *Location) Name() string {
	return m.LocName
}

func (m *Location) NoName() bool {
	return m.LocName == ""
}

func (m *Location) Category() string {
	return m.LocCategory
}

func (m *Location) NoCategory() bool {
	return m.LocCategory == ""
}

func (m *Location) Suburb() string {
	return m.LocSuburb
}

func (m *Location) NoSuburb() bool {
	return m.LocSuburb == ""
}

func (m *Location) Label() string {
	return m.Place.Label()
}

func (m *Location) City() string {
	return m.Place.City()
}

func (m *Location) LongCity() bool {
	return len(m.City()) > 16
}

func (m *Location) NoCity() bool {
	return m.City() == ""
}

func (m *Location) CityContains(text string) bool {
	return strings.Contains(text, m.City())
}

func (m *Location) State() string {
	return m.Place.State()
}

func (m *Location) NoState() bool {
	return m.Place.State() == ""
}

func (m *Location) CountryCode() string {
	return m.Place.CountryCode()
}

func (m *Location) CountryName() string {
	return m.Place.CountryName()
}

func (m *Location) Notes() string {
	return m.Place.Notes()
}

func (m *Location) Source() string {
	return m.LocSource
}

func (m *Location) LatLng() (lat, lng float64) {
	return s2.LatLng(m.ID)
}

// TimeZone returns the location time zone as string.
func (m *Location) Timezone() string {
	lat, lng := m.LatLng()

	if lat == 0 && lng == 0 {
		return "UTC"
	}

	zones, err := tz.GetZone(tz.Point{
		Lon: lng, Lat: lat,
	})

	if err != nil {
		return "UTC"
	}

	return zones[0]
}
