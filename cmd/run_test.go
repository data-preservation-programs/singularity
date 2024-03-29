package cmd

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/parnurzeal/gorequest"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestRunDealTracker(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		_, _, err := NewRunner().Run(ctx, "singularity run deal-tracker")
		require.ErrorIs(t, err, context.DeadlineExceeded)
	})
}

func TestRunAPI(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		done := make(chan struct{})
		go func() {
			_, _, err := NewRunner().Run(ctx, "singularity run api")
			require.ErrorIs(t, err, context.Canceled)
			close(done)
		}()
		var resp *http.Response
		var errs []error
		// try every 100ms for up to 5 seconds for server to come up
		for i := 0; i < 50; i++ {
			time.Sleep(100 * time.Millisecond)
			resp, _, errs = gorequest.New().
				Get("http://127.0.0.1:9090/health").End()
			if resp != nil && resp.StatusCode == http.StatusOK {
				break
			}
		}
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		cancel()
		<-done
	})
}

func TestRunDatasetWorker(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		_, _, err := NewRunner().Run(ctx, "singularity run dataset-worker")
		require.ErrorIs(t, err, context.DeadlineExceeded)
	})
}

func TestRunContentProvider(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		_, _, err := NewRunner().Run(ctx, "singularity run content-provider --http-bind "+contentProviderBind)
		require.ErrorIs(t, err, context.DeadlineExceeded)
	})
}

func TestRunDealPusher(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		_, _, err := NewRunner().Run(ctx, "singularity run deal-pusher")
		require.ErrorIs(t, err, context.DeadlineExceeded)
	})
}

func TestRunDownloadServer(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	_, _, err := NewRunner().Run(ctx, "singularity run download-server")
	require.ErrorIs(t, err, context.DeadlineExceeded)
}
