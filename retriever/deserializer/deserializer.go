package deserializer

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/ipfs/go-unixfsnode/file"
	"github.com/ipld/go-car/v2"
	dagpb "github.com/ipld/go-codec-dagpb"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/linking"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"go.uber.org/multierr"
)

var (
	// ErrMalformedCar indicates an error while trying to read a block from a car file
	ErrMalformedCar = errors.New("malformed CAR")
	// ErrUnexpectedBlock means we read the next block in the CAR file but it wasn't the one we were expecting
	ErrUnexpectedBlock = errors.New("unexpected block in CAR")
)

// Deserialize is a function to deserialize a UnixFS file byte range from a CAR block reader
func Deserialize(ctx context.Context, r *car.BlockReader, c cid.Cid, start int64, end int64, out io.Writer) (int64, error) {
	// set up an ipld prime link system to read blocks from the car
	lsys := cidlink.DefaultLinkSystem()
	lsys.TrustedStorage = true
	lsys.StorageReadOpener = func(lc linking.LinkContext, l datamodel.Link) (io.Reader, error) {
		data, err := readNextBlock(ctx, r, l.(cidlink.Link).Cid)
		if err != nil {
			return nil, err
		}
		return bytes.NewReader(data), nil
	}
	// load the first new node
	node, err := loadNode(ctx, c, lsys)
	if err != nil {
		return 0, fmt.Errorf("deserializing, unable to load root node: %w", err)
	}
	// reify it as unixfs
	fnode, err := file.NewUnixFSFile(ctx, node, &lsys)
	if err != nil {
		return 0, fmt.Errorf("deserializing, reifying as unix fs: %w", err)
	}
	// read the byte range
	nlr, err := fnode.AsLargeBytes()
	if err != nil {
		return 0, fmt.Errorf("deserializing, reading as large bytes node: %w", err)
	}
	_, err = nlr.Seek(start, io.SeekStart)
	if err != nil {
		return 0, fmt.Errorf("deserializing, seeking to start of range: %w", err)
	}
	rangeLeftReader := io.LimitReader(nlr, end-start)
	written, err := io.Copy(out, rangeLeftReader)
	if err != nil {
		return written, fmt.Errorf("deserializing, reading file: %w", err)
	}
	return written, nil
}

var protoChooser = dagpb.AddSupportToChooser(basicnode.Chooser)

// load an individual node (required to start a selector traversal)
func loadNode(ctx context.Context, rootCid cid.Cid, lsys linking.LinkSystem) (datamodel.Node, error) {
	lnk := cidlink.Link{Cid: rootCid}
	lnkCtx := linking.LinkContext{Ctx: ctx}
	proto, err := protoChooser(lnk, lnkCtx)
	if err != nil {
		return nil, fmt.Errorf("failed to choose prototype for CID %s: %w", rootCid.String(), err)
	}
	rootNode, err := lsys.Load(lnkCtx, lnk, proto)
	if err != nil {
		return nil, fmt.Errorf("failed to load root CID: %w", err)
	}
	return rootNode, nil
}

// read the next block from a car reader, verifying it matches the next expected block in traversal
func readNextBlock(ctx context.Context, bs *car.BlockReader, expected cid.Cid) ([]byte, error) {
	blk, err := bs.Next()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil, format.ErrNotFound{Cid: expected}
		}
		return nil, multierr.Combine(ErrMalformedCar, err)
	}

	// compare by multihash only
	if !bytes.Equal(blk.Cid().Hash(), expected.Hash()) {
		return nil, fmt.Errorf("%w: %s != %s", ErrUnexpectedBlock, blk.Cid(), expected)
	}

	return blk.RawData(), nil
}
