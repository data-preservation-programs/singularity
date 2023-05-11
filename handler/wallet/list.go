package wallet

import (
	"github.com/data-preservation-programs/go-singularity/handler"
	"github.com/data-preservation-programs/go-singularity/model"
	"gorm.io/gorm"
)

// ListHandler godoc
// @Summary List all imported wallets
// @Tags Wallet
// @Produce json
// @Success 200 {array} model.Wallet
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /wallets [get]
func ListHandler(
	db *gorm.DB,
) ([]model.Wallet, *handler.Error) {
	var wallets []model.Wallet

	err := db.Find(&wallets).Error
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	for i, _ := range wallets {
		privateKey, err := model.DecryptFromBase64String(wallets[i].PrivateKey)
		if err != nil {
			return nil, handler.NewHandlerError(err)
		}
		wallets[i].PrivateKey = string(privateKey)
	}

	return wallets, nil
}
