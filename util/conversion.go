package util

import "math"

// SafeInt64ToUint64 safely converts int64 to uint64, handling negative values
func SafeInt64ToUint64(val int64) uint64 {
	if val < 0 {
		return 0
	}
	return uint64(val)
}

// SafeUint64ToInt64 safely converts uint64 to int64, ensuring no overflow
func SafeUint64ToInt64(val uint64) int64 {
	if val > math.MaxInt64 {
		return math.MaxInt64
	}
	return int64(val)
}