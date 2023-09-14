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
	PieceCID  string `binding:"required" json:"pieceCid"`      // CID of the piece
	PieceSize string `binding:"required" json:"pieceSize"`     // Size of the piece
	FilePath  string `json:"filePath"    swaggerignore:"true"` // Path to the CAR file, used to determine the size of the file and root CID
	RootCID   string `json:"rootCid"`                          // Root CID of the CAR file, used to populate the label field of storage deal
	FileSize  int64  `json:"fileSize"`                         // File size of the CAR file, this is required for boost online deal
}

type PieceList struct {
	AttachmentID    *model.SourceAttachmentID `json:"attachmentId"`
	SourceStorageID *model.StorageID          `json:"storageId"`
	SourceStorage   *model.Storage            `json:"source"       table:"expand"`
	Pieces          []model.Car               `json:"pieces"       table:"expand"`
}

// ListPiecesHandler retrieves the list of pieces associated with a particular preparation and its source attachments.
//
// This function retrieves the SourceAttachment associated with a given preparation ID. For each source attachment,
// the associated pieces (represented by the Car model) are fetched and grouped. If there are pieces that are not
// associated with any source attachment but are linked to the preparation, they are also fetched.
//
// Parameters:
//   - ctx: The context for database transactions and other operations.
//   - db: A pointer to the gorm.DB instance representing the database connection.
//   - id: The ID or name for the desired Preparation record.
//
// Returns:
//   - A slice of PieceList, each representing a source attachment and its associated pieces.
//   - An error, if any occurred during the operation.
func (DefaultHandler) ListPiecesHandler(
	ctx context.Context,
	db *gorm.DB,
	id string,
) ([]PieceList, error) {
	db = db.WithContext(ctx)
	var preparation model.Preparation
	err := preparation.FindByIDOrName(db, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "preparation '%s' does not exist", id)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var sourceAttachments []model.SourceAttachment
	err = db.Preload("Storage").Where("preparation_id = ?", preparation.ID).Find(&sourceAttachments).Error
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
			AttachmentID:    ptr.Of(sourceAttachment.ID),
			SourceStorageID: ptr.Of(sourceAttachment.StorageID),
			SourceStorage:   sourceAttachment.Storage,
			Pieces:          cars,
		})
	}

	var cars []model.Car
	err = db.Where("attachment_id IS NULL AND preparation_id = ?", preparation.ID).Find(&cars).Error
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

// @ID ListPieces
// @Summary List all prepared pieces for a preparation
// @Tags Piece
// @Accept json
// @Produce json
// @Param id path string true "Preparation ID or name"
// @Success 200 {array} PieceList
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation/{id}/piece [get]
func _() {}

// @ID AddPiece
// @Summary Add a piece to a preparation
// @Tags Piece
// @Accept json
// @Produce json
// @Param id path string true "Preparation ID or name"
// @Param request body AddPieceRequest true "Piece information"
// @Success 200 {object} model.Car
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation/{id}/piece [post]
func _() {}

// AddPieceHandler adds a new piece (represented by the Car model) to a given preparation.
//
// This function fetches a preparation based on the provided ID. It then parses and validates the
// provided piece CID and size from the request. If a root CID is provided in the request, it is
// parsed; if a file path is provided instead, the root CID is extracted from the file.
//
// Once the necessary information is extracted and validated, a new piece (Car model) is created in
// the database associated with the given preparation.
//
// Parameters:
//   - ctx: The context for database transactions and other operations.
//   - db: A pointer to the gorm.DB instance representing the database connection.
//   - id: The ID or name for the desired Preparation record.
//   - request: A struct of AddPieceRequest which contains the information for the piece to be added.
//
// Returns:
//   - A pointer to the newly created Car model.
//   - An error, if any occurred during the operation.
func (DefaultHandler) AddPieceHandler(
	ctx context.Context,
	db *gorm.DB,
	id string,
	request AddPieceRequest,
) (*model.Car, error) {
	db = db.WithContext(ctx)
	var preparation model.Preparation
	err := preparation.FindByIDOrName(db, id)
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
	rootCID := packutil.EmptyFileCid
	fileSize := request.FileSize
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

	if fileSize == 0 && request.FilePath != "" {
		stat, err := os.Stat(request.FilePath)
		if err != nil {
			return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "failed to stat file %s", request.FilePath)
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
