package clean

import (
	"errors"
	"regexp"
)

func CleanCnpj(cnpj string) (string, error) {
	re := regexp.MustCompile(`\D`)
	cleanedCnpj := re.ReplaceAllString(cnpj, "")

	if len(cleanedCnpj) > 14 {
		return "", errors.New("CNPJ must have only 14 numeric digits")
	}

	return cleanedCnpj, nil
}
