package errorcategorization

import (
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCategorizeError(t *testing.T) {
	tests := []struct {
		name              string
		errorMessage      string
		expectedCategory  ErrorCategory
		expectedDealState model.DealState
		expectedSeverity  ErrorSeverity
		expectedRetryable bool
	}{
		// Network timeout errors
		{
			name:              "context deadline exceeded",
			errorMessage:      "context deadline exceeded",
			expectedCategory:  NetworkTimeout,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityMedium,
			expectedRetryable: true,
		},
		{
			name:              "connection timeout",
			errorMessage:      "connection timeout during handshake",
			expectedCategory:  NetworkTimeout,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityMedium,
			expectedRetryable: true,
		},
		{
			name:              "read timeout",
			errorMessage:      "read timeout on socket",
			expectedCategory:  NetworkTimeout,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityMedium,
			expectedRetryable: true,
		},

		// Network connection errors
		{
			name:              "connection refused",
			errorMessage:      "connection refused by peer",
			expectedCategory:  NetworkConnectionError,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityMedium,
			expectedRetryable: true,
		},
		{
			name:              "network unreachable",
			errorMessage:      "network is unreachable",
			expectedCategory:  NetworkConnectionError,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityHigh,
			expectedRetryable: true,
		},
		{
			name:              "no route to host",
			errorMessage:      "no route to host available",
			expectedCategory:  NetworkConnectionError,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityHigh,
			expectedRetryable: true,
		},

		// Network protocol errors
		{
			name:              "no supported protocols",
			errorMessage:      "no supported protocols found",
			expectedCategory:  NetworkProtocolError,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityHigh,
			expectedRetryable: false,
		},
		{
			name:              "protocol not supported",
			errorMessage:      "protocol version not supported",
			expectedCategory:  NetworkProtocolError,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityHigh,
			expectedRetryable: false,
		},

		// DNS errors
		{
			name:              "dns resolution failed",
			errorMessage:      "DNS resolution failed for hostname",
			expectedCategory:  NetworkDNSError,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityMedium,
			expectedRetryable: true,
		},
		{
			name:              "no such host",
			errorMessage:      "no such host exists",
			expectedCategory:  NetworkDNSError,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityMedium,
			expectedRetryable: true,
		},

		// Deal rejection errors
		{
			name:              "deal rejected",
			errorMessage:      "deal rejected by storage provider",
			expectedCategory:  DealRejectedByProvider,
			expectedDealState: model.DealRejected,
			expectedSeverity:  SeverityHigh,
			expectedRetryable: false,
		},
		{
			name:              "price too low",
			errorMessage:      "price too low for storage",
			expectedCategory:  DealInvalidPrice,
			expectedDealState: model.DealRejected,
			expectedSeverity:  SeverityMedium,
			expectedRetryable: false,
		},
		{
			name:              "duration invalid",
			errorMessage:      "duration is invalid for this provider",
			expectedCategory:  DealInvalidDuration,
			expectedDealState: model.DealRejected,
			expectedSeverity:  SeverityMedium,
			expectedRetryable: false,
		},
		{
			name:              "start time invalid",
			errorMessage:      "deal start time is too soon",
			expectedCategory:  DealInvalidStartTime,
			expectedDealState: model.DealRejected,
			expectedSeverity:  SeverityMedium,
			expectedRetryable: false,
		},
		{
			name:              "duplicate proposal",
			errorMessage:      "proposal is identical to existing one",
			expectedCategory:  DealDuplicateProposal,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityLow,
			expectedRetryable: false,
		},

		// Storage provider errors
		{
			name:              "provider unavailable",
			errorMessage:      "provider is unavailable at this time",
			expectedCategory:  ProviderUnavailable,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityHigh,
			expectedRetryable: true,
		},
		{
			name:              "storage full",
			errorMessage:      "storage capacity is full",
			expectedCategory:  ProviderCapacityFull,
			expectedDealState: model.DealRejected,
			expectedSeverity:  SeverityMedium,
			expectedRetryable: true,
		},
		{
			name:              "provider config error",
			errorMessage:      "provider configuration error detected",
			expectedCategory:  ProviderConfigError,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityHigh,
			expectedRetryable: false,
		},
		{
			name:              "version mismatch",
			errorMessage:      "version mismatch between client and provider",
			expectedCategory:  ProviderVersionError,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityHigh,
			expectedRetryable: false,
		},

		// Client/wallet errors
		{
			name:              "insufficient funds",
			errorMessage:      "insufficient funds in wallet",
			expectedCategory:  ClientInsufficientFunds,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityHigh,
			expectedRetryable: true,
		},
		{
			name:              "invalid address",
			errorMessage:      "invalid wallet address provided",
			expectedCategory:  ClientInvalidAddress,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityHigh,
			expectedRetryable: false,
		},
		{
			name:              "authentication failed",
			errorMessage:      "authentication failed for client",
			expectedCategory:  ClientAuthenticationError,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityHigh,
			expectedRetryable: false,
		},

		// Data/piece errors
		{
			name:              "piece not found",
			errorMessage:      "piece data not found in storage",
			expectedCategory:  PieceNotFound,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityHigh,
			expectedRetryable: false,
		},
		{
			name:              "data corrupted",
			errorMessage:      "piece data is corrupted or invalid",
			expectedCategory:  PieceCorrupted,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityCritical,
			expectedRetryable: false,
		},
		{
			name:              "invalid CID",
			errorMessage:      "invalid CID format provided",
			expectedCategory:  PieceInvalidCID,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityHigh,
			expectedRetryable: false,
		},
		{
			name:              "size mismatch",
			errorMessage:      "piece size mismatch detected",
			expectedCategory:  PieceInvalidSize,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityHigh,
			expectedRetryable: false,
		},

		// System/internal errors
		{
			name:              "database error",
			errorMessage:      "database connection error occurred",
			expectedCategory:  SystemDatabaseError,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityCritical,
			expectedRetryable: true,
		},
		{
			name:              "memory error",
			errorMessage:      "out of memory during processing",
			expectedCategory:  SystemMemoryError,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityCritical,
			expectedRetryable: true,
		},
		{
			name:              "disk error",
			errorMessage:      "disk I/O error during write operation",
			expectedCategory:  SystemDiskError,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityCritical,
			expectedRetryable: true,
		},
		{
			name:              "config error",
			errorMessage:      "configuration file is invalid",
			expectedCategory:  SystemConfigError,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityHigh,
			expectedRetryable: false,
		},

		// Unknown error
		{
			name:              "unknown error",
			errorMessage:      "some completely unknown error occurred",
			expectedCategory:  UnknownError,
			expectedDealState: model.DealErrored,
			expectedSeverity:  SeverityMedium,
			expectedRetryable: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CategorizeError(tt.errorMessage)
			require.NotNil(t, result)
			
			assert.Equal(t, tt.expectedCategory, result.Category)
			assert.Equal(t, tt.expectedDealState, result.DealState)
			assert.Equal(t, tt.expectedSeverity, result.Severity)
			assert.Equal(t, tt.expectedRetryable, result.Retryable)
			assert.NotEmpty(t, result.Description)
			assert.NotNil(t, result.Metadata)
		})
	}
}

func TestCategorizeErrorWithContext(t *testing.T) {
	errorMessage := "connection timeout during handshake"
	now := time.Now()
	pieceSize := int64(1024)
	attemptNum := 3
	
	contextMetadata := &ErrorMetadata{
		NetworkEndpoint:  "tcp://192.168.1.100:8080",
		NetworkLatency:   func() *int64 { latency := int64(500); return &latency }(),
		ProviderID:       "f01234",
		ProviderVersion:  "1.2.3",
		PieceCID:         "bafkqaaa",
		PieceSize:        &pieceSize,
		DealPrice:        "100",
		AttemptNumber:    &attemptNum,
		LastAttemptTime:  &now,
		ClientAddress:    "f3abc123",
		WalletBalance:    "1000",
		CustomFields: map[string]interface{}{
			"region": "us-west-1",
			"zone":   "us-west-1a",
		},
	}

	result := CategorizeErrorWithContext(errorMessage, contextMetadata)
	require.NotNil(t, result)
	
	// Verify basic categorization
	assert.Equal(t, NetworkTimeout, result.Category)
	assert.Equal(t, model.DealErrored, result.DealState)
	assert.Equal(t, SeverityMedium, result.Severity)
	assert.True(t, result.Retryable)
	
	// Verify metadata is properly merged
	require.NotNil(t, result.Metadata)
	assert.Equal(t, "tcp://192.168.1.100:8080", result.Metadata.NetworkEndpoint)
	assert.Equal(t, int64(500), *result.Metadata.NetworkLatency)
	assert.Equal(t, "f01234", result.Metadata.ProviderID)
	assert.Equal(t, "1.2.3", result.Metadata.ProviderVersion)
	assert.Equal(t, "bafkqaaa", result.Metadata.PieceCID)
	assert.Equal(t, int64(1024), *result.Metadata.PieceSize)
	assert.Equal(t, "100", result.Metadata.DealPrice)
	assert.Equal(t, 3, *result.Metadata.AttemptNumber)
	assert.Equal(t, now, *result.Metadata.LastAttemptTime)
	assert.Equal(t, "f3abc123", result.Metadata.ClientAddress)
	assert.Equal(t, "1000", result.Metadata.WalletBalance)
	assert.Equal(t, "us-west-1", result.Metadata.CustomFields["region"])
	assert.Equal(t, "us-west-1a", result.Metadata.CustomFields["zone"])
}

func TestCategorizeErrorWithNilContext(t *testing.T) {
	errorMessage := "deal rejected by storage provider"
	
	result := CategorizeErrorWithContext(errorMessage, nil)
	require.NotNil(t, result)
	
	assert.Equal(t, DealRejectedByProvider, result.Category)
	assert.Equal(t, model.DealRejected, result.DealState)
	assert.Equal(t, SeverityHigh, result.Severity)
	assert.False(t, result.Retryable)
	assert.NotNil(t, result.Metadata)
}

func TestIsRetryableError(t *testing.T) {
	tests := []struct {
		category         ErrorCategory
		expectedRetryable bool
	}{
		{NetworkTimeout, true},
		{NetworkConnectionError, true},
		{NetworkProtocolError, false},
		{DealRejectedByProvider, false},
		{DealInvalidPrice, false},
		{ProviderUnavailable, true},
		{ProviderCapacityFull, true},
		{ClientInsufficientFunds, true},
		{ClientInvalidAddress, false},
		{PieceNotFound, false},
		{PieceCorrupted, false},
		{SystemDatabaseError, true},
		{SystemMemoryError, true},
		{SystemConfigError, false},
		{UnknownError, true},
	}

	for _, tt := range tests {
		t.Run(string(tt.category), func(t *testing.T) {
			result := IsRetryableError(tt.category)
			assert.Equal(t, tt.expectedRetryable, result)
		})
	}
}

func TestGetErrorSeverity(t *testing.T) {
	tests := []struct {
		category         ErrorCategory
		expectedSeverity ErrorSeverity
	}{
		{NetworkTimeout, SeverityMedium},
		{NetworkConnectionError, SeverityMedium},
		{NetworkProtocolError, SeverityHigh},
		{DealRejectedByProvider, SeverityHigh},
		{DealInvalidPrice, SeverityMedium},
		{ProviderUnavailable, SeverityHigh},
		{ClientInsufficientFunds, SeverityHigh},
		{PieceCorrupted, SeverityCritical},
		{SystemDatabaseError, SeverityCritical},
		{SystemMemoryError, SeverityCritical},
		{SystemDiskError, SeverityCritical},
		{UnknownError, SeverityMedium},
	}

	for _, tt := range tests {
		t.Run(string(tt.category), func(t *testing.T) {
			result := GetErrorSeverity(tt.category)
			assert.Equal(t, tt.expectedSeverity, result)
		})
	}
}

func TestGetSupportedCategories(t *testing.T) {
	categories := GetSupportedCategories()
	
	// Verify we have a reasonable number of categories
	assert.Greater(t, len(categories), 10)
	
	// Verify some expected categories are present
	expectedCategories := []ErrorCategory{
		NetworkTimeout,
		NetworkConnectionError,
		NetworkProtocolError,
		DealRejectedByProvider,
		ProviderUnavailable,
		ClientInsufficientFunds,
		PieceNotFound,
		SystemDatabaseError,
		UnknownError,
	}
	
	for _, expected := range expectedCategories {
		found := false
		for _, category := range categories {
			if category == expected {
				found = true
				break
			}
		}
		assert.True(t, found, "Expected category %s not found in supported categories", expected)
	}
}

func TestErrorPatternsCoverage(t *testing.T) {
	// Test that our error patterns have good coverage
	
	// Test case-insensitive matching
	result := CategorizeError("DEAL REJECTED BY PROVIDER")
	assert.Equal(t, DealRejectedByProvider, result.Category)
	
	// Test partial matches
	result = CategorizeError("error: connection refused by remote host")
	assert.Equal(t, NetworkConnectionError, result.Category)
	
	// Test multiple keywords
	result = CategorizeError("timeout occurred during network operation")
	assert.Equal(t, NetworkTimeout, result.Category)
	
	// Test complex error messages
	result = CategorizeError("failed to establish connection: no route to host 192.168.1.1")
	assert.Equal(t, NetworkConnectionError, result.Category)
}

func TestErrorCategorizationFields(t *testing.T) {
	result := CategorizeError("connection timeout")
	require.NotNil(t, result)
	
	// Test all required fields are present
	assert.NotEmpty(t, result.Category)
	assert.NotEmpty(t, result.DealState)
	assert.NotEmpty(t, result.Description)
	assert.NotEmpty(t, result.Severity)
	assert.NotNil(t, result.Metadata)
}

// Benchmark tests for performance
func BenchmarkCategorizeError(b *testing.B) {
	errorMessage := "connection timeout during deal negotiation"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CategorizeError(errorMessage)
	}
}

func BenchmarkCategorizeErrorWithContext(b *testing.B) {
	errorMessage := "connection timeout during deal negotiation"
	contextMetadata := &ErrorMetadata{
		ProviderID:    "f01234",
		PieceCID:      "bafkqaaa",
		ClientAddress: "f3abc123",
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CategorizeErrorWithContext(errorMessage, contextMetadata)
	}
}