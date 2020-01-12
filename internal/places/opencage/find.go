package opencage

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/photoprism/photoprism-places/internal/places/client"
	"github.com/photoprism/photoprism/pkg/s2"

	"github.com/tidwall/gjson"
)

// API docs see https://opencagedata.com/api#reverse-resp
func FindLocation(id string) (result Location, err error) {
	if ProviderKey == "" {
		return result, errors.New("opencage: no provider key")
	}

	if len(id) > 16 || len(id) == 0 {
		return result, errors.New("opencage: invalid location id")
	}

	lat, lng := s2.LatLng(id)

	if lat == 0.0 || lng == 0.0 {
		return result, fmt.Errorf("opencage: skipping lat %f, lng %f", lat, lng)
	}

	url := fmt.Sprintf("%sgeocode/v1/json?key=%s&q=%f+%f&pretty=0&no_annotations=1", ProviderUrl, ProviderKey, lat, lng)

	log.Debugf("opencage: query %s", url)

	res, err := client.Request(url)

	if err != nil {
		log.Errorf("opencage: %s", err.Error())
		return result, err
	}

	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		log.Errorf("opencage: %s", readErr.Error())
		return result, readErr
	}

	log.Debug(string(body))

	j := gjson.Parse(string(body))

	if j.Get("total_results").Int() < 1 {
		result.ID = ""

		return result, fmt.Errorf("opencage: no result for %s", id)
	}

	result.ID = id
	result.LocCity = j.Get("results.0.components.city").String()
	result.LocState = j.Get("results.0.components.state").String()
	result.LocCategory = j.Get("results.0.components._type").String()

	if n := j.Get("results.0.components.unknown").String(); n != "" {
		result.LocName = n
	} else if n := j.Get("results.0.components." + result.LocCategory).String(); n != "" {
		result.LocName = n
	}

	result.LocCountry = j.Get("results.0.components.country_code").String()

	return result, nil
}
