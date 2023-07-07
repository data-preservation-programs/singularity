package wallet

import (
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// ListHandler godoc
// @Summary List all imported wallets
// @Tags Wallet
// @Produce json
// @Success 200 {array} model.Wallet
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /wallet [get]
func ListHandler(
	db *gorm.DB,
) ([]model.Wallet, error) {
	var wallets []model.Wallet

	err := db.Find(&wallets).Error
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	return wallets, nil
}
