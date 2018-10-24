package rsa

import (
	"strconv"
	"strings"
)

func Encode(str string) (string, *KeyPair, error) {
	keys, err := generateKeys()

	if err != nil {
		return "", nil, err
	}

	result := make([]string, len(str))

	var b rune = 301

	for j, char := range str {
		var c int32 = 1
		var i int32 = 0
		asciiCode := char + b

		for i < keys.E {
			c *= asciiCode
			c %= keys.N
			i++
		}

		result[j] = strconv.Itoa(int(c))

		b++
	}

	return strings.Join(result, charsSeparator), keys, nil
}
