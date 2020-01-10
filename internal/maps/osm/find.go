package osm

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/photoprism/photoprism-places/internal/maps/client"
	"github.com/photoprism/photoprism-places/internal/s2"
	"github.com/photoprism/photoprism-places/internal/util"
)

// API docs see https://wiki.openstreetmap.org/wiki/Nominatim#Reverse_Geocoding
func FindLocation(id string) (result Location, err error) {
	if NominatimUrl == "" {
		return result, errors.New("osm: no provider url")
	}

	if len(id) > 16 || len(id) == 0 {
		return result, errors.New("osm: invalid location id")
	}

	lat, lng := s2.LatLng(id)

	if lat == 0.0 || lng == 0.0 {
		return result, fmt.Errorf("osm: skipping lat %f, lng %f", lat, lng)
	}

	url := fmt.Sprintf("%sreverse?lat=%f&lon=%f&format=jsonv2&accept-language=en&zoom=18", NominatimUrl, lat, lng)

	log.Debugf("osm: query %s", url)

	res, err := client.Request(url)

	if err != nil {
		log.Errorf("osm: %s", err.Error())
		return result, err
	}

	err = json.NewDecoder(res.Body).Decode(&result)

	if err != nil {
		log.Errorf("osm: %s", err.Error())
		return result, err
	}

	if result.PlaceID == 0 {
		result.ID = ""

		return result, fmt.Errorf("osm: no result for %s", id)
	}

	result.ID = id

	return result, nil
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
	return util.Keywords(l.LocDisplayName)
}

func (l Location) Source() string {
	return ProviderName
}
