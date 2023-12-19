package policyHandler

import (
	routesConsts "github.com/anaclaraddias/brick/adapter/http/constants"
	"github.com/anaclaraddias/brick/adapter/http/routes"
	"github.com/anaclaraddias/brick/core/domain/helper"
	policyDomain "github.com/anaclaraddias/brick/core/domain/policy"
	"github.com/anaclaraddias/brick/core/port"
	repositoryInterface "github.com/anaclaraddias/brick/core/port/repository"
	policyUsecase "github.com/anaclaraddias/brick/core/usecase/policy"
	policySharedmethod "github.com/anaclaraddias/brick/core/usecase/policy/sharedMethod"
	vehicleSharedMethod "github.com/anaclaraddias/brick/core/usecase/vehicle/sharedMethod"
	"github.com/anaclaraddias/brick/infra/database/repository"
	"github.com/anaclaraddias/brick/infra/requestEntity"
	"github.com/gin-gonic/gin"
)

type CreatePolicyVehicleHandler struct {
	connection      port.DatabaseConnectionInterface
	policyDatabase  repositoryInterface.PolicyRepositoryInterface
	vehicleDatabase repositoryInterface.VehicleRepositoryInterface
	uuid            port.UuidInterface
}

func NewCreatePolicyVehicleHandler(
	connection port.DatabaseConnectionInterface,
	uuid port.UuidInterface,
) port.HandlerInterface {
	return &CreatePolicyVehicleHandler{
		connection: connection,
		uuid:       uuid,
	}
}

func (createPolicyVehicleHandler *CreatePolicyVehicleHandler) Handle(context *gin.Context) {
	jsonResponse := routes.NewJsonResponse(context, createPolicyVehicleHandler.connection)

	decodedPolicyVehicle, err := createPolicyVehicleHandler.decodeRequest(context)

	if err != nil {
		jsonResponse.ThrowError(
			routesConsts.MessageKeyConst,
			err,
			routesConsts.BadRequestConst,
		)
		return
	}

	policyVehicle := createPolicyVehicleHandler.parseDataToDomain(decodedPolicyVehicle)

	if err != nil {
		jsonResponse.ThrowError(
			routesConsts.MessageKeyConst,
			err,
			routesConsts.BadRequestConst,
		)
		return
	}

	createPolicyVehicleHandler.openDatabaseConnection()

	policySharedmethod := policySharedmethod.NewPolicySharedMethod(
		createPolicyVehicleHandler.policyDatabase,
	)

	vehicleSharedMethod := vehicleSharedMethod.NewVehicleSharedMethod(
		createPolicyVehicleHandler.vehicleDatabase,
	)

	err = policyUsecase.NewCreatePolicyVehicle(
		createPolicyVehicleHandler.policyDatabase,
		createPolicyVehicleHandler.vehicleDatabase,
		policySharedmethod,
		vehicleSharedMethod,
		policyVehicle,
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
		helper.PostPolicyVehicleConst,
		routesConsts.StatusOk,
	)
}

func (createPolicyVehicleHandler *CreatePolicyVehicleHandler) decodeRequest(
	context *gin.Context,
) (*requestEntity.PostPolicyVehicleRequestEntity, error) {
	policy, err := requestEntity.DecodePolicyVehicleRequest(context.Request)

	if err != nil {
		return nil, err
	}

	if err := policy.Validate(); err != nil {
		return nil, err
	}

	return policy, nil
}

func (createPolicyVehicleHandler *CreatePolicyVehicleHandler) parseDataToDomain(
	decodedPolicyVehicle *requestEntity.PostPolicyVehicleRequestEntity,
) *policyDomain.PolicyVehicle {
	return policyDomain.NewPolicyVehicle(
		createPolicyVehicleHandler.uuid.GenerateUuid(),
		decodedPolicyVehicle.VehicleId,
		decodedPolicyVehicle.PolicyId,
	)
}

func (createPolicyVehicleHandler *CreatePolicyVehicleHandler) openDatabaseConnection() {
	createPolicyVehicleHandler.vehicleDatabase = repository.NewVehicleDatabase(
		createPolicyVehicleHandler.connection,
	)

	createPolicyVehicleHandler.policyDatabase = repository.NewPolicyDatabase(
		createPolicyVehicleHandler.connection,
	)
}
