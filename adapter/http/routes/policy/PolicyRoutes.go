package policyRoutes

import (
	policyHandler "github.com/anaclaraddias/brick/adapter/handler/policy"
	routesConsts "github.com/anaclaraddias/brick/adapter/http/constants"
	"github.com/anaclaraddias/brick/core/port"
	"github.com/anaclaraddias/brick/infra/database"
	"github.com/anaclaraddias/brick/infra/utils"
	"github.com/gin-gonic/gin"
)

const (
	CreatePolicyConst         = "CreatePolicy"
	CreatePolicyVehicleConst  = "CreatePolicyVehicle"
	CreatePolicyCoverageConst = "CreatePolicyCoverage"
)

type PolicyRoutes struct {
	*gin.Engine
	policyHandlers map[string]port.HandlerInterface
}

func NewPolicyRoutes(
	app *gin.Engine,
) *PolicyRoutes {
	return &PolicyRoutes{
		app,
		createPolicyHandlerMap(),
	}
}

func (policyRoutes *PolicyRoutes) Register() {
	policyRoutes.POST(
		routesConsts.PostPolicyConst,
		policyRoutes.policyHandlers[CreatePolicyConst].Handle,
	)

	policyRoutes.POST(
		routesConsts.PostPolicyVehicleConst,
		policyRoutes.policyHandlers[CreatePolicyVehicleConst].Handle,
	)

	policyRoutes.POST(
		routesConsts.PostPolicyCoverageConst,
		policyRoutes.policyHandlers[CreatePolicyCoverageConst].Handle,
	)
}

func createPolicyHandlerMap() map[string]port.HandlerInterface {
	uuid := utils.NewUuid()
	connection := database.NewDatabaseConnection()

	return map[string]port.HandlerInterface{
		CreatePolicyConst:         policyHandler.NewCreatePolicyHandler(connection, uuid),
		CreatePolicyVehicleConst:  policyHandler.NewCreatePolicyVehicleHandler(connection, uuid),
		CreatePolicyCoverageConst: policyHandler.NewCreatePolicyCoverageHandler(connection, uuid),
	}
}
