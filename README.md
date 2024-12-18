# cpfcnpj

Package responsible to check if a CPF or CNPJ is valid or not.

NOTE: This package is able to handle the new **alphanumeric** format of CNPJ that is expected to be used in Brazil in July 2026.

Click [here](https://www.gov.br/receitafederal/pt-br/acesso-a-informacao/acoes-e-programas/programas-e-atividades/cnpj-alfanumerico#:~:text=O%20CNPJ%20Alfanum%C3%A9rico%20ser%C3%A1%20atribu%C3%ADdo,com%20o%20seu%20n%C3%BAmero%20v%C3%A1lido!) for more information about the new format.

Download the package:

```bash
go get github.com/jsfelipearaujo/cpfcnpj@latest
```

## How to use - CPF Validation

Import the package:

```go
import "github.com/jsfelipearaujo/cpf"
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
import "github.com/jsfelipearaujo/cpfcnpj/cnpj"
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
