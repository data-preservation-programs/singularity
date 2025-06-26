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

			_, err = Default.ImportHandler(ctx, db, lotusClient, ImportRequest{
				PrivateKey: testutil.TestPrivateKeyHex,
			})
			require.ErrorIs(t, err, handlererror.ErrDuplicateRecord)
		})

		t.Run("invalid key", func(t *testing.T) {
			_, err := Default.ImportHandler(ctx, db, lotusClient, ImportRequest{
				PrivateKey: "7b2254797065223a22736563703235366b31222c22507269766174654b6579223a22414141414141414141414141414141414141414141414141414141414141414141414141414141414141413d227d", // Valid hex, valid base64, but all zeros private key
			})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		})

		t.Run("invalid response", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
			defer cancel()
			lotusClient := util.NewLotusClient("http://invalid-url-that-does-not-exist.local", "")
			_, err := Default.ImportHandler(ctx, db, lotusClient, ImportRequest{
				PrivateKey: testutil.TestPrivateKeyHex,
			})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		})
	})
}
