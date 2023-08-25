package dataprep

import (
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/urfave/cli/v2"
)

var StartDagGenCmd = &cli.Command{
	Name:      "start-daggen",
	Usage:     "Start a DAG generation that creates a snapshot of all folder structures",
	Category:  "Job Management",
	ArgsUsage: "<preparation_id> <storage_name>",
	Before:    cliutil.CheckNArgs,
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()
		id, err := strconv.ParseUint(c.Args().Get(0), 10, 32)
		if err != nil {
			return errors.Wrapf(err, "invalid preparation ID '%s'", c.Args().Get(0))
		}
		job, err := dataprep.StartDagGenHandler(c.Context, db, uint32(id), c.Args().Get(1))
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.PrintToConsole(c, job)
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
		defer closer.Close()
		id, err := strconv.ParseUint(c.Args().Get(0), 10, 32)
		if err != nil {
			return errors.Wrapf(err, "invalid preparation ID '%s'", c.Args().Get(0))
		}
		job, err := dataprep.PauseDagGenHandler(c.Context, db, uint32(id), c.Args().Get(1))
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.PrintToConsole(c, job)
		return nil
	},
}
