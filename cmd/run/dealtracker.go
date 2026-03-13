package run

import (
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/data-preservation-programs/singularity/service/dealtracker"
	"github.com/data-preservation-programs/singularity/service/epochutil"
	"github.com/urfave/cli/v2"
)

var DealTrackerCmd = &cli.Command{
	Name:  "deal-tracker",
	Usage: "Start a deal tracker that tracks the deal for all relevant wallets",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "market-deal-url",
			Usage:   "The URL for ZST compressed state market deals json. Set to empty to use Lotus API.",
			Aliases: []string{"m"},
			EnvVars: []string{"MARKET_DEAL_URL"},
			Value:   "https://marketdeals.s3.amazonaws.com/StateMarketDeals.json.zst",
		},
		&cli.DurationFlag{
			Name:    "interval",
			Usage:   "How often to check for new deals",
			Aliases: []string{"i"},
			Value:   1 * time.Hour,
		},
		&cli.BoolFlag{
			Name:  "once",
			Usage: "Run once and exit",
			Value: false,
		},
		&cli.StringFlag{
			Name:    "eth-rpc",
			Usage:   "Ethereum RPC endpoint for FEVM (required for DDO allocation tracking)",
			EnvVars: []string{"ETH_RPC_URL"},
		},
		&cli.StringFlag{
			Name:    "ddo-contract",
			Usage:   "DDO Diamond proxy contract address",
			EnvVars: []string{"DDO_CONTRACT_ADDRESS"},
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

		var opts []dealtracker.DealTrackerOption
		if ddoContract := c.String("ddo-contract"); ddoContract != "" {
			rpcURL := c.String("eth-rpc")
			if rpcURL == "" {
				return errors.New("--eth-rpc is required when --ddo-contract is set")
			}
			ddoClient, err := dealtracker.NewDDOTrackingClient(rpcURL, ddoContract)
			if err != nil {
				return errors.Wrap(err, "failed to initialize DDO tracking client")
			}
			defer ddoClient.Close()
			opts = append(opts, dealtracker.WithDDOAllocationTracker(ddoClient))
		}

		tracker := dealtracker.NewDealTracker(db,
			c.Duration("interval"),
			c.String("market-deal-url"),
			c.String("lotus-api"),
			c.String("lotus-token"),
			c.Bool("once"),
			opts...,
		)

		return service.StartServers(c.Context, dealtracker.Logger, &tracker)
	},
}
