package dataprep

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/rclone/rclone/fs"
	"gorm.io/gorm"
)

type RemoveRequest struct {
	RemoveCars bool `json:"removeCars"`
}

func (DefaultHandler) RemovePreparationHandler(ctx context.Context, db *gorm.DB, name string, request RemoveRequest) error {
	db = db.WithContext(ctx)

	var preparation model.Preparation
	err := preparation.FindByIDOrName(db, name)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrapf(handlererror.ErrNotFound, "preparation %s does not exist", name)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	var cars []model.Car
	if request.RemoveCars {
		err = db.Preload("Storage").Where("preparation_id = ?", preparation.ID).Find(&cars).Error
		if err != nil {
			return errors.WithStack(err)
		}
	}

	err = database.DoRetry(ctx, func() error {
		return db.Transaction(func(db *gorm.DB) error {
			return db.Delete(&preparation).Error
		})
	})

	if err != nil {
		return errors.WithStack(err)
	}

	storageHandlers := make(map[model.StorageID]storagesystem.Handler)
	var errs []error
	for _, car := range cars {
		if car.StorageID == nil {
			continue
		}
		handler, ok := storageHandlers[*car.StorageID]
		if !ok {
			var err error
			handler, err = storagesystem.NewRCloneHandler(ctx, *car.Storage)
			if err != nil {
				errs = append(errs, errors.Wrapf(err, "Unable to create rclone handler"))
				continue
			}
			storageHandlers[*car.StorageID] = handler
		}
		entry, err := handler.Check(ctx, car.StoragePath)
		if err != nil {
			errs = append(errs, errors.Wrapf(err, "Unable to check file %s", car.StoragePath))
			continue
		}
		obj, ok := entry.(fs.Object)
		if !ok {
			errs = append(errs, errors.Wrapf(err, "%s is not an object", car.StoragePath))
			continue
		}
		err = handler.Remove(ctx, obj)
		if err != nil {
			errs = append(errs, errors.Wrapf(err, "Failed to delete %s", car.StoragePath))
		}
	}

	if len(errs) > 0 {
		return util.AggregateError{Errors: errs}
	}

	return nil
}

// @ID RemovePreparation
// @Summary Remove a preparation
// @Tags Preparation
// @Param name path string true "Preparation ID or name"
// @Param request body RemoveRequest true "Remove Request"
// @Accept json
// @Produce json
// @Success 204
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation/{name} [delete]
func _() {}
