package cpf

// Represents a service that validates CPF (Brazilian taxpayer identification number for individuals)
type Service interface {
	// Validates if CPF is valid or not
	//
	// Example:
	//
	//	svc := cpf.New("123.456.789-10")
	//	err := svc.IsValid()
	//	if err != nil {
	//		// Handle invalid CPF
	//	}
	IsValid() error
}
