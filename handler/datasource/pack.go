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
	db = db.WithContext(ctx)
	var outDir string
	if len(chunk.Source.Dataset.OutputDirs) > 0 {
		outDir = chunk.Source.Dataset.OutputDirs[0]
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

	// Update all ItemPart and Item CID that are not split
	splitItemIDs := make(map[uint64]model.Item)
	var updatedItems []model.Item
	splitItemBlks := make(map[uint64][]blocks.Block)
	for i, itemPart := range chunk.ItemParts {
		itemPartID := itemPart.ID
		itemPartCID := result.ItemPartCIDs[itemPartID]
		chunk.ItemParts[i].CID = model.CID(itemPartCID)
		logger.Debugw("update item part CID", "itemPartID", itemPartID, "CID", itemPartCID.String())
		err = database.DoRetry(ctx, func() error {
			return db.Model(&model.ItemPart{}).Where("id = ?", itemPartID).
				Update("cid", model.CID(itemPartCID)).Error
		})
		if err != nil {
			return nil, errors.Wrap(err, "failed to update cid of item part")
		}
		if itemPart.Offset == 0 && itemPart.Length == itemPart.Item.Size {
			itemPart.Item.CID = model.CID(itemPartCID)
			logger.Debugw("update item CID", "itemID", itemPart.ItemID, "CID", itemPartCID.String())
			err = database.DoRetry(ctx, func() error {
				return db.Model(&model.Item{}).Where("id = ?", itemPart.ItemID).
					Update("cid", model.CID(itemPartCID)).Error
			})
			if err != nil {
				return nil, errors.Wrap(err, "failed to update cid of item")
			}
			updatedItems = append(updatedItems, *itemPart.Item)
		} else {
			splitItemIDs[itemPart.ItemID] = *itemPart.Item
		}
	}

	// Now check if all item parts of an item are ready. If so, update item CID
	for itemID := range splitItemIDs {
		err = database.DoRetry(ctx, func() error {
			return db.Transaction(func(db *gorm.DB) error {
				var allParts []model.ItemPart
				err = db.Where("item_id = ?", itemID).Order("offset asc").Find(&allParts).Error
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
					blks, node, err := pack.AssembleItemFromLinks(links)
					if err != nil {
						return errors.Wrap(err, "failed to assemble item from links")
					}
					nodeCid := node.Cid()
					err = db.Model(&model.Item{}).Where("id = ?", itemID).Update("cid", model.CID(nodeCid)).Error
					if err != nil {
						return errors.Wrap(err, "failed to update cid of item")
					}
					item := splitItemIDs[itemID]
					item.CID = model.CID(nodeCid)
					updatedItems = append(updatedItems, item)
					splitItemBlks[itemID] = blks
				}
				return nil
			})
		})
	}

	logger.Debugw("create car for finished chunk", "chunkID", chunk.ID)
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
	err = database.DoRetry(ctx, func() error {
		if len(updatedItems) == 0 {
			return nil
		}
		return db.Transaction(func(db *gorm.DB) error {
			var err error
			tree := daggen.NewDirectoryTree()
			var rootDirID uint64
			for _, item := range updatedItems {
				dirID := item.DirectoryID
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
					if dirID == item.DirectoryID {
						err = dirDetail.Data.AddItem(ctx, item.Name(), cid.Cid(item.CID), uint64(item.Size))
						if err != nil {
							return errors.Wrap(err, "failed to add item to directory")
						}
						if blks, ok := splitItemBlks[item.ID]; ok {
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

	logger.With("chunk_id", chunk.ID).Info("finished packing")
	if chunk.Source.DeleteAfterExport && result.CarResults[0].CarFilePath != "" {
		logger.Info("Deleting original data source")
		for _, item := range updatedItems {
			object := result.Objects[item.ID]
			logger.Debugw("removing object", "path", object.Remote())
			err = object.Remove(ctx)
			if err != nil {
				logger.Warnw("failed to remove object", "error", err)
			}
		}
	}
	return cars, nil
}
