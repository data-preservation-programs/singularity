package ez

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/admin"
	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/data-preservation-programs/singularity/handler/job"
	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/data-preservation-programs/singularity/service/datasetworker"
	"github.com/data-preservation-programs/singularity/service/workflow"
	"github.com/urfave/cli/v2"
)

var PrepCmd = &cli.Command{
	Name:      "ez-prep",
	Category:  "Utility",
	Before:    cliutil.CheckNArgs,
	ArgsUsage: "<path>",
	Usage:     "Prepare a dataset from a local path",
	Description: "This commands can be used to prepare a dataset from a local path with minimum configurable parameters.\n" +
		"For more advanced usage, please use the subcommands under `storage` and `data-prep`.\n" +
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
			Aliases:     []string{"f"},
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
				databaseFile = "file::memory:"
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
		db, closer, err := database.OpenWithLogger("sqlite:" + databaseFile)
		if err != nil {
			return errors.Wrapf(err, "failed to open database %s", databaseFile)
		}

		defer func() { _ = closer.Close() }()

		// Step 1, initialize the database
		err = admin.Default.InitHandler(c.Context, db)
		if err != nil {
			return errors.WithStack(err)
		}

		// Disable workflow orchestrator to prevent automatic job progression
		// We manage job progression manually in ez-prep
		workflow.DefaultOrchestrator.SetEnabled(false)
		fmt.Println("⚠️  Workflow orchestrator disabled: manual job progression enabled for ez-prep.")

		// Step 2, create a preparation
		outputDir := c.String("output-dir")
		var outputStorages []string
		if outputDir != "" {
			err = os.MkdirAll(outputDir, 0o750)
			if err != nil {
				return errors.Wrap(err, "failed to create output directory")
			}

			_, err = storage.Default.CreateStorageHandler(c.Context, db, "local", storage.CreateRequest{
				Name: "output",
				Path: outputDir,
			})
			if err != nil {
				return errors.Wrap(err, "failed to create output storage")
			}
			outputStorages = []string{"output"}
		}

		_, err = storage.Default.CreateStorageHandler(c.Context, db, "local", storage.CreateRequest{
			Name: "source",
			Path: path,
		})
		if err != nil {
			return errors.Wrap(err, "failed to create source storage")
		}

		_, err = dataprep.Default.CreatePreparationHandler(c.Context, db, dataprep.CreateRequest{
			SourceStorages: []string{"source"},
			OutputStorages: outputStorages,
			MaxSizeStr:     c.String("max-size"),
			Name:           "preparation",
		})
		if err != nil {
			return errors.Wrap(err, "failed to create preparation")
		}

		_, err = job.Default.StartScanHandler(c.Context, db, "preparation", "source")
		if err != nil {
			return errors.Wrap(err, "failed to start scan")
		}

		// Step 3, start dataset worker
		worker := datasetworker.NewWorker(
			db,
			datasetworker.Config{
				Concurrency:    1,
				EnableScan:     true,
				ExitOnComplete: true,
				ExitOnError:    true,
			})
		err = worker.Run(c.Context)
		if err != nil {
			return errors.Wrap(err, "failed to run dataset worker for scanning")
		}

		worker = datasetworker.NewWorker(
			db,
			datasetworker.Config{
				Concurrency:    c.Int("concurrency"),
				EnablePack:     true,
				ExitOnComplete: true,
				ExitOnError:    true,
			})
		err = worker.Run(c.Context)
		if err != nil {
			return errors.Wrap(err, "failed to run dataset worker for packing")
		}

		// Step 4, Initiate dag gen
		_, err = job.Default.StartDagGenHandler(c.Context, db, "preparation", "source")
		if err != nil {
			return errors.Wrap(err, "failed to start dag gen")
		}

		// Step 5, start dataset worker again
		worker = datasetworker.NewWorker(
			db,
			datasetworker.Config{
				Concurrency:    1,
				EnableDag:      true,
				ExitOnComplete: true,
				ExitOnError:    true,
			})
		err = worker.Run(c.Context)
		if err != nil {
			return errors.Wrap(err, "failed to run dataset worker")
		}

		// Step 6, print all information
		pieceLists, err := dataprep.Default.ListPiecesHandler(
			c.Context, db, "preparation",
		)
		if err != nil {
			return errors.Wrap(err, "failed to list pieces")
		}

		cliutil.Print(c, pieceLists[0].Pieces)
		return nil
	},
}
