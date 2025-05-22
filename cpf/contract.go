package cpf

// Represents a service that validates CPF (Brazilian taxpayer identification number for individuals)
type Service interface {
	// Validates if CPF is valid or not
	//
	// Example:
	//
	//	svc := cpf.New()
	//	err := svc.IsValid("123.456.789-10")
	//	if err != nil {
	//		// Handle invalid CPF
	//	}
	IsValid(cpf string) error

	// Generates a random CPF
	//
	// Example:
	//
	//	svc := cpf.New()
	//	cpf := svc.Generate(true)	// true for pretty format
	//	fmt.Println(cpf)		// 123.456.789-10
	Generate(pretty bool) string
}
