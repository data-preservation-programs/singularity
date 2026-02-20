package run

import (
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/data-preservation-programs/singularity/service/dealpusher"
	"github.com/data-preservation-programs/singularity/service/epochutil"
	"github.com/urfave/cli/v2"
)

var DealPusherCmd = &cli.Command{
	Name:  "deal-pusher",
	Usage: "Start a deal pusher that monitors deal schedules and pushes deals to storage providers",
	Flags: []cli.Flag{
		&cli.UintFlag{
			Name:    "deal-attempts",
			Usage:   "Number of times to attempt a deal before giving up",
			Aliases: []string{"d"},
			Value:   3,
		},
		&cli.UintFlag{
			Name:        "max-replication-factor",
			Usage:       "Max number of replicas for each individual PieceCID across all clients and providers",
			Aliases:     []string{"M"},
			DefaultText: "Unlimited",
		},
		&cli.IntFlag{
			Name:  "pdp-batch-size",
			Usage: "Number of roots to include in each PDP add-roots transaction",
			Value: 128,
		},
		&cli.Uint64Flag{
			Name:  "pdp-gas-limit",
			Usage: "Gas limit for PDP on-chain transactions",
			Value: 5000000,
		},
		&cli.Uint64Flag{
			Name:  "pdp-confirmation-depth",
			Usage: "Number of block confirmations required for PDP transactions",
			Value: 5,
		},
		&cli.DurationFlag{
			Name:  "pdp-poll-interval",
			Usage: "Polling interval for PDP transaction confirmation checks",
			Value: 30 * time.Second,
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()
		lotusAPI := c.String("lotus-api")
		lotusToken := c.String("lotus-token")
		err = epochutil.Initialize(c.Context, lotusAPI, lotusToken)
		if err != nil {
			return errors.WithStack(err)
		}

		pdpCfg := dealpusher.PDPSchedulingConfig{
			BatchSize:         c.Int("pdp-batch-size"),
			GasLimit:          c.Uint64("pdp-gas-limit"),
			ConfirmationDepth: c.Uint64("pdp-confirmation-depth"),
			PollingInterval:   c.Duration("pdp-poll-interval"),
		}
		if err := pdpCfg.Validate(); err != nil {
			return errors.WithStack(err)
		}

		dm, err := dealpusher.NewDealPusher(
			db,
			c.String("lotus-api"),
			c.String("lotus-token"),
			c.Uint("deal-attempts"),
			c.Uint("max-replication-factor"),
			dealpusher.WithPDPSchedulingConfig(pdpCfg),
		)
		if err != nil {
			return errors.WithStack(err)
		}
		return service.StartServers(c.Context, dealpusher.Logger, dm)
	},
}
