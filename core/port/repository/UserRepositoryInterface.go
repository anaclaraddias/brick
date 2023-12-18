package repositoryInterface

import userDomain "github.com/anaclaraddias/brick/core/domain/user"

type UserRepositoryInterface interface {
	CreateUser(user *userDomain.User) error
}
