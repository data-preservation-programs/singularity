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
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		
		done := make(chan error, 1)
		go func() {
			_, _, err := NewRunner().Run(ctx, "singularity run deal-tracker")
			done <- err
		}()
		
		// Give the service time to start and initialize
		time.Sleep(2 * time.Second)
		cancel()
		
		select {
		case err := <-done:
			require.ErrorIs(t, err, context.Canceled)
		case <-time.After(5 * time.Second):
			t.Fatal("Service did not shut down within timeout")
		}
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
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		
		done := make(chan error, 1)
		go func() {
			_, _, err := NewRunner().Run(ctx, "singularity run dataset-worker")
			done <- err
		}()
		
		// Give the service time to start and initialize
		time.Sleep(2 * time.Second)
		cancel()
		
		select {
		case err := <-done:
			require.ErrorIs(t, err, context.Canceled)
		case <-time.After(5 * time.Second):
			t.Fatal("Service did not shut down within timeout")
		}
	})
}

func TestRunContentProvider(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		
		done := make(chan error, 1)
		go func() {
			_, _, err := NewRunner().Run(ctx, "singularity run content-provider --http-bind "+contentProviderBind)
			done <- err
		}()
		
		// Give the service time to start and initialize
		time.Sleep(2 * time.Second)
		cancel()
		
		select {
		case err := <-done:
			require.ErrorIs(t, err, context.Canceled)
		case <-time.After(5 * time.Second):
			t.Fatal("Service did not shut down within timeout")
		}
	})
}

func TestRunDealPusher(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		
		done := make(chan error, 1)
		go func() {
			_, _, err := NewRunner().Run(ctx, "singularity run deal-pusher")
			done <- err
		}()
		
		// Give the service time to start and initialize
		time.Sleep(2 * time.Second)
		cancel()
		
		select {
		case err := <-done:
			require.ErrorIs(t, err, context.Canceled)
		case <-time.After(5 * time.Second):
			t.Fatal("Service did not shut down within timeout")
		}
	})
}

func TestRunDownloadServer(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	
	done := make(chan error, 1)
	go func() {
		_, _, err := NewRunner().Run(ctx, "singularity run download-server")
		done <- err
	}()
	
	// Give the service time to start and initialize
	time.Sleep(2 * time.Second)
	cancel()
	
	select {
	case err := <-done:
		require.ErrorIs(t, err, context.Canceled)
	case <-time.After(5 * time.Second):
		t.Fatal("Service did not shut down within timeout")
	}
}
