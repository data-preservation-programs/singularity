package testutil

import (
	"context"
	"testing"

	"gorm.io/gorm"
)

func TestTestDB(t *testing.T) {
	All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {})
}
