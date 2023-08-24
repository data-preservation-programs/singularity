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
	"github.com/data-preservation-programs/singularity/pack"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-car"
	"gorm.io/gorm"
)

type AddPieceRequest struct {
	PieceCID  string `json:"pieceCid"`  // CID of the piece
	PieceSize string `json:"pieceSize"` // Size of the piece
	FilePath  string `json:"filePath"`  // Path to the CAR file, used to determine the size of the file and root CID
	RootCID   string `json:"rootCid"`   // Root CID of the CAR file, if not provided, will be determined by the CAR file header. Used to populate the label field of storage deal
}

type PieceList struct {
	Source *model.SourceAttachment `json:"source"`
	Pieces []model.Car             `json:"pieces"`
}

func ListPiecesHandler(
	ctx context.Context,
	db *gorm.DB,
	id int32,
) ([]PieceList, error) {
	db = db.WithContext(ctx)
	var sourceAttachments []model.SourceAttachment
	err := db.Preload("Storage").Where("preparation_id = ?", id).Find(&sourceAttachments).Error
	if len(sourceAttachments) == 0 {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "preparation %d not found", id)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var pieceLists []PieceList
	for _, sourceAttachment := range sourceAttachments {
		var cars []model.Car
		err = db.Where("attachment_id = ?", sourceAttachment.ID).Find(&cars).Error
		if err != nil {
			return nil, errors.WithStack(err)
		}
		pieceLists = append(pieceLists, PieceList{
			Source: &sourceAttachment,
			Pieces: cars,
		})
	}

	var cars []model.Car
	err = db.Where("attachment_id IS NULL AND preparation_id = ?", id).Find(&cars).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if len(cars) > 0 {
		pieceLists = append(pieceLists, PieceList{
			Pieces: cars,
		})
	}

	return pieceLists, nil
}

func AddPieceHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint32,
	request AddPieceRequest,
) (*model.Car, error) {
	db = db.WithContext(ctx)
	var preparation model.Preparation
	err := db.Where("id = ?", id).First(&preparation).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "preparation %d not found", id)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	pieceCID, err := cid.Parse(request.PieceCID)
	if err != nil {
		return nil, errors.Join(handlererror.ErrInvalidParameter, errors.Wrapf(err, "invalid piece CID %s", request.PieceCID))
	}
	if pieceCID.Type() != cid.FilCommitmentUnsealed {
		return nil, errors.Wrap(handlererror.ErrInvalidParameter, "piece CID must be commp")
	}
	pieceSize, err := strconv.ParseInt(request.PieceSize, 10, 64)
	if err != nil {
		return nil, errors.Join(handlererror.ErrInvalidParameter, errors.Wrapf(err, "invalid piece size %s", request.PieceSize))
	}
	if (pieceSize & (pieceSize - 1)) != 0 {
		return nil, errors.Wrap(handlererror.ErrInvalidParameter, "piece size must be a power of 2")
	}
	rootCID := pack.EmptyFileCid
	if request.RootCID != "" {
		rootCID, err = cid.Parse(request.RootCID)
		if err != nil {
			return nil, errors.Join(handlererror.ErrInvalidParameter, errors.Wrapf(err, "invalid root CID %s", request.RootCID))
		}
	} else if request.FilePath != "" {
		file, err := os.Open(request.FilePath)
		if err != nil {
			return nil, errors.Join(handlererror.ErrInvalidParameter, errors.Wrapf(err, "failed to open file %s", request.FilePath))
		}
		defer file.Close()
		header, err := car.ReadHeader(bufio.NewReader(file))
		if err != nil {
			return nil, errors.Join(handlererror.ErrInvalidParameter, errors.Wrapf(err, "failed to read CAR header from file %s", request.FilePath))
		}
		if len(header.Roots) != 1 {
			return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "CAR file has %d roots, expected 1", len(header.Roots))
		}
		rootCID = header.Roots[0]
	}

	mCar := model.Car{
		PieceCID:      model.CID(pieceCID),
		PieceSize:     pieceSize,
		RootCID:       model.CID(rootCID),
		StoragePath:   request.FilePath,
		PreparationID: preparation.ID,
	}

	err = database.DoRetry(ctx, func() error { return db.Create(&mCar).Error })
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &mCar, nil
}
