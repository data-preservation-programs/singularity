package run

import (
	"github.com/urfave/cli/v2"
)

var PrepWorkerCmd = &cli.Command{
	Name:  "prep-worker",
	Usage: "Start a dataset preparation worker to process dataset scanning and preparation tasks",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "concurrency",
			Usage:   "Number of concurrent workers to run",
			EnvVars: []string{"PREP_WORKER_CONCURRENCY"},
		},
	},
	Action: func(c *cli.Context) error {
		return nil
	},
}
