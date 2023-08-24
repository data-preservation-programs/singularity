package wallet

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
)

func TestListHandler(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		err := db.Create(&model.Wallet{}).Error
		require.NoError(t, err)
		wallets, err := ListHandler(ctx, db)
		require.NoError(t, err)
		require.Len(t, wallets, 1)
	})
}
