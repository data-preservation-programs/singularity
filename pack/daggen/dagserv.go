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

// Get retrieves a format.Node based on a given CID from the RecordedDagService.
//
// The RecordedDagService keeps track of blocks with their visitation status.
// If a block with the provided CID is found and hasn't been visited yet, it's marked as visited.
// If the block is a dummy block, a dummy node is returned; otherwise, the raw data of the block is decoded
// into a format.Node using the IPLD legacy decoder.
//
// Parameters:
//   - ctx: A context to allow for timeout or cancellation of operations.
//   - c: A CID representing the content identifier of the desired block.
//
// Returns:
//   - A format.Node representing the block associated with the given CID.
//     This could be a dummy node or a decoded version of the actual block data.
//   - An error if the CID does not match any blocks in the service's blockstore or if other issues arise.
//     If the CID is not found, format.ErrNotFound with the CID is returned.
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

// Visit marks a block with a given CID as visited in the RecordedDagService's blockstore.
//
// The purpose of this function is to track the visitation status of blocks.
// If a block with the provided CID exists in the blockstore and hasn't been visited yet,
// this method will mark it as visited. If the block is already marked as visited or doesn't exist,
// the method does nothing.
//
// Parameters:
//   - ctx: A context to allow for timeout or cancellation of operations.
//   - c: A CID representing the content identifier of the block to be marked as visited.
func (r *RecordedDagService) Visit(ctx context.Context, c cid.Cid) {
	if data, ok := r.blockstore[c]; ok {
		if !data.visited {
			data.visited = true
			r.blockstore[c] = data
		}
	}
}

// Add adds a block to the RecordedDagService's blockstore.
//
// If the node is of type *DummyNode, it is added to the blockstore with its size and marked as a dummy.
// Otherwise, the node's raw data is added to the blockstore.
//
// Parameters:
//   - ctx: A context to allow for timeout or cancellation of operations.
//   - node: The format.Node representing the block to be added to the blockstore.
//
// Returns:
//   - An error if any issues arise during the addition process; otherwise, it returns nil.
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

// GetMany fetches multiple nodes from the RecordedDagService's blockstore concurrently.
//
// This method returns a channel where the caller can consume the results.
// Each result is wrapped in a *format.NodeOption which contains the format.Node
// and an error, if any occurred while fetching the node.
//
// Parameters:
//   - ctx: A context to allow for timeout or cancellation of operations.
//   - cids: A slice of cid.Cid representing the identifiers of the nodes to be fetched.
//
// Returns:
//   - A channel of *format.NodeOption. Each *format.NodeOption contains a fetched node
//     and an error indicating any issues that occurred while fetching.
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
