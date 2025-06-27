package contentprovider

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/cockroachdb/errors/oserror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/data-preservation-programs/singularity/store"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/fxamacker/cbor/v2"
	"github.com/ipfs/go-cid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type HTTPServer struct {
	dbNoContext         *gorm.DB
	bind                string
	enablePiece         bool
	enablePieceMetadata bool
}

func (*HTTPServer) Name() string {
	return "HTTPServer"
}

// Start is a method on the HTTPServer struct that starts the HTTP server.
//
// It sets up the Echo framework with various middleware for gzip compression, request logging, and panic recovery.
// It also sets up routes for getting piece metadata and the piece itself.
//
// The server runs in its own goroutine until the provided context is cancelled. When the context is cancelled,
// the server is shut down gracefully.
//
// The method returns two channels: a Done channel that is closed when the server has stopped, and a Fail channel
// that receives an error if the server fails to start or stop.
//
// Parameters:
//   - ctx: The context for the server. This can be used to cancel the server or set a deadline.
//
// Returns:
//   - A Done channel slice that are closed when the server has stopped.
//   - A Fail channel that receives an error if the server fails to start or stop.
//   - An error if the server fails to start.
func (s *HTTPServer) Start(ctx context.Context, exitErr chan<- error) error {
	e := echo.New()
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{}))
	e.Use(
		middleware.RequestLoggerWithConfig(
			middleware.RequestLoggerConfig{
				LogStatus: true,
				LogURI:    true,
				LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
					uri := v.URI
					status := v.Status
					latency := time.Since(v.StartTime)
					err := v.Error
					method := c.Request().Method
					if err != nil {
						logger.With(
							"status",
							status,
							"latency_ms",
							latency.Milliseconds(),
							"err",
							err,
						).Error(method + " " + uri)
					} else {
						logger.With("status", status, "latency_ms", latency.Milliseconds()).Info(method + " " + uri)
					}
					return nil
				},
			},
		),
	)
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
	if s.enablePieceMetadata {
		e.GET("/piece/metadata/:id", s.getMetadataHandler)
		e.HEAD("/piece/metadata/:id", s.getMetadataHandler)
	}
	if s.enablePiece {
		e.GET("/piece/:id", s.handleGetPiece)
		e.HEAD("/piece/:id", s.handleGetPiece)
	}
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	forceShutdown := make(chan struct{})
	shutdownErr := make(chan error, 1)

	go func() {
		err := e.Start(s.bind)
		if errors.Is(err, http.ErrServerClosed) {
			err = nil
		}
		close(forceShutdown)
		closeErr := <-shutdownErr
		if exitErr != nil {
			if err == nil {
				err = closeErr
			}
			exitErr <- err
		}
	}()

	go func() {
		select {
		case <-ctx.Done():
		case <-forceShutdown:
		}
		//nolint:contextcheck
		shutdownErr <- e.Shutdown(context.Background())
	}()

	return nil
}

func getPieceMetadata(ctx context.Context, db *gorm.DB, car model.Car) (*PieceMetadata, error) {
	db = db.WithContext(ctx)

	var attachment model.SourceAttachment
	err := db.Model(&car).Preload("Storage").Association("Attachment").Find(&attachment)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var carBlocks []model.CarBlock
	err = db.Where("car_id = ?", car.ID).Order("id ASC").Find(&carBlocks).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var files []model.File
	err = db.Where("id IN (?)", db.Model(&model.CarBlock{}).Select("file_id").Where("car_id = ?", car.ID)).Find(&files).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &PieceMetadata{
		Car:       car,
		Storage:   *attachment.Storage,
		CarBlocks: carBlocks,
		Files:     files,
	}, nil
}

// GetMetadataHandler is a function that handles HTTP requests to get the metadata of a piece.
// It takes an Echo context and a Gorm DBNoContext connection as arguments.
//
// The function first parses the piece CID from the URL parameters. If the CID is invalid, it returns a 400 Bad Request response.
//
// Then, it queries the database for the car associated with the CID. If no car is found, it returns a 404 Not Found response.
//
// Next, it retrieves the metadata of the piece. If there's an error, it returns a 500 Internal Server Error response.
//
// Finally, it removes any sensitive information from the metadata and returns it in the response.
// The format of the response depends on the "Accept" header of the request: if it's "application/cbor", the metadata is encoded as CBOR;
// otherwise, it's encoded as JSON.
//
// Parameters:
//   - c: The Echo context for the HTTP request.
//   - dbNoContext: The Gorm DBNoContext connection to use for database queries.
//
// Returns:
//   - An error if there was a problem handling the request.
func GetMetadataHandler(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	pieceCid, err := cid.Parse(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "failed to parse piece CID: "+err.Error())
	}

	var car model.Car
	ctx := c.Request().Context()
	err = db.WithContext(ctx).Where("piece_cid = ?", model.CID(pieceCid)).First(&car).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "piece not found")
	}

	if car.AttachmentID == nil {
		return c.String(http.StatusNotFound, "piece metadata not found")
	}

	metadata, err := getPieceMetadata(ctx, db, car)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}

	// Remove all credentials
	for k := range metadata.Storage.Config {
		if model.IsSecretConfigName(k) {
			delete(metadata.Storage.Config, k)
		}
	}

	acceptHeader := c.Request().Header.Get("Accept")
	switch acceptHeader {
	case "application/cbor":
		c.Response().WriteHeader(http.StatusOK)
		c.Response().Header().Set(echo.HeaderContentType, "application/cbor")
		encoder := cbor.NewEncoder(c.Response().Writer)
		return encoder.Encode(metadata)
	default:
		return c.JSON(http.StatusOK, metadata)
	}
}

func (s *HTTPServer) getMetadataHandler(c echo.Context) error {
	return GetMetadataHandler(c, s.dbNoContext.WithContext(c.Request().Context()))
}

type PieceMetadata struct {
	Car       model.Car        `cbor:"1,keyasint,omitempty" json:"car"`
	Storage   model.Storage    `cbor:"2,keyasint,omitempty" json:"storage"`
	CarBlocks []model.CarBlock `cbor:"3,keyasint,omitempty" json:"carBlocks"`
	Files     []model.File     `cbor:"4,keyasint,omitempty" json:"files"`
}

// findPiece is a method on the HTTPServer struct that finds a piece by its CID.
//
// It first queries the database for cars associated with the CID. If there's an error querying the database,
// it returns the error wrapped with additional context.
//
// If no cars are found, it returns os.ErrNotExist.
//
// Then, it tries to open each car's file. If it can't open a file or the file size doesn't match the car's file size,
// it records the error and continues with the next car.
//
// If it successfully opens a file, it returns the file, its modification time, and nil error.
//
// If it can't open any of the files, it tries to create a piece reader for each car. If it can't create a reader,
// it records the error and continues with the next car.
//
// If it successfully creates a reader, it returns the reader, the car's creation time, and nil error.
//
// If it can't create a reader for any of the cars, it returns nil, the zero time, and an aggregate error of all recorded errors.
//
// Parameters:
//   - ctx: The context for the operation. This can be used to cancel the operation or set a deadline.
//   - pieceCid: The CID of the piece to find.
//
// Returns:
//   - A ReadSeekCloser that can be used to read the piece content.
//   - -The modification time of the piece content.
//   - An error if there was a problem finding the piece.
func (s *HTTPServer) findPiece(ctx context.Context, pieceCid cid.Cid) (
	io.ReadSeekCloser,
	time.Time,
	error,
) {
	db := s.dbNoContext.WithContext(ctx)
	var cars []model.Car
	err := db.Preload("Storage").Where("piece_cid = ?", model.CID(pieceCid)).Find(&cars).Error
	if err != nil {
		return nil, time.Time{}, errors.WithStack(err)
	}

	if len(cars) == 0 {
		return nil, time.Time{}, oserror.ErrNotExist
	}

	var errs []error
	for _, car := range cars {
		if car.StoragePath == "" {
			continue
		}

		if car.Storage != nil {
			rclone, err := storagesystem.NewRCloneHandler(ctx, *car.Storage)
			if err != nil {
				errs = append(errs, errors.Wrapf(err, "failed to create rclone handler with storage %d", car.Storage.ID))
				continue
			}
			seeker, obj, err := storagesystem.Open(rclone, ctx, car.StoragePath)
			if err != nil {
				errs = append(errs, errors.Wrapf(err, "failed to open storage path %s", car.StoragePath))
				continue
			}
			return seeker, obj.ModTime(ctx), nil
		}

		file, err := os.Open(car.StoragePath)
		if err != nil {
			errs = append(errs, errors.Wrapf(err, "failed to open file %s", car.StoragePath))
			continue
		}
		fileInfo, err := file.Stat()
		if err != nil {
			_ = file.Close()
			errs = append(errs, errors.Wrapf(err, "failed to stat file %s", car.StoragePath))
			continue
		}
		if fileInfo.Size() != car.FileSize {
			_ = file.Close()
			errs = append(errs, errors.Wrapf(err, "CAR file size mismatch for %s. expected %d, actual %d.", car.StoragePath, car.FileSize, fileInfo.Size()))
			continue
		}
		return file, fileInfo.ModTime(), nil
	}

	for _, car := range cars {
		if car.AttachmentID == nil {
			continue
		}
		metadata, err := getPieceMetadata(ctx, s.dbNoContext.WithContext(ctx), car)
		if err != nil {
			errs = append(errs, errors.Wrap(err, "failed to get piece metadata"))
			continue
		}
		reader, err := store.NewPieceReader(ctx, metadata.Car, metadata.Storage, metadata.CarBlocks, metadata.Files)
		if err != nil {
			errs = append(errs, errors.Wrap(err, "failed to create piece reader"))
			continue
		}
		return reader, car.CreatedAt, nil
	}

	return nil, time.Time{}, &util.AggregateError{Errors: errs}
}

func SetCommonHeaders(c echo.Context, pieceCid string) {
	c.Response().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", pieceCid+".car"))
	c.Response().Header().Set("Content-Type", "application/vnd.ipld.car; version=1")
	c.Response().Header().Set("Accept-Ranges", "bytes")
	c.Response().Header().Set("Etag", "\""+pieceCid+"\"")
}

// handleGetPiece is a method on the HTTPServer struct that handles HTTP requests to get a piece.
//
// It first parses the piece CID from the URL parameters. If the CID is invalid, it returns a 400 Bad Request response.
//
// Then, it tries to find the piece in the storage. If the piece is not found, it returns a 404 Not Found response.
// If there's an error finding the piece, it returns a 500 Internal Server Error response.
//
// If the piece is found, it sets common headers on the response and serves the piece content using http.ServeContent.
// The name of the served content is the string representation of the piece CID with a ".car" extension.
//
// Parameters:
//   - c: The Echo context for the HTTP request.
//
// Returns:
//   - An error if there was a problem handling the request.
func (s *HTTPServer) handleGetPiece(c echo.Context) error {
	id := c.Param("id")
	pieceCid, err := cid.Parse(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "failed to parse piece CID: "+err.Error())
	}
	if pieceCid.Type() != cid.FilCommitmentUnsealed {
		return c.String(http.StatusBadRequest, "CID is not a commp")
	}

	reader, lastModified, err := s.findPiece(c.Request().Context(), pieceCid)
	if oserror.IsNotExist(err) {
		return c.String(http.StatusNotFound, "piece not found")
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to find piece: "+err.Error())
	}

	defer func() { _ = reader.Close() }()
	SetCommonHeaders(c, pieceCid.String())
	http.ServeContent(
		c.Response(),
		c.Request(),
		pieceCid.String()+".car",
		lastModified,
		reader,
	)

	return nil
}
