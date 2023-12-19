package vehicleSharedMethod

import (
	"fmt"

	"github.com/anaclaraddias/brick/core/domain/helper"
	repositoryInterface "github.com/anaclaraddias/brick/core/port/repository"
	sharedMethodInterface "github.com/anaclaraddias/brick/core/port/sharedMethod"
)

type VehicleSharedMethod struct {
	vehicleDatabase repositoryInterface.VehicleRepositoryInterface
}

func NewVehicleSharedMethod(
	vehicleDatabase repositoryInterface.VehicleRepositoryInterface,
) sharedMethodInterface.VehicleSharedMethodInterface {
	return &VehicleSharedMethod{
		vehicleDatabase: vehicleDatabase,
	}
}

func (vehicleSharedMethod *VehicleSharedMethod) VerifyIfVehicleExists(
	vehicleId string,
) error {
	vehicle, err := vehicleSharedMethod.vehicleDatabase.FindVehicleById(
		vehicleId,
	)

	if err != nil {
		return err
	}

	if vehicle == nil {
		return fmt.Errorf(helper.VehicleNotFoundConst)
	}

	return nil
}
