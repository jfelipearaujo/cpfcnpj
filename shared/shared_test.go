package shared_test

import (
	"regexp"
	"testing"

	"github.com/jfelipearaujo/cpfcnpj/shared"
)

func TestCleanString(t *testing.T) {
	t.Run("Should replace all occurrences of the pattern with the replaceBy string", func(t *testing.T) {
		// Arrange
		pattern := regexp.MustCompile(`[a-zA-Z]`)
		input := "abcdefghijklmnopqrstuvwxyz"
		replaceBy := "X"

		// Act
		res := shared.CleanString(pattern, input, replaceBy)

		// Assert
		if res != "XXXXXXXXXXXXXXXXXXXXXXXXXX" {
			t.Errorf("Expected result to be %v, got %v", "XXXXXXXXXXXXXXXXXXXXXXXXXX", res)
		}
	})
}

func TestStringToIntSlice(t *testing.T) {
	t.Run("Should convert a string to a slice of ints", func(t *testing.T) {
		// Arrange
		input := "1234567890"

		// Act
		res := shared.StringToIntSlice(input)

		// Assert
		if len(res) != 10 {
			t.Errorf("Expected result to be %v, got %v", 10, len(res))
		}
	})

	t.Run("Should handle when input contains non-numeric characters", func(t *testing.T) {
		// Arrange
		input := "1234567890a"

		// Act
		res := shared.StringToIntSlice(input)

		// Assert
		if len(res) != 10 {
			t.Errorf("Expected result to be %v, got %v", 10, len(res))
		}
	})
}

func TestStringToRuneSlice(t *testing.T) {
	t.Run("Should convert a string to a slice of runes", func(t *testing.T) {
		// Arrange
		input := "abcdefghijklmnopqrstuvwxyz"

		// Act
		res := shared.StringToRuneSlice(input)

		// Assert
		if len(res) != 26 {
			t.Errorf("Expected result to be %v, got %v", 26, len(res))
		}
	})
}

func TestRandomNumbers(t *testing.T) {
	t.Run("Should generate a slice of random numbers", func(t *testing.T) {
		// Arrange
		length := 10

		// Act
		res := shared.RandomNumbers(length)

		// Assert
		if len(res) != length {
			t.Errorf("Expected result to be %v, got %v", length, len(res))
		}
	})
}

func TestRandomLettersAndNumbers(t *testing.T) {
	t.Run("Should generate a slice of random letters and numbers", func(t *testing.T) {
		// Arrange
		length := 10

		// Act
		res := shared.RandomLettersAndNumbers(length)

		// Assert
		if len(res) != length {
			t.Errorf("Expected result to be %v, got %v", length, len(res))
		}
	})
}
