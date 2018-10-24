package array

import "strconv"

func BytesArrToStringArr(b []byte) []string {
	result := make([]string, len(b))

	for i, char := range b {
		result[i] = strconv.Itoa(int(char))
	}

	return result
}
