package helper

import "regexp"

func UnmaskLicensePlate(licensePlate string) string {
	re := regexp.MustCompile(`[-]`)
	return re.ReplaceAllString(licensePlate, "")
}

func UnmaskCnpj(cnpj string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(cnpj, "")
}

func UnmaskCpf(cpf string) string {
	re := regexp.MustCompile(`[.-]`)
	return re.ReplaceAllString(cpf, "")
}
