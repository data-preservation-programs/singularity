package dataset

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/dustin/go-humanize"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

type CreateRequest struct {
	Name                 string   `json:"name"                 validate:"required"`                     // Name must be a unique identifier for a dataset
	MaxSizeStr           string   `default:"31.5GiB"           json:"maxSize"      validate:"required"` // Maximum size of the CAR files to be created
	PieceSizeStr         string   `default:""                  json:"pieceSize"    validate:"optional"` // Target piece size of the CAR files used for piece commitment calculation
	OutputDirs           []string `json:"outputDirs"           validate:"optional"`                     // Output directory for CAR files. Do not set if using inline preparation
	EncryptionRecipients []string `json:"encryptionRecipients" validate:"optional"`                     // Public key of the encryption recipient
}

func parseCreateRequest(request CreateRequest) (*model.Preparation, error) {
	maxSize, err := humanize.ParseBytes(request.MaxSizeStr)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("invalid value for max-size: " + err.Error())
	}

	pieceSize := util.NextPowerOfTwo(maxSize)
	if request.PieceSizeStr != "" {
		pieceSize, err = humanize.ParseBytes(request.PieceSizeStr)
		if err != nil {
			return nil, handler.NewInvalidParameterErr("invalid value for piece-size: " + err.Error())
		}

		if pieceSize != util.NextPowerOfTwo(pieceSize) {
			return nil, handler.NewInvalidParameterErr("piece size must be a power of two")
		}
	}

	if pieceSize > 1<<36 {
		return nil, handler.NewInvalidParameterErr("piece size cannot be larger than 64 GiB")
	}

	if maxSize*128/127 >= pieceSize {
		return nil, handler.NewInvalidParameterErr("max size needs to be reduced to leave space for padding")
	}

	outDirs := make([]string, len(request.OutputDirs))
	if len(request.OutputDirs) > 1 {
		return nil, handler.NewInvalidParameterErr("multiple output directories will not supported in the future hence no longer allowed")
	}

	for i, outputDir := range request.OutputDirs {
		info, err := os.Stat(outputDir)
		if err != nil || !info.IsDir() {
			return nil, handler.NewInvalidParameterErr("output directory does not exist: " + outputDir)
		}
		abs, err := filepath.Abs(outputDir)
		if err != nil {
			return nil, handler.NewInvalidParameterErr("could not get absolute path for output directory: " + err.Error())
		}
		outDirs[i] = abs
	}

	if len(request.EncryptionRecipients) > 0 && len(request.OutputDirs) == 0 {
		return nil, handler.NewInvalidParameterErr(
			"encryption is not compatible with inline preparation and " +
				"requires at least one output directory",
		)
	}

	return &model.Preparation{
		Name:                 request.Name,
		MaxSize:              int64(maxSize),
		PieceSize:            int64(pieceSize),
		OutputDirs:           outDirs,
		EncryptionRecipients: request.EncryptionRecipients,
	}, nil
}

func CreateHandler(
	ctx context.Context,
	db *gorm.DB,
	request CreateRequest,
) (*model.Preparation, error) {
	return createHandler(ctx, db.WithContext(ctx), request)
}

// @Summary Create a new dataset
// @Tags Preparation
// @Accept json
// @Produce json
// @Description The dataset is a top level object to distinguish different dataset.
// @Param request body CreateRequest true "Request body"
// @Success 200 {object} model.Preparation
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /dataset [post]
func createHandler(
	ctx context.Context,
	db *gorm.DB,
	request CreateRequest,
) (*model.Preparation, error) {
	logger := log.Logger("cli")
	if request.Name == "" {
		return nil, handler.NewInvalidParameterErr("name is required")
	}

	dataset, err := parseCreateRequest(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err2 := database.DoRetry(ctx, func() error { return db.Create(dataset).Error })
	if errors.Is(err2, gorm.ErrDuplicatedKey) || (err2 != nil && strings.Contains(err2.Error(), "constraint failed")) {
		return nil, handler.NewDuplicateRecordError("dataset with this name already exists")
	}

	if err2 != nil {
		return nil, err2
	}

	logger.Infof("Preparation created with ID: %d", dataset.ID)
	return dataset, nil
}
