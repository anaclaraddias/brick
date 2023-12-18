package helper

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"
)

const (
	EmailMaxLengthConst = 80
)

func ValidateEmail(email string) error {
	if len(email) == 0 {
		return fmt.Errorf(EmailCanNotBeEmptyConst)
	}

	if len(email) > EmailMaxLengthConst {
		return fmt.Errorf(EmailTooLongConst)
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return fmt.Errorf(InvalidEmailFormatConst)
	}

	return nil
}

func ValidateUUID(hash string, fieldName string) error {
	_, err := uuid.Parse(hash)

	if err != nil {
		return fmt.Errorf(InformFieldConst, fieldName)
	}

	return nil
}
