package pack

import (
	"context"
	"io"
	"os"
	"path"

	"github.com/data-preservation-programs/singularity/pack/encryption"
	"github.com/ipfs/boxo/util"
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
	FileRangeCIDs map[uint64]cid.Cid
	Objects       map[uint64]fs.Object
	CarResults    []CarResult
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

var EmptyItemCid = cid.NewCidV1(cid.Raw, util.Hash([]byte("")))

var logger = log.Logger("pack")

// GetMultiWriter constructs a writer that writes to multiple outputs
// (i.e., a MultiWriter) based on the specified output directory.
// This function is designed to handle the use case where you want
// to calculate a piece commitment (via a commp.Calc instance) while
// writing data, and potentially write that data to a file at the same
// time.
//
// Parameters:
//   - outDir: A string representing the path of the output directory. If this
//     string is empty, the function will return a writer that only writes to
//     the commp.Calc instance.
//
// Returns:
//   - io.WriteCloser: An interface that groups the basic Write and Close methods.
//     The Write method is implemented by an io.MultiWriter that writes to both
//     a commp.Calc instance and an os.File (if an output directory was specified).
//     The Close method is implemented by the os.File, or a nopCloser if no output
//     directory was specified.
//   - *commp.Calc: A pointer to a commp.Calc instance. This is used to calculate
//     the piece commitment of the data that is written.
//   - string: The path of the file that the data is being written to. This is an
//     empty string if no output directory was specified.
//   - error: An error that can occur when creating the file at the specified path,
//     or nil if the operation was successful.
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

// AssembleCar assembles a Content Addressable aRchive (CAR) file from a list of data blocks.
// It writes these blocks to the specified output directory, and it attempts to ensure that
// the size of the CAR file matches a given target piece size. The function can handle
// assembling multiple CAR files if the blocks don't fit into a single CAR file of the
// target size.
//
// Parameters:
// - ctx: Context used to handle cancellation and deadlines.
// - handler: The ReadHandler interface which handles reading data from a data source.
// - dataset: Information about the dataset being processed.
// - fileRanges: List of data blocks that need to be included in the CAR file.
// - outDir: Directory where the CAR file(s) will be written.
// - pieceSize: Target size of each CAR file, in bytes.
//
// Returns:
//   - *Result: A structure that contains results of the CAR file creation,
//     including CIDs of the items, file paths, and sizes.
//   - error: An error that occurred during the CAR assembly process, or nil if the operation was successful.
func AssembleCar(
	ctx context.Context,
	handler datasource.ReadHandler,
	dataset model.Dataset,
	fileRanges []model.FileRange,
	outDir string,
	pieceSize int64,
) (*Result, error) {
	logger.Debugw("assembling car", "dataset",
		dataset.ID, "fileRanges", len(fileRanges), "outDir", outDir, "pieceSize", pieceSize)
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
		FileRangeCIDs: make(map[uint64]cid.Cid),
		Objects:       make(map[uint64]fs.Object),
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
		writeCloser.Close()
		writeCloser = nil
		if filepath != "" {
			err = os.Rename(filepath, current.CarFilePath)
			if err != nil {
				return errors.Wrap(err, "failed to create symlink")
			}
		}
		result.CarResults = append(result.CarResults, current)
		current = CarResult{}
		offset = 0
		calc = nil
		filepath = ""
		return nil
	}

	for _, fileRange := range fileRanges {
		fileRange := fileRange
		links := make([]format.Link, 0)
		encryptor, err := encryption.GetEncryptor(dataset)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get encryptor")
		}
		if encryptor != nil && outDir == "" {
			return nil, errors.New("encryption is not supported without an output directory")
		}
		blockChan, object, err := GetBlockStreamFromItem(ctx, handler, fileRange, encryptor)
		if err != nil {
			return nil, errors.Wrap(err, "failed to stream fileRange")
		}

		result.Objects[fileRange.FileID] = object

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
							FileID:        &fileRange.FileID,
							FileOffset:    block.Offset,
							FileEncrypted: encryptor != nil,
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
			result.FileRangeCIDs[fileRange.ID] = EmptyItemCid
			continue
		}
		if len(links) == 1 {
			result.FileRangeCIDs[fileRange.ID] = links[0].Cid
			continue
		}

		blks, rootNode, err := AssembleItemFromLinks(links)
		if err != nil {
			return nil, errors.Wrap(err, "failed to assemble fileRange")
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

		result.FileRangeCIDs[fileRange.ID] = rootNode.Cid()
	}

	err = checkResult(true)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check result")
	}
	return result, nil
}

// GetCommp calculates the data commitment (CommP) and the piece size based on the
// provided commp.Calc instance and target piece size. It ensures that the
// calculated piece size matches the target piece size specified. If necessary,
// it pads the data to meet the target piece size.
//
// Parameters:
//   - calc: A pointer to a commp.Calc instance, which has been used to write data
//     and will be used to calculate the piece commitment for that data.
//   - targetPieceSize: The desired size of the piece, specified in bytes.
//
// Returns:
//   - cid.Cid: A CID (Content Identifier) representing the data commitment (CommP).
//   - uint64: The size of the piece, in bytes, after potential padding.
//   - error: An error indicating issues during the piece commitment calculation,
//     padding, or CID conversion, or nil if the operation was successful.
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
