package dealpusher

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/data-preservation-programs/go-synapse/signer"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-cid"
)

// DDOSchedulingConfig holds DDO-specific scheduling knobs for on-chain operations.
type DDOSchedulingConfig struct {
	BatchSize         int           // pieces per createAllocationRequests tx
	ConfirmationDepth uint64        // block confirmations before considering tx final
	PollingInterval   time.Duration // confirmation polling interval
	TermMin           int64         // min term in epochs, default 518400 (~6 months)
	TermMax           int64         // max term in epochs, default 5256000 (~5 years)
	ExpirationOffset  int64         // expiration offset in epochs, default 172800
}

// Validate validates DDO scheduling configuration.
func (c DDOSchedulingConfig) Validate() error {
	if c.BatchSize <= 0 {
		return errors.New("ddo batch size must be greater than 0")
	}
	if c.ConfirmationDepth == 0 {
		return errors.New("ddo confirmation depth must be greater than 0")
	}
	if c.PollingInterval <= 0 {
		return errors.New("ddo polling interval must be greater than 0")
	}
	if c.TermMin <= 0 {
		return errors.New("ddo min term must be greater than 0")
	}
	if c.TermMax <= 0 {
		return errors.New("ddo max term must be greater than 0")
	}
	if c.TermMin > c.TermMax {
		return errors.New("ddo min term must be less than or equal to max term")
	}
	if c.ExpirationOffset < 0 {
		return errors.New("ddo expiration offset must be non-negative")
	}
	return nil
}

// DDODealManager defines DDO allocation lifecycle operations needed by scheduling.
// Path B implements this using the ddo-client SDK.
type DDODealManager interface {
	// ValidateSP checks that the provider is registered and active in the
	// DDO contract, and returns its on-chain config.
	ValidateSP(ctx context.Context, providerActorID uint64) (*DDOSPConfig, error)

	// CheckBalance queries the wallet's native FIL and payment token balances,
	// as well as the payments contract account status. Returns a summary that
	// the scheduler uses for pre-flight logging and low-balance warnings.
	CheckBalance(ctx context.Context, walletAddr common.Address) (*DDOBalanceStatus, error)

	// EnsurePayments checks account balance and operator approval, deposits
	// and approves if needed. Takes the actual pieces (not aggregated totals)
	// because the SDK computes per-piece lockup: allocationLockupAmount * len(pieces).
	EnsurePayments(ctx context.Context, evmSigner signer.EVMSigner,
		pieces []DDOPieceSubmission, cfg DDOSchedulingConfig) error

	// CreateAllocations submits a batch of pieces as DDO allocations.
	CreateAllocations(ctx context.Context, evmSigner signer.EVMSigner,
		pieces []DDOPieceSubmission, cfg DDOSchedulingConfig) (*DDOQueuedTx, error)

	// WaitForConfirmations polls for tx confirmation to the specified depth.
	WaitForConfirmations(ctx context.Context, txHash string,
		depth uint64, pollInterval time.Duration) (*DDOTransactionReceipt, error)

	// ParseAllocationIDs extracts allocation IDs from a confirmed tx receipt.
	ParseAllocationIDs(ctx context.Context, txHash string) ([]uint64, error)
}

type DDOSPConfig struct {
	IsActive     bool
	MinPieceSize uint64
	MaxPieceSize uint64
	MinTermLen   int64
	MaxTermLen   int64
}

type DDOPieceSubmission struct {
	PieceCID    cid.Cid
	PieceSize   uint64
	ProviderID  uint64
	DownloadURL string
}

type DDOQueuedTx struct {
	Hash string
}

type DDOTransactionReceipt struct {
	Hash        string
	BlockNumber uint64
	GasUsed     uint64
	Status      uint64
}

// DDOBalanceStatus summarizes a wallet's balance state for DDO deal-making.
type DDOBalanceStatus struct {
	WalletAddr     common.Address
	NativeFIL      *big.Int // native FIL balance (for gas)
	TokenBalance   *big.Int // payment token balance held in wallet
	DepositedFunds *big.Int // funds already deposited in payments contract
	LockupCurrent  *big.Int // current lockup in payments contract
	Available      *big.Int // deposited - lockup (spendable for new deals)
}
