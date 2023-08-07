package migrate

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

func migrateDataset(ctx context.Context, mg *mongo.Client, db *gorm.DB, scanning ScanningRequest, skipItems bool) error {
	_, err := os.Stat(scanning.OutDir)
	if err != nil {
		log.Printf("[Warning] Output directory %s does not exist\n", scanning.OutDir)
	}
	ds, err := dataset.CreateHandler(db, dataset.CreateRequest{
		Name:                 scanning.Name,
		MaxSizeStr:           fmt.Sprintf("%d", scanning.MaxSize),
		OutputDir:            "",
		EncryptionRecipients: nil,
		EncryptionScript:     "",
	})
	if err != nil {
		return errors.Wrap(err, "failed to create dataset")
	}
	ds.OutputDir = scanning.OutDir
	err = db.Save(ds).Error
	if err != nil {
		return errors.Wrap(err, "failed to save dataset")
	}
	log.Printf("-- Created dataset %s\n", ds.Name)
	cliutil.PrintToConsole(ds, false, nil)

	sourceType := "local"
	path := scanning.Path
	metadata := map[string]any{
		"sourcePath":        path,
		"deleteAfterExport": false,
		"rescanInterval":    "0",
		"scanningState":     "complete",
	}
	if strings.HasPrefix(scanning.Path, "s3://") {
		sourceType = "s3"
		path = strings.TrimPrefix(scanning.Path, "s3://")
	}
	src, err := datasource.CreateDatasourceHandler(db, ctx, nil, sourceType, ds.Name, metadata)
	if err != nil {
		return errors.Wrap(err, "failed to create datasource")
	}
	log.Printf("-- Created datasource %s - %s\n", src.Type, path)
	cliutil.PrintToConsole(src, false, nil)
	rootDirectoryID, err := src.RootDirectoryID(db)
	if err != nil {
		return errors.Wrap(err, "failed to get root directory")
	}

	cursor, err := mg.Database("singularity").Collection("generationrequests").Find(
		ctx, bson.M{"datasetName": scanning.Name},
	)
	if err != nil {
		return errors.Wrap(err, "failed to query mongo for generation requests")
	}

	directoryCache := map[string]uint64{}
	var lastItem model.Item
	for cursor.Next(ctx) {
		var generation GenerationRequest
		err = cursor.Decode(&generation)
		if err != nil {
			return errors.Wrap(err, "failed to decode generation request")
		}

		chunk := model.Chunk{
			CreatedAt:    generation.CreatedAt,
			SourceID:     src.ID,
			PackingState: model.Complete,
			ErrorMessage: generation.ErrorMessage,
		}
		err = db.Create(&chunk).Error
		if err != nil {
			return errors.Wrap(err, "failed to create chunk")
		}
		log.Printf("-- Created chunk %d for %s\n", chunk.ID, ds.Name)

		pieceCID, err := cid.Parse(generation.PieceCID)
		if err != nil {
			log.Printf("failed to parse piece cid %s\n", generation.PieceCID)
			pieceCID = cid.Undef
		}
		dataCID, err := cid.Parse(generation.DataCID)
		if err != nil {
			log.Printf("failed to parse data cid %s\n", generation.DataCID)
			dataCID = cid.Undef
		}
		fileName := fmt.Sprintf("%s.car", generation.PieceCID)
		if generation.FilenameOverride != "" {
			fileName = generation.FilenameOverride
		}
		car := model.Car{
			CreatedAt: generation.CreatedAt,
			PieceCID:  model.CID(pieceCID),
			PieceSize: int64(generation.PieceSize),
			RootCID:   model.CID(dataCID),
			FileSize:  int64(generation.CarSize),
			FilePath:  filepath.Join(scanning.OutDir, fileName),
			DatasetID: ds.ID,
			SourceID:  &src.ID,
			ChunkID:   &chunk.ID,
		}
		err = db.Create(&car).Error
		if err != nil {
			return errors.Wrap(err, "failed to create car")
		}
		log.Printf("-- Created car %s for %s\n", generation.PieceCID, ds.Name)

		if skipItems {
			continue
		}
		cursor, err := mg.Database("singularity").Collection("outputfilelists").Find(
			ctx,
			bson.M{"generationId": generation.ID.Hex()},
			options.Find().SetSort(bson.M{"index": 1}))
		if err != nil {
			return errors.Wrap(err, "failed to query mongo for output file lists")
		}
		var items []model.Item
		for cursor.Next(ctx) {
			var fileList OutputFileList
			err = cursor.Decode(&fileList)
			if err != nil {
				return errors.Wrap(err, "failed to decode output file list")
			}
			for _, file := range fileList.GeneratedFileList {
				if file.CID == "unrecoverable" {
					continue
				}
				if file.Dir {
					continue
				}
				fileCID, err := cid.Parse(file.CID)
				if err != nil {
					return errors.Wrapf(err, "failed to parse file cid %s", file.CID)
				}

				var item model.Item
				if file.IsComplete() {
					item = model.Item{
						CreatedAt: generation.CreatedAt,
						SourceID:  src.ID,
						Path:      file.Path,
						Size:      int64(file.Size),
						CID:       model.CID(fileCID),
						ItemParts: []model.ItemPart{
							{
								Offset:  0,
								Length:  int64(file.Size),
								CID:     model.CID(fileCID),
								ChunkID: &chunk.ID,
							},
						},
					}
				} else if file.Start == 0 {
					lastItem = model.Item{
						CreatedAt: generation.CreatedAt,
						SourceID:  src.ID,
						Path:      file.Path,
						Size:      int64(file.Size),
						CID:       model.CID(cid.Undef),
						ItemParts: []model.ItemPart{
							{
								Offset:  0,
								Length:  int64(file.End),
								CID:     model.CID(fileCID),
								ChunkID: &chunk.ID,
							},
						},
					}
					continue
				} else {
					lastItem.ItemParts = append(lastItem.ItemParts, model.ItemPart{
						Offset:  int64(file.Start),
						Length:  int64(file.End - file.Start),
						CID:     model.CID(fileCID),
						ChunkID: &chunk.ID,
					})
					if file.End < file.Size {
						continue
					} else {
						item = lastItem
						lastItem = model.Item{}
						links := make([]format.Link, 0)
						for _, part := range item.ItemParts {
							links = append(links, format.Link{
								Size: uint64(part.Length),
								Cid:  cid.Cid(part.CID),
							})
						}
						_, root, err := pack.AssembleItemFromLinks(links)
						if err != nil {
							return errors.Wrap(err, "failed to assemble item from links")
						}
						item.CID = model.CID(root.Cid())
					}
				}
				err = datasource.EnsureParentDirectories(db, &item, rootDirectoryID, directoryCache)
				if err != nil {
					return errors.Wrap(err, "failed to ensure parent directories")
				}
				directory, ok := directoryCache[filepath.Dir(item.Path)]
				if !ok {
					return errors.Errorf("directory %s not found in cache", filepath.Dir(item.Path))
				}
				item.DirectoryID = &directory
				items = append(items, item)
			}
		}
		if len(items) > 0 {
			err = db.CreateInBatches(&items, 1000).Error
			if err != nil {
				return errors.Wrap(err, "failed to create items")
			}
		}
		log.Printf("-- Created %d items for %s\n", len(items), ds.Name)
	}

	return nil
}

func MigrateDataset(cctx *cli.Context) error {
	skipItems := cctx.Bool("skip-items")
	datasource.ValidateSource = false
	log.Println("Migrating dataset from old singularity database")
	mongoConnectionString := cctx.String("mongo-connection-string")
	sqlConnectionString := cctx.String("database-connection-string")
	log.Printf("Using mongo connection string: %s\n", mongoConnectionString)
	log.Printf("Using sql connection string: %s\n", sqlConnectionString)
	db, closer, err := database.OpenFromCLI(cctx)
	if err != nil {
		return err
	}
	defer closer.Close()
	ctx := cctx.Context
	db = db.WithContext(ctx)
	mg, err := mongo.Connect(ctx, options.Client().ApplyURI(cctx.String("mongo-connection-string")))
	if err != nil {
		return errors.Wrap(err, "failed to connect to mongo")
	}

	err = model.AutoMigrate(db)
	if err != nil {
		return errors.Wrap(err, "failed to auto-migrate database")
	}

	resp, err := mg.Database("singularity").Collection("scanningrequests").Find(ctx, bson.M{})
	if err != nil {
		return errors.Wrap(err, "failed to query mongo for scanning requests")
	}

	var scannings []ScanningRequest
	err = resp.All(ctx, &scannings)
	if err != nil {
		return errors.Wrap(err, "failed to decode mongo response")
	}

	for _, scanning := range scannings {
		var datasetExists int64
		err = db.Model(&model.Dataset{}).Where("name = ?", scanning.Name).Count(&datasetExists).Error
		if err != nil {
			return errors.Wrapf(err, "failed to query for dataset %s", scanning.Name)
		}
		if datasetExists > 0 {
			log.Printf("Dataset %s already exists, skipping\n", scanning.Name)
			continue
		}
		log.Printf("Migrating Dataset: %s\n", scanning.Name)
		err = migrateDataset(ctx, mg, db, scanning, skipItems)
		if err != nil {
			return errors.Wrapf(err, "failed to migrate dataset %s", scanning.Name)
		}
	}
	return nil
}
