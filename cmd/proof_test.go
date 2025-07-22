package cmd

import (
	"context"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/proof"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/gotidy/ptr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func swapProofHandler(mockHandler proof.Handler) func() {
	actual := proof.Default
	proof.Default = mockHandler
	return func() {
		proof.Default = actual
	}
}

func TestProofListHandler(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(proof.MockProof)
		defer swapProofHandler(mockHandler)()

		mockHandler.On("ListHandler", mock.Anything, mock.Anything, mock.Anything).Return([]model.Proof{
			{
				ID:        1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DealID:    ptr.Of(uint64(100)),
				ProofType: model.ProofOfReplication,
				MessageID: "bafy2bzacea1",
				BlockCID:  "bafy2bzaceb1",
				Height:    1000000,
				Method:    "ProveCommitSector",
				Verified:  true,
				SectorID:  ptr.Of(uint64(456)),
				Provider:  "f01000",
				ErrorMsg:  "",
			},
			{
				ID:        2,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DealID:    ptr.Of(uint64(200)),
				ProofType: model.ProofOfSpacetime,
				MessageID: "bafy2bzacea2",
				BlockCID:  "bafy2bzaceb2",
				Height:    2000000,
				Method:    "SubmitWindowedPoSt",
				Verified:  false,
				SectorID:  ptr.Of(uint64(789)),
				Provider:  "f01001",
				ErrorMsg:  "validation failed",
			},
		}, nil)

		// Test basic list command
		_, _, err := runner.Run(ctx, "singularity proof list")
		require.NoError(t, err)

		// Test with verbose output
		_, _, err = runner.Run(ctx, "singularity --verbose proof list")
		require.NoError(t, err)

		// Test with JSON output
		_, _, err = runner.Run(ctx, "singularity --json proof list")
		require.NoError(t, err)

		// Test with filters
		_, _, err = runner.Run(ctx, "singularity proof list --deal-id 100")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity proof list --proof-type replication")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity proof list --provider f01000")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity proof list --verified")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity proof list --unverified")
		require.NoError(t, err)

		// Test with pagination
		_, _, err = runner.Run(ctx, "singularity proof list --limit 10 --offset 5")
		require.NoError(t, err)

		// Test with multiple filters
		_, _, err = runner.Run(ctx, "singularity proof list --deal-id 100 --proof-type replication --verified")
		require.NoError(t, err)
	})
}

func TestProofSyncHandler(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(proof.MockProof)
		defer swapProofHandler(mockHandler)()

		mockHandler.On("SyncHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

		// Test basic sync command
		_, _, err := runner.Run(ctx, "singularity proof sync")
		require.NoError(t, err)

		// Test sync for specific deal
		_, _, err = runner.Run(ctx, "singularity proof sync --deal-id 123")
		require.NoError(t, err)

		// Test sync for specific provider
		_, _, err = runner.Run(ctx, "singularity proof sync --provider f01000")
		require.NoError(t, err)

		// Test with verbose output
		_, _, err = runner.Run(ctx, "singularity --verbose proof sync")
		require.NoError(t, err)

		// Test with JSON output
		_, _, err = runner.Run(ctx, "singularity --json proof sync")
		require.NoError(t, err)

		// Test with both deal-id and provider (should work)
		_, _, err = runner.Run(ctx, "singularity proof sync --deal-id 123 --provider f01000")
		require.NoError(t, err)
	})
}

func TestProofCommandHelp(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)

		// Test help commands
		_, _, err := runner.Run(ctx, "singularity proof --help")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity proof list --help")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity proof sync --help")
		require.NoError(t, err)
	})
}

func TestProofListHandlerWithErrors(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(proof.MockProof)
		defer swapProofHandler(mockHandler)()

		// Mock an error response
		mockHandler.On("ListHandler", mock.Anything, mock.Anything, mock.Anything).Return([]model.Proof{},
			errors.New("database connection failed"))

		// Test that error is properly handled
		_, _, err := runner.Run(ctx, "singularity proof list")
		require.Error(t, err)
		require.Contains(t, err.Error(), "database connection failed")
	})
}

func TestProofSyncHandlerWithErrors(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(proof.MockProof)
		defer swapProofHandler(mockHandler)()

		// Mock an error response
		mockHandler.On("SyncHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
			errors.New("lotus client connection failed"))

		// Test that error is properly handled
		_, _, err := runner.Run(ctx, "singularity proof sync")
		require.Error(t, err)
		require.Contains(t, err.Error(), "lotus client connection failed")
	})
}

func TestProofCommandValidation(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)

		// Test invalid proof type
		_, _, err := runner.Run(ctx, "singularity proof list --proof-type invalid")
		// This should not error at CLI level since validation happens in handler
		require.NoError(t, err)

		// Test invalid deal ID format - this should be caught by CLI flag parsing
		_, _, err = runner.Run(ctx, "singularity proof list --deal-id invalid")
		require.Error(t, err)

		// Test invalid limit format
		_, _, err = runner.Run(ctx, "singularity proof list --limit invalid")
		require.Error(t, err)

		// Test invalid offset format
		_, _, err = runner.Run(ctx, "singularity proof list --offset invalid")
		require.Error(t, err)

		// Test invalid provider format for sync
		_, _, err = runner.Run(ctx, "singularity proof sync --deal-id invalid")
		require.Error(t, err)
	})
}
