package service

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/filecoin-project/go-address"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(&model.DealProof{}, &model.ProofVerification{})
	require.NoError(t, err)

	return db
}

func TestProofService_StoreProof(t *testing.T) {
	db := setupTestDB(t)
	service := NewProofService(db)
	ctx := context.Background()

	clientAddr, err := address.NewIDAddress(1000)
	require.NoError(t, err)

	providerAddr, err := address.NewIDAddress(2000)
	require.NoError(t, err)

	proof := &model.DealProof{
		DealID:          1,
		PieceCID:        "bafk2bzacedbootlz5dgj5migqrjmwjrqygrmyqwqrpxs7x5n5uncgqkqrw75y",
		ProofBytes:      []byte("test proof"),
		SectorID:        1,
		ClientAddress:   clientAddr.String(),
		ProviderAddress: providerAddr.String(),
	}

	err = service.StoreProof(ctx, proof)
	require.NoError(t, err)

	// Verify the proof was stored
	storedProof, err := service.GetProofByDealID(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, proof.DealID, storedProof.DealID)
	require.Equal(t, proof.PieceCID, storedProof.PieceCID)
	require.Equal(t, model.ProofStatusPending, storedProof.ProofStatus)
}

func TestProofService_VerifyProof(t *testing.T) {
	db := setupTestDB(t)
	service := NewProofService(db)
	ctx := context.Background()

	clientAddr, err := address.NewIDAddress(1000)
	require.NoError(t, err)

	providerAddr, err := address.NewIDAddress(2000)
	require.NoError(t, err)

	verifierAddr, err := address.NewIDAddress(3000)
	require.NoError(t, err)

	// Store a proof first
	proof := &model.DealProof{
		DealID:          1,
		PieceCID:        "bafk2bzacedbootlz5dgj5migqrjmwjrqygrmyqwqrpxs7x5n5uncgqkqrw75y",
		ProofBytes:      []byte("test proof"),
		SectorID:        1,
		ClientAddress:   clientAddr.String(),
		ProviderAddress: providerAddr.String(),
	}

	err = service.StoreProof(ctx, proof)
	require.NoError(t, err)

	// Verify the proof
	verification, err := service.VerifyProof(ctx, 1, verifierAddr)
	require.NoError(t, err)
	require.True(t, verification.VerificationResult)

	// Check the proof status was updated
	storedProof, err := service.GetProofByDealID(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, model.ProofStatusVerified, storedProof.ProofStatus)
	require.NotNil(t, storedProof.VerificationTime)
}

func TestProofService_ListClientProofs(t *testing.T) {
	db := setupTestDB(t)
	service := NewProofService(db)
	ctx := context.Background()

	clientAddr, err := address.NewIDAddress(1000)
	require.NoError(t, err)

	providerAddr, err := address.NewIDAddress(2000)
	require.NoError(t, err)

	// Store multiple proofs
	for i := uint64(1); i <= 3; i++ {
		proof := &model.DealProof{
			DealID:          i,
			PieceCID:        "bafk2bzacedbootlz5dgj5migqrjmwjrqygrmyqwqrpxs7x5n5uncgqkqrw75y",
			ProofBytes:      []byte("test proof"),
			SectorID:        i,
			ClientAddress:   clientAddr.String(),
			ProviderAddress: providerAddr.String(),
		}
		err = service.StoreProof(ctx, proof)
		require.NoError(t, err)
	}

	// List proofs for the client
	proofs, err := service.ListClientProofs(ctx, clientAddr)
	require.NoError(t, err)
	require.Len(t, proofs, 3)
}

func TestProofService_GetProofVerificationHistory(t *testing.T) {
	db := setupTestDB(t)
	service := NewProofService(db)
	ctx := context.Background()

	clientAddr, err := address.NewIDAddress(1000)
	require.NoError(t, err)

	providerAddr, err := address.NewIDAddress(2000)
	require.NoError(t, err)

	verifierAddr, err := address.NewIDAddress(3000)
	require.NoError(t, err)

	// Store a proof
	proof := &model.DealProof{
		DealID:          1,
		PieceCID:        "bafk2bzacedbootlz5dgj5migqrjmwjrqygrmyqwqrpxs7x5n5uncgqkqrw75y",
		ProofBytes:      []byte("test proof"),
		SectorID:        1,
		ClientAddress:   clientAddr.String(),
		ProviderAddress: providerAddr.String(),
	}

	err = service.StoreProof(ctx, proof)
	require.NoError(t, err)

	// Verify the proof multiple times
	for i := 0; i < 3; i++ {
		_, err = service.VerifyProof(ctx, 1, verifierAddr)
		require.NoError(t, err)
	}

	// Get verification history
	history, err := service.GetProofVerificationHistory(ctx, 1)
	require.NoError(t, err)
	require.Len(t, history, 3)
}
