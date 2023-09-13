package cmd

import (
	"context"
	"fmt"
	http2 "net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/client/swagger/http"
	"github.com/data-preservation-programs/singularity/client/swagger/http/file"
	"github.com/data-preservation-programs/singularity/client/swagger/http/job"
	"github.com/data-preservation-programs/singularity/client/swagger/http/piece"
	"github.com/data-preservation-programs/singularity/client/swagger/http/preparation"
	"github.com/data-preservation-programs/singularity/client/swagger/http/storage"
	"github.com/data-preservation-programs/singularity/client/swagger/models"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/gotidy/ptr"
	"github.com/parnurzeal/gorequest"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// The bind address needs to be different for different test package so that they don't conflict.
const apiBind = "127.0.0.1:9091"

func runAPI(t *testing.T, ctx context.Context) func() {
	t.Helper()
	done := make(chan struct{})
	go func() {
		NewRunner().Run(ctx, fmt.Sprintf("singularity run api --bind %s", apiBind))
		close(done)
	}()

	var resp *http2.Response
	var body string
	// try every 100ms for up to 5 seconds for server to come up
	for i := 0; i < 50; i++ {
		time.Sleep(100 * time.Millisecond)
		resp, body, _ = gorequest.New().
			Get(fmt.Sprintf("http://%s/robots.txt", apiBind)).End()
		if resp != nil && resp.StatusCode == http2.StatusOK {
			break
		}
	}
	require.NotNil(t, resp)
	require.Equal(t, http2.StatusOK, resp.StatusCode)
	require.Contains(t, body, "robotstxt.org")
	return func() {
		select {
		case <-done:
		case <-ctx.Done():
		}
	}
}

// TestMotionIntegration tests the basic data preparation flow with Motion.
// 1. Create a local source storage
// 2. Create a local output storage
// 3. Create a preparation with the source and output storage
// 4. Push a file to the source storage
// 5. Prepare this file for packing
// 6. Check this file status
// 7. Get all jobs
// 8. Pack each job
func TestMotionIntegration(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ctx, cancel := context.WithCancel(ctx)
		client, done := setupPreparation(t, ctx, true)
		defer done()
		defer cancel()
		// Push a file
		pushResp, err := client.File.PushFile(&file.PushFileParams{
			File: &models.FileInfo{
				Path: "test.txt",
			},
			ID:      "prep",
			Name:    "source",
			Context: ctx,
		})
		require.NoError(t, err)
		require.True(t, pushResp.IsSuccess())
		require.NotZero(t, pushResp.Payload.ID)
		// Prepare a file
		prepareResp, err := client.File.PrepareToPackFile(&file.PrepareToPackFileParams{
			ID:      pushResp.Payload.ID,
			Context: ctx,
		})
		require.NoError(t, err)
		require.True(t, prepareResp.IsSuccess())
		require.NotZero(t, prepareResp.Payload)
		// Get this file
		getFileResp, err := client.File.GetFile(&file.GetFileParams{
			ID:      pushResp.Payload.ID,
			Context: ctx,
		})
		require.NoError(t, err)
		require.True(t, getFileResp.IsSuccess())
		require.NotNil(t, getFileResp.Payload)
		require.Len(t, getFileResp.Payload.FileRanges, 1)
		// Get all jobs
		jobs, err := client.Preparation.GetPreparationStatus(&preparation.GetPreparationStatusParams{
			ID:      "prep",
			Context: ctx,
		})
		require.NoError(t, err)
		require.True(t, jobs.IsSuccess())
		require.Len(t, jobs.Payload, 1)
		require.Len(t, jobs.Payload[0].Jobs, 1)
		// Pack that job
		car, err := client.Job.Pack(&job.PackParams{
			ID:      jobs.Payload[0].Jobs[0].ID,
			Context: ctx,
		})
		require.NoError(t, err)
		require.True(t, car.IsSuccess())
		require.NotZero(t, car.Payload.ID)
		require.Equal(t, "baga6ea4seaqoahdvfwkrp64ecsxbjvyuqcwpz3o7ctxrjanlv2x4u2cq2qjf2ji", car.Payload.PieceCid)
	})
}

func setupPreparation(t *testing.T, ctx context.Context, disableDagInline bool) (*http.SingularityAPI, func()) {
	t.Helper()
	source := t.TempDir()
	err := os.WriteFile(filepath.Join(source, "test.txt"), []byte("hello world"), 0644)
	require.NoError(t, err)
	output := t.TempDir()
	done := runAPI(t, ctx)
	client := http.NewHTTPClientWithConfig(nil, &http.TransportConfig{
		Host:     apiBind,
		BasePath: http.DefaultBasePath,
	})
	// Create source storage
	response, err := client.Storage.CreateLocalStorage(&storage.CreateLocalStorageParams{
		Request: &models.StorageCreateLocalStorageRequest{
			Name: "source",
			Path: source,
		},
		Context: ctx,
	})
	require.NoError(t, err)
	require.True(t, response.IsSuccess())
	require.NotZero(t, response.Payload.ID)
	// Create output storage
	response, err = client.Storage.CreateLocalStorage(&storage.CreateLocalStorageParams{
		Request: &models.StorageCreateLocalStorageRequest{
			Name: "output",
			Path: output,
		},
		Context: ctx,
	})
	require.NoError(t, err)
	require.True(t, response.IsSuccess())
	require.NotZero(t, response.Payload.ID)
	// Create preparation
	createRequest := &preparation.CreatePreparationParams{
		Request: &models.DataprepCreateRequest{
			MaxSize:        ptr.Of("3MB"),
			Name:           "prep",
			OutputStorages: []string{"output"},
			SourceStorages: []string{"source"},
		},
		Context: ctx,
	}
	if disableDagInline {
		createRequest.Request.NoDag = ptr.Bool(true)
		createRequest.Request.NoInline = ptr.Bool(true)
	}
	createPrepResp, err := client.Preparation.CreatePreparation(createRequest)
	require.NoError(t, err)
	require.True(t, createPrepResp.IsSuccess())
	require.NotZero(t, createPrepResp.Payload.ID)
	return client, done
}

// TestBasicDataPrep tests the basic data preparation flow.
// 1. Create a local source storage
// 2. Create a local output storage
// 3. Create a preparation with the source and output storage
// 4. Start scanning the source storage
// 5. Run the dataset worker
// 6. List the pieces
// 7. Start daggen
// 8. Run the dataset worker
// 9. List the pieces
func TestBasicDataPrep(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ctx, cancel := context.WithCancel(ctx)
		client, done := setupPreparation(t, ctx, false)
		defer done()
		defer cancel()
		// Start Scanning
		startScanResp, err := client.Job.StartScan(&job.StartScanParams{
			ID:      "prep",
			Name:    "source",
			Context: ctx,
		})
		require.NoError(t, err)
		require.True(t, startScanResp.IsSuccess())
		require.NotZero(t, startScanResp.Payload.ID)
		// Run dataset worker
		_, _, err = NewRunner().Run(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		// List pieces
		listPiecesResp, err := client.Piece.ListPieces(&piece.ListPiecesParams{
			ID:      "prep",
			Context: ctx,
		})
		require.NoError(t, err)
		require.True(t, listPiecesResp.IsSuccess())
		require.Len(t, listPiecesResp.Payload, 1)
		require.Len(t, listPiecesResp.Payload[0].Pieces, 1)
		require.Equal(t, "baga6ea4seaqoahdvfwkrp64ecsxbjvyuqcwpz3o7ctxrjanlv2x4u2cq2qjf2ji", listPiecesResp.Payload[0].Pieces[0].PieceCid)
		// Start daggen
		startDagGenResp, err := client.Job.StartDagGen(&job.StartDagGenParams{
			ID:      "prep",
			Name:    "source",
			Context: ctx,
		})
		require.NoError(t, err)
		require.True(t, startDagGenResp.IsSuccess())
		require.NotZero(t, startDagGenResp.Payload.ID)
		// Run dataset worker
		_, _, err = NewRunner().Run(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		// List pieces
		listPiecesResp, err = client.Piece.ListPieces(&piece.ListPiecesParams{
			ID:      "prep",
			Context: ctx,
		})
		require.NoError(t, err)
		require.True(t, listPiecesResp.IsSuccess())
		require.Len(t, listPiecesResp.Payload, 1)
		require.Len(t, listPiecesResp.Payload[0].Pieces, 2)
		require.Equal(t, "baga6ea4seaqoahdvfwkrp64ecsxbjvyuqcwpz3o7ctxrjanlv2x4u2cq2qjf2ji", listPiecesResp.Payload[0].Pieces[0].PieceCid)
		require.Equal(t, "baga6ea4seaqikgg6djdjtgro2pwsipxkjo754gwykmyglrv5yhjkgqi23dtqubq", listPiecesResp.Payload[0].Pieces[1].PieceCid)
	})
}
