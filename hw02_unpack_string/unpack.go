package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	// Place your code here.
	if len(strings.TrimSpace(str)) == 0 {
		return "", nil
	}

	if unicode.IsDigit(rune(str[0])) {
		return "", ErrInvalidString
	}

	var b strings.Builder
	for i, v := range str {
		next := i + 1

		if unicode.IsDigit(v) {
			if len(str) != next {
				if unicode.IsDigit(rune(str[next])) {
					return "", ErrInvalidString
				}
			}
			continue
		}

		if len(str) == next {
			b.WriteRune(v)
			continue
		}

		if unicode.IsDigit(rune(str[next])) {
			d := str[next]

			count, _ := strconv.Atoi(string(d))
			calculated := strings.Repeat(string(str[i]), count)
			println(calculated)

			b.WriteString(calculated)
		} else {
			b.WriteRune(v)
		}
	}
	res := b.String()

	return res, nil
}
