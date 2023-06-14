package datasetworker

import (
	"context"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack"
	"github.com/data-preservation-programs/singularity/pack/daggen"
	"github.com/data-preservation-programs/singularity/pack/device"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-varint"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"os"
	"path"
	"time"
)

func (w *DatasetWorkerThread) dag(ctx context.Context, source model.Source) error {
	var rootCID cid.Cid
	var outDir string
	var headerBytes []byte
	var carBlocks []model.CarBlock
	var car model.Car
	if len(source.Dataset.OutputDirs) > 0 {
		var err error
		outDir, err = device.GetPathWithMostSpace(source.Dataset.OutputDirs)
		if err != nil {
			w.logger.Warnw("failed to get path with most space. using the first one", "error", err)
			outDir = source.Dataset.OutputDirs[0]
		}
	}

	writeCloser, calc, filepath, err := pack.GetMultiWriter(outDir)
	if err != nil {
		return errors.Wrap(err, "failed to get multi writer")
	}
	defer writeCloser.Close()
	offset := int64(0)
	rows, err := w.db.Model(&model.Directory{}).Where("source_id = ? AND exported = ?", source.ID, false).Order("id asc").Rows()
	if err != nil {
		return errors.Wrap(err, "failed to get directories")
	}
	defer rows.Close()
	dirUpdateTimes := map[uint64]time.Time{}
	for rows.Next() {
		var dir model.Directory
		err = w.db.ScanRows(rows, &dir)
		if err != nil {
			return errors.Wrap(err, "failed to scan directory")
		}
		dirUpdateTimes[dir.ID] = dir.UpdatedAt

		_, blks, err := daggen.UnmarshallToBlocks(dir.Data)
		if err != nil {
			return errors.Wrap(err, "failed to unmarshall to blocks")
		}
		for _, blk := range blks {
			if len(blk.RawData()) == 0 {
				// This is dummy node. skip putting into car file
				continue
			}
			if offset == 0 {
				rootCID = blk.Cid()
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
	if outDir != "" {
		car.FilePath = path.Join(outDir, pieceCid.String()+".car")
	}
	if car.FilePath != "" {
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
	err = database.DoRetry(func() error {
		return w.db.Transaction(func(db *gorm.DB) error {
			err := db.Create(&car).Error
			if err != nil {
				return err
			}
			for i, _ := range carBlocks {
				carBlocks[i].CarID = car.ID
			}
			err = db.CreateInBatches(carBlocks, 100).Error
			if err != nil {
				return err
			}
			for dirID, updatedAt := range dirUpdateTimes {
				result := db.Model(&model.Directory{}).Where("id = ? AND updated_at = ?", dirID, updatedAt).Update("exported", true)
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
