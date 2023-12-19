package sharedMethodInterface

type VehicleSharedMethodInterface interface {
	VerifyIfVehicleExists(vehicleId string) error
}
