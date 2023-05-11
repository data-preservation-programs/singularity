package migrate

import (
	"context"
	"path/filepath"
	"strings"
	"time"

	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/data-preservation-programs/go-singularity/replication"
	"github.com/data-preservation-programs/go-singularity/util"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/exp/slices"
)

func Migrate(cctx *cli.Context) error {
	logger := log.Logger("cli")
	db := database.MustOpenFromCLI(cctx)
	ctx := context.TODO()
	mg, err := mongo.Connect(ctx, options.Client().ApplyURI(cctx.String("mongo-connection-string")))
	if err != nil {
		return cli.Exit("Failed to connect to mongo: "+err.Error(), 1)
	}

	cursor, err := mg.Database("singularity").Collection("dealstates").Find(ctx, bson.D{})
	if err != nil {
		return cli.Exit("Failed to query mongo: "+err.Error(), 1)
	}

	var deals []model.Deal
	for cursor.Next(ctx) {
		var deal DealState
		err = cursor.Decode(&deal)
		if err != nil {
			return cli.Exit("Failed to decode mongo document: "+err.Error(), 1)
		}
		d := model.Deal{
			CreatedAt:     time.Time{},
			UpdatedAt:     time.Time{},
			DealID:        &deal.DealID,
			State:         model.DealState(deal.State),
			ClientAddress: deal.Client,
			Provider:      deal.Provider,
			ProposalID:    deal.DealCID,
			PieceCID:      deal.PieceCID,
			PieceSize:     2 << 35,
			Start:         replication.EpochToTime(abi.ChainEpoch(deal.StartEpoch)),
			Duration:      time.Duration(deal.Duration) * 30 * time.Second,
			End:           replication.EpochToTime(abi.ChainEpoch(deal.StartEpoch + deal.Duration)),
			SectorStart:   replication.EpochToTime(abi.ChainEpoch(deal.StartEpoch)).Add(time.Hour * 24),
			Price:         0,
			Verified:      true,
		}
		deals = append(deals, d)
	}
	err = db.CreateInBatches(deals, 1000).Error
	if err != nil {
		return cli.Exit("Failed to insert deals: "+err.Error(), 1)
	}

	cursor, err = mg.Database("singularity").Collection("scanningrequests").Find(ctx, bson.D{})
	if err != nil {
		return cli.Exit("Failed to query mongo: "+err.Error(), 1)
	}

	for cursor.Next(ctx) {
		directoryCache := map[string]model.Directory{}
		var request ScanningRequest
		err = cursor.Decode(&request)
		if slices.Contains([]string{"allen-direct", "ford-mas", "worldbank", "worldbank2"}, request.Name) {
			continue
		}

		if err != nil {
			return cli.Exit("Failed to decode mongo document: "+err.Error(), 1)
		}

		dataset := model.Dataset{
			Name:                 request.Name,
			CreatedAt:            request.UpdatedAt,
			UpdatedAt:            request.UpdatedAt,
			MinSize:              request.MinSize,
			MaxSize:              request.MaxSize,
			PieceSize:            util.NextPowerOfTwo(request.MaxSize),
			OutputDirs:           []string{request.OutDir},
			EncryptionRecipients: nil,
			EncryptionScript:     "",
		}
		logger.Info("Creating dataset: ", request.Name)
		err = db.Where("name = ?", request.Name).Attrs(dataset).FirstOrCreate(&dataset).Error
		if err != nil {
			return cli.Exit("Failed to create dataset: "+err.Error(), 1)
		}

		rootDirectory := model.Directory{
			Name: "",
		}
		logger.Info("Creating root directory: ", request.Name)
		err = db.Create(&rootDirectory).Error
		if err != nil {
			return cli.Exit("Failed to create root directory: "+err.Error(), 1)
		}

		sourceType := model.Dir
		fileType := model.File
		if strings.HasPrefix(request.Path, "s3") {
			sourceType = model.S3Path
			fileType = model.S3Object
		}
		source := model.Source{
			DatasetID:            dataset.ID,
			CreatedAt:            request.UpdatedAt,
			UpdatedAt:            request.UpdatedAt,
			Type:                 sourceType,
			Path:                 request.Path,
			ScanningState:        model.Complete,
			LastScannedTimestamp: request.UpdatedAt.Unix(),
			RootDirectory:        &rootDirectory,
		}
		logger.Info("Creating source: ", request.Name)
		err = db.Where("dataset_id = ?", dataset.ID).Attrs(source).FirstOrCreate(&source).Error
		if err != nil {
			return cli.Exit("Failed to create source: "+err.Error(), 1)
		}
		directoryCache[""] = *source.RootDirectory
		directoryCache["."] = *source.RootDirectory

		cursor, err := mg.Database("singularity").Collection("generationrequests").Find(
			ctx, bson.D{
				{Key: "datasetName", Value: request.Name},
			},
		)
		if err != nil {
			return cli.Exit("Failed to query mongo: "+err.Error(), 1)
		}

		for cursor.Next(ctx) {
			var generation GenerationRequest
			err = cursor.Decode(&generation)
			if err != nil {
				return cli.Exit("Failed to decode mongo document: "+err.Error(), 1)
			}

			chunk := model.Chunk{
				CreatedAt:    generation.CreatedAt,
				SourceID:     source.ID,
				PackingState: model.Complete,
			}

			logger.Info("Creating chunk: ", generation.PieceCID)
			err = db.Create(&chunk).Error
			if err != nil {
				return cli.Exit("Failed to create chunk: "+err.Error(), 1)
			}

			car := model.Car{
				CreatedAt: generation.CreatedAt,
				PieceCID:  generation.PieceCID,
				PieceSize: generation.PieceSize,
				RootCID:   generation.DataCID,
				FileSize:  generation.CarSize,
				FilePath:  generation.OutDir + "/" + generation.PieceCID + ".car",
				DatasetID: dataset.ID,
				ChunkID:   &chunk.ID,
			}
			logger.Info("Creating car: ", generation.PieceCID)
			err = db.Where("dataset_id = ? AND piece_cid = ?", dataset.ID, generation.PieceCID).
				Attrs(car).
				FirstOrCreate(&car).Error
			if err != nil {
				return cli.Exit("Failed to create car: "+err.Error(), 1)
			}

			cursor, err := mg.Database("singularity").Collection("outputfilelists").
				Find(
					ctx, bson.D{{Key: "generationId", Value: generation.ID}},
					options.Find().SetSort(bson.D{{Key: "index", Value: 1}}),
				)
			if err != nil {
				return cli.Exit("Failed to query mongo: "+err.Error(), 1)
			}

			var items []model.Item
			for cursor.Next(ctx) {
				var fileList OutputFileList
				err = cursor.Decode(&fileList)
				if err != nil {
					return cli.Exit("Failed to decode mongo document: "+err.Error(), 1)
				}

				logger.Info("Got list: ", fileList.ID)
				for _, file := range fileList.GeneratedFileList {
					if file.CID == "unrecoverable" {
						continue
					}
					if file.Dir {
						if file.Path == "" {
							continue
						}
						if _, ok := directoryCache[file.Path]; !ok {
							parentPath := filepath.Dir(file.Path)
							parent, ok := directoryCache[parentPath]
							if !ok {
								return cli.Exit("Failed to find parent directory: "+parentPath, 1)
							}

							directory := model.Directory{
								Name:     filepath.Base(file.Path),
								ParentID: &parent.ID,
							}
							logger.Info("Creating directory: ", directory.Name)
							err = db.Create(&directory).Error
							if err != nil {
								return cli.Exit("Failed to create directory: "+err.Error(), 1)
							}
							directoryCache[file.Path] = directory
						}
						continue
					}

					directory, ok := directoryCache[filepath.Dir(file.Path)]
					if !ok {
						return cli.Exit("Failed to find directory of : "+filepath.Dir(file.Path), 1)
					}
					item := model.Item{
						ScannedAt:    generation.CreatedAt,
						ChunkID:      &chunk.ID,
						Type:         fileType,
						Path:         file.Path,
						Size:         file.Size,
						Offset:       file.Start,
						Length:       file.End - file.Start,
						LastModified: nil,
						CID:          file.CID,
						DirectoryID:  &directory.ID,
					}
					items = append(items, item)
				}
			}
			if len(items) == 0 {
				logger.Warn("No items found for chunk: ", chunk.ID)
			} else {
				logger.Info("Creating items: ", len(items))
				err = db.CreateInBatches(&items, 1000).Error
				if err != nil {
					return cli.Exit("Failed to create items: "+err.Error(), 1)
				}
			}
		}

		cursor, err = mg.Database("singularity").Collection("replicationrequests").Find(
			ctx, bson.D{
				{Key: "datasetId", Value: request.ID},
			},
		)
		if err != nil {
			return cli.Exit("Failed to query mongo: "+err.Error(), 1)
		}

		for cursor.Next(ctx) {
			var replication ReplicationRequest
			err = cursor.Decode(&replication)
			if err != nil {
				return cli.Exit("Failed to decode mongo document: "+err.Error(), 1)
			}

			max := replication.CronMaxDeals
			if max < replication.MaxNumberOfDeals {
				max = replication.MaxNumberOfDeals
			}
			schedule := model.Schedule{
				CreatedAt:            replication.CreatedAt,
				UpdatedAt:            replication.UpdatedAt,
				DatasetID:            dataset.ID,
				URLTemplate:          replication.URLPrefix + "/{PIECE_CID}.car",
				Provider:             replication.StorageProviders,
				TotalDealNumber:      max,
				Verified:             replication.IsVerified,
				KeepUnsealed:         true,
				AnnounceToIPNI:       true,
				StartDelay:           time.Second * time.Duration(replication.StartDelay) * 30,
				Duration:             time.Second * time.Duration(replication.Duration) * 30,
				State:                model.ScheduleCompleted,
				SchedulePattern:      replication.CronSchedule,
				ScheduleDealNumber:   replication.MaxNumberOfDeals,
				MaxPendingDealNumber: replication.CronMaxPendingDeals,
				Notes:                replication.Notes,
				PieceCIDListPath:     replication.FileListPath,
			}
			logger.Info("Creating schedule: ", replication.ID)
			err = db.Create(&schedule).Error
			if err != nil {
				return cli.Exit("Failed to create schedule: "+err.Error(), 1)
			}
		}
	}
	return nil
}
