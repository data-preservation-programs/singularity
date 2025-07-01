package dataprep

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/dustin/go-humanize"
	"gorm.io/gorm"
)

type CreateRequest struct {
	Name              string   `binding:"required"    json:"name"`              // Name of the preparation
	SourceStorages    []string `json:"sourceStorages"`                          // Name of Source storage systems to be used for the source
	OutputStorages    []string `json:"outputStorages"`                          // Name of Output storage systems to be used for the output
	MaxSizeStr        string   `default:"31.5GiB"     json:"maxSize"`           // Maximum size of the CAR files to be created
	PieceSizeStr      string   `default:""            json:"pieceSize"`         // Target piece size of the CAR files used for piece commitment calculation
	MinPieceSizeStr   string   `default:"1MiB"        json:"minPieceSize"`      // Minimum piece size for the preparation, applies only to DAG and remainer pieces
	DeleteAfterExport bool     `default:"false"       json:"deleteAfterExport"` // Whether to delete the source files after export
	NoInline          bool     `default:"false"       json:"noInline"`          // Whether to disable inline storage for the preparation. Can save database space but requires at least one output storage.
	NoDag             bool     `default:"false"       json:"noDag"`             // Whether to disable maintaining folder dag structure for the sources. If disabled, DagGen will not be possible and folders will not have an associated CID.
}

// ValidateCreateRequest processes and validates the creation request parameters.
// The function checks the validity of the input parameters such as maxSize, pieceSize, and
// the existence of source and output storages. The function also ensures that provided
// parameters meet certain criteria, like the pieceSize being a power of two, and maxSize
// allowing for padding. The encryption and storages compatibility is also validated.
//
// Parameters:
//   - ctx: The context for database transactions and other operations.
//   - db: A pointer to the gorm.DB instance representing the database connection.
//   - request: The CreateRequest structure containing the parameters for the creation request.
//
// Returns:
//   - A pointer to the validated Preparation model which can be used for subsequent operations.
//   - An error, if any occurred during the validation. This includes errors such as invalid
//     parameter values, storage not found, or incompatibility between encryption and storage options.
//
// Note:
// If certain parameters are not provided in the request, they are computed based on certain
// defaults or constraints, like the pieceSize defaulting to a power of two value.
func ValidateCreateRequest(ctx context.Context, db *gorm.DB, request CreateRequest) (*model.Preparation, error) {
	db = db.WithContext(ctx)
	if request.MaxSizeStr == "" {
		request.MaxSizeStr = "31.5GiB"
	}

	if util.IsAllDigits(request.Name) || request.Name == "" {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "preparation name '%s' cannot be all digits or empty", request.Name)
	}

	maxSize, err := humanize.ParseBytes(request.MaxSizeStr)
	if err != nil {
		return nil, errors.Join(handlererror.ErrInvalidParameter, errors.Wrapf(err, "invalid value for maxSize: %s", request.MaxSizeStr))
	}

	pieceSize := util.NextPowerOfTwo(maxSize)
	if request.PieceSizeStr != "" {
		pieceSize, err = humanize.ParseBytes(request.PieceSizeStr)
		if err != nil {
			return nil, errors.Join(handlererror.ErrInvalidParameter, errors.Wrapf(err, "invalid value for pieceSize: %s", request.PieceSizeStr))
		}

		if pieceSize != util.NextPowerOfTwo(pieceSize) {
			return nil, errors.Wrap(handlererror.ErrInvalidParameter, "pieceSize must be a power of two")
		}
	}

	if pieceSize > 1<<36 {
		return nil, errors.Wrap(handlererror.ErrInvalidParameter, "pieceSize cannot be larger than 64 GiB")
	}

	if maxSize*128/127 >= pieceSize {
		return nil, errors.Wrap(handlererror.ErrInvalidParameter, "maxSize needs to be reduced to leave space for padding")
	}

	minPieceSizeStr := request.MinPieceSizeStr
	if minPieceSizeStr == "" {
		minPieceSizeStr = "1MiB"
	}

	minPieceSize, err := humanize.ParseBytes(minPieceSizeStr)
	if err != nil {
		return nil, errors.Join(handlererror.ErrInvalidParameter, errors.Wrapf(err, "invalid value for minPieceSize: %s", minPieceSizeStr))
	}

	if minPieceSize > pieceSize {
		return nil, errors.Wrap(handlererror.ErrInvalidParameter, "minPieceSize cannot be larger than pieceSize")
	}

	if minPieceSize != util.NextPowerOfTwo(minPieceSize) {
		return nil, errors.Wrap(handlererror.ErrInvalidParameter, "minPieceSize must be a power of two")
	}

	var sources []model.Storage
	for _, name := range request.SourceStorages {
		var source model.Storage
		err = source.FindByIDOrName(db, name)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(handlererror.ErrNotFound, "source storage %s does not exist", name)
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}
		sources = append(sources, source)
	}

	var outputs []model.Storage
	for _, name := range request.OutputStorages {
		var output model.Storage
		err = output.FindByIDOrName(db, name)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(handlererror.ErrNotFound, "output storage %s does not exist", name)
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}
		outputs = append(outputs, output)
	}

	if len(outputs) == 0 && request.DeleteAfterExport {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "deleteAfterExport cannot be set without output storages")
	}

	if len(outputs) == 0 && request.NoInline {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "inline preparation cannot be disabled without output storages")
	}

	return &model.Preparation{
		MaxSize:           int64(maxSize),
		PieceSize:         int64(pieceSize),
		MinPieceSize:      int64(minPieceSize),
		SourceStorages:    sources,
		OutputStorages:    outputs,
		DeleteAfterExport: request.DeleteAfterExport,
		Name:              request.Name,
		NoInline:          request.NoInline,
		NoDag:             request.NoDag,
	}, nil
}

// CreatePreparationHandler handles the creation of a new Preparation entity based on the provided
// CreateRequest parameters. Initially, it validates the request parameters and, if valid,
// creates a new Preparation record in the database.
//
// Parameters:
//   - ctx: The context for database transactions and other operations.
//   - db: A pointer to the gorm.DB instance representing the database connection.
//   - request: The CreateRequest structure containing the parameters for the creation request.
//
// Returns:
//   - A pointer to the newly created Preparation model.
//   - An error, if any occurred during the validation or creation process.
//
// Note:
// This function relies on the ValidateCreateRequest function to ensure that the provided
// parameters meet the required criteria before creating a Preparation record.
func (DefaultHandler) CreatePreparationHandler(
	ctx context.Context,
	db *gorm.DB,
	request CreateRequest,
) (*model.Preparation, error) {
	db = db.WithContext(ctx)
	preparation, err := ValidateCreateRequest(ctx, db, request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = database.DoRetry(ctx, func() error {
		err := db.Create(preparation).Error
		if err != nil {
			return errors.WithStack(err)
		}

		var attachments []model.SourceAttachment
		err = db.Where("preparation_id = ?", preparation.ID).Find(&attachments).Error
		if err != nil {
			return errors.WithStack(err)
		}

		for _, attachment := range attachments {
			err = db.Create(&model.Directory{
				AttachmentID: attachment.ID,
			}).Error
			if err != nil {
				return errors.WithStack(err)
			}
		}

		return nil
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return preparation, nil
}

// @ID CreatePreparation
// @Summary Create a new preparation
// @Tags Preparation
// @Accept json
// @Produce json
// @Param request body CreateRequest true "Create Request"
// @Success 200 {object} model.Preparation
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation [post]
func _() {}
