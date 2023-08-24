package wallet

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
)

func TestImportHandler(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	ctx := context.Background()
	lotusClient := util.NewLotusClient("https://api.node.glif.io/rpc/v0", "")

	t.Run("success", func(t *testing.T) {
		w, err := ImportHandler(ctx, db, lotusClient, ImportRequest{
			PrivateKey: testutil.TestPrivateKeyHex,
		})
		require.NoError(t, err)
		require.Equal(t, testutil.TestWalletAddr, w.Address)

		w, err = ImportHandler(ctx, db, lotusClient, ImportRequest{
			PrivateKey: testutil.TestPrivateKeyHex,
		})
		require.ErrorIs(t, err, handlererror.ErrDuplicateRecord)
	})

	t.Run("invalid key", func(t *testing.T) {
		_, err := ImportHandler(ctx, db, lotusClient, ImportRequest{
			PrivateKey: "xxxx",
		})
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	})

	t.Run("invalid response", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		lotusClient := util.NewLotusClient("http://127.0.0.1", "")
		_, err := ImportHandler(ctx, db, lotusClient, ImportRequest{
			PrivateKey: testutil.TestPrivateKeyHex,
		})
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	})
}
