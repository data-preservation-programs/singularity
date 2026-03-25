package sppool

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/deal/sppool"
	"github.com/urfave/cli/v2"
)

var CreateCmd = &cli.Command{
	Name:  "create",
	Usage: "Create a new SP Pool with default deal parameters",
	Flags: append([]cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Usage:    "Unique name for the pool",
			Required: true,
		},
		&cli.StringFlag{
			Name:  "notes",
			Usage: "Notes",
		},
	}, commonDealFlags()...),
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		request := sppool.CreateRequest{
			Name:                  c.String("name"),
			Notes:                 c.String("notes"),
			HTTPHeaders:           c.StringSlice("http-header"),
			URLTemplate:           c.String("url-template"),
			PricePerGBEpoch:       c.Float64("price-per-gb-epoch"),
			PricePerGB:            c.Float64("price-per-gb"),
			PricePerDeal:          c.Float64("price-per-deal"),
			Verified:              c.Bool("verified"),
			IPNI:                  c.Bool("ipni"),
			KeepUnsealed:          c.Bool("keep-unsealed"),
			ScheduleCron:          c.String("schedule-cron"),
			StartDelay:            c.String("start-delay"),
			Duration:              c.String("duration"),
			ScheduleCronPerpetual: c.Bool("schedule-cron-perpetual"),
			ScheduleDealNumber:    c.Int("schedule-deal-number"),
			ScheduleDealSize:      c.String("schedule-deal-size"),
			MaxPendingDealSize:    c.String("max-pending-deal-size"),
			MaxPendingDealNumber:  c.Int("max-pending-deal-number"),
			Force:                 c.Bool("force"),
		}
		pool, err := sppool.Default.CreateHandler(c.Context, db, request)
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, pool)
		return nil
	},
}
