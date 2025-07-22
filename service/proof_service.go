package service

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/filecoin-project/go-address"
)

// ProofService defines the interface for proof-related operations
type ProofService interface {
	// StoreProof stores a new proof for a deal
	StoreProof(ctx context.Context, proof *model.DealProof) error

	// VerifyProof verifies a stored proof
	VerifyProof(ctx context.Context, dealID uint64, verifier address.Address) (*model.ProofVerification, error)

	// GetProofByDealID retrieves a proof by deal ID
	GetProofByDealID(ctx context.Context, dealID uint64) (*model.DealProof, error)

	// ListClientProofs lists all proofs for a client
	ListClientProofs(ctx context.Context, clientAddr address.Address) ([]*model.DealProof, error)

	// ListProviderProofs lists all proofs for a provider
	ListProviderProofs(ctx context.Context, providerAddr address.Address) ([]*model.DealProof, error)

	// GetProofVerificationHistory gets the verification history for a proof
	GetProofVerificationHistory(ctx context.Context, dealID uint64) ([]*model.ProofVerification, error)
}
