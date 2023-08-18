package migrate

import (
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var pieceCidRegex = regexp.MustCompile("baga[0-9a-z]+")

func MigrateSchedule(cctx *cli.Context) error {
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
	mg, err := mongo.Connect(ctx, options.Client().ApplyURI(cctx.String("mongo-connection-string")))
	if err != nil {
		return errors.Wrap(err, "failed to connect to mongo")
	}

	var count int64
	err = db.Model(&model.Schedule{}).Count(&count).Error
	if err != nil {
		return errors.Wrap(err, "failed to count schedules")
	}

	if count > 0 {
		log.Println("Schedules already exist, skipping")
		return nil
	}

	resp, err := mg.Database("singularity").Collection("replicationrequests").Find(ctx, bson.M{})
	if err != nil {
		return errors.Wrap(err, "failed to query mongo for scanning requests")
	}

	var replications []ReplicationRequest
	err = resp.All(ctx, &replications)
	if err != nil {
		return errors.Wrap(err, "failed to decode mongo response")
	}

	var schedules []model.Schedule
	for _, replication := range replications {
		var scanning ScanningRequest
		oid, err := primitive.ObjectIDFromHex(replication.DatasetID)
		if err != nil {
			return errors.Wrapf(err, "failed to parse dataset id %s", replication.DatasetID)
		}
		findResult := mg.Database("singularity").Collection("scanningrequests").FindOne(ctx, bson.M{"_id": oid})
		if findResult.Err() != nil {
			return errors.Wrapf(err, "failed to find dataset %s", replication.DatasetID)
		}

		err = findResult.Decode(&scanning)
		if err != nil {
			return errors.Wrapf(err, "failed to decode dataset %s", replication.DatasetID)
		}

		var dataset model.Preparation
		err = db.Where("name = ?", scanning.Name).First(&dataset).Error
		if err != nil {
			return errors.Wrapf(err, "failed to find dataset %s", scanning.Name)
		}

		var urlTemplate string
		if replication.URLPrefix != "" {
			if !strings.HasSuffix(replication.URLPrefix, "/") {
				replication.URLPrefix += "/"
			}
			urlTemplate = replication.URLPrefix + "{PIECE_CID}"
		}
		totalDealNumber := replication.MaxNumberOfDeals
		var scheduleDealNumber int
		var maxPendingDealNumber int
		if replication.CronSchedule != "" {
			totalDealNumber = replication.CronMaxDeals
			scheduleDealNumber = int(replication.MaxNumberOfDeals)
			maxPendingDealNumber = int(replication.CronMaxPendingDeals)
		}
		var allowedCIDs model.StringSlice
		if replication.FileListPath != "" {
			content, err := os.ReadFile(replication.FileListPath)
			if err != nil {
				log.Printf("failed to read file list %s. Skipping...", replication.FileListPath)
			} else {
				allowedCIDs = pieceCidRegex.FindAllString(string(content), -1)
			}
		}

		for _, provider := range strings.Split(replication.StorageProviders, ",") {
			if provider == "" {
				continue
			}
			schedule := model.Schedule{
				CreatedAt:            replication.CreatedAt,
				UpdatedAt:            replication.UpdatedAt,
				DatasetID:            dataset.ID,
				URLTemplate:          urlTemplate,
				Provider:             provider,
				PricePerGBEpoch:      replication.MaxPrice,
				TotalDealNumber:      int(totalDealNumber),
				TotalDealSize:        0,
				Verified:             replication.IsVerified,
				KeepUnsealed:         true,
				AnnounceToIPNI:       true,
				StartDelay:           time.Second * time.Duration(replication.StartDelay) * 30,
				Duration:             time.Second * time.Duration(replication.Duration) * 30,
				State:                model.SchedulePaused,
				ScheduleCron:         replication.CronSchedule,
				ScheduleDealNumber:   scheduleDealNumber,
				ScheduleDealSize:     0,
				MaxPendingDealNumber: maxPendingDealNumber,
				MaxPendingDealSize:   0,
				Notes:                replication.Notes,
				ErrorMessage:         replication.ErrorMessage,
				AllowedPieceCIDs:     allowedCIDs,
			}
			schedules = append(schedules, schedule)
		}
	}

	err = db.CreateInBatches(&schedules, util.BatchSize).Error
	if err != nil {
		return errors.Wrap(err, "failed to create schedules")
	}

	cliutil.PrintToConsole(schedules, false, nil)
	return nil
}
