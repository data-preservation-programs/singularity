package dealpusher

import (
	"context"
	"errors"
	"time"

	"github.com/data-preservation-programs/go-synapse/signer"
	"github.com/ipfs/go-cid"
)

// PDPSchedulingConfig holds PDP-specific scheduling knobs for the FWSS-pull
// flow.
type PDPSchedulingConfig struct {
	// BatchSize bounds pieces per /pdp/piece/pull request.
	BatchSize int
	// MaxPiecesPerProofSet bounds pieces per data set; the scheduler starts
	// a new set when the current one fills.
	MaxPiecesPerProofSet int
	// PullTimeout bounds the time we wait for Curio to finish pulling a
	// batch (per request, not aggregate).
	PullTimeout time.Duration
}

// Validate validates PDP scheduling configuration.
func (c PDPSchedulingConfig) Validate() error {
	if c.BatchSize <= 0 {
		return errors.New("pdp batch size must be greater than 0")
	}
	if c.MaxPiecesPerProofSet <= 0 {
		return errors.New("pdp max pieces per proof set must be greater than 0")
	}
	if c.PullTimeout <= 0 {
		return errors.New("pdp pull timeout must be greater than 0")
	}
	return nil
}

// PDPPieceInput names a piece the scheduler wants pushed to the SP. The
// implementation constructs the SP-side source URL.
type PDPPieceInput struct {
	PieceCID cid.Cid
	// PieceSize is the padded piece size, needed for CommPv2 conversion
	// before signing.
	PieceSize int64
}

// PDPPullResult reports the outcome of a /pdp/piece/pull batch.
type PDPPullResult struct {
	// DataSetID is the FWSS-listened data set the pieces ended up in.
	// For new sets, this is the SetID Curio returns after the
	// createDataSet+addPieces tx confirms.
	DataSetID uint64
}

// PDPProofSetManager pushes pieces to an SP's Curio via /pdp/piece/pull,
// then triggers the SP's on-chain commit via /pdp/data-sets/create-and-add
// (new sets) or /pdp/data-sets/{id}/pieces (existing). It encapsulates SP
// service-URL discovery, clientDataSetId persistence, EIP-712 signing,
// and post-completion bookkeeping.
type PDPProofSetManager interface {
	// PullPiecesToFWSS pushes a batch of pieces to the SP via the FWSS-pull
	// flow. If no assembling proof set has room, a new FWSS-listened set is
	// created atomically with the first batch; otherwise pieces are added
	// to the existing assembling set. Blocks until Curio reports the SP
	// transfer is complete and the on-chain tx confirms (giving us the
	// dataSetId), or the configured timeout elapses. The returned
	// DataSetID is the set the pieces landed in.
	//
	// evmSigner is the client's secp256k1 wallet; its EVMAddress is the
	// on-chain payer, and its raw key (via signer.EVMSigner.ECDSAKey)
	// is used for the EIP-712 extraData signing.
	PullPiecesToFWSS(
		ctx context.Context,
		evmSigner signer.EVMSigner,
		provider string,
		pieces []PDPPieceInput,
		cfg PDPSchedulingConfig,
	) (PDPPullResult, error)
}
