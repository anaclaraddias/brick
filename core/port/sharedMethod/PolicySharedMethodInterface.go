package sharedMethodInterface

import policyDomain "github.com/anaclaraddias/brick/core/domain/policy"

type PolicySharedMethodInterface interface {
	VerifyIfVehicleIsInPolicy(vehicleId string) (bool, *policyDomain.PolicyVehicle, error)
	VerifyIfPolicyExists(policyId string) error
}
