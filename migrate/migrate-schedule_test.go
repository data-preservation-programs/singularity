package migrate

import (
	"context"
	"flag"
	"os"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

func TestMigrateSchedule_DatasetNotExist(t *testing.T) {
	err := setupMongoDBSchedule()
	if err != nil {
		t.Log(err)
		t.Skip("Skipping test because MongoDB is not available")
	}
	defer os.Remove("1.txt") // Clean up the test file

	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		flagSet := flag.NewFlagSet("", 0)
		flagSet.String("mongo-connection-string", localMongoDB, "")
		flagSet.String("database-connection-string", os.Getenv("DATABASE_CONNECTION_STRING"), "")
		cctx := cli.NewContext(&cli.App{
			Writer: os.Stdout,
		}, flagSet, nil)

		err = db.Create(&model.Preparation{
			Name: "test",
			SourceStorages: []model.Storage{{
				Name: "test-source",
			}},
			OutputStorages: []model.Storage{{
				Name: "test-output",
			}},
		}).Error
		require.NoError(t, err)

		err = MigrateSchedule(cctx)
		require.NoError(t, err)

		// Migrate again does nothing
		err = MigrateSchedule(cctx)
		require.NoError(t, err)

		var schedules []model.Schedule
		err = db.Find(&schedules).Error
		require.NoError(t, err)
		require.Len(t, schedules, 2)
		require.EqualValues(t, 1, schedules[0].PreparationID)
		require.Equal(t, "http://localhost:8080/{PIECE_CID}", schedules[0].URLTemplate)
		require.Equal(t, "f0miner1", schedules[0].Provider)
		require.Equal(t, 100, schedules[0].TotalDealNumber)
		require.True(t, schedules[0].Verified)
		require.True(t, schedules[0].KeepUnsealed)
		require.True(t, schedules[0].AnnounceToIPNI)
		require.Equal(t, time.Hour*24, schedules[0].StartDelay)
		require.Equal(t, time.Minute/2*150000, schedules[0].Duration)
		require.Equal(t, model.SchedulePaused, schedules[0].State)
		require.Equal(t, 10, schedules[0].ScheduleDealNumber)
		require.Equal(t, 10, schedules[0].MaxPendingDealNumber)
		require.Equal(t, "notes", schedules[0].Notes)
		require.Equal(t, "error message", schedules[0].ErrorMessage)
	})
}

func setupMongoDBSchedule() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(localMongoDB))
	if err != nil {
		return errors.WithStack(err)
	}
	defer db.Disconnect(context.Background())
	err = db.Database("singularity").Drop(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	// Create the file list file that the replication request references
	err = os.WriteFile("1.txt", []byte("baga6ea4seaqexample1234567890abcdef\nbaga6ea4seaqexample0987654321fedcba\n"), 0644)
	if err != nil {
		return errors.WithStack(err)
	}

	insertedDatasetResult, err := db.Database("singularity").Collection("scanningrequests").InsertOne(ctx, ScanningRequest{
		Name: "test",
	})
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = db.Database("singularity").Collection("replicationrequests").InsertMany(ctx, []any{ReplicationRequest{
		DatasetID:           insertedDatasetResult.InsertedID.(primitive.ObjectID).Hex(),
		MaxReplicas:         10,
		StorageProviders:    "f0miner1,f0miner2",
		Client:              "f0client",
		URLPrefix:           "http://localhost:8080",
		MaxPrice:            0,
		MaxNumberOfDeals:    10,
		IsVerified:          true,
		StartDelay:          2880,
		Duration:            150000,
		IsOffline:           false,
		Status:              ReplicationStatusActive,
		CronSchedule:        "* * * * *",
		CronMaxDeals:        100,
		CronMaxPendingDeals: 10,
		FileListPath:        "1.txt",
		Notes:               "notes",
		ErrorMessage:        "error message",
	}})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
