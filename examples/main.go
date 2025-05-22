package main

import (
	"fmt"
	"log"

	"github.com/jfelipearaujo/cpfcnpj/cnpj"
	"github.com/jfelipearaujo/cpfcnpj/cpf"
)

func main() {
	cpfData := "983.533.180-40"
	cnpjData := "12.ABC.345/01DE-35"

	cpfService := cpf.New()
	if err := cpfService.IsValid(cpfData); err != nil {
		log.Fatal(err)
	}

	cnpjService := cnpj.New()
	if err := cnpjService.IsValid(cnpjData); err != nil {
		log.Fatal(err)
	}

	aNewCPF := cpfService.Generate(cpf.WithPrettyFormat())
	fmt.Printf("CPF:\t\t%s\n", aNewCPF)

	aNewCNPJ := cnpjService.Generate(cnpj.WithPrettyFormat())
	fmt.Printf("CNPJ:\t\t%s\n", aNewCNPJ)

	aNewCNPJV2 := cnpjService.Generate(
		cnpj.WithVersion(cnpj.V2),
		cnpj.WithPrettyFormat(),
	)
	fmt.Printf("CNPJ (V2):\t%s\n", aNewCNPJV2)
}
