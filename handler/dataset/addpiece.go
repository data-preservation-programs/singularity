package dataset

import (
	"bufio"
	"context"
	"os"
	"strconv"

	"github.com/ipfs/boxo/util"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-log/v2"
	"github.com/ipld/go-car"
	"gorm.io/gorm"
)

type AddPieceRequest struct {
	PieceCID  string `json:"pieceCid"`  // CID of the piece
	PieceSize string `json:"pieceSize"` // Size of the piece
	FilePath  string `json:"filePath"`  // Path to the CAR file, used to determine the size of the file and root CID
	RootCID   string `json:"rootCid"`   // Root CID of the CAR file, if not provided, will be determined by the CAR file header. Used to populate the label field of storage deal
}

func AddPieceHandler(
	ctx context.Context,
	db *gorm.DB,
	datasetName string,
	request AddPieceRequest,
) (*model.Car, error) {
	return addPieceHandler(ctx, db.WithContext(ctx), datasetName, request)
}

// @Summary Manually register a piece (CAR file) with the dataset for deal making purpose
// @Tags Preparation
// @Produce json
// @Accept json
// @Param datasetName path string true "Preparation name"
// @Param request body AddPieceRequest true "Request body"
// @Success 200 {object} model.Car
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /dataset/{datasetName}/piece [post]
func addPieceHandler(
	ctx context.Context,
	db *gorm.DB,
	datasetName string,
	request AddPieceRequest,
) (*model.Car, error) {
	logger := log.Logger("cli")
	if datasetName == "" {
		return nil, handler.NewInvalidParameterErr("dataset name is required")
	}

	dataset, err := database.FindDatasetByName(db, datasetName)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("failed to find dataset: " + err.Error())
	}

	pieceCID, err := cid.Parse(request.PieceCID)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("invalid piece CID")
	}
	if pieceCID.Type() != cid.FilCommitmentUnsealed {
		return nil, handler.NewInvalidParameterErr("piece CID must be commp")
	}
	pieceSize, err := strconv.ParseInt(request.PieceSize, 10, 64)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("invalid piece size")
	}
	if (pieceSize & (pieceSize - 1)) != 0 {
		return nil, handler.NewInvalidParameterErr("piece size must be a power of 2")
	}
	rootCID := cid.NewCidV1(cid.Raw, util.Hash([]byte(""))).String()
	if request.FilePath != "" {
		file, err := os.Open(request.FilePath)
		if err != nil {
			return nil, handler.NewInvalidParameterErr("failed to open file: " + err.Error())
		}
		defer file.Close()
		header, err := car.ReadHeader(bufio.NewReader(file))
		if err != nil {
			return nil, handler.NewInvalidParameterErr("failed to read CAR header: " + err.Error())
		}
		if len(header.Roots) != 1 {
			logger.Warnf("CAR file has %d roots, expected 1", len(header.Roots))
		}
		if len(header.Roots) > 0 {
			rootCID = header.Roots[0].String()
		}
	}
	if request.RootCID != "" {
		rootCID = request.RootCID
	}

	pCid, err := cid.Decode(request.PieceCID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	rCid, err := cid.Decode(rootCID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	mCar := model.Car{
		PieceCID:    model.CID(pCid),
		PieceSize:   pieceSize,
		RootCID:     model.CID(rCid),
		StoragePath: request.FilePath,
		DatasetID:   dataset.ID,
	}

	err = database.DoRetry(ctx, func() error { return db.Create(&mCar).Error })
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &mCar, nil
}
