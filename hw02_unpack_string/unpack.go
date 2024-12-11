package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var result strings.Builder
	var prev string

	for _, str := range input {
		if count, err := strconv.Atoi(string(str)); err == nil {
			if prev == "" {
				return "", ErrInvalidString
			}

			result.WriteString(strings.Repeat(prev, count))
			prev = ""
		} else {
			result.WriteString(prev)
			prev = string(str)
		}
	}

	result.WriteString(prev)

	return result.String(), nil
}
