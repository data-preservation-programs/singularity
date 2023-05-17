package pack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPathWithMostSpace(t *testing.T) {
	assert := assert.New(t)

	paths := []string{"/", "/tmp", "/var"}
	path, err := GetPathWithMostSpace(paths)
	assert.NoError(err)
	assert.NotEmpty(path)
}
