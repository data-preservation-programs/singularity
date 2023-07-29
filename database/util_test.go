package database

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAutoMigrate(t *testing.T) {
	db, closer, err := OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	require.NotNil(t, db)
}
