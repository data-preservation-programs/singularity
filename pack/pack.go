package pack

import (
	"context"
	"io"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/analytics"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack/daggen"
	"github.com/data-preservation-programs/singularity/pack/packutil"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/data-preservation-programs/singularity/util"
	commcid "github.com/filecoin-project/go-fil-commcid"
	commp "github.com/filecoin-project/go-fil-commp-hashhash"
	"github.com/google/uuid"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/ipfs/go-log/v2"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var logger = log.Logger("pack")

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
		return cid.Undef, 0, errors.WithStack(err)
	}

	if rawPieceSize < targetPieceSize {
		rawCommp, err = commp.PadCommP(rawCommp, rawPieceSize, targetPieceSize)
		if err != nil {
			return cid.Undef, 0, errors.WithStack(err)
		}

		rawPieceSize = targetPieceSize
	} else if rawPieceSize > targetPieceSize {
		logger.Warn("piece size is larger than the target piece size")
	}

	commCid, err := commcid.DataCommitmentV1ToCID(rawCommp)
	if err != nil {
		return cid.Undef, 0, errors.WithStack(err)
	}

	return commCid, rawPieceSize, nil
}

var ErrNoContent = errors.New("no content to pack")

// Pack takes in a Job and processes its attachment by reading it, possibly encrypting it,
// splitting it into manageable chunks, and then storing those chunks into a designated storage.
// If the job's attachment requires encryption, it will be encrypted using the specified encryption method.
// The function returns a slice of Car objects which represent the stored chunks and an error if any occurred.
//
// Parameters:
//   - ctx: The context which controls the lifetime of the operation.
//   - db: The gorm database instance used for querying and updating database records.
//   - job: The Job model instance which contains information about the attachment to be processed.
//
// Returns:
//   - A slice of model.Car instances representing stored chunks.
//   - An error, if any occurred during the operation.
func Pack(
	ctx context.Context,
	db *gorm.DB,
	job model.Job,
) (*model.Car, error) {
	db = db.WithContext(ctx)
	pieceSize := job.Attachment.Preparation.PieceSize
	// storageWriter can be nil for inline preparation
	storageID, storageWriter, err := storagesystem.GetRandomOutputWriter(ctx, job.Attachment.Preparation.OutputStorages)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	storageReader, err := storagesystem.NewRCloneHandler(ctx, *job.Attachment.Storage)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get storage handler for %s", job.Attachment.Storage.Name)
	}

	var skipInaccessibleFile bool
	if job.Attachment.Storage.ClientConfig.SkipInaccessibleFile != nil {
		skipInaccessibleFile = *job.Attachment.Storage.ClientConfig.SkipInaccessibleFile
	}
	assembler := NewAssembler(ctx, storageReader, job.FileRanges, job.Attachment.Preparation.NoInline, skipInaccessibleFile)
	defer assembler.Close()
	var filename string
	calc := &commp.Calc{}
	var pieceCid cid.Cid
	var finalPieceSize uint64
	var fileSize int64
	if storageWriter != nil {
		var carGenerated bool
		reader := io.TeeReader(assembler, calc)
		filename = uuid.NewString() + ".car"
		obj, err := storageWriter.Write(ctx, filename, reader)
		defer func() {
			if !carGenerated && obj != nil {
				removeCtx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
				err := storageWriter.Remove(removeCtx, obj)
				if err != nil {
					logger.Errorf("failed to remove temporary CAR file %s: %v", filename, err)
				}
				cancel()
			}
		}()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		fileSize = obj.Size()

		if assembler.carOffset <= 65 {
			return nil, errors.WithStack(ErrNoContent)
		}
		pieceCid, finalPieceSize, err = GetCommp(calc, uint64(pieceSize))
		if err != nil {
			return nil, errors.WithStack(err)
		}
		_, err = storageWriter.Move(ctx, obj, pieceCid.String()+".car")
		if err != nil && !errors.Is(err, storagesystem.ErrMoveNotSupported) {
			logger.Errorf("failed to move car file from %s to %s: %s", filename, pieceCid.String()+".car", err)
		}
		if err == nil {
			filename = pieceCid.String() + ".car"
		}
		carGenerated = true
	} else {
		fileSize, err = io.Copy(calc, assembler)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if assembler.carOffset <= 65 {
			return nil, errors.WithStack(ErrNoContent)
		}
		pieceCid, finalPieceSize, err = GetCommp(calc, uint64(pieceSize))
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}
	car := &model.Car{
		PieceCID:      model.CID(pieceCid),
		PieceSize:     int64(finalPieceSize),
		RootCID:       model.CID(assembler.rootCID),
		FileSize:      fileSize,
		StorageID:     storageID,
		StoragePath:   filename,
		AttachmentID:  &job.AttachmentID,
		PreparationID: job.Attachment.PreparationID,
		JobID:         &job.ID,
		PieceType:     model.DataPiece,
	}

	// Update all Files and FileRanges that have size == -1
	for fileID, length := range assembler.fileLengthCorrection {
		err := database.DoRetry(ctx, func() error {
			return db.Model(&model.File{}).Where("id = ?", fileID).Update("size", length).Error
		})
		if err != nil {
			return nil, errors.WithStack(err)
		}
		err = database.DoRetry(ctx, func() error {
			return db.Model(&model.FileRange{}).Where("file_id = ?", fileID).Update("length", length).Error
		})
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}
	for i := range job.FileRanges {
		if job.FileRanges[i].Length == -1 {
			job.FileRanges[i].Length = assembler.fileLengthCorrection[job.FileRanges[i].FileID]
			job.FileRanges[i].File.Size = assembler.fileLengthCorrection[job.FileRanges[i].FileID]
			logger.Warnw("correcting unknown file size", "path", job.FileRanges[i].File.Path, "length", job.FileRanges[i].Length)
		}
	}

	// Update all FileRange and file CID that are not split
	splitFileIDs := make(map[model.FileID]model.File)
	var updatedFiles []model.File
	splitFileBlks := make(map[model.FileID][]blocks.Block)
	for _, fileRange := range job.FileRanges {
		err = database.DoRetry(ctx, func() error {
			return db.Model(&model.FileRange{}).Where("id = ?", fileRange.ID).
				Update("cid", fileRange.CID).Error
		})
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if fileRange.Offset == 0 && fileRange.Length == fileRange.File.Size {
			err = database.DoRetry(ctx, func() error {
				return db.Model(&model.File{}).Where("id = ?", fileRange.FileID).
					Update("cid", fileRange.CID).Error
			})
			if err != nil {
				return nil, errors.WithStack(err)
			}
			fileRange.File.CID = fileRange.CID
			updatedFiles = append(updatedFiles, *fileRange.File)
		} else {
			splitFileIDs[fileRange.FileID] = *fileRange.File
		}
	}

	// Now check if all file ranges of a file are ready. If so, update file CID
	for fileID := range splitFileIDs {
		err = database.DoRetry(ctx, func() error {
			return db.Transaction(func(db *gorm.DB) error {
				var allParts []model.FileRange
				err = db.Where("file_id = ?", fileID).Order(clause.OrderByColumn{Column: clause.Column{Name: "offset"}}).Find(&allParts).Error
				if err != nil {
					return errors.WithStack(err)
				}
				if underscore.All(allParts, func(p model.FileRange) bool {
					return p.CID != model.CID(cid.Undef)
				}) {
					links := underscore.Map(allParts, func(p model.FileRange) format.Link {
						return format.Link{
							Size: uint64(p.Length),
							Cid:  cid.Cid(p.CID),
						}
					})
					blks, node, err := packutil.AssembleFileFromLinks(links)
					if err != nil {
						return errors.Wrap(err, "failed to assemble file from links")
					}
					nodeCid := node.Cid()
					err = db.Model(&model.File{}).Where("id = ?", fileID).Update("cid", model.CID(nodeCid)).Error
					if err != nil {
						return errors.Wrap(err, "failed to update cid of file")
					}
					file := splitFileIDs[fileID]
					file.CID = model.CID(nodeCid)
					updatedFiles = append(updatedFiles, file)
					splitFileBlks[fileID] = blks
				}
				return nil
			})
		})
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	car.NumOfFiles = int64(len(updatedFiles))
	err = database.DoRetry(ctx, func() error {
		return db.Transaction(
			func(db *gorm.DB) error {
				err := db.Create(&car).Error
				if err != nil {
					return errors.WithStack(err)
				}
				if !job.Attachment.Preparation.NoInline {
					for j := range assembler.carBlocks {
						assembler.carBlocks[j].CarID = car.ID
					}
					err = db.CreateInBatches(assembler.carBlocks, util.BatchSize).Error
					if err != nil {
						return errors.WithStack(err)
					}
				}
				return nil
			},
		)
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if !job.Attachment.Preparation.NoDag {
		err = database.DoRetry(ctx, func() error {
			if len(updatedFiles) == 0 {
				return nil
			}
			return db.Transaction(func(db *gorm.DB) error {
				var err error
				tree := daggen.NewDirectoryTree()
				var rootDirID model.DirectoryID
				for _, file := range updatedFiles {
					dirID := file.DirectoryID
					for {
						// Add the directory to tree if it is not there
						if !tree.Has(*dirID) {
							var dir model.Directory
							err = db.Where("id = ?", dirID).First(&dir).Error
							if err != nil {
								return errors.Wrap(err, "failed to get directory")
							}
							err = tree.Add(ctx, &dir)
							if err != nil {
								return errors.Wrap(err, "failed to add directory to tree")
							}
						}

						dirDetail := tree.Get(*dirID)

						// Update the directory for first iteration
						if dirID == file.DirectoryID {
							err = dirDetail.Data.AddFile(ctx, file.FileName(), cid.Cid(file.CID), uint64(file.Size))
							if err != nil {
								return errors.Wrap(err, "failed to add file to directory")
							}
							if blks, ok := splitFileBlks[file.ID]; ok {
								dirDetail.Data.AddBlocks(ctx, blks)
							}
						}

						// Next iteration
						dirID = dirDetail.Dir.ParentID
						if dirID == nil {
							rootDirID = dirDetail.Dir.ID
							break
						}
					}
				}
				// Recursively update all directory internal structure
				_, err = tree.Resolve(ctx, rootDirID)
				if err != nil {
					return errors.Wrap(err, "failed to resolve directory tree")
				}

				// Update all directories in the database
				for dirID, dirDetail := range tree.Cache() {
					bytes, err := dirDetail.Data.MarshalBinary(ctx)
					if err != nil {
						return errors.Wrap(err, "failed to marshall directory data")
					}
					node, _ := dirDetail.Data.Node()
					err = db.Model(&model.Directory{}).Where("id = ?", dirID).Updates(map[string]any{
						"cid":      model.CID(node.Cid()),
						"data":     bytes,
						"exported": false,
					}).Error
					if err != nil {
						return errors.Wrap(err, "failed to update directory")
					}
				}
				return nil
			})
		})
		if err != nil {
			return nil, errors.Wrap(err, "failed to update directory CIDs")
		}
	}

	logger.With("jobsID", job.ID).Info("finished packing")
	if job.Attachment.Preparation.DeleteAfterExport && len(job.Attachment.Preparation.OutputStorages) > 0 {
		logger.Info("Deleting original data source")
		for _, file := range updatedFiles {
			object := assembler.objects[file.ID]
			logger.Debugw("removing object", "path", object.Remote())
			err = object.Remove(ctx)
			if err != nil {
				logger.Warnw("failed to remove object", "error", err)
			}
		}
	}

	sourceType := job.Attachment.Storage.Type
	if job.Attachment.Storage.Config != nil {
		provider, ok := job.Attachment.Storage.Config["provider"]
		if ok {
			sourceType = sourceType + "-" + provider
		}
	}
	outputType := "inline"
	if len(job.Attachment.Preparation.OutputStorages) > 0 {
		outputType = job.Attachment.Preparation.OutputStorages[0].Type
	}

	packJobEvent := analytics.PackJobEvent{
		SourceType: sourceType,
		OutputType: outputType,
		PieceSize:  car.PieceSize,
		PieceCID:   car.PieceCID.String(),
		CarSize:    car.FileSize,
		NumOfFiles: car.NumOfFiles,
	}
	analytics.Default.QueuePushJobEvent(packJobEvent)
	return car, nil
}
