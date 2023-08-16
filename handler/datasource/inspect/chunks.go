package inspect

import (
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type GetSourcePackingManifestsRequest struct {
	State model.WorkState `json:"state"`
}

func GetSourcePackingManifestsHandler(
	db *gorm.DB,
	id uint32,
	request GetSourcePackingManifestsRequest,
) ([]model.PackingManifest, error) {
	return getSourcePackingManifestsHandler(db, id, request)
}

// @Summary Get all packing manifest details of a data source
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Param request body GetSourcePackingManifestsRequest true "GetSourcePackingManifestsRequest"
// @Success 200 {array} model.PackingManifest
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /source/{id}/packingmanifests [get]
func getSourcePackingManifestsHandler(
	db *gorm.DB,
	sourceID uint32,
	request GetSourcePackingManifestsRequest,
) ([]model.PackingManifest, error) {
	var source model.Source
	err := db.Where("id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewInvalidParameterErr("source not found")
	}
	if err != nil {
		return nil, err
	}

	var packingManifests []model.PackingManifest
	if request.State == "" {
		err = db.Where("source_id = ?", sourceID).Find(&packingManifests).Error
	} else {
		err = db.Where("source_id = ? AND packing_state = ?", sourceID, request.State).Find(&packingManifests).Error
	}

	if err != nil {
		return nil, err
	}

	return packingManifests, nil
}
