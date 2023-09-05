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
	require.Equal(t, http2.StatusOK, resp.StatusCode)
	require.Contains(t, body, "robotstxt.org")
	return func() {
		select {
		case <-done:
		case <-ctx.Done():
		}
	}
}

func TestBasicDataPrep(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		source := t.TempDir()
		err := os.WriteFile(filepath.Join(source, "test.txt"), []byte("hello world"), 0644)
		require.NoError(t, err)
		output := t.TempDir()
		ctx, cancel := context.WithCancel(ctx)
		defer runAPI(t, ctx)()
		defer cancel()
		client := http.NewHTTPClientWithConfig(nil, &http.TransportConfig{
			Host:     apiBind,
			BasePath: http.DefaultBasePath,
		})
		// Create source storage
		response, err := client.Storage.PostStorageLocal(&storage.PostStorageLocalParams{
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
		response, err = client.Storage.PostStorageLocal(&storage.PostStorageLocalParams{
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
		createPrepResp, err := client.Preparation.PostPreparation(&preparation.PostPreparationParams{
			Request: &models.DataprepCreateRequest{
				MaxSize:        ptr.Of("3MB"),
				Name:           ptr.Of("prep"),
				OutputStorages: []string{"output"},
				SourceStorages: []string{"source"},
			},
			Context: ctx,
		})
		require.NoError(t, err)
		require.True(t, createPrepResp.IsSuccess())
		require.NotZero(t, createPrepResp.Payload.ID)
		// Start Scanning
		startScanResp, err := client.Job.PostPreparationIDSourceNameStartScan(&job.PostPreparationIDSourceNameStartScanParams{
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
		listPiecesResp, err := client.Piece.GetPreparationIDPiece(&piece.GetPreparationIDPieceParams{
			ID:      "prep",
			Context: ctx,
		})
		require.NoError(t, err)
		require.True(t, listPiecesResp.IsSuccess())
		require.Len(t, listPiecesResp.Payload, 1)
		require.Len(t, listPiecesResp.Payload[0].Pieces, 1)
		require.Equal(t, "baga6ea4seaqoahdvfwkrp64ecsxbjvyuqcwpz3o7ctxrjanlv2x4u2cq2qjf2ji", listPiecesResp.Payload[0].Pieces[0].PieceCid.(string))
	})
}
