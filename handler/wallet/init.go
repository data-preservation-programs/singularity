package wallet

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

// @ID InitWallet
// @Summary Initialize a newly created wallet
// @Tags Wallet
// @Produce json
// @Param address path string true "Address"
// @Success 200 {object} model.Wallet
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /wallet/{address}/init [post]
func _() {}

// InitHandler marks a new wallet created from offline keypair generation as initialized by updating its ActorID
//
// Parameters:
//   - ctx: The context for database transactions and other operations.
//   - db: A pointer to the gorm.DB instance representing the database connection.
//   - lotusClient: The RPC client used to interact with a Lotus node for actor lookup.
//   - address: The address or ID of the wallet to be initialized.
//
// Returns:
//   - A pointer to the initialized Wallet model if successful.
//   - An error, if any occurred during the database insert operation.
func (DefaultHandler) InitHandler(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	address string,
) (*model.Wallet, error) {
	db = db.WithContext(ctx)
	var wallet model.Wallet
	err := wallet.FindByIDOrAddr(db, address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find wallet")
	}

	if wallet.ActorID != "" {
		return &wallet, nil
	}

	var result string
	err = lotusClient.CallFor(ctx, &result, "Filecoin.StateLookupID", wallet.Address, nil)
	if err != nil {
		return nil, errors.Join(handlererror.ErrInvalidParameter, errors.Wrap(err, "failed to lookup actor ID"))
	}

	wallet.ActorID = result
	err = database.DoRetry(ctx, func() error {
		return db.Save(&wallet).Error
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &wallet, nil
}
