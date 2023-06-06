package daggen

import (
	"context"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack"
	"github.com/fxamacker/cbor"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-blockservice"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-ipfs-blockstore"
	format "github.com/ipfs/go-ipld-format"
	"github.com/ipfs/go-merkledag"
	uio "github.com/ipfs/go-unixfs/io"
	"github.com/pkg/errors"
)

func ResolveDirectoryTree(currentID uint64,
	dirCache map[uint64]*DirectoryData,
	childrenCache map[uint64][]uint64,
) (*format.Link, error) {
	current, ok := dirCache[currentID]
	if !ok {
		return nil, errors.Errorf("no directory data for current %d", currentID)
	}
	children, _ := childrenCache[currentID]

	for _, child := range children {
		link, err := ResolveDirectoryTree(child, dirCache, childrenCache)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to resolve child %d", child)
		}
		err = current.AddItem(link.Name, link.Cid, link.Size)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to add child %d to directory", child)
		}
	}

	node, err := current.dir.GetNode()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Node from directory")
	}
	size, err := node.Size()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get size from Node")
	}
	current.Node = node
	return &format.Link{
		Name: current.Directory.Name,
		Size: size,
		Cid:  node.Cid(),
	}, nil
}

type DirectoryData struct {
	Directory model.Directory
	Node      format.Node
	dir       uio.Directory
	bstore    blockstore.Blockstore
}

type serialized struct {
	Root   []byte
	Blocks [][2][]byte
}

func NewDirectoryData() DirectoryData {
	ds := datastore.NewMapDatastore()
	bs := blockstore.NewBlockstore(ds)
	dagServ := merkledag.NewDAGService(blockservice.New(bs, nil))
	dir := uio.NewDirectory(dagServ)
	dir.SetCidBuilder(merkledag.V1CidPrefix())
	return DirectoryData{
		dir:    dir,
		bstore: bs,
	}
}

func (d *DirectoryData) AddItem(name string, c cid.Cid, length uint64) error {
	return d.dir.AddChild(context.Background(), name, NewDummyNode(length, c))
}

func (d *DirectoryData) AddItemFromLinks(name string, links []format.Link) error {
	ctx := context.Background()
	blks, node, err := pack.AssembleItemFromLinks(links)
	if err != nil {
		return errors.Wrap(err, "failed to assemble item from links")
	}
	err = d.dir.AddChild(ctx, name, node)
	if err != nil {
		return errors.Wrap(err, "failed to add child to directory")
	}
	err = d.bstore.PutMany(ctx, blks)
	if err != nil {
		return errors.Wrap(err, "failed to put blocks into blockstore")
	}
	return nil
}

func (d *DirectoryData) MarshalBinary() ([]byte, error) {
	root, err := d.dir.GetNode()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get root Node")
	}
	s := serialized{
		Root: root.Cid().Bytes(),
	}
	ctx := context.Background()
	err = d.bstore.Put(ctx, root)
	if err != nil {
		return nil, errors.Wrap(err, "failed to put root Node into blockstore")
	}
	d.bstore.HashOnRead(false)
	ch, err := d.bstore.AllKeysChan(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all keys from blockstore")
	}
	for k := range ch {
		data, err := d.bstore.Get(ctx, k)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get data from blockstore")
		}
		s.Blocks = append(s.Blocks, [2][]byte{k.Bytes(), data.RawData()})
	}
	result, err := cbor.Marshal(s, cbor.CanonicalEncOptions())
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal data")
	}
	return result, nil
}

func (d *DirectoryData) UnmarshallBinary(data []byte) error {
	ds := datastore.NewMapDatastore()
	bs := blockstore.NewBlockstore(ds)
	dagServ := merkledag.NewDAGService(blockservice.New(bs, nil))
	if len(data) == 0 {
		dir := uio.NewDirectory(dagServ)
		dir.SetCidBuilder(merkledag.V1CidPrefix())
		node, err := dir.GetNode()
		if err != nil {
			return errors.Wrap(err, "failed to get Node from directory")
		}
		*d = DirectoryData{
			dir:    dir,
			bstore: bs,
			Node:   node,
		}
		return nil
	}

	ctx := context.Background()
	var s serialized
	err := cbor.Unmarshal(data, &s)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal data")
	}
	dirCID := cid.MustParse(s.Root)
	for _, b := range s.Blocks {
		blk, err := blocks.NewBlockWithCid(b[1], cid.MustParse(b[0]))
		if err != nil {
			return errors.Wrap(err, "failed to create block")
		}
		err = bs.Put(ctx, blk)
		if err != nil {
			return errors.Wrap(err, "failed to put data into blockstore")
		}
	}
	dirNode, err := dagServ.Get(ctx, dirCID)
	if err != nil {
		return errors.Wrap(err, "failed to get root Node")
	}
	dir, err := uio.NewDirectoryFromNode(dagServ, dirNode)
	if err != nil {
		return errors.Wrap(err, "failed to create directory from Node")
	}
	dir.SetCidBuilder(merkledag.V1CidPrefix())
	*d = DirectoryData{
		dir:    dir,
		bstore: bs,
		Node:   dirNode,
	}
	return nil
}
