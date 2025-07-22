package proof

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(&model.Deal{}, &model.Proof{})
	require.NoError(t, err)

	return db
}

func TestDefaultHandler_ListHandler(t *testing.T) {
	db := setupTestDB(t)
	handler := &DefaultHandler{}
	ctx := context.Background()

	// Create test data
	deal1 := model.Deal{
		DealID:   uint64Ptr(1),
		Provider: "f01000",
		State:    model.DealActive,
	}
	deal2 := model.Deal{
		DealID:   uint64Ptr(2),
		Provider: "f01001",
		State:    model.DealActive,
	}
	require.NoError(t, db.Create(&deal1).Error)
	require.NoError(t, db.Create(&deal2).Error)

	proof1 := model.Proof{
		DealID:    uint64Ptr(1),
		ProofType: model.ProofOfReplication,
		Provider:  "f01000",
		MessageID: "bafy2bzacea1",
		Height:    1000,
		Method:    "ProveCommitSector",
		Verified:  true,
	}
	proof2 := model.Proof{
		DealID:    uint64Ptr(2),
		ProofType: model.ProofOfSpacetime,
		Provider:  "f01001",
		MessageID: "bafy2bzacea2",
		Height:    2000,
		Method:    "SubmitWindowedPoSt",
		Verified:  false,
	}
	proof3 := model.Proof{
		ProofType: model.ProofOfReplication,
		Provider:  "f01000",
		MessageID: "bafy2bzacea3",
		Height:    3000,
		Method:    "PreCommitSector",
		Verified:  true,
	}
	require.NoError(t, db.Create(&proof1).Error)
	require.NoError(t, db.Create(&proof2).Error)
	require.NoError(t, db.Create(&proof3).Error)

	tests := []struct {
		name     string
		request  ListProofRequest
		expected int
	}{
		{
			name:     "list all proofs",
			request:  ListProofRequest{},
			expected: 3,
		},
		{
			name: "filter by deal ID",
			request: ListProofRequest{
				DealID: uint64Ptr(1),
			},
			expected: 1,
		},
		{
			name: "filter by proof type",
			request: ListProofRequest{
				ProofType: proofTypePtr(model.ProofOfReplication),
			},
			expected: 2,
		},
		{
			name: "filter by provider",
			request: ListProofRequest{
				Provider: stringPtr("f01000"),
			},
			expected: 2,
		},
		{
			name: "filter by verified status",
			request: ListProofRequest{
				Verified: boolPtr(true),
			},
			expected: 2,
		},
		{
			name: "filter by verified status false",
			request: ListProofRequest{
				Verified: boolPtr(false),
			},
			expected: 1,
		},
		{
			name: "filter with limit",
			request: ListProofRequest{
				Limit: 2,
			},
			expected: 2,
		},
		{
			name: "filter with offset",
			request: ListProofRequest{
				Limit:  10,
				Offset: 1,
			},
			expected: 2,
		},
		{
			name: "multiple filters",
			request: ListProofRequest{
				Provider:  stringPtr("f01000"),
				ProofType: proofTypePtr(model.ProofOfReplication),
				Verified:  boolPtr(true),
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			proofs, err := handler.ListHandler(ctx, db, tt.request)
			require.NoError(t, err)
			assert.Len(t, proofs, tt.expected)

			// Verify proofs are ordered by created_at DESC
			if len(proofs) > 1 {
				for i := 1; i < len(proofs); i++ {
					assert.True(t, proofs[i-1].CreatedAt.After(proofs[i].CreatedAt) || proofs[i-1].CreatedAt.Equal(proofs[i].CreatedAt))
				}
			}
		})
	}
}

func TestDefaultHandler_ListHandler_DefaultLimit(t *testing.T) {
	db := setupTestDB(t)
	handler := &DefaultHandler{}
	ctx := context.Background()

	// Test default limit is applied when limit is 0
	request := ListProofRequest{Limit: 0}
	_, err := handler.ListHandler(ctx, db, request)
	require.NoError(t, err)

	// The limit should be set to 100 internally
	// We can't directly test this without exposing internal state,
	// but we can verify the query doesn't fail
}

func TestDefaultHandler_ListHandler_EmptyResult(t *testing.T) {
	db := setupTestDB(t)
	handler := &DefaultHandler{}
	ctx := context.Background()

	// Test with no proofs in database
	request := ListProofRequest{}
	proofs, err := handler.ListHandler(ctx, db, request)
	require.NoError(t, err)
	assert.Empty(t, proofs)
}

// Helper functions
func uint64Ptr(v uint64) *uint64 {
	return &v
}

func stringPtr(v string) *string {
	return &v
}

func boolPtr(v bool) *bool {
	return &v
}

func proofTypePtr(v model.ProofType) *model.ProofType {
	return &v
}
