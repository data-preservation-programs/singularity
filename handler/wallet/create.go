package wallet

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/filecoin-project/go-address"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

type ExportResponse struct {
	PrivateKey string `json:"privateKey"` // This is the exported private key from lotus wallet export
}

// @ID CreateWallet
// @Summary Create new wallet
// @Tags Wallet
// @Produce json
// @Success 200 {array} model.Wallet
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /wallet/create [post]
func _() {}

// CreateHandler creates a new wallet using the provided Lotus RPC client and creates a new wallet record in the local database.
//
// Parameters:
//   - ctx: The context for database transactions and other operations.
//   - db: A pointer to the gorm.DB instance representing the database connection.
//   - lotusClient: The RPC client used to interact with a Lotus node for actor lookup.
//
// Returns:
//   - A pointer to the created Wallet model if successful.
//   - An error, if any occurred during the database insert operation.
func (DefaultHandler) CreateHandler(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
) (*model.Wallet, error) {
	db = db.WithContext(ctx)

	var result string
	err := lotusClient.CallFor(ctx, &result, "WalletNew")
	if err != nil {
		logger.Errorw("failed to create new wallet", "err", err)
		return nil, errors.WithStack(err)
	}

	addr, err := address.NewFromString(result)
	if err != nil {
		return nil, errors.Wrap(handlererror.ErrInvalidParameter, "invalid actor ID")
	}

	var export ExportResponse
	err = lotusClient.CallFor(ctx, &export, "WalletExport", addr.String())
	if err != nil {
		logger.Errorw("failed to export wallet", "addr", addr, "err", err)
		return nil, errors.WithStack(err)
	}

	wallet := model.Wallet{
		ID: result,
		// HACK: this ensures the address starts with the correct network prefix
		// see ./import.go for more details
		Address:    result[:1] + addr.String()[1:],
		PrivateKey: export.PrivateKey,
	}
	err = database.DoRetry(ctx, func() error {
		return db.Create(&wallet).Error
	})
	if util.IsDuplicateKeyError(err) {
		return nil, errors.Wrap(handlererror.ErrDuplicateRecord, "wallet already exists")
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &wallet, nil
}
