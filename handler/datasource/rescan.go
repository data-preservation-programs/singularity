package datasource

import (
	"strconv"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
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
	db *gorm.DB,
	id string,
) (*model.Source, error) {
	sourceID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("invalid source id")
	}
	var source model.Source
	err = db.Transaction(func(db *gorm.DB) error {
		err = db.Where("id = ?", sourceID).First(&source).Error
		if err != nil {
			return err
		}
		if source.ScanningState == model.Error || source.ScanningState == model.Complete {
			return database.DoRetry(func() error {
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
		return nil, handler.NewInvalidParameterErr("source not found")
	}
	if err != nil {
		return nil, err
	}

	return &source, nil
}

func RescanSourceHandler(
	db *gorm.DB,
	id string,
) (*model.Source, error) {
	return rescanSourceHandler(db, id)
}
