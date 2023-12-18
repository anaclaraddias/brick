package userDomain

type User struct {
	Id        string
	Name      string
	Email     string
	Cellphone string
	Cpf       string
	Cnpj      *string
}

func NewUser(
	id string,
	name string,
	email string,
	cellphone string,
	cpf string,
	cnpj *string,
) *User {
	return &User{
		Id:        id,
		Name:      name,
		Email:     email,
		Cellphone: cellphone,
		Cpf:       cpf,
		Cnpj:      cnpj,
	}
}
