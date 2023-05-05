package dataset

import (
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/data-preservation-programs/go-singularity/util"
	"github.com/dustin/go-humanize"
	"github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

type CreateRequest struct {
	Name         string
	MinSizeStr   string
	MaxSizeStr   string
	PieceSizeStr string
	OutputDirs   []string
	Recipients   []string
	Script       string
}

func CreateHandler(
	db *gorm.DB,
	request CreateRequest,
) (*model.Dataset, error) {
	logger := log.Logger("cli")
	log.SetAllLoggers(log.LevelInfo)
	if request.Name == "" {
		return nil, cli.Exit("dataset name is required", 1)
	}

	minSize, err := humanize.ParseBytes(request.MinSizeStr)
	if err != nil {
		return nil, cli.Exit("invalid value for min-size: "+err.Error(), 1)
	}

	maxSize, err := humanize.ParseBytes(request.MaxSizeStr)
	if err != nil {
		return nil, cli.Exit("invalid value for max-size: "+err.Error(), 1)
	}

	pieceSize := util.NextPowerOfTwo(maxSize)
	if request.PieceSizeStr != "" {
		pieceSize, err = humanize.ParseBytes(request.PieceSizeStr)
		if err != nil {
			return nil, cli.Exit("invalid value for piece-size: "+err.Error(), 1)
		}

		if pieceSize != util.NextPowerOfTwo(pieceSize) {
			return nil, cli.Exit("piece size must be a power of two", 1)
		}
	}

	if pieceSize > 1<<36 {
		return nil, cli.Exit("piece size cannot be larger than 64 GiB", 1)
	}

	if maxSize*128/127 >= pieceSize {
		return nil, cli.Exit("max size needs to be reduced to leave space for padding", 1)
	}

	outDirs := make([]string, 0, len(request.OutputDirs))
	for i, outputDir := range request.OutputDirs {
		info, err := os.Stat(outputDir)
		if err != nil || !info.IsDir() {
			return nil, cli.Exit("output directory does not exist: "+outputDir, 1)
		}
		abs, err := filepath.Abs(outputDir)
		if err != nil {
			return nil, cli.Exit("could not get absolute path for output directory: "+err.Error(), 1)
		}
		outDirs[i] = abs
	}

	if len(request.Recipients) > 0 && request.Script != "" {
		return nil, cli.Exit("encryption recipients and script cannot be used together", 1)
	}

	if (len(request.Recipients) > 0 || request.Script != "") && len(request.OutputDirs) == 0 {
		return nil, cli.Exit("encryption is not compatible with inline preparation and "+
			"requires at least one output directory", 1)
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
		return nil, cli.Exit(err.Error(), 1)
	}

	logger.Infof("Dataset created with ID: %d", dataset.ID)
	return &dataset, nil
}
