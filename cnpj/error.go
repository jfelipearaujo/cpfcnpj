package cnpj

import "errors"

var (
	ErrInvalidLength = errors.New("invalid length, expected 14 numbers")
	ErrInvalidCnpj   = errors.New("invalid cnpj")
)
