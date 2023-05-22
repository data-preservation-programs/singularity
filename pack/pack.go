package pack

import (
	"bytes"
	"context"
	"filippo.io/age"
	"io"
	"os"
	"path"

	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/model"
	commcid "github.com/filecoin-project/go-fil-commcid"
	commp "github.com/filecoin-project/go-fil-commp-hashhash"
	"github.com/google/uuid"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	chunk "github.com/ipfs/go-ipfs-chunker"
	util "github.com/ipfs/go-ipfs-util"
	cbor "github.com/ipfs/go-ipld-cbor"
	format "github.com/ipfs/go-ipld-format"
	"github.com/ipfs/go-log/v2"
	"github.com/ipfs/go-merkledag"
	"github.com/ipfs/go-unixfs"
	unixfs_pb "github.com/ipfs/go-unixfs/pb"
	"github.com/ipld/go-car"
	"github.com/multiformats/go-varint"
	"github.com/pkg/errors"
)

type BlockResult struct {
	CID    cid.Cid
	Offset int64
	Length int64
	Raw    []byte
	Error  error
}
type Result struct {
	CarFilePath string
	CarFileSize int64
	PieceCID    cid.Cid
	PieceSize   int64
	RootCID     cid.Cid
	Header      []byte
	CarBlocks   []model.CarBlock
	ItemCIDs    map[uint64]cid.Cid
}

const ChunkSize int64 = 1 << 20
const NumLinkPerNode = 1024

var EmptyBlockCID = cid.NewCidV1(cid.Raw, util.Hash([]byte{}))

type Link struct {
	format.Link
	ChunkSize uint64
}

func createParentNode(links []Link) (format.Node, uint64, error) {
	node := unixfs.NewFSNode(unixfs_pb.Data_File)
	total := uint64(0)
	for _, link := range links {
		node.AddBlockSize(link.ChunkSize)
		total += link.ChunkSize
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
	for _, link := range links {
		err = pbNode.AddRawLink("", &link.Link)
		if err != nil {
			return nil, 0, errors.Wrap(err, "failed to add link to node")
		}
	}
	return pbNode, total, nil
}

func writeCarBlock(writer io.Writer, block blocks.BasicBlock) (int64, error) {
	written := int64(0)
	varintBytes := varint.ToUvarint(uint64(len(block.RawData()) + block.Cid().ByteLen()))
	n, err := io.Copy(writer, bytes.NewReader(varintBytes))
	if err != nil {
		return written, errors.Wrap(err, "failed to write varint")
	}
	written += int64(n)

	n, err = io.Copy(writer, bytes.NewReader(block.Cid().Bytes()))
	if err != nil {
		return written, errors.Wrap(err, "failed to write cid")
	}
	written += int64(n)

	n, err = io.Copy(writer, bytes.NewReader(block.RawData()))
	if err != nil {
		return written, errors.Wrap(err, "failed to write raw")
	}
	written += int64(n)
	return written, nil
}

func ProcessItems(
	ctx context.Context,
	handler datasource.Handler,
	items []model.Item,
	outDir string,
	pieceSize int64,
	recipients []string,
) (*Result, error) {
	result := &Result{
		ItemCIDs: make(map[uint64]cid.Cid),
	}
	offset := int64(0)
	var headerBytes []byte

	calc := &commp.Calc{}
	var writer io.Writer = calc
	var filepath string
	if outDir != "" {
		filename := uuid.NewString() + ".car"
		filepath = path.Join(outDir, filename)
		file, err := os.Create(filepath)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create file at "+filepath)
		}
		defer file.Close()
		writer = io.MultiWriter(calc, file)
	}

	for _, item := range items {
		item := item
		links := make([]Link, 0)
		blockChan, err := streamItem(ctx, handler, item, recipients)
		if err != nil {
			return nil, errors.Wrap(err, "failed to stream item")
		}

		for block := range blockChan {
			links = append(
				links, Link{
					Link: format.Link{
						Name: "",
						Size: uint64(block.Length),
						Cid:  block.CID,
					},
					ChunkSize: uint64(block.Length),
				},
			)
			if block.Error != nil {
				return nil, errors.Wrap(block.Error, "failed to stream block")
			}

			if offset == 0 {
				result.RootCID = block.CID
				header := car.CarHeader{
					Roots:   []cid.Cid{block.CID},
					Version: 1,
				}

				headerBytes, err = cbor.DumpObject(&header)
				if err != nil {
					return nil, errors.Wrap(err, "failed to dump header")
				}
				headerBytesVarint := varint.ToUvarint(uint64(len(headerBytes)))

				result.Header = append(headerBytesVarint, headerBytes...)
				n, err := io.Copy(writer, bytes.NewReader(result.Header))
				if err != nil {
					return nil, errors.Wrap(err, "failed to write header")
				}

				offset += int64(n)
			}

			basicBlock, _ := blocks.NewBlockWithCid(block.Raw, block.CID)
			written, err := writeCarBlock(writer, *basicBlock)
			if err != nil {
				return nil, errors.Wrap(err, "failed to write block")
			}

			offset += written

			result.CarBlocks = append(
				result.CarBlocks, model.CarBlock{
					CID:            block.CID.String(),
					CarOffset:      offset - written,
					CarBlockLength: written,
					Varint:         uint64(block.Length) + uint64(block.CID.ByteLen()),
					ItemID:         &item.ID,
					ItemOffset:     block.Offset,
					BlockLength:    block.Length,
				},
			)
		}

		for len(links) > 1 {
			newLinks := make([]Link, 0)
			for start := 0; start < len(links); start += NumLinkPerNode {
				newNode, total, err := createParentNode(links[start:Min(start+NumLinkPerNode, len(links))])
				if err != nil {
					return nil, errors.Wrap(err, "failed to create parent node")
				}

				basicBlock, _ := blocks.NewBlockWithCid(newNode.RawData(), newNode.Cid())
				written, err := writeCarBlock(writer, *basicBlock)
				if err != nil {
					return nil, errors.Wrap(err, "failed to write block")
				}
				offset += written

				result.CarBlocks = append(
					result.CarBlocks, model.CarBlock{
						CID:            basicBlock.Cid().String(),
						CarOffset:      offset - written,
						CarBlockLength: written,
						Varint:         uint64(len(basicBlock.RawData()) + basicBlock.Cid().ByteLen()),
						RawBlock:       basicBlock.RawData(),
						BlockLength:    int64(len(basicBlock.RawData())),
					},
				)

				newLinks = append(
					newLinks, Link{
						ChunkSize: total,
						Link: format.Link{
							Name: "",
							Size: total,
							Cid:  newNode.Cid(),
						},
					},
				)
			}

			links = newLinks
		}

		if len(links) == 0 {
			result.ItemCIDs[item.ID] = EmptyBlockCID
		} else {
			result.ItemCIDs[item.ID] = links[0].Cid
		}
	}

	rawCommp, rawPieceSize, err := calc.Digest()
	if err != nil {
		return nil, errors.Wrap(err, "failed to calculate commp")
	}

	if rawPieceSize < uint64(pieceSize) {
		rawCommp, err = commp.PadCommP(rawCommp, rawPieceSize, uint64(pieceSize))
		if err != nil {
			return nil, errors.Wrap(err, "failed to pad commp")
		}

		rawPieceSize = uint64(pieceSize)
	} else if rawPieceSize > uint64(pieceSize) {
		log.Logger("packing").Warn("piece size is larger than the target piece size")
	}

	commCid, err := commcid.DataCommitmentV1ToCID(rawCommp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert commp to cid")
	}

	result.PieceCID = commCid
	result.PieceSize = int64(rawPieceSize)
	result.CarFileSize = offset

	if outDir != "" {
		result.CarFilePath = path.Join(outDir, commCid.String()+".car")
	}
	if filepath != "" {
		err = os.Rename(filepath, result.CarFilePath)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create symlink")
		}
	}

	return result, nil
}

func Min(i int, i2 int) int {
	if i < i2 {
		return i
	}
	return i2
}

func streamItem(ctx context.Context, handler datasource.Handler, item model.Item, recipients []string) (<-chan BlockResult, error) {
	readStream, _, err := handler.Read(ctx, item.Path, item.Offset, item.Length)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open stream")
	}

	if len(recipients) > 0 {
		var rs []age.Recipient
		for _, recipient := range recipients {
			r, err := age.ParseX25519Recipient(recipient)
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse recipient")
			}
			rs = append(rs, r)
		}
		reader, writer := io.Pipe()
		target, err := age.Encrypt(writer, rs...)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create encrypt stream")
		}
		readStream2 := readStream
		go func() {
			defer target.Close()
			io.Copy(target, readStream2)
		}()
		readStream = reader
	}

	blockChan := make(chan BlockResult)
	chunker := chunk.NewSizeSplitter(readStream, ChunkSize)
	go func() {
		defer readStream.Close()
		defer close(blockChan)
		offset := item.Offset
		for {
			chunkerBytes, err := chunker.NextBytes()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return
				}
				blockChan <- BlockResult{Error: errors.Wrap(err, "failed to read chunk")}
				return
			}

			hash := util.Hash(chunkerBytes)
			c := cid.NewCidV1(cid.Raw, hash)
			blockChan <- BlockResult{
				CID:    c,
				Offset: offset,
				Length: int64(len(chunkerBytes)),
				Raw:    chunkerBytes,
				Error:  nil,
			}

			offset += int64(len(chunkerBytes))
		}
	}()

	return blockChan, err
}
