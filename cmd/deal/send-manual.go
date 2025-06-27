package deal

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/deal"
	"github.com/data-preservation-programs/singularity/replication"
	"github.com/data-preservation-programs/singularity/service/epochutil"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/urfave/cli/v2"
)

var SendManualCmd = &cli.Command{
	Name:  "send-manual",
	Usage: "Send a manual deal proposal to boost or legacy market",
	Description: `Send a manual deal proposal to boost or legacy market
  Example: singularity deal send-manual --client f01234 --provider f05678 --piece-cid bagaxxxx --piece-size 32GiB
Notes:
  * The client address must have been imported to the wallet using 'singularity wallet import'
  * The deal proposal will not be saved in the database however will eventually be tracked if the deal tracker is running
  * There is a quick address verification using GLIF API which can be made faster by setting LOTUS_API and LOTUS_TOKEN to your own lotus node`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "save",
			Usage: "Whether to save the deal proposal to the database for tracking purpose",
		},
		&cli.StringFlag{
			Name:     "client",
			Category: "Deal Proposal",
			Usage:    "Client address to send deal from",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "provider",
			Category: "Deal Proposal",
			Usage:    "Storage Provider ID to send deal to",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "piece-cid",
			Category: "Deal Proposal",
			Usage:    "Piece CID of the deal",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "piece-size",
			Category: "Deal Proposal",
			Usage:    "Piece Size of the deal",
			Value:    "32GiB",
			Required: true,
		},
		&cli.StringSliceFlag{
			Name:     "http-header",
			Category: "Boost Only",
			Usage:    "http headers to be passed with the request (i.e. key=value)",
		},
		&cli.StringFlag{
			Name:     "http-url",
			Category: "Boost Only",
			Usage:    "URL or URL template with PIECE_CID placeholder for boost to fetch the CAR file, e.g. http://127.0.0.1/piece/{PIECE_CID}.car",
			Aliases:  []string{"url-template"},
			Value:    "",
		},
		&cli.Uint64Flag{
			Name:     "file-size",
			Category: "Boost Only",
			Usage:    "File size in bytes for boost to fetch the CAR file",
			Value:    0,
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
			Usage:    "Price in FIL  per GiB",
			Value:    0,
		},
		&cli.Float64Flag{
			Name:     "price-per-deal",
			Category: "Deal Proposal",
			Usage:    "Price in FIL per deal",
			Value:    0,
		},
		&cli.StringFlag{
			Name:        "root-cid",
			Category:    "Deal Proposal",
			Usage:       "Root CID that is required as part of the deal proposal, if empty, will be set to empty CID",
			Value:       "bafkqaaa",
			DefaultText: "Empty CID",
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
		&cli.DurationFlag{
			Name:        "timeout",
			Usage:       "Timeout for the deal proposal",
			Value:       time.Minute,
			DefaultText: "1m",
		},
	},
	Action: func(c *cli.Context) error {
		lotusAPI := c.String("lotus-api")
		lotusToken := c.String("lotus-token")
		err := epochutil.Initialize(c.Context, lotusAPI, lotusToken)
		if err != nil {
			return errors.WithStack(err)
		}
		proposal := deal.Proposal{
			HTTPHeaders:     c.StringSlice("http-header"),
			URLTemplate:     c.String("url-template"),
			PricePerGBEpoch: c.Float64("price-per-gb-epoch"),
			PricePerGB:      c.Float64("price-per-gb"),
			PricePerDeal:    c.Float64("price-per-deal"),
			RootCID:         c.String("root-cid"),
			Verified:        c.Bool("verified"),
			IPNI:            c.Bool("ipni"),
			KeepUnsealed:    c.Bool("keep-unsealed"),
			StartDelay:      c.String("start-delay"),
			Duration:        c.String("duration"),
			ClientAddress:   c.String("client"),
			ProviderID:      c.String("provider"),
			PieceCID:        c.String("piece-cid"),
			PieceSize:       c.String("piece-size"),
			FileSize:        c.Uint64("file-size"),
		}
		timeout := c.Duration("timeout")
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()

		ctx, cancel := context.WithTimeout(c.Context, timeout)
		defer cancel()
		h, err := util.InitHost(nil)
		if err != nil {
			return errors.Wrap(err, "failed to init host")
		}
		defer func() { _ = h.Close() }()
		dealMaker := replication.NewDealMaker(
			util.NewLotusClient(c.String("lotus-api"), c.String("lotus-token")),
			h,
			10*timeout,
			timeout,
		)
		dealModel, err := deal.Default.SendManualHandler(ctx, db, dealMaker, proposal)
		if err != nil {
			return errors.WithStack(err)
		}
		if c.Bool("save") {
			db = db.WithContext(ctx)
			err = database.DoRetry(ctx, func() error {
				return db.Create(dealModel).Error
			})
			if err != nil {
				return errors.WithStack(err)
			}
		}
		cliutil.Print(c, dealModel)
		return nil
	},
}
