package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/filecoin-project/go-address"
	"github.com/urfave/cli/v2"
)

func StoreProof(c *cli.Context) error {
	db, closer, err := database.OpenFromCLI(c)
	if err != nil {
		return err
	}
	defer closer.Close()

	proofService := NewProofService(db)
	proofBytes, err := base64.StdEncoding.DecodeString(c.String("proof-bytes"))
	if err != nil {
		return fmt.Errorf("failed to decode proof bytes: %w", err)
	}

	proof := &model.DealProof{
		DealID:          uint64(c.Int64("deal-id")),
		PieceCID:        c.String("piece-cid"),
		SectorID:        uint64(c.Int64("sector-id")),
		ProofBytes:      proofBytes,
		ProofStatus:     "pending",
		ClientAddress:   c.String("client-address"),
		ProviderAddress: c.String("provider-address"),
	}

	if err := proofService.StoreProof(c.Context, proof); err != nil {
		return fmt.Errorf("failed to store proof: %w", err)
	}

	return json.NewEncoder(os.Stdout).Encode(proof)
}

func VerifyProof(c *cli.Context) error {
	db, closer, err := database.OpenFromCLI(c)
	if err != nil {
		return err
	}
	defer closer.Close()

	proofService := NewProofService(db)
	dealID := uint64(c.Int64("deal-id"))
	verifier, err := address.NewFromString(c.String("verifier"))
	if err != nil {
		return fmt.Errorf("invalid verifier address: %w", err)
	}

	verification, err := proofService.VerifyProof(c.Context, dealID, verifier)
	if err != nil {
		return fmt.Errorf("failed to verify proof: %w", err)
	}

	return json.NewEncoder(os.Stdout).Encode(verification)
}

func GetProof(c *cli.Context) error {
	db, closer, err := database.OpenFromCLI(c)
	if err != nil {
		return err
	}
	defer closer.Close()

	proofService := NewProofService(db)
	dealID := uint64(c.Int64("deal-id"))

	proof, err := proofService.GetProofByDealID(c.Context, dealID)
	if err != nil {
		return fmt.Errorf("failed to get proof: %w", err)
	}

	return json.NewEncoder(os.Stdout).Encode(proof)
}

func ListProofs(c *cli.Context) error {
	db, closer, err := database.OpenFromCLI(c)
	if err != nil {
		return err
	}
	defer closer.Close()

	proofService := NewProofService(db)

	// List all proofs
	proofs := []*model.DealProof{}
	
	// If client address is provided, filter by client
	if c.String("client") != "" {
		clientAddr, err := address.NewFromString(c.String("client"))
		if err != nil {
			return fmt.Errorf("invalid client address: %w", err)
		}
		proofs, err = proofService.ListClientProofs(c.Context, clientAddr)
		if err != nil {
			return fmt.Errorf("failed to list client proofs: %w", err)
		}
	} else if c.String("provider") != "" {
		// If provider address is provided, filter by provider
		providerAddr, err := address.NewFromString(c.String("provider"))
		if err != nil {
			return fmt.Errorf("invalid provider address: %w", err)
		}
		proofs, err = proofService.ListProviderProofs(c.Context, providerAddr)
		if err != nil {
			return fmt.Errorf("failed to list provider proofs: %w", err)
		}
	} else {
		return fmt.Errorf("must provide either --client or --provider flag")
	}

	return json.NewEncoder(os.Stdout).Encode(proofs)
}

func GetProofHistory(c *cli.Context) error {
	db, closer, err := database.OpenFromCLI(c)
	if err != nil {
		return err
	}
	defer closer.Close()

	proofService := NewProofService(db)
	dealID := uint64(c.Int64("deal-id"))

	history, err := proofService.GetProofVerificationHistory(c.Context, dealID)
	if err != nil {
		return fmt.Errorf("failed to get proof history: %w", err)
	}

	return json.NewEncoder(os.Stdout).Encode(history)
}
