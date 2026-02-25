package wallet

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
)

// detaches wallet from preparation
// accepts wallet address or ID
func (DefaultHandler) DetachHandler(
	ctx context.Context,
	db *gorm.DB,
	preparationID string,
	walletAddressOrID string,
) (*model.Preparation, error) {
	db = db.WithContext(ctx)
	var preparation model.Preparation
	err := preparation.FindByIDOrName(db, preparationID, "SourceStorages", "OutputStorages", "Wallets")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "preparation %d not found", preparationID)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	found, err := underscore.Find(preparation.Wallets, func(w model.Wallet) bool {
		return w.Address == walletAddressOrID || fmt.Sprint(w.ID) == walletAddressOrID
	})
	if err != nil {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "wallet %s not attached to preparation %s", walletAddressOrID, preparationID)
	}

	err = database.DoRetry(ctx, func() error {
		return db.Model(&preparation).Association("Wallets").Delete(&found)
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &preparation, nil
}

// @ID DetachWallet
// @Summary Detach a new wallet from a preparation
// @Tags Wallet Association
// @Produce json
// @Accept json
// @Param id path string true "Preparation ID or name"
// @Param wallet path string true "Wallet Address"
// @Success 200 {object} model.Preparation
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation/{id}/wallet/{wallet} [delete]
func _() {}
