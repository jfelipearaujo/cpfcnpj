package cpf

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jfelipearaujo/cpfcnpj/shared"
)

var patternOnlyNumbers = regexp.MustCompile(PATTERN_ONLY_NUMBERS)

type service struct{}

func New() Service {
	return &service{}
}

func (svc *service) IsValid(cpf string) error {
	cpf = shared.CleanString(patternOnlyNumbers, cpf, "")

	if len(cpf) != EXPECTED_LENGTH {
		return ErrInvalidLength
	}

	if !svc.checkDigits(cpf) {
		return ErrInvalidCpf
	}

	return nil
}

func (svc *service) Generate(opts ...func(*generatorOptions)) string {
	genOpts := &generatorOptions{}
	for _, opt := range opts {
		opt(genOpts)
	}

	cpf := shared.RandomNumbers(EXPECTED_LENGTH - 2)
	cpf = append(cpf, svc.calculateDigit(cpf, len(cpf)))
	cpf = append(cpf, svc.calculateDigit(cpf, len(cpf)))

	var out string

	for _, digit := range cpf {
		out += strconv.Itoa(digit)
	}

	if genOpts.prettyFormat {
		out = fmt.Sprintf("%s.%s.%s-%s", out[0:3], out[3:6], out[6:9], out[9:11])
	}

	return out
}

func (svc *service) checkDigits(cpf string) bool {
	return svc.checkDigit(cpf, 9) && svc.checkDigit(cpf, 10)
}

func (svc *service) checkDigit(cpf string, digitIndex int) bool {
	digit := svc.calculateDigit(shared.StringToIntSlice(cpf), digitIndex)

	x, err := strconv.Atoi(string(cpf[digitIndex]))
	if err != nil {
		return false
	}
	return digit == x
}

func (svc *service) calculateDigit(data []int, n int) int {
	var total int
	for i := range n {
		total += data[i] * (n + 1 - i)
	}
	total = total % 11
	if total < 2 {
		return 0
	}
	return 11 - total
}
