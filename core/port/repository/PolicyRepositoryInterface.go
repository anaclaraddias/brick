package repositoryInterface

import policyDomain "github.com/anaclaraddias/brick/core/domain/policy"

type PolicyRepositoryInterface interface {
	CreatePolicy(policy *policyDomain.Policy) error
	FindPolicyById(policyId string) (*policyDomain.Policy, error)
	CreatePolicyVehicle(policyVehicle *policyDomain.PolicyVehicle) error
	CreatePolicyCoverage(policyCoverage *policyDomain.PolicyCoverage) error
	FindPolicyVehicleByVehicleId(vehicleId string) (*policyDomain.PolicyVehicle, error)
	DeletePolicyVehicle(policyVehicleId string) error
}
