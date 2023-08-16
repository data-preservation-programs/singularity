package inspect

import (
	"strconv"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func GetSourcePackingManifestDetailHandler(
	db *gorm.DB,
	id string,
) (*model.PackingManifest, error) {
	return getSourcePackingManifestDetailHandler(db, id)
}

// @Summary Get detail of a specific packing manifest
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Packing manifest ID"
// @Success 200 {object} model.PackingManifest
// @Failure 500 {object} api.HTTPError
// @Router /packingmanifest/{id} [get]
func getSourcePackingManifestDetailHandler(
	db *gorm.DB,
	id string,
) (*model.PackingManifest, error) {
	packingManifestID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("invalid packing manifest id")
	}
	var packingManifest model.PackingManifest
	err = db.Preload("Cars").Preload("FileRanges").Where("id = ?", packingManifestID).First(&packingManifest).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewInvalidParameterErr("packing manifest not found")
	}
	if err != nil {
		return nil, err
	}

	fileMap := make(map[uint64]*model.File)
	for _, part := range packingManifest.FileRanges {
		fileMap[part.FileID] = nil
	}

	fileIDChunks := util.ChunkMapKeys(fileMap, util.BatchSize)
	for _, fileIDChunk := range fileIDChunks {
		var files []model.File
		err = db.Where("id IN ?", fileIDChunk).Find(&files).Error
		if err != nil {
			return nil, err
		}
		for i, file := range files {
			fileMap[file.ID] = &files[i]
		}
	}

	for i, part := range packingManifest.FileRanges {
		file, ok := fileMap[part.FileID]
		if ok {
			packingManifest.FileRanges[i].File = file
		}
	}

	return &packingManifest, nil
}
