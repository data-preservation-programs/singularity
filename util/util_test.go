package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextPowerOfTwo(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(uint64(1), NextPowerOfTwo(0))
	assert.Equal(uint64(1), NextPowerOfTwo(1))
	assert.Equal(uint64(2), NextPowerOfTwo(2))
	assert.Equal(uint64(4), NextPowerOfTwo(3))
	assert.Equal(uint64(4), NextPowerOfTwo(4))
	assert.Equal(uint64(8), NextPowerOfTwo(5))
	assert.Equal(uint64(16), NextPowerOfTwo(9))
	assert.Equal(uint64(32), NextPowerOfTwo(17))
	assert.Equal(uint64(64), NextPowerOfTwo(33))
}
