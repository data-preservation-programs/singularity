package repl

import "github.com/urfave/cli/v2"

var ScheduleCmd = &cli.Command{
	Name:      "schedule",
	Usage:     "Schedule a deal making batch request optionally with a schedule",
	ArgsUsage: "DATASET_NAME PROVIDER_ID",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "url-template",
			Category: "Deal Proposal",
			Aliases:  []string{"u"},
			Usage:    "URL template with PIECE_CID placeholder for boost to fetch the CAR file, i.e. http://127.0.0.1/piece/{PIECE_CID}.car",
			Value:    "",
		},
		&cli.Float64Flag{
			Name:     "price",
			Category: "Deal Proposal",
			Aliases:  []string{"p"},
			Usage:    "Price per 32GiB Deal in Fil",
			Value:    0,
		},
		&cli.BoolFlag{
			Name:     "verified",
			Category: "Deal Proposal",
			Usage:    "Whether to propose deals as verified",
			Value:    true,
		},
		&cli.DurationFlag{
			Name:        "batch-interval",
			Category:    "Scheduling",
			Aliases:     []string{"i"},
			Usage:       "Interval to send out batch deals",
			DefaultText: "disabled",
			Value:       0,
		},
		&cli.Float64Flag{
			Name:        "start-delay",
			Category:    "Deal Proposal",
			Aliases:     []string{"s"},
			Usage:       "Deal start delay in days",
			DefaultText: "7.0",
			Value:       7,
		},
		&cli.DurationFlag{
			Name:        "duration",
			Category:    "Deal Proposal",
			Aliases:     []string{"d"},
			Usage:       "Duration in days for deal length",
			DefaultText: "530.0",
			Value:       530,
		},
		&cli.Uint64Flag{
			Name:        "batch-deal-size",
			Category:    "Scheduling",
			Aliases:     []string{"m"},
			Usage:       "Max deal sizes per batch measured in GiB",
			DefaultText: "Unlimited",
			Value:       0,
		},
		&cli.Uint64Flag{
			Name:        "max-deal-size",
			Category:    "Restrictions",
			Aliases:     []string{"M"},
			Usage:       "Max deal sizes overall measured in GiB",
			DefaultText: "Unlimited",
			Value:       0,
		},
		&cli.StringFlag{
			Name:     "notes",
			Category: "Tracking",
			Aliases:  []string{"n"},
			Usage:    "Any notes or tag to store along with the request, for tracking purpose",
			Value:    "",
		},
		&cli.StringFlag{
			Name:     "max-pending-deal-size",
			Category: "Restrictions",
			Aliases:  []string{"Mp"},
			Usage:    "Max pending deal sizes overall measured in GiB",
		},
		&cli.StringSliceFlag{
			Name:     "client",
			Category: "Deal Proposal",
			Aliases:  []string{"c"},
			Usage:    "Client wallet address to use for deal making",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "piece-list",
			Category: "Restrictions",
			Aliases:  []string{"l"},
			Usage:    "Path to a file that contains a list of piece CIDs to be included, one CID per line",
		},
		&cli.StringFlag{
			Name:     "order",
			Category: "Scheduling",
			Aliases:  []string{"o"},
			Usage:    "Order to send out deals, must be one of: random, min_replicas, time_asc, time_desc, piece_size_asc, piece_size_desc, file",
			Value:    "min_replicas",
		},
	},
	Action: func(c *cli.Context) error {
		return nil
	},
}
