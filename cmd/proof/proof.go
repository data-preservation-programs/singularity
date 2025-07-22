package cmd

import (
	"github.com/data-preservation-programs/singularity/model"
	"fmt"
	"io/ioutil"

	"github.com/data-preservation-programs/singularity/service"
	"github.com/filecoin-project/go-address"
	"github.com/urfave/cli/v2"
)

var ProofCmd = &cli.Command{
	Name:        "proof",
	Usage:       "Manage and verify storage proofs",
	Subcommands: []*cli.Command{
		storeProofCmd,
		verifyProofCmd,
		listProofsCmd,
		getProofCmd,
		proofHistoryCmd,
	},
}

var storeProofCmd = &cli.Command{
	Name:  "store",
	Usage: "Store a new proof for a deal",
	Flags: []cli.Flag{
		&cli.Uint64Flag{
			Name:     "deal-id",
			Usage:    "Deal ID to store proof for",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "piece-cid",
			Usage:    "Piece CID",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "proof-file",
			Usage:    "Path to the proof file",
			Required: true,
		},
		&cli.Uint64Flag{
			Name:     "sector-id",
			Usage:    "Sector ID",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "client",
			Usage:    "Client address",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "provider",
			Usage:    "Provider address",
			Required: true,
		},
	},
	Action: func(cctx *cli.Context) error {
		// Read proof file
		proofBytes, err := ioutil.ReadFile(cctx.String("proof-file"))
		if err != nil {
			return fmt.Errorf("reading proof file: %w", err)
		}

		ctx := cctx.Context
		if db == nil {
			return fmt.Errorf("database not initialized")
		}
		proofService := service.NewProofService(db)

		proof := &model.DealProof{
			DealID:          cctx.Uint64("deal-id"),
			PieceCID:        cctx.String("piece-cid"),
			ProofBytes:      proofBytes,
			SectorID:        cctx.Uint64("sector-id"),
			ClientAddress:   cctx.String("client"),
			ProviderAddress: cctx.String("provider"),
		}

		if err := proofService.StoreProof(ctx, proof); err != nil {
			return fmt.Errorf("storing proof: %w", err)
		}

		fmt.Printf("Successfully stored proof for deal %d\n", proof.DealID)
		return nil
	},
}

var verifyProofCmd = &cli.Command{
	Name:  "verify",
	Usage: "Verify a stored proof",
	Flags: []cli.Flag{
		&cli.Uint64Flag{
			Name:     "deal-id",
			Usage:    "Deal ID of the proof to verify",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "verifier",
			Usage:    "Verifier address",
			Required: true,
		},
	},
	Action: func(cctx *cli.Context) error {
		ctx := cctx.Context
		if db == nil {
			return fmt.Errorf("database not initialized")
		}
		proofService := service.NewProofService(db)

		verifierAddr, err := address.NewFromString(cctx.String("verifier"))
		if err != nil {
			return fmt.Errorf("invalid verifier address: %w", err)
		}

		verification, err := proofService.VerifyProof(ctx, cctx.Uint64("deal-id"), verifierAddr)
		if err != nil {
			return fmt.Errorf("verifying proof: %w", err)
		}

		fmt.Printf("Verification Result: %v\nVerified By: %s\nVerification Time: %s\n",
			verification.VerificationResult,
			verification.VerifiedBy,
			verification.VerificationTime.Format("2006-01-02 15:04:05"),
		)
		return nil
	},
}

var listProofsCmd = &cli.Command{
	Name:  "list",
	Usage: "List proofs by client or provider",
	Subcommands: []*cli.Command{
		{
			Name:  "client",
			Usage: "List proofs by client address",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "address",
					Usage:    "Client address",
					Required: true,
				},
			},
			Action: func(cctx *cli.Context) error {
				ctx := cctx.Context
				proofService := service.NewProofService(db)

				clientAddr, err := address.NewFromString(cctx.String("address"))
				if err != nil {
					return fmt.Errorf("invalid client address: %w", err)
				}

				proofs, err := proofService.ListClientProofs(ctx, clientAddr)
				if err != nil {
					return fmt.Errorf("listing proofs: %w", err)
				}

				for _, proof := range proofs {
					fmt.Printf("Deal ID: %d\nPiece CID: %s\nStatus: %s\nCreated: %s\n\n",
						proof.DealID,
						proof.PieceCID,
						proof.ProofStatus,
						proof.CreatedAt.Format("2006-01-02 15:04:05"),
					)
				}
				return nil
			},
		},
		{
			Name:  "provider",
			Usage: "List proofs by provider address",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "address",
					Usage:    "Provider address",
					Required: true,
				},
			},
			Action: func(cctx *cli.Context) error {
				ctx := cctx.Context
				proofService := service.NewProofService(db)

				providerAddr, err := address.NewFromString(cctx.String("address"))
				if err != nil {
					return fmt.Errorf("invalid provider address: %w", err)
				}

				proofs, err := proofService.ListProviderProofs(ctx, providerAddr)
				if err != nil {
					return fmt.Errorf("listing proofs: %w", err)
				}

				for _, proof := range proofs {
					fmt.Printf("Deal ID: %d\nPiece CID: %s\nStatus: %s\nCreated: %s\n\n",
						proof.DealID,
						proof.PieceCID,
						proof.ProofStatus,
						proof.CreatedAt.Format("2006-01-02 15:04:05"),
					)
				}
				return nil
			},
		},
	},
}

var getProofCmd = &cli.Command{
	Name:  "get",
	Usage: "Get proof details by deal ID",
	Flags: []cli.Flag{
		&cli.Uint64Flag{
			Name:     "deal-id",
			Usage:    "Deal ID",
			Required: true,
		},
	},
	Action: func(cctx *cli.Context) error {
		ctx := cctx.Context
		proofService := service.NewProofService(db)

		proof, err := proofService.GetProofByDealID(ctx, cctx.Uint64("deal-id"))
		if err != nil {
			return fmt.Errorf("getting proof: %w", err)
		}

		fmt.Printf("Deal ID: %d\nPiece CID: %s\nSector ID: %d\nClient: %s\nProvider: %s\nStatus: %s\nCreated: %s\n",
			proof.DealID,
			proof.PieceCID,
			proof.SectorID,
			proof.ClientAddress,
			proof.ProviderAddress,
			proof.ProofStatus,
			proof.CreatedAt.Format("2006-01-02 15:04:05"),
		)
		return nil
	},
}

var proofHistoryCmd = &cli.Command{
	Name:  "history",
	Usage: "Get verification history for a proof",
	Flags: []cli.Flag{
		&cli.Uint64Flag{
			Name:     "deal-id",
			Usage:    "Deal ID",
			Required: true,
		},
	},
	Action: func(cctx *cli.Context) error {
		ctx := cctx.Context
		proofService := service.NewProofService(db)

		history, err := proofService.GetProofVerificationHistory(ctx, cctx.Uint64("deal-id"))
		if err != nil {
			return fmt.Errorf("getting verification history: %w", err)
		}

		for _, v := range history {
			fmt.Printf("Verified By: %s\nResult: %v\nTime: %s\n\n",
				v.VerifiedBy,
				v.VerificationResult,
				v.VerificationTime.Format("2006-01-02 15:04:05"),
			)
		}
		return nil
	},
}
