package storage

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/notification"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

var logger = log.Logger("sp-validator")

type SPValidationResult struct {
	IsValid         bool            `json:"isValid"`
	ProviderID      string          `json:"providerId"`
	ProviderAddress string          `json:"providerAddress,omitempty"`
	PeerID          string          `json:"peerId,omitempty"`
	Multiaddrs      []string        `json:"multiaddrs,omitempty"`
	IsOnline        bool            `json:"isOnline"`
	Power           string          `json:"power,omitempty"`
	SectorSize      string          `json:"sectorSize,omitempty"`
	AcceptingDeals  bool            `json:"acceptingDeals"`
	Message         string          `json:"message"`
	Warnings        []string        `json:"warnings,omitempty"`
	Metadata        model.ConfigMap `json:"metadata,omitempty"`
}

// MinerInfo represents storage provider information
type MinerInfo struct {
	PeerID     *peer.ID              `json:"peerId,omitempty"`
	Multiaddrs []multiaddr.Multiaddr `json:"multiaddrs"`
	SectorSize abi.SectorSize        `json:"sectorSize"`
}

// MinerPower represents storage provider power information
type MinerPower struct {
	MinerPower Claim `json:"minerPower"`
}

// Claim represents power claim information
type Claim struct {
	QualityAdjPower abi.StoragePower `json:"qualityAdjPower"`
}

type DefaultSPEntry struct {
	ProviderID      string          `json:"providerId"`
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	Verified        bool            `json:"verified"`
	RecommendedUse  string          `json:"recommendedUse"`
	DefaultSettings model.ConfigMap `json:"defaultSettings"`
}

type SPValidator struct {
	notificationHandler *notification.Handler
	defaultSPs          []DefaultSPEntry
}

func NewSPValidator() *SPValidator {
	return &SPValidator{
		notificationHandler: notification.Default,
		defaultSPs:          getDefaultStorageProviders(),
	}
}

var DefaultSPValidator = NewSPValidator()

// ValidateStorageProvider checks if a storage provider is available and accepting deals
func (v *SPValidator) ValidateStorageProvider(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	providerID string,
	preparationID string,
) (*SPValidationResult, error) {
	result := &SPValidationResult{
		ProviderID: providerID,
		Metadata: model.ConfigMap{
			"preparation_id": preparationID,
			"provider_id":    providerID,
		},
	}

	// Parse provider ID
	providerAddr, err := address.NewFromString(providerID)
	if err != nil {
		result.IsValid = false
		result.Message = "Invalid storage provider ID format"
		v.logError(ctx, db, "Invalid Storage Provider ID", result.Message, result.Metadata)
		return result, nil
	}

	result.ProviderAddress = providerAddr.String()

	// Check if provider exists in the network
	minerInfo, err := v.getMinerInfo(ctx, lotusClient, providerAddr)
	if err != nil {
		result.IsValid = false
		result.Message = "Storage provider not found on network"
		result.Metadata["error"] = err.Error()
		v.logError(ctx, db, "Storage Provider Not Found", result.Message, result.Metadata)
		return result, nil
	}

	// Extract peer ID and multiaddrs
	if minerInfo.PeerID != nil {
		result.PeerID = minerInfo.PeerID.String()
	}

	result.Multiaddrs = make([]string, len(minerInfo.Multiaddrs))
	for i, addr := range minerInfo.Multiaddrs {
		result.Multiaddrs[i] = addr.String()
	}

	// Check if provider is online
	isOnline, connectWarnings := v.checkProviderConnectivity(ctx, lotusClient, result.PeerID, result.Multiaddrs)
	result.IsOnline = isOnline
	result.Warnings = append(result.Warnings, connectWarnings...)

	// Get provider power and sector size
	power, err := v.getMinerPower(ctx, lotusClient, providerAddr)
	if err != nil {
		result.Warnings = append(result.Warnings, "Could not retrieve miner power information")
	} else {
		result.Power = power.MinerPower.QualityAdjPower.String()
	}

	result.SectorSize = fmt.Sprintf("%d", minerInfo.SectorSize)

	// Check if provider is accepting deals
	acceptingDeals, dealWarnings := v.checkDealAcceptance(ctx, lotusClient, providerAddr)
	result.AcceptingDeals = acceptingDeals
	result.Warnings = append(result.Warnings, dealWarnings...)

	// Determine overall validity
	if result.IsOnline && result.AcceptingDeals {
		result.IsValid = true
		result.Message = "Storage provider is available and accepting deals"
		v.logInfo(ctx, db, "Storage Provider Validation Successful", result.Message, result.Metadata)
	} else {
		result.IsValid = false
		issues := []string{}
		if !result.IsOnline {
			issues = append(issues, "not online")
		}
		if !result.AcceptingDeals {
			issues = append(issues, "not accepting deals")
		}
		result.Message = fmt.Sprintf("Storage provider validation failed: %s", strings.Join(issues, ", "))
		v.logWarning(ctx, db, "Storage Provider Validation Failed", result.Message, result.Metadata)
	}

	return result, nil
}

// GetDefaultStorageProviders returns a list of recommended default storage providers
func (v *SPValidator) GetDefaultStorageProviders() []DefaultSPEntry {
	return v.defaultSPs
}

// GetDefaultStorageProvider returns a recommended storage provider for auto-creation
func (v *SPValidator) GetDefaultStorageProvider(ctx context.Context, db *gorm.DB, criteria string) (*DefaultSPEntry, error) {
	// For now, return the first available default SP
	// In the future, this could be more sophisticated based on criteria
	if len(v.defaultSPs) == 0 {
		return nil, errors.New("no default storage providers configured")
	}

	defaultSP := v.defaultSPs[0]

	// Log the selection
	metadata := model.ConfigMap{
		"selected_provider": defaultSP.ProviderID,
		"criteria":          criteria,
	}
	v.logInfo(ctx, db, "Default Storage Provider Selected", fmt.Sprintf("Selected %s for auto-creation", defaultSP.ProviderID), metadata)

	return &defaultSP, nil
}

// ValidateAndGetDefault validates a provider, and if it fails, returns a default one
func (v *SPValidator) ValidateAndGetDefault(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	providerID string,
	preparationID string,
) (*SPValidationResult, *DefaultSPEntry, error) {
	// First try to validate the specified provider
	if providerID != "" {
		result, err := v.ValidateStorageProvider(ctx, db, lotusClient, providerID, preparationID)
		if err != nil {
			return nil, nil, err
		}
		if result.IsValid {
			return result, nil, nil
		}
	}

	// If validation failed or no provider specified, get a default one
	defaultSP, err := v.GetDefaultStorageProvider(ctx, db, "fallback")
	if err != nil {
		return nil, nil, err
	}

	// Validate the default provider
	defaultResult, err := v.ValidateStorageProvider(ctx, db, lotusClient, defaultSP.ProviderID, preparationID)
	if err != nil {
		return nil, nil, err
	}

	return defaultResult, defaultSP, nil
}

// getMinerInfo retrieves miner information from the Lotus API
func (v *SPValidator) getMinerInfo(ctx context.Context, lotusClient jsonrpc.RPCClient, minerAddr address.Address) (*MinerInfo, error) {
	var minerInfo MinerInfo
	err := lotusClient.CallFor(ctx, &minerInfo, "Filecoin.StateMinerInfo", minerAddr, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &minerInfo, nil
}

// getMinerPower retrieves miner power information
func (v *SPValidator) getMinerPower(ctx context.Context, lotusClient jsonrpc.RPCClient, minerAddr address.Address) (*MinerPower, error) {
	var power MinerPower
	err := lotusClient.CallFor(ctx, &power, "Filecoin.StateMinerPower", minerAddr, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &power, nil
}

// checkProviderConnectivity checks if the provider is reachable
func (v *SPValidator) checkProviderConnectivity(ctx context.Context, lotusClient jsonrpc.RPCClient, peerID string, multiaddrs []string) (bool, []string) {
	var warnings []string

	if peerID == "" {
		warnings = append(warnings, "No peer ID available for connectivity check")
		return false, warnings
	}

	// Try to connect to the peer
	_, err := peer.Decode(peerID)
	if err != nil {
		warnings = append(warnings, fmt.Sprintf("Invalid peer ID format: %v", err))
		return false, warnings
	}

	// Check if we can connect (this is a simplified check)
	// In a real implementation, you might want to use libp2p to actually connect
	connected := v.checkPeerConnectivity(ctx, multiaddrs)
	if !connected {
		warnings = append(warnings, "Could not establish connection to storage provider")
	}

	return connected, warnings
}

// checkPeerConnectivity performs basic connectivity checks to multiaddrs
func (v *SPValidator) checkPeerConnectivity(ctx context.Context, multiaddrs []string) bool {
	for _, addr := range multiaddrs {
		if v.testConnection(ctx, addr) {
			return true
		}
	}
	return false
}

// testConnection tests if we can connect to a multiaddr
func (v *SPValidator) testConnection(ctx context.Context, multiaddr string) bool {
	// Parse multiaddr and extract IP and port
	// This is a simplified implementation
	parts := strings.Split(multiaddr, "/")
	if len(parts) < 5 {
		return false
	}

	var host, port string
	for i, part := range parts {
		if part == "ip4" && i+1 < len(parts) {
			host = parts[i+1]
		}
		if part == "tcp" && i+1 < len(parts) {
			port = parts[i+1]
		}
	}

	if host == "" || port == "" {
		return false
	}

	// Test TCP connection
	timeout := 5 * time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// checkDealAcceptance checks if the provider is accepting storage deals
func (v *SPValidator) checkDealAcceptance(ctx context.Context, lotusClient jsonrpc.RPCClient, minerAddr address.Address) (bool, []string) {
	var warnings []string

	// This is a placeholder - in a real implementation, you would check:
	// 1. Miner's ask price
	// 2. Deal acceptance policies
	// 3. Available storage capacity
	// 4. Reputation/past performance

	// For now, we'll do a basic check if the miner has any deals
	// You could implement more sophisticated checks here

	// Simple heuristic: if miner has power, they're likely accepting deals
	power, err := v.getMinerPower(ctx, lotusClient, minerAddr)
	if err != nil {
		warnings = append(warnings, "Could not verify deal acceptance status")
		return false, warnings
	}

	// If miner has quality adjusted power > 0, assume they're accepting deals
	if power.MinerPower.QualityAdjPower.Sign() > 0 {
		return true, warnings
	}

	warnings = append(warnings, "Storage provider appears to have no active storage power")
	return false, warnings
}

// getDefaultStorageProviders returns hardcoded list of reliable SPs
func getDefaultStorageProviders() []DefaultSPEntry {
	return []DefaultSPEntry{
		{
			ProviderID:     "f01000", // Example provider ID
			Name:           "Example SP 1",
			Description:    "Reliable storage provider with good track record",
			Verified:       true,
			RecommendedUse: "General purpose storage deals",
			DefaultSettings: model.ConfigMap{
				"price_per_gb_epoch": "0.0000000001",
				"verified":           "true",
				"duration":           "535 days",
				"start_delay":        "72h",
			},
		},
		{
			ProviderID:     "f01001", // Example provider ID
			Name:           "Example SP 2",
			Description:    "Fast retrieval focused storage provider",
			Verified:       true,
			RecommendedUse: "Fast retrieval scenarios",
			DefaultSettings: model.ConfigMap{
				"price_per_gb_epoch": "0.0000000002",
				"verified":           "true",
				"duration":           "535 days",
				"start_delay":        "48h",
			},
		},
	}
}

// Helper methods for logging
func (v *SPValidator) logError(ctx context.Context, db *gorm.DB, title, message string, metadata model.ConfigMap) {
	_, err := v.notificationHandler.LogError(ctx, db, "sp-validator", title, message, metadata)
	if err != nil {
		logger.Errorf("Failed to log error notification: %v", err)
	}
}

func (v *SPValidator) logWarning(ctx context.Context, db *gorm.DB, title, message string, metadata model.ConfigMap) {
	_, err := v.notificationHandler.LogWarning(ctx, db, "sp-validator", title, message, metadata)
	if err != nil {
		logger.Errorf("Failed to log warning notification: %v", err)
	}
}

func (v *SPValidator) logInfo(ctx context.Context, db *gorm.DB, title, message string, metadata model.ConfigMap) {
	_, err := v.notificationHandler.LogInfo(ctx, db, "sp-validator", title, message, metadata)
	if err != nil {
		logger.Errorf("Failed to log info notification: %v", err)
	}
}
