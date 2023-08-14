package pack

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack/encryption"
	"github.com/ipfs/boxo/util"
	"github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	chunk "github.com/ipfs/go-ipfs-chunker"
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

// CreateParentNode creates a new UnixFS parent node from child links.
// This function does not handle constructing layers of parent nodes.
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

// AssembleItemFromLinks creates a new UnixFS parent node from child links.
// This function handles constructing layers of parent nodes.
// It returns the additional blocks, the root node, and an error if any.
func AssembleItemFromLinks(links []format.Link) ([]blocks.Block, *merkledag.ProtoNode, error) {
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

var ErrItemModified = errors.New("item has been modified")

func IsSameEntry(ctx context.Context, item model.Item, object fs.Object) (bool, string) {
	if item.Size != object.Size() {
		return false, fmt.Sprintf("size mismatch: %d != %d", item.Size, object.Size())
	}
	var err error
	// last modified can be time.Now() if fetch failed so it may not be reliable.
	// This usually won't happen for most cloud provider i.e. S3
	// Because during scanning, the modified time is already fetched.
	lastModified := object.ModTime(ctx)
	// If last modified is not reliable, we will skip using it as a way to determine if the file has already scanned
	lastModifiedReliable := !lastModified.IsZero() && lastModified.Before(time.Now().Add(-time.Millisecond))
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
	if item.Hash != "" && hashValue != "" && item.Hash != hashValue {
		return false, fmt.Sprintf("hash mismatch: %s != %s", item.Hash, hashValue)
	}
	if lastModifiedReliable {
		return lastModified.UnixNano() == item.LastModifiedTimestampNano,
			fmt.Sprintf("last modified mismatch: %d != %d",
				lastModified.UnixNano(),
				item.LastModifiedTimestampNano)
	}
	return true, ""
}

// GetBlockStreamFromItem streams an item from the handler and encrypts it.
// It returns a channel of blocks, the object, and an error if any.
func GetBlockStreamFromItem(ctx context.Context,
	handler datasource.ReadHandler,
	itemPart model.ItemPart,
	encryptor encryption.Encryptor) (<-chan BlockResult, fs.Object, error) {
	if encryptor != nil && (itemPart.Offset != 0 || itemPart.Length != itemPart.Item.Size) {
		return nil, nil, errors.New("encryption is not supported for partial reads")
	}
	readStream, object, err := handler.Read(ctx, itemPart.Item.Path, itemPart.Offset, itemPart.Length)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to open stream")
	}

	if object != nil {
		lastModified := object.ModTime(ctx).UTC()
		lastModifiedReliable := !lastModified.IsZero() && lastModified.Before(time.Now().Add(-time.Millisecond))
		size := object.Size()
		var hashValue string
		supportedHash := object.Fs().Hashes().GetOne()
		if supportedHash != hash.None && object.Fs().Name() != "local" {
			hashValue, err = object.Hash(ctx, supportedHash)
			if err != nil {
				logger.Errorw("failed to hash", "error", err)
			}
		}
		switch {
		case hashValue != "" && hashValue != itemPart.Item.Hash:
			return nil, object, errors.Wrapf(ErrItemModified,
				"itemPart has been modified: %s, oldHash: %s, newHash: %s",
				itemPart.Item.Path, itemPart.Item.Hash, hashValue)
		case hashValue == "" && lastModifiedReliable && (lastModified.UnixNano() != itemPart.Item.LastModifiedTimestampNano || size != itemPart.Item.Size):
			return nil, object, errors.Wrapf(ErrItemModified,
				"itemPart has been modified: %s, oldSize: %d, newSize: %d, oldLastModified: %d, newLastModified: %d",
				itemPart.Item.Path, itemPart.Item.Size, size, itemPart.Item.LastModifiedTimestampNano, lastModified.UnixNano())
		case hashValue == "" && !lastModifiedReliable && size != itemPart.Item.Size:
			return nil, object, errors.Wrapf(ErrItemModified, "itemPart has been modified: %s, oldSize: %d, newSize: %d",
				itemPart.Item.Path, itemPart.Item.Size, size)
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
	chunker := chunk.NewSizeSplitter(readCloser, ChunkSize)
	go func() {
		if readStream != readCloser {
			defer readStream.Close()
		}
		defer readCloser.Close()
		defer close(blockChan)
		offset := itemPart.Offset
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

	return blockChan, object, err
}
