package database

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAutoMigrate(t *testing.T) {
	db := OpenInMemory()
	require.NotNil(t, db)
}
