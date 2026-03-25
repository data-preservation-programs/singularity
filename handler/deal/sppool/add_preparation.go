package sppool

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"gorm.io/gorm"
)

type AddPreparationRequest struct {
	Preparation string `json:"preparation" validation:"required"` // Preparation ID or name
}

func (DefaultHandler) AddPreparationHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint32,
	request AddPreparationRequest,
) (*model.SPPoolPreparation, error) {
	db = db.WithContext(ctx)

	// Validate pool exists.
	var pool model.SPPool
	err := db.Preload("Providers").Preload("Preparations").First(&pool, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "sp pool %d not found", id)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Resolve preparation by ID or name.
	var preparation model.Preparation
	err = preparation.FindByIDOrName(db, request.Preparation, "Wallet")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "preparation %s not found", request.Preparation)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if preparation.WalletID == nil {
		return nil, errors.Wrap(handlererror.ErrInvalidParameter, "preparation has no wallet attached")
	}

	poolPrep := model.SPPoolPreparation{
		PoolID:        model.SPPoolID(id),
		PreparationID: preparation.ID,
	}

	if err := database.DoRetry(ctx, func() error {
		return db.Create(&poolPrep).Error
	}); err != nil {
		if util.IsDuplicateKeyError(err) {
			return nil, errors.Wrapf(handlererror.ErrDuplicateRecord, "preparation %d already in pool %d", preparation.ID, id)
		}
		return nil, errors.WithStack(err)
	}

	// Reconcile to create schedules for the new preparation.
	pool.Preparations = append(pool.Preparations, poolPrep)
	if err := reconcile(ctx, db, &pool); err != nil {
		return nil, err
	}

	return &poolPrep, nil
}

// @ID AddSPPoolPreparation
// @Summary Add a preparation to an SP Pool
// @Tags SP Pool
// @Accept json
// @Produce json
// @Param id path int true "SP Pool ID"
// @Param request body AddPreparationRequest true "AddPreparationRequest"
// @Success 200 {object} model.SPPoolPreparation
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Failure 409 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /sp-pool/{id}/preparation [post]
func _() {}
