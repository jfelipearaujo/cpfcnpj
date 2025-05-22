# cpfcnpj
[![continuous integration](https://github.com/jfelipearaujo/cpfcnpj/actions/workflows/ci.yml/badge.svg)](https://github.com/jfelipearaujo/cpfcnpj/actions/workflows/ci.yml)
[![codecov](https://codecov.io/github/jfelipearaujo/cpfcnpj/graph/badge.svg?token=BR99CZ9VZ8)](https://codecov.io/github/jfelipearaujo/cpfcnpj)
[![version](https://img.shields.io/github/v/release/jfelipearaujo/cpfcnpj.svg)](https://github.com/jfelipearaujo/cpfcnpj/releases/latest)
[![license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/jfelipearaujo/cpfcnpj/blob/main/LICENSE)

Package responsible to check if a CPF or CNPJ is valid or not.

NOTE: This package is able to handle the new **alphanumeric** format of CNPJ that is expected to be used in Brazil in July 2026.

Click [here](https://www.gov.br/receitafederal/pt-br/acesso-a-informacao/acoes-e-programas/programas-e-atividades/cnpj-alfanumerico#:~:text=O%20CNPJ%20Alfanum%C3%A9rico%20ser%C3%A1%20atribu%C3%ADdo,com%20o%20seu%20n%C3%BAmero%20v%C3%A1lido!) for more information about the new format.

Download the package:

```bash
go get github.com/jfelipearaujo/cpfcnpj@latest
```
# How to use

- [cpfcnpj](#cpfcnpj)
- [How to use](#how-to-use)
  - [CPF Validation](#cpf-validation)
  - [CNPJ Validation](#cnpj-validation)
  - [CPF Generator](#cpf-generator)
  - [CNPJ Generator](#cnpj-generator)
  - [Contributing](#contributing)
  - [LICENSE](#license)

## CPF Validation

Import the package:

```go
import "github.com/jfelipearaujo/cpf"
```

Create a new instance of the service:

```go
svc := cpf.New()
```

Validate if the CPF is valid or not:

```go
err := svc.IsValid("123.456.789-10")
if err != nil {
    // Handle invalid CPF
}
```

## CNPJ Validation

Import the package:

```go
import "github.com/jfelipearaujo/cpfcnpj/cnpj"
```

Create a new instance of the service:

```go
svc := cnpj.New()
```

Validate if the CNPJ is valid or not:

```go
err := svc.IsValid("12.ABC.345/01DE-35")
if err != nil {
    // Handle invalid CNPJ
}
```

## CPF Generator

Import the package:

```go
import "github.com/jfelipearaujo/cpf"
```

Create a new instance of the service:

```go
svc := cpf.New()
```

Generate a new CPF:

```go
cpf := svc.Generate(true) // true to generate with pretty format
```

## CNPJ Generator

Import the package:

```go
import "github.com/jfelipearaujo/cpfcnpj/cnpj"
```

Create a new instance of the service:

```go
svc := cnpj.New()
```

Generate a new CNPJ:

```go
cnpj := svc.Generate(true) // true to generate with pretty format
```

## Contributing

Contributions are welcome!

## LICENSE

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
