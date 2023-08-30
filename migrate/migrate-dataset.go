package migrate

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	util2 "github.com/data-preservation-programs/singularity/pack/packutil"
	"github.com/data-preservation-programs/singularity/pack/push"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

//nolint:gocritic
func migrateDataset(ctx context.Context, mg *mongo.Client, db *gorm.DB, scanning ScanningRequest, skipFiles bool) error {
	sourceType := "local"
	path := scanning.Path
	config := make(map[string]string)
	if strings.HasPrefix(scanning.Path, "s3://") {
		sourceType = "s3"
		path = strings.TrimPrefix(scanning.Path, "s3://")
		config["provider"] = "AWS"
	}

	preparation := model.Preparation{
		MaxSize:   int64(scanning.MaxSize),
		PieceSize: int64(util.NextPowerOfTwo(scanning.MaxSize)),
		SourceStorages: []model.Storage{{
			Name:   scanning.Name + "-source",
			Type:   sourceType,
			Path:   path,
			Config: config,
		}},
		OutputStorages: []model.Storage{{
			Name: scanning.Name + "-output",
			Type: "local",
			Path: scanning.OutDir,
		}},
	}

	err := db.Create(&preparation).Error
	if err != nil {
		return errors.WithStack(err)
	}

	var attachment model.SourceAttachment
	err = db.Where("preparation_id = ? AND storage_id = ?", preparation.ID, preparation.SourceStorages[0].ID).First(&attachment).Error
	if err != nil {
		return errors.WithStack(err)
	}

	rootDir := model.Directory{
		AttachmentID: attachment.ID,
		Name:         path,
	}
	err = db.Create(&rootDir).Error
	if err != nil {
		return errors.WithStack(err)
	}

	log.Printf("-- Created preparation %s\n", scanning.Name)
	cursor, err := mg.Database("singularity").Collection("generationrequests").Find(
		ctx, bson.M{"datasetName": scanning.Name},
	)
	if err != nil {
		return errors.Wrap(err, "failed to query mongo for generation requests")
	}

	directoryCache := map[string]uint64{}
	directoryCache[""] = rootDir.ID
	directoryCache["."] = rootDir.ID
	var lastFile model.File
	for cursor.Next(ctx) {
		var generation GenerationRequest
		err = cursor.Decode(&generation)
		if err != nil {
			return errors.Wrap(err, "failed to decode generation request")
		}

		packJob := model.Job{
			Type:            model.Pack,
			State:           model.Complete,
			ErrorMessage:    generation.ErrorMessage,
			ErrorStackTrace: "",
			AttachmentID:    attachment.ID,
		}
		err = db.Create(&packJob).Error
		if err != nil {
			return errors.WithStack(err)
		}
		log.Printf("-- Created pack job %d for %s\n", packJob.ID, scanning.Name)

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
			CreatedAt:     generation.CreatedAt,
			PieceCID:      model.CID(pieceCID),
			PieceSize:     int64(generation.PieceSize),
			RootCID:       model.CID(dataCID),
			FileSize:      int64(generation.CarSize),
			StoragePath:   filepath.Join(scanning.OutDir, fileName),
			AttachmentID:  &attachment.ID,
			PreparationID: preparation.ID,
		}
		err = db.Create(&car).Error
		if err != nil {
			return errors.Wrap(err, "failed to create car")
		}
		log.Printf("-- Created car %s for %s\n", generation.PieceCID, scanning.Name)

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
						Path: generatedFile.Path,
						Size: int64(generatedFile.Size),
						CID:  model.CID(fileCID),
						FileRanges: []model.FileRange{
							{
								Offset: 0,
								Length: int64(generatedFile.Size),
								CID:    model.CID(fileCID),
								JobID:  &packJob.ID,
							},
						},
						AttachmentID:     attachment.ID,
						LastModifiedNano: generation.CreatedAt.UnixNano(),
					}
				} else if generatedFile.Start == 0 {
					lastFile = model.File{
						Path: generatedFile.Path,
						Size: int64(generatedFile.Size),
						CID:  model.CID(cid.Undef),
						FileRanges: []model.FileRange{
							{
								Offset: 0,
								Length: int64(generatedFile.End),
								CID:    model.CID(fileCID),
								JobID:  &packJob.ID,
							},
						},
						AttachmentID:     attachment.ID,
						LastModifiedNano: generation.CreatedAt.UnixNano(),
					}
					continue
				} else {
					lastFile.FileRanges = append(lastFile.FileRanges, model.FileRange{
						Offset: int64(generatedFile.Start),
						Length: int64(generatedFile.End - generatedFile.Start),
						CID:    model.CID(fileCID),
						JobID:  &packJob.ID,
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
						_, root, err := util2.AssembleFileFromLinks(links)
						if err != nil {
							return errors.Wrap(err, "failed to assemble file from links")
						}
						file.CID = model.CID(root.Cid())
					}
				}
				err = push.EnsureParentDirectories(ctx, db, &file, rootDir.ID, directoryCache)
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
		log.Printf("-- Created %d files for %s\n", len(files), scanning.Name)
	}

	return nil
}

func MigrateDataset(cctx *cli.Context) error {
	skipFiles := cctx.Bool("skip-files")
	log.Println("Migrating dataset from old singularity database")
	mongoConnectionString := cctx.String("mongo-connection-string")
	sqlConnectionString := cctx.String("database-connection-string")
	log.Printf("Using mongo connection string: %s\n", mongoConnectionString)
	log.Printf("Using sql connection string: %s\n", sqlConnectionString)
	db, closer, err := database.OpenFromCLI(cctx)
	if err != nil {
		return errors.WithStack(err)
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
		err = db.Model(&model.Storage{}).Where("name = ?", scanning.Name+"-source").Count(&datasetExists).Error
		if err != nil {
			return errors.Wrapf(err, "failed to query for dataset %s", scanning.Name)
		}
		if datasetExists > 0 {
			log.Printf("Preparation %s already exists, skipping\n", scanning.Name)
			continue
		}
		log.Printf("Migrating Preparation: %s\n", scanning.Name)
		err = migrateDataset(ctx, mg, db, scanning, skipFiles)
		if err != nil {
			return errors.Wrapf(err, "failed to migrate dataset %s", scanning.Name)
		}
	}
	return nil
}
