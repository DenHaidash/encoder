package string

func PadRight(str string, suffix rune, length int) string {
	for len(str) < length {
		str += string(suffix)
	}

	return str
}
