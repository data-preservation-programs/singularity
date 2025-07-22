package model

import (
	"time"

	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
)

type ProofStatus string

const (
	ProofStatusPending  ProofStatus = "pending"
	ProofStatusVerified ProofStatus = "verified"
	ProofStatusFailed   ProofStatus = "failed"
)

// DealProof represents a proof associated with a storage deal
type DealProof struct {
	tableName struct{} `gorm:"table:deal_proofs"` // specify the table name for GORM
	DealID            uint64         `json:"deal_id" gorm:"primaryKey"`
	PieceCID          string         `json:"piece_cid" gorm:"index"`
	ProofBytes        []byte         `json:"proof_bytes"`
	SectorID          uint64         `json:"sector_id"`
	ClientAddress     string         `json:"client_address" gorm:"index"`
	ProviderAddress   string         `json:"provider_address" gorm:"index"`
	ProofStatus       ProofStatus    `json:"proof_status"`
	VerificationTime  *time.Time     `json:"verification_time"`
	CreatedAt         time.Time      `json:"created_at"`
	Verifications     []ProofVerification `json:"verifications" gorm:"foreignKey:DealID"`
}

// ProofVerification represents a verification attempt of a proof
type ProofVerification struct {
	tableName struct{} `gorm:"table:proof_verifications"` // specify the table name for GORM
	ID                uint64         `json:"id" gorm:"primaryKey"`
	DealID           uint64         `json:"deal_id"`
	VerifiedBy       string         `json:"verified_by"`
	VerificationResult bool          `json:"verification_result"`
	VerificationTime  time.Time      `json:"verification_time"`
}

// ValidateClientAddress validates and normalizes the client address
func (p *DealProof) ValidateClientAddress() error {
	addr, err := address.NewFromString(p.ClientAddress)
	if err != nil {
		return err
	}
	p.ClientAddress = addr.String()
	return nil
}

// ValidateProviderAddress validates and normalizes the provider address
func (p *DealProof) ValidateProviderAddress() error {
	addr, err := address.NewFromString(p.ProviderAddress)
	if err != nil {
		return err
	}
	p.ProviderAddress = addr.String()
	return nil
}

// ValidatePieceCID validates the piece CID
func (p *DealProof) ValidatePieceCID() error {
	_, err := cid.Decode(p.PieceCID)
	return err
}
