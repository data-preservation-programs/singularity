package ez

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/admin"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/datasetworker"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

var PrepCmd = &cli.Command{
	Name:      "ez-prep",
	Category:  "Easy Commands",
	ArgsUsage: "<path>",
	Usage:     "Prepare a dataset from a local path",
	Description: "This commands can be used to prepare a dataset from a local path with minimum configurable parameters.\n" +
		"For more advanced usage, please use the subcommands under `dataset` and `datasource`.\n" +
		"You can also use this command for benchmarking with in-memory database and inline preparation, i.e.\n" +
		"  mkdir dataset\n" +
		"  truncate -s 1024G dataset/1T.bin\n" +
		"  singularity ez-prep --output-dir '' --database-file '' -j $(($(nproc) / 4 + 1)) ./dataset",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "max-size",
			Aliases: []string{"M"},
			Usage:   "Maximum size of the CAR files to be created",
			Value:   "31.5GiB",
		},
		&cli.StringFlag{
			Name:    "output-dir",
			Aliases: []string{"o"},
			Usage:   "Output directory for CAR files. To use inline preparation, use an empty string",
			Value:   "./cars",
		},
		&cli.IntFlag{
			Name:    "concurrency",
			Aliases: []string{"j"},
			Usage:   "Concurrency for packing",
			Value:   1,
		},
		&cli.StringFlag{
			Name:        "database-file",
			Usage:       "The database file to store the metadata. To use in memory database, use an empty string.",
			DefaultText: "./ezprep-<name>.db",
		},
	},
	Action: func(c *cli.Context) error {
		t := time.Now().Unix()
		path := c.Args().Get(0)
		if path == "" {
			return errors.New("path is required")
		}
		databaseFile := c.String("database-file")
		if databaseFile == "" {
			if c.IsSet("database-file") {
				databaseFile = "file::memory:?cache=shared"
			} else {
				databaseFile = fmt.Sprintf("./ezprep-%d.db", t)
			}
		}
		var err error
		if !strings.HasPrefix(databaseFile, "file::memory") {
			databaseFile, err = filepath.Abs(databaseFile)
			if err != nil {
				return errors.Wrap(err, "failed to get absolute path")
			}
		}
		db, closer, err := database.Open("sqlite:"+databaseFile, &gorm.Config{})
		if err != nil {
			return errors.Wrapf(err, "failed to open database %s", databaseFile)
		}

		defer closer.Close()

		// Step 1, initialize the database
		err = admin.InitHandler(c.Context, db)
		if err != nil {
			return err
		}

		// Step 2, create a dataset
		var outputDirs []string
		if c.String("output-dir") != "" {
			outputDirs = []string{c.String("output-dir")}
			err = os.MkdirAll(outputDirs[0], 0755)
			if err != nil {
				return errors.Wrap(err, "failed to create output directory")
			}
		}
		ds, err2 := dataset.CreateHandler(c.Context, db, dataset.CreateRequest{
			Name:       "ez",
			MaxSizeStr: c.String("max-size"),
			OutputDirs: outputDirs,
		})
		if err2 != nil {
			return err2
		}

		// Step 3, add a local data source
		path, err = filepath.Abs(path)
		if err != nil {
			return errors.Wrap(err, "failed to get absolute path")
		}
		source := model.Source{
			DatasetID:     ds.ID,
			Type:          "local",
			Path:          path,
			Metadata:      model.Metadata(nil),
			ScanningState: model.Ready,
			DagGenState:   model.Created,
		}
		err = db.Create(&source).Error
		if err != nil {
			return errors.Wrap(err, "failed to create source")
		}

		root := model.Directory{
			SourceID: source.ID,
			Name:     path,
		}
		err = db.Create(&root).Error
		if err != nil {
			return errors.Wrap(err, "failed to create root directory")
		}

		// Step 3, start dataset worker
		worker := datasetworker.NewWorker(
			db,
			datasetworker.Config{
				Concurrency:    c.Int("concurrency"),
				EnableScan:     true,
				EnablePack:     true,
				EnableDag:      true,
				ExitOnComplete: true,
			})
		err = worker.Run(c.Context)
		if err != nil {
			return err
		}

		// Step 4, Initiate dag gen
		_, err2 = datasource.DagGenHandler(c.Context, db, strconv.Itoa(int(source.ID)))
		if err2 != nil {
			return err2
		}

		// Step 5, start dataset worker again
		err = worker.Run(c.Context)
		if err != nil {
			return err
		}

		// Step 6, print all information
		cars, err2 := dataset.ListPiecesHandler(
			c.Context, db, ds.Name,
		)
		if err2 != nil {
			return err2
		}

		cliutil.PrintToConsole(cars, false, nil)
		return nil
	},
}
