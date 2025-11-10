package wallet

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
)

// detaches actor from preparation
// accepts actor ID or address
func (DefaultHandler) DetachHandler(
	ctx context.Context,
	db *gorm.DB,
	preparationID string,
	actorIDOrAddress string,
) (*model.Preparation, error) {
	db = db.WithContext(ctx)
	var preparation model.Preparation
	err := preparation.FindByIDOrName(db, preparationID, "SourceStorages", "OutputStorages", "Actors")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "preparation %d not found", preparationID)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	found, err := underscore.Find(preparation.Actors, func(a model.Actor) bool {
		return a.ID == actorIDOrAddress || a.Address == actorIDOrAddress
	})
	if err != nil {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "actor %s not attached to preparation %s", actorIDOrAddress, preparationID)
	}

	err = database.DoRetry(ctx, func() error {
		return db.Model(&preparation).Association("Actors").Delete(&found)
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &preparation, nil
}

// @ID DetachWallet
// @Summary Detach a new wallet from a preparation
// @Tags Wallet Association
// @Produce json
// @Accept json
// @Param id path string true "Preparation ID or name"
// @Param wallet path string true "Wallet Address"
// @Success 200 {object} model.Preparation
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation/{id}/wallet/{wallet} [delete]
func _() {}
