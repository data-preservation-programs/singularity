package wallet

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func ListHandler(
	ctx context.Context,
	db *gorm.DB,
) ([]model.Wallet, error) {
	return listHandler(db.WithContext(ctx))
}

// @Summary List all imported wallets
// @Tags Wallet
// @Produce json
// @Success 200 {array} model.Wallet
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /wallet [get]
func listHandler(
	db *gorm.DB,
) ([]model.Wallet, error) {
	var wallets []model.Wallet

	err := db.Find(&wallets).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return wallets, nil
}
