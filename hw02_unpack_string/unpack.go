package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

// Unpack function unpacking of string.
func Unpack(str string) (string, error) {
	var value string
	var b strings.Builder

	for _, v := range str {
		if unicode.IsDigit(v) {
			if value == "" {
				return "", ErrInvalidString
			}

			d, _ := strconv.Atoi(string(v))
			b.WriteString(strings.Repeat(value, d))
			value = ""

			continue
		}

		b.WriteString(value)
		value = string(v)
	}

	b.WriteString(value)

	return b.String(), nil
}
