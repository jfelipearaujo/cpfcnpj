package cpf

import "errors"

var (
	ErrInvalidLength = errors.New("invalid length, expected 11 numbers")
	ErrInvalidCpf    = errors.New("invalid cpf")
)
