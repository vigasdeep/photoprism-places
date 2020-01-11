package photon

import (
	"strings"
)

type Location struct {
	ID         string
	LocName    string
	LocCity    string
	LocState   string
	LocCountry string
	OsmKey     string
	OsmValue   string
}

func (l Location) CellID() string {
	return l.ID
}

func (l Location) City() string {
	result := strings.TrimSpace(l.LocCity)
	lower := strings.ToLower(result)
	parts := strings.Split(lower, " ")

	if len(parts) < 3 {
		return result
	}

	last := len(parts) - 1

	if parts[last] != "municipality" {
		return result
	}

	if municipalityBlacklist[parts[0]] {
		return ""
	}

	if parts[last - 1] == "local" {
		i := strings.Index(lower, "local municipality") - 1
		return result[:i]
	}

	i := strings.Index(lower, "municipality") - 1

	return result[:i]
}

func (l Location) State() string {
	return strings.TrimSpace(l.LocState)
}

func (l Location) CountryCode() string {
	return strings.TrimSpace(l.LocCountry)
}

func (l Location) Source() string {
	return ProviderName
}
