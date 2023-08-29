package client_test

import (
	"context"
	"crypto/rand"
	"os"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/client"
	"github.com/data-preservation-programs/singularity/client/testutil"
	"github.com/data-preservation-programs/singularity/model"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/data-preservation-programs/singularity/handler/deal/schedule"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/stretchr/testify/require"
)

func TestClients(t *testing.T) {
	ctx := context.Background()
	testutil.TestWithAllClients(ctx, t, func(t *testing.T, client client.Client) {
		// createDataset
		ds, err := client.CreateDataset(ctx, dataset.CreateRequest{
			Name:       "test",
			MaxSizeStr: "31.5GiB",
		})
		require.NoError(t, err)
		require.Equal(t, "test", ds.Name)

		// cannot create dataset with same name

		dupDataset, err := client.CreateDataset(ctx, dataset.CreateRequest{
			Name:       "test",
			MaxSizeStr: "31.5GiB",
		})
		var asConflict handler.DuplicateRecordError
		require.ErrorAs(t, err, &asConflict)
		require.Nil(t, dupDataset)

		// cannot create dataset with invalid parameter
		invalidDataset, err := client.CreateDataset(ctx, dataset.CreateRequest{})
		var asInvalidParameter handler.InvalidParameterError
		require.ErrorAs(t, err, &asInvalidParameter)
		require.Nil(t, invalidDataset)

		path := t.TempDir()
		// create datasource
		source, err := client.CreateLocalSource(ctx, "test", datasource.LocalRequest{
			SourcePath:     path,
			RescanInterval: "0h",
			ScanningState:  model.Ready,
		})
		require.NoError(t, err)
		require.Equal(t, "local", source.Type)
		require.Equal(t, ds.ID, source.DatasetID)
		require.Equal(t, path, source.Path)
		require.Equal(t, model.Ready, source.ScanningState)

		// list sources for dataset
		sources, err := client.ListSourcesByDataset(ctx, "test")
		require.NoError(t, err)
		require.Len(t, sources, 1)
		require.Equal(t, ds.ID, sources[0].DatasetID)
		require.Equal(t, path, sources[0].Path)
		require.Equal(t, model.Ready, sources[0].ScanningState)

		// create datasource when dataset not found
		notFoundSource, err := client.CreateLocalSource(ctx, "apples", datasource.LocalRequest{
			SourcePath:     path,
			RescanInterval: "0h",
		})
		var asNotFoundError handler.NotFoundError
		require.ErrorAs(t, err, &asNotFoundError)
		require.Nil(t, notFoundSource)

		// push osFile
		osFile, err := os.CreateTemp(path, "push-*")
		require.NoError(t, err)
		buf := make([]byte, 1000)
		_, _ = rand.Read(buf)
		osFile.Write(buf)
		name := osFile.Name()
		err = osFile.Close()
		require.NoError(t, err)
		fileA, err := client.PushFile(ctx, source.ID, datasource.FileInfo{Path: filepath.Base(name)})
		require.NoError(t, err)
		require.Equal(t, filepath.Base(name), fileA.Path)

		// get file
		returnedFile, err := client.GetFile(ctx, fileA.ID)
		require.NoError(t, err)
		require.Equal(t, fileA.Path, returnedFile.Path)

		// push another file
		osFile, err = os.CreateTemp(path, "push-*")
		require.NoError(t, err)
		buf = make([]byte, 1000)
		_, _ = rand.Read(buf)
		osFile.Write(buf)
		name = osFile.Name()
		err = osFile.Close()
		require.NoError(t, err)
		fileB, err := client.PushFile(ctx, source.ID, datasource.FileInfo{Path: filepath.Base(name)})
		require.NoError(t, err)
		require.Equal(t, filepath.Base(name), fileB.Path)

		// Prepare the source (create pack jobs)
		client.PrepareToPackSource(ctx, source.ID)

		// Check that pack job exists
		packJobs, err := client.GetSourcePackJobs(ctx, source.ID, inspect.GetSourcePackJobsRequest{})
		require.NoError(t, err)

		require.Len(t, packJobs, 1)

		// import wallet
		key := "7b2254797065223a22736563703235366b31222c22507269766174654b6579223a226b35507976337148327349586343595a58594f5775453149326e32554539436861556b6c4e36695a5763453d227d"
		wallet, err := client.ImportWallet(ctx, wallet.ImportRequest{
			PrivateKey: key,
		})
		require.NoError(t, err)

		// add to dataset
		_, err = client.AddWalletToDataset(ctx, "test", wallet.ID)
		require.NoError(t, err)

		// verify wallet is present in dataset
		wallets, err := client.ListWalletsByDataset(ctx, "test")
		require.NoError(t, err)
		require.Len(t, wallets, 1)
		require.Equal(t, *wallet, wallets[0])
		// create schedule
		_, err = client.CreateSchedule(ctx, schedule.CreateRequest{
			DatasetName:        "test",
			StartDelay:         "72h",
			Duration:           "2400h",
			TotalDealSize:      "100TiB",
			ScheduleDealSize:   "100TiB",
			MaxPendingDealSize: "100TiB",
			Provider:           "f01",
		})
		require.NoError(t, err)

		// list schedules
		schedules, err := client.ListSchedules(ctx)
		require.NoError(t, err)
		require.Len(t, schedules, 1)
	})
}
