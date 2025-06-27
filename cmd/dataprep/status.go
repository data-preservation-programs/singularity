package dataprep

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/job"
	"github.com/urfave/cli/v2"
)

var StatusCmd = &cli.Command{
	Name:      "status",
	Usage:     "Get the preparation job status of a preparation",
	Category:  "Job Management",
	ArgsUsage: "<preparation id|name>",
	Before:    cliutil.CheckNArgs,
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()

		status, err := job.Default.GetStatusHandler(c.Context, db, c.Args().Get(0))
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, status)
		return nil
	},
}
