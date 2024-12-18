package cnpj

import "regexp"

var patternNumbersLetters = regexp.MustCompile(PATTERN_NUMBERS_LETTERS)

type service struct {
	CNPJ string
}

func New(cnpj string) Service {
	return &service{
		CNPJ: cnpj,
	}
}

func (svc *service) IsValid() error {
	svc.clean()

	if len(svc.CNPJ) != EXPECTED_LENGTH {
		return ErrInvalidLength
	}

	firstVerificationDigit := svc.calculateDigitFromFirstVersion(true)
	secondVerificationDigit := svc.calculateDigitFromFirstVersion(false)

	if firstVerificationDigit != int(svc.CNPJ[12]-'0') || secondVerificationDigit != int(svc.CNPJ[13]-'0') {
		return ErrInvalidCnpj
	}

	return nil
}

func (svc *service) clean() {
	svc.CNPJ = patternNumbersLetters.ReplaceAllString(svc.CNPJ, "")
}

func (svc *service) calculateDigitFromFirstVersion(isFirst bool) int {
	maxIndex := 12
	weights := 2

	if !isFirst {
		maxIndex = 13
	}

	total := 0

	for i := maxIndex; i > 0; i-- {
		total += int(svc.CNPJ[i-1]-'0') * weights
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

	return digit
}
