package bits

func ROR(a uint64, offset uint64) uint64 {
	return a>>offset | a<<(64-offset)
}
