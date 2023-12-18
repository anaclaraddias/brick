package policyHandler

import (
	routesConsts "github.com/anaclaraddias/brick/adapter/http/constants"
	"github.com/anaclaraddias/brick/adapter/http/routes"
	"github.com/anaclaraddias/brick/core/domain/helper"
	policyDomain "github.com/anaclaraddias/brick/core/domain/policy"
	"github.com/anaclaraddias/brick/core/port"
	repositoryInterface "github.com/anaclaraddias/brick/core/port/repository"
	policyUsecase "github.com/anaclaraddias/brick/core/usecase/policy"
	"github.com/anaclaraddias/brick/infra/database/repository"
	"github.com/anaclaraddias/brick/infra/requestEntity"
	"github.com/gin-gonic/gin"
)

type CreatePolicyHandler struct {
	connection     port.DatabaseConnectionInterface
	policyDatabase repositoryInterface.PolicyRepositoryInterface
	uuid           port.UuidInterface
}

func NewCreatePolicyHandler(
	connection port.DatabaseConnectionInterface,
	uuid port.UuidInterface,
) port.HandlerInterface {
	return &CreatePolicyHandler{
		connection: connection,
		uuid:       uuid,
	}
}

func (createPolicyHandler *CreatePolicyHandler) Handle(context *gin.Context) {
	jsonResponse := routes.NewJsonResponse(context)

	decodedPolicy, err := createPolicyHandler.decodeRequest(context)

	if err != nil {
		jsonResponse.ThrowError(
			routesConsts.MessageKeyConst,
			err,
			routesConsts.BadRequestConst,
		)
		return
	}

	policy := createPolicyHandler.parseDataToDomain(decodedPolicy)

	if err != nil {
		jsonResponse.ThrowError(
			routesConsts.MessageKeyConst,
			err,
			routesConsts.BadRequestConst,
		)
		return
	}

	createPolicyHandler.openDatabaseConnection()

	err = policyUsecase.NewCreatePolicy(
		createPolicyHandler.policyDatabase,
		policy,
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
		helper.PostPolicyConst,
		routesConsts.StatusOk,
	)
}

func (createPolicyHandler *CreatePolicyHandler) decodeRequest(
	context *gin.Context,
) (*requestEntity.PostPolicyRequestEntity, error) {
	policy, err := requestEntity.DecodePolicyRequest(context.Request)

	if err != nil {
		return nil, err
	}

	if err := policy.Validate(); err != nil {
		return nil, err
	}

	return policy, nil
}

func (createPolicyHandler *CreatePolicyHandler) parseDataToDomain(
	decodedPolicy *requestEntity.PostPolicyRequestEntity,
) *policyDomain.Policy {
	return policyDomain.NewPolicy(
		createPolicyHandler.uuid.GenerateUuid(),
		decodedPolicy.Status,
		nil,
		nil,
		nil,
		nil,
		decodedPolicy.UserId,
	)
}

func (createPolicyHandler *CreatePolicyHandler) openDatabaseConnection() {
	createPolicyHandler.policyDatabase = repository.NewPolicyDatabase(createPolicyHandler.connection)
}
