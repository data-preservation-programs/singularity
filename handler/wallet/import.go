package wallet

import (
	"context"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/filecoin-project/go-address"
	"github.com/jsign/go-filsigner/wallet"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

type ImportRequest struct {
	PrivateKey string `json:"privateKey"` // This is the exported private key from lotus wallet export
	LotusAPI   string `json:"lotusApi" swaggerignore:"true"`
	LotusToken string `json:"lotusToken" swaggerignore:"true"`
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

	var lotusClient jsonrpc.RPCClient
	if request.LotusToken == "" {
		lotusClient = jsonrpc.NewClient(request.LotusAPI)
	} else {
		lotusClient = jsonrpc.NewClientWithOpts(request.LotusAPI, &jsonrpc.RPCClientOpts{
			CustomHeaders: map[string]string{
				"Authorization": "Bearer " + request.LotusToken,
			},
		})
	}

	var result string
	err = lotusClient.CallFor(context.Background(), &result, "Filecoin.StateLookupID", addr.String(), nil)
	if err != nil {
		return handler.NewBadRequestString("invalid private key")
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
		return handler.NewHandlerError(err)
	}

	return nil
}
