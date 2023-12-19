package repositoryInterface

import vehicleDomain "github.com/anaclaraddias/brick/core/domain/vehicle"

type VehicleRepositoryInterface interface {
	CreateVehicle(vehicle *vehicleDomain.Vehicle) error
	FindVehicleByRenavamOrLicensePlate(renavam string, licensePlate string) (*vehicleDomain.Vehicle, error)
	FindVehicleById(vehicleId string) (*vehicleDomain.Vehicle, error)
	DeleteVehicle(vehicleId string) error
}
