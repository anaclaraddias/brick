package repository

import (
	"database/sql"
	"time"

	vehicleDomain "github.com/anaclaraddias/brick/core/domain/vehicle"
	"github.com/anaclaraddias/brick/core/port"
	repositoryInterface "github.com/anaclaraddias/brick/core/port/repository"
	"github.com/anaclaraddias/brick/infra/models"
)

type VehicleDatabase struct {
	connection port.DatabaseConnectionInterface
}

func NewVehicleDatabase(
	connection port.DatabaseConnectionInterface,
) repositoryInterface.VehicleRepositoryInterface {
	connection.Open()

	return &VehicleDatabase{connection: connection}
}

func (vehicleDatabase *VehicleDatabase) CreateVehicle(
	vehicle *vehicleDomain.Vehicle,
) error {
	var dbVehicle *models.VehicleModel

	createdAt := time.Now()
	updatedAt := sql.NullTime{Valid: false}

	query := `INSERT INTO vehicle (
			id,
			brand,
			model,
			vehicle_year,
			license_plate,
			renavam,
			vehicle_value,
			cargo,
			vehicle_height,
			vehicle_width,
			vehicle_length,
			vehicle_type,
			update_date,
			creation_date
		) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?);`

	if err := vehicleDatabase.connection.Raw(
		query,
		&dbVehicle,
		vehicle.Id,
		vehicle.Brand,
		vehicle.Model,
		vehicle.Year,
		vehicle.LicensePlate,
		vehicle.Renavam,
		vehicle.Value,
		vehicle.Cargo,
		vehicle.Height,
		vehicle.Width,
		vehicle.Length,
		vehicle.Type,
		updatedAt,
		createdAt,
	); err != nil {
		return err
	}

	return nil
}

func (vehicleDatabase *VehicleDatabase) FindVehicleByRenavam(
	renavam string,
) ([]map[string]interface{}, error) {
	query := `SELECT * FROM vehicle WHERE renavam = ?;`

	dbVehicle, err := vehicleDatabase.connection.Rows(
		query,
		renavam,
	)

	if err != nil {
		return nil, err
	}

	return dbVehicle, nil
}

func (vehicleDatabase *VehicleDatabase) FindVehicleById(
	vehicleId string,
) (*vehicleDomain.Vehicle, error) {
	var dbVehicle *models.VehicleModel

	query := `SELECT * FROM vehicle WHERE id = ?;`

	if err := vehicleDatabase.connection.Raw(
		query,
		&dbVehicle,
		vehicleId,
	); err != nil {
		return nil, err
	}

	if dbVehicle == nil {
		return nil, nil
	}

	vehicle := vehicleDomain.NewVehicle(
		dbVehicle.Id,
		dbVehicle.Brand,
		dbVehicle.Model,
		dbVehicle.Year,
		dbVehicle.Renavam,
		dbVehicle.LicensePlate,
		dbVehicle.Value,
		dbVehicle.Cargo,
		dbVehicle.Height,
		dbVehicle.Width,
		dbVehicle.Length,
		dbVehicle.Type,
	)

	return vehicle, nil
}
