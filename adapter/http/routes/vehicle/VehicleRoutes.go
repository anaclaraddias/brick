package vehicleRoutes

import (
	vehicleHandler "github.com/anaclaraddias/brick/adapter/handler/vehicle"
	routesConsts "github.com/anaclaraddias/brick/adapter/http/constants"
	"github.com/anaclaraddias/brick/core/port"
	"github.com/anaclaraddias/brick/infra/database"
	"github.com/anaclaraddias/brick/infra/utils"
	"github.com/gin-gonic/gin"
)

const (
	CreateVehicleConst string = "CreateVehicle"
)

type VehicleRoutes struct {
	*gin.Engine
	vehicleHandlers map[string]port.HandlerInterface
}

func NewVehicleRoutes(
	app *gin.Engine,
) *VehicleRoutes {
	return &VehicleRoutes{
		app,
		createVehicleHandlerMap(),
	}
}

func (vehicleRoutes *VehicleRoutes) Register() {
	vehicleRoutes.POST(
		routesConsts.PostVehicleConst,
		vehicleRoutes.vehicleHandlers[CreateVehicleConst].Handle,
	)
}

func createVehicleHandlerMap() map[string]port.HandlerInterface {
	uuid := utils.NewUuid()
	connection := database.NewDatabaseConnection()

	return map[string]port.HandlerInterface{
		CreateVehicleConst: vehicleHandler.NewCreateVehicleHandler(connection, uuid),
	}
}
