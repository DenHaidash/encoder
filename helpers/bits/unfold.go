package bits

func Uint64ToBytes(a uint64, b *[]byte, position int32) {
	for i := int32(0); i < 7; i++ {
		(*b)[position+i] = byte(a & 0xFF)
		a >>= 8
	}

	(*b)[position+7] = byte(a & 0xFF)
}
