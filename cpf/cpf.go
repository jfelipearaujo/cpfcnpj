package cpf

import (
	"regexp"
)

var patternOnlyNumbers = regexp.MustCompile(PATTERN_ONLY_NUMBERS)

type service struct {
	CPF string
}

func New(cpf string) Service {
	return &service{
		CPF: cpf,
	}
}

func (svc *service) IsValid() error {
	svc.clean()

	if len(svc.CPF) != EXPECTED_LENGTH {
		return ErrInvalidLength
	}

	firstVerificationDigit := svc.calculateDigit(true)
	secondVerificationDigit := svc.calculateDigit(false)

	if firstVerificationDigit != int(svc.CPF[9]-'0') || secondVerificationDigit != int(svc.CPF[10]-'0') {
		return ErrInvalidCpf
	}

	return nil
}

func (svc *service) clean() {
	svc.CPF = patternOnlyNumbers.ReplaceAllString(svc.CPF, "")
}

func (svc *service) calculateDigit(isFirst bool) int {
	maxIndex := 9
	weights := 10

	if !isFirst {
		weights = 11
		maxIndex = 10
	}

	total := 0

	for _, c := range svc.CPF[0:maxIndex] {
		total += int(c-'0') * weights
		weights--
	}

	rest := total % 11
	digit := 11 - rest

	if digit >= 10 {
		digit = 0
	}

	return digit
}
