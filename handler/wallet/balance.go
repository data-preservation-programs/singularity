package wallet

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

type BalanceResponse struct {
	Address        string  `json:"address"`
	Balance        string  `json:"balance"`        // FIL balance in FIL units
	BalanceAttoFIL string  `json:"balanceAttoFIL"` // Raw balance in attoFIL
	DataCap        string  `json:"dataCap"`        // FIL+ datacap balance
	DataCapBytes   int64   `json:"dataCapBytes"`   // Raw datacap in bytes
	Error          *string `json:"error,omitempty"` // Error message if any
}

// GetBalanceHandler retrieves the FIL and FIL+ balance for a specific wallet
//
// Parameters:
//   - ctx: The context for the request
//   - db: Database connection
//   - lotusClient: Lotus JSON-RPC client
//   - address: Wallet address to query
//
// Returns:
//   - BalanceResponse containing balance information
//   - Error if the operation fails
func (DefaultHandler) GetBalanceHandler(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	walletAddress string,
) (*BalanceResponse, error) {
	// Validate wallet address format
	addr, err := address.NewFromString(walletAddress)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid wallet address format: %s", walletAddress)
	}

	// Initialize response
	response := &BalanceResponse{
		Address:      walletAddress,
		Balance:      "0",
		DataCap:      "0",
		DataCapBytes: 0,
	}

	// Get FIL balance using Lotus API
	balance, err := getWalletBalance(ctx, lotusClient, addr)
	if err != nil {
		errMsg := fmt.Sprintf("failed to get wallet balance: %v", err)
		response.Error = &errMsg
	} else {
		response.Balance = formatFILFromAttoFIL(balance.Int)
		response.BalanceAttoFIL = balance.Int.String()
	}

	// Get FIL+ datacap balance
	datacap, err := getDatacapBalance(ctx, lotusClient, addr)
	if err != nil {
		// Always show datacap errors to help debug
		if response.Error != nil {
			errMsg := fmt.Sprintf("%s; failed to get datacap balance: %v", *response.Error, err)
			response.Error = &errMsg
		} else {
			errMsg := fmt.Sprintf("failed to get datacap balance: %v", err)
			response.Error = &errMsg
		}
		datacap = 0
	}
	
	response.DataCap = formatDatacap(datacap)
	response.DataCapBytes = datacap

	return response, nil
}

// getWalletBalance retrieves FIL balance from Lotus
func getWalletBalance(ctx context.Context, lotusClient jsonrpc.RPCClient, addr address.Address) (abi.TokenAmount, error) {
	var balance string
	// Pass address as string to avoid parameter marshaling issues
	err := lotusClient.CallFor(ctx, &balance, "Filecoin.WalletBalance", addr.String())
	if err != nil {
		return abi.TokenAmount{}, errors.WithStack(err)
	}

	balanceInt, ok := new(big.Int).SetString(balance, 10)
	if !ok {
		return abi.TokenAmount{}, errors.New("failed to parse balance")
	}

	return abi.TokenAmount{Int: balanceInt}, nil
}

// getDatacapBalance retrieves FIL+ datacap balance from Lotus
func getDatacapBalance(ctx context.Context, lotusClient jsonrpc.RPCClient, addr address.Address) (int64, error) {
	var result string
	// Use Filecoin.StateVerifiedClientStatus to get datacap allocation
	// This is the API method that corresponds to "lotus filplus check-client-datacap"
	err := lotusClient.CallFor(ctx, &result, "Filecoin.StateVerifiedClientStatus", addr.String(), nil)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	
	// If result is empty or null, client has no datacap
	if result == "" {
		return 0, nil
	}
	
	// Parse the datacap balance string
	datacap, err := strconv.ParseInt(result, 10, 64)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	
	return datacap, nil
}

// formatFILFromAttoFIL converts attoFIL to human-readable FIL
func formatFILFromAttoFIL(attoFIL *big.Int) string {
	if attoFIL == nil {
		return "0.000000 FIL"
	}
	
	filValue := new(big.Float).SetInt(attoFIL)
	filValue.Quo(filValue, big.NewFloat(1e18))
	
	return filValue.Text('f', 6) + " FIL" // 6 decimal places with FIL unit
}

// formatDatacap formats datacap bytes to human-readable format
func formatDatacap(bytes int64) string {
	if bytes == 0 {
		return "0.00 GiB"
	}
	
	// Convert bytes to GiB for display
	gib := float64(bytes) / (1024 * 1024 * 1024)
	return fmt.Sprintf("%.2f GiB", gib)
}

// @ID GetWalletBalance
// @Summary Get wallet FIL and FIL+ balance
// @Description Retrieves the FIL balance and FIL+ datacap balance for a specific wallet address
// @Tags Wallet
// @Param address path string true "Wallet address"
// @Produce json
// @Success 200 {object} BalanceResponse
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /wallet/{address}/balance [get]
func _() {}
