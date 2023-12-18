package repository

import (
	"database/sql"
	"time"

	userDomain "github.com/anaclaraddias/brick/core/domain/user"
	"github.com/anaclaraddias/brick/core/port"
	repositoryInterface "github.com/anaclaraddias/brick/core/port/repository"
	"github.com/anaclaraddias/brick/infra/models"
)

type UserDatabase struct {
	connection port.DatabaseConnectionInterface
}

func NewUserDatabase(
	connection port.DatabaseConnectionInterface,
) repositoryInterface.UserRepositoryInterface {
	connection.Open()

	return &UserDatabase{connection: connection}
}

func (userDatabase *UserDatabase) CreateUser(user *userDomain.User) error {
	var dbUser *models.UserModel

	createdAt := time.Now()
	updatedAt := sql.NullTime{Valid: false}

	query := `INSERT INTO users (
		id,
		name,
		email,
		cellphone,
		cpf,
		cnpj,
		update_date,
		creation_date
	) values (?,?,?,?,?,?,?,?)`

	if err := userDatabase.connection.Raw(
		query,
		&dbUser,
		user.Id,
		user.Name,
		user.Email,
		user.Cellphone,
		user.Cpf,
		user.Cnpj,
		updatedAt,
		createdAt,
	); err != nil {
		return err
	}

	return nil
}
