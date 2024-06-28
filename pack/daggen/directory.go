package daggen

import (
	"bytes"
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack/packutil"
	"github.com/fxamacker/cbor/v2"
	"github.com/ipfs/boxo/ipld/merkledag"
	uio "github.com/ipfs/boxo/ipld/unixfs/io"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/klauspost/compress/zstd"
)

var compressor, _ = zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedDefault))
var decompressor, _ = zstd.NewReader(nil, zstd.WithDecoderConcurrency(0))

type DirectoryDetail struct {
	Dir  *model.Directory
	Data *DirectoryData
}

type DirectoryTree struct {
	cache         map[model.DirectoryID]*DirectoryDetail
	childrenCache map[model.DirectoryID][]model.DirectoryID // This is known children for this pack only
}

func NewDirectoryTree() DirectoryTree {
	return DirectoryTree{
		cache:         make(map[model.DirectoryID]*DirectoryDetail),
		childrenCache: make(map[model.DirectoryID][]model.DirectoryID),
	}
}

func (t DirectoryTree) Cache() map[model.DirectoryID]*DirectoryDetail {
	return t.cache
}

func (t DirectoryTree) Has(dirID model.DirectoryID) bool {
	_, ok := t.cache[dirID]
	return ok
}

func (t DirectoryTree) Get(dirID model.DirectoryID) *DirectoryDetail {
	return t.cache[dirID]
}

// Add inserts a new directory into the DirectoryTree.
//
// Parameters:
//   - ctx: Context that allows for asynchronous task cancellation.
//   - dir: A pointer to a model.Directory object that needs to be added to the tree.
//
// Returns:
//   - error: The error encountered during the operation, if any
func (t DirectoryTree) Add(ctx context.Context, dir *model.Directory) error {
	data := &DirectoryData{}
	err := data.UnmarshalBinary(ctx, dir.Data)
	if err != nil {
		return errors.WithStack(err)
	}
	if dir.ParentID != nil {
		t.childrenCache[*dir.ParentID] = append(t.childrenCache[*dir.ParentID], dir.ID)
	}
	t.cache[dir.ID] = &DirectoryDetail{
		Dir:  dir,
		Data: data,
	}
	return nil
}

// Resolve recursively constructs the IPLD (InterPlanetary Linked Data) structure for a directory and its subdirectories,
// and returns a link pointing to the root of this structure.
//
// Parameters:
//   - ctx: Context that allows for asynchronous task cancellation.
//   - dirID: The ID of the directory that needs to be resolved.
//
// Returns:
//   - *format.Link: A link that points to the root of the IPLD structure for the directory.
//   - error: The error encountered during the operation, if any.
func (t DirectoryTree) Resolve(ctx context.Context, dirID model.DirectoryID) (*format.Link, error) {
	detail, ok := t.cache[dirID]
	if !ok {
		return nil, errors.Errorf("no directory detail for dir %d", dirID)
	}

	for _, child := range t.childrenCache[dirID] {
		link, err := t.Resolve(ctx, child)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to resolve child %d", child)
		}
		err = detail.Data.AddFile(ctx, link.Name, link.Cid, link.Size)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to add child %d to directory", child)
		}
	}

	node, err := detail.Data.Node()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	size, err := node.Size()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &format.Link{
		Name: detail.Dir.Name,
		Size: size,
		Cid:  node.Cid(),
	}, nil
}

// DirectoryData represents a structured directory in a content-addressed file system.
// It manages the underlying data and provides methods for interacting with this data
// as a hierarchical directory structure.
//
// Fields:
//
//   - dir: The current representation of the directory, implementing the uio.Directory interface.
//   - bstore: The blockstore used to store and retrieve blocks of data associated with the directory.
//   - node: The cached format.Node representation of the current directory.
//   - nodeDirty : A flag indicating whether the cached node representation is potentially outdated
//     and needs to be refreshed from the internal directory representation.
type DirectoryData struct {
	dir        uio.Directory
	dagServ    *RecordedDagService
	node       format.Node
	nodeDirty  bool
	additional map[cid.Cid][]byte
}

// Node retrieves the format.Node representation of the current DirectoryData.
// If the node representation is marked as dirty (meaning it is potentially outdated),
// this method:
//  1. Calls GetNode on the internal directory to refresh the node representation.
//  2. Updates the internal node field with this new node.
//  3. Resets the dirty flag to false, indicating that the node is now up to date.
//
// Returns:
//
//   - format.Node : The current node representation of the directory, or nil if an error occurs.
//   - error       : An error is returned if getting the Node from the internal directory fails.
//     Otherwise, it returns nil.
func (d *DirectoryData) Node() (format.Node, error) {
	if d.nodeDirty {
		node, err := d.dir.GetNode()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		d.node = node
		d.nodeDirty = false
	}
	return d.node, nil
}

// NewDirectoryData creates and initializes a new DirectoryData instance.
// This function:
//  1. Creates a new in-memory map datastore.
//  2. Initializes a new blockstore with the created datastore.
//  3. Initializes a new DAG service with the blockstore.
//  4. Creates a new directory with the DAG service and sets its CID (Content Identifier) builder.
//
// Returns:
//
//   - DirectoryData : A new DirectoryData instance with the initialized directory, blockstore, and a dirty node flag set to true.
func NewDirectoryData() DirectoryData {
	dagServ := NewRecordedDagService()
	dir := uio.NewDirectory(dagServ)
	dir.SetCidBuilder(merkledag.V1CidPrefix())
	return DirectoryData{
		dir:        dir,
		nodeDirty:  true,
		dagServ:    dagServ,
		additional: make(map[cid.Cid][]byte),
	}
}

// AddFile adds a new file to the directory with the specified name, content identifier (CID), and length.
// It creates a new dummy node with the provided length and CID, and then adds this node as a child
// to the current directory under the given name.
//
// Parameters:
//
//   - ctx    : Context used to control cancellations or timeouts.
//   - name   : Name of the file to be added to the directory.
//   - c      : Content Identifier (CID) of the file to be added.
//   - length : The length of the file in bytes.
//
// Returns:
//
//	error  : An error is returned if adding the child to the directory fails, otherwise it returns nil.
func (d *DirectoryData) AddFile(ctx context.Context, name string, c cid.Cid, length uint64) error {
	d.nodeDirty = true
	node := NewDummyNode(length, c)
	_ = d.dagServ.Add(ctx, node)
	return d.dir.AddChild(ctx, name, node)
}

// AddFileFromLinks constructs a new file from a set of links and adds it to the directory.
// It first assembles the file from the provided links, then adds this file as a child to
// the current directory with the specified name. The assembled file and its constituent
// blocks are stored in the associated blockstore.
//
// Parameters:
//
//   - ctx   : Context used to control cancellations or timeouts.
//   - name  : Name of the file to be added to the directory.
//   - links : Slice of format.Link that define the file to be assembled and added.
//
// Returns:
//
//   - cid.Cid : Content Identifier (CID) of the added file if successful.
//   - error   : An error is returned if assembling the file from links fails,
//     adding the child to the directory fails, or putting blocks into the blockstore fails.
//     Otherwise, it returns nil.
func (d *DirectoryData) AddFileFromLinks(ctx context.Context, name string, links []format.Link) (cid.Cid, error) {
	blks, node, err := packutil.AssembleFileFromLinks(links)
	if err != nil {
		return cid.Undef, errors.WithStack(err)
	}
	err = d.dir.AddChild(ctx, name, node)
	if err != nil {
		return cid.Undef, errors.WithStack(err)
	}
	d.AddBlocks(ctx, blks)
	d.nodeDirty = true
	return node.Cid(), nil
}

// AddBlocks adds an array of blocks to the underlying blockstore of the DirectoryData instance.
//
// Parameters:
//   - ctx: Context that allows for asynchronous task cancellation.
//   - blks: An array of blocks that need to be added to the blockstore.
//
// Returns:
//   - error: The error encountered during the operation, if any.
//
// This function is a wrapper that delegates the block adding task to the blockstore instance
// associated with the DirectoryData instance.
func (d *DirectoryData) AddBlocks(ctx context.Context, blks []blocks.Block) {
	for _, blk := range blks {
		d.additional[blk.Cid()] = blk.RawData()
	}
}

type directoryData struct {
	_          struct{} `cbor:",toarray"`
	Root       cid.Cid
	Dummies    map[cid.Cid]uint32
	Reals      map[cid.Cid][]byte
	Additional map[cid.Cid][]byte
}

// MarshalBinary encodes the DirectoryData into a binary format using CBOR and then compresses the result.
//
// The method reconstructs the directory using the DagService, determines which blocks have been visited,
// and constructs a representation containing both dummy and real data from the directory.
//
// Parameters:
//   - ctx: A context to allow for timeout or cancellation of operations.
//
// Returns:
//   - A byte slice representing the compressed CBOR encoded binary of the DirectoryData.
//   - An error, if any occurred during the encoding or compression process.
func (d *DirectoryData) MarshalBinary(ctx context.Context) ([]byte, error) {
	buf := &bytes.Buffer{}
	newNode, err := d.Node()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	_ = d.dagServ.Add(ctx, newNode)

	// Reconstruct the directory from dagServ and figure out the visited blocks
	d.dagServ.ResetVisited()
	d.dagServ.Visit(ctx, newNode.Cid())
	d.dir, err = uio.NewDirectoryFromNode(d.dagServ, newNode)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = d.dir.ForEachLink(ctx, func(link *format.Link) error {
		d.dagServ.Visit(ctx, link.Cid)
		return nil
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	data := directoryData{
		Dummies:    make(map[cid.Cid]uint32),
		Reals:      make(map[cid.Cid][]byte),
		Additional: d.additional,
		Root:       newNode.Cid(),
	}

	for c, d := range d.dagServ.blockstore {
		if !d.visited {
			continue
		}
		if d.dummy {
			data.Dummies[c] = d.size
		} else {
			data.Reals[c] = d.raw
		}
	}
	encoder := cbor.NewEncoder(buf)
	err = encoder.Encode(data)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return compressor.EncodeAll(buf.Bytes(), make([]byte, 0, len(buf.Bytes()))), nil
}

// UnmarshalToBlocks decodes a byte slice into a slice of blocks. The input byte slice is expected to
// represent compressed CBOR encoded binary data of directoryData.
//
// The function first decompresses the input byte slice and then decodes it using CBOR to obtain
// directoryData which contains information about real and additional blocks. These blocks are then
// reconstructed and returned.
//
// Parameters:
//   - in: A byte slice representing the compressed CBOR encoded binary of the directoryData.
//
// Returns:
//   - A slice of blocks reconstructed from the input byte slice.
//   - An error, if any occurred during the decompression, decoding or block reconstruction process.
func UnmarshalToBlocks(in []byte) ([]blocks.Block, error) {
	if len(in) == 0 {
		return nil, nil
	}

	decompressed, err := decompressor.DecodeAll(in, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	decoder, err := cbor.DecOptions{
		MaxMapPairs: 2147483647,
	}.DecMode()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var data directoryData
	err = decoder.Unmarshal(decompressed, &data)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	blks := make([]blocks.Block, 0, len(data.Reals)+len(data.Additional))
	for c, d := range data.Reals {
		blk, _ := blocks.NewBlockWithCid(d, c)
		blks = append(blks, blk)
	}
	for c, d := range data.Additional {
		blk, _ := blocks.NewBlockWithCid(d, c)
		blks = append(blks, blk)
	}
	return blks, nil
}

// UnmarshalBinary deserializes binary data into the current DirectoryData object.
// This method:
//  1. Creates a new blockstore and DAG service.
//  2. Checks if the input data is empty. If it is, initializes the DirectoryData with
//     a new empty directory and returns.
//  3. Otherwise, it unmarshalls the input data into blocks and a root CID.
//  4. Puts these blocks into the blockstore.
//  5. Retrieves the root directory node from the DAG service using the root CID.
//  6. Constructs a new directory from the retrieved node and sets this as the current directory.
//
// Parameters:
//
//   - data : Binary data representing a serialized DirectoryData object.
//
// Returns:
//
//   - error : An error is returned if unmarshalling the data, putting blocks into blockstore,
//     retrieving the root directory node, or creating a new directory from the node fails.
//     Otherwise, it returns nil.
func (d *DirectoryData) UnmarshalBinary(ctx context.Context, in []byte) error {
	dagServ := NewRecordedDagService()
	if len(in) == 0 {
		dir := uio.NewDirectory(dagServ)
		dir.SetCidBuilder(merkledag.V1CidPrefix())
		*d = DirectoryData{
			dir:        dir,
			nodeDirty:  true,
			dagServ:    dagServ,
			additional: make(map[cid.Cid][]byte),
		}
		return nil
	}

	var data directoryData

	decompressed, err := decompressor.DecodeAll(in, nil)
	if err != nil {
		return errors.WithStack(err)
	}
	decoder, err := cbor.DecOptions{
		MaxMapPairs: 2147483647,
	}.DecMode()
	if err != nil {
		return errors.WithStack(err)
	}
	err = decoder.Unmarshal(decompressed, &data)
	if err != nil {
		return errors.WithStack(err)
	}

	for c, d := range data.Dummies {
		dagServ.blockstore[c] = blockData{
			dummy: true,
			size:  d,
		}
	}
	for c, d := range data.Reals {
		dagServ.blockstore[c] = blockData{
			raw: d,
		}
	}

	root, err := dagServ.Get(ctx, data.Root)
	if err != nil {
		return errors.WithStack(err)
	}
	dir, err := uio.NewDirectoryFromNode(dagServ, root)
	if err != nil {
		return errors.WithStack(err)
	}
	dir.SetCidBuilder(merkledag.V1CidPrefix())
	*d = DirectoryData{
		dir:        dir,
		node:       root,
		nodeDirty:  false,
		dagServ:    dagServ,
		additional: data.Additional,
	}
	return nil
}
