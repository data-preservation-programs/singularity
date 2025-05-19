// Final cleaned up version of piece.go

package dataprep

import (
	"bufio"
	"context"
	"os"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack/packutil"
	"github.com/gotidy/ptr"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-car"
	"gorm.io/gorm"
)

type AddPieceRequest struct {
	PieceCID  string `binding:"required" json:"pieceCid"`
	PieceSize string `binding:"required" json:"pieceSize"`
	FilePath  string `json:"filePath" swaggerignore:"true"`
	RootCID   string `json:"rootCid"`
	FileSize  int64  `json:"fileSize"`
}

type PieceList struct {
	AttachmentID    *model.SourceAttachmentID `json:"attachmentId"`
	SourceStorageID *model.StorageID          `json:"storageId"`
	SourceStorage   *model.Storage            `json:"source" table:"expand"`
	Pieces          []model.Car               `json:"pieces" table:"expand"`
}

func batchCountFilesForPieces(ctx context.Context, db *gorm.DB, preparationID uint32) (map[string]int, error) {
	type Result struct {
		PieceCID   []byte
		NumOfFiles int
	}

	var results []Result
	err := db.WithContext(ctx).
		Table("cars AS c").
		Select("c.piece_cid, COUNT(DISTINCT cb.file_id) AS num_of_files").
		Joins("JOIN car_blocks cb ON cb.car_id = c.id").
		Where("c.preparation_id = ? AND cb.file_id IS NOT NULL", preparationID).
		Group("c.piece_cid").
		Scan(&results).
		Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	fileCounts := make(map[string]int)
	for _, r := range results {
		cidObj, err := cid.Cast(r.PieceCID)
		if err != nil {
			continue
		}
		fileCounts[cidObj.String()] = r.NumOfFiles
	}
	return fileCounts, nil
}

func (DefaultHandler) ListPiecesHandler(ctx context.Context, db *gorm.DB, id string) ([]PieceList, error) {
	db = db.WithContext(ctx)
	var preparation model.Preparation
	err := preparation.FindByIDOrName(db, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "preparation '%s' does not exist", id)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	fileCounts, err := batchCountFilesForPieces(ctx, db, uint32(preparation.ID))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var sourceAttachments []model.SourceAttachment
	err = db.Preload("Storage").Where("preparation_id = ?", preparation.ID).Find(&sourceAttachments).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var pieceLists []PieceList
	for _, sa := range sourceAttachments {
		var cars []model.Car
		err = db.Where("attachment_id = ?", sa.ID).Find(&cars).Error
		if err != nil {
			return nil, errors.WithStack(err)
		}
		for i := range cars {
			cars[i].NumOfFiles = int64(fileCounts[cars[i].PieceCID.String()])
		}
		pieceLists = append(pieceLists, PieceList{
			AttachmentID:    ptr.Of(sa.ID),
			SourceStorageID: ptr.Of(sa.StorageID),
			SourceStorage:   sa.Storage,
			Pieces:          cars,
		})
	}

	var cars []model.Car
	err = db.Where("attachment_id IS NULL AND preparation_id = ?", preparation.ID).Find(&cars).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if len(cars) > 0 {
		for i := range cars {
			cars[i].NumOfFiles = int64(fileCounts[cars[i].PieceCID.String()])
		}
		pieceLists = append(pieceLists, PieceList{
			Pieces: cars,
		})
	}
	return pieceLists, nil
}

func (DefaultHandler) AddPieceHandler(ctx context.Context, db *gorm.DB, id string, request AddPieceRequest) (*model.Car, error) {
	db = db.WithContext(ctx)
	var preparation model.Preparation
	err := preparation.FindByIDOrName(db, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "preparation '%s' does not exist", id)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	pieceCID, err := cid.Parse(request.PieceCID)
	if err != nil {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid piece CID: %s", request.PieceCID)
	}
	if pieceCID.Type() != cid.FilCommitmentUnsealed {
		return nil, errors.Wrap(handlererror.ErrInvalidParameter, "piece CID must be commp")
	}

	pieceSize, err := strconv.ParseInt(request.PieceSize, 10, 64)
	if err != nil || (pieceSize&(pieceSize-1)) != 0 {
		return nil, errors.Wrap(handlererror.ErrInvalidParameter, "piece size must be a power of 2")
	}

	rootCID := packutil.EmptyFileCid
	if request.RootCID != "" {
		rootCID, err = cid.Parse(request.RootCID)
		if err != nil {
			return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "invalid root CID: %s", request.RootCID)
		}
	} else if request.FilePath != "" {
		file, err := os.Open(request.FilePath)
		if err != nil {
			return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "failed to open file: %s", request.FilePath)
		}
		defer file.Close()
		header, err := car.ReadHeader(bufio.NewReader(file))
		if err != nil || len(header.Roots) != 1 {
			return nil, errors.Wrap(handlererror.ErrInvalidParameter, "invalid CAR file header")
		}
		rootCID = header.Roots[0]
	}

	fileSize := request.FileSize
	if fileSize == 0 && request.FilePath != "" {
		stat, err := os.Stat(request.FilePath)
		if err != nil {
			return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "failed to stat file: %s", request.FilePath)
		}
		fileSize = stat.Size()
	}

	mCar := model.Car{
		PieceCID:      model.CID(pieceCID),
		PieceSize:     pieceSize,
		RootCID:       model.CID(rootCID),
		StoragePath:   request.FilePath,
		PreparationID: preparation.ID,
		FileSize:      fileSize,
	}

	err = database.DoRetry(ctx, func() error { return db.Create(&mCar).Error })
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &mCar, nil
}
