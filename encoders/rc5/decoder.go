package rc5

import (
	bitsHelper "encoder/helpers/bits"
	arrayConv "encoder/helpers/convert/array"
	"strings"
)

func Decode(str string, key string) (string, error) {
	results := make([]string, 0)

	extKey := extendKey(key)

	for _, block := range strings.Split(str, blockSeparator) {
		bytesBlock, err := arrayConv.StringArrToBytesArr(strings.Split(block, charSeparator))

		if err != nil {
			return "", err
		}

		decodedBlock := decodeBlock(bytesBlock, extKey)
		results = append(results, string(decodedBlock))
	}

	return strings.TrimSpace(strings.Join(results, "")), nil
}

func decodeBlock(block []byte, s []uint64) []byte {
	a := bitsHelper.BytesToUint64(block, 0)
	b := bitsHelper.BytesToUint64(block, blockSize/2)

	for i := r; i > 0; i-- {
		b = bitsHelper.ROR(b-s[2*i+1], a&63) ^ a
		a = bitsHelper.ROR(a-s[2*i], b&63) ^ b
	}

	b -= s[1]
	a -= s[0]

	result := make([]byte, len(block))
	bitsHelper.Uint64ToBytes(a, &result, 0)
	bitsHelper.Uint64ToBytes(b, &result, blockSize/2)

	return result
}
