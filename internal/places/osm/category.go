package osm

import "fmt"

func (l Location) Category() (result string) {
	key := fmt.Sprintf("%s=%s", l.LocCategory, l.LocType)
	catKey := fmt.Sprintf("%s=*", l.LocCategory)
	typeKey := fmt.Sprintf("*=%s", l.LocType)

	if result, ok := categories[key]; ok {
		return result
	} else if result, ok := categories[catKey]; ok {
		return result
	} else if result, ok := categories[typeKey]; ok {
		return result
	}

	// log.Debugf("osm: no label found for %s", key)

	return ""
}
