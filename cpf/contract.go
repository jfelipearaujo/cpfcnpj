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
	//	cpf := svc.Generate()
	//	fmt.Println(cpf) // 12345678910
	Generate(opts ...func(*generatorOptions)) string
}

type generatorOptions struct {
	prettyFormat bool
}

// Sets the pretty format of the CPF when generating a new one
//
// Example:
//
//	cpf := svc.Generate(cpf.WithPrettyFormat())
func WithPrettyFormat() func(*generatorOptions) {
	return func(opts *generatorOptions) {
		opts.prettyFormat = true
	}
}
