# Validate CPF/CNPJ

[![GoDoc](https://godoc.org/github.com/JoaoPedr0Maciel/validate-cpf-cnpj?status.svg)](https://pkg.go.dev/github.com/JoaoPedr0Maciel/validate-cpf-cnpj)
[![Go Report Card](https://goreportcard.com/badge/github.com/JoaoPedr0Maciel/validate-cpf-cnpj)](https://goreportcard.com/report/github.com/JoaoPedr0Maciel/validate-cpf-cnpj)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


## Overview
A Go package for validating Brazilian CPF and CNPJ numbers with concurrent processing and comprehensive error handling.

## Features
- **CPF Validation**: Validates CPF numbers using concurrent processing for better performance
- **CNPJ Validation**: Validates CNPJ numbers with proper algorithm checks
- **Multiple Formats**: Accepts both formatted (000.000.000-00 or 00.000.000/0000-00) and unformatted (00000000000 or 00000000000000) numbers
- **Error Handling**: Provides detailed error messages for invalid inputs
- **Cleaning Functions**: Utilities to clean CPF and CNPJ numbers

## Installation
```bash
go get github.com/JoaoPedr0Maciel/validate-cpf-cnpj
```

## Usage

### CPF Validation
```go
package main

import (
    "fmt"
    "github.com/JoaoPedr0Maciel/validate-cpf-cnpj/pkg/validator"
)

func main() {
    // Valid CPF
    isValid, err := validator.ValidateCPF("951.765.520-78")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Is CPF valid:", isValid) // true

    // Invalid CPF
    isValid, err = validator.ValidateCPF("12345678909")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Is CPF valid:", isValid) // false
}
```

### CNPJ Validation
```go
package main

import (
    "fmt"
    "github.com/JoaoPedr0Maciel/validate-cpf-cnpj/pkg/validator"
)

func main() {
    // Valid CNPJ
    isValid, err := validator.ValidateCNPJ("13.347.016/0001-17")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Is CNPJ valid:", isValid) // true

    // Invalid CNPJ
    isValid, err = validator.ValidateCNPJ("11111111111111")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Is CNPJ valid:", isValid) // false
}
```

### Cleaning Functions
```go
package main

import (
    "fmt"
    "github.com/JoaoPedr0Maciel/validate-cpf-cnpj/pkg/clean"
)

func main() {
    // Clean CPF (remove special characters)
    cleanedCPF := clean.CleanCPF("123.456.789-09")
    fmt.Println("Cleaned CPF:", cleanedCPF) // 12345678909
    
    // Clean CNPJ (remove special characters)
    cleanedCNPJ := clean.CleanCNPJ("11.222.333/0001-44")
    fmt.Println("Cleaned CNPJ:", cleanedCNPJ) // 11222333000144
}
```

### Error Handling
```go
package main

import (
    "fmt"
    "github.com/JoaoPedr0Maciel/validate-cpf-cnpj/pkg/validator"
)

func main() {
    // Empty CPF
    isValid, err := validator.ValidateCPF("")
    if err != nil {
        fmt.Println("Error:", err) // "cpf must be a valid string (000.000.000-00) or (00000000000)"
        return
    }

    // Invalid Format
    isValid, err = validator.ValidateCPF("invalid")
    if err != nil {
        fmt.Println("Error:", err) // "this CPF is invalid"
        return
    }
    
    // Empty CNPJ
    isValid, err = validator.ValidateCNPJ("")
    if err != nil {
        fmt.Println("Error:", err) // "cnpj must be a valid string (00.000.000/0000-00) or (00000000000000)"
        return
    }
}
```

## Parameters

### ValidateCPF
| Parameter | Type | Description | Example |
| --- | --- | --- | --- |
| cpf | string | CPF number to validate | "00000000000" or "000.000.000-00" |

### ValidateCNPJ
| Parameter | Type | Description | Example |
| --- | --- | --- | --- |
| cnpj | string | CNPJ number to validate | "00000000000000" or "00.000.000/0000-00" |

### Clean Functions
| Function | Parameter | Return | Description |
| --- | --- | --- | --- |
| CleanCPF | string (CPF) | string | Removes formatting from CPF |
| CleanCNPJ | string (CNPJ) | string | Removes formatting from CNPJ |

## Returns
| Value | Type | Description |
| --- | --- | --- |
| isValid | bool | true if CPF/CNPJ is valid, false otherwise |
| err | error | Error message if validation fails |

## Contributing
Contributions are welcome! Please:
1. Fork the repository
2. Create a feature branch
3. Add your changes
4. Run tests with `go test`
5. Submit a pull request

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.