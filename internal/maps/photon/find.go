package photon

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/photoprism/photoprism-places/internal/maps/client"
	"github.com/photoprism/photoprism-places/internal/s2"

	"github.com/tidwall/gjson"
)

// API docs see https://github.com/komoot/photon
func FindLocation(id string) (result Location, err error) {
	if ProviderUrl == "" {
		return result, errors.New("photon: no provider url")
	}

	if len(id) > 16 || len(id) == 0 {
		return result, errors.New("photon: invalid location id")
	}

	lat, lng := s2.LatLng(id)

	if lat == 0.0 || lng == 0.0 {
		return result, fmt.Errorf("photon: skipping lat %f, lng %f", lat, lng)
	}

	url := fmt.Sprintf("%sreverse?lat=%f&lon=%f", ProviderUrl, lat, lng)

	log.Debugf("photon: query %s", url)

	res, err := client.Request(url)

	if err != nil {
		log.Errorf("photon: %s", err.Error())
		return result, err
	}

	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		log.Errorf("photon: %s", readErr.Error())
		return result, readErr
	}

	log.Debug(string(body))

	j := gjson.Parse(string(body))

	if !j.Get("features.0").Exists() {
		result.ID = ""

		return result, fmt.Errorf("photon: no result for %s", id)
	}

	result.ID = id
	result.LocCity = j.Get("features.0.properties.city").String()
	result.LocState = j.Get("features.0.properties.state").String()

	result.OsmKey = j.Get("features.0.properties.osm_key").String()
	result.OsmValue = j.Get("features.0.properties.osm_value").String()

	result.LocName = j.Get("features.0.properties.name").String()

	result.LocCountry = Countries[strings.ToLower(j.Get("features.0.properties.country").String())]

	return result, nil
}
