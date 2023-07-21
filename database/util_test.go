package database

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAutoMigrate(t *testing.T) {
	db, err := OpenInMemory()
	require.NoError(t, err)
	require.NotNil(t, db)
}
