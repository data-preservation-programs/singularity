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
	"github.com/data-preservation-programs/singularity/util"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

var CreateCmd = &cli.Command{
	Name:     "create",
	Usage:    "Create a new preparation",
	Category: "Preparation Management",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "name",
			Usage:       "The name for the preparation",
			DefaultText: "Auto generated",
		},
		&cli.StringSliceFlag{
			Name:  "source",
			Usage: "The id or name of the source storage to be used for the preparation",
		},
		&cli.StringSliceFlag{
			Name:  "output",
			Usage: "The id or name of the output storage to be used for the preparation",
		},
		&cli.StringSliceFlag{
			Name:     "local-source",
			Category: "Quick creation with local source paths",
			Usage:    "The local source path to be used for the preparation. This is a convenient flag that will create a source storage with the provided path",
		},
		&cli.StringSliceFlag{
			Name:     "local-output",
			Category: "Quick creation with local output paths",
			Usage:    "The local output path to be used for the preparation. This is a convenient flag that will create a output storage with the provided path",
		},
		&cli.StringFlag{
			Name:  "max-size",
			Usage: "The maximum size of a single CAR file",
			Value: "31.5GiB",
		},
		&cli.StringFlag{
			Name:        "piece-size",
			Usage:       "The target piece size of the CAR files used for piece commitment calculation",
			Value:       "",
			DefaultText: "Determined by --max-size",
		},
		&cli.BoolFlag{
			Name:  "delete-after-export",
			Usage: "Whether to delete the source files after export to CAR files",
		},
		&cli.BoolFlag{
			Name:  "no-inline",
			Usage: "Whether to disable inline storage for the preparation. Can save database space but requires at least one output storage.",
		},
		&cli.BoolFlag{
			Name:  "no-dag",
			Usage: "Whether to disable maintaining folder dag structure for the sources. If disabled, DagGen will not be possible and folders will not have an associated CID.",
		},
		&cli.BoolFlag{
			Name:  "auto",
			Usage: "Whether to automatically start pack and daggen jobs after scan. If disabled, jobs will need to be manually started.",
			Value: true,
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()
		db = db.WithContext(c.Context)
		name := c.String("name")
		if name == "" {
			name = util.RandomName()
		}
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

		prep, err := dataprep.Default.CreatePreparationHandler(c.Context, db, dataprep.CreateRequest{
			SourceStorages:    sourceStorages,
			OutputStorages:    outputStorages,
			MaxSizeStr:        maxSizeStr,
			PieceSizeStr:      pieceSizeStr,
			DeleteAfterExport: c.Bool("delete-after-export"),
			Name:              name,
			NoInline:          c.Bool("no-inline"),
			NoDag:             c.Bool("no-dag"),
			Auto:              c.Bool("auto"),
		})
		if err != nil {
			return errors.WithStack(err)
		}

		cliutil.Print(c, *prep)
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
		existing, err = storage.Default.CreateStorageHandler(
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
		//nolint:gosec
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
