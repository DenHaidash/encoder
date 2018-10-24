package rsa

import (
	"strconv"
	"strings"
)

const charsSeparator = " "

func Decode(str string, key string) (string, error) {
	n, d, err := splitKey(key)

	if err != nil {
		return "", err
	}

	result := make([]rune, 0)

	var b int32 = 301

	for _, charStr := range strings.Split(str, charsSeparator) {
		var m int32 = 1
		var i int32 = 0

		char, err := strconv.Atoi(charStr)

		if err != nil {
			return "", err
		}

		for i < d {
			m *= int32(char)
			m %= n

			i++
		}

		m -= b

		result = append(result, rune(m))

		b++
	}

	return string(result), nil
}
