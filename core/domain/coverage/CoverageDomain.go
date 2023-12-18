package coverageDomain

type Coverage struct {
	Id          string
	Name        string
	Description string
	RateValue   float64
}

func NewCoverage(
	id string,
	name string,
	description string,
	rateValue float64,
) *Coverage {
	return &Coverage{
		Id:          id,
		Name:        name,
		Description: description,
		RateValue:   rateValue,
	}
}
