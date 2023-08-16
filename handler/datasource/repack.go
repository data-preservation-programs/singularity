package datasource

import (
	"strconv"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// @Summary Trigger a repack of a packing manifest or all errored packing manifests of a data source
// @Tags Data Source
// @Produce json
// @Param id path string true "Source ID"
// @Param request body RepackRequest true "Request body"
// @Success 200 {array} model.PackingManifest
// @Failure 500 {object} api.HTTPError
// @Router /source/{id}/repack [post]
func repackHandler(
	db *gorm.DB,
	id string,
	request RepackRequest,
) ([]model.PackingManifest, error) {
	if id == "" && request.PackingManifestID == nil {
		return nil, handler.NewInvalidParameterErr("either source id or packing manifest id must be provided")
	}

	var sourceID int
	var err error
	if id != "" {
		sourceID, err = strconv.Atoi(id)
		if err != nil {
			return nil, handler.NewInvalidParameterErr("invalid source id")
		}
	}

	if request.PackingManifestID != nil {
		packingManifestID := *request.PackingManifestID
		var packingManifest model.PackingManifest
		statement := db.Where("id = ?", packingManifestID)
		if sourceID != 0 {
			statement = statement.Where("source_id = ?", sourceID)
		}
		err = statement.First(&packingManifest).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, handler.NewInvalidParameterErr("packing manifest not found")
		}
		if err != nil {
			return nil, err
		}
		if packingManifest.PackingState == model.Error || packingManifest.PackingState == model.Complete {
			err = database.DoRetry(func() error {
				return db.Model(&packingManifest).Updates(map[string]any{
					"packing_state": model.Ready,
					"error_message": "",
				}).Error
			})
			if err != nil {
				return nil, err
			}
		} else {
			return nil, handler.NewInvalidParameterErr("packing manifest is not in error or complete state")
		}
		return []model.PackingManifest{packingManifest}, nil
	}

	var packingManifests []model.PackingManifest
	err = db.Transaction(func(db *gorm.DB) error {
		err = db.Where("source_id = ? and packing_state = ?", sourceID, model.Error).Find(&packingManifests).Error
		if err != nil {
			return err
		}
		err = db.Model(&model.PackingManifest{}).Where("source_id = ? and packing_state = ?", sourceID, model.Error).Updates(map[string]any{
			"packing_state": model.Ready,
			"error_message": "",
		}).Error
		if err != nil {
			return err
		}
		for i := range packingManifests {
			packingManifests[i].PackingState = model.Ready
			packingManifests[i].ErrorMessage = ""
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return packingManifests, nil
}

type RepackRequest struct {
	PackingManifestID *uint64 `json:"packingManifestId"`
}

func RepackHandler(
	db *gorm.DB,
	id string,
	request RepackRequest,
) ([]model.PackingManifest, error) {
	return repackHandler(db, id, request)
}
