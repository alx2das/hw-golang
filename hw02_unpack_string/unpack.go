package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	// пустая строка - коррект
	if input == "" {
		return "", nil
	}

	// начало с цифры - ошибка
	if unicode.IsDigit(rune(input[0])) {
		return "", ErrInvalidString
	}

	var result strings.Builder
	var isEscaped bool

	for i, item := range input {
		// если пытаемся повторить неэкранированную цифру
		if unicode.IsDigit(item) && unicode.IsDigit(rune(input[i-1])) && input[i-2] != '\\' {
			return "", ErrInvalidString
		}

		// если пытаемся повторить букву
		if unicode.IsLetter(item) && isEscaped {
			return "", ErrInvalidString
		}

		// начало экранирование
		if item == '\\' && !isEscaped {
			isEscaped = true
			continue
		}

		// записываем экранированный символ
		if isEscaped {
			result.WriteRune(item)
			isEscaped = false
			continue
		}

		// повторим предыдущий символ `count` раз
		if unicode.IsDigit(item) {
			count, _ := strconv.Atoi(string(item))

			// удалим предыдущий символ
			if count == 0 {
				str := result.String()
				result.Reset()
				result.WriteString(str[:len(str)-1])
				continue
			}

			// дублируем предыдущий символ `count` раз
			result.WriteString(strings.Repeat(string(input[i-1]), count-1))
			continue
		}

		result.WriteRune(item)
	}

	return result.String(), nil
}
