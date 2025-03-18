package validator

import (
	"errors"

	"github.com/JoaoPedr0Maciel/validate-cpf-cnpj/pkg/clean"
)

var (
	weightCnpj1 = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	weightCnpj2 = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
)

func ValidateCNPJ(cnpj string) (bool, error) {
	cleanedCnpj, err := clean.CleanCnpj(cnpj)
	if err != nil {
		return false, err
	}

	firstDigitUser := int(cleanedCnpj[12] - '0')
	secondDigitUser := int(cleanedCnpj[13] - '0')

	firstDigitCalculated := calculateDigitCNPJ(cleanedCnpj, true)
	if firstDigitCalculated != firstDigitUser {
		return false, errors.New("this CNPJ is invalid")
	}

	secondDigitCalculated := calculateDigitCNPJ(cleanedCnpj, false)
	if secondDigitCalculated != secondDigitUser {
		return false, errors.New("this CNPJ is invalid")
	}

	return true, nil
}

func calculateDigitCNPJ(cleanedCnpj string, isFirstDigit bool) int {
	weight := map[bool][]int{true: weightCnpj1, false: weightCnpj2}[isFirstDigit]
	cleanedCnpjPosition := cleanedCnpj[:12+boolToInt(!isFirstDigit)]

	var sum int
	for i, num := range cleanedCnpjPosition {
		sum += int(num-'0') * weight[i]
	}

	rest := sum % 11
	return map[bool]int{true: 0, false: 11 - rest}[rest <= 2]
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
