package model

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/ipfs/go-log/v2"
)

var dealConfigLogger = log.Logger("dealconfig")

// DealConfig encapsulates all deal-related configuration parameters
type DealConfig struct {
	// AutoCreateDeals enables automatic deal creation after preparation completes
	AutoCreateDeals bool `json:"autoCreateDeals" gorm:"default:false"`

	// DealProvider specifies the Storage Provider ID for deals
	DealProvider string `json:"dealProvider" gorm:"type:varchar(255)"`

	// DealTemplate specifies the deal template name or ID to use (optional)
	DealTemplate string `json:"dealTemplate" gorm:"type:varchar(255)"`

	// DealVerified indicates whether deals should be verified
	DealVerified bool `json:"dealVerified" gorm:"default:false"`

	// DealKeepUnsealed indicates whether to keep unsealed copy
	DealKeepUnsealed bool `json:"dealKeepUnsealed" gorm:"default:false"`

	// DealAnnounceToIpni indicates whether to announce to IPNI
	DealAnnounceToIpni bool `json:"dealAnnounceToIpni" gorm:"default:true"`

	// DealDuration specifies the deal duration (time.Duration for backward compatibility)
	DealDuration time.Duration `json:"dealDuration" swaggertype:"primitive,integer" gorm:"default:15552000000000000"` // ~180 days in nanoseconds

	// DealStartDelay specifies the deal start delay (time.Duration for backward compatibility)
	DealStartDelay time.Duration `json:"dealStartDelay" swaggertype:"primitive,integer" gorm:"default:86400000000000"` // ~1 day in nanoseconds

	// DealPricePerDeal specifies the price in FIL per deal
	DealPricePerDeal float64 `json:"dealPricePerDeal" gorm:"default:0"`

	// DealPricePerGb specifies the price in FIL per GiB
	DealPricePerGb float64 `json:"dealPricePerGb" gorm:"default:0"`

	// DealPricePerGbEpoch specifies the price in FIL per GiB per epoch
	DealPricePerGbEpoch float64 `json:"dealPricePerGbEpoch" gorm:"default:0"`

	// DealHTTPHeaders contains HTTP headers for deals
	DealHTTPHeaders ConfigMap `json:"dealHttpHeaders" gorm:"type:text" swaggertype:"object"`

	// DealURLTemplate specifies the URL template for deals
	DealURLTemplate string `json:"dealUrlTemplate" gorm:"type:text"`
}

// Validate validates the deal configuration and returns any errors
func (dc *DealConfig) Validate() error {
	// Validate numeric fields for negative values
	if dc.DealPricePerDeal < 0 {
		return fmt.Errorf("dealPricePerDeal cannot be negative: %f", dc.DealPricePerDeal)
	}
	if dc.DealPricePerGb < 0 {
		return fmt.Errorf("dealPricePerGb cannot be negative: %f", dc.DealPricePerGb)
	}
	if dc.DealPricePerGbEpoch < 0 {
		return fmt.Errorf("dealPricePerGbEpoch cannot be negative: %f", dc.DealPricePerGbEpoch)
	}
	if dc.DealDuration <= 0 {
		return fmt.Errorf("dealDuration must be positive: %v", dc.DealDuration)
	}
	if dc.DealStartDelay < 0 {
		return fmt.Errorf("dealStartDelay cannot be negative: %v", dc.DealStartDelay)
	}

	// Validate that at least one pricing model is used
	if dc.DealPricePerDeal == 0 && dc.DealPricePerGb == 0 && dc.DealPricePerGbEpoch == 0 {
		// This might be valid for free deals, so we don't error but could warn
	}

	// Validate provider format if specified
	if dc.DealProvider != "" {
		if len(dc.DealProvider) < 4 || dc.DealProvider[:1] != "f" {
			return fmt.Errorf("dealProvider must be a valid miner ID starting with 'f': %s", dc.DealProvider)
		}
		// Try to parse the number part
		if _, err := strconv.Atoi(dc.DealProvider[1:]); err != nil {
			return fmt.Errorf("dealProvider must be a valid miner ID (f<number>): %s", dc.DealProvider)
		}
	}

	return nil
}

// IsEmpty returns true if the deal config has no meaningful configuration
func (dc *DealConfig) IsEmpty() bool {
	return !dc.AutoCreateDeals &&
		dc.DealProvider == "" &&
		dc.DealTemplate == "" &&
		dc.DealPricePerDeal == 0 &&
		dc.DealPricePerGb == 0 &&
		dc.DealPricePerGbEpoch == 0 &&
		dc.DealURLTemplate == ""
}

// SetDurationFromString parses a duration string and converts it to time.Duration
// Supports formats like "180d", "24h", "30s" or direct epoch numbers
func (dc *DealConfig) SetDurationFromString(durationStr string) error {
	// First try to parse as a direct number (epochs)
	if epochs, err := strconv.ParseInt(durationStr, 10, 64); err == nil {
		if epochs <= 0 {
			return fmt.Errorf("duration must be positive: %d", epochs)
		}
		// Convert epochs to time.Duration (assuming 30 second epoch time)
		const epochDuration = 30 * time.Second
		dc.DealDuration = time.Duration(epochs) * epochDuration
		return nil
	}

	// Try to parse as a Go duration
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("invalid duration format: %s (use format like '180d', '24h', or epoch number)", durationStr)
	}

	if duration <= 0 {
		return fmt.Errorf("duration must be positive: %s", durationStr)
	}

	dc.DealDuration = duration
	return nil
}

// SetStartDelayFromString parses a start delay string and converts it to time.Duration
func (dc *DealConfig) SetStartDelayFromString(delayStr string) error {
	// First try to parse as a direct number (epochs)
	if epochs, err := strconv.ParseInt(delayStr, 10, 64); err == nil {
		if epochs < 0 {
			return fmt.Errorf("start delay cannot be negative: %d", epochs)
		}
		// Convert epochs to time.Duration (assuming 30 second epoch time)
		const epochDuration = 30 * time.Second
		dc.DealStartDelay = time.Duration(epochs) * epochDuration
		return nil
	}

	// Try to parse as a Go duration
	duration, err := time.ParseDuration(delayStr)
	if err != nil {
		return fmt.Errorf("invalid delay format: %s (use format like '1d', '2h', or epoch number)", delayStr)
	}

	if duration < 0 {
		return fmt.Errorf("start delay cannot be negative: %s", delayStr)
	}

	dc.DealStartDelay = duration
	return nil
}

// ToMap converts the DealConfig to a map for template override operations
func (dc *DealConfig) ToMap() map[string]interface{} {
	result := make(map[string]interface{})

	// Use reflection-like approach with json marshaling/unmarshaling
	jsonData, _ := json.Marshal(dc)
	json.Unmarshal(jsonData, &result)

	return result
}

// ApplyOverrides applies template values to zero-value fields in the deal config
func (dc *DealConfig) ApplyOverrides(template *DealConfig) {
	if template == nil {
		return
	}

	dealConfigLogger.Debug("Applying template overrides to DealConfig")

	// Apply template values only to zero-value fields
	if !dc.AutoCreateDeals && template.AutoCreateDeals {
		dealConfigLogger.Debugf("Overriding AutoCreateDeals: %v -> %v", dc.AutoCreateDeals, template.AutoCreateDeals)
		dc.AutoCreateDeals = template.AutoCreateDeals
	}
	if dc.DealProvider == "" && template.DealProvider != "" {
		dealConfigLogger.Debugf("Overriding DealProvider: '%s' -> '%s'", dc.DealProvider, template.DealProvider)
		dc.DealProvider = template.DealProvider
	}
	if dc.DealTemplate == "" && template.DealTemplate != "" {
		dealConfigLogger.Debugf("Overriding DealTemplate: '%s' -> '%s'", dc.DealTemplate, template.DealTemplate)
		dc.DealTemplate = template.DealTemplate
	}
	if !dc.DealVerified && template.DealVerified {
		dealConfigLogger.Debugf("Overriding DealVerified: %v -> %v", dc.DealVerified, template.DealVerified)
		dc.DealVerified = template.DealVerified
	}
	if !dc.DealKeepUnsealed && template.DealKeepUnsealed {
		dealConfigLogger.Debugf("Overriding DealKeepUnsealed: %v -> %v", dc.DealKeepUnsealed, template.DealKeepUnsealed)
		dc.DealKeepUnsealed = template.DealKeepUnsealed
	}
	if !dc.DealAnnounceToIpni && template.DealAnnounceToIpni {
		dealConfigLogger.Debugf("Overriding DealAnnounceToIpni: %v -> %v", dc.DealAnnounceToIpni, template.DealAnnounceToIpni)
		dc.DealAnnounceToIpni = template.DealAnnounceToIpni
	}
	if dc.DealDuration == 0 && template.DealDuration != 0 {
		dealConfigLogger.Debugf("Overriding DealDuration: %v -> %v", dc.DealDuration, template.DealDuration)
		dc.DealDuration = template.DealDuration
	}
	if dc.DealStartDelay == 0 && template.DealStartDelay != 0 {
		dealConfigLogger.Debugf("Overriding DealStartDelay: %v -> %v", dc.DealStartDelay, template.DealStartDelay)
		dc.DealStartDelay = template.DealStartDelay
	}
	if dc.DealPricePerDeal == 0 && template.DealPricePerDeal != 0 {
		dealConfigLogger.Debugf("Overriding DealPricePerDeal: %v -> %v", dc.DealPricePerDeal, template.DealPricePerDeal)
		dc.DealPricePerDeal = template.DealPricePerDeal
	}
	if dc.DealPricePerGb == 0 && template.DealPricePerGb != 0 {
		dealConfigLogger.Debugf("Overriding DealPricePerGb: %v -> %v", dc.DealPricePerGb, template.DealPricePerGb)
		dc.DealPricePerGb = template.DealPricePerGb
	}
	if dc.DealPricePerGbEpoch == 0 && template.DealPricePerGbEpoch != 0 {
		dealConfigLogger.Debugf("Overriding DealPricePerGbEpoch: %v -> %v", dc.DealPricePerGbEpoch, template.DealPricePerGbEpoch)
		dc.DealPricePerGbEpoch = template.DealPricePerGbEpoch
	}
	if dc.DealURLTemplate == "" && template.DealURLTemplate != "" {
		dealConfigLogger.Debugf("Overriding DealURLTemplate: '%s' -> '%s'", dc.DealURLTemplate, template.DealURLTemplate)
		dc.DealURLTemplate = template.DealURLTemplate
	}
	if len(dc.DealHTTPHeaders) == 0 && len(template.DealHTTPHeaders) > 0 {
		dealConfigLogger.Debugf("Overriding DealHTTPHeaders: %d headers -> %d headers", len(dc.DealHTTPHeaders), len(template.DealHTTPHeaders))
		dc.DealHTTPHeaders = template.DealHTTPHeaders
	}

	dealConfigLogger.Debug("Template override application completed")
}
