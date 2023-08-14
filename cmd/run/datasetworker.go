package run

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/datasetworker"
	"github.com/urfave/cli/v2"
)

var DatasetWorkerCmd = &cli.Command{
	Name:  "dataset-worker",
	Usage: "Start a dataset preparation worker to process dataset scanning and preparation tasks",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "concurrency",
			Usage:   "Number of concurrent workers to run",
			EnvVars: []string{"DATASET_WORKER_CONCURRENCY"},
			Value:   1,
		},
		&cli.BoolFlag{
			Name:    "enable-scan",
			Usage:   "Enable scanning of datasets",
			EnvVars: []string{"DATASET_WORKER_ENABLE_SCAN"},
			Value:   true,
		},
		&cli.BoolFlag{
			Name:    "enable-pack",
			Usage:   "Enable packing of datasets that calculates CIDs and packs them into CAR files",
			EnvVars: []string{"DATASET_WORKER_ENABLE_PACK"},
			Value:   true,
		},
		&cli.BoolFlag{
			Name:    "enable-dag",
			Usage:   "Enable dag generation of datasets that maintains the directory structure of datasets",
			EnvVars: []string{"DATASET_WORKER_ENABLE_DAG"},
			Value:   true,
		},
		&cli.BoolFlag{
			Name:    "exit-on-complete",
			Usage:   "Exit the worker when there is no more work to do",
			EnvVars: []string{"DATASET_WORKER_EXIT_ON_COMPLETE"},
			Value:   false,
		},
		&cli.BoolFlag{
			Name:    "exit-on-error",
			Usage:   "Exit the worker when there is any error",
			EnvVars: []string{"DATASET_WORKER_EXIT_ON_ERROR"},
			Value:   false,
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		defer closer.Close()
		if err := model.AutoMigrate(db); err != nil {
			return err
		}
		worker := datasetworker.NewWorker(
			db,
			datasetworker.Config{
				Concurrency:    c.Int("concurrency"),
				EnableScan:     c.Bool("enable-scan"),
				EnablePack:     c.Bool("enable-pack"),
				EnableDag:      c.Bool("enable-dag"),
				ExitOnComplete: c.Bool("exit-on-complete"),
				ExitOnError:    c.Bool("exit-on-error"),
			})
		err = worker.Run(c.Context)
		if err != nil {
			return err
		}
		return nil
	},
}
