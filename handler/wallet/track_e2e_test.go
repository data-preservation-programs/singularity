package wallet

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/util"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// TestTrackWalletE2E tests TrackedWallet creation with real Lotus API calls
// This test will be skipped if Lotus API is not available
func TestTrackWalletE2E(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Use real Lotus API with default endpoint
		lotusAPI := "https://api.node.glif.io/rpc/v1"
		lotusClient := util.NewLotusClient(lotusAPI, "")

		// Test with a known client ActorID from mainnet
		// f01131298 is a known client wallet
		knownClientActorID := "f01131298"

		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		// First verify the Lotus API is accessible
		var version interface{}
		err := lotusClient.CallFor(ctx, &version, "Filecoin.Version")
		if err != nil {
			t.Skipf("Skipping e2e test because Lotus API is not available: %v", err)
			return
		}
		t.Logf("Connected to Lotus API, version: %v", version)

		// Test TrackedWallet creation with real API call
		w, err := Default.CreateHandler(ctx, db, lotusClient, CreateRequest{
			ActorID:   knownClientActorID,
			TrackOnly: true,
			Name:      "E2E Test Client Wallet",
			Contact:   "test@example.com",
			Location:  "Global",
		})

		require.NoError(t, err, "Should successfully create TrackedWallet with real Lotus API")
		require.Equal(t, knownClientActorID, w.ActorID)
		require.NotEmpty(t, w.Address, "Address should be resolved from ActorID")
		require.Equal(t, "TrackedWallet", string(w.WalletType))
		require.Equal(t, "E2E Test Client Wallet", w.ActorName)
		require.Equal(t, "test@example.com", w.ContactInfo)
		require.Equal(t, "Global", w.Location)
		require.Empty(t, w.PrivateKey, "TrackedWallet should not have private key")

		t.Logf("Successfully created TrackedWallet:")
		t.Logf("  ActorID: %s", w.ActorID)
		t.Logf("  Address: %s", w.Address)
		t.Logf("  Type: %s", w.WalletType)

		// Verify we can track multiple different wallets
		// Use a second known client wallet for testing
		w2, err := Default.CreateHandler(ctx, db, lotusClient, CreateRequest{
			ActorID:   "f03510418", // Another client wallet
			TrackOnly: true,
			Name:      "Second E2E Test Client",
		})

		require.NoError(t, err, "Should be able to create second TrackedWallet")
		require.Equal(t, "f03510418", w2.ActorID)
		require.NotEmpty(t, w2.Address)
		require.NotEqual(t, w.Address, w2.Address, "Different ActorIDs should resolve to different addresses")

		// Test error case with invalid ActorID
		_, err = Default.CreateHandler(ctx, db, lotusClient, CreateRequest{
			ActorID:   "f0999999999", // Non-existent ActorID
			TrackOnly: true,
		})
		require.Error(t, err, "Should fail with non-existent ActorID")
		require.Contains(t, err.Error(), "does not exist on the network")

		// Test error case with storage provider ActorID
		_, err = Default.CreateHandler(ctx, db, lotusClient, CreateRequest{
			ActorID:   "f01000", // Known storage provider
			TrackOnly: true,
		})
		require.Error(t, err, "Should fail when trying to track storage provider")
		require.Contains(t, err.Error(), "is a storage provider, not a client wallet")
		require.Contains(t, err.Error(), "Wallet tracking is for client wallets that make deals")

		t.Logf("Successfully validated error handling for non-client ActorIDs")
	})
}
