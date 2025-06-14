package wallet

import (
	"context"
	"fmt"
	"math/big"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/notification"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-log/v2"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

var validatorLogger = log.Logger("wallet-validator")

// formatFIL converts attoFIL (big.Int) to human-readable FIL string
func formatFIL(attoFIL *big.Int) string {
	if attoFIL == nil {
		return "0 FIL"
	}

	// Convert attoFIL to FIL (divide by 10^18)
	filValue := new(big.Float).SetInt(attoFIL)
	filValue.Quo(filValue, big.NewFloat(1e18))

	// Format with appropriate precision
	return fmt.Sprintf("%.9g FIL", filValue)
}

type ValidationResult struct {
	IsValid          bool            `json:"isValid"`
	WalletAddress    string          `json:"walletAddress"`
	CurrentBalance   string          `json:"currentBalance"`   // FIL amount as string
	RequiredBalance  string          `json:"requiredBalance"`  // FIL amount as string
	AvailableBalance string          `json:"availableBalance"` // FIL amount after pending deals
	Message          string          `json:"message"`
	Warnings         []string        `json:"warnings,omitempty"`
	Metadata         model.ConfigMap `json:"metadata,omitempty"`
}

type BalanceValidator struct {
	notificationHandler *notification.Handler
}

func NewBalanceValidator() *BalanceValidator {
	return &BalanceValidator{
		notificationHandler: notification.Default,
	}
}

var DefaultBalanceValidator = NewBalanceValidator()

// ValidateWalletBalance checks if a wallet has sufficient FIL balance for deals
func (v *BalanceValidator) ValidateWalletBalance(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	walletAddress string,
	requiredAmountAttoFIL *big.Int,
	preparationID string,
) (*ValidationResult, error) {
	result := &ValidationResult{
		WalletAddress:   walletAddress,
		RequiredBalance: formatFIL(requiredAmountAttoFIL),
		Metadata: model.ConfigMap{
			"preparation_id": preparationID,
			"wallet_address": walletAddress,
		},
	}

	// Parse wallet address
	addr, err := address.NewFromString(walletAddress)
	if err != nil {
		result.IsValid = false
		result.Message = "Invalid wallet address format"
		v.logError(ctx, db, "Invalid Wallet Address", result.Message, result.Metadata)
		return result, errors.WithStack(err)
	}

	// Get current wallet balance
	balance, err := v.getWalletBalance(ctx, lotusClient, addr)
	if err != nil {
		result.IsValid = false
		result.Message = "Failed to retrieve wallet balance"
		result.Metadata["error"] = err.Error()
		v.logError(ctx, db, "Wallet Balance Query Failed", result.Message, result.Metadata)
		return result, errors.WithStack(err)
	}

	result.CurrentBalance = formatFIL(balance.Int)

	// Get pending deals amount for this wallet
	pendingAmount, err := v.getPendingDealsAmount(ctx, db, walletAddress)
	if err != nil {
		logger.Warnf("Failed to get pending deals amount for wallet %s: %v", walletAddress, err)
		result.Warnings = append(result.Warnings, "Could not calculate pending deals amount")
		pendingAmount = big.NewInt(0)
	}

	// Calculate available balance (current - pending)
	availableBalance := new(big.Int).Sub(balance.Int, pendingAmount)
	if availableBalance.Sign() < 0 {
		availableBalance = big.NewInt(0)
	}
	result.AvailableBalance = formatFIL(availableBalance)

	// Check if available balance is sufficient
	if availableBalance.Cmp(requiredAmountAttoFIL) >= 0 {
		result.IsValid = true
		result.Message = "Wallet has sufficient balance for deal"
		v.logInfo(ctx, db, "Wallet Validation Successful", result.Message, result.Metadata)
	} else {
		result.IsValid = false
		shortage := new(big.Int).Sub(requiredAmountAttoFIL, availableBalance)
		result.Message = "Insufficient wallet balance. Shortage: " + formatFIL(shortage)
		result.Metadata["shortage_fil"] = formatFIL(shortage)
		result.Metadata["pending_deals_fil"] = formatFIL(pendingAmount)

		v.logWarning(ctx, db, "Insufficient Wallet Balance", result.Message, result.Metadata)
	}

	return result, nil
}

// ValidateWalletExists checks if a wallet exists and is accessible
func (v *BalanceValidator) ValidateWalletExists(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	walletAddress string,
	preparationID string,
) (*ValidationResult, error) {
	result := &ValidationResult{
		WalletAddress: walletAddress,
		Metadata: model.ConfigMap{
			"preparation_id": preparationID,
			"wallet_address": walletAddress,
		},
	}

	// Parse wallet address
	addr, err := address.NewFromString(walletAddress)
	if err != nil {
		result.IsValid = false
		result.Message = "Invalid wallet address format"
		v.logError(ctx, db, "Invalid Wallet Address", result.Message, result.Metadata)
		return result, errors.WithStack(err)
	}

	// Try to get wallet balance (this verifies wallet exists and is accessible)
	balance, err := v.getWalletBalance(ctx, lotusClient, addr)
	if err != nil {
		result.IsValid = false
		result.Message = "Wallet not found or not accessible"
		result.Metadata["error"] = err.Error()
		v.logError(ctx, db, "Wallet Not Accessible", result.Message, result.Metadata)
		return result, errors.WithStack(err)
	}

	result.IsValid = true
	result.CurrentBalance = formatFIL(balance.Int)
	result.Message = "Wallet exists and is accessible"
	v.logInfo(ctx, db, "Wallet Validation Successful", result.Message, result.Metadata)

	return result, nil
}

// CalculateRequiredBalance calculates the total FIL needed for deals based on parameters
func (v *BalanceValidator) CalculateRequiredBalance(
	pricePerGBEpoch float64,
	pricePerGB float64,
	pricePerDeal float64,
	totalSizeBytes int64,
	durationEpochs int64,
	numberOfDeals int,
) *big.Int {
	totalCost := big.NewFloat(0)

	// Price per GB epoch
	if pricePerGBEpoch > 0 {
		sizeGB := float64(totalSizeBytes) / (1024 * 1024 * 1024)
		epochCost := big.NewFloat(pricePerGBEpoch * sizeGB * float64(durationEpochs))
		totalCost.Add(totalCost, epochCost)
	}

	// Price per GB
	if pricePerGB > 0 {
		sizeGB := float64(totalSizeBytes) / (1024 * 1024 * 1024)
		gbCost := big.NewFloat(pricePerGB * sizeGB)
		totalCost.Add(totalCost, gbCost)
	}

	// Price per deal
	if pricePerDeal > 0 {
		dealCost := big.NewFloat(pricePerDeal * float64(numberOfDeals))
		totalCost.Add(totalCost, dealCost)
	}

	// Convert FIL to attoFIL (1 FIL = 10^18 attoFIL)
	attoFILPerFIL := big.NewFloat(1e18)
	totalAttoFIL := new(big.Float).Mul(totalCost, attoFILPerFIL)

	// Convert to big.Int
	result, _ := totalAttoFIL.Int(nil)
	return result
}

// getWalletBalance retrieves the current balance of a wallet
func (v *BalanceValidator) getWalletBalance(ctx context.Context, lotusClient jsonrpc.RPCClient, addr address.Address) (abi.TokenAmount, error) {
	var balance string
	err := lotusClient.CallFor(ctx, &balance, "Filecoin.WalletBalance", addr)
	if err != nil {
		return abi.TokenAmount{}, errors.WithStack(err)
	}

	// Parse balance string to big.Int
	balanceInt, ok := new(big.Int).SetString(balance, 10)
	if !ok {
		return abi.TokenAmount{}, errors.New("failed to parse balance")
	}

	return abi.TokenAmount{Int: balanceInt}, nil
}

// getPendingDealsAmount calculates the total amount locked in pending deals for a wallet
func (v *BalanceValidator) getPendingDealsAmount(ctx context.Context, db *gorm.DB, walletAddress string) (*big.Int, error) {
	var deals []model.Deal
	err := db.WithContext(ctx).Where("client_id = ? AND state IN (?)", walletAddress, []string{
		string(model.DealProposed),
		string(model.DealPublished),
	}).Find(&deals).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	totalPending := big.NewInt(0)
	for _, deal := range deals {
		// Parse deal price to big.Int (assuming it's in attoFIL)
		priceInt, ok := new(big.Int).SetString(deal.Price, 10)
		if ok {
			totalPending.Add(totalPending, priceInt)
		}
	}

	return totalPending, nil
}

// Helper methods for logging
func (v *BalanceValidator) logError(ctx context.Context, db *gorm.DB, title, message string, metadata model.ConfigMap) {
	_, err := v.notificationHandler.LogError(ctx, db, "wallet-validator", title, message, metadata)
	if err != nil {
		validatorLogger.Errorf("Failed to log error notification: %v", err)
	}
}

func (v *BalanceValidator) logWarning(ctx context.Context, db *gorm.DB, title, message string, metadata model.ConfigMap) {
	_, err := v.notificationHandler.LogWarning(ctx, db, "wallet-validator", title, message, metadata)
	if err != nil {
		validatorLogger.Errorf("Failed to log warning notification: %v", err)
	}
}

func (v *BalanceValidator) logInfo(ctx context.Context, db *gorm.DB, title, message string, metadata model.ConfigMap) {
	_, err := v.notificationHandler.LogInfo(ctx, db, "wallet-validator", title, message, metadata)
	if err != nil {
		validatorLogger.Errorf("Failed to log info notification: %v", err)
	}
}
