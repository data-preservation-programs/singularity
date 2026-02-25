package wallet

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/keystore"
	"gorm.io/gorm"
)

// removes wallet record and keystore file
// does not remove the associated actor — actors may be shared or tracked independently
func (DefaultHandler) RemoveHandler(
	ctx context.Context,
	db *gorm.DB,
	ks keystore.KeyStore,
	address string,
) error {
	db = db.WithContext(ctx)
	var w model.Wallet
	err := db.Where("address = ?", address).First(&w).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrapf(handlererror.ErrNotFound, "wallet %s not found", address)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	err = database.DoRetry(ctx, func() error {
		return db.Delete(&w).Error
	})
	if err != nil {
		return errors.WithStack(err)
	}

	// best-effort keystore cleanup
	if ks != nil && ks.Has(w.KeyPath) {
		if delErr := ks.Delete(w.KeyPath); delErr != nil {
			logger.Warnw("failed to delete key file", "path", w.KeyPath, "err", delErr)
		}
	}

	return nil
}

// @ID RemoveWallet
// @Summary Remove a wallet
// @Tags Wallet
// @Param address path string true "Address"
// @Success 204
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /wallet/{address} [delete]
func _() {}
