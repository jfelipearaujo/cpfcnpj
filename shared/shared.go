package shared

import (
	"math/rand"
	"regexp"
	"strconv"
)

func CleanString(pattern *regexp.Regexp, input string, replaceBy string) string {
	return pattern.ReplaceAllString(input, replaceBy)
}

func StringToIntSlice(data string) []int {
	res := make([]int, 0)
	for _, d := range data {
		x, err := strconv.Atoi(string(d))
		if err != nil {
			continue
		}
		res = append(res, x)
	}
	return res
}

func StringToRuneSlice(data string) []rune {
	res := make([]rune, 0)
	for _, d := range data {
		res = append(res, d)
	}
	return res
}

var numberCharset = []rune("0123456789")
var letterNumberCharset = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomNumbers(length int) []int {
	out := make([]int, length)
	for i := range out {
		out[i] = int(numberCharset[rand.Intn(len(numberCharset))] - '0')
	}
	return out
}

func NumbersToRunes(data []int) []rune {
	out := make([]rune, len(data))
	for i, d := range data {
		out[i] = rune(d + '0')
	}
	return out
}

func RandomLettersAndNumbers(length int) []rune {
	out := make([]rune, length)
	for i := range out {
		out[i] = letterNumberCharset[rand.Intn(len(letterNumberCharset))]
	}
	return out
}
