package dataprep

import (
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/urfave/cli/v2"
)

var StatusCmd = &cli.Command{
	Name:      "status",
	Usage:     "Get the preparation job status of a preparation",
	Category:  "Job Management",
	ArgsUsage: "<preparation_id>",
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

		status, err := dataprep.GetStatusHandler(c.Context, db, uint32(id))
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.PrintToConsole(c, status)
		return nil
	},
}
