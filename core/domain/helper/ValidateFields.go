package helper

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

const (
	EmailMaxLengthConst = 80
	RenavamLengthConst  = 11
	CpfLengthConst      = 11
	CnpjLengthConst     = 14
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

func ValidateLicensePlate(licensePlate string) error {
	licensePlate = UnmaskLicensePlate(licensePlate)

	plateRegex := regexp.MustCompile(`^[A-Z]{3}[0-9]{4}$`)

	if !plateRegex.MatchString(licensePlate) {
		return fmt.Errorf(InvalidLicensePlateFormatConst)
	}

	return nil
}

func ValidateRenavam(renavam string) error {
	if len(renavam) != RenavamLengthConst {
		return fmt.Errorf(RenavamMustHave11digitsConst)
	}

	_, err := strconv.Atoi(renavam)
	if err != nil {
		return fmt.Errorf(RenavamMustHaveOnlyNumbersConst)
	}

	weights := []int{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	var sum int

	for i := 0; i < 10; i++ {
		digit, _ := strconv.Atoi(string(renavam[i]))
		sum += digit * weights[i]
	}

	remainder := sum % 11
	var expectedDigit int

	if remainder < 2 {
		expectedDigit = 0
	} else {
		expectedDigit = 11 - remainder
	}

	lastDigit, _ := strconv.Atoi(string(renavam[10]))

	if lastDigit != expectedDigit {
		return fmt.Errorf(InvalidRenavamFormatConst)
	}

	return nil
}

func ValidateCNPJ(cnpj string) error {
	cnpj = UnmaskCnpj(cnpj)

	if len(cnpj) != CnpjLengthConst {
		return fmt.Errorf(CnpjMustHave14digitsConst)
	}

	sum := 0
	factor := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	for i := 0; i < 12; i++ {
		digit, err := strconv.Atoi(string(cnpj[i]))
		if err != nil {
			return fmt.Errorf(InvalidCnpjFormatConst)
		}
		sum += digit * factor[i]
	}
	remainder := sum % 11
	digit1 := 11 - remainder
	if digit1 >= 10 {
		digit1 = 0
	}
	if digit1 != int(cnpj[12]-'0') {
		return fmt.Errorf(InvalidCnpjFormatConst)
	}

	sum = 0
	factor = append([]int{6}, factor...)
	for i := 0; i < 13; i++ {
		digit, err := strconv.Atoi(string(cnpj[i]))
		if err != nil {
			return fmt.Errorf(InvalidCnpjFormatConst)
		}
		sum += digit * factor[i]
	}
	remainder = sum % 11
	digit2 := 11 - remainder
	if digit2 >= 10 {
		digit2 = 0
	}
	if digit2 != int(cnpj[13]-'0') {
		return fmt.Errorf(InvalidCnpjFormatConst)
	}

	return nil
}

func ValidateCPF(cpf string) error {
	cpf = UnmaskCpf(cpf)

	cpf = strings.Trim(cpf, " ")
	cpf = fmt.Sprintf("%011s", cpf)

	_, err := strconv.Atoi(cpf)

	if err != nil {
		return fmt.Errorf(CpfShouldContainOnlyNumbersCont)
	}

	cpf = regexp.MustCompile(`\D`).ReplaceAllString(cpf, "")

	if len(cpf) != CpfLengthConst {
		return fmt.Errorf(InvalidCpfLengthErrorConst)
	}

	digits := make([]int, CpfLengthConst)
	for i, char := range cpf {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return fmt.Errorf(CpfIsInvalidErrorConst)
		}
		digits[i] = digit
	}

	sum1 := 0
	for i := 0; i < 9; i++ {
		sum1 += digits[i] * (10 - i)
	}
	mod1 := (sum1 * 10) % 11
	if mod1 == 10 {
		mod1 = 0
	}

	sum2 := 0
	for i := 0; i < 9; i++ {
		sum2 += digits[i] * (11 - i)
	}
	sum2 += mod1 * 2
	mod2 := (sum2 * 10) % 11
	if mod2 == 10 {
		mod2 = 0
	}

	if mod1 != digits[9] || mod2 != digits[10] {
		return fmt.Errorf(CpfIsInvalidErrorConst)
	}

	return nil
}
