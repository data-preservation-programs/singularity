package downloadserver

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/data-preservation-programs/singularity/service/contentprovider"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/data-preservation-programs/singularity/store"
	"github.com/fxamacker/cbor/v2"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-log/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rjNemo/underscore"
)

const shutdownTimeout = 5 * time.Second

type DownloadServer struct {
	bind         string
	api          string
	config       map[string]string
	clientConfig model.ClientConfig
	usageCache   *UsageCache[contentprovider.PieceMetadata]
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
		return c.String(http.StatusBadRequest, "failed to parse piece CID: "+err.Error())
	}
	if pieceCid.Type() != cid.FilCommitmentUnsealed {
		return c.String(http.StatusBadRequest, "CID is not a commp")
	}
	var pieceMetadata *contentprovider.PieceMetadata
	var ok bool
	pieceMetadata, ok = d.usageCache.Get(pieceCid.String())
	if !ok {
		var statusCode int
		pieceMetadata, statusCode, err = GetMetadata(c.Request().Context(), d.api, d.config, d.clientConfig, pieceCid.String())
		if err != nil && statusCode >= 400 {
			return c.String(statusCode, "failed to query metadata API: "+err.Error())
		}
		if err != nil {
			return c.String(http.StatusInternalServerError, "failed to query metadata API: "+err.Error())
		}
		d.usageCache.Set(pieceCid.String(), *pieceMetadata)
	}
	defer func() {
		d.usageCache.Done(pieceCid.String())
	}()
	pieceReader, err := store.NewPieceReader(c.Request().Context(), pieceMetadata.Car, pieceMetadata.Storage, pieceMetadata.CarBlocks, pieceMetadata.Files)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to create piece reader: "+err.Error())
	}
	defer pieceReader.Close()
	contentprovider.SetCommonHeaders(c, pieceCid.String())
	http.ServeContent(
		c.Response(),
		c.Request(),
		pieceCid.String()+".car",
		pieceMetadata.Car.CreatedAt,
		pieceReader,
	)

	return nil
}

func GetMetadata(
	ctx context.Context,
	api string,
	config map[string]string,
	clientConfig model.ClientConfig,
	pieceCid string) (*contentprovider.PieceMetadata, int, error) {
	api = strings.TrimSuffix(api, "/")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, api+"/piece/metadata/"+pieceCid, nil)
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}

	req.Header.Add("Accept", "application/cbor")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode, errors.Errorf("failed to get metadata: %s", resp.Status)
	}

	var pieceMetadata contentprovider.PieceMetadata
	err = cbor.NewDecoder(resp.Body).Decode(&pieceMetadata)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to decode metadata")
	}

	cfg := make(map[string]string)
	backend, ok := storagesystem.BackendMap[pieceMetadata.Storage.Type]
	if !ok {
		return nil, 0, errors.Newf("storage type %s is not supported", pieceMetadata.Storage.Type)
	}

	prefix := pieceMetadata.Storage.Type + "-"
	provider := pieceMetadata.Storage.Config["provider"]
	providerOptions, err := underscore.Find(backend.ProviderOptions, func(providerOption storagesystem.ProviderOptions) bool {
		return providerOption.Provider == provider
	})
	if err != nil {
		return nil, 0, errors.Newf("provider '%s' is not supported", provider)
	}

	for _, option := range providerOptions.Options {
		if option.Default != nil {
			cfg[option.Name] = fmt.Sprintf("%v", option.Default)
		}
	}

	for key, value := range pieceMetadata.Storage.Config {
		cfg[key] = value
	}

	for key, value := range config {
		if strings.HasPrefix(key, prefix) {
			trimmed := strings.TrimPrefix(key, prefix)
			snake := strings.ReplaceAll(trimmed, "-", "_")
			cfg[snake] = value
		}
	}

	pieceMetadata.Storage.Config = cfg
	storage.OverrideStorageWithClientConfig(&pieceMetadata.Storage, clientConfig)
	return &pieceMetadata, 0, nil
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
