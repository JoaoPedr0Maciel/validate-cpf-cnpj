package clean

import (
	"errors"
	"regexp"
)

func CleanCpf(cpf string) (string, error) {
	re := regexp.MustCompile(`\D`)
	cleanedCPF := re.ReplaceAllString(cpf, "")

	if len(cleanedCPF) != 11 {
		return "", errors.New("CPF must have only 11 numeric digits")
	}

	return cleanedCPF, nil
}
