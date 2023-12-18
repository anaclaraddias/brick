package policyUsecase

import (
	"fmt"

	"github.com/anaclaraddias/brick/core/domain/helper"
	policyDomain "github.com/anaclaraddias/brick/core/domain/policy"
	repositoryInterface "github.com/anaclaraddias/brick/core/port/repository"
)

type CreatePolicyVehicle struct {
	policyDatabase  repositoryInterface.PolicyRepositoryInterface
	vehicleDatabase repositoryInterface.VehicleRepositoryInterface
	policyVehicle   *policyDomain.PolicyVehicle
}

func NewCreatePolicyVehicle(
	policyDatabase repositoryInterface.PolicyRepositoryInterface,
	vehicleDatabase repositoryInterface.VehicleRepositoryInterface,
	policyVehicle *policyDomain.PolicyVehicle,
) *CreatePolicyVehicle {
	return &CreatePolicyVehicle{
		policyDatabase:  policyDatabase,
		vehicleDatabase: vehicleDatabase,
		policyVehicle:   policyVehicle,
	}
}

func (createPolicyVehicle *CreatePolicyVehicle) Execute() error {
	if err := createPolicyVehicle.verifyIfPolicyExists(); err != nil {
		return err
	}

	if err := createPolicyVehicle.verifyIfVehicleExists(); err != nil {
		return err
	}

	if err := createPolicyVehicle.policyDatabase.CreatePolicyVehicle(
		createPolicyVehicle.policyVehicle,
	); err != nil {
		return err
	}

	return nil
}

func (createPolicyVehicle *CreatePolicyVehicle) verifyIfPolicyExists() error {
	policy, err := createPolicyVehicle.policyDatabase.FindPolicyById(
		createPolicyVehicle.policyVehicle.PolicyId,
	)

	if err != nil {
		return err
	}

	if policy == nil {
		return fmt.Errorf(helper.PolicyNotFoundConst)
	}

	return nil
}

func (createPolicyVehicle *CreatePolicyVehicle) verifyIfVehicleExists() error {
	vehicle, err := createPolicyVehicle.vehicleDatabase.FindVehicleById(
		createPolicyVehicle.policyVehicle.VehicleId,
	)

	if err != nil {
		return err
	}

	if vehicle == nil {
		return fmt.Errorf(helper.VehicleNotFoundConst)
	}

	return nil
}
