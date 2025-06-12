package wallet

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

type UpdateRequest struct {
	Name     *string `json:"actorName,omitempty"`   // Name is readable label for the wallet
	Contact  *string `json:"contactInfo,omitempty"` // Contact is optional email for SP wallets
	Location *string `json:"location,omitempty"`    // Location is optional region, country for SP wallets
}

// @ID UpdateWallet
// @Summary Update wallet details
// @Tags Wallet
// @Accept json
// @Produce json
// @Param address path string true "Wallet address"
// @Param request body UpdateRequest true "Request body"
// @Success 200 {object} model.Wallet
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /wallet/{address}/update [patch]
func _() {}

// UpdateHandler updates non-essential details of an existing wallet.
// Only fields provided in the request will be updated (partial update).
// Essential fields like Address, PrivateKey, and Balance cannot be modified.
//
// Parameters:
//   - ctx: The context for database transactions and other operations.
//   - db: A pointer to the gorm.DB instance representing the database connection.
//   - address: The wallet address to update.
//   - request: UpdateRequest containing the fields to update.
//
// Returns:
//   - A pointer to the updated Wallet model if successful.
//   - An error, if any occurred during the database operation.
func (DefaultHandler) UpdateHandler(
	ctx context.Context,
	db *gorm.DB,
	address string,
	request UpdateRequest,
) (*model.Wallet, error) {
	db = db.WithContext(ctx)

	// Find the existing wallet
	var wallet model.Wallet
	err := wallet.FindByIDOrAddr(db, address)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrap(handlererror.ErrNotFound, "wallet not found")
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Prepare update data - only include fields that are provided
	updates := make(map[string]interface{})

	if request.Name != nil {
		updates["actor_name"] = *request.Name
	}

	if request.Contact != nil {
		updates["contact_info"] = *request.Contact
	}

	if request.Location != nil {
		updates["location"] = *request.Location
	}

	// If no fields to update, return the existing wallet
	if len(updates) == 0 {
		return &wallet, nil
	}

	// Perform the update
	err = database.DoRetry(ctx, func() error {
		return db.Model(&wallet).Updates(updates).Error
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Fetch the updated wallet to return
	err = database.DoRetry(ctx, func() error {
		return db.Where("address = ?", address).First(&wallet).Error
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &wallet, nil
}
