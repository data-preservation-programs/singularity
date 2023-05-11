package wallet

import (
	"github.com/data-preservation-programs/go-singularity/handler"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/filecoin-project/go-address"
	"github.com/jsign/go-filsigner/wallet"
	"gorm.io/gorm"
)

type ImportRequest struct {
	PrivateKey string `json:"privateKey"`
}

// ImportHandler godoc
// @Summary Import a private key
// @Tags Wallet
// @Accept json
// @Param request body ImportRequest true "Request body"
// @Success 204
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /wallet [post]
func ImportHandler(
	db *gorm.DB,
	request ImportRequest,
) *handler.Error {
	address.CurrentNetwork = address.Mainnet
	addr, err := wallet.PublicKey(request.PrivateKey)
	if err != nil {
		return handler.NewBadRequestString("invalid private key")
	}
	encryptedPrivateKey, err := model.EncryptToBase64String([]byte(request.PrivateKey))
	if err != nil {
		return handler.NewHandlerError(err)
	}

	wallet := model.Wallet{
		ID:         addr.String(),
		PrivateKey: encryptedPrivateKey,
	}

	err = db.Transaction(func(db *gorm.DB) error {
		return db.Where("ID = ?", wallet.ID).Attrs(wallet).FirstOrCreate(&wallet).Error
	})
	if err != nil {
		return handler.NewHandlerError(err)
	}

	return nil
}
