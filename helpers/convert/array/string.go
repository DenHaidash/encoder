package array

import "strconv"

func StringArrToBytesArr(s []string) ([]byte, error) {
	results := make([]byte, len(s))

	for i, str := range s {
		num, err := strconv.Atoi(str)

		if err != nil {
			return nil, err
		}

		results[i] = byte(num)
	}

	return results, nil
}
