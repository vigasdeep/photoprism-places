package api

import (
	"github.com/photoprism/photoprism-places/internal/entity"
	"gopkg.in/ugjka/go-tz.v2/tz"
)

type LocationResponse struct {
	ID       string        `json:"id"`
	Name     string        `json:"name"`
	Category string        `json:"category"`
	Timezone string        `json:"timezone"`
	Lat      float64       `json:"lat"`
	Lng      float64       `json:"lng"`
	Place    *entity.Place `json:"place"`
	Licence  string        `json:"licence"`
}

// Timezone returns the location time zone as string.
func Timezone(lat, lng float64) string {
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

// NewLocationResponse returns an initialized LocationResponse.
func NewLocationResponse(el *entity.Location) *LocationResponse {
	lat, lng := el.LatLng()

	l := &LocationResponse{
		ID:       el.ID,
		Name:     el.Name(),
		Category: el.Category(),
		Timezone: Timezone(lat, lng),
		Lat:      lat,
		Lng:      lng,
		Place:    el.Place,
		Licence:  "Data © OpenStreetMap contributors, ODbL 1.0, see https://osm.org/copyright",
	}

	return l
}
