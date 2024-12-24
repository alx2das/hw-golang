package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var result strings.Builder
	var previous rune
	var shielding bool

	for _, current := range []rune(input) {
		if shielding {
			result.WriteRune(current)

			previous = current
			shielding = false

			continue
		}

		if current == '\\' {
			shielding = true
			continue
		}

		if unicode.IsDigit(current) {
			count, _ := strconv.Atoi(string(current))

			if previous == 0 {
				return "", ErrInvalidString
			}

			if count == 0 {
				str := removeLastChar(result.String())
				result.Reset()
				result.WriteString(str)

				continue
			}

			result.WriteString(strings.Repeat(string(previous), count-1))
			previous = 0

			continue
		}

		previous = current
		result.WriteRune(current)
	}

	return result.String(), nil
}

func removeLastChar(input string) string {
	if len(input) == 0 {
		return input
	}

	r := []rune(input)
	return string(r[:len(r)-1])
}
