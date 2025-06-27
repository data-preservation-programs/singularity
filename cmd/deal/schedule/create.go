package schedule

import (
	"bufio"
	"os"
	"regexp"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/deal/schedule"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/urfave/cli/v2"
)

var re = regexp.MustCompile(`\bbaga[a-z2-7]+\b`)

var CreateCmd = &cli.Command{
	Name:  "create",
	Usage: "Create a schedule to send out deals to a storage provider",
	Description: `CRON pattern '--schedule-cron': The CRON pattern can either be a descriptor or a standard CRON pattern with optional second field
  Standard CRON:
    ┌───────────── minute (0 - 59)
    │ ┌───────────── hour (0 - 23)
    │ │ ┌───────────── day of the month (1 - 31)
    │ │ │ ┌───────────── month (1 - 12)
    │ │ │ │ ┌───────────── day of the week (0 - 6) (Sunday to Saturday)
    │ │ │ │ │                                   
    │ │ │ │ │
    │ │ │ │ │
    * * * * *

  Optional Second field:
    ┌─────────────  second (0 - 59)
    │ ┌─────────────  minute (0 - 59)
    │ │ ┌─────────────  hour (0 - 23)
    │ │ │ ┌─────────────  day of the month (1 - 31)
    │ │ │ │ ┌─────────────  month (1 - 12)
    │ │ │ │ │ ┌─────────────  day of the week (0 - 6) (Sunday to Saturday)
    │ │ │ │ │ │
    │ │ │ │ │ │
    * * * * * *

  Descriptor:
    @yearly, @annually - Equivalent to 0 0 1 1 *
    @monthly           - Equivalent to 0 0 1 * *
    @weekly            - Equivalent to 0 0 * * 0
    @daily,  @midnight - Equivalent to 0 0 * * *
    @hourly            - Equivalent to 0 * * * *`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "preparation",
			Usage:    "Preparation ID or name",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "provider",
			Usage:    "Storage Provider ID to send deals to",
			Required: true,
		},
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
			Usage:    "URL template with PIECE_CID placeholder for boost to fetch the CAR file, i.e. http://127.0.0.1/piece/{PIECE_CID}.car",
			Value:    "",
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
			Value:       "",
		},
		&cli.StringFlag{
			Name:        "start-delay",
			Category:    "Deal Proposal",
			Aliases:     []string{"s"},
			Usage:       "Deal start delay in epoch or in duration format, i.e. 1000, 72h",
			Value:       "72h",
			DefaultText: "72h[3 days]",
		},
		&cli.StringFlag{
			Name:        "duration",
			Category:    "Deal Proposal",
			Aliases:     []string{"d"},
			Usage:       "Duration in epoch or in duration format, i.e. 1500000, 2400h",
			Value:       "12840h",
			DefaultText: "12840h[535 days]",
		},
		&cli.IntFlag{
			Name:        "schedule-deal-number",
			Category:    "Scheduling",
			Aliases:     []string{"number"},
			Usage:       "Max deal number per triggered schedule, i.e. 30",
			DefaultText: "Unlimited",
		},
		&cli.IntFlag{
			Name:        "total-deal-number",
			Category:    "Restrictions",
			Aliases:     []string{"total-number"},
			Usage:       "Max total deal number for this request, i.e. 1000",
			DefaultText: "Unlimited",
		},
		&cli.BoolFlag{
			Name:     "force",
			Category: "Restrictions",
			Usage:    "Force to send out deals regardless of replication restriction",
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
			Name:        "total-deal-size",
			Category:    "Restrictions",
			Aliases:     []string{"total-size"},
			Usage:       "Max total deal sizes for this request, i.e. 100TiB",
			DefaultText: "Unlimited",
			Value:       "0",
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
			Aliases:     []string{"pending-size"},
			Usage:       "Max pending deal sizes overall for this request, i.e. 1000",
			DefaultText: "Unlimited",
			Value:       "0",
		},
		&cli.IntFlag{
			Name:        "max-pending-deal-number",
			Category:    "Restrictions",
			Aliases:     []string{"pending-number"},
			Usage:       "Max pending deal number overall for this request, i.e. 100TiB",
			DefaultText: "Unlimited",
		},
		&cli.StringSliceFlag{
			Name:        "allowed-piece-cid",
			Category:    "Restrictions",
			Aliases:     []string{"piece-cid"},
			Usage:       "List of allowed piece CIDs in this schedule",
			DefaultText: "Any",
		},
		&cli.StringSliceFlag{
			Name:     "allowed-piece-cid-file",
			Category: "Restrictions",
			Aliases:  []string{"piece-cid-file"},
			Usage:    "List of files that contains a list of piece CIDs to allow",
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()
		allowedPieceCIDs := c.StringSlice("allowed-piece-cid")
		for _, f := range c.StringSlice("allowed-piece-cid-file") {
			cidsFromFile, err := readCIDsFromFile(f)
			if err != nil {
				return errors.WithStack(err)
			}
			allowedPieceCIDs = append(allowedPieceCIDs, cidsFromFile...)
		}
		request := schedule.CreateRequest{
			Preparation:          c.String("preparation"),
			Provider:             c.String("provider"),
			HTTPHeaders:          c.StringSlice("http-header"),
			URLTemplate:          c.String("url-template"),
			PricePerGBEpoch:      c.Float64("price-per-gb-epoch"),
			PricePerGB:           c.Float64("price-per-gb"),
			PricePerDeal:         c.Float64("price-per-deal"),
			Verified:             c.Bool("verified"),
			IPNI:                 c.Bool("ipni"),
			KeepUnsealed:         c.Bool("keep-unsealed"),
			ScheduleCron:         c.String("schedule-cron"),
			StartDelay:           c.String("start-delay"),
			Duration:             c.String("duration"),
			ScheduleDealNumber:   c.Int("schedule-deal-number"),
			TotalDealNumber:      c.Int("total-deal-number"),
			ScheduleDealSize:     c.String("schedule-deal-size"),
			TotalDealSize:        c.String("total-deal-size"),
			Notes:                c.String("notes"),
			MaxPendingDealSize:   c.String("max-pending-deal-size"),
			MaxPendingDealNumber: c.Int("max-pending-deal-number"),
			AllowedPieceCIDs:     allowedPieceCIDs,
			Force:                c.Bool("force"),
		}
		lotusClient := util.NewLotusClient(c.String("lotus-api"), c.String("lotus-token"))
		schedule, err := schedule.Default.CreateHandler(c.Context, db, lotusClient, request)
		if err != nil {
			return errors.WithStack(err)
		}

		cliutil.Print(c, schedule)
		return nil
	},
}

func readCIDsFromFile(f string) ([]string, error) {
	var result []string
	file, err := os.Open(f)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open file")
	}
	defer func() { _ = file.Close() }()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindAllString(line, -1)
		result = append(result, matches...)
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.Wrap(err, "failed to read file")
	}
	return result, nil
}
