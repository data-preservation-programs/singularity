package dataset

import (
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/handler"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

type AddPieceRequest struct {
	PieceCID  string `json:"pieceCID"`
	PieceSize string `json:"pieceSize"`
	FilePath  uint64 `json:"filePath"`
	FileSize  uint64 `json:"fileSize"`
	RootCID   string `json:"rootCID"`
}

// AddPieceHandler godoc
// @Summary Register a CAR file piece with a dataset
// @Tags Dataset
// @Produce json
// @Accept json
// @Param name path string true "Dataset name"
// @Param request body AddPieceRequest true "Request body"
// @Success 200 {object} model.Car
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /dataset/{name}/piece [post]
func AddPieceHandler(
	db *gorm.DB,
	name string,
	request AddPieceRequest,
) (*model.Source, *handler.Error) {
	logger := log.Logger("cli")
	log.SetAllLoggers(log.LevelInfo)
	if name == "" {
		return nil, handler.NewBadRequestString("dataset name is required")
	}

	dataset, err := database.FindDatasetByName(db, name)
	if err != nil {
		return nil, handler.NewBadRequestString("failed to find dataset: " + err.Error())
	}

	cid.Parse(request.PieceCID)

	err = db.Create(&source).Error
	if err != nil {
		return nil, handler.NewBadRequestString("failed to create source: " + err.Error())
	}
	logger.Infof("created source %d", source.ID)
	return &source, nil
}
