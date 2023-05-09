package dataset

import (
	"github.com/data-preservation-programs/go-singularity/handler"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/data-preservation-programs/go-singularity/util"
	"github.com/dustin/go-humanize"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

type CreateRequest struct {
	Name         string   `json:"name"`
	MinSizeStr   string   `json:"minSize"`
	MaxSizeStr   string   `json:"maxSize"`
	PieceSizeStr string   `json:"pieceSize"`
	OutputDirs   []string `json:"outputDirs"`
	Recipients   []string `json:"recipients"`
	Script       string   `json:"script"`
}

// CreateHandler godoc
// @Summary Create a new dataset
// @Tags Dataset
// @Accept json
// @Produce json
// @Param request body CreateRequest true "Request body"
// @Success 200 {object} model.Dataset
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /dataset [post]
func CreateHandler(
	db *gorm.DB,
	request CreateRequest,
) (*model.Dataset, *handler.Error) {
	logger := log.Logger("cli")
	log.SetAllLoggers(log.LevelInfo)
	if request.Name == "" {
		return nil, handler.NewBadRequestString("name is required")
	}

	minSize, err := humanize.ParseBytes(request.MinSizeStr)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid value for min-size: " + err.Error())
	}

	maxSize, err := humanize.ParseBytes(request.MaxSizeStr)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid value for max-size: " + err.Error())
	}

	pieceSize := util.NextPowerOfTwo(maxSize)
	if request.PieceSizeStr != "" {
		pieceSize, err = humanize.ParseBytes(request.PieceSizeStr)
		if err != nil {
			return nil, handler.NewBadRequestString("invalid value for piece-size: " + err.Error())
		}

		if pieceSize != util.NextPowerOfTwo(pieceSize) {
			return nil, handler.NewBadRequestString("piece size must be a power of two")
		}
	}

	if pieceSize > 1<<36 {
		return nil, handler.NewBadRequestString("piece size cannot be larger than 64 GiB")
	}

	if maxSize*128/127 >= pieceSize {
		return nil, handler.NewBadRequestString("max size needs to be reduced to leave space for padding")
	}

	outDirs := make([]string, 0, len(request.OutputDirs))
	for i, outputDir := range request.OutputDirs {
		info, err := os.Stat(outputDir)
		if err != nil || !info.IsDir() {
			return nil, handler.NewBadRequestString("output directory does not exist: " + outputDir)
		}
		abs, err := filepath.Abs(outputDir)
		if err != nil {
			return nil, handler.NewBadRequestString("could not get absolute path for output directory: " + err.Error())
		}
		outDirs[i] = abs
	}

	if len(request.Recipients) > 0 && request.Script != "" {
		return nil, handler.NewBadRequestString("encryption recipients and script cannot be used together")
	}

	if (len(request.Recipients) > 0 || request.Script != "") && len(request.OutputDirs) == 0 {
		return nil, handler.NewBadRequestString("encryption is not compatible with inline preparation and " +
			"requires at least one output directory")
	}

	dataset := model.Dataset{
		Name:                 request.Name,
		MinSize:              minSize,
		MaxSize:              maxSize,
		PieceSize:            pieceSize,
		OutputDirs:           request.OutputDirs,
		EncryptionRecipients: request.Recipients,
		EncryptionScript:     request.Script,
	}

	err = db.Create(&dataset).Error
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	logger.Infof("Dataset created with ID: %d", dataset.ID)
	return &dataset, nil
}
