package dataprep

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/urfave/cli/v2"
)

var AutoDealCmd = &cli.Command{
	Name:     "autodeal",
	Usage:    "Auto-deal management commands",
	Category: "Auto-Deal Management",
	Subcommands: []*cli.Command{
		{
			Name:  "create",
			Usage: "Create automatic deal schedule for a specific preparation",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "preparation",
					Usage:    "Preparation ID or name",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				db, closer, err := database.OpenFromCLI(c)
				if err != nil {
					return errors.WithStack(err)
				}
				defer closer.Close()

				lotusClient := util.NewLotusClient(c.String("lotus-api"), c.String("lotus-token"))

				schedule, err := dataprep.DefaultAutoDealService.CreateAutomaticDealSchedule(
					c.Context,
					db.WithContext(c.Context),
					lotusClient,
					c.String("preparation"),
				)
				if err != nil {
					return errors.WithStack(err)
				}

				if schedule == nil {
					cliutil.Print(c, map[string]interface{}{
						"message": "Auto-deal creation not enabled for this preparation",
					})
				} else {
					cliutil.Print(c, *schedule)
				}
				return nil
			},
		},
		{
			Name:  "process",
			Usage: "Process all ready preparations for auto-deal creation",
			Action: func(c *cli.Context) error {
				db, closer, err := database.OpenFromCLI(c)
				if err != nil {
					return errors.WithStack(err)
				}
				defer closer.Close()

				lotusClient := util.NewLotusClient(c.String("lotus-api"), c.String("lotus-token"))

				err = dataprep.DefaultAutoDealService.ProcessReadyPreparations(
					c.Context,
					db.WithContext(c.Context),
					lotusClient,
				)
				if err != nil {
					return errors.WithStack(err)
				}

				cliutil.Print(c, map[string]interface{}{
					"message": "Auto-deal processing completed successfully",
				})
				return nil
			},
		},
		{
			Name:  "check",
			Usage: "Check if a preparation is ready for auto-deal creation",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "preparation",
					Usage:    "Preparation ID or name",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				db, closer, err := database.OpenFromCLI(c)
				if err != nil {
					return errors.WithStack(err)
				}
				defer closer.Close()

				isReady, err := dataprep.DefaultAutoDealService.CheckPreparationReadiness(
					c.Context,
					db.WithContext(c.Context),
					c.String("preparation"),
				)
				if err != nil {
					return errors.WithStack(err)
				}

				cliutil.Print(c, map[string]interface{}{
					"preparation": c.String("preparation"),
					"ready":       isReady,
				})
				return nil
			},
		},
	},
}