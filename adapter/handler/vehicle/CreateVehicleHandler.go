package vehicleHandler

import (
	routesConsts "github.com/anaclaraddias/brick/adapter/http/constants"
	"github.com/anaclaraddias/brick/adapter/http/routes"
	"github.com/anaclaraddias/brick/core/domain/helper"
	vehicleDomain "github.com/anaclaraddias/brick/core/domain/vehicle"
	"github.com/anaclaraddias/brick/core/port"
	repositoryInterface "github.com/anaclaraddias/brick/core/port/repository"
	vehicleUsecase "github.com/anaclaraddias/brick/core/usecase/vehicle"
	"github.com/anaclaraddias/brick/infra/database/repository"
	"github.com/anaclaraddias/brick/infra/requestEntity"
	"github.com/gin-gonic/gin"
)

type CreateVehicleHandler struct {
	connection      port.DatabaseConnectionInterface
	vehicleDatabase repositoryInterface.VehicleRepositoryInterface
	uuid            port.UuidInterface
}

func NewCreateVehicleHandler(
	connection port.DatabaseConnectionInterface,
	uuid port.UuidInterface,
) port.HandlerInterface {
	return &CreateVehicleHandler{
		connection: connection,
		uuid:       uuid,
	}
}

func (createVehicleHandler *CreateVehicleHandler) Handle(context *gin.Context) {
	jsonResponse := routes.NewJsonResponse(context)

	decodedVehicle, err := createVehicleHandler.decodeRequest(context)

	if err != nil {
		jsonResponse.ThrowError(
			routesConsts.MessageKeyConst,
			err,
			routesConsts.BadRequestConst,
		)
		return
	}

	vehicle := createVehicleHandler.parseDataToDomain(decodedVehicle)

	if err != nil {
		jsonResponse.ThrowError(
			routesConsts.MessageKeyConst,
			err,
			routesConsts.BadRequestConst,
		)
		return
	}

	createVehicleHandler.openDatabaseConnection()

	err = vehicleUsecase.NewCreateVehicle(
		vehicle,
		createVehicleHandler.vehicleDatabase,
	).Execute()

	if err != nil {
		jsonResponse.ThrowError(
			routesConsts.MessageKeyConst,
			err,
			routesConsts.BadRequestConst,
		)
		return
	}

	jsonResponse.SendJson(
		routesConsts.MessageKeyConst,
		helper.PostVehicleConst,
		routesConsts.StatusOk,
	)
}

func (createVehicleHandler *CreateVehicleHandler) decodeRequest(
	context *gin.Context,
) (*requestEntity.PostVehicleRequestEntity, error) {
	vehicle, err := requestEntity.DecodeVehicleRequest(context.Request)

	if err != nil {
		return nil, err
	}

	if err := vehicle.Validate(); err != nil {
		return nil, err
	}

	return vehicle, nil
}

func (createVehicleHandler *CreateVehicleHandler) parseDataToDomain(
	decodedVehicle *requestEntity.PostVehicleRequestEntity,
) *vehicleDomain.Vehicle {
	return vehicleDomain.NewVehicle(
		createVehicleHandler.uuid.GenerateUuid(),
		decodedVehicle.Brand,
		decodedVehicle.Model,
		decodedVehicle.Year,
		decodedVehicle.Renavam,
		decodedVehicle.LicensePlate,
		decodedVehicle.Value,
		decodedVehicle.Cargo,
		decodedVehicle.Height,
		decodedVehicle.Width,
		decodedVehicle.Length,
		decodedVehicle.Type,
	)
}

func (createVehicleHandler *CreateVehicleHandler) openDatabaseConnection() {
	createVehicleHandler.vehicleDatabase = repository.NewVehicleDatabase(createVehicleHandler.connection)
}
