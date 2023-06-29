package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/data-preservation-programs/singularity/util"
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
	"time"

	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/store"
	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var logger = logging.Logger("contentprovider")

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

type ContentProviderConfig struct {
	EnableHTTP        bool
	HTTPBind          string
	EnableBitswap     bool
	Libp2pIdentityKey string
	Libp2pListenAddrs []string
}

func NewContentProviderService(db *gorm.DB, config ContentProviderConfig) (*ContentProviderService, error) {
	bind := ""
	if config.EnableHTTP {
		bind = config.HTTPBind
	}
	if config.EnableBitswap {
		var private []byte
		if config.Libp2pIdentityKey == "" {
			var err error
			private, _, _, err = GenerateNewPeer()
			if err != nil {
				return nil, err
			}
		} else {
			var err error
			private, err = base64.StdEncoding.DecodeString(config.Libp2pIdentityKey)
			if err != nil {
				return nil, err
			}
		}
		identityKey, err := crypto.UnmarshalPrivateKey(private)
		if err != nil {
			return nil, err
		}
		var listenAddrs []multiaddr.Multiaddr
		for _, addr := range config.Libp2pListenAddrs {
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
		return &ContentProviderService{DB: db, bind: bind, Resolver: datasource.DefaultHandlerResolver{}, host: h}, nil
	}
	return &ContentProviderService{DB: db, bind: bind, Resolver: datasource.DefaultHandlerResolver{}}, nil
}

func (s *ContentProviderService) StartBitswap(ctx context.Context) error {
	nilRouter, err := nilrouting.ConstructNilRouting(ctx, nil, nil, nil)
	if err != nil {
		return err
	}

	net := bsnetwork.NewFromIpfsHost(s.host, nilRouter)
	bs := store.ItemReferenceBlockStore{DB: s.DB, HandlerResolver: datasource.DefaultHandlerResolver{}}
	bsserver := server.New(ctx, net, bs)
	net.Start(bsserver)
	return nil
}

func (s *ContentProviderService) Start(ctx context.Context) {
	if s.host != nil {
		err := s.StartBitswap(ctx)
		if err != nil {
			logging.Logger("contentprovider").Fatal(err)
		}
	}
	logger := logging.Logger("contentprovider")
	if s.bind != "" {
		e := echo.New()
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
			StackSize:         4 << 10, // 4 KB
			DisableStackAll:   false,
			DisablePrintStack: false,
			LogLevel:          0,
			LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
				logger.Errorw("panic", "err", err, "stack", string(stack))
				return nil
			},
		}))
		e.GET("/piece/:id", s.handleGetPiece)
		e.HEAD("/piece/:id", s.handleHeadPiece)
		e.GET("/ipfs/:cid", s.handleGetCid)
		err := e.Start(s.bind)
		if err != nil {
			panic(err)
		}
	}

	<-ctx.Done()
}

func (s *ContentProviderService) headPiece(ctx context.Context, pieceCid cid.Cid) (int64, error) {
	var cars []model.Car
	err := s.DB.WithContext(ctx).Where("piece_cid = ?", model.CID(pieceCid)).Find(&cars).Error
	if err != nil {
		return 0, errors.Wrap(err, "failed to query for CARs")
	}

	if len(cars) == 0 {
		return 0, os.ErrNotExist
	}

	return cars[0].FileSize, nil
}

func (s *ContentProviderService) FindPiece(ctx context.Context, pieceCid cid.Cid) (
	io.ReadSeekCloser,
	time.Time,
	error,
) {
	var cars []model.Car
	err := s.DB.WithContext(ctx).Where("piece_cid = ?", model.CID(pieceCid)).Find(&cars).Error
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("failed to query for CARs: %w", err)
	}

	if len(cars) == 0 {
		return nil, time.Time{}, os.ErrNotExist
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
		return file, fileInfo.ModTime(), nil
	}

	car := cars[0]
	var source model.Source
	err = s.DB.WithContext(ctx).Where("id = ?", car.SourceID).Find(&source).Error
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("failed to query for source: %w", err)
	}
	var carBlocks []model.CarBlock
	err = s.DB.WithContext(ctx).Where("car_id = ?", car.ID).
		Find(&carBlocks).Error
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("failed to query for CAR blocks: %w", err)
	}
	itemIDSet := make(map[uint64]struct{})
	for _, carBlock := range carBlocks {
		if carBlock.ItemID != nil {
			itemIDSet[*carBlock.ItemID] = struct{}{}
		}
	}
	var itemIDs []uint64
	for itemID := range itemIDSet {
		itemIDs = append(itemIDs, itemID)
	}
	var items []model.Item
	err = s.DB.WithContext(ctx).Where("id IN ?", itemIDs).Find(&items).Error
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("failed to query for items: %w", err)
	}
	reader, err := store.NewPieceReader(ctx, car, source, carBlocks, items, s.Resolver)
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("failed to create piece reader: %w", err)
	}
	return reader, car.CreatedAt, nil
}

func (s *ContentProviderService) setCommonHeaders(c echo.Context, pieceCid string) {
	c.Response().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", pieceCid+".car"))
	c.Response().Header().Set("Content-Type", "application/vnd.ipld.car; version=1")
	c.Response().Header().Set("Accept-Ranges", "bytes")
	c.Response().Header().Set("Etag", "\""+pieceCid+"\"")
}

func (s *ContentProviderService) handleHeadPiece(c echo.Context) error {
	id := c.Param("id")
	pieceCid, err := cid.Parse(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "failed to parse piece CID: "+err.Error())
	}

	size, err := s.headPiece(c.Request().Context(), pieceCid)
	if os.IsNotExist(err) {
		return c.String(http.StatusNotFound, "piece not found")
	}
	if err != nil {
		logger.Errorw("failed to find piece", "pieceCid", pieceCid, "error", err)
		return c.String(http.StatusInternalServerError, "failed to find piece: "+err.Error())
	}

	s.setCommonHeaders(c, pieceCid.String())
	c.Response().Header().Set("Content-Length", strconv.FormatInt(size, 10))
	return c.NoContent(http.StatusOK)
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
	handler, err := s.Resolver.Resolve(c.Request().Context(), *item.Source)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to get handler: "+err.Error())
	}

	handle, _, err := handler.Read(c.Request().Context(), item.Path, 0, item.Size)
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

	reader, lastModified, err := s.FindPiece(c.Request().Context(), pieceCid)
	if os.IsNotExist(err) {
		return c.String(http.StatusNotFound, "piece not found")
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to find piece: "+err.Error())
	}

	defer reader.Close()
	s.setCommonHeaders(c, pieceCid.String())
	http.ServeContent(
		c.Response(),
		c.Request(),
		pieceCid.String()+".car",
		lastModified,
		reader,
	)

	return nil
}
