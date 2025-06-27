package wallet

import (
	"context"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestImportHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		t.Run("success", func(t *testing.T) {
			mockClient := testutil.NewMockLotusClient()
			mockClient.SetResponse("Filecoin.StateLookupID", testutil.TestWalletActorID)

			w, err := Default.ImportHandler(ctx, db, mockClient, ImportRequest{
				PrivateKey: testutil.TestPrivateKeyHex,
			})
			require.NoError(t, err)
			require.Equal(t, testutil.TestWalletAddr, w.Address)

			_, err = Default.ImportHandler(ctx, db, mockClient, ImportRequest{
				PrivateKey: testutil.TestPrivateKeyHex,
			})
			require.ErrorIs(t, err, handlererror.ErrDuplicateRecord)
		})

		t.Run("invalid key", func(t *testing.T) {
			mockClient := testutil.NewMockLotusClient()
			// Mock the RPC call to return "actor not found" for the invalid key
			mockClient.SetError("Filecoin.StateLookupID", errors.New("3: actor not found"))
			_, err := Default.ImportHandler(ctx, db, mockClient, ImportRequest{
				PrivateKey: "7b2254797065223a22736563703235366b31222c22507269766174654b6579223a22414141414141414141414141414141414141414141414141414141414141414141414141414141414141413d227d", // Valid hex, valid base64, but all zeros private key
			})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		})

		t.Run("invalid response", func(t *testing.T) {
			mockClient := testutil.NewMockLotusClient()
			mockClient.SetError("Filecoin.StateLookupID", errors.New("rpc call failed"))
			_, err := Default.ImportHandler(ctx, db, mockClient, ImportRequest{
				PrivateKey: testutil.TestPrivateKeyHex,
			})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		})
	})
}
