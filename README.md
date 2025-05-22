# cpfcnpj
[![continuous integration](https://github.com/jfelipearaujo/cpfcnpj/actions/workflows/ci.yml/badge.svg)](https://github.com/jfelipearaujo/cpfcnpj/actions/workflows/ci.yml)
[![codecov](https://codecov.io/github/jfelipearaujo/cpfcnpj/graph/badge.svg?token=BR99CZ9VZ8)](https://codecov.io/github/jfelipearaujo/cpfcnpj)
[![version](https://img.shields.io/github/v/release/jfelipearaujo/cpfcnpj.svg)](https://github.com/jfelipearaujo/cpfcnpj/releases/latest)
[![license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/jfelipearaujo/cpfcnpj/blob/main/LICENSE)

Package responsible to check if a CPF or CNPJ is valid or not.

NOTE: This package is able to handle the new **alphanumeric** format of CNPJ that is expected to be used in Brazil from July 2026.

Click [here](https://www.gov.br/receitafederal/pt-br/acesso-a-informacao/acoes-e-programas/programas-e-atividades/cnpj-alfanumerico#:~:text=O%20CNPJ%20Alfanum%C3%A9rico%20ser%C3%A1%20atribu%C3%ADdo,com%20o%20seu%20n%C3%BAmero%20v%C3%A1lido!) for more information about the new format.

Download the package:

```bash
go get github.com/jfelipearaujo/cpfcnpj@latest
```
# How to use

- [cpfcnpj](#cpfcnpj)
- [How to use](#how-to-use)
  - [Examples](#examples)
  - [CPF Validation](#cpf-validation)
  - [CNPJ Validation](#cnpj-validation)
  - [CPF Generator](#cpf-generator)
  - [CNPJ Generator](#cnpj-generator)
    - [CNPJ Version](#cnpj-version)
  - [Contributing](#contributing)
  - [LICENSE](#license)

## Examples

You can find some examples in the [examples](examples) folder.

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
cpf := svc.Generate()
```

If you want to generate the CPF with pretty format, do the following:

```go
cpf := svc.Generate(cpf.WithPrettyFormat())
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
cnpj := svc.Generate()
```

If you want to generate the CNPJ with pretty format, do the following:

```go
cnpj := svc.Generate(cnpj.WithPrettyFormat())
```

### CNPJ Version

This package is able to handle the new **alphanumeric** format of CNPJ that is expected to be used in Brazil from July 2026.

By default every CNPJ generate will use the V1 version with only numbers, but if you want to generate a CNPJ with the new version, do the following:

```go
cnpj := svc.Generate(cnpj.WithVersion(cnpj.V2))
```

## Contributing

Contributions are welcome!

Please read the [CONTRIBUTING](CONTRIBUTING.md) file for details on our code of conduct, and the process for submitting pull requests.

## LICENSE

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
