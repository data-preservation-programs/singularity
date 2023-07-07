package datasetworker

import (
	"context"
	"strings"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack"
	"github.com/data-preservation-programs/singularity/pack/daggen"
	"github.com/data-preservation-programs/singularity/pack/device"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/pkg/errors"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
)

func (w *DatasetWorkerThread) pack(
	ctx context.Context, chunk model.Chunk,
) error {
	var outDir string
	if len(chunk.Source.Dataset.OutputDirs) > 0 {
		var err error
		outDir, err = device.GetPathWithMostSpace(chunk.Source.Dataset.OutputDirs)
		if err != nil {
			w.logger.Warnw("failed to get path with most space. using the first one", "error", err)
			outDir = chunk.Source.Dataset.OutputDirs[0]
		}
	}
	handler, err := w.datasourceHandlerResolver.Resolve(ctx, *chunk.Source)
	if err != nil {
		return errors.Wrap(err, "failed to get datasource handler")
	}
	result, err := pack.AssembleCar(ctx, handler, *chunk.Source.Dataset,
		chunk.ItemParts, outDir, chunk.Source.Dataset.PieceSize)
	if err != nil {
		return errors.Wrap(err, "failed to pack items")
	}

	for _, itemPart := range chunk.ItemParts {
		itemPartID := itemPart.ID
		itemPartCID, ok := result.ItemPartCIDs[itemPartID]
		if !ok {
			return errors.New("item part not found in result")
		}
		err = database.DoRetry(func() error {
			return w.db.Model(&model.ItemPart{}).Where("id = ?", itemPartID).
				Update("cid", model.CID(itemPartCID)).Error
		})
		if err != nil {
			return errors.Wrap(err, "failed to update cid of item")
		}
		if itemPart.Offset == 0 && itemPart.Length == itemPart.Item.Size {
			err = database.DoRetry(func() error {
				return w.db.Model(&model.Item{}).Where("id = ?", itemPart.ItemID).
					Update("cid", model.CID(itemPartCID)).Error
			})
			if err != nil {
				return errors.Wrap(err, "failed to update cid of item")
			}
		}
	}

	err = database.DoRetry(func() error {
		return w.db.Transaction(
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
					err = db.CreateInBatches(&result.CarBlocks, 1000).Error
					if err != nil {
						return errors.Wrap(err, "failed to create car blocks")
					}
				}
				return nil
			},
		)
	})
	if err != nil {
		return errors.Wrap(err, "failed to save car")
	}

	// Update all directory CIDs
	err = database.DoRetry(func() error {
		return w.db.Transaction(func(db *gorm.DB) error {
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
			rootDirID, err := chunk.Source.RootDirectoryID(w.db)
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
				err = db.Model(&model.Directory{}).Where("id = ?", dirID).Updates(map[string]interface{}{
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
		return errors.Wrap(err, "failed to update directory CIDs")
	}

	w.logger.With("chunk_id", chunk.ID).Info("finished packing")
	if chunk.Source.DeleteAfterExport && result.CarResults[0].CarFilePath != "" {
		w.logger.Info("Deleting original data source")
		handled := map[uint64]struct{}{}
		for _, itemPart := range chunk.ItemParts {
			if _, ok := handled[itemPart.ItemID]; ok {
				continue
			}
			handled[itemPart.ItemID] = struct{}{}
			object := result.Objects[itemPart.ItemID]
			if itemPart.Offset == 0 && itemPart.Length == itemPart.Item.Size {
				w.logger.Debugw("removing object", "path", object.Remote())
				err = object.Remove(ctx)
				if err != nil {
					w.logger.Warnw("failed to remove object", "error", err)
				}
				continue
			}
			// Make sure all parts of this file has been exported before deleting
			var unfinishedCount int64
			err = w.db.Model(&model.ItemPart{}).
				Where("item_id = ? AND cid IS NULL", itemPart.ItemID).Count(&unfinishedCount).Error
			if err != nil {
				w.logger.Warnw("failed to get count for unfinished item parts", "error", err)
				continue
			}
			if unfinishedCount > 0 {
				w.logger.Info("not all items have been exported yet, skipping delete")
				continue
			}
			w.logger.Debugw("removing object", "path", object.Remote())
			err = object.Remove(ctx)
			if err != nil {
				w.logger.Warnw("failed to remove object", "error", err)
			}
		}
	}
	return nil
}
