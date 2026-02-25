package wallet

import (
	"context"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// attaches wallet to preparation for deal-making
// accepts wallet address or wallet ID
func (DefaultHandler) AttachHandler(
	ctx context.Context,
	db *gorm.DB,
	preparationID string,
	walletAddressOrID string,
) (*model.Preparation, error) {
	db = db.WithContext(ctx)
	var preparation model.Preparation
	err := preparation.FindByIDOrName(db, preparationID, "SourceStorages", "OutputStorages", "Wallets")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "preparation %s not found", preparationID)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var w model.Wallet
	q := db.Where("address = ?", walletAddressOrID)
	if id, parseErr := strconv.ParseUint(walletAddressOrID, 10, 32); parseErr == nil {
		q = db.Where("address = ? OR id = ?", walletAddressOrID, id)
	}
	err = q.First(&w).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "wallet %s not found", walletAddressOrID)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = database.DoRetry(ctx, func() error {
		return db.Model(&preparation).Association("Wallets").Append(&w)
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &preparation, nil
}

// @ID AttachWallet
// @Summary Attach a new wallet with a preparation
// @Tags Wallet Association
// @Produce json
// @Accept json
// @Param id path string true "Preparation ID or name"
// @Param wallet path string true "Wallet Address"
// @Success 200 {object} model.Preparation
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation/{id}/wallet/{wallet} [post]
func _() {}
