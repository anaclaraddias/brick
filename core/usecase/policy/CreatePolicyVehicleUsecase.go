package policyUsecase

import (
	"fmt"

	"github.com/anaclaraddias/brick/core/domain/helper"
	policyDomain "github.com/anaclaraddias/brick/core/domain/policy"
	repositoryInterface "github.com/anaclaraddias/brick/core/port/repository"
	sharedMethodInterface "github.com/anaclaraddias/brick/core/port/sharedMethod"
)

type CreatePolicyVehicle struct {
	policyDatabase      repositoryInterface.PolicyRepositoryInterface
	vehicleDatabase     repositoryInterface.VehicleRepositoryInterface
	policySharedmethod  sharedMethodInterface.PolicySharedMethodInterface
	vehicleSharedmethod sharedMethodInterface.VehicleSharedMethodInterface
	policyVehicle       *policyDomain.PolicyVehicle
}

func NewCreatePolicyVehicle(
	policyDatabase repositoryInterface.PolicyRepositoryInterface,
	vehicleDatabase repositoryInterface.VehicleRepositoryInterface,
	policySharedmethod sharedMethodInterface.PolicySharedMethodInterface,
	vehicleSharedmethod sharedMethodInterface.VehicleSharedMethodInterface,
	policyVehicle *policyDomain.PolicyVehicle,
) *CreatePolicyVehicle {
	return &CreatePolicyVehicle{
		policyDatabase:      policyDatabase,
		vehicleDatabase:     vehicleDatabase,
		policySharedmethod:  policySharedmethod,
		vehicleSharedmethod: vehicleSharedmethod,
		policyVehicle:       policyVehicle,
	}
}

func (createPolicyVehicle *CreatePolicyVehicle) Execute() error {
	if err := createPolicyVehicle.policySharedmethod.VerifyIfPolicyExists(
		createPolicyVehicle.policyVehicle.PolicyId,
	); err != nil {
		return err
	}

	if err := createPolicyVehicle.vehicleSharedmethod.VerifyIfVehicleExists(
		createPolicyVehicle.policyVehicle.VehicleId,
	); err != nil {
		return err
	}

	isVehicleInPolicy, _, err := createPolicyVehicle.policySharedmethod.VerifyIfVehicleIsInPolicy(
		createPolicyVehicle.policyVehicle.VehicleId,
	)

	if err != nil {
		return err
	}

	if isVehicleInPolicy {
		return fmt.Errorf(helper.VehicleAlreadyIsInPolicyConst)
	}

	if err := createPolicyVehicle.policyDatabase.CreatePolicyVehicle(
		createPolicyVehicle.policyVehicle,
	); err != nil {
		return err
	}

	return nil
}
