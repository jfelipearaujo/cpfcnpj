package cnpj

import (
	"fmt"
	"regexp"

	"github.com/jfelipearaujo/cpfcnpj/shared"
)

var patternNumbersLetters = regexp.MustCompile(PATTERN_NUMBERS_LETTERS)

type service struct{}

func New() Service {
	return &service{}
}

func (svc *service) IsValid(cnpj string) error {
	cnpj = shared.CleanString(patternNumbersLetters, cnpj, "")

	if len(cnpj) != EXPECTED_LENGTH {
		return ErrInvalidLength
	}

	if !svc.checkDigits(cnpj) {
		return ErrInvalidCnpj
	}

	return nil
}

func (svc *service) Generate(opts ...func(*generatorOptions)) string {
	genOpts := &generatorOptions{
		version: V1,
	}
	for _, opt := range opts {
		opt(genOpts)
	}

	var cnpj []rune
	length := EXPECTED_LENGTH - 2

	switch genOpts.version {
	case V1:
		cnpj = shared.NumbersToRunes(shared.RandomNumbers(length))
	case V2:
		cnpj = shared.RandomLettersAndNumbers(length)
	}

	cnpj = append(cnpj, svc.calculateDigit(cnpj, len(cnpj)))
	cnpj = append(cnpj, svc.calculateDigit(cnpj, len(cnpj)))

	var out string

	for _, digit := range cnpj {
		out += string(digit)
	}

	if genOpts.prettyFormat {
		out = fmt.Sprintf("%s.%s.%s/%s-%s", out[0:2], out[2:5], out[5:8], out[8:12], out[12:14])
	}

	return out
}

func (svc *service) checkDigits(cnpj string) bool {
	return svc.checkDigit(cnpj, 12) && svc.checkDigit(cnpj, 13)
}

func (svc *service) checkDigit(cnpj string, digitIndex int) bool {
	digit := svc.calculateDigit(shared.StringToRuneSlice(cnpj), digitIndex)
	x := rune(cnpj[digitIndex])

	return digit == x
}

func (svc *service) calculateDigit(data []rune, n int) rune {
	weights := 2
	total := 0

	for i := n - 1; i >= 0; i-- {
		total += int(data[i]-'0') * weights
		weights++
		if weights > 9 {
			weights = 2
		}
	}

	rest := total % 11
	digit := 11 - rest

	if digit >= 10 {
		digit = 0
	}

	return '0' + rune(digit)
}
