package proposal110

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
)

type DataRef struct {
	TransferType string
	Root         cid.Cid

	PieceCid     *cid.Cid              // Optional for non-manual transfer, will be recomputed from the data if not given
	PieceSize    abi.UnpaddedPieceSize // Optional for non-manual transfer, will be recomputed from the data if not given
	RawBlockSize uint64                // Optional: used as the denominator when calculating transfer %
}
