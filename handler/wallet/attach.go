package wallet

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// attaches actor to preparation for deal-making
// accepts actor ID (f0...) or wallet address/ID
// wallet must already be linked to on-chain actor
func (DefaultHandler) AttachHandler(
	ctx context.Context,
	db *gorm.DB,
	preparationID string,
	actorOrWallet string,
) (*model.Preparation, error) {
	db = db.WithContext(ctx)
	var preparation model.Preparation
	err := preparation.FindByIDOrName(db, preparationID, "SourceStorages", "OutputStorages", "Actors")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "preparation %s not found", preparationID)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// try to find as actor ID first
	var actor model.Actor
	err = db.Where("id = ?", actorOrWallet).First(&actor).Error
	if err == nil {
		// found actor directly
		err = database.DoRetry(ctx, func() error {
			return db.Model(&preparation).Association("Actors").Append(&actor)
		})
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return &preparation, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.WithStack(err)
	}

	// not found as actor, try as wallet address or ID
	var wallet model.Wallet
	err = db.Where("address = ? OR id = ?", actorOrWallet, actorOrWallet).First(&wallet).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "actor or wallet %s not found", actorOrWallet)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// wallet found - check if it has an actor
	if wallet.ActorID == nil || *wallet.ActorID == "" {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter,
			"wallet %s not yet linked to on-chain actor - fund the wallet first", wallet.Address)
	}

	// get the actor
	err = db.Where("id = ?", *wallet.ActorID).First(&actor).Error
	if err != nil {
		return nil, errors.Wrapf(err, "actor %s not found", *wallet.ActorID)
	}

	err = database.DoRetry(ctx, func() error {
		return db.Model(&preparation).Association("Actors").Append(&actor)
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &preparation, nil
}

// @ID AttachWallet
// @Summary Attach a new wallet with a preparation
// @Tags Wallet Association
// @Produce json
// @Accept json
// @Param id path string true "Preparation ID or name"
// @Param wallet path string true "Wallet Address"
// @Success 200 {object} model.Preparation
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation/{id}/wallet/{wallet} [post]
func _() {}
