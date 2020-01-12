package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/photoprism/photoprism-places/internal/places"
)

// Photo place
type Place struct {
	ID         string    `gorm:"type:varbinary(16);primary_key;auto_increment:false;" json:"id"`
	LocLabel   string    `gorm:"type:varbinary(500);unique_index;" json:"label"`
	LocCity    string    `gorm:"type:varchar(100);" json:"city"`
	LocState   string    `gorm:"type:varchar(100);" json:"state"`
	LocCountry string    `gorm:"type:binary(2);" json:"country"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	New        bool      `gorm:"-" json:"-"`
}

var UnknownPlace = NewPlace("-", "Unknown", "Unknown", "Unknown", "zz")

func CreateUnknownPlace(db *gorm.DB) {
	UnknownPlace.FirstOrCreate(db)
}

func (m *Place) AfterCreate(scope *gorm.Scope) error {
	return scope.SetColumn("New", true)
}

func FindPlace(token string, db *gorm.DB) *Place {
	place := &Place{}

	if err := db.First(place, "id = ?", token).Error; err != nil {
		log.Debugf("place: %s for token %s", err.Error(), token)
	}

	return place
}

func FindPlaceByLabel(token string, label string, db *gorm.DB) *Place {
	place := &Place{}

	if err := db.First(place, "id = ? OR loc_label = ?", token, label).Error; err != nil {
		log.Debugf("place: %s for token %s or label \"%s\"", err.Error(), token, label)
	}

	return place
}

func NewPlace(token, label, city, state, countryCode string) *Place {
	result := &Place{
		ID:         token,
		LocLabel:   label,
		LocCity:    city,
		LocState:   state,
		LocCountry: countryCode,
	}

	return result
}

func (m *Place) Find(db *gorm.DB) error {
	if err := db.First(m, "id = ?", m.ID).Error; err != nil {
		return err
	}

	return nil
}

func (m *Place) FirstOrCreate(db *gorm.DB) *Place {
	if err := db.FirstOrCreate(m, "id = ? OR loc_label = ?", m.ID, m.LocLabel).Error; err != nil {
		log.Debugf("place: %s for token %s or label \"%s\"", err.Error(), m.ID, m.LocLabel)
	}

	return m
}

func (m *Place) NoID() bool {
	return m.ID == ""
}

func (m *Place) Label() string {
	return m.LocLabel
}

func (m *Place) City() string {
	return m.LocCity
}

func (m *Place) State() string {
	return m.LocState
}

func (m *Place) CountryCode() string {
	return m.LocCountry
}

func (m *Place) CountryName() string {
	return places.CountryNames[m.LocCountry]
}
