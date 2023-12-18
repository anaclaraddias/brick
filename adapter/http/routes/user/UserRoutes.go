package userRoutes

import (
	userHandler "github.com/anaclaraddias/brick/adapter/handler/user"
	routesConsts "github.com/anaclaraddias/brick/adapter/http/constants"
	"github.com/anaclaraddias/brick/core/port"
	"github.com/anaclaraddias/brick/infra/database"
	"github.com/anaclaraddias/brick/infra/utils"
	"github.com/gin-gonic/gin"
)

const (
	CreateUserConst = "CreateUser"
)

type UserRoutes struct {
	*gin.Engine
	userHandlers map[string]port.HandlerInterface
}

func NewUserRoutes(
	app *gin.Engine,
) *UserRoutes {
	return &UserRoutes{
		app,
		createUserHandlerMap(),
	}
}

func (UserRoutes *UserRoutes) Register() {
	UserRoutes.POST(
		routesConsts.PostUserConst,
		UserRoutes.userHandlers[CreateUserConst].Handle,
	)
}

func createUserHandlerMap() map[string]port.HandlerInterface {
	uuid := utils.NewUuid()
	connection := database.NewDatabaseConnection()

	return map[string]port.HandlerInterface{
		CreateUserConst: userHandler.NewCreateUserHandler(connection, uuid),
	}
}
