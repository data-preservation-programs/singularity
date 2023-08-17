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
	"github.com/data-preservation-programs/singularity/util"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

//nolint:gocritic
func migrateDataset(ctx context.Context, mg *mongo.Client, db *gorm.DB, scanning ScanningRequest, skipFiles bool) error {
	_, err := os.Stat(scanning.OutDir)
	if err != nil {
		log.Printf("[Warning] Output directory %s does not exist\n", scanning.OutDir)
	}
	ds, err := dataset.CreateHandler(ctx, db, dataset.CreateRequest{
		Name:                 scanning.Name,
		MaxSizeStr:           fmt.Sprintf("%d", scanning.MaxSize),
		OutputDirs:           nil,
		EncryptionRecipients: nil,
	})
	if err != nil {
		return errors.Wrap(err, "failed to create dataset")
	}
	ds.OutputDirs = []string{scanning.OutDir}
	err = db.Save(ds).Error
	if err != nil {
		return errors.Wrap(err, "failed to save dataset")
	}
	log.Printf("-- Created dataset %s\n", ds.Name)
	cliutil.PrintToConsole(ds, false, nil)

	sourceType := "local"
	path := scanning.Path
	if strings.HasPrefix(scanning.Path, "s3://") {
		sourceType = "s3"
		path = strings.TrimPrefix(scanning.Path, "s3://")
	}
	metadata := map[string]any{
		"sourcePath":        path,
		"deleteAfterExport": false,
		"rescanInterval":    "0",
		"scanningState":     "complete",
	}
	src, err := datasource.CreateDatasourceHandler(ctx, db, sourceType, ds.Name, metadata)
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
	directoryCache[""] = rootDirectoryID
	directoryCache["."] = rootDirectoryID
	var lastFile model.File
	for cursor.Next(ctx) {
		var generation GenerationRequest
		err = cursor.Decode(&generation)
		if err != nil {
			return errors.Wrap(err, "failed to decode generation request")
		}

		packJob := model.PackJob{
			CreatedAt:    generation.CreatedAt,
			SourceID:     src.ID,
			PackingState: model.Complete,
			ErrorMessage: generation.ErrorMessage,
		}
		err = db.Create(&packJob).Error
		if err != nil {
			return errors.Wrap(err, "failed to create pack job")
		}
		log.Printf("-- Created pack job %d for %s\n", packJob.ID, ds.Name)

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
			PackJobID: &packJob.ID,
		}
		err = db.Create(&car).Error
		if err != nil {
			return errors.Wrap(err, "failed to create car")
		}
		log.Printf("-- Created car %s for %s\n", generation.PieceCID, ds.Name)

		if skipFiles {
			continue
		}
		cursor, err := mg.Database("singularity").Collection("outputfilelists").Find(
			ctx,
			bson.M{"generationId": generation.ID.Hex()},
			options.Find().SetSort(bson.M{"index": 1}))
		if err != nil {
			return errors.Wrap(err, "failed to query mongo for output file lists")
		}
		var files []model.File
		for cursor.Next(ctx) {
			var fileList OutputFileList
			err = cursor.Decode(&fileList)
			if err != nil {
				return errors.Wrap(err, "failed to decode output file list")
			}
			for _, generatedFile := range fileList.GeneratedFileList {
				if generatedFile.CID == "unrecoverable" {
					continue
				}
				if generatedFile.Dir {
					continue
				}
				fileCID, err := cid.Parse(generatedFile.CID)
				if err != nil {
					return errors.Wrapf(err, "failed to parse file cid %s", generatedFile.CID)
				}

				var file model.File
				if generatedFile.IsComplete() {
					file = model.File{
						CreatedAt: generation.CreatedAt,
						SourceID:  src.ID,
						Path:      generatedFile.Path,
						Size:      int64(generatedFile.Size),
						CID:       model.CID(fileCID),
						FileRanges: []model.FileRange{
							{
								Offset:    0,
								Length:    int64(generatedFile.Size),
								CID:       model.CID(fileCID),
								PackJobID: &packJob.ID,
							},
						},
					}
				} else if generatedFile.Start == 0 {
					lastFile = model.File{
						CreatedAt: generation.CreatedAt,
						SourceID:  src.ID,
						Path:      generatedFile.Path,
						Size:      int64(generatedFile.Size),
						CID:       model.CID(cid.Undef),
						FileRanges: []model.FileRange{
							{
								Offset:    0,
								Length:    int64(generatedFile.End),
								CID:       model.CID(fileCID),
								PackJobID: &packJob.ID,
							},
						},
					}
					continue
				} else {
					lastFile.FileRanges = append(lastFile.FileRanges, model.FileRange{
						Offset:    int64(generatedFile.Start),
						Length:    int64(generatedFile.End - generatedFile.Start),
						CID:       model.CID(fileCID),
						PackJobID: &packJob.ID,
					})
					if generatedFile.End < generatedFile.Size {
						continue
					} else {
						file = lastFile
						lastFile = model.File{}
						links := make([]format.Link, 0)
						for _, part := range file.FileRanges {
							links = append(links, format.Link{
								Size: uint64(part.Length),
								Cid:  cid.Cid(part.CID),
							})
						}
						_, root, err := pack.AssembleItemFromLinks(links)
						if err != nil {
							return errors.Wrap(err, "failed to assemble file from links")
						}
						file.CID = model.CID(root.Cid())
					}
				}
				err = datasource.EnsureParentDirectories(ctx, db, &file, rootDirectoryID, directoryCache)
				if err != nil {
					return errors.Wrap(err, "failed to ensure parent directories")
				}
				directory, ok := directoryCache[filepath.Dir(file.Path)]
				if !ok {
					return errors.Errorf("directory %s not found in cache", filepath.Dir(file.Path))
				}
				file.DirectoryID = &directory
				files = append(files, file)
			}
		}
		if len(files) > 0 {
			err = db.CreateInBatches(&files, util.BatchSize).Error
			if err != nil {
				return errors.Wrap(err, "failed to create files")
			}
		}
		log.Printf("-- Created %d files for %s\n", len(files), ds.Name)
	}

	return nil
}

func MigrateDataset(cctx *cli.Context) error {
	skipFiles := cctx.Bool("skip-files")
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
	mg, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConnectionString))
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
		err = migrateDataset(ctx, mg, db, scanning, skipFiles)
		if err != nil {
			return errors.Wrapf(err, "failed to migrate dataset %s", scanning.Name)
		}
	}
	return nil
}
