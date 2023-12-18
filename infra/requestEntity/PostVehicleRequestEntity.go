package requestEntity

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/anaclaraddias/brick/core/domain/helper"
	vehicleDomain "github.com/anaclaraddias/brick/core/domain/vehicle"
)

const (
	BrandFieldConst        = "a marca"
	ModelFieldConst        = "o modelo"
	YearFieldConst         = "o ano"
	LicensePlateFieldConst = "a placa"
	RenavamFieldConst      = "o renavam"
	ValueFieldConst        = "o valor"
	CargoFieldConst        = "a carga máxima"
	HeightFieldConst       = "a altura"
	WidthFieldConst        = "a largura"
	LengthFieldConst       = "o comprimento"
	TypeFieldConst         = "o tipo"
)

type PostVehicleRequestEntity struct {
	Brand        string  `json:"brand"`
	Model        string  `json:"model"`
	Year         string  `json:"year"`
	LicensePlate string  `json:"license_plate"`
	Renavam      string  `json:"renavam"`
	Value        float64 `json:"value"`
	Cargo        float64 `json:"cargo"`
	Height       float64 `json:"height"`
	Width        float64 `json:"width"`
	Length       float64 `json:"length"`
	Type         string  `json:"type"`
}

func DecodeVehicleRequest(request *http.Request) (*PostVehicleRequestEntity, error) {
	var vehicle *PostVehicleRequestEntity

	if err := json.NewDecoder(request.Body).Decode(&vehicle); err != nil {
		return nil, err
	}

	return vehicle, nil
}

func (vehicle *PostVehicleRequestEntity) Validate() error {
	if vehicle.Brand == "" {
		return fmt.Errorf(helper.VehicleFieldCanNotBeEmptyConst, BrandFieldConst)
	}

	if vehicle.Model == "" {
		return fmt.Errorf(helper.VehicleFieldCanNotBeEmptyConst, ModelFieldConst)
	}

	if vehicle.Year == "" {
		return fmt.Errorf(helper.VehicleFieldCanNotBeEmptyConst, YearFieldConst)
	}

	if vehicle.LicensePlate == "" {
		return fmt.Errorf(helper.VehicleFieldCanNotBeEmptyConst, LicensePlateFieldConst)
	}

	if vehicle.Renavam == "" {
		return fmt.Errorf(helper.VehicleFieldCanNotBeEmptyConst, RenavamFieldConst)
	}

	if vehicle.Value == 0 {
		return fmt.Errorf(helper.VehicleFieldCanNotBeEmptyConst, ValueFieldConst)
	}

	if vehicle.Cargo == 0 {
		return fmt.Errorf(helper.VehicleFieldCanNotBeEmptyConst, CargoFieldConst)
	}

	if vehicle.Height == 0 {
		return fmt.Errorf(helper.VehicleFieldCanNotBeEmptyConst, HeightFieldConst)
	}

	if vehicle.Width == 0 {
		return fmt.Errorf(helper.VehicleFieldCanNotBeEmptyConst, WidthFieldConst)
	}

	if vehicle.Length == 0 {
		return fmt.Errorf(helper.VehicleFieldCanNotBeEmptyConst, LengthFieldConst)
	}

	if vehicle.Type == "" {
		return fmt.Errorf(helper.VehicleFieldCanNotBeEmptyConst, TypeFieldConst)
	}

	if !slices.Contains(vehicleDomain.VehicleTypes, vehicle.Type) {
		return fmt.Errorf(helper.VehicleTypeIsNotInTheEnumConst)
	}

	return nil
}
