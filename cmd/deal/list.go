package deal

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/deal"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/rjNemo/underscore"
	"github.com/urfave/cli/v2"
)

var ListCmd = &cli.Command{
	Name:  "list",
	Usage: "List all deals",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:  "preparation",
			Usage: "Filter deals by preparation id or name",
		},
		&cli.StringSliceFlag{
			Name:  "source",
			Usage: "Filter deals by source storage id or name",
		},
		&cli.UintSliceFlag{
			Name:  "schedule",
			Usage: "Filter deals by schedule",
		},
		&cli.StringSliceFlag{
			Name:  "provider",
			Usage: "Filter deals by provider",
		},
		&cli.StringSliceFlag{
			Name:  "state",
			Usage: "Filter deals by state: proposed, published, active, expired, proposal_expired, slashed",
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()
		deals, err := deal.Default.ListHandler(c.Context, db, deal.ListDealRequest{
			Preparations: c.StringSlice("preparation"),
			Sources:      c.StringSlice("source"),
			Schedules:    underscore.Map(c.IntSlice("schedules"), func(i int) uint32 { return uint32(i) }),
			Providers:    c.StringSlice("provider"),
			States:       underscore.Map(c.StringSlice("state"), func(s string) model.DealState { return model.DealState(s) }),
		})
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, deals)
		return nil
	},
}
