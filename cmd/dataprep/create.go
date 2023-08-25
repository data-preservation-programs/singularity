package dataprep

import (
	"context"
	"math/rand"
	"path/filepath"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

var CreateCmd = &cli.Command{
	Name:     "create",
	Usage:    "Create a new preparation",
	Category: "Preparation Management",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:  "source",
			Usage: "The name of the source storage to be used for the preparation",
		},
		&cli.StringSliceFlag{
			Name:  "output",
			Usage: "The name of the output storage to be used for the preparation",
		},
		&cli.StringSliceFlag{
			Name:     "local-source",
			Category: "Quick creation with local paths",
			Usage:    "The local source path to be used for the preparation. This is a convenient flag that will create a source storage with the provided path",
		},
		&cli.StringSliceFlag{
			Name:     "local-output",
			Category: "Quick creation with local paths",
			Usage:    "The local output path to be used for the preparation. This is a convenient flag that will create a output storage with the provided path",
		},
		&cli.StringFlag{
			Name:  "max-size",
			Usage: "The maximum size of a single CAR file",
			Value: "31.5GiB",
		},
		&cli.StringFlag{
			Name:  "piece-size",
			Usage: "The target piece size of the CAR files used for piece commitment calculation",
			Value: "32GiB",
		},
		&cli.StringSliceFlag{
			Name:  "encryption-recipient",
			Usage: "The public key of the encryption recipient",
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenInMemory()
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()
		db = db.WithContext(c.Context)
		sourceStorages := c.StringSlice("source")
		outputStorages := c.StringSlice("output")
		maxSizeStr := c.String("max-size")
		pieceSizeStr := c.String("piece-size")
		for _, sourcePath := range c.StringSlice("local-source") {
			source, err := createStorageIfNotExist(c.Context, db, sourcePath)
			if err != nil {
				return errors.WithStack(err)
			}
			sourceStorages = append(sourceStorages, source.Name)
		}
		for _, outputPath := range c.StringSlice("local-output") {
			output, err := createStorageIfNotExist(c.Context, db, outputPath)
			if err != nil {
				return errors.WithStack(err)
			}
			outputStorages = append(outputStorages, output.Name)
		}

		prep, err := dataprep.CreatePreparationHandler(c.Context, db, dataprep.CreateRequest{
			SourceStorages:       sourceStorages,
			OutputStorages:       outputStorages,
			MaxSizeStr:           maxSizeStr,
			PieceSizeStr:         pieceSizeStr,
			EncryptionRecipients: c.StringSlice("encryption-recipient"),
		})
		if err != nil {
			return errors.WithStack(err)
		}

		cliutil.PrintToConsole(c, prep)
		return nil
	},
}

func createStorageIfNotExist(ctx context.Context, db *gorm.DB, sourcePath string) (*model.Storage, error) {
	db = db.WithContext(ctx)
	path, err := filepath.Abs(sourcePath)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid path: %s", sourcePath)
	}
	existing := &model.Storage{}
	err = db.Where("type = ? AND path = ?", "local", path).
		First(existing).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.WithStack(err)
	}
	if err != nil {
		name := filepath.Base(path)
		if name == "." {
			name = ""
		}
		name += "-" + randomReadableString(4)
		existing, err = storage.CreateStorageHandler(
			ctx,
			db,
			"local", storage.CreateRequest{
				Name: name,
				Path: path,
			})
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return existing, nil

}

func randomReadableString(length int) string {
	const charset = "0123456789abcdef"

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
