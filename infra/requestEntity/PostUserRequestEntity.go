package requestEntity

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anaclaraddias/brick/core/domain/helper"
)

const (
	NameFieldConst      = "o nome"
	CellphoneFieldConst = "o celular"
	CpfFieldConst       = "o cpf"
)

type PostUserRequestEntity struct {
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Cellphone string  `json:"cellphone"`
	Cpf       string  `json:"cpf"`
	Cnpj      *string `json:"cnpj"`
}

func DecodeUserRequest(request *http.Request) (*PostUserRequestEntity, error) {
	var user *PostUserRequestEntity

	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}

func (user *PostUserRequestEntity) Validate() error {
	if user.Name == "" {
		return fmt.Errorf(helper.UserFieldCanNotBeEmptyConst, NameFieldConst)
	}

	if err := helper.ValidateEmail(user.Email); err != nil {
		return err
	}

	if user.Cellphone == "" {
		return fmt.Errorf(helper.UserFieldCanNotBeEmptyConst, CellphoneFieldConst)
	}

	if user.Cpf == "" {
		return fmt.Errorf(helper.UserFieldCanNotBeEmptyConst, CpfFieldConst)
	}

	if err := helper.ValidateCPF(user.Cpf); err != nil {
		return err
	}

	user.Cpf = helper.UnmaskCpf(user.Cpf)

	if user.Cnpj != nil {
		if err := helper.ValidateCNPJ(*user.Cnpj); err != nil {
			return err
		}

		unmaskedCnpj := helper.UnmaskCnpj(*user.Cnpj)
		user.Cnpj = &unmaskedCnpj
	}

	return nil
}
