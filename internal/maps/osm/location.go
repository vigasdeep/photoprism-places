package osm

type Location struct {
	ID             string  `json:"-"`
	PlaceID        int     `json:"place_id"`
	LocName        string  `json:"name"`
	LocCategory    string  `json:"category"`
	LocType        string  `json:"type"`
	LocDisplayName string  `json:"display_name"`
	Address        Address `json:"address"`
}
