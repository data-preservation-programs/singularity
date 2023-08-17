package datasource

import (
	"context"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack"
	"github.com/data-preservation-programs/singularity/pack/daggen"
	"github.com/data-preservation-programs/singularity/util"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/pkg/errors"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func PackHandler(
	db *gorm.DB,
	ctx context.Context,
	resolver datasource.HandlerResolver,
	packJobID uint64,
) ([]model.Car, error) {
	return packHandler(db, ctx, resolver, packJobID)
}

// @Summary Pack a pack job into car files
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Pack job ID"
// @Success 201 {object} []model.Car
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /packjob/{id}/pack [post]
func packHandler(
	db *gorm.DB,
	ctx context.Context,
	resolver datasource.HandlerResolver,
	packJobID uint64,
) ([]model.Car, error) {
	var packJob model.PackJob
	err := db.Where("id = ?", packJobID).Find(&packJob).Error
	if err != nil {
		return nil, err
	}

	return Pack(ctx, db, packJob, resolver)
}

func Pack(
	ctx context.Context,
	db *gorm.DB,
	packJob model.PackJob,
	resolver datasource.HandlerResolver,
) ([]model.Car, error) {
	db = db.WithContext(ctx)
	var outDir string
	if len(packJob.Source.Dataset.OutputDirs) > 0 {
		outDir = packJob.Source.Dataset.OutputDirs[0]
	}
	logger.Debugw("Use output dir", "dir", outDir)
	handler, err := resolver.Resolve(ctx, *packJob.Source)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get datasource handler")
	}
	result, err := pack.AssembleCar(ctx, handler, *packJob.Source.Dataset,
		packJob.FileRanges, outDir, packJob.Source.Dataset.PieceSize)
	if err != nil {
		return nil, errors.Wrap(err, "failed to pack items")
	}

	// Update all FileRange and Item CID that are not split
	splitFileIDs := make(map[uint64]model.File)
	var updatedFiles []model.File
	splitFileBlks := make(map[uint64][]blocks.Block)
	for i, fileRange := range packJob.FileRanges {
		fileRangeID := fileRange.ID
		fileRangeCID := result.FileRangeCIDs[fileRangeID]
		packJob.FileRanges[i].CID = model.CID(fileRangeCID)
		logger.Debugw("update file range CID", "fileRangeID", fileRangeID, "CID", fileRangeCID.String())
		err = database.DoRetry(ctx, func() error {
			return db.Model(&model.FileRange{}).Where("id = ?", fileRangeID).
				Update("cid", model.CID(fileRangeCID)).Error
		})
		if err != nil {
			return nil, errors.Wrap(err, "failed to update cid of file part")
		}
		if fileRange.Offset == 0 && fileRange.Length == fileRange.File.Size {
			fileRange.File.CID = model.CID(fileRangeCID)
			logger.Debugw("update file CID", "fileID", fileRange.FileID, "CID", fileRangeCID.String())
			err = database.DoRetry(ctx, func() error {
				return db.Model(&model.File{}).Where("id = ?", fileRange.FileID).
					Update("cid", model.CID(fileRangeCID)).Error
			})
			if err != nil {
				return nil, errors.Wrap(err, "failed to update cid of file")
			}
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
					return errors.Wrap(err, "failed to get all item parts")
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
					blks, node, err := pack.AssembleFileFromLinks(links)
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
	}

	logger.Debugw("create car for finished packJob", "packJobID", packJob.ID)
	var cars []model.Car
	err = database.DoRetry(ctx, func() error {
		return db.Transaction(
			func(db *gorm.DB) error {
				for _, result := range result.CarResults {
					car := model.Car{
						PieceCID:  model.CID(result.PieceCID),
						PieceSize: result.PieceSize,
						RootCID:   model.CID(result.RootCID),
						FileSize:  result.CarFileSize,
						FilePath:  result.CarFilePath,
						PackJobID: &packJob.ID,
						DatasetID: packJob.Source.DatasetID,
						SourceID:  &packJob.SourceID,
						Header:    result.Header,
					}
					err := db.Create(&car).Error
					if err != nil {
						return errors.Wrap(err, "failed to create car")
					}
					for i := range result.CarBlocks {
						result.CarBlocks[i].CarID = car.ID
					}
					err = db.CreateInBatches(&result.CarBlocks, util.BatchSize).Error
					if err != nil {
						return errors.Wrap(err, "failed to create car blocks")
					}
					cars = append(cars, car)
				}
				return nil
			},
		)
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to save car")
	}

	logger.Debugw("update directory data", "packJobID", packJob.ID)
	err = database.DoRetry(ctx, func() error {
		if len(updatedFiles) == 0 {
			return nil
		}
		return db.Transaction(func(db *gorm.DB) error {
			var err error
			tree := daggen.NewDirectoryTree()
			var rootDirID uint64
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
						err = dirDetail.Data.AddItem(ctx, file.Name(), cid.Cid(file.CID), uint64(file.Size))
						if err != nil {
							return errors.Wrap(err, "failed to add file to directory")
						}
						if blks, ok := splitFileBlks[file.ID]; ok {
							err = dirDetail.Data.AddBlocks(ctx, blks)
							if err != nil {
								return errors.Wrap(err, "failed to add blocks to directory")
							}
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

	logger.With("pack_job", packJob.ID).Info("finished packing")
	if packJob.Source.DeleteAfterExport && result.CarResults[0].CarFilePath != "" {
		logger.Info("Deleting original data source")
		for _, file := range updatedFiles {
			object := result.Objects[file.ID]
			logger.Debugw("removing object", "path", object.Remote())
			err = object.Remove(ctx)
			if err != nil {
				logger.Warnw("failed to remove object", "error", err)
			}
		}
	}
	return cars, nil
}
