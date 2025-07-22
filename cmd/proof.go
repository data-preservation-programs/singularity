package cmd

import (
	"github.com/data-preservation-programs/singularity/service"
	"github.com/urfave/cli/v2"
)

var ProofCmd = &cli.Command{
	Name:     "proof",
	Usage:    "Manage sector deal proofs",
	Category: "Operations",
	Subcommands: []*cli.Command{
		{
			Name:  "store",
			Usage: "Store a deal proof for a specific deal ID",
			Flags: []cli.Flag{
				&cli.Int64Flag{
					Name:     "deal-id",
					Usage:    "The ID of the deal to store proof for",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "piece-cid",
					Usage:    "The CID of the piece",
					Required: true,
				},
				&cli.Int64Flag{
					Name:     "sector-id",
					Usage:    "The ID of the sector containing the deal",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "proof-bytes",
					Usage:    "Base64 encoded proof bytes",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "client-address",
					Usage:    "The client's Filecoin address",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "provider-address",
					Usage:    "The storage provider's Filecoin address",
					Required: true,
				},
			},
			Action: service.StoreProof,
		},
		{
			Name:  "verify",
			Usage: "Verify a stored deal proof",
			Flags: []cli.Flag{
				&cli.Int64Flag{
					Name:     "deal-id",
					Usage:    "The ID of the deal to verify proof for",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "verifier",
					Usage:    "The Filecoin address of the verifier",
					Required: true,
				},
			},
			Action: service.VerifyProof,
		},
		{
			Name:  "get",
			Usage: "Get details of a stored deal proof",
			Flags: []cli.Flag{
				&cli.Int64Flag{
					Name:     "deal-id",
					Usage:    "The ID of the deal to get proof for",
					Required: true,
				},
			},
			Action: service.GetProof,
		},
		{
			Name:  "list",
			Usage: "List stored deal proofs",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "client",
					Usage:    "Filter proofs by client address",
				},
				&cli.StringFlag{
					Name:     "provider",
					Usage:    "Filter proofs by provider address",
				},
			},
			Action: service.ListProofs,
		},
		{
			Name:  "history",
			Usage: "Get verification history for a deal proof",
			Flags: []cli.Flag{
				&cli.Int64Flag{
					Name:     "deal-id",
					Usage:    "The ID of the deal to get history for",
					Required: true,
				},
			},
			Action: service.GetProofHistory,
		},
	},
}
