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
	//	cnpj := svc.Generate()
	//	fmt.Println(cnpj) // 12ABC34501DE35
	Generate(opts ...func(*generatorOptions)) string
}

type generatorOptions struct {
	version      Version
	prettyFormat bool
}

// Sets the version of the CNPJ when generating a new one
//
// # Note: This function by default will generate a CNPJ with the version V1
//
// Example:
//
//	cnpj := svc.Generate(cnpj.WithVersion(cnpj.V2))
func WithVersion(version Version) func(*generatorOptions) {
	return func(opts *generatorOptions) {
		opts.version = version
	}
}

// Sets the pretty format of the CNPJ when generating a new one
//
// Example:
//
//	cnpj := svc.Generate(cnpj.WithPrettyFormat())
func WithPrettyFormat() func(*generatorOptions) {
	return func(opts *generatorOptions) {
		opts.prettyFormat = true
	}
}
