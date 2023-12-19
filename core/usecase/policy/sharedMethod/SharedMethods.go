package policySharedmethod

import (
	"fmt"

	"github.com/anaclaraddias/brick/core/domain/helper"
	policyDomain "github.com/anaclaraddias/brick/core/domain/policy"
	repositoryInterface "github.com/anaclaraddias/brick/core/port/repository"
	sharedMethodInterface "github.com/anaclaraddias/brick/core/port/sharedMethod"
)

type PolicySharedMethod struct {
	policyDatabase repositoryInterface.PolicyRepositoryInterface
}

func NewPolicySharedMethod(
	policyDatabase repositoryInterface.PolicyRepositoryInterface,
) sharedMethodInterface.PolicySharedMethodInterface {
	return &PolicySharedMethod{
		policyDatabase: policyDatabase,
	}
}

func (policySharedMethod *PolicySharedMethod) VerifyIfVehicleIsInPolicy(
	vehicleId string,
) (bool, *policyDomain.PolicyVehicle, error) {
	linkedPolicyVehicle, err := policySharedMethod.policyDatabase.FindPolicyVehicleByVehicleId(
		vehicleId,
	)

	if err != nil {
		return false, nil, err
	}

	if linkedPolicyVehicle == nil {
		return false, nil, nil
	}

	return true, linkedPolicyVehicle, nil
}

func (policySharedMethod *PolicySharedMethod) VerifyIfPolicyExists(
	policyId string,
) error {
	policy, err := policySharedMethod.policyDatabase.FindPolicyById(
		policyId,
	)

	if err != nil {
		return err
	}

	if policy == nil {
		return fmt.Errorf(helper.PolicyNotFoundConst)
	}

	return nil
}
