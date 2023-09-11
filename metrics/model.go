package metrics

type EventType int

const (
	PackJobComplete EventType = iota
	DealProposal
)

type Events struct {
	_             struct{} `cbor:",toarray"`
	PackJobEvents []PackJobEvent
	DealEvents    []DealProposalEvent
}

type PackJobEvent struct {
	_         struct{} `cbor:",toarray"`
	Timestamp int64
	Instance  string

	SourceType string
	OutputType string
	PieceSize  int64
	PieceCID   string
	CarSize    int64
	NumOfFiles int64
}

type DealProposalEvent struct {
	_         struct{} `cbor:",toarray"`
	Timestamp int64
	Instance  string

	PieceCID   string
	DataCID    string
	PieceSize  int64
	Provider   string
	Client     string
	Verified   bool
	StartEpoch int32
	EndEpoch   int32
}
