package inspect

import (
	"strconv"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func GetSourceChunkDetailHandler(
	db *gorm.DB,
	id string,
) (*model.Chunk, error) {
	return getSourceChunkDetailHandler(db, id)
}

// @Summary Get detail of a specific chunk
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Chunk ID"
// @Success 200 {object} model.Chunk
// @Failure 500 {object} api.HTTPError
// @Router /chunk/{id} [get]
func getSourceChunkDetailHandler(
	db *gorm.DB,
	id string,
) (*model.Chunk, error) {
	chunkID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("invalid chunk id")
	}
	var chunk model.Chunk
	err = db.Preload("Cars").Preload("FileRanges").Where("id = ?", chunkID).First(&chunk).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewInvalidParameterErr("chunk not found")
	}
	if err != nil {
		return nil, err
	}

	fileMap := make(map[uint64]*model.File)
	for _, part := range chunk.FileRanges {
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

	for i, part := range chunk.FileRanges {
		file, ok := fileMap[part.FileID]
		if ok {
			chunk.FileRanges[i].File = file
		}
	}

	return &chunk, nil
}
