package run

import (
	"time"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
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
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		defer closer.Close()
		if err := model.AutoMigrate(db); err != nil {
			return err
		}

		lotusAPI := c.String("lotus-api")
		lotusToken := c.String("lotus-token")
		err = epochutil.Initialize(c.Context, lotusAPI, lotusToken)
		if err != nil {
			return err
		}

		tracker := dealtracker.NewDealTracker(db,
			c.Duration("interval"),
			c.String("market-deal-url"),
			c.String("lotus-api"),
			c.String("lotus-token"),
		)

		return tracker.Run(c.Context)
	},
}
