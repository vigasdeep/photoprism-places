package opencage

import (
	"strings"
)

func (l Location) Category() (result string) {
	key := strings.ToLower(strings.TrimSpace(l.LocCategory))

	if result, ok := ocTypes[key]; ok {
		return result
	}

	return ""
}
