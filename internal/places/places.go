/*
This package encapsulates the geo location APIs.

Additional information can be found in our Developer Guide:

https://github.com/photoprism/photoprism/wiki
*/
package places

import (
	"github.com/photoprism/photoprism-places/internal/event"
)

//go:generate go run gen.go

var log = event.Log
