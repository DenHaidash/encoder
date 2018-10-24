package rc5

import (
	bitsHelper "encoder/helpers/bits"
	"math"
)

const pw uint64 = 0xB7E151628AED2A6B
const qw uint64 = 0x9E3779B97F4A7C15

func extendKey(key string) []uint64 {
	u := w >> 3

	c := len(key) / u

	if len(key)%u > 0 {
		c = len(key)/u + 1
	}

	l := make([]uint64, c)

	for i := len(key) - 1; i >= 0; i-- {
		l[i/u] = bitsHelper.ROL(l[i/u], 8) + uint64(key[i])
	}

	t := 2 * (r + 1)
	s := make([]uint64, t)

	s[0] = pw

	for i := 1; i < t; i++ {
		s[i] = s[i-1] + qw
	}

	var x, y uint64 = 0, 0
	i, j := 0, 0

	n := int32(3 * math.Max(float64(t), float64(c)))

	for k := int32(0); k < n; k++ {
		s[i] = bitsHelper.ROL(s[i]+x+y, 3)
		x = s[i]

		l[j] = bitsHelper.ROL(l[j]+x+y, (x+y)&31)
		y = l[j]

		i = (i + 1) % t
		j = (j + 1) % c
	}

	return s
}
