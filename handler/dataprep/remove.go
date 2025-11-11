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
	"github.com/rjNemo/underscore"
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

	attachments, err := preparation.SourceAttachments(db)
	if err != nil {
		return errors.WithStack(err)
	}
	attachmentIDs := underscore.Map(attachments, func(attachment model.SourceAttachment) model.SourceAttachmentID { return attachment.ID })
	var activeCount int64
	err = db.Model(&model.Job{}).Where("attachment_id in ? and state = ?", attachmentIDs, model.Processing).Count(&activeCount).Error
	if err != nil {
		return errors.WithStack(err)
	}
	if activeCount > 0 {
		return errors.Wrapf(handlererror.ErrInvalidParameter, "preparation %s has %d active jobs", name, activeCount)
	}

	var cars []model.Car
	if request.RemoveCars {
		err = db.Preload("Storage").Where("preparation_id = ?", preparation.ID).Find(&cars).Error
		if err != nil {
			return errors.WithStack(err)
		}
	}

	err = database.DoRetry(ctx, func() error {
		return db.Transaction(func(tx *gorm.DB) error {
			// Explicitly delete child records to avoid CASCADE deadlocks with concurrent operations.
			// We materialize IDs first to avoid nested subqueries that cause MySQL deadlocks.

			// Step 1: Get all attachment IDs for this preparation
			var attachmentIDs []model.SourceAttachmentID
			err := tx.Table("source_attachments").Select("id").
				Where("preparation_id = ?", preparation.ID).
				Find(&attachmentIDs).Error
			if err != nil {
				return errors.WithStack(err)
			}

			if len(attachmentIDs) == 0 {
				// No attachments, just delete the preparation
				return tx.Select("Wallets", "SourceStorages", "OutputStorages").Delete(&preparation).Error
			}

			// Step 2: Get all car IDs
			var carIDs []model.CarID
			err = tx.Table("cars").Select("id").
				Where("preparation_id = ?", preparation.ID).
				Find(&carIDs).Error
			if err != nil {
				return errors.WithStack(err)
			}

			// Step 3: Get all job IDs
			var jobIDs []model.JobID
			err = tx.Table("jobs").Select("id").
				Where("attachment_id IN ?", attachmentIDs).
				Find(&jobIDs).Error
			if err != nil {
				return errors.WithStack(err)
			}

			// Step 4: Get all file IDs
			var fileIDs []model.FileID
			err = tx.Table("files").Select("id").
				Where("attachment_id IN ?", attachmentIDs).
				Find(&fileIDs).Error
			if err != nil {
				return errors.WithStack(err)
			}

			// Now delete in leaf-to-root order using materialized IDs:

			// 1. Delete car_blocks (leaf node)
			if len(carIDs) > 0 {
				err = tx.Where("car_id IN ?", carIDs).Delete(&model.CarBlock{}).Error
				if err != nil {
					return errors.WithStack(err)
				}
			}

			// 2. Delete cars
			if len(carIDs) > 0 {
				err = tx.Where("id IN ?", carIDs).Delete(&model.Car{}).Error
				if err != nil {
					return errors.WithStack(err)
				}
			}

			// 3. Delete file_ranges (from jobs)
			if len(jobIDs) > 0 {
				err = tx.Where("job_id IN ?", jobIDs).Delete(&model.FileRange{}).Error
				if err != nil {
					return errors.WithStack(err)
				}
			}

			// 4. Delete file_ranges (from files)
			if len(fileIDs) > 0 {
				err = tx.Where("file_id IN ?", fileIDs).Delete(&model.FileRange{}).Error
				if err != nil {
					return errors.WithStack(err)
				}
			}

			// 5. Delete files (before directories to avoid circular cascade)
			if len(fileIDs) > 0 {
				err = tx.Where("id IN ?", fileIDs).Delete(&model.File{}).Error
				if err != nil {
					return errors.WithStack(err)
				}
			}

			// 6. Delete directories
			err = tx.Where("attachment_id IN ?", attachmentIDs).Delete(&model.Directory{}).Error
			if err != nil {
				return errors.WithStack(err)
			}

			// 7. Delete jobs
			if len(jobIDs) > 0 {
				err = tx.Where("id IN ?", jobIDs).Delete(&model.Job{}).Error
				if err != nil {
					return errors.WithStack(err)
				}
			}

			// 8. Now delete the preparation itself, which will cascade to:
			//    - wallet_assignments (many2many, small table)
			//    - source_attachments (now empty, no more cascades)
			//    - output_attachments (many2many, small table)
			// These cascades are safe because we've already deleted all the heavy child tables.
			return tx.Select("Wallets", "SourceStorages", "OutputStorages").Delete(&preparation).Error
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
