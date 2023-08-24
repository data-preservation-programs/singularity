package inspect

import (
	"context"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"gorm.io/gorm"
)

func GetSourcePackJobDetailHandler(
	ctx context.Context,
	db *gorm.DB,
	id string,
) (*model.PackJob, error) {
	return getSourcePackJobDetailHandler(db.WithContext(ctx), id)
}

// @Summary Get detail of a specific pack job
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Pack job ID"
// @Success 200 {object} model.PackJob
// @Failure 500 {object} api.HTTPError
// @Router /packjob/{id} [get]
func getSourcePackJobDetailHandler(
	db *gorm.DB,
	id string,
) (*model.PackJob, error) {
	packJobID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handlererror.NewInvalidParameterErr("invalid pack job id")
	}
	var packJob model.PackJob
	err = db.Preload("Cars").Preload("FileRanges").Where("id = ?", packJobID).First(&packJob).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handlererror.NewInvalidParameterErr("pack job not found")
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	fileMap := make(map[uint64]*model.File)
	for _, part := range packJob.FileRanges {
		fileMap[part.FileID] = nil
	}

	fileIDChunks := util.ChunkMapKeys(fileMap, util.BatchSize)
	for _, fileIDChunk := range fileIDChunks {
		var files []model.File
		err = db.Where("id IN ?", fileIDChunk).Find(&files).Error
		if err != nil {
			return nil, errors.WithStack(err)
		}
		for i, file := range files {
			fileMap[file.ID] = &files[i]
		}
	}

	for i, part := range packJob.FileRanges {
		file, ok := fileMap[part.FileID]
		if ok {
			packJob.FileRanges[i].File = file
		}
	}

	return &packJob, nil
}
