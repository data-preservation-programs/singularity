package pack

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack/encryption"
	"github.com/ipfs/boxo/util"
	"github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	packJob "github.com/ipfs/go-ipfs-chunker"
	cbor "github.com/ipfs/go-ipld-cbor"
	"github.com/ipfs/go-ipld-format"
	"github.com/ipfs/go-merkledag"
	"github.com/ipfs/go-unixfs"
	"github.com/ipfs/go-unixfs/pb"
	"github.com/ipld/go-car"
	"github.com/multiformats/go-varint"
	"github.com/pkg/errors"
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/hash"
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
		return nil, 0, errors.Wrap(err, "failed to get bytes from node")
	}
	pbNode := merkledag.NodeWithData(nodeBytes)
	err = pbNode.SetCidBuilder(merkledag.V1CidPrefix())
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to set cid builder")
	}
	for i := range links {
		err = pbNode.AddRawLink("", &links[i])
		if err != nil {
			return nil, 0, errors.Wrap(err, "failed to add link to node")
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
				return nil, nil, errors.Wrap(err, "failed to create parent node")
			}

			basicBlock, err := blocks.NewBlockWithCid(newNode.RawData(), newNode.Cid())
			if err != nil {
				return nil, nil, errors.Wrap(err, "failed to create block")
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

// GenerateCarHeader generates the CAR (Content Addressable aRchive) format header
// based on the given root CID (Content Identifier).
//
// Parameters:
//   - root: The root CID of the MerkleDag that the CAR file represents.
//
// Returns:
//   - []byte: The byte representation of the CAR header.
//   - error: An error that can occur during the header generation process, or nil if successful.
func GenerateCarHeader(root cid.Cid) ([]byte, error) {
	header := car.CarHeader{
		Roots:   []cid.Cid{root},
		Version: 1,
	}

	headerBytes, err := cbor.DumpObject(&header)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dump header")
	}
	headerBytesVarint := varint.ToUvarint(uint64(len(headerBytes)))
	headerBytes = append(headerBytesVarint, headerBytes...)
	return headerBytes, nil
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
	headerBytes, err := GenerateCarHeader(root)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate header")
	}
	_, err = io.Copy(writer, bytes.NewReader(headerBytes))
	if err != nil {
		return nil, errors.Wrap(err, "failed to write header")
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
		return written, errors.Wrap(err, "failed to write varint")
	}
	written += n

	n, err = io.Copy(writer, bytes.NewReader(block.Cid().Bytes()))
	if err != nil {
		return written, errors.Wrap(err, "failed to write cid")
	}
	written += n

	n, err = io.Copy(writer, bytes.NewReader(block.RawData()))
	if err != nil {
		return written, errors.Wrap(err, "failed to write raw")
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

// IsSameEntry checks if a given model.File and a given fs.Object represent the same entry,
// based on their size, hash, and last modification time.
//
// Parameters:
//   - ctx: Context that allows for asynchronous task cancellation.
//   - file: A model.File instance representing a file in the model.
//   - object: An fs.Object instance representing a file or object in the filesystem.
//
// Returns:
//   - bool: A boolean indicating whether the given model.File and fs.Object are considered the same entry.
//   - string: A string providing details in case of a mismatch.
//
// The function performs the following checks:
//  1. Compares the sizes of 'file' and 'object'. If there is a mismatch, it returns false along with a
//     formatted string that shows the mismatched sizes.
//  2. Retrieves the last modified time of 'object'.
//  3. Identifies a supported hash type for the storage backend of 'object' and computes its hash value.
//  4. The hash computation is skipped if the storage backend is a local file system or does not support any hash types.
//  5. If both 'file' and 'object' have non-empty hash values and these values do not match,
//     it returns false along with a formatted string that shows the mismatched hash values.
//  6. Compares the last modified times of 'file' and 'object' at the nanosecond precision.
//     If there is a mismatch, it returns false along with a formatted string that shows the mismatched times.
//  7. If all the checks pass (sizes, hashes, and last modified times match), the function returns true,
//     indicating that 'file' and 'object' are considered to be the same entry.
//
// Note:
// - In certain cases (e.g., failures during fetch), the last modified time might not be reliable.
// - For local file systems, hash computation is skipped to avoid inefficient operations.
func IsSameEntry(ctx context.Context, file model.File, object fs.Object) (bool, string) {
	if file.Size != object.Size() {
		return false, fmt.Sprintf("size mismatch: %d != %d", file.Size, object.Size())
	}
	var err error
	// last modified can be time.Now() if fetch failed so it may not be reliable.
	// This usually won't happen for most cloud provider i.e. S3
	// Because during scanning, the modified time is already fetched.
	lastModified := object.ModTime(ctx)
	supportedHash := object.Fs().Hashes().GetOne()
	// For local file system, rclone is actually hashing the file stream which is not efficient.
	// So we skip hashing for local file system.
	// For some of the remote storage, there may not have any supported hash type.
	var hashValue string
	if supportedHash != hash.None && object.Fs().Name() != "local" {
		hashValue, err = object.Hash(ctx, supportedHash)
		if err != nil {
			logger.Errorw("failed to hash", "error", err)
		}
	}
	if file.Hash != "" && hashValue != "" && file.Hash != hashValue {
		return false, fmt.Sprintf("hash mismatch: %s != %s", file.Hash, hashValue)
	}
	return lastModified.UnixNano() == file.LastModifiedTimestampNano,
		fmt.Sprintf("last modified mismatch: %d != %d",
			lastModified.UnixNano(),
			file.LastModifiedTimestampNano)
}

// GetBlockStreamFromFileRange streams a file range from the handler and encrypts it.
// It returns a channel of blocks, the object, and an error if any.
func GetBlockStreamFromFileRange(ctx context.Context,
	handler datasource.ReadHandler,
	fileRange model.FileRange,
	encryptor encryption.Encryptor) (<-chan BlockResult, fs.Object, error) {
	if encryptor != nil && (fileRange.Offset != 0 || fileRange.Length != fileRange.File.Size) {
		return nil, nil, errors.New("encryption is not supported for partial reads")
	}
	readStream, object, err := handler.Read(ctx, fileRange.File.Path, fileRange.Offset, fileRange.Length)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to open stream")
	}

	if object != nil {
		same, detail := IsSameEntry(ctx, *fileRange.File, object)
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
		return nil, object, errors.Wrap(err, "failed to encrypt stream")
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
				result = BlockResult{Error: errors.Wrap(err, "failed to read packJob")}
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

	return blockChan, object, err
}
