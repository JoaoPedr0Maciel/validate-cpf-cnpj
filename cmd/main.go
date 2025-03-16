package main

import (
	"fmt"

	"github.com/JoaoPedr0Maciel/validate-cpf-cnpj/pkg/validator"
)

func main() {
	isValid, err := validator.ValidateCPF("000.000.000-00")

	if err != nil {
		panic(err)
	}

	fmt.Printf("is valid => %v", isValid)
}
