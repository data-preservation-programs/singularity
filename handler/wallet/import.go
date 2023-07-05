package wallet

import (
	"context"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/filecoin-project/go-address"
	"github.com/jsign/go-filsigner/wallet"
	"gorm.io/gorm"
)

type ImportRequest struct {
	PrivateKey string `json:"privateKey"` // This is the exported private key from lotus wallet export
	LotusAPI   string `json:"lotusApi"   swaggerignore:"true"`
	LotusToken string `json:"lotusToken" swaggerignore:"true"`
}

// ImportHandler godoc
// @Summary Import a private key
// @Tags Wallet
// @Accept json
// @Produce json
// @Param request body ImportRequest true "Request body"
// @Success 200 {object} model.Wallet
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /wallet [post]
func ImportHandler(
	db *gorm.DB,
	request ImportRequest,
) (*model.Wallet, *handler.Error) {
	address.CurrentNetwork = address.Mainnet
	addr, err := wallet.PublicKey(request.PrivateKey)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid private key")
	}

	lotusClient := util.NewLotusClient(request.LotusAPI, request.LotusToken)
	var result string
	err = lotusClient.CallFor(context.Background(), &result, "Filecoin.StateLookupID", addr.String(), nil)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid private key")
	}

	wallet := model.Wallet{
		ID:         result,
		Address:    addr.String(),
		PrivateKey: request.PrivateKey,
	}

	err = db.Transaction(func(db *gorm.DB) error {
		return db.Create(&wallet).Error
	})
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	return &wallet, nil
}
