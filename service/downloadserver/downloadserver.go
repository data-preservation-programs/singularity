package downloadserver

import (
	"context"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
	"io"
	"bufio"


	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/data-preservation-programs/singularity/service/contentprovider"
	"github.com/data-preservation-programs/singularity/store"
	"github.com/fxamacker/cbor/v2"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-log/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const shutdownTimeout = 5 * time.Second

// Maximum retry attempts
const maxRetries = 5

// Initial wait duration before retrying (doubles each retry)
const initialBackoff = 120 * time.Second // Relax for Fail2Ban rules

type DownloadServer struct {
    bind         string
    api          string
    config       map[string]string
    clientConfig model.ClientConfig
    usageCache   *UsageCache[contentprovider.PieceMetadata]

    metadataCache sync.Map // Cache for ongoing metadata requests
}

type cacheItem[C any] struct {
	item         C
	usageCount   int
	lastAccessed time.Time
}

type UsageCache[C any] struct {
	data   map[string]*cacheItem[C]
	mu     sync.Mutex
	ttl    time.Duration
	cancel context.CancelFunc
}

func (c *UsageCache[C]) Close() {
	c.cancel()
}

func NewUsageCache[C any](ttl time.Duration) *UsageCache[C] {
	ctx, cancel := context.WithCancel(context.Background())
	cache := &UsageCache[C]{
		data:   make(map[string]*cacheItem[C]),
		ttl:    ttl,
		cancel: cancel,
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(ttl):
			}
			now := time.Now()
			cache.mu.Lock()
			for key, item := range cache.data {
				if item.usageCount <= 0 && now.Sub(item.lastAccessed) > cache.ttl {
					delete(cache.data, key)
				}
			}
			cache.mu.Unlock()
		}
	}()
	return cache
}

func (c *UsageCache[C]) Get(key string) (*C, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	item, ok := c.data[key]
	if !ok {
		return nil, false
	}
	item.usageCount++
	item.lastAccessed = time.Now()
	return &item.item, true
}

func (c *UsageCache[C]) Set(key string, item C) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = &cacheItem[C]{
		item:         item,
		usageCount:   1,
		lastAccessed: time.Now(),
	}
}

func (c *UsageCache[C]) Done(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	item, ok := c.data[key]
	if !ok {
		return
	}
	item.usageCount--
}

func (d *DownloadServer) handleGetPiece(c echo.Context) error {
	id := c.Param("id")
	pieceCid, err := cid.Parse(id)
	if err != nil {
		Logger.Errorw("Invalid piece CID", "id", id, "error", err)
		return c.String(http.StatusBadRequest, "failed to parse piece CID: "+err.Error())
	}

	if pieceCid.Type() != cid.FilCommitmentUnsealed {
		Logger.Warnw("Received invalid CID type", "id", id, "type", pieceCid.Type())
		return c.String(http.StatusBadRequest, "CID is not a commp")
	}

	Logger.Infow("Processing request", "pieceCID", pieceCid.String())

	// **Check usage cache first**
	var pieceMetadata *contentprovider.PieceMetadata
	var ok bool
	pieceMetadata, ok = d.usageCache.Get(pieceCid.String())

	if !ok {
		Logger.Infow("Metadata not found in cache, checking for ongoing fetch", "pieceCID", pieceCid.String())

		// **Check if another request is already fetching this metadata**
		metadataChan := make(chan *contentprovider.PieceMetadata, 1)
		actual, loaded := d.metadataCache.LoadOrStore(pieceCid.String(), metadataChan)

		if loaded {
			// **Wait for the existing metadata fetch**
			Logger.Infow("Waiting for ongoing metadata fetch", "pieceCID", pieceCid.String())
			metadataChan = actual.(chan *contentprovider.PieceMetadata)
			pieceMetadata, ok = <-metadataChan
			if pieceMetadata != nil {
				Logger.Infow("Received metadata from another request", "pieceCID", pieceCid.String())
				d.usageCache.Set(pieceCid.String(), *pieceMetadata)
				goto ServeContent
			} else {
				Logger.Errorw("Metadata fetch failed, returning error", "pieceCID", pieceCid.String())
				return c.String(http.StatusInternalServerError, "Failed to fetch metadata")
			}
		}

		// **Fetch metadata (call standalone function)**
		var statusCode int
		Logger.Infow("Fetching metadata from API", "pieceCID", pieceCid.String())

		pieceMetadata, statusCode, err = GetMetadata(c.Request().Context(), d.api, d.config, d.clientConfig, pieceCid.String())

		// **Handle errors**
		if err != nil {
			Logger.Errorw("Failed to query metadata API", "pieceCID", pieceCid.String(), "statusCode", statusCode, "error", err)

			// **Remove metadata cache entry to allow retries**
			d.metadataCache.Delete(pieceCid.String())

			close(metadataChan) // **Close channel so waiting goroutines don't hang**
			if statusCode >= 400 {
				return c.String(statusCode, "failed to query metadata API: "+err.Error())
			}
			return c.String(http.StatusInternalServerError, "failed to query metadata API: "+err.Error())
		}

		Logger.Infow("Successfully fetched metadata", "pieceCID", pieceCid.String())

		// **Ensure metadata is cached before notifying other requests**
		d.usageCache.Set(pieceCid.String(), *pieceMetadata)
		metadataChan <- pieceMetadata
		close(metadataChan) // Close channel after setting metadata
	}

ServeContent:
	defer func() {
		d.usageCache.Done(pieceCid.String())
	}()

	// **Log CAR file and storage details**
	Logger.Infow("Attempting to create piece reader",
		"pieceCID", pieceCid.String(),
		"carFile", pieceMetadata.Car,
		"storageType", pieceMetadata.Storage.Type)

	pieceReader, err := store.NewPieceReader(c.Request().Context(), pieceMetadata.Car, pieceMetadata.Storage, pieceMetadata.CarBlocks, pieceMetadata.Files)
	if err != nil {
		Logger.Errorw("Failed to create piece reader", "pieceCID", pieceCid.String(), "error", err)
		return c.String(http.StatusInternalServerError, "failed to create piece reader: "+err.Error())
	}
	defer pieceReader.Close()

	Logger.Infow("Serving content with increased buffer", "pieceCID", pieceCid.String(), "filename", pieceCid.String()+".car")

	contentprovider.SetCommonHeaders(c, pieceCid.String())

	// **Use buffered writer (16MB buffer) to improve streaming performance**
	bufferedWriter := bufio.NewWriterSize(c.Response().Writer, 16*1024*1024) // 16MB buffer
	defer bufferedWriter.Flush()

	_, err = io.Copy(bufferedWriter, pieceReader)
	if err != nil {
		Logger.Errorw("Error streaming content", "pieceCID", pieceCid.String(), "error", err)
		return c.String(http.StatusInternalServerError, "Error streaming content: "+err.Error())
	}

	return nil
}

func GetMetadata(
	ctx context.Context,
	api string,
	config map[string]string,
	clientConfig model.ClientConfig,
	pieceCid string) (*contentprovider.PieceMetadata, int, error) {
	api = strings.TrimSuffix(api, "/")

	var lastErr error
	var lastStatusCode int

	for attempt := 1; attempt <= maxRetries; attempt++ {
		Logger.Infow("Fetching metadata from API", "pieceCID", pieceCid, "attempt", attempt)

		// Set request timeout (increase timeout to 120s per attempt)
		reqCtx, cancel := context.WithTimeout(ctx, 120*time.Second)
		defer cancel()

		req, err := http.NewRequestWithContext(reqCtx, http.MethodGet, api+"/piece/metadata/"+pieceCid, nil)
		if err != nil {
			Logger.Errorw("Failed to create request", "pieceCID", pieceCid, "error", err)
			return nil, 0, errors.WithStack(err)
		}

		req.Header.Add("Accept", "application/cbor")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			Logger.Warnw("Failed to reach metadata API", "pieceCID", pieceCid, "attempt", attempt, "error", err)
			lastErr = err
			time.Sleep(exponentialBackoff(attempt))
			continue // Retry
		}
		defer resp.Body.Close()

		// Handle known HTTP errors (don't retry on these)
		if resp.StatusCode == http.StatusForbidden || resp.StatusCode == http.StatusNotFound {
			Logger.Warnw("Permanent metadata API error", "pieceCID", pieceCid, "statusCode", resp.StatusCode)
			return nil, resp.StatusCode, errors.Errorf("permanent error: %s", resp.Status)
		}

		if resp.StatusCode != http.StatusOK {
			Logger.Errorw("Metadata API returned error", "pieceCID", pieceCid, "statusCode", resp.StatusCode, "attempt", attempt)
			lastStatusCode = resp.StatusCode
			lastErr = errors.Errorf("failed to get metadata: %s", resp.Status)
			time.Sleep(exponentialBackoff(attempt))
			continue // Retry
		}

		// Decode metadata response
		var pieceMetadata contentprovider.PieceMetadata
		err = cbor.NewDecoder(resp.Body).Decode(&pieceMetadata)
		if err != nil {
			Logger.Errorw("Failed to decode metadata", "pieceCID", pieceCid, "attempt", attempt, "error", err)
			lastErr = errors.Wrap(err, "failed to decode metadata")
			time.Sleep(exponentialBackoff(attempt))
			continue // Retry
		}

		Logger.Infow("Successfully fetched metadata", "pieceCID", pieceCid)
		return &pieceMetadata, 0, nil
	}

	// If all retries fail, return the last error
	Logger.Errorw("All metadata fetch attempts failed", "pieceCID", pieceCid, "error", lastErr)
	return nil, lastStatusCode, lastErr
}

// Exponential backoff function (2^attempt with jitter)
func exponentialBackoff(attempt int) time.Duration {
	baseDelay := time.Duration(1<<attempt) * initialBackoff
	jitter := time.Duration(rand.Intn(500)) * time.Millisecond
	return baseDelay + jitter
}

func (d *DownloadServer) Start(ctx context.Context, exitErr chan<- error) error {
	e := echo.New()
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		Skipper:           middleware.DefaultSkipper,
		StackSize:         4 << 10, // 4 KiB
		DisableStackAll:   false,
		DisablePrintStack: false,
		LogLevel:          0,
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			Logger.Errorw("panic", "err", err, "stack", string(stack))
			return nil
		},
	}))
	e.GET("/piece/:id", d.handleGetPiece)
	e.HEAD("/piece/:id", d.handleGetPiece)
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	shutdownErr := make(chan error, 1)
	forceShutdown := make(chan struct{})

	go func() {
		runErr := e.Start(d.bind)
		close(forceShutdown)

		err := <-shutdownErr
		d.usageCache.Close()

		if exitErr != nil {
			if runErr != nil {
				err = runErr
			}
			exitErr <- err
		}
	}()

	go func() {
		select {
		case <-ctx.Done():
		case <-forceShutdown:
		}
		ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()
		//nolint:contextcheck
		shutdownErr <- e.Shutdown(ctx)
	}()

	return nil
}

func (d *DownloadServer) Name() string {
	return "DownloadServer"
}

var Logger = log.Logger("downloadserver")

var _ service.Server = &DownloadServer{}

func NewDownloadServer(bind string, api string, config map[string]string, clientConfig model.ClientConfig) *DownloadServer {
	return &DownloadServer{
		bind:         bind,
		api:          api,
		config:       config,
		clientConfig: clientConfig,
		usageCache:   NewUsageCache[contentprovider.PieceMetadata](time.Minute),
	}
}
