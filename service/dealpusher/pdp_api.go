package dealpusher

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/go-cid"
)

// PDPSchedulingConfig holds PDP-specific scheduling knobs for on-chain operations.
type PDPSchedulingConfig struct {
	BatchSize         int
	GasLimit          uint64
	ConfirmationDepth uint64
	PollingInterval   time.Duration
}

// Validate validates PDP scheduling configuration.
func (c PDPSchedulingConfig) Validate() error {
	if c.BatchSize <= 0 {
		return errors.New("pdp batch size must be greater than 0")
	}
	if c.ConfirmationDepth == 0 {
		return errors.New("pdp confirmation depth must be greater than 0")
	}
	if c.PollingInterval <= 0 {
		return errors.New("pdp polling interval must be greater than 0")
	}
	return nil
}

// PDPProofSetManager defines proof set lifecycle operations needed by scheduling.
type PDPProofSetManager interface {
	// EnsureProofSet returns an existing proof set ID or creates one for this client/provider pair.
	EnsureProofSet(ctx context.Context, wallet model.Wallet, provider string) (uint64, error)
	// QueueAddRoots submits root additions for a proof set and returns the queued tx reference.
	QueueAddRoots(ctx context.Context, proofSetID uint64, pieceCIDs []cid.Cid, cfg PDPSchedulingConfig) (*PDPQueuedTx, error)
}

// PDPTransactionConfirmer defines confirmation checks for queued on-chain transactions.
type PDPTransactionConfirmer interface {
	WaitForConfirmations(ctx context.Context, txHash string, depth uint64, pollInterval time.Duration) (*PDPTransactionReceipt, error)
}

// PDPQueuedTx represents an on-chain transaction submitted by PDP scheduling.
type PDPQueuedTx struct {
	Hash string
}

// PDPTransactionReceipt represents a confirmed on-chain transaction result.
type PDPTransactionReceipt struct {
	Hash        string
	BlockNumber uint64
	GasUsed     uint64
	Status      uint64
	CostAttoFIL *big.Int
}
