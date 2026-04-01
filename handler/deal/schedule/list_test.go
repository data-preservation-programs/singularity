package schedule

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestListHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		w := model.Wallet{Address: "f01", KeyPath: "/tmp/key", KeyStore: "local"}
		require.NoError(t, db.Create(&w).Error)
		err := db.Create(&model.Preparation{
			WalletID: &w.ID,
		}).Error
		require.NoError(t, err)
		err = db.Create(&model.Schedule{
			PreparationID: 1,
		}).Error
		require.NoError(t, err)
		schedules, err := Default.ListHandler(ctx, db, ListRequest{})
		require.NoError(t, err)
		require.Len(t, schedules, 1)

		// Test filtering by group
		err = db.Model(&model.Schedule{}).Where("id = ?", 1).UpdateColumn("group", "test-group").Error
		require.NoError(t, err)
		filtered, err := Default.ListHandler(ctx, db, ListRequest{Group: "test-group"})
		require.NoError(t, err)
		require.Len(t, filtered, 1)

		empty, err := Default.ListHandler(ctx, db, ListRequest{Group: "nonexistent"})
		require.NoError(t, err)
		require.Len(t, empty, 0)
	})
}
