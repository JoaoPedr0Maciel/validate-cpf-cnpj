package main

import (
	"fmt"

	"github.com/JoaoPedr0Maciel/validate-cpf-cnpj/pkg/validator"
)

func main() {
	isCPFvalid, err := validator.ValidateCPF("000.000.000-00")

	if err != nil {
		panic(err)
	}

	isCNPJvalid, err := validator.ValidateCNPJ("00.000.000/0001-00")

	if err != nil {
		panic(err)
	}

	fmt.Printf("is valid CPF => %v\n\n", isCPFvalid)
	fmt.Printf("is valid CNPJ => %v\n\n", isCNPJvalid)
}
