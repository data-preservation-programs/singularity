package handler

import (
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/data-preservation-programs/go-singularity/util"
	"github.com/dustin/go-humanize"
	"github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

func CreateHandler(c *cli.Context) error {
	logger := log.Logger("cli")
	log.SetAllLoggers(log.LevelInfo)
	name := c.Args().Get(0)
	if name == "" {
		return cli.Exit("dataset name is required", 1)
	}

	minSize, err := humanize.ParseBytes(c.String("min-size"))
	if err != nil {
		return cli.Exit("invalid value for min-size: "+err.Error(), 1)
	}

	maxSize, err := humanize.ParseBytes(c.String("max-size"))
	if err != nil {
		return cli.Exit("invalid value for max-size: "+err.Error(), 1)
	}

	pieceSize := util.NextPowerOfTwo(maxSize)
	if c.String("piece-size") != "" {
		pieceSize, err = humanize.ParseBytes(c.String("piece-size"))
		if err != nil {
			return cli.Exit("invalid value for piece-size: "+err.Error(), 1)
		}
	}

	if pieceSize > 1<<36 {
		return cli.Exit("piece size cannot be larger than 64 GiB", 1)
	}

	if maxSize*128/127 >= pieceSize {
		return cli.Exit("max size needs to be reduced to leave space for padding", 1)
	}

	outputDirs := c.StringSlice("output-dir")
	rescanInterval := c.Duration("rescan-interval")
	if len(outputDirs) == 0 && rescanInterval > 0 {
		return cli.Exit("output directory is required for live monitoring", 1)
	}

	maxWait := c.Duration("max-wait")

	recipients := c.StringSlice("encryption-recipient")
	script := c.String("encryption-script")
	if len(recipients) > 0 && script != "" {
		return cli.Exit("encryption recipients and script cannot be used together", 1)
	}

	sourceArgs := c.Args().Slice()[1:]
	sources := make([]model.Source, len(sourceArgs))
	for i, sourceArg := range sourceArgs {
		source, err := model.NewSource(sourceArg)
		if err != nil {
			return cli.Exit("invalid dataset source: "+err.Error(), 1)
		}

		source.ScanInterval = rescanInterval
		source.ScanningState = model.Ready
		source.MaxWait = maxWait
		sources[i] = *source
	}
	db := database.MustOpenFromCLI(c)
	dataset := model.Dataset{
		Name:                 name,
		MinSize:              minSize,
		MaxSize:              maxSize,
		PieceSize:            pieceSize,
		OutputDirs:           outputDirs,
		EncryptionRecipients: recipients,
		EncryptionScript:     script,
	}

	return db.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&dataset)
		if result.Error != nil {
			return cli.Exit(result.Error.Error(), 1)
		}

		logger.Infof("Dataset created with ID: %d", dataset.ID)

		for _, source := range sources {
			rootDirectory := model.Directory{}
			err := tx.Create(&rootDirectory).Error
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}
			source.DatasetID = dataset.ID
			source.RootDirectoryID = rootDirectory.ID
			err = tx.Create(&source).Error
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}

			logger.Infof("Dataset source created with ID: %d, path: %s", source.ID, source.Path)
		}

		return nil
	})
}
