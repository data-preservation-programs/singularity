package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/data-preservation-programs/go-singularity/util"
	nilrouting "github.com/ipfs/go-ipfs-routing/none"
	bsnetwork "github.com/ipfs/go-libipfs/bitswap/network"
	"github.com/ipfs/go-libipfs/bitswap/server"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/data-preservation-programs/go-singularity/datasource"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/data-preservation-programs/go-singularity/store"
	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func GenerateNewPeer() ([]byte, []byte, peer.ID, error) {
	private, public, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, nil, "", errors.Wrap(err, "cannot generate new peer")
	}

	peerID, err := peer.IDFromPublicKey(public)
	if err != nil {
		return nil, nil, "", errors.Wrap(err, "cannot generate peer id")
	}

	privateBytes, err := crypto.MarshalPrivateKey(private)
	if err != nil {
		return nil, nil, "", errors.Wrap(err, "cannot marshal private key")
	}

	publicBytes, err := crypto.MarshalPublicKey(public)
	if err != nil {
		return nil, nil, "", errors.Wrap(err, "cannot marshal public key")
	}
	return privateBytes, publicBytes, peerID, nil
}

type ContentProviderService struct {
	Resolver datasource.HandlerResolver
	DB       *gorm.DB
	bind     string
	host     host.Host
}

func NewContentProviderService(db *gorm.DB, bind string, identity string, listen []string) (*ContentProviderService, error) {
	var private []byte
	if identity == "" {
		var err error
		private, _, _, err = GenerateNewPeer()
		if err != nil {
			return nil, err
		}
	} else {
		var err error
		private, err = base64.StdEncoding.DecodeString(identity)
		if err != nil {
			return nil, err
		}
	}
	identityKey, err := crypto.UnmarshalPrivateKey(private)
	if err != nil {
		return nil, err
	}
	var listenAddrs []multiaddr.Multiaddr
	for _, addr := range listen {
		ma, err := multiaddr.NewMultiaddr(addr)
		if err != nil {
			return nil, err
		}
		listenAddrs = append(listenAddrs, ma)
	}
	h, err := util.InitHost(context.Background(), []libp2p.Option{libp2p.Identity(identityKey)}, listenAddrs...)
	if err != nil {
		return nil, err
	}
	for _, m := range h.Addrs() {
		logging.Logger("contentprovider").Info("listening on " + m.String())
	}
	logging.Logger("contentprovider").Info("peerID: " + h.ID().String())
	return &ContentProviderService{DB: db, bind: bind, Resolver: datasource.NewDefaultHandlerResolver(), host: h}, nil
}

func (s *ContentProviderService) StartBitswap() error {
	ctx := context.Background()
	nilRouter, err := nilrouting.ConstructNilRouting(ctx, nil, nil, nil)
	if err != nil {
		return err
	}

	net := bsnetwork.NewFromIpfsHost(s.host, nilRouter)
	bs := store.ItemReferenceBlockStore{DB: s.DB, HandlerResolver: datasource.NewDefaultHandlerResolver()}
	bsserver := server.New(ctx, net, bs)
	net.Start(bsserver)
	return nil
}

func (s *ContentProviderService) Start() {
	s.StartBitswap()
	logger := logging.Logger("contentprovider")
	e := echo.New()
	current := logging.GetConfig().Level
	if logging.LevelInfo < current {
		logging.SetAllLoggers(logging.LevelInfo)
	}

	e.Use(
		middleware.RequestLoggerWithConfig(
			middleware.RequestLoggerConfig{
				LogStatus: true,
				LogURI:    true,
				LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
					uri := v.URI
					status := v.Status
					latency := time.Now().Sub(v.StartTime)
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
	e.Use(middleware.Recover())
	e.GET("/piece/:id", s.handleGetPiece)
	e.HEAD("/piece/:id", s.handleHeadPiece)
	e.GET("/ipfs/:cid", s.handleGetCid)

	err := e.Start(s.bind)
	if err != nil {
		panic(err)
	}
}

func (s *ContentProviderService) FindPieceAsPieceReader(ctx context.Context, pieceCid cid.Cid) (
	*store.PieceReader,
	*model.Car,
	error,
) {
	var car model.Car
	err := s.DB.WithContext(ctx).Where("piece_cid = ?", pieceCid.String()).First(&car).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, os.ErrNotExist
	}
	if err != nil {
		return nil, nil, fmt.Errorf("failed to query for CARs: %w", err)
	}

	var carBlocks []model.CarBlock
	err = s.DB.WithContext(ctx).Preload("Source").Preload("Item").Where("car_id = ?", car.ID).
		Find(&carBlocks).Error
	if err != nil {
		return nil, nil, fmt.Errorf("failed to query for CAR items: %w", err)
	}
	reader, err := store.NewPieceReader(ctx, car, carBlocks, s.Resolver)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create piece reader: %w", err)
	}
	return reader, &car, nil
}

func (s *ContentProviderService) findPiece(ctx context.Context, pieceCid cid.Cid) (*os.File, os.FileInfo, error) {
	var cars []model.Car
	err := s.DB.WithContext(ctx).Where("piece_cid = ?", pieceCid.String()).Find(&cars).Error
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

	_, car, err := s.FindPieceAsPieceReader(c.Request().Context(), pieceCid)
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

func (s *ContentProviderService) handleGetCid(c echo.Context) error {
	id := c.Param("cid")
	cid, err := cid.Parse(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "failed to parse CID: "+err.Error())
	}

	var item model.Item
	err = s.DB.WithContext(c.Request().Context()).Preload("Source").Where("cid = ?", cid.String()).First(&item).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "CID not found")
	}
	handler, err := s.Resolver.GetHandler(*item.Source)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to get handler: "+err.Error())
	}

	handle, err := handler.Read(c.Request().Context(), item.Path, item.Offset, item.Length)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to open handler: "+err.Error())
	}
	defer handle.Close()
	return c.Stream(http.StatusOK, "application/octet-stream", handle)
}

func (s *ContentProviderService) handleGetPiece(c echo.Context) error {
	id := c.Param("id")
	pieceCid, err := cid.Parse(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "failed to parse piece CID: "+err.Error())
	}

	var reader io.ReaderAt
	var pieceReader *store.PieceReader
	var lastModified time.Time
	var fileSize int64
	fil, fInfo, err := s.findPiece(c.Request().Context(), pieceCid)
	if err == nil {
		defer fil.Close()
		reader = fil
		lastModified = fInfo.ModTime()
		fileSize = fInfo.Size()
	} else {
		var car *model.Car
		pieceReader, car, err = s.FindPieceAsPieceReader(c.Request().Context(), pieceCid)
		switch {
		case err == nil:
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
		if reader != nil {
			http.ServeContent(
				c.Response(),
				c.Request(),
				pieceCid.String()+".car",
				lastModified,
				io.NewSectionReader(reader, 0, fileSize),
			)
		} else {
			c.Response().Header().Set("Last-Modified", lastModified.UTC().Format(http.TimeFormat))
			_, err := io.Copy(c.Response().Writer, pieceReader)
			if err != nil {
				return c.String(http.StatusInternalServerError, "failed to copy piece reader: "+err.Error())
			}
		}
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
	if reader != nil {
		http.ServeContent(
			c.Response(),
			c.Request(),
			pieceCid.String()+".car",
			lastModified,
			io.NewSectionReader(reader, start, end-start+1),
		)
	} else {
		_, err := io.CopyN(c.Response().Writer, pieceReader, end-start+1)
		if err != nil {
			return c.String(http.StatusInternalServerError, "failed to copy piece reader: "+err.Error())
		}
	}
	return nil
}
