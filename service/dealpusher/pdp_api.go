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

// PDPSchedulingConfig holds PDP-specific scheduling knobs for on-chain operations.
type PDPSchedulingConfig struct {
	BatchSize            int // pieces per addPieces transaction (gas constraint)
	MaxPiecesPerProofSet int // pieces per proof set before handoff to SP
	ConfirmationDepth    uint64
	PollingInterval      time.Duration
}

// Validate validates PDP scheduling configuration.
func (c PDPSchedulingConfig) Validate() error {
	if c.BatchSize <= 0 {
		return errors.New("pdp batch size must be greater than 0")
	}
	if c.MaxPiecesPerProofSet <= 0 {
		return errors.New("pdp max pieces per proof set must be greater than 0")
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
// All methods take an EVMSigner because they submit FEVM transactions.
type PDPProofSetManager interface {
	// EnsureProofSet returns an existing proof set ID in assembling state or creates one.
	EnsureProofSet(ctx context.Context, evmSigner signer.EVMSigner, provider string, cfg PDPSchedulingConfig) (uint64, error)
	// QueueAddRoots submits root additions for a proof set and returns the queued tx reference.
	// pieceSizes are the padded piece sizes corresponding to each CID, needed for CommPv2 conversion.
	QueueAddRoots(ctx context.Context, evmSigner signer.EVMSigner, proofSetID uint64, pieceCIDs []cid.Cid, pieceSizes []int64, cfg PDPSchedulingConfig) (*PDPQueuedTx, error)
	// ProposeTransfer proposes the proof set for handoff to the given SP.
	ProposeTransfer(ctx context.Context, evmSigner signer.EVMSigner, proofSetID uint64, spEVMAddress common.Address) error
	// IncrementPieceCount atomically increments the durable piece count after confirmed on-chain add.
	IncrementPieceCount(ctx context.Context, proofSetID uint64, count int) error
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
