// Package pdptracker - Event-driven ingestion interfaces for scalable PDP tracking
//
// This file defines the interfaces needed to replace scan-based RPC reads with
// event-driven ingestion for improved scalability. Part of issue #615.
package pdptracker

import (
	"context"
	"math/big"
	"time"
)

// EventSource provides cursor-based event ingestion from contract logs.
// This interface abstracts the event source to allow different implementations
// (e.g., direct RPC, external indexer, etc.)
type EventSource interface {
	// GetEventsByRange returns events within the specified block range
	// cursor is the last processed block number (exclusive)
	// Returns events and the latest block number processed
	GetEventsByRange(ctx context.Context, cursor uint64, limit int) ([]ContractEvent, uint64, error)
	
	// GetLatestBlockNumber returns the current latest block number
	GetLatestBlockNumber(ctx context.Context) (uint64, error)
	
	// Close releases any resources held by the event source
	Close() error
}

// ContractEvent represents a parsed event from the PDP contract
type ContractEvent struct {
	BlockNumber   uint64
	TxHash        string
	EventType     EventType
	ProofSetID    uint64
	ClientAddress string
	Data          map[string]interface{} // Event-specific data
	Timestamp     time.Time
}

// EventType represents the type of PDP contract event
type EventType string

const (
	EventTypeProofSetCreated   EventType = "ProofSetCreated"
	EventTypeProofSetUpdated   EventType = "ProofSetUpdated"
	EventTypeChallengePosted   EventType = "ChallengePosted"
	EventTypeChallengeAnswered EventType = "ChallengeAnswered"
	EventTypeProofSetExpired   EventType = "ProofSetExpired"
)

// HydrationBatcher batches RPC calls to fetch current state for events.
// This is used to get up-to-date state information that may not be
// available in the event data alone.
type HydrationBatcher interface {
	// BatchHydrate takes a batch of events and enriches them with current state
	// Returns hydrated events with current on-chain state
	BatchHydrate(ctx context.Context, events []ContractEvent) ([]HydratedEvent, error)
	
	// SetBatchSize configures the maximum number of events to process in a single batch
	SetBatchSize(size int)
	
	// GetBatchSize returns the current batch size configuration
	GetBatchSize() int
}

// HydratedEvent extends ContractEvent with current state information
type HydratedEvent struct {
	ContractEvent
	CurrentState ProofSetState
}

// ProofSetState represents the current on-chain state of a proof set
type ProofSetState struct {
	IsLive             bool
	NextChallengeEpoch int32
	LastChallengeEpoch int32
	PieceCIDs          []string
	ProviderAddress    string
	TotalSize          *big.Int
}

// CursorStore manages persistent checkpoint state for resume-safe processing
type CursorStore interface {
	// GetCursor returns the last processed block number
	// Returns 0 if no cursor is stored (first run)
	GetCursor(ctx context.Context) (uint64, error)
	
	// SetCursor atomically updates the cursor position
	// This should be called after successfully processing events up to blockNumber
	SetCursor(ctx context.Context, blockNumber uint64) error
	
	// GetCursorWithMetadata returns cursor with additional processing metadata
	GetCursorWithMetadata(ctx context.Context) (*CursorMetadata, error)
	
	// SetCursorWithMetadata atomically updates cursor and metadata
	SetCursorWithMetadata(ctx context.Context, metadata *CursorMetadata) error
}

// CursorMetadata contains checkpoint state and processing metrics
type CursorMetadata struct {
	BlockNumber      uint64    // Last processed block number
	LastUpdate       time.Time // When cursor was last updated
	EventsProcessed  uint64    // Total events processed
	LastEventHash    string    // Hash of last processed event (for verification)
	ProcessingLagMs  int64     // Processing lag in milliseconds
	BatchesProcessed uint64    // Total batches processed
}

// EventDrivenConfig contains configuration for the event-driven pipeline
type EventDrivenConfig struct {
	// Enable toggles event-driven mode vs scan-based mode
	Enabled bool `json:"enabled"`
	
	// BatchSize for event processing
	BatchSize int `json:"batch_size"`
	
	// PollingInterval for checking new events
	PollingInterval time.Duration `json:"polling_interval"`
	
	// MaxRetries for failed event processing
	MaxRetries int `json:"max_retries"`
	
	// RetryBackoff configuration
	RetryBackoff time.Duration `json:"retry_backoff"`
	
	// MaxProcessingLag threshold for alerting
	MaxProcessingLag time.Duration `json:"max_processing_lag"`
}

// EventProcessor orchestrates the event-driven pipeline
type EventProcessor interface {
	// Start begins event processing from the stored cursor position
	Start(ctx context.Context) error
	
	// Stop gracefully stops event processing
	Stop(ctx context.Context) error
	
	// GetProcessingStats returns current processing metrics
	GetProcessingStats() ProcessingStats
	
	// HealthCheck returns the current health status
	HealthCheck(ctx context.Context) error
}

// ProcessingStats contains metrics about event processing
type ProcessingStats struct {
	CurrentCursor     uint64
	LatestBlockNumber uint64
	ProcessingLag     time.Duration
	EventsProcessed   uint64
	BatchesProcessed  uint64
	LastProcessedAt   time.Time
	ErrorCount        uint64
	RetryCount        uint64
}