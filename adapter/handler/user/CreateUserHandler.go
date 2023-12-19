package userHandler

import (
	routesConsts "github.com/anaclaraddias/brick/adapter/http/constants"
	"github.com/anaclaraddias/brick/adapter/http/routes"
	"github.com/anaclaraddias/brick/core/domain/helper"
	userDomain "github.com/anaclaraddias/brick/core/domain/user"
	"github.com/anaclaraddias/brick/core/port"
	repositoryInterface "github.com/anaclaraddias/brick/core/port/repository"
	userUsecase "github.com/anaclaraddias/brick/core/usecase/user"
	"github.com/anaclaraddias/brick/infra/database/repository"
	"github.com/anaclaraddias/brick/infra/requestEntity"
	"github.com/gin-gonic/gin"
)

type CreateUserHandler struct {
	connection   port.DatabaseConnectionInterface
	userDatabase repositoryInterface.UserRepositoryInterface
	uuid         port.UuidInterface
}

func NewCreateUserHandler(
	connection port.DatabaseConnectionInterface,
	uuid port.UuidInterface,
) port.HandlerInterface {
	return &CreateUserHandler{
		connection: connection,
		uuid:       uuid,
	}
}

func (createUserHandler *CreateUserHandler) Handle(context *gin.Context) {
	jsonResponse := routes.NewJsonResponse(context)

	decodedUser, err := createUserHandler.decodeRequest(context)

	if err != nil {
		jsonResponse.ThrowError(
			routesConsts.MessageKeyConst,
			err,
			routesConsts.BadRequestConst,
		)
		return
	}

	user := createUserHandler.parseDataToDomain(decodedUser)

	if err != nil {
		jsonResponse.ThrowError(
			routesConsts.MessageKeyConst,
			err,
			routesConsts.BadRequestConst,
		)
		return
	}

	createUserHandler.openDatabaseConnection()

	err = userUsecase.NewCreateUser(
		createUserHandler.userDatabase,
		user,
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
		helper.PostUserConst,
		routesConsts.StatusOk,
	)
}

func (createUserHandler *CreateUserHandler) decodeRequest(
	context *gin.Context,
) (*requestEntity.PostUserRequestEntity, error) {
	decodedUser, err := requestEntity.DecodeUserRequest(context.Request)

	if err != nil {
		return nil, err
	}

	if err := decodedUser.Validate(); err != nil {
		return nil, err
	}

	return decodedUser, nil
}

func (createUserHandler *CreateUserHandler) parseDataToDomain(
	decodedUser *requestEntity.PostUserRequestEntity,
) *userDomain.User {
	return userDomain.NewUser(
		createUserHandler.uuid.GenerateUuid(),
		decodedUser.Name,
		decodedUser.Email,
		decodedUser.Cellphone,
		decodedUser.Cpf,
		decodedUser.Cnpj,
	)
}

func (createUserHandler *CreateUserHandler) openDatabaseConnection() {
	createUserHandler.userDatabase = repository.NewUserDatabase(createUserHandler.connection)
}
