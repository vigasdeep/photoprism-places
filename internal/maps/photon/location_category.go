package photon

import (
	"fmt"
)


func (l Location) Category() (result string) {
	key := fmt.Sprintf("%s=%s", l.OsmKey, l.OsmValue)
	catKey := fmt.Sprintf("%s=*", l.OsmKey)
	typeKey := fmt.Sprintf("*=%s", l.OsmValue)

	if result, ok := categories[key]; ok {
		return result
	} else if result, ok := categories[catKey]; ok {
		return result
	} else if result, ok := categories[typeKey]; ok {
		return result
	}

	return ""
}
