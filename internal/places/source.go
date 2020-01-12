package places

type LocationSource interface {
	CellID() string
	Category() string
	Name() string
	City() string
	State() string
	CountryCode() string
	Source() string
}

type StructSource struct {
	LocCellID string
	LocCountryCode string
	LocCategory string
	LocName string
	LocCity string
	LocState string
	LocSource string
}

func (s StructSource) CellID() string {
	return s.LocCellID
}

func (s StructSource) Category() string {
	return s.LocCategory
}

func (s StructSource) Name() string {
	return s.LocName
}

func (s StructSource) City() string {
	return s.LocCity
}

func (s StructSource) State() string {
	return s.LocState
}

func (s StructSource) CountryCode() string {
	return s.LocCountryCode
}

func (s StructSource) Source() string {
	return s.LocSource
}

