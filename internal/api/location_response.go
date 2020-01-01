package api

import (
	"github.com/photoprism/photoprism-places/internal/entity"
)

type LocationResponse struct {
	ID       string        `json:"id"`
	Name     string        `json:"name"`
	Category string        `json:"category"`
	Suburb   string        `json:"suburb"`
	Timezone string        `json:"timezone"`
	Lat      float64       `json:"lat"`
	Lng      float64       `json:"lng"`
	Place    *entity.Place `json:"place"`
}

func NewLocationResponse(el *entity.Location) *LocationResponse {
	lat, lng := el.LatLng()

	l := &LocationResponse{
		ID:       el.ID,
		Name:     el.Name(),
		Category: el.Category(),
		Suburb:   el.Suburb(),
		Timezone: el.Timezone(),
		Lat:      lat,
		Lng:      lng,
		Place:    el.Place,
	}

	return l
}
