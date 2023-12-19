package vehicleHandler

import (
	"fmt"

	routesConsts "github.com/anaclaraddias/brick/adapter/http/constants"
	"github.com/anaclaraddias/brick/adapter/http/routes"
	"github.com/anaclaraddias/brick/core/domain/helper"
	"github.com/anaclaraddias/brick/core/port"
	repositoryInterface "github.com/anaclaraddias/brick/core/port/repository"
	policySharedmethod "github.com/anaclaraddias/brick/core/usecase/policy/sharedMethod"
	vehicleUsecase "github.com/anaclaraddias/brick/core/usecase/vehicle"
	vehicleSharedMethod "github.com/anaclaraddias/brick/core/usecase/vehicle/sharedMethod"
	"github.com/anaclaraddias/brick/infra/database/repository"
	"github.com/gin-gonic/gin"
)

const (
	VehicleIdFieldConst = "o veículo"
)

type DeleteVehicleHandler struct {
	connection      port.DatabaseConnectionInterface
	policyDatabase  repositoryInterface.PolicyRepositoryInterface
	vehicleDatabase repositoryInterface.VehicleRepositoryInterface
	uuid            port.UuidInterface
}

func NewDeleteVehicleHandler(
	connection port.DatabaseConnectionInterface,
	uuid port.UuidInterface,
) port.HandlerInterface {
	return &DeleteVehicleHandler{
		connection: connection,
		uuid:       uuid,
	}
}

func (deleteVehicleHandler *DeleteVehicleHandler) Handle(context *gin.Context) {
	jsonResponse := routes.NewJsonResponse(context)

	vehicleId, err := deleteVehicleHandler.verifyVehicleId(context)

	if err != nil {
		jsonResponse.ThrowError(
			routesConsts.MessageKeyConst,
			err,
			routesConsts.BadRequestConst,
		)
		return
	}

	deleteVehicleHandler.openDatabaseConnection()

	policySharedmethod := policySharedmethod.NewPolicySharedMethod(
		deleteVehicleHandler.policyDatabase,
	)

	vehicleSharedMethod := vehicleSharedMethod.NewVehicleSharedMethod(
		deleteVehicleHandler.vehicleDatabase,
	)

	err = vehicleUsecase.NewDeleteVehicle(
		deleteVehicleHandler.policyDatabase,
		deleteVehicleHandler.vehicleDatabase,
		policySharedmethod,
		vehicleSharedMethod,
		vehicleId,
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
		helper.DeleteVehicleConst,
		routesConsts.StatusOk,
	)
}

func (deleteVehicleHandler *DeleteVehicleHandler) verifyVehicleId(
	context *gin.Context,
) (string, error) {
	vehicleId := context.Query("vehicle-id")

	if vehicleId == "" {
		return "", fmt.Errorf(helper.InformFieldConst, VehicleIdFieldConst)
	}

	return vehicleId, nil
}

func (deleteVehicleHandler *DeleteVehicleHandler) openDatabaseConnection() {
	deleteVehicleHandler.vehicleDatabase = repository.NewVehicleDatabase(
		deleteVehicleHandler.connection,
	)

	deleteVehicleHandler.policyDatabase = repository.NewPolicyDatabase(
		deleteVehicleHandler.connection,
	)
}
