package wallet

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-crypto"
	g1 "github.com/phoreproject/bls/g1pubs"
	"github.com/ybbus/jsonrpc/v3"
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

		// Format the private key as a Lotus exported key (JSON format)
		privKeyJSON := map[string]interface{}{
			"Type":       "secp256k1",
			"PrivateKey": base64.StdEncoding.EncodeToString(kb),
		}

		privKeyBytes, err := json.Marshal(privKeyJSON)
		if err != nil {
			return "", "", xerrors.Errorf("failed to marshal private key to JSON: %w", err)
		}
		privKey = hex.EncodeToString(privKeyBytes)

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

		// Format the private key as a Lotus exported key (JSON format)
		// Convert the private key to base64 format
		privKeyBytes := priv.Serialize()
		privKeyJSON := map[string]interface{}{
			"Type":       "bls",
			"PrivateKey": base64.StdEncoding.EncodeToString(privKeyBytes[:]),
		}

		privKeyJSONBytes, err := json.Marshal(privKeyJSON)
		if err != nil {
			return "", "", xerrors.Errorf("failed to marshal private key to JSON: %w", err)
		}
		privKey = hex.EncodeToString(privKeyJSONBytes)

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
	// For UserWallet creation (generates new keypair)
	KeyType string `json:"keyType,omitempty"` // This is either "secp256k1" or "bls"

	// For SPWallet creation
	Address string `json:"address,omitempty"`
	ActorID string `json:"actorId,omitempty"`

	// Optional fields for adding details to Wallet
	Name     string `json:"name,omitempty"`
	Contact  string `json:"contact,omitempty"`
	Location string `json:"location,omitempty"`
}

// @ID CreateWallet
// @Summary Create new wallet
// @Tags Wallet
// @Accept json
// @Produce json
// @Param request body CreateRequest true "Request body"
// @Success 200 {object} model.Wallet
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /wallet/create [post]
func _() {}

// CreateHandler creates a new wallet using offline keypair generation and a new record in the local database.
// CreateHandler creates a new wallet and stores it in the local database.
// The wallet type is automatically inferred from the provided parameters:
// - If KeyType is provided: creates a UserWallet with generated keypair
// - If Address is provided: creates an SPWallet contact entry
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
//   - lotusClient: The RPC client used to interact with a Lotus node for actor lookup (only used for SP wallets).
//   - request: CreateRequest with either KeyType (for UserWallet) or Address (for SPWallet)
//
// Returns:
//   - A pointer to the created Wallet model if successful.
//   - An error, if any occurred during validation or database operations.
func (DefaultHandler) CreateHandler(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	request CreateRequest,
) (*model.Wallet, error) {
	db = db.WithContext(ctx)

	// Generate a new keypair
	privateKey, address, err := GenerateKey(request.KeyType)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	wallet := model.Wallet{
		Address:    address,
		PrivateKey: privateKey,
	}
	err = database.DoRetry(ctx, func() error {
	// Infer wallet type from provided parameters
	hasKeyType := request.KeyType != ""
	hasAddress := request.Address != ""
	hasActorID := request.ActorID != ""

	// Validate that only one wallet type is specified
	switch {
	case !hasKeyType && !hasAddress && !hasActorID:
		return nil, errors.New("must specify either KeyType (for UserWallet) or Address/ActorID (for SPWallet)")
	case !hasKeyType && !(hasAddress && hasActorID):
		return nil, errors.New("must specify both Address and ActorID (for SPWallet)")
	case hasKeyType && (hasAddress || hasActorID):
		return nil, errors.New("cannot specify both KeyType (for UserWallet) and Address/ActorID (for SPWallet) - please specify parameters for one wallet type")
	}

	var wallet model.Wallet

	if hasKeyType {
		// Create UserWallet: generate a new keypair
		privateKey, address, err := GenerateKey(request.KeyType)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		wallet = model.Wallet{
			Address:    address,
			PrivateKey: privateKey,
			WalletType: model.UserWallet,
			// ActorID is empty for UserWallets until initialized
		}
	} else {
		// Validate the address and actor ID with Lotus
		addr, err := address.NewFromString(request.Address)
		if err != nil {
			return nil, errors.Wrap(handlererror.ErrInvalidParameter, "invalid address format")
		}

		var result string
		err = lotusClient.CallFor(ctx, &result, "Filecoin.StateLookupID", addr.String(), nil)
		if err != nil {
			logger.Errorw("failed to lookup state for wallet address", "addr", addr, "err", err)
			return nil, errors.Join(handlererror.ErrInvalidParameter, errors.Wrap(err, "failed to lookup actor ID"))
		}

		_, err = address.NewFromString(result)
		if err != nil {
			return nil, errors.Wrap(handlererror.ErrInvalidParameter, "invalid actor ID")
		} else if result != request.ActorID {
			return nil, errors.Wrap(handlererror.ErrInvalidParameter, "provided actor ID is not associated with address")
		}

		wallet = model.Wallet{
			ActorID:    result,
			Address:    result[:1] + addr.String()[1:],
			WalletType: model.SPWallet,
			// PrivateKey is empty for SP wallets
		}
	}

	// Update wallet details
	wallet.ActorName = request.Name
	wallet.ContactInfo = request.Contact
	wallet.Location = request.Location

	err := database.DoRetry(ctx, func() error {
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
