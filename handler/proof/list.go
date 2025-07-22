package proof

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

type ListProofRequest struct {
	DealID    *uint64          `json:"dealId"`    // deal ID filter
	ProofType *model.ProofType `json:"proofType"` // proof type filter
	Provider  *string          `json:"provider"`  // provider filter
	Verified  *bool            `json:"verified"`  // verified status filter
	Limit     int              `json:"limit"`     // limit number of results
	Offset    int              `json:"offset"`    // offset for pagination
}

// ListHandler retrieves a list of proofs from the database based on the specified filtering criteria in ListProofRequest.
//
// The function takes advantage of the conditional nature of the ListProofRequest to construct the final query. It
// filters proofs based on various conditions such as deal ID, proof type, provider, and verification status
// as specified in the request.
//
// The function begins by associating the provided context with the database connection. It then successively builds
// upon a GORM statement by appending where clauses based on the parameters in the request.
//
// Parameters:
//   - ctx:      The context for the operation which provides facilities for timeouts and cancellations.
//   - db:       The database connection for performing CRUD operations related to proofs.
//   - request:  The request object which contains the filtering criteria for the proofs retrieval.
//
// Returns:
//   - A slice of model.Proof objects matching the filtering criteria.
//   - An error indicating any issues that occurred during the database operation.
func (DefaultHandler) ListHandler(ctx context.Context, db *gorm.DB, request ListProofRequest) ([]model.Proof, error) {
	db = db.WithContext(ctx)

	// Set default limit if not provided
	if request.Limit == 0 {
		request.Limit = 100
	}

	statement := db

	if request.DealID != nil {
		statement = statement.Where("deal_id = ?", *request.DealID)
	}

	if request.ProofType != nil {
		statement = statement.Where("proof_type = ?", *request.ProofType)
	}

	if request.Provider != nil {
		statement = statement.Where("provider = ?", *request.Provider)
	}

	if request.Verified != nil {
		statement = statement.Where("verified = ?", *request.Verified)
	}

	var proofs []model.Proof
	err := statement.Preload("Deal").
		Limit(request.Limit).
		Offset(request.Offset).
		Order("created_at DESC").
		Find(&proofs).Error

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return proofs, nil
}

// @ID ListProofs
// @Summary List all proofs
// @Description List all proofs with optional filtering
// @Tags Proof
// @Accept json
// @Produce json
// @Param request body ListProofRequest true "ListProofRequest"
// @Success 200 {array} model.Proof
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /proof [get]
func _() {}
