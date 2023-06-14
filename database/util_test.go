package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAutoMigrate(t *testing.T) {
	db := OpenInMemory()
	assert.NotNil(t, db)
}
