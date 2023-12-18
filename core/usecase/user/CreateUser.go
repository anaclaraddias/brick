package userUsecase

import (
	userDomain "github.com/anaclaraddias/brick/core/domain/user"
	repositoryInterface "github.com/anaclaraddias/brick/core/port/repository"
)

type CreateUser struct {
	userDatabase repositoryInterface.UserRepositoryInterface
}

func NewCreateUser(
	userDatabase repositoryInterface.UserRepositoryInterface,
) *CreateUser {
	return &CreateUser{
		userDatabase: userDatabase,
	}
}

func (createUser *CreateUser) Execute(user *userDomain.User) error {
	err := createUser.userDatabase.CreateUser(user)

	if err != nil {
		return err
	}

	return nil
}
