package cnpj

type Version string

const (
	V1 Version = "06-1998"
	V2 Version = "07-2026"
)

const (
	PATTERN_NUMBERS_LETTERS string = "[^0-9A-Z]"
	EXPECTED_LENGTH         int    = 14
)
