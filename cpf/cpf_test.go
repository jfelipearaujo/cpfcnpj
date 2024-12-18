package cpf_test

import (
	"testing"

	"github.com/jfelipearaujo/cpfcnpj/cpf"
	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	t.Run("Should return no error if CPF is valid", func(t *testing.T) {
		// Arrange
		input := "145.382.206-20"

		sut := cpf.New(input)

		// Act
		res := sut.IsValid()

		// Assert
		assert.NoError(t, res)
	})

	t.Run("Should return error if CPF is invalid", func(t *testing.T) {
		// Arrange
		input := "123.456.789-10"

		sut := cpf.New(input)

		// Act
		res := sut.IsValid()

		// Assert
		assert.Error(t, res)
		assert.ErrorIs(t, res, cpf.ErrInvalidCpf)
	})

	t.Run("Should return error if CPF is smaller than 11 numbers", func(t *testing.T) {
		// Arrange
		input := "123.456.789"

		sut := cpf.New(input)

		// Act
		res := sut.IsValid()

		// Assert
		assert.Error(t, res)
		assert.ErrorIs(t, res, cpf.ErrInvalidLength)
	})

	t.Run("Should return error if CPF is bigger than 11 numbers", func(t *testing.T) {
		// Arrange
		input := "123.456.789-123"

		sut := cpf.New(input)

		// Act
		res := sut.IsValid()

		// Assert
		assert.Error(t, res)
		assert.ErrorIs(t, res, cpf.ErrInvalidLength)
	})
}
