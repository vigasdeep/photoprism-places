package osm

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/photoprism/photoprism-places/internal/places/client"
	"github.com/photoprism/photoprism/pkg/s2"
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
