package wallet

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestImportHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		lotusClient := util.NewLotusClient("https://api.node.glif.io/rpc/v0", "")

		t.Run("success", func(t *testing.T) {
			w, err := Default.ImportHandler(ctx, db, lotusClient, ImportRequest{
				PrivateKey: testutil.TestPrivateKeyHex,
			})
			require.NoError(t, err)
			require.Equal(t, testutil.TestWalletAddr, w.Address)

			w, err = Default.ImportHandler(ctx, db, lotusClient, ImportRequest{
				PrivateKey: testutil.TestPrivateKeyHex,
			})
			require.ErrorIs(t, err, handlererror.ErrDuplicateRecord)
		})

		t.Run("invalid key", func(t *testing.T) {
			_, err := Default.ImportHandler(ctx, db, lotusClient, ImportRequest{
				PrivateKey: "xxxx",
			})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		})

		t.Run("invalid response", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			lotusClient := util.NewLotusClient("http://127.0.0.1", "")
			_, err := Default.ImportHandler(ctx, db, lotusClient, ImportRequest{
				PrivateKey: testutil.TestPrivateKeyHex,
			})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		})
	})
}
