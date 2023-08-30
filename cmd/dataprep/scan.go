package dataprep

import (
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/urfave/cli/v2"
)

var StartScanCmd = &cli.Command{
	Name:      "start-scan",
	Usage:     "Start scanning of the source storage",
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
		job, err := dataprep.Default.StartScanHandler(c.Context, db, uint32(id), c.Args().Get(1))
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, job)
		return nil
	},
}

var PauseScanCmd = &cli.Command{
	Name:      "pause-scan",
	Usage:     "Pause a scanning job",
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
		job, err := dataprep.Default.PauseScanHandler(c.Context, db, uint32(id), c.Args().Get(1))
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, job)
		return nil
	},
}
