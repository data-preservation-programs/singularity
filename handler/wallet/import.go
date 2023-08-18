package wallet

import (
	"context"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/filecoin-project/go-address"
	"github.com/jsign/go-filsigner/wallet"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

type ImportRequest struct {
	PrivateKey string `json:"privateKey"` // This is the exported private key from lotus wallet export
}

func ImportHandler(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	request ImportRequest,
) (*model.Wallet, error) {
	return importHandler(ctx, db.WithContext(ctx), lotusClient, request)
}

// @Summary Import a private key
// @Tags Wallet
// @Accept json
// @Produce json
// @Param request body ImportRequest true "Request body"
// @Success 200 {object} model.Wallet
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /wallet [post]
func importHandler(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	request ImportRequest,
) (*model.Wallet, error) {
	addr, err := wallet.PublicKey(request.PrivateKey)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("invalid private key")
	}

	var result string
	err = lotusClient.CallFor(ctx, &result, "Filecoin.StateLookupID", addr.String(), nil)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("invalid private key")
	}

	_, err = address.NewFromString(result)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("invalid actor ID from GLIF result: " + result)
	}

	wallet := model.Wallet{
		ID:         result,
		Address:    result[:1] + addr.String()[1:],
		PrivateKey: request.PrivateKey,
	}

	err = database.DoRetry(ctx, func() error {
		return db.Create(&wallet).Error
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &wallet, nil
}
