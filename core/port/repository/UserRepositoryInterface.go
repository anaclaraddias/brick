package repositoryInterface

import userDomain "github.com/anaclaraddias/brick/core/domain/user"

type UserRepositoryInterface interface {
	CreateUser(user *userDomain.User) error
	FindUserByCpfOrCnpj(cpf string, cnpj *string) (*userDomain.User, error)
}
