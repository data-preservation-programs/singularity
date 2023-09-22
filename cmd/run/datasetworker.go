package run

import (
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/service/datasetworker"
	"github.com/urfave/cli/v2"
)

var DatasetWorkerCmd = &cli.Command{
	Name:  "dataset-worker",
	Usage: "Start a dataset preparation worker to process dataset scanning and preparation tasks",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:  "concurrency",
			Usage: "Number of concurrent workers to run",
			Value: 1,
		},
		&cli.BoolFlag{
			Name:  "enable-scan",
			Usage: "Enable scanning of datasets",
			Value: true,
		},
		&cli.BoolFlag{
			Name:  "enable-pack",
			Usage: "Enable packing of datasets that calculates CIDs and packs them into CAR files",
			Value: true,
		},
		&cli.BoolFlag{
			Name:  "enable-dag",
			Usage: "Enable dag generation of datasets that maintains the directory structure of datasets",
			Value: true,
		},
		&cli.BoolFlag{
			Name:  "exit-on-complete",
			Usage: "Exit the worker when there is no more work to do",
			Value: false,
		},
		&cli.BoolFlag{
			Name:  "exit-on-error",
			Usage: "Exit the worker when there is any error",
			Value: false,
		},
		&cli.DurationFlag{
			Name:  "min-interval",
			Usage: "How often to check for new jobs (minimum)",
			Value: 5 * time.Second,
		},
		&cli.DurationFlag{
			Name:  "max-interval",
			Usage: "How often to check for new jobs (maximum)",
			Value: 160 * time.Second,
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()
		worker := datasetworker.NewWorker(
			db,
			datasetworker.Config{
				Concurrency:    c.Int("concurrency"),
				EnableScan:     c.Bool("enable-scan"),
				EnablePack:     c.Bool("enable-pack"),
				EnableDag:      c.Bool("enable-dag"),
				ExitOnComplete: c.Bool("exit-on-complete"),
				ExitOnError:    c.Bool("exit-on-error"),
				MinInterval:    c.Duration("min-interval"),
				MaxInterval:    c.Duration("max-interval"),
			})
		err = worker.Run(c.Context)
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	},
}
