package database

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type TestModel struct {
	ID   uint   `gorm:"primarykey"`
	Name string
}

func TestModernCDriver(t *testing.T) {
	db, closer, err := OpenDatabase("sqlite:file::memory:?cache=shared", &gorm.Config{})
	require.NoError(t, err)
	defer closer.Close()

	// Create test table
	err = db.AutoMigrate(&TestModel{})
	require.NoError(t, err)

	// Test insert
	result := db.Create(&TestModel{Name: "test"})
	require.NoError(t, result.Error)
	require.Equal(t, int64(1), result.RowsAffected)

	// Test query
	var model TestModel
	result = db.First(&model)
	require.NoError(t, result.Error)
	require.Equal(t, "test", model.Name)
}
