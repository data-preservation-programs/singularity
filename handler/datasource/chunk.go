package datasource

import (
	"strconv"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type PackJobRequest struct {
	FileIDs []uint64 `json:"fileIDs" validation:"required"`
}

func CreatePackJobHandler(
	db *gorm.DB,
	sourceID string,
	request PackJobRequest,
) (*model.PackJob, error) {
	return createPackJobHandler(db, sourceID, request)
}

// @Summary Create a pack job for the specified files
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Param request body PackJobRequest true "Request body"
// @Success 201 {object} model.File
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /source/{id}/packjob [post]
func createPackJobHandler(
	db *gorm.DB,
	sourceID string,
	request PackJobRequest,
) (*model.PackJob, error) {
	sourceIDInt, err := strconv.Atoi(sourceID)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("invalid source id")
	}

	packJob := model.PackJob{
		SourceID:     uint32(sourceIDInt),
		PackingState: model.Ready,
	}

	err = database.DoRetry(func() error {
		return db.Transaction(
			func(db *gorm.DB) error {
				err := db.Create(&packJob).Error
				if err != nil {
					return errors.Wrap(err, "failed to create pack job")
				}
				fileRangeIDChunks := util.ChunkSlice(request.FileIDs, util.BatchSize)
				for _, fileRangeIDChunks := range fileRangeIDChunks {
					err = db.Model(&model.FileRange{}).
						Where("id IN ?", fileRangeIDChunks).Update("packing_manifest_id", packJob.ID).Error
					if err != nil {
						return errors.Wrap(err, "failed to update files")
					}
				}
				return nil
			},
		)
	})
	if err != nil {
		return nil, err
	}

	return &packJob, nil
}
