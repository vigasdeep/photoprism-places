/*
This package encapsulates the OpenStreetMap API.

Additional information can be found in our Developer Guide:

https://github.com/photoprism/photoprism/wiki
*/
package osm

import (
	"github.com/photoprism/photoprism-places/internal/event"
)

var ProviderName = "osm"
var NominatimUrl = "https://nominatim.openstreetmap.org/"

var log = event.Log
