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

type CreatePolicyCoverageHandler struct {
	connection       port.DatabaseConnectionInterface
	policyDatabase   repositoryInterface.PolicyRepositoryInterface
	coverageDatabase repositoryInterface.CoverageRepositoryInterface
	uuid             port.UuidInterface
}

func NewCreatePolicyCoverageHandler(
	connection port.DatabaseConnectionInterface,
	uuid port.UuidInterface,
) port.HandlerInterface {
	return &CreatePolicyCoverageHandler{
		connection: connection,
		uuid:       uuid,
	}
}

func (createPolicyCoverageHandler *CreatePolicyCoverageHandler) Handle(context *gin.Context) {
	jsonResponse := routes.NewJsonResponse(context)

	decodedPolicyCoverage, err := createPolicyCoverageHandler.decodeRequest(context)

	if err != nil {
		jsonResponse.ThrowError(
			routesConsts.MessageKeyConst,
			err,
			routesConsts.BadRequestConst,
		)
		return
	}

	policyCoverage := createPolicyCoverageHandler.parseDataToDomain(decodedPolicyCoverage)

	if err != nil {
		jsonResponse.ThrowError(
			routesConsts.MessageKeyConst,
			err,
			routesConsts.BadRequestConst,
		)
		return
	}

	createPolicyCoverageHandler.openDatabaseConnection()

	err = policyUsecase.NewCreatePolicyCoverage(
		createPolicyCoverageHandler.coverageDatabase,
		createPolicyCoverageHandler.policyDatabase,
		policyCoverage,
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
		helper.PostPolicyCoverageConst,
		routesConsts.StatusOk,
	)
}

func (createPolicyCoverageHandler *CreatePolicyCoverageHandler) decodeRequest(
	context *gin.Context,
) (*requestEntity.PostPolicyCoverageRequestEntity, error) {
	policy, err := requestEntity.DecodePolicyCoverageRequest(context.Request)

	if err != nil {
		return nil, err
	}

	if err := policy.Validate(); err != nil {
		return nil, err
	}

	return policy, nil
}

func (createPolicyCoverageHandler *CreatePolicyCoverageHandler) parseDataToDomain(
	decodedPolicyCoverage *requestEntity.PostPolicyCoverageRequestEntity,
) *policyDomain.PolicyCoverage {
	return policyDomain.NewPolicyCoverage(
		createPolicyCoverageHandler.uuid.GenerateUuid(),
		decodedPolicyCoverage.CoverageId,
		decodedPolicyCoverage.PolicyId,
	)
}

func (createPolicyCoverageHandler *CreatePolicyCoverageHandler) openDatabaseConnection() {
	createPolicyCoverageHandler.coverageDatabase = repository.NewCoverageDatabase(
		createPolicyCoverageHandler.connection,
	)

	createPolicyCoverageHandler.policyDatabase = repository.NewPolicyDatabase(
		createPolicyCoverageHandler.connection,
	)
}
