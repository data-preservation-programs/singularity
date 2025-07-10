package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDealConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		config  DealConfig
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid config",
			config: DealConfig{
				AutoCreateDeals:     true,
				DealProvider:        "f01000",
				DealDuration:        180 * 24 * time.Hour,
				DealStartDelay:      24 * time.Hour,
				DealPricePerDeal:    0.1,
				DealPricePerGb:      0.01,
				DealPricePerGbEpoch: 0.001,
			},
			wantErr: false,
		},
		{
			name: "negative price per deal",
			config: DealConfig{
				DealPricePerDeal: -1.0,
			},
			wantErr: true,
			errMsg:  "dealPricePerDeal cannot be negative",
		},
		{
			name: "negative price per gb",
			config: DealConfig{
				DealPricePerGb: -1.0,
			},
			wantErr: true,
			errMsg:  "dealPricePerGb cannot be negative",
		},
		{
			name: "negative price per gb epoch",
			config: DealConfig{
				DealPricePerGbEpoch: -1.0,
			},
			wantErr: true,
			errMsg:  "dealPricePerGbEpoch cannot be negative",
		},
		{
			name: "zero duration",
			config: DealConfig{
				DealDuration: 0,
			},
			wantErr: true,
			errMsg:  "dealDuration must be positive",
		},
		{
			name: "negative start delay",
			config: DealConfig{
				DealDuration:   time.Hour,
				DealStartDelay: -time.Hour,
			},
			wantErr: true,
			errMsg:  "dealStartDelay cannot be negative",
		},
		{
			name: "invalid provider format",
			config: DealConfig{
				DealDuration: time.Hour,
				DealProvider: "invalid",
			},
			wantErr: true,
			errMsg:  "dealProvider must be a valid miner ID starting with 'f'",
		},
		{
			name: "valid provider format",
			config: DealConfig{
				DealDuration: time.Hour,
				DealProvider: "f01234",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDealConfig_IsEmpty(t *testing.T) {
	tests := []struct {
		name   string
		config DealConfig
		want   bool
	}{
		{
			name: "default config (matches CLI defaults, should not be empty)",
			config: DealConfig{
				DealVerified:     true,
				DealKeepUnsealed: true,
				// Add other new defaults here if needed
			},
			want: false,
		},
		{
			name: "config with auto create deals",
			config: DealConfig{
				AutoCreateDeals: true,
			},
			want: false,
		},
		{
			name: "config with provider",
			config: DealConfig{
				DealProvider: "f01000",
			},
			want: false,
		},
		{
			name: "config with template",
			config: DealConfig{
				DealTemplate: "template1",
			},
			want: false,
		},
		{
			name: "config with pricing",
			config: DealConfig{
				DealPricePerDeal: 0.1,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.config.IsEmpty())
		})
	}
}

func TestDealConfig_SetDurationFromString(t *testing.T) {
	tests := []struct {
		name        string
		durationStr string
		expectDur   time.Duration
		expectErr   bool
		errMsg      string
	}{
		{
			name:        "valid epoch number",
			durationStr: "518400", // 180 days in epochs
			expectDur:   518400 * 30 * time.Second,
			expectErr:   false,
		},
		{
			name:        "valid duration string",
			durationStr: "24h",
			expectDur:   24 * time.Hour,
			expectErr:   false,
		},
		{
			name:        "valid duration with days (converted)",
			durationStr: "180d",
			expectErr:   true, // Go duration doesn't support 'd' unit
			errMsg:      "invalid duration format",
		},
		{
			name:        "zero epochs",
			durationStr: "0",
			expectErr:   true,
			errMsg:      "duration must be positive",
		},
		{
			name:        "negative epochs",
			durationStr: "-100",
			expectErr:   true,
			errMsg:      "duration must be positive",
		},
		{
			name:        "invalid format",
			durationStr: "invalid",
			expectErr:   true,
			errMsg:      "invalid duration format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &DealConfig{}
			err := config.SetDurationFromString(tt.durationStr)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectDur, config.DealDuration)
			}
		})
	}
}

func TestDealConfig_SetStartDelayFromString(t *testing.T) {
	tests := []struct {
		name        string
		delayStr    string
		expectDelay time.Duration
		expectErr   bool
		errMsg      string
	}{
		{
			name:        "valid epoch number",
			delayStr:    "2880", // 1 day in epochs
			expectDelay: 2880 * 30 * time.Second,
			expectErr:   false,
		},
		{
			name:        "valid duration string",
			delayStr:    "2h",
			expectDelay: 2 * time.Hour,
			expectErr:   false,
		},
		{
			name:        "zero delay",
			delayStr:    "0",
			expectDelay: 0,
			expectErr:   false,
		},
		{
			name:      "negative epochs",
			delayStr:  "-100",
			expectErr: true,
			errMsg:    "start delay cannot be negative",
		},
		{
			name:      "invalid format",
			delayStr:  "invalid",
			expectErr: true,
			errMsg:    "invalid delay format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &DealConfig{}
			err := config.SetStartDelayFromString(tt.delayStr)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectDelay, config.DealStartDelay)
			}
		})
	}
}

func TestDealConfig_ApplyOverrides(t *testing.T) {
	config := &DealConfig{
		AutoCreateDeals:  false,
		DealProvider:     "",
		DealPricePerDeal: 0,
		DealDuration:     0,
	}

	template := &DealConfig{
		AutoCreateDeals:  true,
		DealProvider:     "f01000",
		DealPricePerDeal: 0.1,
		DealDuration:     24 * time.Hour,
		DealTemplate:     "template1",
	}

	config.ApplyOverrides(template)

	// Should apply template values to zero-value fields
	assert.True(t, config.AutoCreateDeals)
	assert.Equal(t, "f01000", config.DealProvider)
	assert.Equal(t, 0.1, config.DealPricePerDeal)
	assert.Equal(t, 24*time.Hour, config.DealDuration)
	assert.Equal(t, "template1", config.DealTemplate)

	// Test with existing values - should not override
	config2 := &DealConfig{
		AutoCreateDeals:  true, // This should stay true (explicit)
		DealProvider:     "f02000",
		DealPricePerDeal: 0.2,
		DealDuration:     48 * time.Hour,
	}

	config2.ApplyOverrides(template)

	// Should not override existing non-zero values
	assert.True(t, config2.AutoCreateDeals) // Stays true (explicit)
	assert.Equal(t, "f02000", config2.DealProvider)
	assert.Equal(t, 0.2, config2.DealPricePerDeal)
	assert.Equal(t, 48*time.Hour, config2.DealDuration)
}

func TestDealConfig_ToMap(t *testing.T) {
	config := &DealConfig{
		AutoCreateDeals:    true,
		DealProvider:       "f01000",
		DealPricePerDeal:   0.1,
		DealDuration:       24 * time.Hour,
		DealAnnounceToIpni: true,
	}

	result := config.ToMap()

	assert.NotNil(t, result)
	assert.Equal(t, true, result["autoCreateDeals"])
	assert.Equal(t, "f01000", result["dealProvider"])
	assert.Equal(t, 0.1, result["dealPricePerDeal"])
	assert.Equal(t, true, result["dealAnnounceToIpni"])
}

func TestDealConfig_ApplyOverrides_NilTemplate(t *testing.T) {
	config := &DealConfig{
		DealProvider: "f01000",
	}

	// Should not panic or change anything
	config.ApplyOverrides(nil)
	assert.Equal(t, "f01000", config.DealProvider)
}
