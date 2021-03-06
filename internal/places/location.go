package places

import (
	"errors"
	"strings"

	"github.com/photoprism/photoprism-places/internal/places/opencage"
	"github.com/photoprism/photoprism-places/internal/places/osm"
	"github.com/photoprism/photoprism-places/internal/places/photon"
)

// Photo location
type Location struct {
	ID          string
	LocName     string
	LocCategory string
	LocLabel    string
	LocCity     string
	LocState    string
	LocCountry  string
	LocSource   string
}

func NewLocation(id string) *Location {
	result := &Location{
		ID: id,
	}

	return result
}

func (l *Location) QueryApi(api string) error {
	switch api {
	case photon.ProviderName:
		return l.QueryPhoton()
	case osm.ProviderName:
		return l.QueryOSM()
	case opencage.ProviderName:
		return l.QueryOpenCage()
	}

	return l.Query()
}

func (l *Location) Query() error {
	if err := l.QueryPhoton(); err == nil {
		return nil
	}

	if opencage.ProviderKey == "" {
		return l.QueryOSM()
	}

	if err := l.QueryOpenCage(); err != nil {
		return l.QueryOSM()
	}

	return nil
}

func (l *Location) QueryOSM() error {
	s, err := osm.FindLocation(l.ID)

	if err != nil {
		return err
	}

	return l.Assign(s)
}

func (l *Location) QueryPhoton() error {
	s, err := photon.FindLocation(l.ID)

	if err != nil {
		return err
	}

	return l.Assign(s)
}

func (l *Location) QueryOpenCage() error {
	s, err := opencage.FindLocation(l.ID)

	if err != nil {
		return err
	}

	return l.Assign(s)
}

func (l *Location) Assign(s LocationSource) error {
	l.LocSource = s.Source()

	l.ID = s.CellID()

	if l.Unknown() {
		l.LocCategory = "unknown"
		return errors.New("maps: unknown location")
	}

	cat := s.Category()

	if ignoreCategories[cat] {
		l.LocName = ""
		l.LocCategory = ""
	} else {
		l.LocName = s.Name()
		l.LocCategory = s.Category()
	}

	l.LocCity = s.City()
	l.LocState = s.State()
	l.LocCountry = s.CountryCode()
	l.LocLabel = l.label()

	return nil
}

func (l *Location) Unknown() bool {
	return l.ID == ""
}

func (l *Location) label() string {
	if l.Unknown() {
		return "Unknown"
	}

	var countryName = l.CountryName()
	var loc []string

	shortCountry := len([]rune(countryName)) <= 20

	if l.LocCity != "" {
		loc = append(loc, l.LocCity)
	}

	if shortCountry && l.LocState != "" && l.LocCity != l.LocState {
		loc = append(loc, l.LocState)
	}

	if countryName != "" {
		loc = append(loc, countryName)
	}

	return strings.Join(loc[:], ", ")
}

func (l Location) Name() string {
	return l.LocName
}

func (l Location) Category() string {
	return l.LocCategory
}

func (l Location) Label() string {
	return l.LocLabel
}

func (l Location) City() string {
	return l.LocCity
}

func (l Location) State() string {
	return l.LocState
}

func (l Location) CountryCode() string {
	return l.LocCountry
}

func (l Location) CountryName() string {
	return CountryNames[l.LocCountry]
}

func (l Location) Source() string {
	return l.LocSource
}
