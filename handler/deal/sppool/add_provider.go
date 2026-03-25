package sppool

import (
	"context"
	"slices"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

type AddProviderRequest struct {
	Provider    string                  `json:"provider"    validation:"required"` // Storage Provider address
	Policy      model.ReplicationPolicy `json:"policy"      validation:"required"` // Replication policy, e.g. {"market": 1, "pdp": 1}
	URLTemplate *string                 `json:"urlTemplate,omitempty"`             // Optional per-provider URL template override
}

func (DefaultHandler) AddProviderHandler(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	id uint32,
	request AddProviderRequest,
) (*model.SPPoolProvider, error) {
	db = db.WithContext(ctx)

	// Validate pool exists.
	var pool model.SPPool
	err := db.Preload("Providers").Preload("Preparations").First(&pool, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "sp pool %d not found", id)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Validate policy.
	if len(request.Policy) == 0 {
		return nil, errors.Wrap(handlererror.ErrInvalidParameter, "policy must not be empty")
	}
	for dt, count := range request.Policy {
		if !slices.Contains(model.DealTypes, dt) {
			return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "unknown deal type %q in policy", dt)
		}
		if count < 1 {
			return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "policy count for %q must be at least 1", dt)
		}
	}

	// Resolve provider via Lotus.
	var providerActor string
	err = lotusClient.CallFor(ctx, &providerActor, "Filecoin.StateLookupID", request.Provider, nil)
	if err != nil {
		return nil, errors.Join(handlererror.ErrInvalidParameter, errors.Wrapf(err, "provider %s cannot be resolved", request.Provider))
	}

	provider := model.SPPoolProvider{
		PoolID:      model.SPPoolID(id),
		Provider:    providerActor,
		Policy:      request.Policy,
		URLTemplate: request.URLTemplate,
	}

	if err := database.DoRetry(ctx, func() error {
		return db.Create(&provider).Error
	}); err != nil {
		if util.IsDuplicateKeyError(err) {
			return nil, errors.Wrapf(handlererror.ErrDuplicateRecord, "provider %s already in pool %d", providerActor, id)
		}
		return nil, errors.WithStack(err)
	}

	// Reconcile to create schedules for the new provider.
	pool.Providers = append(pool.Providers, provider)
	if err := reconcile(ctx, db, &pool); err != nil {
		return nil, err
	}

	return &provider, nil
}

// @ID AddSPPoolProvider
// @Summary Add a storage provider to an SP Pool
// @Tags SP Pool
// @Accept json
// @Produce json
// @Param id path int true "SP Pool ID"
// @Param request body AddProviderRequest true "AddProviderRequest"
// @Success 200 {object} model.SPPoolProvider
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Failure 409 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /sp-pool/{id}/provider [post]
func _() {}
