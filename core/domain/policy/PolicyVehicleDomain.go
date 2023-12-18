package policyDomain

type PolicyVehicle struct {
	Id        string
	VehicleId string
	PolicyId  string
}

func NewPolicyVehicle(
	id string,
	vehicleId string,
	policyId string,
) *PolicyVehicle {
	return &PolicyVehicle{
		Id:        id,
		VehicleId: vehicleId,
		PolicyId:  policyId,
	}
}
