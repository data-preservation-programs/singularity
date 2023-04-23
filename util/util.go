package util

func NextPowerOfTwo(x uint64) uint64 {
	if x == 0 {
		return 1
	}

	// Find the position of the highest bit set to 1
	pos := uint(0)
	for shifted := x; shifted > 0; shifted >>= 1 {
		pos++
	}

	// If x is already a power of two, return x
	if x == 1<<(pos-1) {
		return x
	}

	// Otherwise, return the next power of two
	return 1 << pos
}
