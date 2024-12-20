# cpfcnpj
[![tests](https://github.com/jfelipearaujo/cpfcnpj/actions/workflows/tests.yml/badge.svg)](https://github.com/jfelipearaujo/cpfcnpj/actions/workflows/tests.yml)
[![version](https://img.shields.io/github/v/release/jfelipearaujo/cpfcnpj.svg)](https://github.com/jfelipearaujo/cpfcnpj/releases/latest)
[![license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/jfelipearaujo/cpfcnpj/blob/main/LICENSE)

Package responsible to check if a CPF or CNPJ is valid or not.

NOTE: This package is able to handle the new **alphanumeric** format of CNPJ that is expected to be used in Brazil in July 2026.

Click [here](https://www.gov.br/receitafederal/pt-br/acesso-a-informacao/acoes-e-programas/programas-e-atividades/cnpj-alfanumerico#:~:text=O%20CNPJ%20Alfanum%C3%A9rico%20ser%C3%A1%20atribu%C3%ADdo,com%20o%20seu%20n%C3%BAmero%20v%C3%A1lido!) for more information about the new format.

Download the package:

```bash
go get github.com/jfelipearaujo/cpfcnpj@latest
```

## How to use - CPF Validation

Import the package:

```go
import "github.com/jfelipearaujo/cpf"
```

Create a new instance of the service:

```go
svc := cpf.New("123.456.789-10")
```

Validate if the CPF is valid or not:

```go
err := svc.IsValid()
if err != nil {
    // Handle invalid CPF
}
```

## How to use - CNPJ Validation

Import the package:

```go
import "github.com/jfelipearaujo/cpfcnpj/cnpj"
```

Create a new instance of the service:

```go
svc := cnpj.New("12.ABC.345/01DE-35")
```

Validate if the CNPJ is valid or not:

```go
err := svc.IsValid()
if err != nil {
    // Handle invalid CNPJ
}
```

## Contributing

Contributions are welcome!

## LICENSE

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
