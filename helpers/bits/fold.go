package bits

func BytesToUint64(b []byte, position int32) uint64 {
	var r uint64 = 0

	for i := position + 7; i > position; i-- {
		r |= uint64(b[i])
		r <<= 8
	}

	r |= uint64(b[position])

	return r
}
