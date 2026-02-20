package wallet

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/keystore"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

// loads private key from keystore and signs a filecoin message
func SignWithWallet(ks keystore.KeyStore, wallet model.Wallet, msg []byte) (*crypto.Signature, error) {
	s, err := keystore.Signer(ks, wallet)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load signer from keystore")
	}

	signature, err := s.Sign(msg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign message")
	}

	logger.Debugw("signed message", "address", wallet.Address, "msgLen", len(msg))
	return signature, nil
}

// lazy actor lookup and creation for a wallet
// workflow: import wallet offline → fund externally → first deal queries on-chain actor
// returns existing actor if wallet.ActorID already set, otherwise queries lotus and creates record
func GetOrCreateActor(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	wallet *model.Wallet,
) (*model.Actor, error) {
	db = db.WithContext(ctx)

	// return existing actor if already linked
	if wallet.ActorID != nil {
		var actor model.Actor
		err := db.First(&actor, "id = ?", *wallet.ActorID).Error
		if err != nil {
			return nil, errors.Wrapf(err, "actor %s not found in database", *wallet.ActorID)
		}
		logger.Debugw("wallet already linked to actor", "walletID", wallet.ID, "actorID", actor.ID)
		return &actor, nil
	}

	// query lotus for on-chain actor
	logger.Infow("looking up actor on-chain", "address", wallet.Address)

	var actorID string
	err := lotusClient.CallFor(ctx, &actorID, "Filecoin.StateLookupID", wallet.Address, nil)
	if err != nil {
		logger.Warnw("actor not found on-chain", "address", wallet.Address, "err", err)
		return nil, errors.Wrapf(err, "actor for address %s not found on-chain - wallet may need funding", wallet.Address)
	}

	logger.Infow("found actor on-chain", "address", wallet.Address, "actorID", actorID)

	// check if actor already exists in database
	var existingActor model.Actor
	err = db.First(&existingActor, "id = ?", actorID).Error
	if err == nil {
		// actor exists - verify not linked to different wallet
		var otherWallet model.Wallet
		err = db.Where("actor_id = ?", actorID).First(&otherWallet).Error
		if err == nil && otherWallet.ID != wallet.ID {
			logger.Warnw("actor already linked to different wallet",
				"actorID", actorID,
				"existingWalletID", otherWallet.ID,
				"newWalletID", wallet.ID)
			return nil, errors.Errorf("actor %s already linked to wallet %d", actorID, otherWallet.ID)
		}

		// link to this wallet
		wallet.ActorID = &actorID
		err = db.Save(wallet).Error
		if err != nil {
			return nil, errors.Wrap(err, "failed to link wallet to existing actor")
		}

		logger.Infow("linked wallet to existing actor", "walletID", wallet.ID, "actorID", actorID)
		return &existingActor, nil
	}

	// create new actor record
	newActor := model.Actor{
		ID:      actorID,
		Address: wallet.Address,
	}

	err = db.Create(&newActor).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to create actor record")
	}

	// link wallet to new actor
	wallet.ActorID = &actorID
	err = db.Save(wallet).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to link wallet to new actor")
	}

	logger.Infow("created actor and linked to wallet",
		"walletID", wallet.ID,
		"actorID", actorID,
		"address", wallet.Address)

	return &newActor, nil
}

// loads wallet by actor ID for signing operations
func LoadWalletByActorID(ctx context.Context, db *gorm.DB, actorID string) (*model.Wallet, error) {
	db = db.WithContext(ctx)

	var wallet model.Wallet
	err := db.Where("actor_id = ?", actorID).First(&wallet).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Errorf("no wallet found for actor %s - actor may not be controlled by this instance", actorID)
		}
		return nil, errors.Wrap(err, "failed to query wallet by actor ID")
	}

	return &wallet, nil
}
