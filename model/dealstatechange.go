package model

import (
	"time"

	"gorm.io/gorm"
)

// DealStateChange represents a state transition in the deal lifecycle
// This model provides comprehensive tracking of all deal state changes
// for auditing, debugging, and analytics purposes
type DealStateChange struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement"                  json:"id"             table:"verbose"`
	DealID        DealID    `gorm:"index;not null"                           json:"dealId"         table:"verbose"` // Internal singularity deal ID
	PreviousState DealState `gorm:"size:50;index"                            json:"previousState"  table:"verbose"` // Previous deal state (nullable for initial state)
	NewState      DealState `gorm:"size:50;index;not null"                   json:"newState"       table:"verbose"` // New deal state
	Timestamp     time.Time `gorm:"index;not null;default:CURRENT_TIMESTAMP" json:"timestamp"      table:"verbose;format:2006-01-02 15:04:05"`
	EpochHeight   *int32    `gorm:"index"                                    json:"epochHeight"    table:"verbose"` // Filecoin epoch when change occurred
	SectorID      *string   `gorm:"size:100;index"                           json:"sectorId"       table:"verbose"` // Storage provider sector ID
	ProviderID    string    `gorm:"size:20;index"                            json:"providerId"     table:"verbose"` // Storage provider ID
	ClientAddress string    `gorm:"size:86;index"                            json:"clientAddress"  table:"verbose"` // Client wallet address
	Metadata      string    `gorm:"type:TEXT"                                json:"metadata"       table:"verbose"` // Additional metadata as JSON

	// Associations
	Deal *Deal `gorm:"foreignKey:DealID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;association_autoupdate:false;association_autocreate:false" json:"deal,omitempty" swaggerignore:"true" table:"expand"`
}

// DealStateChangeQuery represents query parameters for filtering deal state changes
type DealStateChangeQuery struct {
	DealID        *DealID    `json:"dealId" form:"dealId"`               // Filter by specific deal ID
	State         *DealState `json:"state" form:"state"`                 // Filter by new state
	ProviderID    *string    `json:"providerId" form:"providerId"`       // Filter by provider ID
	ClientAddress *string    `json:"clientAddress" form:"clientAddress"` // Filter by client address
	StartTime     *time.Time `json:"startTime" form:"startTime"`         // Filter changes after this time
	EndTime       *time.Time `json:"endTime" form:"endTime"`             // Filter changes before this time

	// Pagination
	Offset *int `json:"offset" form:"offset"` // Number of records to skip
	Limit  *int `json:"limit" form:"limit"`   // Maximum number of records to return

	// Sorting
	OrderBy *string `json:"orderBy" form:"orderBy"` // Field to sort by (default: timestamp)
	Order   *string `json:"order" form:"order"`     // Sort order: asc/desc (default: desc)
}

// CreateDealStateChange records a new deal state change
func CreateDealStateChange(db *gorm.DB, dealID DealID, previousState *DealState, newState DealState, epochHeight *int32, sectorID *string, providerID string, clientAddress string, metadata string) error {
	change := &DealStateChange{
		DealID:        dealID,
		PreviousState: "",
		NewState:      newState,
		Timestamp:     time.Now(),
		EpochHeight:   epochHeight,
		SectorID:      sectorID,
		ProviderID:    providerID,
		ClientAddress: clientAddress,
		Metadata:      metadata,
	}

	if previousState != nil {
		change.PreviousState = *previousState
	}

	return db.Create(change).Error
}

// GetDealStateChanges retrieves deal state changes with optional filtering
func GetDealStateChanges(db *gorm.DB, query DealStateChangeQuery) ([]DealStateChange, int64, error) {
	var changes []DealStateChange
	var total int64

	// Build the base query
	baseQuery := db.Model(&DealStateChange{})

	// Apply filters
	if query.DealID != nil {
		baseQuery = baseQuery.Where("deal_id = ?", *query.DealID)
	}

	if query.State != nil {
		baseQuery = baseQuery.Where("new_state = ?", *query.State)
	}

	if query.ProviderID != nil {
		baseQuery = baseQuery.Where("provider_id = ?", *query.ProviderID)
	}

	if query.ClientAddress != nil {
		baseQuery = baseQuery.Where("client_address = ?", *query.ClientAddress)
	}

	if query.StartTime != nil {
		baseQuery = baseQuery.Where("timestamp >= ?", *query.StartTime)
	}

	if query.EndTime != nil {
		baseQuery = baseQuery.Where("timestamp <= ?", *query.EndTime)
	}

	// Get total count
	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply sorting
	orderBy := "timestamp"
	if query.OrderBy != nil {
		orderBy = *query.OrderBy
	}

	order := "DESC"
	if query.Order != nil && (*query.Order == "asc" || *query.Order == "ASC") {
		order = "ASC"
	}

	baseQuery = baseQuery.Order(orderBy + " " + order)

	// Apply pagination
	if query.Offset != nil {
		baseQuery = baseQuery.Offset(*query.Offset)
	}

	if query.Limit != nil {
		baseQuery = baseQuery.Limit(*query.Limit)
	} else {
		// Default limit to prevent huge responses
		baseQuery = baseQuery.Limit(100)
	}

	// Preload deal information
	err := baseQuery.Preload("Deal").Find(&changes).Error
	return changes, total, err
}

// GetDealStateChangesByDealID retrieves all state changes for a specific deal
func GetDealStateChangesByDealID(db *gorm.DB, dealID DealID) ([]DealStateChange, error) {
	var changes []DealStateChange
	err := db.Where("deal_id = ?", dealID).
		Order("timestamp ASC").
		Preload("Deal").
		Find(&changes).Error
	return changes, err
}
