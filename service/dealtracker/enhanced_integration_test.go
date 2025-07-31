package dealtracker

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/statetracker"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/ipfs/boxo/util"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestEnhancedDealTracking(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Create wallet for testing
		wallet := model.Wallet{
			ActorID: "t0100",
			Address: "t3xxx",
		}
		err := db.Create(&wallet).Error
		require.NoError(t, err)

		// Create test deal that will be slashed
		dealID := uint64(1)
		cidVal := model.CID(cid.NewCidV1(cid.Raw, util.Hash([]byte("test"))))

		testDeal := model.Deal{
			DealID:           &dealID,
			State:            model.DealActive,
			ClientID:         &wallet.ID,
			ClientActorID:    wallet.ActorID,
			Provider:         "t01000",
			PieceCID:         cidVal,
			PieceSize:        1024,
			StartEpoch:       100,
			EndEpoch:         200,
			SectorStartEpoch: 100,
			Price:            "1000",
			Verified:         true,
		}
		err = db.Create(&testDeal).Error
		require.NoError(t, err)

		// Create mock deal data that shows the deal as slashed
		slashedDeal := Deal{
			Proposal: DealProposal{
				PieceCID:             Cid{Root: cidVal.String()},
				PieceSize:            1024,
				VerifiedDeal:         true,
				Client:               "t0100",
				Provider:             "t01000",
				StartEpoch:           100,
				EndEpoch:             200,
				StoragePricePerEpoch: "1000",
				Label:                "test",
			},
			State: DealState{
				SectorStartEpoch: 100,
				LastUpdatedEpoch: 150,
				SlashEpoch:       125, // Deal was slashed at epoch 125
			},
		}

		// Create test server with slashed deal data
		deals := map[string]Deal{
			"1": slashedDeal,
		}
		body, err := json.Marshal(deals)
		require.NoError(t, err)

		url, server := setupTestServerWithBody(t, string(body))
		defer server.Close()

		// Create deal tracker
		tracker := NewDealTracker(db, time.Minute, url, "https://api.node.glif.io/", "", true)

		// Run tracking once
		err = tracker.runOnce(ctx)
		require.NoError(t, err)

		// Verify deal state was updated to slashed
		var updatedDeal model.Deal
		err = db.First(&updatedDeal, testDeal.ID).Error
		require.NoError(t, err)
		require.Equal(t, model.DealSlashed, updatedDeal.State)

		// Verify state changes were tracked with enhanced metadata
		// We expect 2 state changes: 1 for recovery and 1 for the actual slashing
		var stateChanges []model.DealStateChange
		err = db.Where("deal_id = ?", testDeal.ID).Order("timestamp ASC").Find(&stateChanges).Error
		require.NoError(t, err)
		require.Len(t, stateChanges, 2)

		// First state change should be the recovery
		recoveryChange := stateChanges[0]
		require.Equal(t, testDeal.ID, recoveryChange.DealID)
		require.Equal(t, model.DealState(""), recoveryChange.PreviousState) // No previous state for recovery
		require.Equal(t, model.DealActive, recoveryChange.NewState)

		// Second state change should be the slashing event
		sc := stateChanges[1]
		require.Equal(t, testDeal.ID, sc.DealID)
		require.Equal(t, model.DealActive, sc.PreviousState)
		require.Equal(t, model.DealSlashed, sc.NewState)
		require.Equal(t, "t01000", sc.ProviderID)
		require.Equal(t, "t0100", sc.ClientAddress)

		// Verify enhanced metadata
		var metadata statetracker.StateChangeMetadata
		err = json.Unmarshal([]byte(sc.Metadata), &metadata)
		require.NoError(t, err)
		require.Contains(t, metadata.Reason, "slashed")
		require.Contains(t, metadata.Reason, "125")
		require.Equal(t, "1000", metadata.StoragePrice)
		require.Equal(t, int64(1024), metadata.PieceSize)
		require.True(t, metadata.VerifiedDeal)
		require.NotNil(t, metadata.SlashingEpoch)
		require.Equal(t, int32(125), *metadata.SlashingEpoch)
	})
}

func TestEnhancedExternalDealDiscovery(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Create wallet for testing
		wallet := model.Wallet{
			ActorID: "t0100",
			Address: "t3xxx",
		}
		err := db.Create(&wallet).Error
		require.NoError(t, err)

		cidVal := model.CID(cid.NewCidV1(cid.Raw, util.Hash([]byte("external"))))

		// Create mock external deal data with far future epochs to ensure it's active
		// Use epochs that are guaranteed to be in the far future
		currentEpoch := int32(10000000) // Use a very high epoch to ensure it's future
		externalDeal := Deal{
			Proposal: DealProposal{
				PieceCID:             Cid{Root: cidVal.String()},
				PieceSize:            2048,
				VerifiedDeal:         false,
				Client:               "t0100",
				Provider:             "t01001",
				StartEpoch:           currentEpoch - 1000,    // Started 1000 epochs ago
				EndEpoch:             currentEpoch + 1000000, // Ends 1M epochs in the future
				StoragePricePerEpoch: "2000",
				Label:                "external-deal",
			},
			State: DealState{
				SectorStartEpoch: currentEpoch - 900, // Activated 900 epochs ago
				LastUpdatedEpoch: currentEpoch - 800, // Last updated 800 epochs ago
				SlashEpoch:       -1,
			},
		}

		// Create test server with external deal data
		deals := map[string]Deal{
			"999": externalDeal,
		}
		body, err := json.Marshal(deals)
		require.NoError(t, err)

		url, server := setupTestServerWithBody(t, string(body))
		defer server.Close()

		// Create deal tracker
		tracker := NewDealTracker(db, time.Minute, url, "https://api.node.glif.io/", "", true)

		// Run tracking once
		err = tracker.runOnce(ctx)
		require.NoError(t, err)

		// Verify external deal was discovered and inserted
		var discoveredDeals []model.Deal
		err = db.Where("deal_id = ?", 999).Find(&discoveredDeals).Error
		require.NoError(t, err)
		require.Len(t, discoveredDeals, 1)

		deal := discoveredDeals[0]
		require.Equal(t, uint64(999), *deal.DealID)
		require.Equal(t, model.DealActive, deal.State)
		require.Equal(t, "t01001", deal.Provider)
		require.Equal(t, cidVal, deal.PieceCID)
		require.Equal(t, int64(2048), deal.PieceSize)
		require.False(t, deal.Verified)

		// Verify state change was tracked for external deal
		var stateChanges []model.DealStateChange
		err = db.Where("deal_id = ?", deal.ID).Find(&stateChanges).Error
		require.NoError(t, err)
		require.Len(t, stateChanges, 1)

		sc := stateChanges[0]
		require.Equal(t, deal.ID, sc.DealID)
		require.Equal(t, model.DealState(""), sc.PreviousState) // No previous state for external deals
		require.Equal(t, model.DealActive, sc.NewState)

		// Verify enhanced metadata for external deal
		var metadata statetracker.StateChangeMetadata
		err = json.Unmarshal([]byte(sc.Metadata), &metadata)
		require.NoError(t, err)
		require.Contains(t, metadata.Reason, "External deal discovered")
		require.Equal(t, "2000", metadata.StoragePrice)
		require.Equal(t, int64(2048), metadata.PieceSize)
		require.False(t, metadata.VerifiedDeal)
		require.NotNil(t, metadata.ActivationEpoch)
		require.Equal(t, int32(9999100), *metadata.ActivationEpoch) // currentEpoch - 900
	})
}

func TestEnhancedDealExpiration(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Create wallet for testing
		wallet := model.Wallet{
			ActorID: "t0100",
			Address: "t3xxx",
		}
		err := db.Create(&wallet).Error
		require.NoError(t, err)

		// Create test deal that will expire
		dealID := uint64(1)
		cidVal := model.CID(cid.NewCidV1(cid.Raw, util.Hash([]byte("expiring"))))

		expiringDeal := model.Deal{
			DealID:           &dealID,
			State:            model.DealActive,
			ClientID:         &wallet.ID,
			ClientActorID:    wallet.ActorID,
			Provider:         "t01000",
			PieceCID:         cidVal,
			PieceSize:        1024,
			StartEpoch:       100,
			EndEpoch:         150, // Will expire at epoch 150
			SectorStartEpoch: 100,
			Price:            "1000",
			Verified:         true,
		}
		err = db.Create(&expiringDeal).Error
		require.NoError(t, err)

		// Create mock deal data that simulates the current epoch past expiration
		deals := map[string]Deal{
			"1": {
				Proposal: DealProposal{
					PieceCID:             Cid{Root: cidVal.String()},
					PieceSize:            1024,
					VerifiedDeal:         true,
					Client:               "t0100",
					Provider:             "t01000",
					StartEpoch:           100,
					EndEpoch:             150,
					StoragePricePerEpoch: "1000",
					Label:                "expiring",
				},
				State: DealState{
					SectorStartEpoch: 100,
					LastUpdatedEpoch: 200, // Current epoch is past end epoch
					SlashEpoch:       -1,
				},
			},
		}
		body, err := json.Marshal(deals)
		require.NoError(t, err)

		url, server := setupTestServerWithBody(t, string(body))
		defer server.Close()

		// Create deal tracker
		tracker := NewDealTracker(db, time.Minute, url, "https://api.node.glif.io/", "", true)

		// Run tracking once
		err = tracker.runOnce(ctx)
		require.NoError(t, err)

		// Verify deal state was updated to expired
		var updatedDeal model.Deal
		err = db.First(&updatedDeal, expiringDeal.ID).Error
		require.NoError(t, err)
		require.Equal(t, model.DealExpired, updatedDeal.State)

		// Verify state changes were tracked (one for active state, one for expiration)
		var stateChanges []model.DealStateChange
		err = db.Where("deal_id = ?", expiringDeal.ID).Order("timestamp ASC").Find(&stateChanges).Error
		require.NoError(t, err)
		// Should have both transition and expiration tracking
		require.GreaterOrEqual(t, len(stateChanges), 1)

		// Check the latest state change for expiration
		latestSC := stateChanges[len(stateChanges)-1]
		require.Equal(t, model.DealExpired, latestSC.NewState)

		// Verify enhanced metadata for expiration
		var metadata statetracker.StateChangeMetadata
		err = json.Unmarshal([]byte(latestSC.Metadata), &metadata)
		require.NoError(t, err)
		require.Contains(t, metadata.Reason, "expired")
		require.Equal(t, "1000", metadata.StoragePrice)
		require.Equal(t, int64(1024), metadata.PieceSize)
		require.True(t, metadata.VerifiedDeal)
		require.NotNil(t, metadata.ExpirationEpoch)
		require.Equal(t, int32(150), *metadata.ExpirationEpoch)
	})
}

func TestEnhancedProposalExpiration(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Create wallet for testing
		wallet := model.Wallet{
			ActorID: "t0100",
			Address: "t3xxx",
		}
		err := db.Create(&wallet).Error
		require.NoError(t, err)

		cidVal := model.CID(cid.NewCidV1(cid.Raw, util.Hash([]byte("proposal"))))

		// Create test deal proposal that will expire
		expiringProposal := model.Deal{
			State:         model.DealProposed,
			ClientID:      &wallet.ID,
			ClientActorID: wallet.ActorID,
			Provider:      "t01000",
			PieceCID:      cidVal,
			PieceSize:     1024,
			StartEpoch:    100, // Will expire at epoch 100 without activation
			EndEpoch:      200,
			Price:         "1000",
			Verified:      false,
		}
		err = db.Create(&expiringProposal).Error
		require.NoError(t, err)

		// Create mock deal data with current epoch past start epoch but deal not activated
		deals := map[string]Deal{
			"1": {
				Proposal: DealProposal{
					PieceCID:             Cid{Root: cidVal.String()},
					PieceSize:            1024,
					VerifiedDeal:         false,
					Client:               "t0100",
					Provider:             "t01000",
					StartEpoch:           100,
					EndEpoch:             200,
					StoragePricePerEpoch: "1000",
					Label:                "proposal",
				},
				State: DealState{
					SectorStartEpoch: -1,  // Deal not activated
					LastUpdatedEpoch: 150, // Current epoch past start epoch
					SlashEpoch:       -1,
				},
			},
		}
		body, err := json.Marshal(deals)
		require.NoError(t, err)

		url, server := setupTestServerWithBody(t, string(body))
		defer server.Close()

		// Create deal tracker
		tracker := NewDealTracker(db, time.Minute, url, "https://api.node.glif.io/", "", true)

		// Run tracking once
		err = tracker.runOnce(ctx)
		require.NoError(t, err)

		// Verify proposal state was updated to expired
		var updatedDeal model.Deal
		err = db.First(&updatedDeal, expiringProposal.ID).Error
		require.NoError(t, err)
		require.Equal(t, model.DealProposalExpired, updatedDeal.State)

		// Verify state change was tracked for proposal expiration
		var stateChanges []model.DealStateChange
		err = db.Where("deal_id = ?", expiringProposal.ID).Find(&stateChanges).Error
		require.NoError(t, err)
		require.GreaterOrEqual(t, len(stateChanges), 1)

		// Find the expiration state change
		var expirationSC *model.DealStateChange
		for _, sc := range stateChanges {
			if sc.NewState == model.DealProposalExpired {
				expirationSC = &sc
				break
			}
		}
		require.NotNil(t, expirationSC)

		// Verify enhanced metadata for proposal expiration
		var metadata statetracker.StateChangeMetadata
		err = json.Unmarshal([]byte(expirationSC.Metadata), &metadata)
		require.NoError(t, err)
		require.Contains(t, metadata.Reason, "proposal expired")
		require.Equal(t, "1000", metadata.StoragePrice)
		require.Equal(t, int64(1024), metadata.PieceSize)
		require.False(t, metadata.VerifiedDeal)
		require.NotNil(t, metadata.ExpirationEpoch)
		require.Equal(t, int32(100), *metadata.ExpirationEpoch)
	})
}
