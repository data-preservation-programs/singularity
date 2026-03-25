package sppool

import "github.com/urfave/cli/v2"

// commonDealFlags returns the shared deal-parameter flags used by create and update.
func commonDealFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringSliceFlag{
			Name:     "http-header",
			Category: "Boost Only",
			Aliases:  []string{"H"},
			Usage:    "HTTP headers to be passed with the request (i.e. key=value)",
		},
		&cli.StringFlag{
			Name:     "url-template",
			Category: "Boost Only",
			Aliases:  []string{"u"},
			Usage:    "URL template with PIECE_CID placeholder for boost to fetch the CAR file",
		},
		&cli.Float64Flag{
			Name:     "price-per-gb-epoch",
			Category: "Deal Proposal",
			Usage:    "Price in FIL per GiB per epoch",
			Value:    0,
		},
		&cli.Float64Flag{
			Name:     "price-per-gb",
			Category: "Deal Proposal",
			Usage:    "Price in FIL per GiB",
			Value:    0,
		},
		&cli.Float64Flag{
			Name:     "price-per-deal",
			Category: "Deal Proposal",
			Usage:    "Price in FIL per deal",
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
			Category: "Boost Only",
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
			Name:        "schedule-cron",
			Category:    "Scheduling",
			Aliases:     []string{"cron"},
			Usage:       "Cron schedule to send out batch deals",
			DefaultText: "disabled",
		},
		&cli.BoolFlag{
			Name:     "schedule-cron-perpetual",
			Category: "Scheduling",
			Usage:    "Whether the cron schedule runs indefinitely",
		},
		&cli.StringFlag{
			Name:        "start-delay",
			Category:    "Deal Proposal",
			Aliases:     []string{"s"},
			Usage:       "Deal start delay in epoch or duration format, i.e. 1000, 72h",
			Value:       "72h",
			DefaultText: "72h[3 days]",
		},
		&cli.StringFlag{
			Name:        "duration",
			Category:    "Deal Proposal",
			Aliases:     []string{"d"},
			Usage:       "Duration in epoch or duration format, i.e. 1500000, 2400h",
			Value:       "12840h",
			DefaultText: "12840h[535 days]",
		},
		&cli.IntFlag{
			Name:        "schedule-deal-number",
			Category:    "Scheduling",
			Aliases:     []string{"number"},
			Usage:       "Max deal number per triggered schedule",
			DefaultText: "Unlimited",
		},
		&cli.StringFlag{
			Name:        "schedule-deal-size",
			Category:    "Scheduling",
			Aliases:     []string{"size"},
			Usage:       "Max deal sizes per triggered schedule, i.e. 500GiB",
			DefaultText: "Unlimited",
			Value:       "0",
		},
		&cli.StringFlag{
			Name:        "max-pending-deal-size",
			Category:    "Restrictions",
			Aliases:     []string{"pending-size"},
			Usage:       "Max pending deal sizes overall, i.e. 100TiB",
			DefaultText: "Unlimited",
			Value:       "0",
		},
		&cli.IntFlag{
			Name:        "max-pending-deal-number",
			Category:    "Restrictions",
			Aliases:     []string{"pending-number"},
			Usage:       "Max pending deal number overall",
			DefaultText: "Unlimited",
		},
		&cli.BoolFlag{
			Name:     "force",
			Category: "Restrictions",
			Usage:    "Force to send out deals regardless of replication restriction",
		},
	}
}
