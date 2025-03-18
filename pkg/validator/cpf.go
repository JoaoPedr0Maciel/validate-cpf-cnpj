package validator

import (
	"errors"

	"github.com/JoaoPedr0Maciel/validate-cpf-cnpj/pkg/clean"
)

var (
	weightCpf1 = []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	weightCpf2 = []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}

)

func ValidateCPF(cnpj string) (bool, error) {
	cleanedCnpj, err := clean.CleanCnpj(cnpj)
	if err != nil {
		return false, err
	}

	firstDigitUser := int(cleanedCnpj[9] - '0')
	secondDigitUser := int(cleanedCnpj[10] - '0')

	firstDigitCalculated := calculateDigitCPF(cleanedCnpj, true)
	if firstDigitCalculated != firstDigitUser {
		return false, errors.New("this CPF is invalid")
	}

	secondDigitCalculated := calculateDigitCPF(cleanedCnpj, false)
	if secondDigitCalculated != secondDigitUser {
		return false, errors.New("this CPF is invalid")
	}

	return true, nil
}

func calculateDigitCPF(cleanedCnpj string, isFirstDigit bool) int {
	weight := weightCpf1
	cleanedCnpjPosition := cleanedCnpj[:9]

	if isFirstDigit {
		weight = weightCpf2
		cleanedCnpjPosition = cleanedCnpj[:10]
	}

	sum := 0
	for i, num := range cleanedCnpjPosition {
		numInt := int(num - '0')
		result := numInt * weight[i]
		sum += result
	}

	rest := sum % 11
	if rest <= 2 {
		return 0
	}
	return 11 - rest
}
