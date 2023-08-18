package pack

import (
	"bytes"
	"context"
	"io"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack/encryption"
	"github.com/data-preservation-programs/singularity/storagesystem"
	util2 "github.com/data-preservation-programs/singularity/util"
	"github.com/ipfs/boxo/util"
	"github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	packJob "github.com/ipfs/go-ipfs-chunker"
	"github.com/ipfs/go-ipld-format"
	"github.com/ipfs/go-merkledag"
	"github.com/ipfs/go-unixfs"
	"github.com/ipfs/go-unixfs/pb"
	"github.com/multiformats/go-varint"
	"github.com/rclone/rclone/fs"
)

const ChunkSize int64 = 1 << 20
const NumLinkPerNode = 1024

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

func Min(i int, i2 int) int {
	if i < i2 {
		return i
	}
	return i2
}

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
		return nil, nil, errors.New("links must be more than 1")
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
	headerBytes, err := util2.GenerateCarHeader(root)
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
	varintBytes := varint.ToUvarint(uint64(len(block.RawData()) + block.Cid().ByteLen()))
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

type BlockResult struct {
	// Offset is the offset of the block in the potentially encrypted stream
	Offset int64
	// Raw is the block data which is potentially encrypted
	Raw []byte
	// CID is the CID of the block
	CID   cid.Cid
	Error error
}

var ErrFileModified = errors.New("file has been modified")

// GetBlockStreamFromFileRange reads a file (or a part of a file) identified by fileRange
// from a specified data source. Optionally, it applies encryption to the file's content.
// It then streams the resulting data blocks to the caller through a Go channel.
//
// Parameters:
//   - ctx: A context.Context used to control the lifecycle of the function.
//   - handler: A storagesystem.Reader interface implementation that is capable of reading files from a data source.
//   - fileRange: A model.FileRange struct that specifies the file to read and the range of bytes to read.
//   - encryptor: An encryption.Encryptor interface implementation that is capable of encrypting the stream.
//     If this is nil, the file's content is streamed without encryption.
//
// Returns:
// - A channel that the caller can range over to receive data blocks from the file.
// - An fs.Object representing the metadata of the file being read.
// - An error if any error occurs while processing.
//
// Note:
//   - If encryption is requested (i.e., encryptor is not nil), partial reads (i.e., reading a subrange of the file)
//     are not supported and the function will return an error in this case.
//   - The function is designed to be used concurrently, as it runs a goroutine to read blocks from the file.
func GetBlockStreamFromFileRange(ctx context.Context,
	handler storagesystem.Reader,
	fileRange model.FileRange,
	encryptor encryption.Encryptor) (<-chan BlockResult, fs.Object, error) {
	if encryptor != nil && (fileRange.Offset != 0 || fileRange.Length != fileRange.File.Size) {
		return nil, nil, errors.New("encryption is not supported for partial reads")
	}
	readStream, object, err := handler.Read(ctx, fileRange.File.Path, fileRange.Offset, fileRange.Length)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to open stream for %s at %d with length %d", fileRange.File.Path, fileRange.Offset, fileRange.Length)
	}

	if object != nil {
		same, detail := storagesystem.IsSameEntry(ctx, *fileRange.File, object)
		if !same {
			return nil, nil, errors.Wrapf(ErrFileModified, "fileRange has been modified: %s, %s", fileRange.File.Path, detail)
		}
	}

	var readCloser io.ReadCloser
	if encryptor == nil {
		readCloser = readStream
	} else {
		readCloser, err = encryptor.Encrypt(readStream)
	}
	if err != nil {
		return nil, object, errors.WithStack(err)
	}
	blockChan := make(chan BlockResult)
	chunker := packJob.NewSizeSplitter(readCloser, ChunkSize)
	go func() {
		defer close(blockChan)
		if readStream != readCloser {
			defer readStream.Close()
		}
		defer readCloser.Close()
		offset := fileRange.Offset
		firstChunk := true
		for {
			if ctx.Err() != nil {
				return
			}
			chunkerBytes, err := chunker.NextBytes()
			var result BlockResult
			if err != nil && !(errors.Is(err, io.EOF) && firstChunk) {
				if errors.Is(err, io.EOF) {
					return
				}
				result = BlockResult{Error: errors.Wrap(err, "failed to read chunk")}
			} else {
				firstChunk = false
				hash := util.Hash(chunkerBytes)
				c := cid.NewCidV1(cid.Raw, hash)
				result = BlockResult{
					CID:    c,
					Offset: offset,
					Raw:    chunkerBytes,
					Error:  nil,
				}
				offset += int64(len(chunkerBytes))
			}
			select {
			case <-ctx.Done():
				return
			case blockChan <- result:
			}
		}
	}()

	return blockChan, object, errors.WithStack(err)
}
