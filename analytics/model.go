package analytics

type EventType int

const (
	PackJobComplete EventType = iota
	DealProposal
)

type Events struct {
	PackJobEvents []PackJobEvent      `cbor:"1,keyasint"`
	DealEvents    []DealProposalEvent `cbor:"2,keyasint"`
}

type PackJobEvent struct {
	Timestamp int64  `cbor:"1,keyasint"`
	Instance  string `cbor:"2,keyasint"`

	SourceType string `cbor:"3,keyasint"`
	OutputType string `cbor:"4,keyasint"`
	PieceSize  int64  `cbor:"5,keyasint"`
	PieceCID   string `cbor:"6,keyasint"`
	CarSize    int64  `cbor:"7,keyasint"`
	NumOfFiles int64  `cbor:"8,keyasint"`

	Identity string `cbor:"9,keyasint"`
}

type DealProposalEvent struct {
	Timestamp int64  `cbor:"1,keyasint"`
	Instance  string `cbor:"2,keyasint"`

	PieceCID   string `cbor:"3,keyasint"`
	DataCID    string `cbor:"4,keyasint"`
	PieceSize  int64  `cbor:"5,keyasint"`
	Provider   string `cbor:"6,keyasint"`
	Client     string `cbor:"7,keyasint"`
	Verified   bool   `cbor:"8,keyasint"`
	StartEpoch int32  `cbor:"9,keyasint"`
	EndEpoch   int32  `cbor:"10,keyasint"`

	Identity string `cbor:"11,keyasint"`
}
