package cnpj_test

import (
	"errors"
	"testing"

	"github.com/jfelipearaujo/cpfcnpj/cnpj"
)

func TestIsValid(t *testing.T) {
	t.Run("Should return no error if CNPJ is valid", func(t *testing.T) {
		// Arrange
		input := "59.541.264/0001-03"

		sut := cnpj.New(input)

		// Act
		res := sut.IsValid()

		// Assert
		if res != nil {
			t.Errorf("Expected no error, got %v", res)
		}
	})

	t.Run("Should return error if CNPJ is invalid", func(t *testing.T) {
		// Arrange
		input := "59.541.264/0004-03"

		sut := cnpj.New(input)

		// Act
		res := sut.IsValid()

		// Assert
		if !errors.Is(res, cnpj.ErrInvalidCnpj) {
			t.Errorf("Expected error to be %v, got %v", cnpj.ErrInvalidCnpj, res)
		}
	})

	t.Run("Should return error if CNPJ is smaller than 14 numbers", func(t *testing.T) {
		// Arrange
		input := "59.541.264/0004"

		sut := cnpj.New(input)

		// Act
		res := sut.IsValid()

		// Assert
		if !errors.Is(res, cnpj.ErrInvalidLength) {
			t.Errorf("Expected error to be %v, got %v", cnpj.ErrInvalidLength, res)
		}
	})

	t.Run("Should return error if CNPJ is bigger than 14 numbers", func(t *testing.T) {
		// Arrange
		input := "59.541.264/0004-123"

		sut := cnpj.New(input)

		// Act
		res := sut.IsValid()

		// Assert
		if !errors.Is(res, cnpj.ErrInvalidLength) {
			t.Errorf("Expected error to be %v, got %v", cnpj.ErrInvalidLength, res)
		}
	})

	t.Run("Should return no error if CNPJ (new version) is valid", func(t *testing.T) {
		// Arrange
		input := "12.ABC.345/01DE-35"

		sut := cnpj.New(input)

		// Act
		res := sut.IsValid()

		// Assert
		if res != nil {
			t.Errorf("Expected no error, got %v", res)
		}
	})

	t.Run("Should return error if CNPJ (new version) is invalid", func(t *testing.T) {
		// Arrange
		input := "12.ABC.345/01DF-35"

		sut := cnpj.New(input)

		// Act
		res := sut.IsValid()

		// Assert
		if !errors.Is(res, cnpj.ErrInvalidCnpj) {
			t.Errorf("Expected error to be %v, got %v", cnpj.ErrInvalidCnpj, res)
		}
	})
}
