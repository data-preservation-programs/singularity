package run

import (
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/data-preservation-programs/go-singularity/service"
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
	},
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		if err := model.InitializeEncryption(c.String("password"), db); err != nil {
			return cli.Exit("Cannot initialize encryption"+err.Error(), 1)
		}
		worker := service.NewDatasetWorker(db, c.Int("concurrency"))
		err := worker.Run(c.Context)
		if err != nil {
			return err
		}
		return nil
	},
}
