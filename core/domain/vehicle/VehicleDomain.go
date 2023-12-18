package vehicleDomain

type Vehicle struct {
	Id           string
	Brand        string
	Model        string
	Year         string
	Renavam      string
	LicensePlate string
	Value        float64
	Cargo        float64
	Height       float64
	Width        float64
	Length       float64
	Type         string
}

func NewVehicle(
	Id string,
	Brand string,
	Model string,
	Year string,
	Renavam string,
	LicensePlate string,
	Value float64,
	Cargo float64,
	Height float64,
	Width float64,
	Length float64,
	Type string,
) *Vehicle {
	return &Vehicle{
		Id:           Id,
		Brand:        Brand,
		Model:        Model,
		Year:         Year,
		Renavam:      Renavam,
		LicensePlate: LicensePlate,
		Value:        Value,
		Cargo:        Cargo,
		Height:       Height,
		Width:        Width,
		Length:       Length,
		Type:         Type,
	}
}

const (
	VehiclePersonalTypeConst = "personal"
	VehicleBusinessTypeConst = "business"
)

var VehicleTypes = []string{
	VehiclePersonalTypeConst,
	VehicleBusinessTypeConst,
}
