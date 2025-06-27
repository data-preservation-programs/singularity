package dataprep

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/job"
	"github.com/urfave/cli/v2"
)

var StartScanCmd = &cli.Command{
	Name:      "start-scan",
	Usage:     "Start scanning of the source storage",
	Category:  "Job Management",
	ArgsUsage: "<preparation id|name> <storage id|name>",
	Before:    cliutil.CheckNArgs,
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()
		job, err := job.Default.StartScanHandler(c.Context, db, c.Args().Get(0), c.Args().Get(1))
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
	ArgsUsage: "<preparation id|name> <storage id|name>",
	Before:    cliutil.CheckNArgs,
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()
		job, err := job.Default.PauseScanHandler(c.Context, db, c.Args().Get(0), c.Args().Get(1))
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, job)
		return nil
	},
}
