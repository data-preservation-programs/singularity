package schedule

import (
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/deal/schedule"
	"github.com/gotidy/ptr"
	"github.com/urfave/cli/v2"
)

var UpdateCmd = &cli.Command{
	Name:      "update",
	ArgsUsage: "<schedule_id>",
	Before:    cliutil.CheckNArgs,
	Usage:     "Update an existing schedule",
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
		&cli.StringSliceFlag{
			Name:     "http-header",
			Category: "Boost Only",
			Aliases:  []string{"H"},
			Usage:    "HTTP headers to be passed with the request (i.e. key=value). This will replace the existing header values. To remove a header, use --http-header \"key=\"\". To remove all headers, use --http-header \"\"",
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
			Name:     "force",
			Category: "Restrictions",
			Usage:    "Force to send out deals regardless of replication restriction",
		},
		&cli.BoolFlag{
			Name:     "keep-unsealed",
			Category: "Deal Proposal",
			Usage:    "Whether to keep unsealed copy",
			Value:    true,
		},
		&cli.StringFlag{
			Name:     "schedule-cron",
			Category: "Scheduling",
			Aliases:  []string{"cron"},
			Usage:    "Cron schedule to send out batch deals",
		},
		&cli.StringFlag{
			Name:     "start-delay",
			Category: "Deal Proposal",
			Aliases:  []string{"s"},
			Usage:    "Deal start delay in epoch or in duration format, i.e. 1000, 72h",
		},
		&cli.StringFlag{
			Name:     "duration",
			Category: "Deal Proposal",
			Aliases:  []string{"d"},
			Usage:    "Duration in epoch or in duration format, i.e. 1500000, 2400h",
		},
		&cli.IntFlag{
			Name:     "schedule-deal-number",
			Category: "Scheduling",
			Aliases:  []string{"number"},
			Usage:    "Max deal number per triggered schedule, i.e. 30",
		},
		&cli.IntFlag{
			Name:     "total-deal-number",
			Category: "Restrictions",
			Aliases:  []string{"total-number"},
			Usage:    "Max total deal number for this request, i.e. 1000",
		},
		&cli.StringFlag{
			Name:     "schedule-deal-size",
			Category: "Scheduling",
			Aliases:  []string{"size"},
			Usage:    "Max deal sizes per triggered schedule, i.e. 500GiB",
		},
		&cli.StringFlag{
			Name:     "total-deal-size",
			Category: "Restrictions",
			Aliases:  []string{"total-size"},
			Usage:    "Max total deal sizes for this request, i.e. 100TiB",
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
			Aliases:  []string{"pending-size"},
			Usage:    "Max pending deal sizes overall for this request, i.e. 1000",
		},
		&cli.IntFlag{
			Name:     "max-pending-deal-number",
			Category: "Restrictions",
			Aliases:  []string{"pending-number"},
			Usage:    "Max pending deal number overall for this request, i.e. 100TiB",
		},
		&cli.StringSliceFlag{
			Name:     "allowed-piece-cid",
			Category: "Restrictions",
			Aliases:  []string{"piece-cid"},
			Usage:    "List of allowed piece CIDs in this schedule. Append only.",
		},
		&cli.StringSliceFlag{
			Name:     "allowed-piece-cid-file",
			Category: "Restrictions",
			Aliases:  []string{"piece-cid-file"},
			Usage:    "List of files that contains a list of piece CIDs to allow. Append only.",
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
		request := schedule.UpdateRequest{
			HTTPHeaders:      c.StringSlice("http-header"),
			AllowedPieceCIDs: allowedPieceCIDs,
		}

		if c.IsSet("url-template") {
			request.URLTemplate = ptr.Of(c.String("url-template"))
		}
		if c.IsSet("price-per-gb-epoch") {
			request.PricePerGBEpoch = ptr.Of(c.Float64("price-per-gb-epoch"))
		}
		if c.IsSet("price-per-gb") {
			request.PricePerGB = ptr.Of(c.Float64("price-per-gb"))
		}
		if c.IsSet("price-per-deal") {
			request.PricePerDeal = ptr.Of(c.Float64("price-per-deal"))
		}
		if c.IsSet("verified") {
			request.Verified = ptr.Of(c.Bool("verified"))
		}
		if c.IsSet("ipni") {
			request.IPNI = ptr.Of(c.Bool("ipni"))
		}
		if c.IsSet("keep-unsealed") {
			request.KeepUnsealed = ptr.Of(c.Bool("keep-unsealed"))
		}
		if c.IsSet("schedule-cron") {
			request.ScheduleCron = ptr.Of(c.String("schedule-cron"))
		}
		if c.IsSet("start-delay") {
			request.StartDelay = ptr.Of(c.String("start-delay"))
		}
		if c.IsSet("duration") {
			request.Duration = ptr.Of(c.String("duration"))
		}
		if c.IsSet("schedule-deal-number") {
			request.ScheduleDealNumber = ptr.Of(c.Int("schedule-deal-number"))
		}
		if c.IsSet("total-deal-number") {
			request.TotalDealNumber = ptr.Of(c.Int("total-deal-number"))
		}
		if c.IsSet("schedule-deal-size") {
			request.ScheduleDealSize = ptr.Of(c.String("schedule-deal-size"))
		}
		if c.IsSet("total-deal-size") {
			request.TotalDealSize = ptr.Of(c.String("total-deal-size"))
		}
		if c.IsSet("notes") {
			request.Notes = ptr.Of(c.String("notes"))
		}
		if c.IsSet("max-pending-deal-size") {
			request.MaxPendingDealSize = ptr.Of(c.String("max-pending-deal-size"))
		}
		if c.IsSet("max-pending-deal-number") {
			request.MaxPendingDealNumber = ptr.Of(c.Int("max-pending-deal-number"))
		}
		if c.IsSet("force") {
			request.Force = ptr.Of(c.Bool("force"))
		}

		id, err := strconv.ParseUint(c.Args().Get(0), 10, 32)
		if err != nil {
			return errors.Newf("invalid schedule id %s", c.Args().Get(0))
		}

		schedule, err := schedule.Default.UpdateHandler(c.Context, db, uint32(id), request)
		if err != nil {
			return errors.WithStack(err)
		}

		cliutil.Print(c, schedule)
		return nil
	},
}
