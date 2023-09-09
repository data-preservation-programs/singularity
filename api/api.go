package api

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"time"

	handler2 "github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/data-preservation-programs/singularity/handler/deal"
	"github.com/data-preservation-programs/singularity/handler/deal/schedule"
	"github.com/data-preservation-programs/singularity/handler/file"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/handler/job"
	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/replication"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/data-preservation-programs/singularity/service/contentprovider"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/ybbus/jsonrpc/v3"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/dashboard"
	"github.com/data-preservation-programs/singularity/database"
	_ "github.com/data-preservation-programs/singularity/docs/swagger"
	logging "github.com/ipfs/go-log/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

type Server struct {
	db          *gorm.DB
	listener    net.Listener
	lotusClient jsonrpc.RPCClient
	dealMaker   replication.DealMaker
	closer      io.Closer
}

func (s Server) Name() string {
	return "api"
}

// @Summary Get metadata for a piece
// @Description Get metadata for a piece for how it may be reassembled from the data source
// @Tags Piece
// @Produce json
// @Param id path string true "Piece CID"
// @Success 200 {object} store.PieceReader
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /piece/{id}/metadata [get]
func (s Server) getMetadataHandler(c echo.Context) error {
	return contentprovider.GetMetadataHandler(c, s.db)
}

func Run(c *cli.Context) error {
	connString := c.String("database-connection-string")

	bind := c.String("bind")

	lotusAPI := c.String("lotus-api")
	lotusToken := c.String("lotus-token")

	listener, err := net.Listen("tcp", bind)
	if err != nil {
		return errors.WithStack(err)
	}
	server, err := InitServer(c.Context, APIParams{
		ConnString: connString,
		Listener:   listener,
		LotusAPI:   lotusAPI,
		LotusToken: lotusToken,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	logger.Info("Starting Singularity API server...")
	return service.StartServers(c.Context, logger, server)
}

type APIParams struct {
	Listener   net.Listener
	LotusAPI   string
	LotusToken string
	ConnString string
}

func InitServer(ctx context.Context, params APIParams) (Server, error) {
	db, closer, err := database.OpenWithLogger(params.ConnString)
	if err != nil {
		return Server{}, errors.WithStack(err)
	}
	h, err := util.InitHost(nil)
	if err != nil {
		return Server{}, errors.Wrap(err, "failed to init host")
	}

	return Server{
		db:          db,
		listener:    params.Listener,
		lotusClient: util.NewLotusClient(params.LotusAPI, params.LotusToken),
		dealMaker: replication.NewDealMaker(
			util.NewLotusClient(params.LotusAPI, params.LotusToken),
			h,
			time.Hour,
			time.Minute*5,
		),
		closer: closer,
	}, nil
}

func toEchoHandler[Req any, Resp any](
	s Server,
	handler func(ctx context.Context, request handler2.Request[Req], dep handler2.Dependency) (Resp, error)) echo.HandlerFunc {
	return func(c echo.Context) error {
		paramValues := c.ParamValues()
		var params []string
		for _, param := range paramValues {
			unescaped, err := url.PathUnescape(param)
			if err != nil {
				return c.JSON(http.StatusBadRequest, HTTPError{Err: fmt.Sprintf("failed to decode path parameter: %s", err)})
			}
			params = append(params, unescaped)
		}
		var body Req
		err := c.Bind(&body)
		if err != nil {
			return c.JSON(http.StatusBadRequest, HTTPError{Err: fmt.Sprintf("failed to bind request body: %s", err)})
		}
		response, err := handler(c.Request().Context(), handler2.Request[Req]{
			Params:  params,
			Payload: body,
		}, handler2.Dependency{
			LotusClient: s.lotusClient,
			DealMaker:   s.dealMaker,
			DB:          s.db,
		})
		if err != nil {
			return httpResponseFromError(c, err)
		}
		if response == nil {
			return c.NoContent(http.StatusNoContent)
		}
		return c.JSON(http.StatusOK, response)
	}
}

func (s Server) setupRoutes(e *echo.Echo) {
	// Storage
	e.POST("/api/storage/:type", s.toEchoHandler(storage.Default.CreateStorageHandler))
	e.GET("/api/storage/:name/explore/:path", s.toEchoHandler(storage.Default.ExploreHandler))
	e.GET("/api/storage", s.toEchoHandler(storage.Default.ListStoragesHandler))
	e.DELETE("/api/storage/:name", s.toEchoHandler(storage.Default.RemoveHandler))
	e.PATCH("/api/storage/:name", s.toEchoHandler(storage.Default.UpdateStorageHandler))

	// Preparation
	e.POST("/api/preparation", toEchoHandler(s, dataprep.Default.CreatePreparationHandler))
	e.GET("/api/preparation", s.toEchoHandler(dataprep.Default.ListHandler))
	e.GET("/api/preparation/:id", s.toEchoHandler(job.Default.GetStatusHandler))
	e.GET("/api/preparation/:id/schedules", s.toEchoHandler(dataprep.Default.ListSchedulesHandler))

	// Job management
	e.POST("/api/preparation/:id/source/:name/start-daggen", s.toEchoHandler(job.Default.StartDagGenHandler))
	e.POST("/api/preparation/:id/source/:name/pause-daggen", s.toEchoHandler(job.Default.PauseDagGenHandler))
	e.POST("/api/preparation/:id/source/:name/start-scan", s.toEchoHandler(job.Default.StartScanHandler))
	e.POST("/api/preparation/:id/source/:name/pause-scan", s.toEchoHandler(job.Default.PauseScanHandler))
	e.POST("/api/preparation/:id/source/:name/start-pack", s.toEchoHandler(job.Default.StartPackHandler))
	e.POST("/api/preparation/:id/source/:name/pause-pack", s.toEchoHandler(job.Default.PausePackHandler))
	e.POST("/api/preparation/:id/source/:name/start-pack/:job_id", s.toEchoHandler(job.Default.StartPackHandler))
	e.POST("/api/preparation/:id/source/:name/pause-pack/:job_id", s.toEchoHandler(job.Default.PausePackHandler))
	e.POST("/api/preparation/:id/source/:name/finalize", s.toEchoHandler(job.Default.PrepareToPackSourceHandler))
	e.POST("/api/job/:id/pack", s.toEchoHandler(job.Default.PackHandler))

	// storage attachment
	e.POST("/api/preparation/:id/output/:name", s.toEchoHandler(dataprep.Default.AddOutputStorageHandler))
	e.POST("/api/preparation/:id/source/:name", s.toEchoHandler(dataprep.Default.AddSourceStorageHandler))
	e.DELETE("/api/preparation/:id/output/:name", s.toEchoHandler(dataprep.Default.RemoveOutputStorageHandler))

	// Explore
	e.GET("/api/preparation/:id/source/:name/explore/:path", s.toEchoHandler(dataprep.Default.ExploreHandler))

	// Piece
	e.GET("/api/preparation/:id/piece", s.toEchoHandler(dataprep.Default.ListPiecesHandler))
	e.POST("/api/preparation/:id/piece", s.toEchoHandler(dataprep.Default.AddPieceHandler))

	// Wallet
	e.POST("/api/wallet", s.toEchoHandler(wallet.Default.ImportHandler))
	e.GET("/api/wallet", s.toEchoHandler(wallet.Default.ListHandler))
	e.DELETE("/api/wallet/:address", s.toEchoHandler(wallet.Default.RemoveHandler))

	// Wallet Association
	e.POST("/api/preparation/:id/wallet/:wallet", s.toEchoHandler(wallet.Default.AttachHandler))
	e.GET("/api/preparation/:id/wallet", s.toEchoHandler(wallet.Default.ListAttachedHandler))
	e.DELETE("/api/preparation/:id/wallet/:wallet", s.toEchoHandler(wallet.Default.DetachHandler))

	// Piece metadata
	e.GET("/api/piece/:id/metadata", s.getMetadataHandler)

	// Deal Schedule
	e.POST("/api/send_deal", s.toEchoHandler(deal.Default.SendManualHandler))
	e.POST("/api/schedule", s.toEchoHandler(schedule.Default.CreateHandler))
	e.GET("/api/schedule", s.toEchoHandler(schedule.Default.ListHandler))
	e.POST("/api/schedule/:id/pause", s.toEchoHandler(schedule.Default.PauseHandler))
	e.POST("/api/schedule/:id/resume", s.toEchoHandler(schedule.Default.ResumeHandler))

	// Deal
	e.POST("/api/deal", s.toEchoHandler(deal.Default.ListHandler))

	// File
	e.GET("/api/file/:id/deals", s.toEchoHandler(file.Default.GetFileDealsHandler))
	e.GET("/api/file/:id", s.toEchoHandler(file.Default.GetFileHandler))
	e.POST("/api/file/:id/prepare_to_pack", s.toEchoHandler(file.Default.PrepareToPackFileHandler))
	e.POST("/api/preparation/:id/source/:name/file", s.toEchoHandler(file.Default.PushFileHandler))
}

var logger = logging.Logger("api")

func (s Server) Start(ctx context.Context) ([]service.Done, service.Fail, error) {
	e := echo.New()
	e.Debug = true
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		Skipper:           middleware.DefaultSkipper,
		StackSize:         4 << 10, // 4 KiB
		DisableStackAll:   false,
		DisablePrintStack: false,
		LogLevel:          0,
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			logger.Errorw("panic", "err", err, "stack", string(stack))
			return nil
		},
	}))
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			uri := v.URI
			status := v.Status
			latency := time.Since(v.StartTime)
			err := v.Error
			method := c.Request().Method
			if err != nil {
				logger.With("status", status, "latency_ms", latency.Milliseconds(), "err", err).Error(method + " " + uri)
			} else {
				logger.With("status", status, "latency_ms", latency.Milliseconds()).Info(method + " " + uri)
			}
			return nil
		},
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	//nolint:contextcheck
	s.setupRoutes(e)
	efs, err := fs.Sub(dashboard.DashboardStaticFiles, "build")
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/*", echo.WrapHandler(http.FileServer(http.FS(efs))))
	e.Listener = s.listener

	done := make(chan struct{})
	fail := make(chan error)
	go func() {
		err := e.Start("")
		if err != nil {
			select {
			case <-ctx.Done():
			case fail <- err:
			}
		}
	}()
	go func() {
		<-ctx.Done()
		ctx2, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		//nolint:contextcheck
		err := e.Shutdown(ctx2)
		if err != nil {
			logger.Errorw("failed to shutdown api server", "err", err)
		}
		err = s.closer.Close()
		if err != nil {
			logger.Errorw("failed to close database connection", "err", err)
		}
		close(done)
	}()
	return []service.Done{done}, fail, nil
}

func isIntKind(kind reflect.Kind) bool {
	return kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 || kind == reflect.Int64
}

func isUIntKind(kind reflect.Kind) bool {
	return kind == reflect.Uint || kind == reflect.Uint8 || kind == reflect.Uint16 || kind == reflect.Uint32 || kind == reflect.Uint64
}

type HTTPError struct {
	Err string `json:"err"`
}

func httpResponseFromError(c echo.Context, e error) error {
	if e == nil {
		return c.String(http.StatusOK, "OK")
	}

	httpStatusCode := http.StatusInternalServerError

	if errors.Is(e, handlererror.ErrNotFound) {
		httpStatusCode = http.StatusNotFound
	}

	if errors.Is(e, handlererror.ErrInvalidParameter) {
		httpStatusCode = http.StatusBadRequest
	}

	if errors.Is(e, handlererror.ErrDuplicateRecord) {
		httpStatusCode = http.StatusConflict
	}

	logger.Errorf("%+v", e)
	return c.JSON(httpStatusCode, HTTPError{Err: e.Error()})
}
