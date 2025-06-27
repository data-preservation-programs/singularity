package dataprep

import (
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/job"
	"github.com/urfave/cli/v2"
)

var StartPackCmd = &cli.Command{
	Name:      "start-pack",
	Usage:     "Start / Restart all pack jobs or a specific one",
	Category:  "Job Management",
	ArgsUsage: "<preparation id|name> <storage id|name> [job_id]",
	Before:    cliutil.CheckNArgs,
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()
		var jobID int64
		if c.Args().Get(2) != "" {
			jobID, err = strconv.ParseInt(c.Args().Get(2), 10, 64)
			if err != nil {
				return errors.Wrapf(err, "invalid job ID '%s'", c.Args().Get(2))
			}
		}
		job, err := job.Default.StartPackHandler(c.Context, db, c.Args().Get(0), c.Args().Get(1), jobID)
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, job)
		return nil
	},
}

var PausePackCmd = &cli.Command{
	Name:      "pause-pack",
	Usage:     "Pause all pack jobs or a specific one",
	Category:  "Job Management",
	ArgsUsage: "<preparation id|name> <storage id|name> [job_id]",
	Before:    cliutil.CheckNArgs,
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()
		var jobID int64
		if c.Args().Get(2) != "" {
			jobID, err = strconv.ParseInt(c.Args().Get(2), 10, 64)
			if err != nil {
				return errors.Wrapf(err, "invalid job ID '%s'", c.Args().Get(2))
			}
		}
		job, err := job.Default.PausePackHandler(c.Context, db, c.Args().Get(0), c.Args().Get(1), jobID)
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, job)
		return nil
	},
}
