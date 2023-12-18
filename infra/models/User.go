package models

import (
	"database/sql"
	"time"
)

type UserModel struct {
	Id        string
	Name      string
	Email     string
	Cellphone string
	Cpf       string
	Cnpj      *string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
