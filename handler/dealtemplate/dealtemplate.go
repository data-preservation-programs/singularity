package dealtemplate

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

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
		Name:                request.Name,
		Description:         request.Description,
		DealPricePerGB:      request.DealPricePerGB,
		DealPricePerGBEpoch: request.DealPricePerGBEpoch,
		DealPricePerDeal:    request.DealPricePerDeal,
		DealDuration:        request.DealDuration,
		DealStartDelay:      request.DealStartDelay,
		DealVerified:        request.DealVerified,
		DealKeepUnsealed:    request.DealKeepUnsealed,
		DealAnnounceToIPNI:  request.DealAnnounceToIPNI,
		DealProvider:        request.DealProvider,
		DealHTTPHeaders:     request.DealHTTPHeaders,
		DealURLTemplate:     request.DealURLTemplate,
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
	Name                *string           `json:"name,omitempty"`
	Description         *string           `json:"description,omitempty"`
	DealPricePerGB      *float64          `json:"dealPricePerGb,omitempty"`
	DealPricePerGBEpoch *float64          `json:"dealPricePerGbEpoch,omitempty"`
	DealPricePerDeal    *float64          `json:"dealPricePerDeal,omitempty"`
	DealDuration        *time.Duration    `json:"dealDuration,omitempty"`
	DealStartDelay      *time.Duration    `json:"dealStartDelay,omitempty"`
	DealVerified        *bool             `json:"dealVerified,omitempty"`
	DealKeepUnsealed    *bool             `json:"dealKeepUnsealed,omitempty"`
	DealAnnounceToIPNI  *bool             `json:"dealAnnounceToIpni,omitempty"`
	DealProvider        *string           `json:"dealProvider,omitempty"`
	DealHTTPHeaders     *model.ConfigMap  `json:"dealHttpHeaders,omitempty"`
	DealURLTemplate     *string           `json:"dealUrlTemplate,omitempty"`
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

// ApplyTemplateToPreparation applies deal template parameters to a preparation
func (h *Handler) ApplyTemplateToPreparation(template *model.DealTemplate, prep *model.Preparation) {
	if template == nil {
		return
	}

	// Only apply template values if the preparation doesn't have values set
	if prep.DealPricePerGB == 0 {
		prep.DealPricePerGB = template.DealPricePerGB
	}
	if prep.DealPricePerGBEpoch == 0 {
		prep.DealPricePerGBEpoch = template.DealPricePerGBEpoch
	}
	if prep.DealPricePerDeal == 0 {
		prep.DealPricePerDeal = template.DealPricePerDeal
	}
	if prep.DealDuration == 0 {
		prep.DealDuration = template.DealDuration
	}
	if prep.DealStartDelay == 0 {
		prep.DealStartDelay = template.DealStartDelay
	}
	if !prep.DealVerified {
		prep.DealVerified = template.DealVerified
	}
	if !prep.DealKeepUnsealed {
		prep.DealKeepUnsealed = template.DealKeepUnsealed
	}
	if !prep.DealAnnounceToIPNI {
		prep.DealAnnounceToIPNI = template.DealAnnounceToIPNI
	}
	if prep.DealProvider == "" {
		prep.DealProvider = template.DealProvider
	}
	if prep.DealURLTemplate == "" {
		prep.DealURLTemplate = template.DealURLTemplate
	}
	if len(prep.DealHTTPHeaders) == 0 && len(template.DealHTTPHeaders) > 0 {
		prep.DealHTTPHeaders = template.DealHTTPHeaders
	}
}