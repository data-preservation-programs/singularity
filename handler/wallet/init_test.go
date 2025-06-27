package wallet

import (
	"context"
	"errors"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestInitHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Create mock client for success case
		successMockClient := testutil.NewMockLotusClient()
		successMockClient.SetResponse("Filecoin.StateLookupID", testutil.TestWalletActorID)

		t.Run("success", func(t *testing.T) {
			err := db.Create(&model.Wallet{
				Address:    testutil.TestWalletAddr,
				PrivateKey: testutil.TestPrivateKeyHex,
			}).Error
			require.NoError(t, err)
			w, err := Default.InitHandler(ctx, db, successMockClient, testutil.TestWalletAddr)
			require.NoError(t, err)
			require.NotEmpty(t, w.PrivateKey)
			require.Equal(t, w.Address, testutil.TestWalletAddr)
			require.NotEmpty(t, w.ActorID)

			// Running again on an initialized wallet should not change the wallet
			w2, err := Default.InitHandler(ctx, db, successMockClient, testutil.TestWalletAddr)
			require.NoError(t, err)
			require.Equal(t, w.ActorID, w2.ActorID)
		})

		t.Run("uninitialized-address", func(t *testing.T) {
			// Create mock client that returns an error for uninitialized address
			errorMockClient := testutil.NewMockLotusClient()
			errorMockClient.SetError("Filecoin.StateLookupID", errors.New("actor not found"))

			err := db.Create(&model.Wallet{
				Address: "f100",
			}).Error
			require.NoError(t, err)
			_, err = Default.InitHandler(ctx, db, errorMockClient, "f100")
			require.ErrorContains(t, err, "failed to lookup actor ID")
		})

		t.Run("unknown-address", func(t *testing.T) {
			_, err := Default.InitHandler(ctx, db, successMockClient, "unknown-address")
			require.ErrorContains(t, err, "failed to find wallet")
		})
	})
}
