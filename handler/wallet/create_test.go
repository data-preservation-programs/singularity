package wallet

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestCreateHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		lotusClient := util.NewLotusClient("https://api.node.glif.io/rpc/v0", "")

		t.Run("success-secp256k1", func(t *testing.T) {
			w, err := Default.CreateHandler(ctx, db, CreateRequest{KeyType: KTSecp256k1.String()})
			require.NoError(t, err)
			require.NotEmpty(t, w.Address)
			require.Equal(t, "f1", w.Address[:2])
			require.Equal(t, "hello", w.PrivateKey)

			_, err = Default.ImportHandler(ctx, db, lotusClient, ImportRequest{
				PrivateKey: w.PrivateKey,
			})
			require.ErrorIs(t, err, handlererror.ErrDuplicateRecord)
		})

		t.Run("success-bls", func(t *testing.T) {
			w, err := Default.CreateHandler(ctx, db, CreateRequest{KeyType: KTBLS.String()})
			require.NoError(t, err)
			require.NotEmpty(t, w.Address)
			require.Equal(t, "f3", w.Address[:2])

			_, err = Default.ImportHandler(ctx, db, lotusClient, ImportRequest{
				PrivateKey: w.PrivateKey,
			})
			require.ErrorIs(t, err, handlererror.ErrDuplicateRecord)
		})

		t.Run("invalid-key-type", func(t *testing.T) {
			_, err := Default.CreateHandler(ctx, db, CreateRequest{KeyType: "invalid-type"})
			require.Error(t, err)
		})
	})
}
