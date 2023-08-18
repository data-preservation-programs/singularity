package datasetworker

import (
	"context"
	"io"
	"os"
	"path"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack"
	"github.com/data-preservation-programs/singularity/pack/daggen"
	"github.com/data-preservation-programs/singularity/util"
	commp "github.com/filecoin-project/go-fil-commp-hashhash"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-varint"
	"gorm.io/gorm"
)

// ExportDag exports a Directed Acyclic Graph (DAG) for a given source.
// The function takes a source, iterates through the related directories
// (as rows from the database), and constructs the DAG in the form of a
// CAR (Content Addressable Archive) file. This CAR file represents the
// block structure of the data.
//
// The function:
// - Initializes necessary components like writers and calculators
// - Iterates through the directories linked with the source and fetches blocks
// - Writes the blocks into a CAR file
// - Closes the CAR file and renames it appropriately
// - Saves the CAR meta-information into the database
//
// Parameters:
// - ctx context.Context: The context to control cancellations and timeouts.
// - source model.Source: The source for which the DAG needs to be generated.
//
// The function performs several database and file system operations,
// each of which might result in an error. Errors are wrapped with context
// information and returned.
//
// Returns:
// - error: Standard error interface, returns nil if no error occurred during execution.
func (w *Thread) ExportDag(ctx context.Context, source model.Source) error {
	db := w.dbNoContext.WithContext(ctx)
	rootCID := pack.EmptyFileCid
	var outDir string
	var headerBytes []byte
	var carBlocks []model.CarBlock
	var car model.Car
	if len(source.Dataset.OutputDirs) > 0 {
		outDir = source.Dataset.OutputDirs[0]
	}

	var writeCloser io.WriteCloser
	var calc *commp.Calc
	var filepath string
	var err error
	offset := int64(0)
	rows, err := db.Model(&model.Directory{}).Where("source_id = ? AND exported = ?", source.ID, false).Order("id asc").Rows()
	if err != nil {
		return errors.Wrap(err, "failed to get directories")
	}
	defer rows.Close()
	dirCIDs := make(map[uint64]cid.Cid)
	for rows.Next() {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		var dir model.Directory
		err = db.ScanRows(rows, &dir)
		if err != nil {
			return errors.Wrap(err, "failed to scan directory")
		}
		dirCIDs[dir.ID] = cid.Cid(dir.CID)
		if dir.ParentID == nil && dir.CID != model.CID(cid.Undef) {
			rootCID = cid.Cid(dir.CID)
		}

		logger.Debugw("Reading content of directory", "dir_id", dir.ID, "name", dir.Name)
		_, blks, err := daggen.UnmarshalToBlocks(dir.Data)
		if err != nil {
			return errors.Wrap(err, "failed to unmarshall to blocks")
		}
		for _, blk := range blks {
			if len(blk.RawData()) == 0 && blk.Cid() != pack.EmptyFileCid {
				// This is dummy node. skip putting into car file
				continue
			}
			if offset == 0 {
				rootCID = blk.Cid()
				writeCloser, calc, filepath, err = pack.GetMultiWriter(outDir)
				if err != nil {
					return errors.Wrap(err, "failed to get multi writer")
				}
				defer writeCloser.Close()
				headerBytes, err = pack.WriteCarHeader(writeCloser, rootCID)
				if err != nil {
					return errors.Wrap(err, "failed to write header")
				}

				offset += int64(len(headerBytes))
			}
			written, err := pack.WriteCarBlock(writeCloser, blk)
			if err != nil {
				return errors.Wrap(err, "failed to write block")
			}
			carBlocks = append(carBlocks, model.CarBlock{
				CID:            model.CID(blk.Cid()),
				CarOffset:      offset,
				CarBlockLength: int32(written),
				Varint:         varint.ToUvarint(uint64(len(blk.RawData()) + blk.Cid().ByteLen())),
				RawBlock:       blk.RawData(),
			})
			offset += written
		}
	}

	if offset == 0 {
		logger.Warnw("no blocks to write")
		return nil
	}

	pieceCid, finalPieceSize, err := pack.GetCommp(calc, uint64(source.Dataset.PieceSize))
	if err != nil {
		return errors.Wrap(err, "failed to get commp")
	}
	if outDir != "" {
		car.FilePath = path.Join(outDir, pieceCid.String()+".car")
	}
	if car.FilePath != "" {
		writeCloser.Close()
		err = os.Rename(filepath, car.FilePath)
		if err != nil {
			return errors.Wrap(err, "failed to rename car file")
		}
	}
	car.Header = headerBytes
	car.PieceSize = int64(finalPieceSize)
	car.PieceCID = model.CID(pieceCid)
	car.RootCID = model.CID(rootCID)
	car.FileSize = offset
	car.DatasetID = source.DatasetID
	car.SourceID = &source.ID
	logger.Debugw("Saving car", "car", car)
	err = database.DoRetry(ctx, func() error {
		return db.Transaction(func(db *gorm.DB) error {
			err := db.Create(&car).Error
			if err != nil {
				return errors.WithStack(err)
			}
			for i := range carBlocks {
				carBlocks[i].CarID = car.ID
			}
			err = db.CreateInBatches(carBlocks, util.BatchSize).Error
			if err != nil {
				return errors.WithStack(err)
			}
			for dirID, dirCID := range dirCIDs {
				result := db.Model(&model.Directory{}).Where("id = ? AND cid = ?", dirID, model.CID(dirCID)).Update("exported", true)
				if result.Error != nil {
					return errors.Wrap(result.Error, "failed to update directory")
				}
				if result.RowsAffected == 0 {
					logger.Warnw("directory info has changed since we started. skipping update", "directory_id", dirID)
				}
			}
			return nil
		})
	})
	if err != nil {
		return errors.Wrap(err, "failed to save car")
	}
	return nil
}
