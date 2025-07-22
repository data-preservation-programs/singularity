package handler

import (
	"encoding/base64"
	"fmt"

	"github.com/data-preservation-programs/singularity/model"
)

// StoreProofRequest represents the request to store a new proof
type StoreProofRequest struct {
	DealID          uint64 `json:"dealId" binding:"required"`
	PieceCID        string `json:"pieceCid" binding:"required"`
	ProofBytes      string `json:"proofBytes" binding:"required"` // base64 encoded
	SectorID        uint64 `json:"sectorId" binding:"required"`
	ClientAddress   string `json:"clientAddress" binding:"required"`
	ProviderAddress string `json:"providerAddress" binding:"required"`
}

// ToModel converts the request to a DealProof model
func (r *StoreProofRequest) ToModel() (*model.DealProof, error) {
	proofBytes, err := base64.StdEncoding.DecodeString(r.ProofBytes)
	if err != nil {
		return nil, fmt.Errorf("invalid proof bytes encoding: %w", err)
	}

	return &model.DealProof{
		DealID:          r.DealID,
		PieceCID:        r.PieceCID,
		ProofBytes:      proofBytes,
		SectorID:        r.SectorID,
		ClientAddress:   r.ClientAddress,
		ProviderAddress: r.ProviderAddress,
	}, nil
}

// VerifyProofRequest represents the request to verify a proof
type VerifyProofRequest struct {
	VerifierAddress string `json:"verifierAddress" binding:"required"`
}

// ProofResponse represents a proof in API responses
type ProofResponse struct {
	DealID            uint64                  `json:"dealId"`
	PieceCID          string                  `json:"pieceCid"`
	SectorID          uint64                  `json:"sectorId"`
	ClientAddress     string                  `json:"clientAddress"`
	ProviderAddress   string                  `json:"providerAddress"`
	ProofStatus       model.ProofStatus       `json:"proofStatus"`
	VerificationTime  *string                 `json:"verificationTime,omitempty"`
	CreatedAt         string                  `json:"createdAt"`
	LastVerification  *VerificationResponse   `json:"lastVerification,omitempty"`
}

// VerificationResponse represents a verification in API responses
type VerificationResponse struct {
	ID                uint64 `json:"id"`
	VerifiedBy        string `json:"verifiedBy"`
	VerificationResult bool   `json:"verificationResult"`
	VerificationTime   string `json:"verificationTime"`
}

// FromModel converts a DealProof model to a ProofResponse
func ProofResponseFromModel(proof *model.DealProof) *ProofResponse {
	resp := &ProofResponse{
		DealID:          proof.DealID,
		PieceCID:        proof.PieceCID,
		SectorID:        proof.SectorID,
		ClientAddress:   proof.ClientAddress,
		ProviderAddress: proof.ProviderAddress,
		ProofStatus:     proof.ProofStatus,
		CreatedAt:       proof.CreatedAt.UTC().Format("2006-01-02T15:04:05Z"),
	}

	if proof.VerificationTime != nil {
		timeStr := proof.VerificationTime.UTC().Format("2006-01-02T15:04:05Z")
		resp.VerificationTime = &timeStr
	}

	return resp
}
