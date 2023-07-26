package wallet

import (
	"context"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/jsign/go-filsigner/wallet"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

type ImportRequest struct {
	PrivateKey string `json:"privateKey"` // This is the exported private key from lotus wallet export
}

func ImportHandler(
	db *gorm.DB,
	ctx context.Context,
	lotusClient jsonrpc.RPCClient,
	request ImportRequest,
) (*model.Wallet, error) {
	return importHandler(db, ctx, lotusClient, request)
}

// @Summary Import a private key
// @Tags Wallet
// @Accept json
// @Produce json
// @Param request body ImportRequest true "Request body"
// @Success 200 {object} model.Wallet
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /wallet [post]
func importHandler(
	db *gorm.DB,
	ctx context.Context,
	lotusClient jsonrpc.RPCClient,
	request ImportRequest,
) (*model.Wallet, error) {
	addr, err := wallet.PublicKey(request.PrivateKey)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid private key")
	}

	var result string
	err = lotusClient.CallFor(ctx, &result, "Filecoin.StateLookupID", addr.String(), nil)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid private key")
	}

	wallet := model.Wallet{
		ID:         result,
		Address:    result[:1] + addr.String()[1:],
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
