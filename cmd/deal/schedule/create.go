package schedule

import (
	"bufio"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/deal/schedule"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"os"
	"regexp"
)

var CreateCmd = &cli.Command{
	Name:      "create",
	Usage:     "Create a schedule to send out deals to a storage provider",
	ArgsUsage: "DATASET_NAME PROVIDER_ID",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:     "http-header",
			Category: "Boost Only",
			Aliases:  []string{"H"},
			Usage:    "http headers to be passed with the request (i.e. key=value)",
		},
		&cli.StringFlag{
			Name:     "url-template",
			Category: "Boost Only",
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
		&cli.DurationFlag{
			Name:        "schedule-interval",
			Category:    "Scheduling",
			Aliases:     []string{"every"},
			Usage:       "Cron schedule to send out batch deals",
			DefaultText: "disabled",
			Value:       0,
		},
		&cli.Float64Flag{
			Name:     "start-delay",
			Category: "Deal Proposal",
			Aliases:  []string{"s"},
			Usage:    "Deal start delay in days",
			Value:    3,
		},
		&cli.Float64Flag{
			Name:     "duration",
			Category: "Deal Proposal",
			Aliases:  []string{"d"},
			Usage:    "Duration in days for deal length",
			Value:    530,
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
			Category:    "Scheduling",
			Aliases:     []string{"total-number"},
			Usage:       "Max total deal number for this request, i.e. 1000",
			DefaultText: "Unlimited",
		},
		&cli.StringFlag{
			Name:        "schedule-deal-size",
			Category:    "Scheduling",
			Aliases:     []string{"size"},
			Usage:       "Max deal sizes per triggered schedule, i.e. 500GB",
			DefaultText: "Unlimited",
			Value:       "0",
		},
		&cli.StringFlag{
			Name:        "total-deal-size",
			Category:    "Restrictions",
			Aliases:     []string{"total-size"},
			Usage:       "Max total deal sizes for this request, i.e. 100TB",
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
			Usage:       "Max pending deal sizes overall for this request",
			DefaultText: "Unlimited",
			Value:       "0",
		},
		&cli.IntFlag{
			Name:        "max-pending-deal-number",
			Category:    "Restrictions",
			Aliases:     []string{"pending-number"},
			Usage:       "Max pending deal number overall for this request",
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
		db := database.MustOpenFromCLI(c)
		cids := map[string]struct{}{}
		for _, f := range c.StringSlice("allowed-piece-cid-file") {
			cidsFromFile, err := readCIDsFromFile(f)
			if err != nil {
				return err
			}
			for _, cid := range cidsFromFile {
				cids[cid] = struct{}{}
			}
		}
		for _, cid := range c.StringSlice("allowed-piece-cid") {
			cids[cid] = struct{}{}
		}
		var allowedPieceCIDs []string
		for cid := range cids {
			allowedPieceCIDs = append(allowedPieceCIDs, cid)
		}
		request := schedule.CreateRequest{
			DatasetName:          c.Args().Get(0),
			Provider:             c.Args().Get(1),
			HTTPHeaders:          c.StringSlice("http-header"),
			URLTemplate:          c.String("url-template"),
			Price:                c.Float64("price"),
			Verified:             c.Bool("verified"),
			IPNI:                 c.Bool("ipni"),
			KeepUnsealed:         c.Bool("keep-unsealed"),
			ScheduleInterval:     c.Duration("schedule-interval"),
			StartDelayDays:       c.Float64("start-delay"),
			DurationDays:         c.Float64("duration"),
			ScheduleDealNumber:   c.Int("schedule-deal-number"),
			TotalDealNumber:      c.Int("total-deal-number"),
			ScheduleDealSize:     c.String("schedule-deal-size"),
			TotalDealSize:        c.String("total-deal-size"),
			Notes:                c.String("notes"),
			MaxPendingDealSize:   c.String("max-pending-deal-size"),
			MaxPendingDealNumber: c.Int("max-pending-deal-number"),
			AllowedPieceCIDs:     allowedPieceCIDs,
		}
		schedule, err := schedule.CreateHandler(db, request)
		if err != nil {
			return err.CliError()
		}

		cliutil.PrintToConsole(schedule, c.Bool("json"))
		return nil
	},
}

func readCIDsFromFile(f string) ([]string, error) {
	var result []string
	file, err := os.Open(f)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open file")
	}
	defer file.Close()
	re := regexp.MustCompile(`\bbaga[a-z2-7]+\b`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			result = append(result, match)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.Wrap(err, "failed to read file")
	}
	return result, nil
}
