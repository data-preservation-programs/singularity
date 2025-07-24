package errorcategorization

import (
	"regexp"
	"time"

	"github.com/data-preservation-programs/singularity/model"
)

// ErrorCategory represents different categories of errors that can occur during deal processing
type ErrorCategory string

const (
	// Network-related errors
	NetworkTimeout         ErrorCategory = "network_timeout"
	NetworkConnectionError ErrorCategory = "network_connection_error"
	NetworkProtocolError   ErrorCategory = "network_protocol_error"
	NetworkDNSError        ErrorCategory = "network_dns_error"

	// Deal proposal errors
	DealRejectedByProvider ErrorCategory = "deal_rejected_by_provider"
	DealInvalidPrice       ErrorCategory = "deal_invalid_price"
	DealInvalidDuration    ErrorCategory = "deal_invalid_duration"
	DealInvalidStartTime   ErrorCategory = "deal_invalid_start_time"
	DealDuplicateProposal  ErrorCategory = "deal_duplicate_proposal"

	// Storage provider errors
	ProviderUnavailable  ErrorCategory = "provider_unavailable"
	ProviderCapacityFull ErrorCategory = "provider_capacity_full"
	ProviderConfigError  ErrorCategory = "provider_config_error"
	ProviderVersionError ErrorCategory = "provider_version_error"

	// Client/wallet errors
	ClientInsufficientFunds   ErrorCategory = "client_insufficient_funds"
	ClientInvalidAddress      ErrorCategory = "client_invalid_address"
	ClientAuthenticationError ErrorCategory = "client_authentication_error"

	// Data/piece errors
	PieceNotFound    ErrorCategory = "piece_not_found"
	PieceCorrupted   ErrorCategory = "piece_corrupted"
	PieceInvalidCID  ErrorCategory = "piece_invalid_cid"
	PieceInvalidSize ErrorCategory = "piece_invalid_size"

	// System/internal errors
	SystemDatabaseError ErrorCategory = "system_database_error"
	SystemMemoryError   ErrorCategory = "system_memory_error"
	SystemDiskError     ErrorCategory = "system_disk_error"
	SystemConfigError   ErrorCategory = "system_config_error"

	// Unknown/unclassified errors
	UnknownError ErrorCategory = "unknown_error"
)

// ErrorPattern represents a pattern for matching and categorizing errors
type ErrorPattern struct {
	Pattern     *regexp.Regexp
	Category    ErrorCategory
	DealState   model.DealState
	Description string
	Severity    ErrorSeverity
	Retryable   bool
}

// ErrorSeverity represents the severity level of an error
type ErrorSeverity string

const (
	SeverityCritical ErrorSeverity = "critical" // System-level errors that require immediate attention
	SeverityHigh     ErrorSeverity = "high"     // Errors that prevent deal completion but are recoverable
	SeverityMedium   ErrorSeverity = "medium"   // Errors that may be temporary and retryable
	SeverityLow      ErrorSeverity = "low"      // Minor errors or warnings
)

// ErrorCategorization holds the result of error categorization
type ErrorCategorization struct {
	Category    ErrorCategory   `json:"category"`
	DealState   model.DealState `json:"dealState"`
	Description string          `json:"description"`
	Severity    ErrorSeverity   `json:"severity"`
	Retryable   bool            `json:"retryable"`
	Metadata    *ErrorMetadata  `json:"metadata,omitempty"`
}

// ErrorMetadata holds additional metadata about the error
type ErrorMetadata struct {
	// Network-related metadata
	NetworkEndpoint string `json:"networkEndpoint,omitempty"`
	NetworkLatency  *int64 `json:"networkLatency,omitempty"` // in milliseconds
	DNSResolution   string `json:"dnsResolution,omitempty"`

	// Provider-related metadata
	ProviderID      string `json:"providerId,omitempty"`
	ProviderVersion string `json:"providerVersion,omitempty"`
	ProviderRegion  string `json:"providerRegion,omitempty"`

	// Deal-related metadata
	ProposalID      string     `json:"proposalId,omitempty"`
	PieceCID        string     `json:"pieceCid,omitempty"`
	PieceSize       *int64     `json:"pieceSize,omitempty"`
	DealPrice       string     `json:"dealPrice,omitempty"`
	DealDuration    *int64     `json:"dealDuration,omitempty"` // in epochs
	StartEpoch      *int32     `json:"startEpoch,omitempty"`
	EndEpoch        *int32     `json:"endEpoch,omitempty"`
	AttemptNumber   *int       `json:"attemptNumber,omitempty"`
	LastAttemptTime *time.Time `json:"lastAttemptTime,omitempty"`

	// Client-related metadata
	ClientAddress string `json:"clientAddress,omitempty"`
	WalletBalance string `json:"walletBalance,omitempty"`

	// System-related metadata
	SystemLoad     *float64 `json:"systemLoad,omitempty"`
	MemoryUsage    *int64   `json:"memoryUsage,omitempty"`   // in bytes
	DiskSpaceUsed  *int64   `json:"diskSpaceUsed,omitempty"` // in bytes
	DatabaseHealth string   `json:"databaseHealth,omitempty"`

	// Additional custom fields
	CustomFields map[string]interface{} `json:"customFields,omitempty"`
}

// errorPatterns contains the comprehensive list of error patterns for categorization
var errorPatterns = []ErrorPattern{
	// Network timeout errors
	{
		Pattern:     regexp.MustCompile(`(?i)(context deadline exceeded|timeout|timed out)`),
		Category:    NetworkTimeout,
		DealState:   model.DealErrored,
		Description: "Network timeout during deal negotiation",
		Severity:    SeverityMedium,
		Retryable:   true,
	},
	{
		Pattern:     regexp.MustCompile(`(?i)(connection.*timeout|read.*timeout|write.*timeout)`),
		Category:    NetworkTimeout,
		DealState:   model.DealErrored,
		Description: "Connection timeout during communication",
		Severity:    SeverityMedium,
		Retryable:   true,
	},

	// Network connection errors
	{
		Pattern:     regexp.MustCompile(`(?i)(connection refused|connection reset|connection aborted)`),
		Category:    NetworkConnectionError,
		DealState:   model.DealErrored,
		Description: "Network connection failure",
		Severity:    SeverityMedium,
		Retryable:   true,
	},
	{
		Pattern:     regexp.MustCompile(`(?i)(network.*unreachable|host.*unreachable|no route to host)`),
		Category:    NetworkConnectionError,
		DealState:   model.DealErrored,
		Description: "Network routing failure",
		Severity:    SeverityHigh,
		Retryable:   true,
	},

	// Network protocol errors
	{
		Pattern:     regexp.MustCompile(`(?i)(no supported protocols?|protocol.*not.*supported)`),
		Category:    NetworkProtocolError,
		DealState:   model.DealErrored,
		Description: "No supported storage protocols found",
		Severity:    SeverityHigh,
		Retryable:   false,
	},
	{
		Pattern:     regexp.MustCompile(`(?i)(invalid.*protocol|unsupported.*protocol|protocol.*error)`),
		Category:    NetworkProtocolError,
		DealState:   model.DealErrored,
		Description: "Protocol compatibility error",
		Severity:    SeverityHigh,
		Retryable:   false,
	},

	// DNS errors
	{
		Pattern:     regexp.MustCompile(`(?i)(dns.*resolution.*failed|no such host|name.*not.*resolved)`),
		Category:    NetworkDNSError,
		DealState:   model.DealErrored,
		Description: "DNS resolution failure",
		Severity:    SeverityMedium,
		Retryable:   true,
	},

	// Deal rejection errors
	{
		Pattern:     regexp.MustCompile(`(?i)(deal rejected|proposal.*rejected|rejected.*proposal)`),
		Category:    DealRejectedByProvider,
		DealState:   model.DealRejected,
		Description: "Deal rejected by storage provider",
		Severity:    SeverityHigh,
		Retryable:   false,
	},
	{
		Pattern:     regexp.MustCompile(`(?i)(price.*too.*low|insufficient.*price|price.*rejected)`),
		Category:    DealInvalidPrice,
		DealState:   model.DealRejected,
		Description: "Deal price too low or invalid",
		Severity:    SeverityMedium,
		Retryable:   false,
	},
	{
		Pattern:     regexp.MustCompile(`(?i)(duration.*invalid|duration.*too.*short|duration.*too.*long)`),
		Category:    DealInvalidDuration,
		DealState:   model.DealRejected,
		Description: "Deal duration is invalid",
		Severity:    SeverityMedium,
		Retryable:   false,
	},
	{
		Pattern:     regexp.MustCompile(`(?i)(start.*time.*invalid|start.*epoch.*invalid|deal.*start.*too.*soon|deal.*start.*too.*late)`),
		Category:    DealInvalidStartTime,
		DealState:   model.DealRejected,
		Description: "Deal start time is invalid",
		Severity:    SeverityMedium,
		Retryable:   false,
	},
	{
		Pattern:     regexp.MustCompile(`(?i)(duplicate.*proposal|proposal.*exists|identical.*proposal)`),
		Category:    DealDuplicateProposal,
		DealState:   model.DealErrored,
		Description: "Duplicate deal proposal",
		Severity:    SeverityLow,
		Retryable:   false,
	},

	// Storage provider errors
	{
		Pattern:     regexp.MustCompile(`(?i)(provider.*unavailable|provider.*offline|provider.*not.*responding)`),
		Category:    ProviderUnavailable,
		DealState:   model.DealErrored,
		Description: "Storage provider is unavailable",
		Severity:    SeverityHigh,
		Retryable:   true,
	},
	{
		Pattern:     regexp.MustCompile(`(?i)(storage.*full|capacity.*exceeded|no.*space.*available|disk.*full)`),
		Category:    ProviderCapacityFull,
		DealState:   model.DealRejected,
		Description: "Storage provider capacity is full",
		Severity:    SeverityMedium,
		Retryable:   true,
	},
	{
		Pattern:     regexp.MustCompile(`(?i)(provider.*configuration.*error|miner.*config.*error|sealing.*config.*error)`),
		Category:    ProviderConfigError,
		DealState:   model.DealErrored,
		Description: "Storage provider configuration error",
		Severity:    SeverityHigh,
		Retryable:   false,
	},
	{
		Pattern:     regexp.MustCompile(`(?i)(unsupported.*version|version.*mismatch|incompatible.*version)`),
		Category:    ProviderVersionError,
		DealState:   model.DealErrored,
		Description: "Storage provider version compatibility error",
		Severity:    SeverityHigh,
		Retryable:   false,
	},

	// Client/wallet errors
	{
		Pattern:     regexp.MustCompile(`(?i)(insufficient.*funds|balance.*too.*low|not.*enough.*funds)`),
		Category:    ClientInsufficientFunds,
		DealState:   model.DealErrored,
		Description: "Client has insufficient funds",
		Severity:    SeverityHigh,
		Retryable:   true,
	},
	{
		Pattern:     regexp.MustCompile(`(?i)(invalid.*address|address.*not.*found|unknown.*address)`),
		Category:    ClientInvalidAddress,
		DealState:   model.DealErrored,
		Description: "Client address is invalid",
		Severity:    SeverityHigh,
		Retryable:   false,
	},
	{
		Pattern:     regexp.MustCompile(`(?i)(authentication.*failed|unauthorized|access.*denied|permission.*denied)`),
		Category:    ClientAuthenticationError,
		DealState:   model.DealErrored,
		Description: "Client authentication failed",
		Severity:    SeverityHigh,
		Retryable:   false,
	},

	// Data/piece errors
	{
		Pattern:     regexp.MustCompile(`(?i)(piece.*not.*found|cid.*not.*found|data.*not.*found)`),
		Category:    PieceNotFound,
		DealState:   model.DealErrored,
		Description: "Piece data not found",
		Severity:    SeverityHigh,
		Retryable:   false,
	},
	{
		Pattern:     regexp.MustCompile(`(?i)(piece.*corrupted|data.*corrupted|checksum.*mismatch|integrity.*check.*failed)`),
		Category:    PieceCorrupted,
		DealState:   model.DealErrored,
		Description: "Piece data is corrupted",
		Severity:    SeverityCritical,
		Retryable:   false,
	},
	{
		Pattern:     regexp.MustCompile(`(?i)(invalid.*cid|malformed.*cid|cid.*parse.*error)`),
		Category:    PieceInvalidCID,
		DealState:   model.DealErrored,
		Description: "Piece CID is invalid",
		Severity:    SeverityHigh,
		Retryable:   false,
	},
	{
		Pattern:     regexp.MustCompile(`(?i)(invalid.*piece.*size|size.*mismatch|wrong.*size)`),
		Category:    PieceInvalidSize,
		DealState:   model.DealErrored,
		Description: "Piece size is invalid",
		Severity:    SeverityHigh,
		Retryable:   false,
	},

	// System/internal errors
	{
		Pattern:     regexp.MustCompile(`(?i)(database.*error|sql.*error|connection.*pool.*exhausted|database.*timeout)`),
		Category:    SystemDatabaseError,
		DealState:   model.DealErrored,
		Description: "Database system error",
		Severity:    SeverityCritical,
		Retryable:   true,
	},
	{
		Pattern:     regexp.MustCompile(`(?i)(out.*of.*memory|memory.*exhausted|allocation.*failed)`),
		Category:    SystemMemoryError,
		DealState:   model.DealErrored,
		Description: "System memory error",
		Severity:    SeverityCritical,
		Retryable:   true,
	},
	{
		Pattern:     regexp.MustCompile(`(?i)(disk.*error|storage.*error|i/o.*error|write.*failed|read.*failed)`),
		Category:    SystemDiskError,
		DealState:   model.DealErrored,
		Description: "System disk error",
		Severity:    SeverityCritical,
		Retryable:   true,
	},
	{
		Pattern:     regexp.MustCompile(`(?i)(configuration.*error|config.*invalid|missing.*config)`),
		Category:    SystemConfigError,
		DealState:   model.DealErrored,
		Description: "System configuration error",
		Severity:    SeverityHigh,
		Retryable:   false,
	},
}

// CategorizeError analyzes an error message and returns comprehensive categorization information
func CategorizeError(errorMessage string) *ErrorCategorization {
	// Try to match against known error patterns
	for _, pattern := range errorPatterns {
		if pattern.Pattern.MatchString(errorMessage) {
			return &ErrorCategorization{
				Category:    pattern.Category,
				DealState:   pattern.DealState,
				Description: pattern.Description,
				Severity:    pattern.Severity,
				Retryable:   pattern.Retryable,
				Metadata:    &ErrorMetadata{},
			}
		}
	}

	// Default categorization for unknown errors
	return &ErrorCategorization{
		Category:    UnknownError,
		DealState:   model.DealErrored,
		Description: "Unknown error occurred during deal processing",
		Severity:    SeverityMedium,
		Retryable:   true,
		Metadata:    &ErrorMetadata{},
	}
}

// CategorizeErrorWithContext analyzes an error message with additional context information
func CategorizeErrorWithContext(errorMessage string, contextMetadata *ErrorMetadata) *ErrorCategorization {
	categorization := CategorizeError(errorMessage)

	// Merge context metadata with the base metadata
	if contextMetadata != nil {
		if categorization.Metadata == nil {
			categorization.Metadata = &ErrorMetadata{}
		}

		// Merge fields from contextMetadata
		if contextMetadata.NetworkEndpoint != "" {
			categorization.Metadata.NetworkEndpoint = contextMetadata.NetworkEndpoint
		}
		if contextMetadata.NetworkLatency != nil {
			categorization.Metadata.NetworkLatency = contextMetadata.NetworkLatency
		}
		if contextMetadata.ProviderID != "" {
			categorization.Metadata.ProviderID = contextMetadata.ProviderID
		}
		if contextMetadata.ProviderVersion != "" {
			categorization.Metadata.ProviderVersion = contextMetadata.ProviderVersion
		}
		if contextMetadata.ProposalID != "" {
			categorization.Metadata.ProposalID = contextMetadata.ProposalID
		}
		if contextMetadata.PieceCID != "" {
			categorization.Metadata.PieceCID = contextMetadata.PieceCID
		}
		if contextMetadata.PieceSize != nil {
			categorization.Metadata.PieceSize = contextMetadata.PieceSize
		}
		if contextMetadata.DealPrice != "" {
			categorization.Metadata.DealPrice = contextMetadata.DealPrice
		}
		if contextMetadata.DealDuration != nil {
			categorization.Metadata.DealDuration = contextMetadata.DealDuration
		}
		if contextMetadata.StartEpoch != nil {
			categorization.Metadata.StartEpoch = contextMetadata.StartEpoch
		}
		if contextMetadata.EndEpoch != nil {
			categorization.Metadata.EndEpoch = contextMetadata.EndEpoch
		}
		if contextMetadata.AttemptNumber != nil {
			categorization.Metadata.AttemptNumber = contextMetadata.AttemptNumber
		}
		if contextMetadata.LastAttemptTime != nil {
			categorization.Metadata.LastAttemptTime = contextMetadata.LastAttemptTime
		}
		if contextMetadata.ClientAddress != "" {
			categorization.Metadata.ClientAddress = contextMetadata.ClientAddress
		}
		if contextMetadata.WalletBalance != "" {
			categorization.Metadata.WalletBalance = contextMetadata.WalletBalance
		}
		if contextMetadata.CustomFields != nil {
			categorization.Metadata.CustomFields = contextMetadata.CustomFields
		}
	}

	return categorization
}

// IsRetryableError determines if an error category is retryable
func IsRetryableError(category ErrorCategory) bool {
	for _, pattern := range errorPatterns {
		if pattern.Category == category {
			return pattern.Retryable
		}
	}
	return false
}

// GetErrorSeverity returns the severity level for an error category
func GetErrorSeverity(category ErrorCategory) ErrorSeverity {
	for _, pattern := range errorPatterns {
		if pattern.Category == category {
			return pattern.Severity
		}
	}
	return SeverityMedium
}

// GetSupportedCategories returns all supported error categories
func GetSupportedCategories() []ErrorCategory {
	categories := make([]ErrorCategory, 0, len(errorPatterns))
	seen := make(map[ErrorCategory]bool)

	for _, pattern := range errorPatterns {
		if !seen[pattern.Category] {
			categories = append(categories, pattern.Category)
			seen[pattern.Category] = true
		}
	}

	return categories
}
