package pack

import (
	"context"
	"io"
	"os"
	"path"

	"github.com/data-preservation-programs/singularity/pack/encryption"
	util "github.com/ipfs/go-ipfs-util"
	"github.com/rclone/rclone/fs"

	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/model"
	commcid "github.com/filecoin-project/go-fil-commcid"
	commp "github.com/filecoin-project/go-fil-commp-hashhash"
	"github.com/google/uuid"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/ipfs/go-log/v2"
	"github.com/multiformats/go-varint"
	"github.com/pkg/errors"
)

type Result struct {
	ItemPartCIDs map[uint64]cid.Cid
	Objects      map[uint64]fs.Object
	CarResults   []CarResult
}

type CarResult struct {
	CarFilePath string
	CarFileSize int64
	PieceCID    cid.Cid
	PieceSize   int64
	RootCID     cid.Cid
	Header      []byte
	CarBlocks   []model.CarBlock
}

type nopCloser struct {
}

func (n nopCloser) Close() error {
	return nil
}

type WriteCloser struct {
	io.Writer
	io.Closer
}

var emptyItemCid = cid.NewCidV1(cid.Raw, util.Hash([]byte("")))

var logger = log.Logger("pack")

func GetMultiWriter(outDir string) (io.WriteCloser, *commp.Calc, string, error) {
	calc := &commp.Calc{}
	if outDir == "" {
		return WriteCloser{Writer: calc, Closer: &nopCloser{}}, calc, "", nil
	}
	var filepath string
	filename := uuid.NewString() + ".car"
	filepath = path.Join(outDir, filename)
	file, err := os.Create(filepath)
	if err != nil {
		return nil, nil, "", errors.Wrap(err, "failed to create file at "+filepath)
	}
	writer := io.MultiWriter(calc, file)
	return WriteCloser{Writer: writer, Closer: file}, calc, filepath, nil
}

func AssembleCar(
	ctx context.Context,
	handler datasource.ReadHandler,
	dataset model.Dataset,
	itemParts []model.ItemPart,
	outDir string,
	pieceSize int64,
) (*Result, error) {
	var writeCloser io.WriteCloser
	var calc *commp.Calc
	var filepath string
	var err error
	defer func() {
		if writeCloser != nil {
			writeCloser.Close()
		}
	}()

	result := &Result{
		ItemPartCIDs: make(map[uint64]cid.Cid),
		Objects:      make(map[uint64]fs.Object),
	}
	offset := int64(0)
	current := CarResult{}
	addHeaderIfNeeded := func(c cid.Cid) error {
		if offset == 0 {
			writeCloser, calc, filepath, err = GetMultiWriter(outDir)
			if err != nil {
				return errors.Wrap(err, "failed to get multi writer")
			}
			current.RootCID = c
			headerBytes, err := WriteCarHeader(writeCloser, c)
			if err != nil {
				return errors.Wrap(err, "failed to write header")
			}

			offset += int64(len(headerBytes))
			current.Header = headerBytes
		}
		return nil
	}
	checkResult := func(force bool) error {
		if !force && offset < pieceSize {
			return nil
		}
		if offset == 0 {
			return nil
		}
		pieceCid, finalPieceSize, err := GetCommp(calc, uint64(pieceSize))
		if err != nil {
			return errors.Wrap(err, "failed to get commp")
		}
		current.PieceCID = pieceCid
		current.PieceSize = int64(finalPieceSize)
		current.CarFileSize = offset

		if outDir != "" {
			current.CarFilePath = path.Join(outDir, pieceCid.String()+".car")
		}
		if filepath != "" {
			err = os.Rename(filepath, current.CarFilePath)
			if err != nil {
				return errors.Wrap(err, "failed to create symlink")
			}
		}
		result.CarResults = append(result.CarResults, current)
		current = CarResult{}
		offset = 0
		writeCloser.Close()
		writeCloser = nil
		calc = nil
		filepath = ""
		return nil
	}

	for _, itemPart := range itemParts {
		itemPart := itemPart
		links := make([]format.Link, 0)
		encryptor, err := encryption.GetEncryptor(dataset)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get encryptor")
		}
		if encryptor != nil && outDir == "" {
			return nil, errors.New("encryption is not supported without an output directory")
		}
		blockChan, object, err := GetBlockStreamFromItem(ctx, handler, itemPart, encryptor)
		if err != nil {
			return nil, errors.Wrap(err, "failed to stream itemPart")
		}

		result.Objects[itemPart.ItemID] = object

	blockChanLoop:
		for {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case block, ok := <-blockChan:
				if !ok {
					break blockChanLoop
				}
				if block.Error != nil {
					return nil, errors.Wrap(block.Error, "failed to stream block")
				}
				basicBlock, _ := blocks.NewBlockWithCid(block.Raw, block.CID)
				err = addHeaderIfNeeded(basicBlock.Cid())
				if err != nil {
					return nil, errors.Wrap(err, "failed to add header")
				}
				written, err := WriteCarBlock(writeCloser, basicBlock)
				if err != nil {
					return nil, errors.Wrap(err, "failed to write block")
				}

				if !dataset.UseEncryption() {
					current.CarBlocks = append(
						current.CarBlocks, model.CarBlock{
							CID:            model.CID(block.CID),
							CarOffset:      offset,
							CarBlockLength: int32(len(block.Raw)) + int32(block.CID.ByteLen()) + int32(varint.UvarintSize(uint64(len(block.Raw))+uint64(block.CID.ByteLen()))),
							Varint:         varint.ToUvarint(uint64(len(block.Raw)) + uint64(block.CID.ByteLen())),
							// nolint:exportloopref
							ItemID:        &itemPart.ItemID,
							ItemOffset:    block.Offset,
							ItemEncrypted: encryptor != nil,
						},
					)
				}

				offset += written
				links = append(
					links, format.Link{
						Name: "",
						Size: uint64(len(block.Raw)),
						Cid:  block.CID,
					},
				)
				err = checkResult(false)
				if err != nil {
					return nil, errors.Wrap(err, "failed to check result")
				}
			}
		}

		if len(links) == 0 {
			result.ItemPartCIDs[itemPart.ID] = emptyItemCid
			continue
		}
		if len(links) == 1 {
			result.ItemPartCIDs[itemPart.ID] = links[0].Cid
			continue
		}

		blks, rootNode, err := AssembleItemFromLinks(links)
		if err != nil {
			return nil, errors.Wrap(err, "failed to assemble itemPart")
		}
		for _, blk := range blks {
			err = addHeaderIfNeeded(blk.Cid())
			if err != nil {
				return nil, errors.Wrap(err, "failed to add header")
			}
			written, err := WriteCarBlock(writeCloser, blk)
			if err != nil {
				return nil, errors.Wrap(err, "failed to write block")
			}

			if !dataset.UseEncryption() {
				current.CarBlocks = append(
					current.CarBlocks, model.CarBlock{
						CID:            model.CID(blk.Cid()),
						CarOffset:      offset,
						CarBlockLength: int32(written),
						Varint:         varint.ToUvarint(uint64(len(blk.RawData()) + blk.Cid().ByteLen())),
						RawBlock:       blk.RawData(),
					},
				)
			}
			offset += written
			err = checkResult(false)
			if err != nil {
				return nil, errors.Wrap(err, "failed to check result")
			}
		}

		result.ItemPartCIDs[itemPart.ID] = rootNode.Cid()
	}

	err = checkResult(true)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check result")
	}
	return result, nil
}

func GetCommp(calc *commp.Calc, targetPieceSize uint64) (cid.Cid, uint64, error) {
	rawCommp, rawPieceSize, err := calc.Digest()
	if err != nil {
		return cid.Undef, 0, errors.Wrap(err, "failed to calculate commp")
	}

	if rawPieceSize < targetPieceSize {
		rawCommp, err = commp.PadCommP(rawCommp, rawPieceSize, targetPieceSize)
		if err != nil {
			return cid.Undef, 0, errors.Wrap(err, "failed to pad commp")
		}

		rawPieceSize = targetPieceSize
	} else if rawPieceSize > targetPieceSize {
		logger.Warn("piece size is larger than the target piece size")
	}

	commCid, err := commcid.DataCommitmentV1ToCID(rawCommp)
	if err != nil {
		return cid.Undef, 0, errors.Wrap(err, "failed to convert commp to cid")
	}

	return commCid, rawPieceSize, nil
}
