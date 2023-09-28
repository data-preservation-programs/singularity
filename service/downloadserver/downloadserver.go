package downloadserver

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"

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

type DownloadServer struct {
	bind               string
	api                string
	config             map[string]string
	clientConfig       model.ClientConfig
	metadataCache      map[cid.Cid]contentprovider.PieceMetadata
	metadataUsageCount map[cid.Cid]int
	mu                 sync.RWMutex
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
	var pieceMetadata contentprovider.PieceMetadata
	d.mu.RLock()
	var ok bool
	pieceMetadata, ok = d.metadataCache[pieceCid]
	d.mu.RUnlock()
	if !ok {
		var statusCode int
		pieceMetadata, statusCode, err = GetMetadata(c.Request().Context(), d.api, d.config, d.clientConfig, pieceCid.String())
		if err != nil && statusCode >= 400 {
			return c.String(statusCode, "failed to query metadata API: "+err.Error())
		}
		if err != nil {
			return c.String(http.StatusInternalServerError, "failed to query metadata API: "+err.Error())
		}
		d.mu.Lock()
		d.metadataCache[pieceCid] = pieceMetadata
		d.mu.Unlock()
	}
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
	pieceCid string) (contentprovider.PieceMetadata, int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, api+"/piece/metadata/"+pieceCid, nil)
	if err != nil {
		return contentprovider.PieceMetadata{}, 0, errors.WithStack(err)
	}

	req.Header.Add("Accept", "application/cbor")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return contentprovider.PieceMetadata{}, 0, errors.WithStack(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return contentprovider.PieceMetadata{}, resp.StatusCode, errors.Errorf("failed to get metadata: %s", resp.Status)
	}

	var pieceMetadata contentprovider.PieceMetadata
	err = cbor.NewDecoder(resp.Body).Decode(&pieceMetadata)
	if err != nil {
		return contentprovider.PieceMetadata{}, 0, errors.Wrap(err, "failed to decode metadata")
	}

	cfg := make(map[string]string)
	backend, ok := storagesystem.BackendMap[pieceMetadata.Storage.Type]
	if !ok {
		return contentprovider.PieceMetadata{}, 0, errors.Newf("storage type %s is not supported", pieceMetadata.Storage.Type)
	}

	prefix := pieceMetadata.Storage.Type + "-"
	provider := pieceMetadata.Storage.Config["provider"]
	providerOptions, err := underscore.Find(backend.ProviderOptions, func(providerOption storagesystem.ProviderOptions) bool {
		return providerOption.Provider == provider
	})
	if err != nil {
		return contentprovider.PieceMetadata{}, 0, errors.Newf("provider '%s' is not supported", provider)
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
	return pieceMetadata, 0, nil
}

func (d *DownloadServer) Start(ctx context.Context) ([]service.Done, service.Fail, error) {
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
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	done := make(chan struct{})
	fail := make(chan error)
	go func() {
		err := e.Start(d.bind)
		if err != nil {
			select {
			case <-ctx.Done():
			case fail <- err:
			}
		}
	}()
	go func() {
		defer close(done)
		<-ctx.Done()
		//nolint:contextcheck
		err := e.Shutdown(context.Background())
		if err != nil {
			fail <- err
		}
	}()
	return []service.Done{done}, fail, nil
}

func (d *DownloadServer) Name() string {
	return "DownloadServer"
}

var Logger = log.Logger("downloadserver")

var _ service.Server = &DownloadServer{}

func NewDownloadServer(bind string, api string, config map[string]string, clientConfig model.ClientConfig) *DownloadServer {
	return &DownloadServer{
		bind:               bind,
		api:                api,
		config:             config,
		clientConfig:       clientConfig,
		metadataCache:      make(map[cid.Cid]contentprovider.PieceMetadata),
		metadataUsageCount: make(map[cid.Cid]int),
	}
}
