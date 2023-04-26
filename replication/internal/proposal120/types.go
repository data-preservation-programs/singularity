package proposal120

import (
	"github.com/filecoin-project/go-state-types/builtin/v9/market"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
)

type DealParams struct {
	DealUUID           uuid.UUID
	IsOffline          bool
	ClientDealProposal market.ClientDealProposal
	DealDataRoot       cid.Cid
	Transfer           Transfer // Transfer params will be the zero value if this is an offline deal
	RemoveUnsealedCopy bool
	SkipIPNIAnnounce   bool
}

type HttpRequest struct {
	// URL can be
	// - an http URL:
	//   "https://example.com/path"
	// - a libp2p URL:
	//   "libp2p:///ip4/104.131.131.82/tcp/4001/ipfs/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ"
	//   Must include a Peer ID
	URL string
	// Headers are the HTTP headers that are sent as part of the request,
	// eg "Authorization"
	Headers map[string]string
}

// Transfer has the parameters for a data transfer
type Transfer struct {
	// The type of transfer eg "http"
	Type string
	// An optional ID that can be supplied by the client to identify the deal
	ClientID string
	// A byte array containing marshalled data specific to the transfer type
	// eg a JSON encoded struct { URL: "<url>", Headers: {...} }
	Params []byte
	// The size of the data transferred in bytes
	Size uint64
}

type DealResponse struct {
	Accepted bool
	// Message is the reason the deal proposal was rejected. It is empty if
	// the deal was accepted.
	Message string
}
