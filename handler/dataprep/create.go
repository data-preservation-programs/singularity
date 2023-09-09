package dataprep

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/dustin/go-humanize"
	"gorm.io/gorm"
)

type CreateRequest struct {
	Name              string   `json:"name"           validate:"required"`                          // Name of the preparation
	SourceStorages    []string `json:"sourceStorages" validate:"required"`                          // Name of Source storage systems to be used for the source
	OutputStorages    []string `json:"outputStorages" validate:"optional"`                          // Name of Output storage systems to be used for the output
	MaxSizeStr        string   `default:"31.5GiB"     json:"maxSize"           validate:"required"` // Maximum size of the CAR files to be created
	PieceSizeStr      string   `default:""            json:"pieceSize"         validate:"optional"` // Target piece size of the CAR files used for piece commitment calculation
	DeleteAfterExport bool     `default:"false"       json:"deleteAfterExport" validate:"optional"` // Whether to delete the source files after export
}

// ValidateCreateRequest processes and validates the creation request parameters.
// The function checks the validity of the input parameters such as maxSize, pieceSize, and
// the existence of source and output storages. The function also ensures that provided
// parameters meet certain criteria, like the pieceSize being a power of two, and maxSize
// allowing for padding. The encryption and storages compatibility is also validated.
//
// Parameters:
// - ctx: The context for database transactions and other operations.
// - db: A pointer to the gorm.DB instance representing the database connection.
// - request: The CreateRequest structure containing the parameters for the creation request.
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

	var sources []model.Storage
	for _, name := range request.SourceStorages {
		var source model.Storage
		err = db.Where("name = ?", name).First(&source).Error
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
		err = db.Where("name = ?", name).First(&output).Error
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

	return &model.Preparation{
		MaxSize:           int64(maxSize),
		PieceSize:         int64(pieceSize),
		SourceStorages:    sources,
		OutputStorages:    outputs,
		DeleteAfterExport: request.DeleteAfterExport,
		Name:              request.Name,
	}, nil
}

// CreatePreparationHandler handles the creation of a new Preparation record based on the provided request.
//
// The function first validates the given request payload. If the validation is successful,
// it will create a new Preparation record in the database. Additionally, for each source
// attachment associated with the created Preparation, it will create a new Directory record.
// All database operations are executed with retries to handle potential transient errors.
//
// Parameters:
//   - ctx: The context for managing timeouts and cancellation.
//   - request: A handler request that wraps the CreateRequest payload.
//   - dep: Contains the handler's dependencies, such as the gorm.DB instance.
//
// Returns:
//   - The created model.Preparation instance if the operation is successful.
//   - An error if any issues occur during the operation, such as validation or database errors.
func (DefaultHandler) CreatePreparationHandler(
	ctx context.Context,
	request handler.Request[CreateRequest],
	dep handler.Dependency,
) (*model.Preparation, error) {
	db := dep.DB.WithContext(ctx)
	preparation, err := ValidateCreateRequest(ctx, db, request.Payload)
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
