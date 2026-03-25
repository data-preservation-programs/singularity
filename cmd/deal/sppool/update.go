package sppool

import (
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/deal/sppool"
	"github.com/urfave/cli/v2"
)

var UpdateCmd = &cli.Command{
	Name:      "update",
	Usage:     "Update an SP Pool's default deal parameters",
	Before:    cliutil.CheckNArgs,
	ArgsUsage: "<pool_id>",
	Flags:     commonDealFlags(),
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		poolID, err := strconv.ParseUint(c.Args().Get(0), 10, 32)
		if err != nil {
			return errors.Wrapf(err, "failed to parse pool ID %s", c.Args().Get(0))
		}

		request := sppool.UpdateRequest{}
		if c.IsSet("url-template") {
			v := c.String("url-template")
			request.URLTemplate = &v
		}
		if c.IsSet("price-per-gb-epoch") {
			v := c.Float64("price-per-gb-epoch")
			request.PricePerGBEpoch = &v
		}
		if c.IsSet("price-per-gb") {
			v := c.Float64("price-per-gb")
			request.PricePerGB = &v
		}
		if c.IsSet("price-per-deal") {
			v := c.Float64("price-per-deal")
			request.PricePerDeal = &v
		}
		if c.IsSet("verified") {
			v := c.Bool("verified")
			request.Verified = &v
		}
		if c.IsSet("ipni") {
			v := c.Bool("ipni")
			request.IPNI = &v
		}
		if c.IsSet("keep-unsealed") {
			v := c.Bool("keep-unsealed")
			request.KeepUnsealed = &v
		}
		if c.IsSet("start-delay") {
			v := c.String("start-delay")
			request.StartDelay = &v
		}
		if c.IsSet("duration") {
			v := c.String("duration")
			request.Duration = &v
		}
		if c.IsSet("schedule-cron") {
			v := c.String("schedule-cron")
			request.ScheduleCron = &v
		}
		if c.IsSet("schedule-cron-perpetual") {
			v := c.Bool("schedule-cron-perpetual")
			request.ScheduleCronPerpetual = &v
		}
		if c.IsSet("schedule-deal-number") {
			v := c.Int("schedule-deal-number")
			request.ScheduleDealNumber = &v
		}
		if c.IsSet("schedule-deal-size") {
			v := c.String("schedule-deal-size")
			request.ScheduleDealSize = &v
		}
		if c.IsSet("max-pending-deal-size") {
			v := c.String("max-pending-deal-size")
			request.MaxPendingDealSize = &v
		}
		if c.IsSet("max-pending-deal-number") {
			v := c.Int("max-pending-deal-number")
			request.MaxPendingDealNumber = &v
		}
		if c.IsSet("force") {
			v := c.Bool("force")
			request.Force = &v
		}

		pool, err := sppool.Default.UpdateHandler(c.Context, db, uint32(poolID), request)
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, pool)
		return nil
	},
}
