package cnpj

// Represents a service that validates CNPJ (Brazilian taxpayer identification number for companies)
type Service interface {
	// Validates if CNPJ is valid or not
	//
	// # Note: This function is able to handle the new alphanumeric format of CNPJ
	//
	// Example:
	//
	//	svc := cnpj.New()
	//	err := svc.IsValid("12.ABC.345/01DE-35")
	//	if err != nil {
	//		// Handle invalid CNPJ
	//	}
	IsValid(cnpj string) error

	// Generates a random CNPJ
	//
	// Example:
	//
	//	svc := cnpj.New()
	//	cnpj := svc.Generate(true)	// true for pretty format
	//	fmt.Println(cnpj)		// 12.ABC.345/01DE-35
	Generate(pretty bool) string
}
