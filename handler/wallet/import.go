package wallet

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/filecoin-project/go-address"
	"github.com/jsign/go-filsigner/wallet"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

type ImportRequest struct {
	PrivateKey string `json:"privateKey"` // This is the exported private key from lotus wallet export
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
func _() {}

// ImportHandler imports a wallet into the system using a given private key. It first verifies the private key's
// validity by generating its associated public address. It then checks for the existence of this address in the
// Lotus system using the provided RPC client. After confirming the actor ID from the Lotus system, it creates a
// new wallet record in the local database.
//
// Parameters:
// - ctx: The context for database transactions and other operations.
// - db: A pointer to the gorm.DB instance representing the database connection.
// - lotusClient: The RPC client used to interact with a Lotus node for actor lookup.
// - request: The request containing the private key for the wallet import operation.
//
// Returns:
// - A pointer to the created Wallet model if successful.
// - An error, if any occurred during the operation.
func ImportHandler(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	request ImportRequest,
) (*model.Wallet, error) {
	db = db.WithContext(ctx)
	addr, err := wallet.PublicKey(request.PrivateKey)
	if err != nil {
		return nil, errors.Wrap(handlererror.ErrInvalidParameter, "invalid private key")
	}

	var result string
	err = lotusClient.CallFor(ctx, &result, "Filecoin.StateLookupID", addr.String(), nil)
	if err != nil {
		return nil, errors.Join(handlererror.ErrInvalidParameter, errors.Wrap(err, "failed to lookup actor ID"))
	}

	_, err = address.NewFromString(result)
	if err != nil {
		return nil, errors.Wrap(handlererror.ErrInvalidParameter, "invalid actor ID")
	}

	wallet := model.Wallet{
		ID:         result,
		Address:    result[:1] + addr.String()[1:],
		PrivateKey: request.PrivateKey,
	}

	err = database.DoRetry(ctx, func() error {
		return db.Create(&wallet).Error
	})
	if util.IsDuplicateKeyError(err) {
		return nil, errors.Wrap(handlererror.ErrDuplicateRecord, "wallet already imported")
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &wallet, nil
}
