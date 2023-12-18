package helper

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	CpfLengthConst                   = 11
	CpfIsInvalidErrorConst           = "o CPF é inválido"
	InvalidCpfLengthErrorConst       = "o CPF não possui o tamanho correto"
	CpfShouldContainOnlyNumbersConts = "o CPF deve conter apenas números"
)

type Cpf struct {
	Number string
}

func NewCpf(number string) *Cpf {
	return &Cpf{
		number,
	}
}

func (cpf *Cpf) Validate() error {
	cpf.Unmask()

	cpf.Number = strings.Trim(cpf.Number, " ")
	cpf.Number = fmt.Sprintf("%011s", cpf.Number)

	_, err := strconv.Atoi(cpf.Number)

	if err != nil {
		return fmt.Errorf(CpfShouldContainOnlyNumbersConts)
	}

	if err := cpf.checkCpfValidity(); err != nil {
		return err
	}

	return nil
}

func (cpf *Cpf) Unmask() {
	re := regexp.MustCompile(`[.-]`)
	cpf.Number = re.ReplaceAllString(cpf.Number, "")
}

func (cpf *Cpf) checkCpfValidity() error {
	cpf.Number = regexp.MustCompile(`\D`).ReplaceAllString(cpf.Number, "")

	if len(cpf.Number) != CpfLengthConst {
		return fmt.Errorf(InvalidCpfLengthErrorConst)
	}

	digits, err := cpf.parseCpfDigits()
	if err != nil {
		return err
	}

	mod1, mod2 := cpf.calculateChecksums(digits)

	if mod1 != digits[9] || mod2 != digits[10] {
		return fmt.Errorf(CpfIsInvalidErrorConst)
	}

	return nil
}

func (cpf *Cpf) parseCpfDigits() ([]int, error) {
	digits := make([]int, CpfLengthConst)
	for i, char := range cpf.Number {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return nil, fmt.Errorf(CpfIsInvalidErrorConst)
		}
		digits[i] = digit
	}
	return digits, nil
}

func (cpf *Cpf) calculateChecksums(digits []int) (int, int) {
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

	return mod1, mod2
}
