package proposal110

import (
	"unicode/utf8"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"

	addr "github.com/filecoin-project/go-address"
)

//go:generate go run github.com/hannahhoward/cbor-gen-for --map-encoding DataRef Proposal SignedResponse Response

// DataRef is a reference for how data will be transferred for a given storage deal
type DataRef struct {
	TransferType string
	Root         cid.Cid

	PieceCid     *cid.Cid              // Optional for non-manual transfer, will be recomputed from the data if not given
	PieceSize    abi.UnpaddedPieceSize // Optional for non-manual transfer, will be recomputed from the data if not given
	RawBlockSize uint64                // Optional: used as the denominator when calculating transfer %
}

// Proposal is the data sent over the network from client to provider when proposing
// a deal
type Proposal struct {
	DealProposal  *ClientDealProposal
	Piece         *DataRef
	FastRetrieval bool
}

type ClientDealProposal struct {
	Proposal        DealProposal
	ClientSignature crypto.Signature
}

type DealProposal struct {
	PieceCID     cid.Cid `checked:"true"` // Checked in validateDeal, CommP
	PieceSize    abi.PaddedPieceSize
	VerifiedDeal bool
	Client       addr.Address
	Provider     addr.Address

	// Label is an arbitrary client chosen label to apply to the deal
	Label DealLabel

	// Nominal start epoch. Deal payment is linear between StartEpoch and EndEpoch,
	// with total amount StoragePricePerEpoch * (EndEpoch - StartEpoch).
	// Storage deal must appear in a sealed (proven) sector no later than StartEpoch,
	// otherwise it is invalid.
	StartEpoch           abi.ChainEpoch
	EndEpoch             abi.ChainEpoch
	StoragePricePerEpoch abi.TokenAmount

	ProviderCollateral abi.TokenAmount
	ClientCollateral   abi.TokenAmount
}

type DealLabel struct {
	bs        []byte
	notString bool
}

func NewLabelFromString(s string) (DealLabel, error) {
	if len(s) > DealMaxLabelSize {
		return EmptyDealLabel, xerrors.Errorf("provided string is too large to be a label (%d), max length (%d)", len(s), DealMaxLabelSize)
	}
	if !utf8.ValidString(s) {
		return EmptyDealLabel, xerrors.Errorf("provided string is invalid utf8")
	}
	return DealLabel{
		bs:        []byte(s),
		notString: false,
	}, nil
}

var EmptyDealLabel = DealLabel{}

const DealMaxLabelSize = 256

const (
	// TTGraphsync means data for a deal will be transferred by graphsync
	TTGraphsync = "graphsync"

	// TTManual means data for a deal will be transferred manually and imported
	// on the provider
	TTManual = "manual"
)

type SignedResponse struct {
	Response Response

	Signature *crypto.Signature
}
type Response struct {
	State StorageDealStatus

	// DealProposalRejected
	Message  string
	Proposal cid.Cid

	// StorageDealProposalAccepted
	PublishMessage *cid.Cid
}
type StorageDealStatus = uint64
