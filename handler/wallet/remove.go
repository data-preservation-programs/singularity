package wallet

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// RemoveHandler deletes a wallet from the database based on its address or ID.
//
// Parameters:
// - ctx: The context for database transactions and other operations.
// - db: A pointer to the gorm.DB instance representing the database connection.
// - address: The address or ID of the wallet to be deleted.
//
// Returns:
// - An error, if any occurred during the database deletion operation.
func RemoveHandler(
	ctx context.Context,
	db *gorm.DB,
	address string,
) error {
	db = db.WithContext(ctx)
	var affected int64
	err := database.DoRetry(ctx, func() error {
		tx := db.Where("address = ? OR id = ?", address, address).Delete(&model.Wallet{})
		affected = tx.RowsAffected
		return tx.Error
	})
	if err != nil {
		return errors.WithStack(err)
	}
	if affected == 0 {
		return errors.Wrapf(handlererror.ErrNotFound, "wallet %s not found", address)
	}
	return nil
}

// @Summary Remove a wallet
// @Tags Wallet
// @Param address path string true "Address"
// @Success 204
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /wallet/{address} [delete]
func _() {}
