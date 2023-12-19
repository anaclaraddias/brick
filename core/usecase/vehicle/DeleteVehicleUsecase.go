package vehicleUsecase

import (
	"fmt"

	"github.com/anaclaraddias/brick/core/domain/helper"
	policyDomain "github.com/anaclaraddias/brick/core/domain/policy"
	repositoryInterface "github.com/anaclaraddias/brick/core/port/repository"
)

type DeleteVehicle struct {
	policyDatabase      repositoryInterface.PolicyRepositoryInterface
	vehicleDatabase     repositoryInterface.VehicleRepositoryInterface
	linkedPolicyVehicle *policyDomain.PolicyVehicle
	vehicleId           string
}

func NewDeleteVehicle(
	policyDatabase repositoryInterface.PolicyRepositoryInterface,
	vehicleDatabase repositoryInterface.VehicleRepositoryInterface,
	vehicleId string,
) *DeleteVehicle {
	return &DeleteVehicle{
		policyDatabase:  policyDatabase,
		vehicleDatabase: vehicleDatabase,
		vehicleId:       vehicleId,
	}
}

func (deleteVehicle *DeleteVehicle) Execute() error {
	if err := deleteVehicle.verifyIfVehicleExists(); err != nil {
		return err
	}

	deleteLinkedPolicyVehicle, err := deleteVehicle.verifyIfVehicleIsInPolicy()

	if err != nil {
		return err
	}

	if deleteLinkedPolicyVehicle {
		deleteVehicle.policyDatabase.DeletePolicyVehicle(
			deleteVehicle.linkedPolicyVehicle.Id,
		)
	}

	if err := deleteVehicle.vehicleDatabase.DeleteVehicle(
		deleteVehicle.vehicleId,
	); err != nil {
		return err
	}

	return nil
}

func (deleteVehicle *DeleteVehicle) verifyIfVehicleExists() error {
	vehicle, err := deleteVehicle.vehicleDatabase.FindVehicleById(
		deleteVehicle.vehicleId,
	)

	if err != nil {
		return err
	}

	if vehicle == nil {
		return fmt.Errorf(helper.VehicleNotFoundConst)
	}

	return nil
}

func (deleteVehicle *DeleteVehicle) verifyIfVehicleIsInPolicy() (bool, error) {
	linkedPolicyVehicle, err := deleteVehicle.policyDatabase.FindPolicyVehicleByVehicleId(
		deleteVehicle.vehicleId,
	)

	if err != nil {
		return false, err
	}

	if linkedPolicyVehicle == nil {
		return false, nil
	}

	deleteVehicle.linkedPolicyVehicle = linkedPolicyVehicle

	return true, nil
}
