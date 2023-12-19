package vehicleUsecase

import (
	repositoryInterface "github.com/anaclaraddias/brick/core/port/repository"
	sharedMethodInterface "github.com/anaclaraddias/brick/core/port/sharedMethod"
)

type DeleteVehicle struct {
	policyDatabase      repositoryInterface.PolicyRepositoryInterface
	vehicleDatabase     repositoryInterface.VehicleRepositoryInterface
	policySharedmethod  sharedMethodInterface.PolicySharedMethodInterface
	vehicleSharedmethod sharedMethodInterface.VehicleSharedMethodInterface
	vehicleId           string
}

func NewDeleteVehicle(
	policyDatabase repositoryInterface.PolicyRepositoryInterface,
	vehicleDatabase repositoryInterface.VehicleRepositoryInterface,
	policySharedmethod sharedMethodInterface.PolicySharedMethodInterface,
	vehicleSharedmethod sharedMethodInterface.VehicleSharedMethodInterface,
	vehicleId string,
) *DeleteVehicle {
	return &DeleteVehicle{
		policyDatabase:      policyDatabase,
		vehicleDatabase:     vehicleDatabase,
		policySharedmethod:  policySharedmethod,
		vehicleSharedmethod: vehicleSharedmethod,
		vehicleId:           vehicleId,
	}
}

func (deleteVehicle *DeleteVehicle) Execute() error {
	if err := deleteVehicle.vehicleSharedmethod.VerifyIfVehicleExists(
		deleteVehicle.vehicleId,
	); err != nil {
		return err
	}

	deleteLinkedPolicyVehicle,
		linkedPolicyVehicle,
		err := deleteVehicle.policySharedmethod.VerifyIfVehicleIsInPolicy(
		deleteVehicle.vehicleId,
	)

	if err != nil {
		return err
	}

	if deleteLinkedPolicyVehicle {
		deleteVehicle.policyDatabase.DeletePolicyVehicle(
			linkedPolicyVehicle.Id,
		)
	}

	if err := deleteVehicle.vehicleDatabase.DeleteVehicle(
		deleteVehicle.vehicleId,
	); err != nil {
		return err
	}

	return nil
}
