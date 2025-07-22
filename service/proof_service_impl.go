package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/filecoin-project/go-address"
	"gorm.io/gorm"
)

// ProofServiceImpl implements the ProofService interface
type ProofServiceImpl struct {
	db *gorm.DB
	mu sync.Mutex // Protects concurrent operations on the same proof
}

// NewProofService creates a new instance of ProofServiceImpl
func NewProofService(db *gorm.DB) ProofService {
	return &ProofServiceImpl{db: db}
}

func (s *ProofServiceImpl) StoreProof(ctx context.Context, proof *model.DealProof) error {
	if err := proof.ValidateClientAddress(); err != nil {
		return fmt.Errorf("invalid client address: %w", err)
	}
	if err := proof.ValidateProviderAddress(); err != nil {
		return fmt.Errorf("invalid provider address: %w", err)
	}
	if err := proof.ValidatePieceCID(); err != nil {
		return fmt.Errorf("invalid piece CID: %w", err)
	}

	proof.ProofStatus = model.ProofStatusPending
	proof.CreatedAt = time.Now()

	return s.db.WithContext(ctx).Create(proof).Error
}

func (s *ProofServiceImpl) VerifyProof(ctx context.Context, dealID uint64, verifier address.Address) (*model.ProofVerification, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var proof model.DealProof
	if err := s.db.WithContext(ctx).First(&proof, dealID).Error; err != nil {
		return nil, fmt.Errorf("proof not found: %w", err)
	}

	// Here we would typically call the actual proof verification logic
	// For now, we'll just create a verification record
	now := time.Now()
	verification := &model.ProofVerification{
		DealID:            dealID,
		VerifiedBy:        verifier.String(),
		VerificationResult: true, // This would be the actual verification result
		VerificationTime:   now,
	}

	if err := s.db.WithContext(ctx).Create(verification).Error; err != nil {
		return nil, fmt.Errorf("failed to create verification record: %w", err)
	}

	// Update proof status
	proof.ProofStatus = model.ProofStatusVerified
	proof.VerificationTime = &now
	if err := s.db.WithContext(ctx).Save(&proof).Error; err != nil {
		return nil, fmt.Errorf("failed to update proof status: %w", err)
	}

	return verification, nil
}

func (s *ProofServiceImpl) GetProofByDealID(ctx context.Context, dealID uint64) (*model.DealProof, error) {
	var proof model.DealProof
	if err := s.db.WithContext(ctx).First(&proof, dealID).Error; err != nil {
		return nil, fmt.Errorf("proof not found: %w", err)
	}
	return &proof, nil
}

func (s *ProofServiceImpl) ListClientProofs(ctx context.Context, clientAddr address.Address) ([]*model.DealProof, error) {
	var proofs []*model.DealProof
	if err := s.db.WithContext(ctx).Where("client_address = ?", clientAddr.String()).Find(&proofs).Error; err != nil {
		return nil, fmt.Errorf("failed to list client proofs: %w", err)
	}
	return proofs, nil
}

func (s *ProofServiceImpl) ListProviderProofs(ctx context.Context, providerAddr address.Address) ([]*model.DealProof, error) {
	var proofs []*model.DealProof
	if err := s.db.WithContext(ctx).Where("provider_address = ?", providerAddr.String()).Find(&proofs).Error; err != nil {
		return nil, fmt.Errorf("failed to list provider proofs: %w", err)
	}
	return proofs, nil
}

func (s *ProofServiceImpl) GetProofVerificationHistory(ctx context.Context, dealID uint64) ([]*model.ProofVerification, error) {
	var verifications []*model.ProofVerification
	if err := s.db.WithContext(ctx).Where("deal_id = ?", dealID).Order("verification_time desc").Find(&verifications).Error; err != nil {
		return nil, fmt.Errorf("failed to get verification history: %w", err)
	}
	return verifications, nil
}
