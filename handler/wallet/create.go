package wallet

import (
	"context"
	"crypto/rand"
	"encoding/hex"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-crypto"
	g1 "github.com/phoreproject/bls/g1pubs"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

type KeyType string

const (
	KTSecp256k1 KeyType = "secp256k1"
	KTBLS       KeyType = "bls"
	// TODO: add support for "delegated" or "secp256k1-ledger" types?
)

func (kt KeyType) String() string {
	return string(kt)
}

// GenerateKey generates a new keypair and returns the private key and address.
// The keypair is generated using the specified key type (secp256k1 or BLS).
func GenerateKey(keyType string) (string, string, error) {
	var privKey string
	var addr address.Address
	var err error

	switch keyType {
	case KTSecp256k1.String():
		kb := make([]byte, 32)
		_, err = rand.Read(kb)
		if err != nil {
			return "", "", xerrors.Errorf("failed to generate %s private key: %w", keyType, err)
		}
		privKey = hex.EncodeToString(kb)

		// Get the public key from private key
		pubKey := crypto.PublicKey(kb)
		addr, err = address.NewSecp256k1Address(pubKey)
		if err != nil {
			return "", "", xerrors.Errorf("failed to generate address from %s key: %w", keyType, err)
		}
	case KTBLS.String():
		priv, err := g1.RandKey(rand.Reader)
		if err != nil {
			return "", "", xerrors.Errorf("failed to generate %s private key: %w", keyType, err)
		}
		privKey = priv.String()

		// Get the public key from private key
		pub := g1.PrivToPub(priv)
		pubKey := pub.Serialize()
		addr, err = address.NewBLSAddress(pubKey[:])
		if err != nil {
			return "", "", xerrors.Errorf("failed to generate address from %s key: %w", keyType, err)
		}
	default:
		return "", "", xerrors.Errorf("unsupported key type: %s", keyType)
	}

	return privKey, addr.String(), nil
}

type CreateRequest struct {
	KeyType string `json:"keyType"` // This is either "secp256k1" or "bls"
}

// @ID CreateWallet
// @Summary Create new wallet
// @Tags Wallet
// @Accept json
// @Produce json
// @Param request body CreateRequest true "Request body"
// @Success 200 {array} model.Wallet
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /wallet/create [post]
func _() {}

// CreateHandler creates a new wallet using offline keypair generation and a new record in the local database.
//
// Parameters:
//   - ctx: The context for database transactions and other operations.
//   - db: A pointer to the gorm.DB instance representing the database connection.
//
// Returns:
//   - A pointer to the created Wallet model if successful.
//   - An error, if any occurred during the database insert operation.
func (DefaultHandler) CreateHandler(
	ctx context.Context,
	db *gorm.DB,
	request CreateRequest,
) (*model.Wallet, error) {
	db = db.WithContext(ctx)

	// Generate a new keypair
	privateKey, address, err := GenerateKey(request.KeyType)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	wallet := model.Wallet{
		ID:         address,
		Address:    address,
		PrivateKey: privateKey,
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
