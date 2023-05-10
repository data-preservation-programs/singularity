package service

import (
	"context"
	"fmt"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type ContentProviderService struct {
	db   *gorm.DB
	bind string
}

func NewContentProviderService(db *gorm.DB, bind string) *ContentProviderService {
	return &ContentProviderService{db: db, bind: bind}
}

func (s *ContentProviderService) Start() {
	logger := logging.Logger("contentprovider")
	e := echo.New()
	current := logging.GetConfig().Level
	if logging.LevelInfo < current {
		logging.SetAllLoggers(logging.LevelInfo)
	}

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			uri := v.URI
			status := v.Status
			latency := time.Now().Sub(v.StartTime)
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
	e.Use(middleware.Recover())
	e.GET("/piece/:id", s.handleGetPiece)
	e.HEAD("/piece/:id", s.handleHeadPiece)

	err := e.Start(s.bind)
	if err != nil {
		panic(err)
	}
}

func (s *ContentProviderService) findPieceAsPieceReader(ctx context.Context, pieceCid cid.Cid) (io.ReaderAt, *model.Car, error) {
	var car model.Car
	err := s.db.WithContext(ctx).Where("piece_cid = ?", pieceCid.String()).First(&car).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, os.ErrNotExist
	}
	if err != nil {
		return nil, nil, fmt.Errorf("failed to query for CARs: %w", err)
	}

	// TODO
	return nil, &car, nil
}

func (s *ContentProviderService) findPiece(ctx context.Context, pieceCid cid.Cid) (*os.File, os.FileInfo, error) {
	var cars []model.Car
	err := s.db.WithContext(ctx).Where("piece_cid = ?", pieceCid.String()).Find(&cars).Error
	if err != nil {
		return nil, nil, fmt.Errorf("failed to query for CARs: %w", err)
	}

	if len(cars) == 0 {
		return nil, nil, os.ErrNotExist
	}

	for _, car := range cars {
		if car.FilePath == "" {
			continue
		}

		file, err := os.Open(car.FilePath)
		if err != nil {
			continue
		}
		fileInfo, err := file.Stat()
		if err != nil {
			file.Close()
			continue
		}
		return file, fileInfo, nil
	}

	return nil, nil, os.ErrNotExist
}

func (s *ContentProviderService) setCommonHeaders(c echo.Context, pieceCid string) {
	c.Response().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", pieceCid+".car"))
	c.Response().Header().Set("Content-Type", "application/piece")
	c.Response().Header().Set("Accept-Ranges", "bytes")
}

func (s *ContentProviderService) handleHeadPiece(c echo.Context) error {
	id := c.Param("id")
	pieceCid, err := cid.Parse(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "failed to parse piece CID: "+err.Error())
	}

	file, fileInfo, err := s.findPiece(c.Request().Context(), pieceCid)
	if err == nil {
		defer file.Close()
		s.setCommonHeaders(c, pieceCid.String())
		c.Response().Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))
		return c.NoContent(http.StatusOK)
	}

	_, car, err := s.findPieceAsPieceReader(c.Request().Context(), pieceCid)
	if err == nil {
		s.setCommonHeaders(c, pieceCid.String())
		c.Response().Header().Set("Content-Length", strconv.FormatInt(int64(car.FileSize), 10))
		return c.NoContent(http.StatusOK)
	}

	if err != nil && !os.IsNotExist(err) {
		return c.String(http.StatusInternalServerError, "failed to find CAR file: "+err.Error())
	}

	return c.String(http.StatusNotFound, "piece not found")
}

func (s *ContentProviderService) handleGetPiece(c echo.Context) error {
	id := c.Param("id")
	pieceCid, err := cid.Parse(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "failed to parse piece CID: "+err.Error())
	}

	var reader io.ReaderAt
	var lastModified time.Time
	var fileSize int64
	fil, fInfo, err := s.findPiece(c.Request().Context(), pieceCid)
	if err == nil {
		defer fil.Close()
		reader = fil
		lastModified = fInfo.ModTime()
		fileSize = fInfo.Size()
	} else {
		pieceReader, car, err := s.findPieceAsPieceReader(c.Request().Context(), pieceCid)
		switch {
		case err == nil:
			reader = pieceReader
			lastModified = car.CreatedAt
			fileSize = int64(car.FileSize)
		case os.IsNotExist(err):
			return c.String(http.StatusNotFound, "piece not found")
		default:
			return c.String(http.StatusInternalServerError, "failed to find CAR file: "+err.Error())
		}
	}

	s.setCommonHeaders(c, pieceCid.String())
	rangeHeader := c.Request().Header.Get("Range")
	if rangeHeader == "" {
		http.ServeContent(c.Response(), c.Request(), pieceCid.String()+".car", lastModified, io.NewSectionReader(reader, 0, fileSize))
		return nil
	}

	// Parse Range header
	rangeStr := strings.TrimPrefix(rangeHeader, "bytes=")
	rangeParts := strings.Split(rangeStr, "-")
	start, err := strconv.ParseInt(rangeParts[0], 10, 64)
	if err != nil {
		return c.String(http.StatusRequestedRangeNotSatisfiable, "invalid range")
	}

	var end int64
	if len(rangeParts) > 1 && rangeParts[1] != "" {
		end, err = strconv.ParseInt(rangeParts[1], 10, 64)
		if err != nil {
			return c.String(http.StatusRequestedRangeNotSatisfiable, "invalid range")
		}
	} else {
		end = fileSize - 1
	}

	if start > end || end >= fileSize {
		return c.String(http.StatusRequestedRangeNotSatisfiable, "invalid range")
	}

	// Set required headers for partial content
	c.Response().Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, fileSize))
	c.Response().Header().Set("Content-Length", strconv.FormatInt(end-start+1, 10))

	// Send the specified range of bytes
	c.Response().WriteHeader(http.StatusPartialContent)
	http.ServeContent(c.Response(), c.Request(), pieceCid.String()+".car", lastModified, io.NewSectionReader(reader, start, end-start+1))
	return nil
}
