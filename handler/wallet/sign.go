package wallet

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/keystore"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/jsign/go-filsigner/wallet"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

// loads private key from keystore and signs message
// new signing flow - loads keys from disk instead of database
func SignWithWallet(ks keystore.KeyStore, walletKey model.WalletKey, msg []byte) (*crypto.Signature, error) {
	privateKey, err := ks.Get(walletKey.KeyPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load private key from keystore")
	}

	// wallet.WalletSign automatically detects key type (secp256k1 or BLS)
	signature, err := wallet.WalletSign(privateKey, msg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign message")
	}

	logger.Debugw("signed message", "address", walletKey.Address, "msgLen", len(msg))
	return signature, nil
}

// lazy actor lookup and creation for a wallet
// workflow: import wallet offline → fund externally → first deal queries on-chain actor
// returns existing actor if wallet.ActorID already set, otherwise queries lotus and creates record
// TODO: after step 6 rename, return type will be *model.Actor instead of *model.Wallet
func GetOrCreateActor(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	walletKey *model.WalletKey,
) (*model.Wallet, error) {
	db = db.WithContext(ctx)

	// return existing actor if already linked
	if walletKey.ActorID != nil {
		var actor model.Wallet
		err := db.First(&actor, "id = ?", *walletKey.ActorID).Error
		if err != nil {
			return nil, errors.Wrapf(err, "actor %s not found in database", *walletKey.ActorID)
		}
		logger.Debugw("wallet already linked to actor", "walletID", walletKey.ID, "actorID", actor.ID)
		return &actor, nil
	}

	// query lotus for on-chain actor
	logger.Infow("looking up actor on-chain", "address", walletKey.Address)

	var actorID string
	err := lotusClient.CallFor(ctx, &actorID, "Filecoin.StateLookupID", walletKey.Address, nil)
	if err != nil {
		logger.Warnw("actor not found on-chain", "address", walletKey.Address, "err", err)
		return nil, errors.Wrapf(err, "actor for address %s not found on-chain - wallet may need funding", walletKey.Address)
	}

	logger.Infow("found actor on-chain", "address", walletKey.Address, "actorID", actorID)

	// check if actor already exists in database
	var existingActor model.Wallet
	err = db.First(&existingActor, "id = ?", actorID).Error
	if err == nil {
		// actor exists - verify not linked to different wallet
		var otherWallet model.WalletKey
		err = db.Where("actor_id = ?", actorID).First(&otherWallet).Error
		if err == nil && otherWallet.ID != walletKey.ID {
			logger.Warnw("actor already linked to different wallet",
				"actorID", actorID,
				"existingWalletID", otherWallet.ID,
				"newWalletID", walletKey.ID)
			return nil, errors.Errorf("actor %s already linked to wallet %d", actorID, otherWallet.ID)
		}

		// link to this wallet
		walletKey.ActorID = &actorID
		err = db.Save(walletKey).Error
		if err != nil {
			return nil, errors.Wrap(err, "failed to link wallet to existing actor")
		}

		logger.Infow("linked wallet to existing actor", "walletID", walletKey.ID, "actorID", actorID)
		return &existingActor, nil
	}

	// create new actor record
	newActor := model.Wallet{
		ID:      actorID,
		Address: walletKey.Address,
		// TODO: after step 6 rename, this becomes model.Actor without PrivateKey field
	}

	err = db.Create(&newActor).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to create actor record")
	}

	// link wallet to new actor
	walletKey.ActorID = &actorID
	err = db.Save(walletKey).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to link wallet to new actor")
	}

	logger.Infow("created actor and linked to wallet",
		"walletID", walletKey.ID,
		"actorID", actorID,
		"address", walletKey.Address)

	return &newActor, nil
}

// loads wallet by actor ID for signing operations
func LoadWalletKeyByActorID(ctx context.Context, db *gorm.DB, actorID string) (*model.WalletKey, error) {
	db = db.WithContext(ctx)

	var walletKey model.WalletKey
	err := db.Where("actor_id = ?", actorID).First(&walletKey).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Errorf("no wallet found for actor %s - actor may not be controlled by this instance", actorID)
		}
		return nil, errors.Wrap(err, "failed to query wallet by actor ID")
	}

	return &walletKey, nil
}
