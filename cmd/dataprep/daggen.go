package dataprep

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/job"
	"github.com/urfave/cli/v2"
)

var StartDagGenCmd = &cli.Command{
	Name:      "start-daggen",
	Usage:     "Start a DAG generation that creates a snapshot of all folder structures",
	Category:  "Job Management",
	ArgsUsage: "<preparation id|name> <storage id|name>",
	Before:    cliutil.CheckNArgs,
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()
		job, err := job.Default.StartDagGenHandler(c.Context, db, c.Args().Get(0), c.Args().Get(1))
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, job)
		return nil
	},
}

var PauseDagGenCmd = &cli.Command{
	Name:      "pause-daggen",
	Usage:     "Pause a DAG generation job",
	Category:  "Job Management",
	ArgsUsage: "<preparation_id> <storage_name>",
	Before:    cliutil.CheckNArgs,
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()
		job, err := job.Default.PauseDagGenHandler(c.Context, db, c.Args().Get(0), c.Args().Get(1))
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, job)
		return nil
	},
}
