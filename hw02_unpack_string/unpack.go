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
	var isEscaped bool

	for i, item := range input {
		if unicode.IsDigit(item) && i == 0 {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(item) && unicode.IsDigit(rune(input[i-1])) && input[i-2] != '\\' {
			return "", ErrInvalidString
		}

		if unicode.IsLetter(item) && isEscaped {
			return "", ErrInvalidString
		}

		if item == '\\' && !isEscaped {
			isEscaped = true
			continue
		}

		if isEscaped {
			result.WriteRune(item)
			isEscaped = false
			continue
		}

		if unicode.IsDigit(item) {
			count, _ := strconv.Atoi(string(item))
			if count == 0 {

				var str = result.String()
				result.Reset()
				result.WriteString(str[:len(str)-1])

				continue
			}

			result.WriteString(strings.Repeat(string(input[i-1]), count-1))
			continue
		}

		result.WriteRune(item)
	}

	return result.String(), nil
}
