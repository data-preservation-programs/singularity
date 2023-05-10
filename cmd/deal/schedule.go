package deal

import "github.com/urfave/cli/v2"

var ScheduleCmd = &cli.Command{
	Name:      "schedule",
	Usage:     "Schedule a deal making batch request optionally with a schedule",
	ArgsUsage: "DATASET_NAME PROVIDER_ID",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:     "http-header",
			Category: "Deal Proposal",
			Aliases:  []string{"H"},
			Usage:    "http headers to be passed with the request (i.e. key=value)",
		},
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
			Usage:    "Price per 32GiB Deal over whole duration in Fil",
			Value:    0,
		},
		&cli.BoolFlag{
			Name:     "verified",
			Category: "Deal Proposal",
			Usage:    "Whether to propose deals as verified",
			Value:    true,
		},
		&cli.BoolFlag{
			Name:     "ipni",
			Category: "Deal Proposal",
			Usage:    "Whether to announce the deal to IPNI",
			Value:    true,
		},
		&cli.BoolFlag{
			Name:     "keep-unsealed",
			Category: "Deal Proposal",
			Usage:    "Whether to keep unsealed copy",
			Value:    true,
		},
		&cli.StringFlag{
			Name:        "cron-schedule",
			Category:    "Scheduling",
			Aliases:     []string{"c"},
			Usage:       "Cron schedule to send out batch deals",
			DefaultText: "disabled",
			Value:       "",
		},
		&cli.Float64Flag{
			Name:     "start-delay",
			Category: "Deal Proposal",
			Aliases:  []string{"s"},
			Usage:    "Deal start delay in days",
			Value:    3,
		},
		&cli.DurationFlag{
			Name:     "duration",
			Category: "Deal Proposal",
			Aliases:  []string{"d"},
			Usage:    "Duration in days for deal length",
			Value:    530,
		},
		&cli.Uint64Flag{
			Name:        "schedule-deal-number",
			Category:    "Scheduling",
			Aliases:     []string{"n"},
			Usage:       "Max deal number per triggered schedule, i.e. 30",
			DefaultText: "Unlimited",
		},
		&cli.StringFlag{
			Name:        "total-deal-number",
			Category:    "Scheduling",
			Aliases:     []string{"N"},
			Usage:       "Max total deal number for this request, i.e. 1000",
			DefaultText: "Unlimited",
		},
		&cli.StringFlag{
			Name:        "schedule-deal-size",
			Category:    "Scheduling",
			Aliases:     []string{"m"},
			Usage:       "Max deal sizes per triggered schedule, i.e. 500GB",
			DefaultText: "Unlimited",
		},
		&cli.StringFlag{
			Name:        "total-deal-size",
			Category:    "Restrictions",
			Aliases:     []string{"M"},
			Usage:       "Max total deal sizes for this request, i.e. 100TB",
			DefaultText: "Unlimited",
		},
		&cli.StringFlag{
			Name:     "notes",
			Category: "Tracking",
			Aliases:  []string{"n"},
			Usage:    "Any notes or tag to store along with the request, for tracking purpose",
			Value:    "",
		},
		&cli.StringFlag{
			Name:        "max-pending-deal-size",
			Category:    "Restrictions",
			Aliases:     []string{"Mp"},
			Usage:       "Max pending deal sizes overall for this request",
			DefaultText: "Unlimited",
		},
		&cli.StringFlag{
			Name:        "max-pending-deal-number",
			Category:    "Restrictions",
			Aliases:     []string{"Np"},
			Usage:       "Max pending deal number overall for this request",
			DefaultText: "Unlimited",
		},
	},
	Action: func(c *cli.Context) error {
		return nil
	},
}
