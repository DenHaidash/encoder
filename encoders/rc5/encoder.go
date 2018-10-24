package rc5

import (
	bitsHelper "encoder/helpers/bits"
	arrayConv "encoder/helpers/convert/array"
	mathHelper "encoder/helpers/math"
	strHelper "encoder/helpers/string"
	"math"
	"strings"
)

const blockSeparator = "\n"
const charSeparator = " "

const w = 64 // machine word, 8 bytes -> block 16 byte
const r = 16 // number of rounds: 0...255

var blockSize int32

func init() {
	blockSize = w * 2 / 8 // bytes - 128 bits
}

func Encode(str string, key string) string {
	str = fixInputStrLength(str)

	numOfBlocks := int32(math.Ceil(float64(len(str)) / float64(blockSize)))

	extKey := extendKey(key)

	blocks := make([]string, 0, numOfBlocks)

	for i := int32(0); i < numOfBlocks; i++ {
		block := str[i*blockSize : i*blockSize+blockSize]

		encodedBlockArr := arrayConv.BytesArrToStringArr(encodeBlock(block, extKey))
		blocks = append(blocks, strings.Join(encodedBlockArr, charSeparator))
	}

	return strings.Join(blocks, blockSeparator)
}

func encodeBlock(block string, s []uint64) []byte {
	a := bitsHelper.BytesToUint64([]byte(block), 0)
	b := bitsHelper.BytesToUint64([]byte(block), blockSize/2)

	a += s[0]
	b += s[1]

	for i := 1; i < r+1; i++ {
		a = bitsHelper.ROL(a^b, b&63) + s[2*i]
		b = bitsHelper.ROL(b^a, a&63) + s[2*i+1]
	}

	result := make([]byte, len(block))
	bitsHelper.Uint64ToBytes(a, &result, 0)
	bitsHelper.Uint64ToBytes(b, &result, blockSize/2)

	return result
}

func fixInputStrLength(str string) string {
	minLength := math.Max(float64(blockSize), float64(len(str)))
	return strHelper.PadRight(str, ' ', int(mathHelper.GetUpperPowerOf2(int32(minLength))))
}
