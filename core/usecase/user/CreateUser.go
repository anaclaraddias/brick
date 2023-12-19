package userUsecase

import (
	"fmt"

	"github.com/anaclaraddias/brick/core/domain/helper"
	userDomain "github.com/anaclaraddias/brick/core/domain/user"
	repositoryInterface "github.com/anaclaraddias/brick/core/port/repository"
)

type CreateUser struct {
	userDatabase repositoryInterface.UserRepositoryInterface
	user         *userDomain.User
}

func NewCreateUser(
	userDatabase repositoryInterface.UserRepositoryInterface,
	user *userDomain.User,
) *CreateUser {
	return &CreateUser{
		userDatabase: userDatabase,
		user:         user,
	}
}

func (createUser *CreateUser) Execute() error {
	if err := createUser.verifyIfUserAlreadyExists(); err != nil {
		return err
	}

	err := createUser.userDatabase.CreateUser(createUser.user)

	if err != nil {
		return err
	}

	return nil
}

func (createUser *CreateUser) verifyIfUserAlreadyExists() error {
	user, err := createUser.userDatabase.FindUserByCpfOrCnpj(
		createUser.user.Cpf,
		createUser.user.Cnpj,
	)

	if err != nil {
		return err
	}

	if user != nil {
		return fmt.Errorf(helper.UserlreadyExistsConst)
	}

	return nil
}
