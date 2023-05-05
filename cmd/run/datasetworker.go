package run

import (
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
		return nil
	},
}
