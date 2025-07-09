package dealtemplate

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

var logger = log.Logger("dealtemplate")

type Handler struct{}

var Default = &Handler{}

// CreateRequest represents the request to create a deal template
type CreateRequest struct {
	Name                string            `json:"name"`
	Description         string            `json:"description"`
	DealPricePerGB      float64           `json:"dealPricePerGb"`
	DealPricePerGBEpoch float64           `json:"dealPricePerGbEpoch"`
	DealPricePerDeal    float64           `json:"dealPricePerDeal"`
	DealDuration        time.Duration     `json:"dealDuration"`
	DealStartDelay      time.Duration     `json:"dealStartDelay"`
	DealVerified        bool              `json:"dealVerified"`
	DealKeepUnsealed    bool              `json:"dealKeepUnsealed"`
	DealAnnounceToIPNI  bool              `json:"dealAnnounceToIpni"`
	DealProvider        string            `json:"dealProvider"`
	DealHTTPHeaders     model.ConfigMap   `json:"dealHttpHeaders"`
	DealURLTemplate     string            `json:"dealUrlTemplate"`
	DealNotes           string            `json:"dealNotes"`
	DealForce           bool              `json:"dealForce"`
	DealAllowedPieceCIDs model.StringSlice `json:"dealAllowedPieceCids"`
}

// CreateHandler creates a new deal template
func (h *Handler) CreateHandler(ctx context.Context, db *gorm.DB, request CreateRequest) (*model.DealTemplate, error) {
	db = db.WithContext(ctx)

	// Check if template with the same name already exists
	var existing model.DealTemplate
	err := db.Where("name = ?", request.Name).First(&existing).Error
	if err == nil {
		return nil, errors.Newf("deal template with name %s already exists", request.Name)
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.WithStack(err)
	}

	template := model.DealTemplate{
		Name:        request.Name,
		Description: request.Description,
		DealConfig: model.DealConfig{
			AutoCreateDeals:      true, // Templates are for auto-creation
			DealPricePerGb:       request.DealPricePerGB,
			DealPricePerGbEpoch:  request.DealPricePerGBEpoch,
			DealPricePerDeal:     request.DealPricePerDeal,
			DealDuration:         request.DealDuration,
			DealStartDelay:       request.DealStartDelay,
			DealVerified:         request.DealVerified,
			DealKeepUnsealed:     request.DealKeepUnsealed,
			DealAnnounceToIpni:   request.DealAnnounceToIPNI,
			DealProvider:         request.DealProvider,
			DealHTTPHeaders:      request.DealHTTPHeaders,
			DealURLTemplate:      request.DealURLTemplate,
			DealNotes:            request.DealNotes,
			DealForce:            request.DealForce,
			DealAllowedPieceCIDs: request.DealAllowedPieceCIDs,
		},
	}

	err = db.Create(&template).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &template, nil
}

// ListHandler lists all deal templates
func (h *Handler) ListHandler(ctx context.Context, db *gorm.DB) ([]model.DealTemplate, error) {
	db = db.WithContext(ctx)

	var templates []model.DealTemplate
	err := db.Find(&templates).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return templates, nil
}

// GetHandler gets a deal template by ID or name
func (h *Handler) GetHandler(ctx context.Context, db *gorm.DB, idOrName string) (*model.DealTemplate, error) {
	db = db.WithContext(ctx)

	var template model.DealTemplate
	err := template.FindByIDOrName(db, idOrName)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &template, nil
}

// UpdateRequest represents the request to update a deal template
type UpdateRequest struct {
	Name                *string            `json:"name,omitempty"`
	Description         *string            `json:"description,omitempty"`
	DealPricePerGB      *float64           `json:"dealPricePerGb,omitempty"`
	DealPricePerGBEpoch *float64           `json:"dealPricePerGbEpoch,omitempty"`
	DealPricePerDeal    *float64           `json:"dealPricePerDeal,omitempty"`
	DealDuration        *time.Duration     `json:"dealDuration,omitempty"`
	DealStartDelay      *time.Duration     `json:"dealStartDelay,omitempty"`
	DealVerified        *bool              `json:"dealVerified,omitempty"`
	DealKeepUnsealed    *bool              `json:"dealKeepUnsealed,omitempty"`
	DealAnnounceToIPNI  *bool              `json:"dealAnnounceToIpni,omitempty"`
	DealProvider        *string            `json:"dealProvider,omitempty"`
	DealHTTPHeaders     *model.ConfigMap   `json:"dealHttpHeaders,omitempty"`
	DealURLTemplate     *string            `json:"dealUrlTemplate,omitempty"`
	DealNotes           *string            `json:"dealNotes,omitempty"`
	DealForce           *bool              `json:"dealForce,omitempty"`
	DealAllowedPieceCIDs *model.StringSlice `json:"dealAllowedPieceCids,omitempty"`
}

// UpdateHandler updates a deal template
func (h *Handler) UpdateHandler(ctx context.Context, db *gorm.DB, idOrName string, request UpdateRequest) (*model.DealTemplate, error) {
	db = db.WithContext(ctx)

	var template model.DealTemplate
	err := template.FindByIDOrName(db, idOrName)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Update only provided fields
	updates := make(map[string]interface{})
	if request.Name != nil {
		updates["name"] = *request.Name
	}
	if request.Description != nil {
		updates["description"] = *request.Description
	}
	if request.DealPricePerGB != nil {
		updates["deal_price_per_gb"] = *request.DealPricePerGB
	}
	if request.DealPricePerGBEpoch != nil {
		updates["deal_price_per_gb_epoch"] = *request.DealPricePerGBEpoch
	}
	if request.DealPricePerDeal != nil {
		updates["deal_price_per_deal"] = *request.DealPricePerDeal
	}
	if request.DealDuration != nil {
		updates["deal_duration"] = *request.DealDuration
	}
	if request.DealStartDelay != nil {
		updates["deal_start_delay"] = *request.DealStartDelay
	}
	if request.DealVerified != nil {
		updates["deal_verified"] = *request.DealVerified
	}
	if request.DealKeepUnsealed != nil {
		updates["deal_keep_unsealed"] = *request.DealKeepUnsealed
	}
	if request.DealAnnounceToIPNI != nil {
		updates["deal_announce_to_ipni"] = *request.DealAnnounceToIPNI
	}
	if request.DealProvider != nil {
		updates["deal_provider"] = *request.DealProvider
	}
	if request.DealHTTPHeaders != nil {
		updates["deal_http_headers"] = *request.DealHTTPHeaders
	}
	if request.DealURLTemplate != nil {
		updates["deal_url_template"] = *request.DealURLTemplate
	}
	if request.DealNotes != nil {
		updates["deal_notes"] = *request.DealNotes
	}
	if request.DealForce != nil {
		updates["deal_force"] = *request.DealForce
	}
	if request.DealAllowedPieceCIDs != nil {
		updates["deal_allowed_piece_cids"] = *request.DealAllowedPieceCIDs
	}

	if len(updates) == 0 {
		return &template, nil
	}

	err = db.Model(&template).Updates(updates).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Reload the template to get updated values
	err = template.FindByIDOrName(db, idOrName)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &template, nil
}

// DeleteHandler deletes a deal template
func (h *Handler) DeleteHandler(ctx context.Context, db *gorm.DB, idOrName string) error {
	db = db.WithContext(ctx)

	var template model.DealTemplate
	err := template.FindByIDOrName(db, idOrName)
	if err != nil {
		return errors.WithStack(err)
	}

	err = db.Delete(&template).Error
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// ApplyTemplateToPreparation applies deal template parameters to a preparation.
// Preparation fields take precedence. Template values are only applied to fields that are unset
// (i.e. zero-value: 0, false, "", or nil). This ensures user-specified values are not overridden.
func (h *Handler) ApplyTemplateToPreparation(template *model.DealTemplate, prep *model.Preparation) {
	if template == nil {
		logger.Debug("No template provided, skipping template application")
		return
	}

	logger.Debugf("Applying deal template %s to preparation %s", template.Name, prep.Name)

	// Use the DealConfig ApplyOverrides method for clean and consistent override logic
	prep.DealConfig.ApplyOverrides(&template.DealConfig)

	logger.Debugf("Applied template %s to preparation %s - template values applied for unset fields only",
		template.Name, prep.Name)
}
