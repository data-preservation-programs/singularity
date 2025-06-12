package wallet

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/gotidy/ptr"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestDefaultHandler_UpdateHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Create a test wallet first
		wallet := model.Wallet{
			ActorID:     "test-id",
			ActorName:   "Original Name",
			Address:     "f1test123address",
			ContactInfo: "original@example.com",
			Location:    "Original Location",
			WalletType:  model.SPWallet,
		}
		err := db.Create(&wallet).Error
		require.NoError(t, err)

		handler := DefaultHandler{}

		t.Run("update all fields", func(t *testing.T) {
			newActorName := "Updated Name"
			newContactInfo := "updated@example.com"
			newLocation := "Updated Location"

			request := UpdateRequest{
				Name:     &newActorName,
				Contact:  &newContactInfo,
				Location: &newLocation,
			}

			updated, err := handler.UpdateHandler(ctx, db, wallet.Address, request)
			require.NoError(t, err)
			require.NotNil(t, updated)

			require.Equal(t, newActorName, updated.ActorName)
			require.Equal(t, newContactInfo, updated.ContactInfo)
			require.Equal(t, newLocation, updated.Location)

			// Verify essential fields are unchanged
			require.Equal(t, wallet.Address, updated.Address)
			require.Equal(t, wallet.PrivateKey, updated.PrivateKey)
			require.Equal(t, wallet.ActorID, updated.ActorID)
			require.Equal(t, wallet.WalletType, updated.WalletType)
		})

		t.Run("update single field", func(t *testing.T) {
			newActorName := "Single Field Update"
			request := UpdateRequest{
				Name: &newActorName,
			}

			updated, err := handler.UpdateHandler(ctx, db, wallet.Address, request)
			require.NoError(t, err)
			require.NotNil(t, updated)

			require.Equal(t, newActorName, updated.ActorName)
			// Other fields should remain as they were from the previous update
			require.Equal(t, "updated@example.com", updated.ContactInfo)
			require.Equal(t, "Updated Location", updated.Location)
			require.Equal(t, model.SPWallet, updated.WalletType)
		})

		t.Run("update with empty string", func(t *testing.T) {
			emptyString := ""
			request := UpdateRequest{
				Contact: &emptyString,
			}

			updated, err := handler.UpdateHandler(ctx, db, wallet.Address, request)
			require.NoError(t, err)
			require.NotNil(t, updated)

			require.Equal(t, "", updated.ContactInfo)
		})

		t.Run("no fields to update", func(t *testing.T) {
			request := UpdateRequest{}

			updated, err := handler.UpdateHandler(ctx, db, wallet.Address, request)
			require.NoError(t, err)
			require.NotNil(t, updated)

			// Should return the existing wallet unchanged
			require.Equal(t, "Single Field Update", updated.ActorName)
		})

		t.Run("wallet not found", func(t *testing.T) {
			request := UpdateRequest{
				Name: ptr.Of("Test"),
			}

			updated, err := handler.UpdateHandler(ctx, db, "nonexistent-address", request)
			require.Error(t, err)
			require.Nil(t, updated)
			require.Contains(t, err.Error(), "wallet not found")
		})
	})
}
