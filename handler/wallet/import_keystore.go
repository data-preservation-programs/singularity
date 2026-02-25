package wallet

import (
	"context"
	"os"
	"path/filepath"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/data-preservation-programs/singularity/util/keystore"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

var logger = log.Logger("singularity/handler/wallet")

type ImportKeystoreRequest struct {
	PrivateKey string `json:"privateKey"` // lotus wallet export format
	Name       string `json:"name"`       // optional human-readable name
}

// @ID ImportWallet
// @Summary Import a private key
// @Tags Wallet
// @Accept json
// @Produce json
// @Param request body ImportKeystoreRequest true "Request body"
// @Success 200 {object} model.Wallet
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /wallet [post]
func _() {}

// imports wallet by saving private key to keystore and creating wallet record
// does not require actor to exist on-chain - wallet can be imported offline
// uses external keystore instead of storing keys in database
func (DefaultHandler) ImportKeystoreHandler(
	ctx context.Context,
	db *gorm.DB,
	ks keystore.KeyStore,
	request ImportKeystoreRequest,
) (*model.Wallet, error) {
	db = db.WithContext(ctx)

	// validate key before touching keystore: derive address to distinguish
	// bad input (400) from keystore I/O failures (500)
	addr, err := keystore.AddressFromExport(request.PrivateKey)
	if err != nil {
		return nil, errors.Wrap(handlererror.ErrInvalidParameter, err.Error())
	}

	keyPath, _, err := ks.Put(request.PrivateKey)
	if err != nil {
		logger.Errorw("failed to save key to keystore", "err", err)
		return nil, errors.WithStack(err)
	}

	logger.Infow("saved key to keystore", "address", addr.String(), "path", keyPath)

	walletRecord := model.Wallet{
		KeyPath:  keyPath,
		KeyStore: "local",
		Address:  addr.String(),
		Name:     request.Name,
		ActorID:  nil, // populated lazily when needed
	}

	err = database.DoRetry(ctx, func() error {
		return db.Create(&walletRecord).Error
	})

	if util.IsDuplicateKeyError(err) {
		// don't delete the key file — it belongs to the existing wallet record
		return nil, errors.Wrap(handlererror.ErrDuplicateRecord, "wallet already imported")
	}

	if err != nil {
		ks.Delete(keyPath) // cleanup only for non-duplicate failures
		return nil, errors.WithStack(err)
	}

	logger.Infow("imported wallet", "id", walletRecord.ID, "address", addr.String())

	return &walletRecord, nil
}

// returns default keystore directory path
// TODO: make configurable via config file
func GetKeystoreDir() string {
	if dir := os.Getenv("SINGULARITY_KEYSTORE"); dir != "" {
		return dir
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return filepath.Join(".", ".singularity", "keystore")
	}
	return filepath.Join(home, ".singularity", "keystore")
}
