package dataprep

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-log/v2"
	"github.com/rclone/rclone/fs"
	"gorm.io/gorm"
)

var logger = log.Logger("dataprep")

// DeletePieceRequest contains options for deleting a piece from a preparation.
type DeletePieceRequest struct {
	DeleteCar bool `json:"deleteCar"` // Delete the physical CAR file from storage (default: true)
	Force     bool `json:"force"`     // Delete even if deals reference this piece
}

// DeletePieceHandler deletes a piece (CAR) from a preparation.
//
// This function handles deletion of both data pieces and DAG pieces, with appropriate
// cleanup for each type:
//   - Data pieces: resets file_ranges.job_id to NULL so ranges can be re-packed
//   - DAG pieces: resets directories.exported to false so DAG can be re-generated
//
// For non-inline preparations with physical CAR files stored in output storage,
// the CAR file can optionally be deleted from storage (default behavior).
//
// Parameters:
//   - ctx: The context for database transactions and other operations.
//   - db: A pointer to the gorm.DB instance representing the database connection.
//   - prepName: The ID or name for the desired Preparation record.
//   - pieceCIDStr: The piece CID (CommP) identifying the piece to delete.
//   - request: Options controlling deletion behavior.
//
// Returns:
//   - An error if the piece doesn't exist, belongs to a different preparation,
//     has active deals (without Force flag), or storage deletion fails.
func (DefaultHandler) DeletePieceHandler(
	ctx context.Context,
	db *gorm.DB,
	prepName string,
	pieceCIDStr string,
	request DeletePieceRequest,
) error {
	db = db.WithContext(ctx)

	// 1. Parse and validate piece CID
	pieceCID, err := cid.Parse(pieceCIDStr)
	if err != nil {
		return errors.Join(handlererror.ErrInvalidParameter, errors.Wrapf(err, "invalid piece CID %s", pieceCIDStr))
	}
	if pieceCID.Type() != cid.FilCommitmentUnsealed {
		return errors.Wrap(handlererror.ErrInvalidParameter, "piece CID must be commp")
	}

	// 2. Find preparation
	var preparation model.Preparation
	err = preparation.FindByIDOrName(db, prepName)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrapf(handlererror.ErrNotFound, "preparation '%s' does not exist", prepName)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	// 3. Find car by piece_cid AND preparation_id
	var car model.Car
	err = db.Preload("Storage").
		Where("piece_cid = ? AND preparation_id = ?", model.CID(pieceCID), preparation.ID).
		First(&car).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrapf(handlererror.ErrNotFound, "piece '%s' not found in preparation '%s'", pieceCIDStr, prepName)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	// 4. Check for active deals (block unless Force)
	var dealCount int64
	err = db.Model(&model.Deal{}).Where("piece_cid = ?", car.PieceCID).Count(&dealCount).Error
	if err != nil {
		return errors.WithStack(err)
	}
	if dealCount > 0 && !request.Force {
		return errors.Wrapf(handlererror.ErrInvalidParameter,
			"piece has %d deals; use --force to delete anyway", dealCount)
	}

	// 5. Handle piece type-specific cleanup
	if car.PieceType == model.DataPiece {
		// For data pieces: reset file_ranges.job_id to allow re-packing
		if car.JobID != nil {
			err = db.Model(&model.FileRange{}).
				Where("job_id = ?", car.JobID).
				Update("job_id", nil).Error
			if err != nil {
				return errors.Wrap(err, "failed to reset file ranges")
			}
			logger.Infow("reset file ranges for re-packing", "job_id", *car.JobID)
		}
	} else if car.PieceType == model.DagPiece {
		// For DAG pieces: reset directories.exported to allow re-generation
		if car.AttachmentID != nil {
			err = db.Model(&model.Directory{}).
				Where("attachment_id = ?", car.AttachmentID).
				Update("exported", false).Error
			if err != nil {
				return errors.Wrap(err, "failed to reset directory export flags")
			}
			logger.Infow("reset directories for DAG re-generation", "attachment_id", *car.AttachmentID)
		}
	}

	// 6. Delete physical CAR file from storage (if requested and applicable)
	if request.DeleteCar && car.StorageID != nil && car.StoragePath != "" {
		handler, err := storagesystem.NewRCloneHandler(ctx, *car.Storage)
		if err != nil {
			return errors.Wrap(err, "failed to connect to storage")
		}

		entry, err := handler.Check(ctx, car.StoragePath)
		if err != nil {
			// File might already be deleted - warn but continue with DB cleanup
			logger.Warnw("CAR file not found in storage, continuing with DB cleanup",
				"path", car.StoragePath, "storage_id", *car.StorageID, "err", err)
		} else {
			obj, ok := entry.(fs.Object)
			if !ok {
				return errors.Errorf("%s is not a file object", car.StoragePath)
			}
			err = handler.Remove(ctx, obj)
			if err != nil {
				return errors.Wrapf(err, "failed to delete CAR file %s", car.StoragePath)
			}
			logger.Infow("deleted CAR file from storage", "path", car.StoragePath, "storage_id", *car.StorageID)
		}
	}

	// 7. Delete car_blocks and car record in a transaction
	err = database.DoRetry(ctx, func() error {
		return db.Transaction(func(tx *gorm.DB) error {
			// Delete car_blocks first
			result := tx.Where("car_id = ?", car.ID).Delete(&model.CarBlock{})
			if result.Error != nil {
				return errors.Wrap(result.Error, "failed to delete car blocks")
			}
			logger.Infow("deleted car blocks", "car_id", car.ID, "count", result.RowsAffected)

			// Delete the car record
			err := tx.Delete(&car).Error
			if err != nil {
				return errors.Wrap(err, "failed to delete car record")
			}
			return nil
		})
	})
	if err != nil {
		return errors.WithStack(err)
	}

	logger.Infow("successfully deleted piece",
		"piece_cid", pieceCIDStr,
		"preparation", prepName,
		"piece_type", car.PieceType,
		"deals_existed", dealCount)

	return nil
}

// @ID DeletePiece
// @Summary Delete a piece from a preparation
// @Description Deletes a piece (CAR) and its associated records. For data pieces, resets file ranges
// @Description to allow re-packing. For DAG pieces, resets directory export flags for re-generation.
// @Tags Piece
// @Accept json
// @Produce json
// @Param id path string true "Preparation ID or name"
// @Param piece_cid path string true "Piece CID"
// @Param request body DeletePieceRequest true "Delete options"
// @Success 204 "No Content"
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation/{id}/piece/{piece_cid} [delete]
func _() {}
