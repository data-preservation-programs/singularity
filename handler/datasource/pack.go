package datasource

import (
	"context"
	"strings"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack"
	"github.com/data-preservation-programs/singularity/pack/daggen"
	"github.com/data-preservation-programs/singularity/pack/device"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/pkg/errors"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
)

func PackHandler(
	db *gorm.DB,
	ctx context.Context,
	resolver datasource.HandlerResolver,
	chunkID uint64,
) ([]model.Car, error) {
	return packHandler(db, ctx, resolver, chunkID)
}

// @Summary Pack a chunk into car files
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Chunk ID"
// @Success 201 {object} []model.Car
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /chunk/{id}/pack [post]
func packHandler(
	db *gorm.DB,
	ctx context.Context,
	resolver datasource.HandlerResolver,
	chunkID uint64,
) ([]model.Car, error) {
	var chunk model.Chunk
	err := db.Where("id = ?", chunkID).Find(&chunk).Error
	if err != nil {
		return nil, err
	}

	return Pack(ctx, db, chunk, resolver)
}

func Pack(
	ctx context.Context,
	db *gorm.DB,
	chunk model.Chunk,
	resolver datasource.HandlerResolver,
) ([]model.Car, error) {
	var outDir string
	if len(chunk.Source.Dataset.OutputDirs) > 0 {
		var err error
		outDir, err = device.GetPathWithMostSpace(chunk.Source.Dataset.OutputDirs)
		if err != nil {
			logger.Warnw("failed to get path with most space. using the first one", "error", err)
			outDir = chunk.Source.Dataset.OutputDirs[0]
		}
	}
	logger.Debugw("Use output dir", "dir", outDir)
	handler, err := resolver.Resolve(ctx, *chunk.Source)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get datasource handler")
	}
	result, err := pack.AssembleCar(ctx, handler, *chunk.Source.Dataset,
		chunk.ItemParts, outDir, chunk.Source.Dataset.PieceSize)
	if err != nil {
		return nil, errors.Wrap(err, "failed to pack items")
	}

	for _, itemPart := range chunk.ItemParts {
		itemPartID := itemPart.ID
		itemPartCID, ok := result.ItemPartCIDs[itemPartID]
		if !ok {
			return nil, errors.New("item part not found in result")
		}
		logger.Debugw("update item part CID", "itemPartID", itemPartID, "CID", itemPartCID.String())
		err = database.DoRetry(func() error {
			return db.Model(&model.ItemPart{}).Where("id = ?", itemPartID).
				Update("cid", model.CID(itemPartCID)).Error
		})
		if err != nil {
			return nil, errors.Wrap(err, "failed to update cid of item")
		}
		logger.Debugw("update item CID", "itemID", itemPart.ItemID, "CID", itemPartCID.String())
		if itemPart.Offset == 0 && itemPart.Length == itemPart.Item.Size {
			err = database.DoRetry(func() error {
				return db.Model(&model.Item{}).Where("id = ?", itemPart.ItemID).
					Update("cid", model.CID(itemPartCID)).Error
			})
			if err != nil {
				return nil, errors.Wrap(err, "failed to update cid of item")
			}
		}
	}

	logger.Debugw("create car for finished chunk", "chunkID", chunk.ID)
	var cars []model.Car
	err = database.DoRetry(func() error {
		return db.Transaction(
			func(db *gorm.DB) error {
				for _, result := range result.CarResults {
					car := model.Car{
						PieceCID:  model.CID(result.PieceCID),
						PieceSize: result.PieceSize,
						RootCID:   model.CID(result.RootCID),
						FileSize:  result.CarFileSize,
						FilePath:  result.CarFilePath,
						ChunkID:   &chunk.ID,
						DatasetID: chunk.Source.DatasetID,
						SourceID:  &chunk.SourceID,
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

	logger.Debugw("update directory data", "chunkID", chunk.ID)
	err = database.DoRetry(func() error {
		return db.Transaction(func(db *gorm.DB) error {
			dirCache := make(map[uint64]*daggen.DirectoryData)
			childrenCache := make(map[uint64][]uint64)
			for _, itemPart := range chunk.ItemParts {
				dirID := itemPart.Item.DirectoryID
				for dirID != nil {
					dirData, ok := dirCache[*dirID]
					if !ok {
						dirData = &daggen.DirectoryData{}
						var dir model.Directory
						err := db.Where("id = ?", dirID).First(&dir).Error
						if err != nil {
							return errors.Wrap(err, "failed to get directory")
						}

						err = dirData.UnmarshallBinary(dir.Data)
						if err != nil {
							return errors.Wrap(err, "failed to unmarshall directory data")
						}
						dirData.Directory = dir
						dirCache[*dirID] = dirData
						if dir.ParentID != nil {
							childrenCache[*dir.ParentID] = append(childrenCache[*dir.ParentID], *dirID)
						}
					}

					// Update the directory for first iteration
					if dirID == itemPart.Item.DirectoryID {
						itemPartID := itemPart.ID
						itemPartCID, ok := result.ItemPartCIDs[itemPartID]
						if !ok {
							return errors.New("item part not found in result")
						}
						err = db.Model(&model.ItemPart{}).Where("id = ?", itemPartID).
							Update("cid", model.CID(itemPartCID)).Error
						if err != nil {
							return errors.Wrap(err, "failed to update cid of item")
						}
						name := itemPart.Item.Path[strings.LastIndex(itemPart.Item.Path, "/")+1:]
						if itemPart.Offset == 0 && itemPart.Length == itemPart.Item.Size {
							partCID := result.ItemPartCIDs[itemPart.ID]
							err = dirData.AddItem(name, partCID, uint64(itemPart.Length))
							if err != nil {
								return errors.Wrap(err, "failed to add item to directory")
							}
							/*
								err = db.Model(&model.Item{}).Where("id = ?", itemPart.ItemID).Update("cid", model.CID(itemPartCID)).Error
								if err != nil {
									return errors.Wrap(err, "failed to update cid of item")
								}
							*/
						} else {
							var allParts []model.ItemPart
							err = db.Where("item_id = ?", itemPart.ItemID).Order("\"offset\" asc").Find(&allParts).Error
							if err != nil {
								return errors.Wrap(err, "failed to get all item parts")
							}
							if underscore.All(allParts, func(p model.ItemPart) bool {
								return p.CID != model.CID(cid.Undef)
							}) {
								links := underscore.Map(allParts, func(p model.ItemPart) format.Link {
									return format.Link{
										Size: uint64(p.Length),
										Cid:  cid.Cid(p.CID),
									}
								})
								c, err := dirData.AddItemFromLinks(name, links)
								if err != nil {
									return errors.Wrap(err, "failed to add item to directory")
								}
								err = db.Model(&model.Item{}).Where("id = ?", itemPart.ItemID).Update("cid", model.CID(c)).Error
								if err != nil {
									return errors.Wrap(err, "failed to update cid of item")
								}
							}
						}
					}

					// Next iteration
					dirID = dirData.Directory.ParentID
				}
			}
			// Recursively update all directory internal structure
			rootDirID, err := chunk.Source.RootDirectoryID(db)
			if err != nil {
				return errors.Wrap(err, "failed to get root directory id")
			}
			_, err = daggen.ResolveDirectoryTree(rootDirID, dirCache, childrenCache)
			if err != nil {
				return errors.Wrap(err, "failed to resolve directory tree")
			}
			// Update all directories in the database
			for dirID, dirData := range dirCache {
				bytes, err := dirData.MarshalBinary()
				if err != nil {
					return errors.Wrap(err, "failed to marshall directory data")
				}
				err = db.Model(&model.Directory{}).Where("id = ?", dirID).Updates(map[string]any{
					"cid":      model.CID(dirData.Node.Cid()),
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

	logger.With("chunk_id", chunk.ID).Info("finished packing")
	if chunk.Source.DeleteAfterExport && result.CarResults[0].CarFilePath != "" {
		logger.Info("Deleting original data source")
		handled := map[uint64]struct{}{}
		for _, itemPart := range chunk.ItemParts {
			if _, ok := handled[itemPart.ItemID]; ok {
				continue
			}
			handled[itemPart.ItemID] = struct{}{}
			object := result.Objects[itemPart.ItemID]
			if itemPart.Offset == 0 && itemPart.Length == itemPart.Item.Size {
				logger.Debugw("removing object", "path", object.Remote())
				err = object.Remove(ctx)
				if err != nil {
					logger.Warnw("failed to remove object", "error", err)
				}
				continue
			}
			// Make sure all parts of this file has been exported before deleting
			var unfinishedCount int64
			err = db.Model(&model.ItemPart{}).
				Where("item_id = ? AND cid IS NULL", itemPart.ItemID).Count(&unfinishedCount).Error
			if err != nil {
				logger.Warnw("failed to get count for unfinished item parts", "error", err)
				continue
			}
			if unfinishedCount > 0 {
				logger.Info("not all items have been exported yet, skipping delete")
				continue
			}
			logger.Debugw("removing object", "path", object.Remote())
			err = object.Remove(ctx)
			if err != nil {
				logger.Warnw("failed to remove object", "error", err)
			}
		}
	}
	return cars, nil
}
