package wallet

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func (DefaultHandler) ListAttachedHandler(
	ctx context.Context,
	db *gorm.DB,
	preparationID uint32,
) ([]model.Wallet, error) {
	var preparation model.Preparation
	err := db.Preload("Wallets").Where("id = ?", preparationID).First(&preparation).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "preparation %d not found", preparationID)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return preparation.Wallets, nil
}

// @Summary List all wallets of a preparation.
// @Tags Wallet Association
// @Produce json
// @Accept json
// @Param id path int true "Preparation ID"
// @Success 200 {object} model.Wallet
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation/{id}/wallet [post]
func _() {}
