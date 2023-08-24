package wallet

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// ListHandler retrieves a list of all the wallets stored in the database.
//
// Parameters:
// - ctx: The context for database transactions and other operations.
// - db: A pointer to the gorm.DB instance representing the database connection.
//
// Returns:
// - A slice containing all Wallet models from the database.
// - An error, if any occurred during the database fetch operation.
func ListHandler(
	ctx context.Context,
	db *gorm.DB,
) ([]model.Wallet, error) {
	db = db.WithContext(ctx)
	var wallets []model.Wallet

	err := db.Find(&wallets).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return wallets, nil
}

// @Summary List all imported wallets
// @Tags Wallet
// @Produce json
// @Success 200 {array} model.Wallet
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /wallet [get]
func _() {}
