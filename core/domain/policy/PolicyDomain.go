package policyDomain

type Policy struct {
	Id            string
	Status        string
	StartDate     *string
	EndDate       *string
	CoverageLimit *float64
	Value         *float64
	UserId        string
}

func NewPolicy(
	id string,
	status string,
	startDate *string,
	endDate *string,
	coverageLimit *float64,
	value *float64,
	userId string,
) *Policy {
	return &Policy{
		Id:            id,
		Status:        status,
		StartDate:     startDate,
		EndDate:       endDate,
		CoverageLimit: coverageLimit,
		Value:         value,
		UserId:        userId,
	}
}

const (
	PolicyPendingStatusConst    = "pending"
	PolicyActiveStatusConst     = "active"
	PolicyCanceledStatusConst   = "canceled"
	PolicyRenovationStatusConst = "renovation"
)

var PolicyStatus = []string{
	PolicyPendingStatusConst,
	PolicyActiveStatusConst,
	PolicyCanceledStatusConst,
	PolicyRenovationStatusConst,
}
