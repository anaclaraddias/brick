package policyDomain

type PolicyCoverage struct {
	Id         string
	CoverageId string
	PolicyId   string
}

func NewPolicyCoverage(
	id string,
	coverageId string,
	policyId string,
) *PolicyCoverage {
	return &PolicyCoverage{
		Id:         id,
		CoverageId: coverageId,
		PolicyId:   policyId,
	}
}
