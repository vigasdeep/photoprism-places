package osm

import (
	"strings"

	"github.com/photoprism/photoprism/pkg/txt"
)

type Location struct {
	ID             string  `json:"-"`
	PlaceID        int     `json:"place_id"`
	LocName        string  `json:"name"`
	LocCategory    string  `json:"category"`
	LocType        string  `json:"type"`
	LocDisplayName string  `json:"display_name"`
	Address        Address `json:"address"`
}

func (l Location) CellID() (result string) {
	return l.ID
}

func (l Location) State() (result string) {
	result = l.Address.State

	return strings.TrimSpace(result)
}

func (l Location) City() (result string) {
	if l.Address.City != "" {
		result = l.Address.City
	} else if l.Address.Town != "" {
		result = l.Address.Town
	} else if l.Address.Village != "" {
		result = l.Address.Village
	} else if l.Address.County != "" {
		result = l.Address.County
	} else if l.Address.State != "" {
		result = l.Address.State
	}

	if len([]rune(result)) > 19 {
		result = ""
	}

	return strings.TrimSpace(result)
}

func (l Location) Suburb() (result string) {
	result = l.Address.Suburb

	return strings.TrimSpace(result)
}

func (l Location) CountryCode() (result string) {
	result = l.Address.CountryCode

	return strings.ToLower(strings.TrimSpace(result))
}

func (l Location) Keywords() (result []string) {
	return txt.Keywords(l.LocDisplayName)
}

func (l Location) Source() string {
	return ProviderName
}
