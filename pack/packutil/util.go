package packutil

import (
	"bytes"
	"io"
	"math"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/ipfs/boxo/ipld/merkledag"
	"github.com/ipfs/boxo/ipld/unixfs"
	unixfs_pb "github.com/ipfs/boxo/ipld/unixfs/pb"
	util2 "github.com/ipfs/boxo/util"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/multiformats/go-varint"
)

var EmptyFileCid = cid.NewCidV1(cid.Raw, util2.Hash([]byte("")))

// safeIntToUint64 safely converts int to uint64, handling negative values
func safeIntToUint64(val int) uint64 {
	if val < 0 {
		return 0
	}
	if val > math.MaxInt {
		return math.MaxUint64
	}
	return uint64(val)
}

var EmptyFileVarint = varint.ToUvarint(uint64(len(EmptyFileCid.Bytes())))

var EmptyCarHeader, _ = util.GenerateCarHeader(EmptyFileCid)

const (
	ChunkSize      int64 = 1 << 20
	NumLinkPerNode       = 1024
)

func Min(i int, i2 int) int {
	if i < i2 {
		return i
	}
	return i2
}

var errLinkLessThanTwo = errors.New("links must be more than 1")

// AssembleFileFromLinks constructs a MerkleDAG from a list of links.
// It organizes the links into a tree structure where each internal node
// can have up to NumLinkPerNode children. This function assembles the DAG
// and returns the blocks that make up the DAG and the root node of the DAG.
//
// Parameters:
//   - links: An array of format.Link objects representing the links to the
//     content that will be part of the MerkleDAG.
//
// Returns:
//   - []blocks.Block: A slice of Block objects representing the blocks of
//     the constructed MerkleDAG. Each block contains a part of the overall data.
//   - *merkledag.ProtoNode: The root node of the created MerkleDAG. This node
//     provides a starting point to navigate through the rest of the DAG.
//   - error: An error that can occur during the MerkleDAG creation process,
//     or nil if the operation was successful.
func AssembleFileFromLinks(links []format.Link) ([]blocks.Block, *merkledag.ProtoNode, error) {
	if len(links) <= 1 {
		return nil, nil, errLinkLessThanTwo
	}
	result := make([]blocks.Block, 0)
	var rootNode *merkledag.ProtoNode
	for len(links) > 1 {
		newLinks := make([]format.Link, 0)
		for start := 0; start < len(links); start += NumLinkPerNode {
			newNode, total, err := createParentNode(links[start:Min(start+NumLinkPerNode, len(links))])
			if err != nil {
				return nil, nil, errors.WithStack(err)
			}

			basicBlock, err := blocks.NewBlockWithCid(newNode.RawData(), newNode.Cid())
			if err != nil {
				return nil, nil, errors.WithStack(err)
			}
			result = append(result, basicBlock)
			newLinks = append(
				newLinks, format.Link{
					Name: "",
					Size: total,
					Cid:  newNode.Cid(),
				},
			)
			rootNode = newNode
		}

		links = newLinks
	}
	return result, rootNode, nil
}

// WriteCarHeader writes the CAR (Content Addressable aRchive) format header to a given io.Writer.
//
// Parameters:
//   - writer: An io.Writer to which the CAR header will be written.
//   - root: The root CID (Content Identifier) for the CAR file.
//
// Returns:
//   - []byte: The byte representation of the CAR header.
//   - error: An error that can occur during the write process, or nil if the write was successful.
func WriteCarHeader(writer io.Writer, root cid.Cid) ([]byte, error) {
	headerBytes, err := util.GenerateCarHeader(root)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	_, err = io.Copy(writer, bytes.NewReader(headerBytes))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return headerBytes, nil
}

// WriteCarBlock writes a block in CAR (Content Addressable aRchive) format to a given io.Writer.
//
// Parameters:
//   - writer: An io.Writer to which the CAR-formatted block will be written.
//   - block: A blocks.Block instance representing the data to be written.
//
// Returns:
//   - int64: The number of bytes written to the writer.
//   - error: An error that can occur during the write process, or nil if the write was successful.
func WriteCarBlock(writer io.Writer, block blocks.Block) (int64, error) {
	written := int64(0)
	varintBytes := varint.ToUvarint(safeIntToUint64(len(block.RawData()) + block.Cid().ByteLen()))
	n, err := io.Copy(writer, bytes.NewReader(varintBytes))
	if err != nil {
		return written, errors.WithStack(err)
	}
	written += n

	n, err = io.Copy(writer, bytes.NewReader(block.Cid().Bytes()))
	if err != nil {
		return written, errors.WithStack(err)
	}
	written += n

	n, err = io.Copy(writer, bytes.NewReader(block.RawData()))
	if err != nil {
		return written, errors.WithStack(err)
	}
	written += n
	return written, nil
}

// createParentNode creates a new parent ProtoNode for a given set of links.
// It constructs a UnixFS node with the type Data_File and adds the sizes of
// the links as block sizes to this UnixFS node. It then creates a new ProtoNode
// with the UnixFS node's data and adds the links to this ProtoNode.
//
// Parameters:
//   - links: An array of format.Link objects. These links will be added as child
//     links to the new ProtoNode.
//
// Returns:
//   - *merkledag.ProtoNode: A pointer to the new parent ProtoNode that has been
//     created. This node contains the data of the UnixFS node and the child links.
//   - uint64: The total size of the data that the new parent node represents. This
//     is the sum of the sizes of all the links.
//   - error: An error that can occur during the creation of the new parent node, or
//     nil if the operation was successful.
func createParentNode(links []format.Link) (*merkledag.ProtoNode, uint64, error) {
	node := unixfs.NewFSNode(unixfs_pb.Data_File)
	total := uint64(0)
	for _, link := range links {
		node.AddBlockSize(link.Size)
		total += link.Size
	}
	nodeBytes, err := node.GetBytes()
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}
	pbNode := merkledag.NodeWithData(nodeBytes)
	err = pbNode.SetCidBuilder(merkledag.V1CidPrefix())
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}
	for i := range links {
		err = pbNode.AddRawLink("", &links[i])
		if err != nil {
			return nil, 0, errors.WithStack(err)
		}
	}
	return pbNode, total, nil
}
