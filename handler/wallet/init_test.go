package wallet

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestInitHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		lotusClient := util.NewLotusClient("https://api.node.glif.io/rpc/v0", "")

		t.Run("success", func(t *testing.T) {
			err := db.Create(&model.Wallet{
				Address:    testutil.TestWalletAddr,
				PrivateKey: testutil.TestPrivateKeyHex,
			}).Error
			require.NoError(t, err)
			w, err := Default.InitHandler(ctx, db, lotusClient, testutil.TestWalletAddr)
			require.NoError(t, err)
			require.NotEmpty(t, w.Address)
			require.Equal(t, "f1", w.Address[:2])
			require.NotEmpty(t, w.PrivateKey)
		})

		t.Run("uninitialized-address", func(t *testing.T) {
			err := db.Create(&model.Wallet{
				Address: "f100",
			}).Error
			require.NoError(t, err)
			_, err = Default.InitHandler(ctx, db, lotusClient, "f100")
			require.ErrorContains(t, err, "failed to lookup actor ID")
		})

		t.Run("unknown-address", func(t *testing.T) {
			_, err := Default.InitHandler(ctx, db, lotusClient, "unknown-address")
			require.ErrorContains(t, err, "failed to find wallet")
		})
	})
}
