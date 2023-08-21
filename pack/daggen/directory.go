package daggen

import (
	"bufio"
	"bytes"
	"context"
	"io"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	util2 "github.com/data-preservation-programs/singularity/pack/util"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-blockservice"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-ipfs-blockstore"
	format "github.com/ipfs/go-ipld-format"
	"github.com/ipfs/go-merkledag"
	uio "github.com/ipfs/go-unixfs/io"
	"github.com/ipld/go-car"
	"github.com/ipld/go-car/util"
	"github.com/klauspost/compress/zstd"
)

var encoder, _ = zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedDefault))
var decoder, _ = zstd.NewReader(nil, zstd.WithDecoderConcurrency(0))

type DirectoryDetail struct {
	Dir  *model.Directory
	Data *DirectoryData
}

type DirectoryTree struct {
	cache         map[uint64]*DirectoryDetail
	childrenCache map[uint64][]uint64 // This is known children for this pack only
}

func NewDirectoryTree() DirectoryTree {
	return DirectoryTree{
		cache:         make(map[uint64]*DirectoryDetail),
		childrenCache: make(map[uint64][]uint64),
	}
}

func (t DirectoryTree) Cache() map[uint64]*DirectoryDetail {
	return t.cache
}

func (t DirectoryTree) Has(dirID uint64) bool {
	_, ok := t.cache[dirID]
	return ok
}

func (t DirectoryTree) Get(dirID uint64) *DirectoryDetail {
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
func (t DirectoryTree) Resolve(ctx context.Context, dirID uint64) (*format.Link, error) {
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
//	dir       : The current representation of the directory, implementing the uio.Directory interface.
//	bstore    : The blockstore used to store and retrieve blocks of data associated with the directory.
//	node      : The cached format.Node representation of the current directory.
//	nodeDirty : A flag indicating whether the cached node representation is potentially outdated
//	            and needs to be refreshed from the internal directory representation.
type DirectoryData struct {
	dir       uio.Directory
	bstore    blockstore.Blockstore
	node      format.Node
	nodeDirty bool
}

// Node retrieves the format.Node representation of the current DirectoryData.
// If the node representation is marked as dirty (meaning it is potentially outdated),
// this method:
// 1. Calls GetNode on the internal directory to refresh the node representation.
// 2. Updates the internal node field with this new node.
// 3. Resets the dirty flag to false, indicating that the node is now up to date.
//
// Returns:
//
//	format.Node : The current node representation of the directory, or nil if an error occurs.
//	error       : An error is returned if getting the Node from the internal directory fails.
//	              Otherwise, it returns nil.
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
// 1. Creates a new in-memory map datastore.
// 2. Initializes a new blockstore with the created datastore.
// 3. Initializes a new DAG service with the blockstore.
// 4. Creates a new directory with the DAG service and sets its CID (Content Identifier) builder.
//
// Returns:
//
//	DirectoryData : A new DirectoryData instance with the initialized directory, blockstore, and a dirty node flag set to true.
func NewDirectoryData() DirectoryData {
	ds := datastore.NewMapDatastore()
	bs := blockstore.NewBlockstore(ds, blockstore.WriteThrough())
	dagServ := merkledag.NewDAGService(blockservice.New(bs, nil))
	dir := uio.NewDirectory(dagServ)
	dir.SetCidBuilder(merkledag.V1CidPrefix())
	return DirectoryData{
		dir:       dir,
		bstore:    bs,
		nodeDirty: true,
	}
}

// AddFile adds a new file to the directory with the specified name, content identifier (CID), and length.
// It creates a new dummy node with the provided length and CID, and then adds this node as a child
// to the current directory under the given name.
//
// Parameters:
//
//	ctx    : Context used to control cancellations or timeouts.
//	name   : Name of the file to be added to the directory.
//	c      : Content Identifier (CID) of the file to be added.
//	length : The length of the file in bytes.
//
// Returns:
//
//	error  : An error is returned if adding the child to the directory fails, otherwise it returns nil.
func (d *DirectoryData) AddFile(ctx context.Context, name string, c cid.Cid, length uint64) error {
	return d.dir.AddChild(ctx, name, NewDummyNode(length, c))
}

// AddFileFromLinks constructs a new file from a set of links and adds it to the directory.
// It first assembles the file from the provided links, then adds this file as a child to
// the current directory with the specified name. The assembled file and its constituent
// blocks are stored in the associated blockstore.
//
// Parameters:
//
//	ctx   : Context used to control cancellations or timeouts.
//	name  : Name of the file to be added to the directory.
//	links : Slice of format.Link that define the file to be assembled and added.
//
// Returns:
//
//	cid.Cid : Content Identifier (CID) of the added file if successful.
//	error   : An error is returned if assembling the file from links fails,
//	          adding the child to the directory fails, or putting blocks into the blockstore fails.
//	          Otherwise, it returns nil.
func (d *DirectoryData) AddFileFromLinks(ctx context.Context, name string, links []format.Link) (cid.Cid, error) {
	blks, node, err := util2.AssembleFileFromLinks(links)
	if err != nil {
		return cid.Undef, errors.WithStack(err)
	}
	err = d.dir.AddChild(ctx, name, node)
	if err != nil {
		return cid.Undef, errors.WithStack(err)
	}
	err = d.bstore.PutMany(ctx, blks)
	if err != nil {
		return cid.Undef, errors.WithStack(err)
	}
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
func (d *DirectoryData) AddBlocks(ctx context.Context, blks []blocks.Block) error {
	return d.bstore.PutMany(ctx, blks)
}

// MarshalBinary serializes the current state of the DirectoryData object into a binary format.
// This method:
//  1. Refreshes the internal representation of the directory (Node).
//  2. Writes the CAR (Content Addressable Archives) header of the new Node to a buffer.
//  3. If an old Node exists, it deletes the old Node from the blockstore.
//  4. Puts the new Node into the blockstore.
//  5. Iterates through all the keys in the blockstore, retrieves the corresponding data,
//     and writes it as CAR blocks to the buffer.
//  6. Returns the entire buffer content encoded.
//
// Parameters:
//
//	ctx : Context used to control cancellations or timeouts.
//
// Returns:
//
//	[]byte : Binary representation of the DirectoryData, or nil if an error occurs.
//	error  : An error is returned if refreshing the Node, writing the CAR header, deleting the old Node,
//	         putting
func (d *DirectoryData) MarshalBinary(ctx context.Context) ([]byte, error) {
	buf := &bytes.Buffer{}
	oldNode := d.node
	newNode, err := d.Node()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	_, err = util2.WriteCarHeader(buf, newNode.Cid())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if oldNode != nil {
		err = d.bstore.DeleteBlock(ctx, oldNode.Cid())
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}
	err = d.bstore.Put(ctx, newNode)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	ch, err := d.bstore.AllKeysChan(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for k := range ch {
		data, err := d.bstore.Get(ctx, k)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		_, err = util2.WriteCarBlock(buf, data)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return encoder.EncodeAll(buf.Bytes(), make([]byte, 0, len(buf.Bytes()))), nil
}

// UnmarshalToBlocks deserializes binary data into a set of blocks and a root content identifier (CID).
// This function:
// 1. Decodes the input binary data.
// 2. Reads the CAR (Content Addressable Archives) header from the decoded data to obtain the root CID.
// 3. Iteratively reads CAR blocks from the data and constructs block objects from them.
//
// Parameters:
//
//	data : Binary data representing a serialized set of blocks and a root CID.
//
// Returns:
//
//	cid.Cid     : The root CID extracted from the CAR header, or an undefined CID if an error occurs.
//	[]blocks.Block : Slice of blocks.Block objects reconstructed from the input data, or nil if an error occurs.
//	error       : An error is returned if decoding the input data, reading the CAR header, or reading CAR blocks fails.
//	              Otherwise, it returns nil.
func UnmarshalToBlocks(data []byte) (cid.Cid, []blocks.Block, error) {
	if len(data) == 0 {
		return cid.Undef, nil, nil
	}
	decoded, err := decoder.DecodeAll(data, nil)
	if err != nil {
		return cid.Undef, nil, errors.WithStack(err)
	}
	reader := bufio.NewReader(bytes.NewReader(decoded))
	ch, err := car.ReadHeader(reader)
	if err != nil {
		return cid.Undef, nil, errors.WithStack(err)
	}
	dirCID := ch.Roots[0]
	var blks []blocks.Block
	for {
		c, data, err := util.ReadNode(reader)
		if err != nil && !errors.Is(err, io.EOF) {
			return cid.Undef, nil, errors.WithStack(err)
		}
		if errors.Is(err, io.EOF) {
			break
		}
		blk, _ := blocks.NewBlockWithCid(data, c)
		blks = append(blks, blk)
	}
	return dirCID, blks, nil
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
//	data : Binary data representing a serialized DirectoryData object.
//
// Returns:
//
//	error : An error is returned if unmarshalling the data, putting blocks into blockstore,
//	        retrieving the root directory node, or creating a new directory from the node fails.
//	        Otherwise, it returns nil.
func (d *DirectoryData) UnmarshalBinary(ctx context.Context, data []byte) error {
	ds := datastore.NewMapDatastore()
	bs := blockstore.NewBlockstore(ds, blockstore.WriteThrough())
	bs.HashOnRead(false)
	dagServ := merkledag.NewDAGService(blockservice.New(bs, nil))
	if len(data) == 0 {
		dir := uio.NewDirectory(dagServ)
		dir.SetCidBuilder(merkledag.V1CidPrefix())
		*d = DirectoryData{
			dir:       dir,
			bstore:    bs,
			nodeDirty: true,
		}
		return nil
	}

	dirCID, blks, err := UnmarshalToBlocks(data)
	if err != nil {
		return errors.WithStack(err)
	}
	err = bs.PutMany(ctx, blks)
	if err != nil {
		return errors.WithStack(err)
	}
	dirNode, err := dagServ.Get(ctx, dirCID)
	if err != nil {
		return errors.WithStack(err)
	}
	dir, err := uio.NewDirectoryFromNode(dagServ, dirNode)
	if err != nil {
		return errors.WithStack(err)
	}
	dir.SetCidBuilder(merkledag.V1CidPrefix())
	*d = DirectoryData{
		dir:       dir,
		bstore:    bs,
		node:      nil,
		nodeDirty: true,
	}
	return nil
}
