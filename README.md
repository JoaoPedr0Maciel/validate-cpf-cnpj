# Validate CPF/CNPJ

[![GoDoc](https://godoc.org/github.com/JoaoPedr0Maciel/validate-cpf-cnpj?status.svg)](https://pkg.go.dev/github.com/JoaoPedr0Maciel/validate-cpf-cnpj)
[![Go Report Card](https://goreportcard.com/badge/github.com/JoaoPedr0Maciel/validate-cpf-cnpj)](https://goreportcard.com/report/github.com/JoaoPedr0Maciel/validate-cpf-cnpj)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Overview

A Go package for validating Brazilian CPF and CNPJ numbers with concurrent processing and comprehensive error handling.

## Features

- **CPF Validation**: Validates CPF numbers using concurrent processing for better performance
- **Multiple Formats**: Accepts both formatted (`000.000.000-00`) and unformatted (`00000000000`) CPF numbers
- **Error Handling**: Provides detailed error messages for invalid inputs
- **CNPJ Support**: Coming soon...

## Installation

```bash
go get github.com/JoaoPedr0Maciel/validate-cpf-cnpj
```

## Usage

### Basic Usage

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
}
```

## Parameters

### ValidateCPF

| Parameter | Type | Description | Example |
| --- | --- | --- | --- |
| `cpf` | `string` | CPF number to validate | `"00000000000"` or `"000.000.000-00"` |

## Returns

| Value | Type | Description |
| --- | --- | --- |
| `isValid` | `bool` | `true` if CPF is valid, `false` otherwise |
| `err` | `error` | Error message if validation fails |

## Contributing

Contributions are welcome! Please:

1. Fork the repository
2. Create a feature branch
3. Add your changes
4. Run tests with `go test`
5. Submit a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
