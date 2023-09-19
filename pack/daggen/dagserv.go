package daggen

import (
	"context"

	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	legacy "github.com/ipfs/go-ipld-legacy"
	"github.com/ipfs/go-merkledag"
	dagpb "github.com/ipld/go-codec-dagpb"
	_ "github.com/ipld/go-ipld-prime/codec/raw"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
)

var ipldLegacyDecoder *legacy.Decoder

// TODO: Don't require global registries
func init() {
	d := legacy.NewDecoder()
	d.RegisterCodec(cid.DagProtobuf, dagpb.Type.PBNode, merkledag.ProtoNodeConverter)
	d.RegisterCodec(cid.Raw, basicnode.Prototype.Bytes, merkledag.RawNodeConverter)
	ipldLegacyDecoder = d
}

type blockData struct {
	raw     []byte
	dummy   bool
	size    uint32
	visited bool
}

// RecordedDagService is a DAGService that records all blocks that are used.
// This struct is only meant to be used internally which achieves
// * Tracks the blocks that are used
// * Directly supports DummyNode
type RecordedDagService struct {
	blockstore map[cid.Cid]blockData
}

type DagServEntry struct {
	_         struct{} `cbor:",toarray"`
	Cid       cid.Cid
	Data      []byte
	DummySize int32
}

var _ format.DAGService = &RecordedDagService{}

func NewRecordedDagService() *RecordedDagService {
	return &RecordedDagService{
		blockstore: make(map[cid.Cid]blockData),
	}
}

func (r *RecordedDagService) ResetVisited() {
	for c, data := range r.blockstore {
		data.visited = false
		r.blockstore[c] = data
	}
}

func (r *RecordedDagService) Get(ctx context.Context, c cid.Cid) (format.Node, error) {
	if data, ok := r.blockstore[c]; ok {
		if !data.visited {
			data.visited = true
			r.blockstore[c] = data
		}
		if data.dummy {
			return NewDummyNode(uint64(data.size), c), nil
		}
		blk, _ := blocks.NewBlockWithCid(data.raw, c)
		return ipldLegacyDecoder.DecodeNode(ctx, blk)
	}
	return nil, format.ErrNotFound{Cid: c}
}

func (r *RecordedDagService) Visit(ctx context.Context, c cid.Cid) {
	if data, ok := r.blockstore[c]; ok {
		if !data.visited {
			data.visited = true
			r.blockstore[c] = data
		}
	}
}

func (r *RecordedDagService) Add(ctx context.Context, node format.Node) error {
	dummy, ok := node.(*DummyNode)
	if ok {
		r.blockstore[dummy.cid] = blockData{
			size:  uint32(dummy.size),
			dummy: true,
		}
		return nil
	}
	r.blockstore[node.Cid()] = blockData{
		raw: node.RawData(),
	}
	return nil
}

func (r *RecordedDagService) GetMany(ctx context.Context, cids []cid.Cid) <-chan *format.NodeOption {
	resultChan := make(chan *format.NodeOption, len(cids))
	go func() {
		defer close(resultChan)
		for _, c := range cids {
			node, err := r.Get(ctx, c)
			resultChan <- &format.NodeOption{
				Node: node,
				Err:  err,
			}
		}
	}()
	return resultChan
}

func (r *RecordedDagService) AddMany(ctx context.Context, nodes []format.Node) error {
	panic("implement me")
}

func (r *RecordedDagService) Remove(ctx context.Context, c cid.Cid) error {
	panic("implement me")
}

func (r *RecordedDagService) RemoveMany(ctx context.Context, cids []cid.Cid) error {
	panic("implement me")
}
