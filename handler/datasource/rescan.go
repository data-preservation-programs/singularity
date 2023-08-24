package datasource

import (
	"context"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// @Summary Trigger a rescan of a data source
// @Tags Data Source
// @Produce json
// @Param id path string true "Source ID"
// @Success 200 {object} model.Source
// @Failure 500 {object} api.HTTPError
// @Router /source/{id}/rescan [post]
func rescanSourceHandler(
	ctx context.Context,
	db *gorm.DB,
	id string,
) (*model.Source, error) {
	sourceID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handlererror.NewInvalidParameterErr("invalid source id")
	}
	var source model.Source
	err = db.Transaction(func(db *gorm.DB) error {
		err = db.Where("id = ?", sourceID).First(&source).Error
		if err != nil {
			return errors.WithStack(err)
		}
		if source.ScanningState == model.Error || source.ScanningState == model.Complete {
			return database.DoRetry(ctx, func() error {
				return db.Model(&source).Updates(map[string]any{
					"scanning_state":         model.Ready,
					"last_scanned_timestamp": 0,
					"error_message":          "",
				}).Error
			})
		}
		return nil
	})

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handlererror.NewInvalidParameterErr("source not found")
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &source, nil
}

func RescanSourceHandler(
	ctx context.Context,
	db *gorm.DB,
	id string,
) (*model.Source, error) {
	return rescanSourceHandler(ctx, db.WithContext(ctx), id)
}
