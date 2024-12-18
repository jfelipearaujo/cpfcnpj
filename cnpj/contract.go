package cnpj

// Represents a service that validates CNPJ (Brazilian taxpayer identification number for companies)
type Service interface {
	// Validates if CNPJ is valid or not
	//
	// # Note: This function is able to handle the new alphanumeric format of CNPJ
	//
	// Example:
	//
	//	svc := cnpj.New("12.ABC.345/01DE-35")
	//	err := svc.IsValid()
	//	if err != nil {
	//		// Handle invalid CNPJ
	//	}
	IsValid() error
}
