package vehicleUsecase

import (
	"fmt"

	"github.com/anaclaraddias/brick/core/domain/helper"
	vehicleDomain "github.com/anaclaraddias/brick/core/domain/vehicle"
	repositoryInterface "github.com/anaclaraddias/brick/core/port/repository"
)

type CreateVehicle struct {
	vehicle         *vehicleDomain.Vehicle
	vehicleDatabase repositoryInterface.VehicleRepositoryInterface
}

func NewCreateVehicle(
	vehicle *vehicleDomain.Vehicle,
	vehicleDatabase repositoryInterface.VehicleRepositoryInterface,
) *CreateVehicle {
	return &CreateVehicle{
		vehicle:         vehicle,
		vehicleDatabase: vehicleDatabase,
	}
}

func (createVehicle *CreateVehicle) Execute() error {
	// if err := createVehicle.verifyIfVehicleAlreadyExists(); err != nil {
	// 	return err
	// }

	if err := createVehicle.vehicleDatabase.CreateVehicle(createVehicle.vehicle); err != nil {
		return err
	}

	return nil
}

func (createVehicle *CreateVehicle) verifyIfVehicleAlreadyExists() error {
	vehicle, err := createVehicle.vehicleDatabase.FindVehicleByRenavam(
		createVehicle.vehicle.Renavam,
	)

	if err != nil {
		return err
	}

	if vehicle != nil {
		return fmt.Errorf(helper.VehicleAlreadyExistsConst)
	}

	return nil
}
