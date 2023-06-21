package dataset

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/dustin/go-humanize"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

type UpdateRequest struct {
	MaxSizeStr           *string  `default:"31.5GiB"           json:"maxSize"      validate:"optional"` // Maximum size of the CAR files to be created
	PieceSizeStr         *string  `default:""                  json:"pieceSize"    validate:"optional"` // Target piece size of the CAR files used for piece commitment calculation
	OutputDirs           []string `json:"outputDirs"           validate:"optional"`                     // Output directory for CAR files. Do not set if using inline preparation
	EncryptionRecipients []string `json:"encryptionRecipients" validate:"optional"`                     // Public key of the encryption recipient
	EncryptionScript     *string  `json:"encryptionScript"     validate:"optional"`                     // EncryptionScript command to run for custom encryption
}

// UpdateHandler godoc
// @Summary Update a dataset
// @Tags Dataset
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Param request body UpdateRequest true "Request body"
// @Success 200 {object} model.Dataset
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /dataset/{datasetName} [patch]
func UpdateHandler(
	db *gorm.DB,
	datasetName string,
	request UpdateRequest,
) (*model.Dataset, *handler.Error) {
	logger := log.Logger("cli")
	if datasetName == "" {
		return nil, handler.NewBadRequestString("name is required")
	}

	dataset, err := database.FindDatasetByName(db, datasetName)
	if err != nil {
		return nil, handler.NewBadRequestString("dataset not found")
	}

	err2 := parseUpdateRequest(request, &dataset)
	if err2 != nil {
		return nil, err2
	}

	err = database.DoRetry(func() error { return db.Save(&dataset).Error })
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	logger.Infof("Dataset created with ID: %d", dataset.ID)
	return &dataset, nil
}

func parseUpdateRequest(request UpdateRequest, dataset *model.Dataset) *handler.Error {
	if request.MaxSizeStr != nil {
		maxSize, err := humanize.ParseBytes(*request.MaxSizeStr)
		if err != nil {
			return handler.NewBadRequestString("invalid value for max-size: " + err.Error())
		}
		dataset.MaxSize = int64(maxSize)
	}

	pieceSize := util.NextPowerOfTwo(uint64(dataset.MaxSize))
	if request.PieceSizeStr != nil && *request.PieceSizeStr != "" {
		var err error
		pieceSize, err = humanize.ParseBytes(*request.PieceSizeStr)
		if err != nil {
			return handler.NewBadRequestString("invalid value for piece-size: " + err.Error())
		}

		if pieceSize != util.NextPowerOfTwo(pieceSize) {
			return handler.NewBadRequestString("piece size must be a power of two")
		}
	}

	dataset.PieceSize = int64(pieceSize)
	if dataset.PieceSize > 1<<36 {
		return handler.NewBadRequestString("piece size cannot be larger than 64 GiB")
	}

	if dataset.MaxSize*128/127 >= dataset.PieceSize {
		return handler.NewBadRequestString("max size needs to be reduced to leave space for padding")
	}

	if request.OutputDirs != nil {
		outDirs := make([]string, len(request.OutputDirs))
		for i, outputDir := range request.OutputDirs {
			info, err := os.Stat(outputDir)
			if err != nil || !info.IsDir() {
				return handler.NewBadRequestString("output directory does not exist: " + outputDir)
			}
			abs, err := filepath.Abs(outputDir)
			if err != nil {
				return handler.NewBadRequestString("could not get absolute path for output directory: " + err.Error())
			}
			outDirs[i] = abs
		}
		dataset.OutputDirs = outDirs
	}

	if request.EncryptionRecipients != nil {
		dataset.EncryptionRecipients = request.EncryptionRecipients
	}

	if request.EncryptionScript != nil {
		dataset.EncryptionScript = *request.EncryptionScript
	}

	if len(dataset.EncryptionRecipients) > 0 && dataset.EncryptionScript != "" {
		return handler.NewBadRequestString("encryption recipients and script cannot be used together")
	}

	if (len(dataset.EncryptionRecipients) > 0 || dataset.EncryptionScript != "") && len(dataset.OutputDirs) == 0 {
		return handler.NewBadRequestString(
			"encryption is not compatible with inline preparation and " +
				"requires at least one output directory",
		)
	}

	return nil
}
