package wallet

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
		t.Run("success", func(t *testing.T) {
			err := db.Create(&model.Wallet{}).Error
			require.NoError(t, err)
			wallets, err := Default.ListHandler(ctx, db)
			require.NoError(t, err)
			require.Len(t, wallets, 1)
		})
	})
}
