package dealpusher

import (
	"context"
	"time"

	"github.com/data-preservation-programs/go-synapse/signer"
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

// DDODealManager defines DDO allocation lifecycle operations needed by scheduling.
// Path B implements this using the ddo-client SDK.
type DDODealManager interface {
	// ValidateSP checks that the provider is registered and active in the
	// DDO contract, and returns its on-chain config.
	ValidateSP(ctx context.Context, providerActorID uint64) (*DDOSPConfig, error)

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

// DDOAllocationTracker polls allocation activation status for deal tracking.
// Path B implements this; we define a local struct to avoid importing ddo-client types.
type DDOAllocationTracker interface {
	GetAllocationInfo(ctx context.Context, allocationID uint64) (*DDOAllocationStatus, error)
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

type DDOAllocationStatus struct {
	Activated    bool
	SectorNumber uint64
}
