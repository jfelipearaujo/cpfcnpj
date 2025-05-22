package cnpj_test

import (
	"errors"
	"regexp"
	"testing"

	"github.com/jfelipearaujo/cpfcnpj/cnpj"
)

func TestIsValid(t *testing.T) {
	t.Run("Should return no error if CNPJ is valid", func(t *testing.T) {
		// Arrange
		input := "59.541.264/0001-03"

		sut := cnpj.New()

		// Act
		res := sut.IsValid(input)

		// Assert
		if res != nil {
			t.Errorf("Expected no error, got %v", res)
		}
	})

	t.Run("Should return error if CNPJ is invalid", func(t *testing.T) {
		// Arrange
		input := "59.541.264/0004-03"

		sut := cnpj.New()

		// Act
		res := sut.IsValid(input)

		// Assert
		if !errors.Is(res, cnpj.ErrInvalidCnpj) {
			t.Errorf("Expected error to be %v, got %v", cnpj.ErrInvalidCnpj, res)
		}
	})

	t.Run("Should return error if CNPJ is smaller than 14 numbers", func(t *testing.T) {
		// Arrange
		input := "59.541.264/0004"

		sut := cnpj.New()

		// Act
		res := sut.IsValid(input)

		// Assert
		if !errors.Is(res, cnpj.ErrInvalidLength) {
			t.Errorf("Expected error to be %v, got %v", cnpj.ErrInvalidLength, res)
		}
	})

	t.Run("Should return error if CNPJ is bigger than 14 numbers", func(t *testing.T) {
		// Arrange
		input := "59.541.264/0004-123"

		sut := cnpj.New()

		// Act
		res := sut.IsValid(input)

		// Assert
		if !errors.Is(res, cnpj.ErrInvalidLength) {
			t.Errorf("Expected error to be %v, got %v", cnpj.ErrInvalidLength, res)
		}
	})

	t.Run("Should return no error if CNPJ with letters is valid", func(t *testing.T) {
		// Arrange
		input := "12.ABC.345/01DE-35"

		sut := cnpj.New()

		// Act
		res := sut.IsValid(input)

		// Assert
		if res != nil {
			t.Errorf("Expected no error, got %v", res)
		}
	})

	t.Run("Should return error if CNPJ with letters is invalid", func(t *testing.T) {
		// Arrange
		input := "12.ABC.345/01DF-35"

		sut := cnpj.New()

		// Act
		res := sut.IsValid(input)

		// Assert
		if !errors.Is(res, cnpj.ErrInvalidCnpj) {
			t.Errorf("Expected error to be %v, got %v", cnpj.ErrInvalidCnpj, res)
		}
	})
}

func TestGenerate(t *testing.T) {
	cnpjRegexPretty := regexp.MustCompile(`^[A-Z0-9]{2}\.[A-Z0-9]{3}\.[A-Z0-9]{3}\/[A-Z0-9]{4}\-\d{2}$`)
	cnpjRegexNonPretty := regexp.MustCompile(`^[A-Z0-9]{12}\d{2}$`)

	t.Run("Should return a valid CNPJ in pretty mode", func(t *testing.T) {
		// Arrange
		sut := cnpj.New()

		// Act
		res := sut.Generate(true)

		// Assert
		err := sut.IsValid(res)

		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if !cnpjRegexPretty.MatchString(res) {
			t.Errorf("Expected CNPJ to match the regex, got: %v", res)
		}
	})

	t.Run("Should return a valid CNPJ in non-pretty mode", func(t *testing.T) {
		// Arrange
		sut := cnpj.New()

		// Act
		res := sut.Generate(false)

		// Assert
		err := sut.IsValid(res)

		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if !cnpjRegexNonPretty.MatchString(res) {
			t.Errorf("Expected CNPJ to match the regex, got: %v", res)
		}
	})
}
