package deal

import (
	"encoding/json"
	"fmt"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/deal"
	"github.com/urfave/cli/v2"
)

var SendManualCmd = &cli.Command{
	Name:      "send-manual",
	Usage:     "Send a manual deal proposal to boost or legacy market",
	ArgsUsage: "CLIENT_ADDRESS PROVIDER_ID PIECE_CID PIECE_SIZE",
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
			Name:     "price",
			Category: "Deal Proposal",
			Aliases:  []string{"p"},
			Usage:    "Price per 32GiB Deal over whole duration in Fil",
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
			Value:    535,
		},
		&cli.StringFlag{
			Name:     "lotus-api",
			Category: "Lotus",
			Usage:    "Lotus RPC API endpoint, only used to get miner info",
			Value:    "https://api.node.glif.io/rpc/v1",
		},
		&cli.StringFlag{
			Name:     "lotus-token",
			Category: "Lotus",
			Usage:    "Lotus RPC API token, only used to get miner info",
			Value:    "",
		},
	},
	Action: func(cctx *cli.Context) error {
		proposal := deal.Proposal{
			HTTPHeaders:    cctx.StringSlice("http-header"),
			URLTemplate:    cctx.String("url-template"),
			Price:          cctx.Float64("price"),
			RootCID:        cctx.String("root-cid"),
			Verified:       cctx.Bool("verified"),
			IPNI:           cctx.Bool("ipni"),
			KeepUnsealed:   cctx.Bool("keep-unsealed"),
			StartDelayDays: cctx.Float64("start-delay"),
			DurationDays:   cctx.Float64("duration"),
			ClientAddress:  cctx.Args().Get(0),
			ProviderID:     cctx.Args().Get(1),
			PieceCID:       cctx.Args().Get(2),
			PieceSize:      cctx.Args().Get(3),
			FileSize:       cctx.Uint64("file-size"),
			LotusAPI:       cctx.String("lotus-api"),
			LotusToken:     cctx.String("lotus-token"),
		}
		db := database.MustOpenFromCLI(cctx)
		proposalID, err2 := deal.SendManualHandler(db, proposal)
		if err2 != nil {
			return err2.CliError()
		}

		if cctx.Bool("json") {
			content, _ := json.Marshal(map[string]string{"proposalId": proposalID})
			fmt.Println(string(content))
			return nil
		} else {
			fmt.Println("Deal proposal sent with proposal ID: ", proposalID)
			return nil
		}
	},
}
