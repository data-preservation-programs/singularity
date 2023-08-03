package deal

import (
	"context"
	"time"

	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/replication"
	"github.com/data-preservation-programs/singularity/service/epochutil"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/pkg/errors"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/deal"
	"github.com/urfave/cli/v2"
)

var SendManualCmd = &cli.Command{
	Name:      "send-manual",
	Usage:     "Send a manual deal proposal to boost or legacy market",
	ArgsUsage: "CLIENT_ADDRESS PROVIDER_ID PIECE_CID PIECE_SIZE",
	Description: `Send a manual deal proposal to boost or legacy market
  Example: singularity deal send-manual f01234 f05678 bagaxxxx 32GiB
Notes:
  * The client address must have been imported to the wallet using 'singularity wallet import'
  * The deal proposal will not be saved in the database however will eventually be tracked if the deal tracker is running
  * There is a quick address verification using GLIF API which can be made faster by setting LOTUS_API and LOTUS_TOKEN to your own lotus node`,
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:     "http-header",
			Category: "Boost Only",
			Usage:    "http headers to be passed with the request (i.e. key=value)",
		},
		&cli.StringFlag{
			Name:     "url-template",
			Category: "Boost Only",
			Usage:    "URL template with PIECE_CID placeholder for boost to fetch the CAR file, i.e. http://127.0.0.1/piece/{PIECE_CID}.car",
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
	Action: func(cctx *cli.Context) error {
		lotusAPI := cctx.String("lotus-api")
		lotusToken := cctx.String("lotus-token")
		err := epochutil.Initialize(cctx.Context, lotusAPI, lotusToken)
		if err != nil {
			return err
		}
		proposal := deal.Proposal{
			HTTPHeaders:     cctx.StringSlice("http-header"),
			URLTemplate:     cctx.String("url-template"),
			PricePerGBEpoch: cctx.Float64("price-per-gb-epoch"),
			PricePerGB:      cctx.Float64("price-per-gb"),
			PricePerDeal:    cctx.Float64("price-per-deal"),
			RootCID:         cctx.String("root-cid"),
			Verified:        cctx.Bool("verified"),
			IPNI:            cctx.Bool("ipni"),
			KeepUnsealed:    cctx.Bool("keep-unsealed"),
			StartDelay:      cctx.String("start-delay"),
			Duration:        cctx.String("duration"),
			ClientAddress:   cctx.Args().Get(0),
			ProviderID:      cctx.Args().Get(1),
			PieceCID:        cctx.Args().Get(2),
			PieceSize:       cctx.Args().Get(3),
			FileSize:        cctx.Uint64("file-size"),
		}
		timeout := cctx.Duration("timeout")
		db, closer, err := database.OpenFromCLI(cctx)
		if err != nil {
			return err
		}
		defer closer.Close()
		ctx, cancel := context.WithTimeout(cctx.Context, timeout)
		defer cancel()
		h, err := util.InitHost(nil)
		if err != nil {
			return errors.Wrap(err, "failed to init host")
		}
		dealMaker := replication.NewDealMaker(
			util.NewLotusClient(cctx.String("lotus-api"), cctx.String("lotus-token")),
			h,
			10*timeout,
			timeout,
		)
		dealModel, err2 := deal.SendManualHandler(db.WithContext(ctx), ctx, dealMaker, proposal)
		if err2 != nil {
			return err2
		}
		cliutil.PrintToConsole(dealModel, cctx.Bool("json"), []string{
			"CreatedAt", "UpdatedAt", "DealID", "DatasetID", "SectorStartEpoch",
			"ID", "State", "ErrorMessage", "ScheduleID"})
		return nil
	},
}
