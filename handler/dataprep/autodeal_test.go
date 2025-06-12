package dataprep

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestCreatePreparationHandler_AutoDealCreation(t *testing.T) {
	tmp1 := t.TempDir()
	tmp2 := t.TempDir()
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {

		// Create a test source storage
		sourceStorage, err := storage.Default.CreateStorageHandler(ctx, db, "local", storage.CreateRequest{
			Name: "test-source",
			Path: tmp1,
		})
		require.NoError(t, err)

		// Create a test output storage
		outputStorage, err := storage.Default.CreateStorageHandler(ctx, db, "local", storage.CreateRequest{
			Name: "test-output",
			Path: tmp2,
		})
		require.NoError(t, err)

		// Test auto-deal creation with all parameters
		request := CreateRequest{
			Name:                "test-prep-autodeal",
			SourceStorages:      []string{sourceStorage.Name},
			OutputStorages:      []string{outputStorage.Name},
			MaxSizeStr:          "2GB",
			AutoCreateDeals:     true,
			DealPricePerGB:      0.1,
			DealPricePerGBEpoch: 0.0000001,
			DealPricePerDeal:    0.01,
			DealDuration:        time.Hour * 24 * 535, // 535 days
			DealStartDelay:      time.Hour * 72,       // 72 hours
			DealVerified:        true,
			DealKeepUnsealed:    true,
			DealAnnounceToIPNI:  true,
			DealProvider:        "f01000",
			DealHTTPHeaders:     model.ConfigMap{"Authorization": "Bearer token"},
			DealURLTemplate:     "https://example.com/deals/{id}",
			WalletValidation:    true,
			SPValidation:        true,
		}

		prep, err := Default.CreatePreparationHandler(ctx, db, request)
		require.NoError(t, err)

		// Verify the preparation was created with auto-deal parameters
		assert.Equal(t, "test-prep-autodeal", prep.Name)
		assert.True(t, prep.AutoCreateDeals)
		assert.Equal(t, 0.1, prep.DealPricePerGB)
		assert.Equal(t, 0.0000001, prep.DealPricePerGBEpoch)
		assert.Equal(t, 0.01, prep.DealPricePerDeal)
		assert.Equal(t, time.Hour*24*535, prep.DealDuration)
		assert.Equal(t, time.Hour*72, prep.DealStartDelay)
		assert.True(t, prep.DealVerified)
		assert.True(t, prep.DealKeepUnsealed)
		assert.True(t, prep.DealAnnounceToIPNI)
		assert.Equal(t, "f01000", prep.DealProvider)
		assert.Equal(t, model.ConfigMap{"Authorization": "Bearer token"}, prep.DealHTTPHeaders)
		assert.Equal(t, "https://example.com/deals/{id}", prep.DealURLTemplate)
		assert.True(t, prep.WalletValidation)
		assert.True(t, prep.SPValidation)

		// Verify that notifications were created during validation
		var notifications []model.Notification
		err = db.Where("source = ?", "dataprep-create").Find(&notifications).Error
		require.NoError(t, err)
		assert.Greater(t, len(notifications), 0, "Expected notifications to be created during auto-deal validation")

		// Check that at least one notification is about starting validation
		foundValidationStart := false
		for _, notification := range notifications {
			if notification.Title == "Starting Auto-Deal Validation" {
				foundValidationStart = true
				break
			}
		}
		assert.True(t, foundValidationStart, "Expected to find validation start notification")
	})
}

func TestCreatePreparationHandler_AutoDealDisabled(t *testing.T) {
	tmp1 := t.TempDir()
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {

		// Create a test source storage
		sourceStorage, err := storage.Default.CreateStorageHandler(ctx, db, "local", storage.CreateRequest{
			Name: "test-source",
			Path: tmp1,
		})
		require.NoError(t, err)

		// Test with auto-deal creation disabled
		request := CreateRequest{
			Name:            "test-prep-no-autodeal",
			SourceStorages:  []string{sourceStorage.Name},
			MaxSizeStr:      "2GB",
			AutoCreateDeals: false,
		}

		prep, err := Default.CreatePreparationHandler(ctx, db, request)
		require.NoError(t, err)

		// Verify the preparation was created without auto-deal parameters
		assert.Equal(t, "test-prep-no-autodeal", prep.Name)
		assert.False(t, prep.AutoCreateDeals)
		assert.Equal(t, 0.0, prep.DealPricePerGB)
		assert.Equal(t, 0.0, prep.DealPricePerGBEpoch)
		assert.Equal(t, 0.0, prep.DealPricePerDeal)
		assert.Equal(t, time.Duration(0), prep.DealDuration)
		assert.Equal(t, time.Duration(0), prep.DealStartDelay)
		assert.False(t, prep.DealVerified)
		assert.False(t, prep.DealKeepUnsealed)
		assert.False(t, prep.DealAnnounceToIPNI)
		assert.Equal(t, "", prep.DealProvider)
		assert.Equal(t, "", prep.DealURLTemplate)
		assert.False(t, prep.WalletValidation)
		assert.False(t, prep.SPValidation)

		// Verify that no auto-deal validation notifications were created
		var notifications []model.Notification
		err = db.Where("source = ? AND title LIKE ?", "dataprep-create", "%Auto-Deal%").Find(&notifications).Error
		require.NoError(t, err)
		assert.Equal(t, 0, len(notifications), "Expected no auto-deal notifications when auto-deal is disabled")
	})
}

func TestCreatePreparationHandler_ValidationOnly(t *testing.T) {
	tmp1 := t.TempDir()
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {

		// Create a test source storage
		sourceStorage, err := storage.Default.CreateStorageHandler(ctx, db, "local", storage.CreateRequest{
			Name: "test-source",
			Path: tmp1,
		})
		require.NoError(t, err)

		// Test with auto-deal creation enabled but only wallet validation
		request := CreateRequest{
			Name:             "test-prep-wallet-validation",
			SourceStorages:   []string{sourceStorage.Name},
			MaxSizeStr:       "2GB",
			AutoCreateDeals:  true,
			WalletValidation: true,
			SPValidation:     false,
		}

		prep, err := Default.CreatePreparationHandler(ctx, db, request)
		require.NoError(t, err)

		// Verify the preparation was created
		assert.Equal(t, "test-prep-wallet-validation", prep.Name)
		assert.True(t, prep.AutoCreateDeals)
		assert.True(t, prep.WalletValidation)
		assert.False(t, prep.SPValidation)

		// Verify that validation notifications were created
		var notifications []model.Notification
		err = db.Where("source = ? AND title LIKE ?", "dataprep-create", "%Wallet%").Find(&notifications).Error
		require.NoError(t, err)
		assert.Greater(t, len(notifications), 0, "Expected wallet validation notifications")
	})
}
