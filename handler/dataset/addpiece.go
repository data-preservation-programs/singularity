package dataset

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-log/v2"
	"github.com/ipld/go-car"
	"gorm.io/gorm"
)

type AddPieceRequest struct {
	PieceCID  string `json:"pieceCID"`
	PieceSize string `json:"pieceSize"`
	FilePath  string `json:"filePath"`
	FileSize  uint64 `json:"fileSize"`
	RootCID   string `json:"rootCID"`
}

// AddPieceHandler godoc
// @Summary Manually register a piece (CAR file) with the dataset for deal making purpose
// @Tags Dataset Piece
// @Produce json
// @Accept json
// @Param datasetName path string true "Dataset name"
// @Param request body AddPieceRequest true "Request body"
// @Success 200 {object} model.Car
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /dataset/{datasetName}/piece [post]
func AddPieceHandler(
	db *gorm.DB,
	datasetName string,
	request AddPieceRequest,
) (*model.Car, *handler.Error) {
	logger := log.Logger("cli")
	log.SetAllLoggers(log.LevelInfo)
	if datasetName == "" {
		return nil, handler.NewBadRequestString("dataset name is required")
	}

	dataset, err := database.FindDatasetByName(db, datasetName)
	if err != nil {
		return nil, handler.NewBadRequestString("failed to find dataset: " + err.Error())
	}

	pieceCID, err := cid.Parse(request.PieceCID)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid piece CID")
	}
	if pieceCID.Type() != cid.FilCommitmentUnsealed {
		return nil, handler.NewBadRequestString("piece CID must be commp")
	}
	pieceSize, err := strconv.ParseInt(request.PieceSize, 10, 64)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid piece size")
	}
	if (pieceSize & (pieceSize - 1)) != 0 {
		return nil, handler.NewBadRequestString("piece size must be a power of 2")
	}
	rootCID := ""
	fileSize := int64(0)
	if request.FilePath != "" {
		file, err := os.Open(request.FilePath)
		if err != nil {
			return nil, handler.NewBadRequestString("failed to open file: " + err.Error())
		}
		defer file.Close()
		header, err := car.ReadHeader(bufio.NewReader(file))
		if err != nil {
			return nil, handler.NewBadRequestString("failed to read CAR header: " + err.Error())
		}
		if len(header.Roots) != 1 {
			logger.Warnf("CAR file has %d roots, expected 1", len(header.Roots))
		}
		if len(header.Roots) > 0 {
			rootCID = header.Roots[0].String()
		}
		resolvedPath, err := filepath.EvalSymlinks(request.FilePath)
		if err != nil {
			return nil, handler.NewBadRequestString("failed to resolve symlinks: " + err.Error())
		}
		stat, err := os.Stat(resolvedPath)
		if err != nil {
			return nil, handler.NewBadRequestString("failed to stat file: " + err.Error())
		}
		fileSize = stat.Size()
		request.FilePath, err = filepath.Abs(request.FilePath)
		if err != nil {
			return nil, handler.NewBadRequestString("failed to get absolute path: " + err.Error())
		}
	}
	if request.FileSize != 0 {
		fileSize = int64(request.FileSize)
	}
	if request.RootCID != "" {
		rootCID = request.RootCID
	}

	if fileSize >= pieceSize {
		return nil, handler.NewBadRequestString("piece size must be larger than file size")
	}

	car := model.Car{
		PieceCID:  model.CID(cid.MustParse(request.PieceCID)),
		PieceSize: pieceSize,
		RootCID:   model.CID(cid.MustParse(rootCID)),
		FileSize:  fileSize,
		FilePath:  request.FilePath,
		DatasetID: dataset.ID,
	}

	err = database.DoRetry(func() error { return db.Create(&car).Error })
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}
	return &car, nil
}
