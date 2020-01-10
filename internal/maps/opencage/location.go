package opencage

import (
	"strings"
)

type Location struct {
	ID          string
	LocName     string
	LocCategory string
	LocSuburb   string
	LocCity     string
	LocState    string
	LocCountry  string
}

func (l Location) CellID() string {
	return l.ID
}

func (l Location) Suburb() string {
	return strings.TrimSpace(l.LocSuburb)
}

func (l Location) City() string {
	return strings.TrimSpace(l.LocCity)
}

func (l Location) State() string {
	return strings.TrimSpace(l.LocState)
}

func (l Location) CountryCode() string {
	return strings.TrimSpace(l.LocCountry)
}

func (l Location) Source() string {
	return ProviderName
}
