package wallet

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)
// WalletWithBalance combines wallet info with live balance fields
type WalletWithBalance struct {
	model.Wallet
	Balance        string `json:"balance"`
	BalanceAttoFIL string `json:"balanceAttoFIL"`
	DataCap        string `json:"dataCap"`
	DataCapBytes   int64  `json:"dataCapBytes"`
	Error          *string `json:"error,omitempty"`
}

// ListWithBalanceHandler retrieves all wallets and fetches their balances from Lotus
func ListWithBalanceHandler(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
) ([]WalletWithBalance, error) {
	db = db.WithContext(ctx)
	var wallets []model.Wallet
	err := db.Find(&wallets).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
   var result []WalletWithBalance
   allFailed := true
   for _, w := range wallets {
	   bal, err := DefaultHandler{}.GetBalanceHandler(ctx, db, lotusClient, w.Address)
	   wb := WalletWithBalance{Wallet: w}
	   if err == nil && bal != nil {
		   wb.Balance = bal.Balance
		   wb.BalanceAttoFIL = bal.BalanceAttoFIL
		   wb.DataCap = bal.DataCap
		   wb.DataCapBytes = bal.DataCapBytes
		   wb.Error = bal.Error
		   allFailed = false
	   } else if err != nil {
		   errMsg := err.Error()
		   wb.Error = &errMsg
	   }
	   result = append(result, wb)
   }
   if allFailed && len(wallets) > 0 {
	   return result, errors.New("Failed to fetch balances for all wallets. Lotus may be unreachable or misconfigured.")
   }
   return result, nil
}



// ListHandler retrieves a list of all the wallets stored in the database.
//
// Parameters:
//   - ctx: The context for database transactions and other operations.
//   - db: A pointer to the gorm.DB instance representing the database connection.
//
// Returns:
//   - A slice containing all Wallet models from the database.
//   - An error, if any occurred during the database fetch operation.
func (DefaultHandler) ListHandler(
	ctx context.Context,
	db *gorm.DB,
) ([]model.Wallet, error) {
	db = db.WithContext(ctx)
	var wallets []model.Wallet

	err := db.Find(&wallets).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return wallets, nil
}

// @ID ListWallets
// @Summary List all imported wallets
// @Tags Wallet
// @Produce json
// @Success 200 {array} model.Wallet
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /wallet [get]
func _() {}
