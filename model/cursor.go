package model

import (
	"time"
	"gorm.io/gorm"
)

// PDPCursor represents persistent checkpoint state for event-driven PDP tracking
// This table stores the last processed block number and processing metadata
// to enable resume-safe event processing after restarts.
type PDPCursor struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	
	// Service identifier to support multiple trackers
	ServiceName       string    `gorm:"uniqueIndex;not null" json:"serviceName"`
	
	// Last successfully processed block number
	BlockNumber       uint64    `gorm:"not null;default:0" json:"blockNumber"`
	
	// Metadata about processing state
	EventsProcessed   uint64    `gorm:"not null;default:0" json:"eventsProcessed"`
	BatchesProcessed  uint64    `gorm:"not null;default:0" json:"batchesProcessed"`
	LastEventHash     string    `json:"lastEventHash,omitempty"`
	ProcessingLagMs   int64     `gorm:"not null;default:0" json:"processingLagMs"`
	
	// Health and error tracking
	ErrorCount        uint64    `gorm:"not null;default:0" json:"errorCount"`
	LastError         string    `json:"lastError,omitempty"`
	LastErrorAt       *time.Time `json:"lastErrorAt,omitempty"`
	
	// Processing status
	IsHealthy         bool      `gorm:"not null;default:true" json:"isHealthy"`
	LastHealthCheck   time.Time `json:"lastHealthCheck"`
}

// TableName returns the table name for PDPCursor
func (PDPCursor) TableName() string {
	return "pdp_cursors"
}

// BeforeCreate sets default values before creating a PDPCursor record
func (c *PDPCursor) BeforeCreate(tx *gorm.DB) error {
	if c.ServiceName == "" {
		c.ServiceName = "pdptracker"
	}
	c.IsHealthy = true
	c.LastHealthCheck = time.Now()
	return nil
}

// UpdateHealth updates the health status and last check time
func (c *PDPCursor) UpdateHealth(isHealthy bool, errorMsg string) {
	c.IsHealthy = isHealthy
	c.LastHealthCheck = time.Now()
	if !isHealthy && errorMsg != "" {
		c.LastError = errorMsg
		now := time.Now()
		c.LastErrorAt = &now
		c.ErrorCount++
	}
}

// UpdateProgress updates processing progress metrics
func (c *PDPCursor) UpdateProgress(blockNumber uint64, eventsCount, batchesCount uint64, lastEventHash string, lagMs int64) {
	c.BlockNumber = blockNumber
	c.EventsProcessed += eventsCount
	c.BatchesProcessed += batchesCount
	c.LastEventHash = lastEventHash
	c.ProcessingLagMs = lagMs
}