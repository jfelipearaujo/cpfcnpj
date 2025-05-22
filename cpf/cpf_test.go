package cpf_test

import (
	"errors"
	"testing"

	"github.com/jfelipearaujo/cpfcnpj/cpf"
)

func TestIsValid(t *testing.T) {
	t.Run("Should return no error if CPF is valid", func(t *testing.T) {
		// Arrange
		input := "145.382.206-20"

		sut := cpf.New(input)

		// Act
		res := sut.IsValid()

		// Assert
		if res != nil {
			t.Errorf("Expected no error, got %v", res)
		}
	})

	t.Run("Should return error if CPF is invalid", func(t *testing.T) {
		// Arrange
		input := "123.456.789-10"

		sut := cpf.New(input)

		// Act
		res := sut.IsValid()

		// Assert
		if !errors.Is(res, cpf.ErrInvalidCpf) {
			t.Errorf("Expected error to be %v, got %v", cpf.ErrInvalidCpf, res)
		}
	})

	t.Run("Should return error if CPF is smaller than 11 numbers", func(t *testing.T) {
		// Arrange
		input := "123.456.789"

		sut := cpf.New(input)

		// Act
		res := sut.IsValid()

		// Assert
		if !errors.Is(res, cpf.ErrInvalidLength) {
			t.Errorf("Expected error to be %v, got %v", cpf.ErrInvalidLength, res)
		}
	})

	t.Run("Should return error if CPF is bigger than 11 numbers", func(t *testing.T) {
		// Arrange
		input := "123.456.789-123"

		sut := cpf.New(input)

		// Act
		res := sut.IsValid()

		// Assert
		if !errors.Is(res, cpf.ErrInvalidLength) {
			t.Errorf("Expected error to be %v, got %v", cpf.ErrInvalidLength, res)
		}
	})
}
