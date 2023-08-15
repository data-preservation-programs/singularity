package dataset

import (
	"os"
	"path/filepath"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/dustin/go-humanize"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

type UpdateRequest struct {
	MaxSizeStr           *string  `default:"31.5GiB"           json:"maxSize"      validate:"optional"` // Maximum size of the CAR files to be created
	PieceSizeStr         *string  `default:""                  json:"pieceSize"    validate:"optional"` // Target piece size of the CAR files used for piece commitment calculation
	OutputDirs           []string `json:"outputDirs"           validate:"optional"`                     // Output directory for CAR files. Do not set if using inline preparation
	EncryptionRecipients []string `json:"encryptionRecipients" validate:"optional"`                     // Public key of the encryption recipient
}

func UpdateHandler(
	db *gorm.DB,
	datasetName string,
	request UpdateRequest,
) (*model.Dataset, error) {
	return updateHandler(db, datasetName, request)
}

// @Summary Update a dataset
// @Tags Dataset
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Param request body UpdateRequest true "Request body"
// @Success 200 {object} model.Dataset
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /dataset/{datasetName} [patch]
func updateHandler(
	db *gorm.DB,
	datasetName string,
	request UpdateRequest,
) (*model.Dataset, error) {
	logger := log.Logger("cli")
	if datasetName == "" {
		return nil, handler.NewInvalidParameterErr("name is required")
	}

	dataset, err := database.FindDatasetByName(db, datasetName)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("dataset not found")
	}

	err2 := parseUpdateRequest(request, &dataset)
	if err2 != nil {
		return nil, err2
	}

	err = database.DoRetry(func() error { return db.Save(&dataset).Error })
	if err != nil {
		return nil, err
	}

	logger.Infof("Dataset created with ID: %d", dataset.ID)
	return &dataset, nil
}

func parseUpdateRequest(request UpdateRequest, dataset *model.Dataset) error {
	if request.MaxSizeStr != nil {
		maxSize, err := humanize.ParseBytes(*request.MaxSizeStr)
		if err != nil {
			return handler.NewInvalidParameterErr("invalid value for max-size: " + err.Error())
		}
		dataset.MaxSize = int64(maxSize)
	}

	pieceSize := util.NextPowerOfTwo(uint64(dataset.MaxSize))
	if request.PieceSizeStr != nil && *request.PieceSizeStr != "" {
		var err error
		pieceSize, err = humanize.ParseBytes(*request.PieceSizeStr)
		if err != nil {
			return handler.NewInvalidParameterErr("invalid value for piece-size: " + err.Error())
		}

		if pieceSize != util.NextPowerOfTwo(pieceSize) {
			return handler.NewInvalidParameterErr("piece size must be a power of two")
		}
	}

	dataset.PieceSize = int64(pieceSize)
	if dataset.PieceSize > 1<<36 {
		return handler.NewInvalidParameterErr("piece size cannot be larger than 64 GiB")
	}

	if dataset.MaxSize*128/127 >= dataset.PieceSize {
		return handler.NewInvalidParameterErr("max size needs to be reduced to leave space for padding")
	}

	if len(request.OutputDirs) > 1 {
		return handler.NewInvalidParameterErr("multiple output directories will not supported in the future hence no longer allowed")
	}

	if request.OutputDirs != nil {
		if len(request.OutputDirs) == 1 && request.OutputDirs[0] == "" {
			dataset.OutputDirs = nil
		} else {
			outDirs := make([]string, len(request.OutputDirs))
			for i, outputDir := range request.OutputDirs {
				info, err := os.Stat(outputDir)
				if err != nil || !info.IsDir() {
					return handler.NewInvalidParameterErr("output directory does not exist: " + outputDir)
				}
				abs, err := filepath.Abs(outputDir)
				if err != nil {
					return handler.NewInvalidParameterErr("could not get absolute path for output directory: " + err.Error())
				}
				outDirs[i] = abs
			}
			dataset.OutputDirs = outDirs
		}
	}

	if request.EncryptionRecipients != nil {
		dataset.EncryptionRecipients = request.EncryptionRecipients
	}

	if len(dataset.EncryptionRecipients) > 0 && len(dataset.OutputDirs) == 0 {
		return handler.NewInvalidParameterErr(
			"encryption is not compatible with inline preparation and " +
				"requires at least one output directory",
		)
	}

	return nil
}
