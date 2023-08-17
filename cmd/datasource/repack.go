package datasource

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/urfave/cli/v2"
)

var RepackCmd = &cli.Command{
	Name:      "repack",
	Usage:     "Retry packing a packjob or all errored packjobs of a data source",
	ArgsUsage: "<source_id> or --pack-job-id <pack_job_id>",
	Flags: []cli.Flag{
		&cli.Uint64Flag{
			Name:  "pack-job-id",
			Usage: "Pack job ID to retry packing",
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		defer closer.Close()
		var packJobID *uint64
		if c.IsSet("pack-job-id") {
			c2 := c.Uint64("pack-job-id")
			packJobID = &c2
		}
		packJobs, err := datasource.RepackHandler(
			c.Context,
			db,
			c.Args().Get(0),
			datasource.RepackRequest{
				PackJobID: packJobID,
			},
		)
		if err != nil {
			return err
		}

		cliutil.PrintToConsole(packJobs, c.Bool("json"), nil)
		return nil
	},
}
