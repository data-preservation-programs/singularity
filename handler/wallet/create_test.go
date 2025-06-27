package wallet

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestCreateHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Create mock client for all tests
		mockClient := testutil.NewMockLotusClient()
		mockClient.SetResponse("Filecoin.StateLookupID", testutil.TestWalletActorID)

		t.Run("success-secp256k1", func(t *testing.T) {
			w, err := Default.CreateHandler(ctx, db, mockClient, CreateRequest{
				KeyType: KTSecp256k1.String(),
			})
			require.NoError(t, err)
			require.Equal(t, "UserWallet", string(w.WalletType))
		})

		t.Run("success-user-wallet-secp256k1", func(t *testing.T) {
			w, err := Default.CreateHandler(ctx, db, mockClient, CreateRequest{
				KeyType: KTSecp256k1.String(),
			})
			require.NoError(t, err)
			require.NotEmpty(t, w.Address)
			require.Equal(t, "f1", w.Address[:2])
			require.NotEmpty(t, w.PrivateKey)
		})

		t.Run("success-bls", func(t *testing.T) {
			w, err := Default.CreateHandler(ctx, db, mockClient, CreateRequest{
				KeyType: KTBLS.String(),
			})
			require.NoError(t, err)
			require.Equal(t, "UserWallet", string(w.WalletType))
		})

		t.Run("success-user-wallet-bls", func(t *testing.T) {
			w, err := Default.CreateHandler(ctx, db, mockClient, CreateRequest{
				KeyType: KTBLS.String(),
			})
			require.NoError(t, err)
			require.NotEmpty(t, w.Address)
			require.Equal(t, "f3", w.Address[:2])
			require.NotEmpty(t, w.PrivateKey)
		})

		t.Run("invalid-key-type", func(t *testing.T) {
			_, err := Default.CreateHandler(ctx, db, mockClient, CreateRequest{
				KeyType: "invalid-type",
			})
			require.Error(t, err)
		})

		t.Run("success-user-wallet-with-details", func(t *testing.T) {
			w, err := Default.CreateHandler(ctx, db, mockClient, CreateRequest{
				KeyType: KTSecp256k1.String(),
				Name:    "my wallet",
			})
			require.NoError(t, err)
			require.NotEmpty(t, w.Address)
			require.Equal(t, "f1", w.Address[:2])
			require.NotEmpty(t, w.PrivateKey)
			require.Equal(t, "UserWallet", string(w.WalletType))
			require.Equal(t, "my wallet", w.ActorName)
		})

		t.Run("success-sp-wallet", func(t *testing.T) {
			w, err := Default.CreateHandler(ctx, db, mockClient, CreateRequest{
				Address:  testutil.TestWalletAddr,
				ActorID:  testutil.TestWalletActorID,
				Name:     "Test SP",
				Contact:  "test@example.com",
				Location: "US",
			})
			require.NoError(t, err)
			require.Equal(t, testutil.TestWalletAddr, w.Address)
			require.Equal(t, testutil.TestWalletActorID, w.ActorID)
			require.Equal(t, "Test SP", w.ActorName)
			require.Equal(t, "test@example.com", w.ContactInfo)
			require.Equal(t, "US", w.Location)
			require.Empty(t, w.PrivateKey)
			require.Equal(t, "SPWallet", string(w.WalletType))
		})

		t.Run("error-no-parameters", func(t *testing.T) {
			_, err := Default.CreateHandler(ctx, db, mockClient, CreateRequest{})
			require.Error(t, err)
			require.Contains(t, err.Error(), "must specify either KeyType (for UserWallet) or Address/ActorID (for SPWallet)")
		})

		t.Run("error-sp-wallet-missing-actorid", func(t *testing.T) {
			_, err := Default.CreateHandler(ctx, db, mockClient, CreateRequest{
				Address: "f123456789",
				Name:    "Test SP",
			})
			require.Error(t, err)
			require.Contains(t, err.Error(), "must specify both Address and ActorID (for SPWallet)")
		})

		t.Run("error-sp-wallet-missing-address", func(t *testing.T) {
			_, err := Default.CreateHandler(ctx, db, mockClient, CreateRequest{
				ActorID: "f1234",
				Name:    "Test SP",
			})
			require.Error(t, err)
			require.Contains(t, err.Error(), "must specify both Address and ActorID (for SPWallet)")
		})

		t.Run("error-sp-wallet-mismatched-id", func(t *testing.T) {
			// Create a mock client that returns a different valid actor ID
			mismatchMockClient := testutil.NewMockLotusClient()
			mismatchMockClient.SetResponse("Filecoin.StateLookupID", "f0123456")

			_, err := Default.CreateHandler(ctx, db, mismatchMockClient, CreateRequest{
				Address: testutil.TestWalletAddr,
				ActorID: "f0999999",
			})
			require.Error(t, err)
			require.Contains(t, err.Error(), "provided actor ID is not associated with address")
		})

		t.Run("error-mixed-parameters", func(t *testing.T) {
			_, err := Default.CreateHandler(ctx, db, mockClient, CreateRequest{
				KeyType: KTSecp256k1.String(),
				Address: "f3abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890",
			})
			require.Error(t, err)
			require.Contains(t, err.Error(), "cannot specify both KeyType (for UserWallet) and Address/ActorID (for SPWallet)")
		})
	})
}
