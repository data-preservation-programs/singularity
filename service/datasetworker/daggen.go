package datasetworker

import (
	"bytes"
	"context"
	"database/sql"
	"io"
	"os"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack"
	"github.com/data-preservation-programs/singularity/pack/daggen"
	"github.com/data-preservation-programs/singularity/pack/packutil"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/data-preservation-programs/singularity/util"
	commp "github.com/filecoin-project/go-fil-commp-hashhash"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-varint"
	"github.com/rclone/rclone/fs"
	"gorm.io/gorm"
)

type DagGenerator struct {
	ctx          context.Context
	db           *gorm.DB
	attachmentID model.SourceAttachmentID
	rows         *sql.Rows
	root         cid.Cid
	dirCIDs      map[model.DirectoryID]model.CID
	buffer       io.Reader
	done         bool
	carBlocks    []model.CarBlock
	offset       int64
	noInline     bool
}

// Read implements the io.Reader interface for the DagGenerator. It generates
// a CAR (Content Addressable Archive) representation of directories from a database,
// which can be read in chunks using the provided byte slice.
//
// Read operation involves several key steps:
//  1. It checks if the context has been canceled or if an error has occurred.
//  2. If there's an existing buffer, it reads from it.
//  3. If reading reaches the end of the current buffer, or if no buffer has been initialized,
//     the method fetches the next directory from the database and processes it.
//  4. The directory data is then converted to CAR format, and the resulting bytes are buffered.
//  5. Finally, the buffered data is read into the provided slice.
//
// Parameters:
//   - p: A byte slice that will be filled with the generated CAR data.
//
// Returns:
//   - The number of bytes read.
//   - An error if there's an issue during the operation. If the end of the data is reached,
//     it returns io.EOF.
func (d *DagGenerator) Read(p []byte) (int, error) {
	if d.ctx.Err() != nil {
		return 0, d.ctx.Err()
	}
	if d.buffer != nil {
		n, err := d.buffer.Read(p)
		if err == io.EOF {
			err = nil
			d.buffer = nil
		}
		return n, err
	}

	if d.done {
		return 0, io.EOF
	}

	db := d.db
	if d.rows == nil {
		rows, err := db.
			Model(&model.Directory{}).
			Where("attachment_id = ? AND exported = ?", d.attachmentID, false).
			Order("id asc").Rows()
		if err != nil {
			return 0, errors.WithStack(err)
		}
		d.rows = rows
		header, err := util.GenerateCarHeader(d.root)
		if err != nil {
			return 0, errors.WithStack(err)
		}
		d.buffer = bytes.NewReader(header)
		d.offset += int64(len(header))
		return 0, nil
	}
	if !d.rows.Next() {
		d.done = true
		return 0, nil
	}
	var dir model.Directory
	err := db.ScanRows(d.rows, &dir)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	d.dirCIDs[dir.ID] = dir.CID
	blks, err := daggen.UnmarshalToBlocks(dir.Data)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to unmarshall directory %d to blocks", dir.ID)
	}
	readers := make([]io.Reader, 0, len(blks)*3)
	for _, blk := range blks {
		if len(blk.RawData()) == 0 && blk.Cid() != packutil.EmptyFileCid {
			// This is dummy node. skip putting into car file
			continue
		}

		carBlockSize := len(blk.RawData()) + blk.Cid().ByteLen()
		vint := varint.ToUvarint(uint64(carBlockSize))
		carBlockSize += len(vint)
		readers = append(readers, bytes.NewReader(vint), bytes.NewReader(blk.Cid().Bytes()), bytes.NewReader(blk.RawData()))
		if !d.noInline {
			d.carBlocks = append(d.carBlocks, model.CarBlock{
				CID:            model.CID(blk.Cid()),
				CarOffset:      d.offset,
				CarBlockLength: int32(carBlockSize),
				Varint:         vint,
				RawBlock:       blk.RawData(),
			})
		}
		d.offset += int64(carBlockSize)
	}
	d.buffer = io.MultiReader(readers...)
	return 0, nil
}

func (d *DagGenerator) Close() error {
	if d.rows != nil {
		return errors.WithStack(d.rows.Close())
	}
	return nil
}

func NewDagGenerator(ctx context.Context, db *gorm.DB, attachmentID model.SourceAttachmentID, root cid.Cid, noInline bool) *DagGenerator {
	return &DagGenerator{
		ctx:          ctx,
		db:           db,
		attachmentID: attachmentID,
		root:         root,
		dirCIDs:      make(map[model.DirectoryID]model.CID),
		noInline:     noInline,
	}
}

var ErrDagNotReady = errors.New("dag is not ready to be generated")

var ErrDagDisabled = errors.New("dag generation is disabled for this preparation")

// ExportDag exports a Directed Acyclic Graph (DAG) for a given source.
// The function takes a source, iterates through the related directories
// (as rows from the database), and constructs the DAG in the form of a
// CAR (Content Addressable Archive) file. This CAR file represents the
// block structure of the data.
//
// The function:
//   - Initializes necessary components like writers and calculators
//   - Iterates through the directories linked with the source and fetches blocks
//   - Writes the blocks into a CAR file
//   - Closes the CAR file and renames it appropriately
//   - Saves the CAR meta-information into the database
//
// Parameters:
//   - ctx context.Context: The context to control cancellations and timeouts.
//   - source model.Source: The source for which the DAG needs to be generated.
//
// The function performs several database and file system operations,
// each of which might result in an error. Errors are wrapped with context
// information and returned.
//
// Returns:
//   - error: Standard error interface, returns nil if no error occurred during execution.
func (w *Thread) ExportDag(ctx context.Context, job model.Job) error {
	if job.Attachment.Preparation.NoDag {
		return errors.WithStack(ErrDagDisabled)
	}

	rootCID, err := job.Attachment.RootDirectoryCID(ctx, w.dbNoContext)
	if err != nil {
		return errors.WithStack(err)
	}

	if rootCID == cid.Undef {
		return ErrDagNotReady
	}

	db := w.dbNoContext.WithContext(ctx)
	pieceSize := job.Attachment.Preparation.GetMinPieceSize()
	// storageWriter can be nil for inline preparation
	storageID, storageWriter, err := storagesystem.GetRandomOutputWriter(ctx, job.Attachment.Preparation.OutputStorages)
	if err != nil {
		return errors.WithStack(err)
	}

	dagGenerator := NewDagGenerator(ctx, db, job.Attachment.ID, rootCID, job.Attachment.Preparation.NoInline)
	defer dagGenerator.Close()

	var filename string
	calc := &commp.Calc{}
	var pieceCid cid.Cid
	var finalPieceSize uint64
	var fileSize int64
	var minPieceSizePadding int64
	if storageWriter != nil {
		// Find the output storage to determine if it's local or remote
		var outputStorage *model.Storage
		for i := range job.Attachment.Preparation.OutputStorages {
			if job.Attachment.Preparation.OutputStorages[i].ID == *storageID {
				outputStorage = &job.Attachment.Preparation.OutputStorages[i]
				break
			}
		}

		filename = uuid.NewString() + ".car"
		var obj fs.Object

		// For remote storage, use temp file to enable padding after writing
		if outputStorage != nil && outputStorage.Type != "local" {
			// Write to temp file first
			tempFile, err := os.CreateTemp("", "dagcar-*.car")
			if err != nil {
				return errors.Wrap(err, "failed to create temp file for DAG CAR")
			}
			tempPath := tempFile.Name()
			defer os.Remove(tempPath)

			reader := io.TeeReader(dagGenerator, calc)
			_, err = io.Copy(tempFile, reader)
			if err != nil {
				tempFile.Close()
				return errors.Wrap(err, "failed to write DAG CAR to temp file")
			}
			tempFile.Close()

			if dagGenerator.offset <= 59 {
				logger.Info("Nothing to export to dag. Skipping.")
				return nil
			}

			// Get file size before padding
			stat, err := os.Stat(tempPath)
			if err != nil {
				return errors.WithStack(err)
			}
			fileSize = stat.Size()

			pieceCid, finalPieceSize, err = pack.GetCommp(calc, uint64(pieceSize))
			if err != nil {
				return errors.WithStack(err)
			}

			// Check if minPieceSize constraint forced larger piece size
			naturalPieceSize := util.NextPowerOfTwo(uint64(fileSize))
			if finalPieceSize > naturalPieceSize {
				// Need to pad to (127/128) × piece_size due to Fr32 padding overhead
				targetCarSize := (int64(finalPieceSize) * 127) / 128
				paddingNeeded := targetCarSize - fileSize

				// Append zeros to temp file
				f, err := os.OpenFile(tempPath, os.O_APPEND|os.O_WRONLY, 0644)
				if err != nil {
					return errors.Wrap(err, "failed to open temp DAG CAR file for padding")
				}

				zeros := make([]byte, paddingNeeded)
				_, err = f.Write(zeros)
				f.Close()
				if err != nil {
					return errors.Wrap(err, "failed to write padding to temp DAG CAR file")
				}

				fileSize = targetCarSize
				logger.Infow("padded DAG CAR file for minPieceSize (remote storage)", "original", fileSize-paddingNeeded, "padded", fileSize, "padding", paddingNeeded, "piece_size", finalPieceSize)
			}

			// Upload complete file to remote storage
			f, err := os.Open(tempPath)
			if err != nil {
				return errors.Wrap(err, "failed to open temp file for upload")
			}
			defer f.Close()

			obj, err = storageWriter.Write(ctx, filename, f)
			if err != nil {
				return errors.WithStack(err)
			}
		} else {
			// Local storage: write directly then append if needed
			reader := io.TeeReader(dagGenerator, calc)
			obj, err = storageWriter.Write(ctx, filename, reader)
			if err != nil {
				return errors.WithStack(err)
			}
			fileSize = obj.Size()

			if dagGenerator.offset <= 59 {
				logger.Info("Nothing to export to dag. Skipping.")
				return nil
			}

			pieceCid, finalPieceSize, err = pack.GetCommp(calc, uint64(pieceSize))
			if err != nil {
				return errors.WithStack(err)
			}

			// Check if minPieceSize constraint forced larger piece size
			naturalPieceSize := util.NextPowerOfTwo(uint64(fileSize))
			if finalPieceSize > naturalPieceSize {
				// Need to pad to (127/128) × piece_size due to Fr32 padding overhead
				targetCarSize := (int64(finalPieceSize) * 127) / 128
				paddingNeeded := targetCarSize - fileSize

				if outputStorage != nil && obj != nil {
					// Build full path to CAR file
					carPath := outputStorage.Path + "/" + filename

					// Reopen file and append zeros
					f, err := os.OpenFile(carPath, os.O_APPEND|os.O_WRONLY, 0644)
					if err != nil {
						return errors.Wrap(err, "failed to open DAG CAR file for padding")
					}

					zeros := make([]byte, paddingNeeded)
					_, err = f.Write(zeros)
					f.Close()
					if err != nil {
						return errors.Wrap(err, "failed to write padding to DAG CAR file")
					}

					fileSize = targetCarSize
					logger.Infow("padded DAG CAR file for minPieceSize (local storage)", "original", fileSize-paddingNeeded, "padded", fileSize, "padding", paddingNeeded, "piece_size", finalPieceSize)
				}
			}
		}

		_, err = storageWriter.Move(ctx, obj, pieceCid.String()+".car")
		if err != nil && !errors.Is(err, storagesystem.ErrMoveNotSupported) {
			logger.Errorf("failed to move car file from %s to %s: %s", filename, pieceCid.String()+".car", err)
		}
		if err == nil {
			filename = pieceCid.String() + ".car"
		}
	} else {
		fileSize, err = io.Copy(calc, dagGenerator)
		if err != nil {
			return errors.WithStack(err)
		}

		if dagGenerator.offset <= 59 {
			logger.Info("Nothing to export to dag. Skipping.")
			return nil
		}

		pieceCid, finalPieceSize, err = pack.GetCommp(calc, uint64(pieceSize))
		if err != nil {
			return errors.WithStack(err)
		}

		// Check if minPieceSize constraint forced larger piece size
		naturalPieceSize := util.NextPowerOfTwo(uint64(fileSize))
		if finalPieceSize > naturalPieceSize {
			// For inline, store virtual padding to (127/128) × piece_size
			targetCarSize := (int64(finalPieceSize) * 127) / 128
			paddingNeeded := targetCarSize - fileSize
			fileSize = targetCarSize
			minPieceSizePadding = paddingNeeded

			logger.Infow("inline DAG CAR needs virtual padding for minPieceSize", "data", fileSize-paddingNeeded, "total", fileSize, "padding", paddingNeeded, "piece_size", finalPieceSize)
		}
	}

	car := model.Car{
		PieceCID:            model.CID(pieceCid),
		PieceSize:           int64(finalPieceSize),
		RootCID:             model.CID(rootCID),
		FileSize:            fileSize,
		MinPieceSizePadding: minPieceSizePadding,
		StorageID:           storageID,
		StoragePath:         filename,
		AttachmentID:        &job.AttachmentID,
		PreparationID:       job.Attachment.PreparationID,
		PieceType:           model.DagPiece,
	}

	err = database.DoRetry(ctx, func() error {
		return db.Transaction(func(db *gorm.DB) error {
			err := db.Create(&car).Error
			if err != nil {
				return errors.WithStack(err)
			}
			if len(dagGenerator.carBlocks) > 0 {
				for i := range dagGenerator.carBlocks {
					dagGenerator.carBlocks[i].CarID = car.ID
				}
				err = db.CreateInBatches(dagGenerator.carBlocks, util.BatchSize).Error
				if err != nil {
					return errors.WithStack(err)
				}
			}
			for dirID, dirCID := range dagGenerator.dirCIDs {
				result := db.Model(&model.Directory{}).Where("id = ? AND cid = ?", dirID, dirCID).Update("exported", true)
				if result.Error != nil {
					return errors.Wrap(result.Error, "failed to update directory")
				}
				if result.RowsAffected == 0 {
					logger.Warnf("directory %d has changed since we started.", dirID)
				}
			}
			return nil
		})
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
