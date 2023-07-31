package contentprovider

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/data-preservation-programs/singularity/util"
	"github.com/fxamacker/cbor/v2"
	nilrouting "github.com/ipfs/go-ipfs-routing/none"
	bsnetwork "github.com/ipfs/go-libipfs/bitswap/network"
	"github.com/ipfs/go-libipfs/bitswap/server"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"

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
		h, err := util.InitHost([]libp2p.Option{libp2p.Identity(identityKey)}, listenAddrs...)
		if err != nil {
			return nil, err
		}
		for _, m := range h.Addrs() {
			logger.Info("listening on " + m.String())
		}
		logger.Info("peerID: " + h.ID().String())
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

func (s *ContentProviderService) Start(ctx context.Context) error {
	if s.host != nil {
		err := s.StartBitswap(ctx)
		if err != nil {
			logger.Fatal(err)
		}
	}
	httpDone := make(chan struct{})
	if s.bind != "" {
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
			StackSize:         4 << 10, // 4 KB
			DisableStackAll:   false,
			DisablePrintStack: false,
			LogLevel:          0,
			LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
				logger.Errorw("panic", "err", err, "stack", string(stack))
				return nil
			},
		}))
		e.GET("/piece/metadata/:id", s.GetMetadataHandler)
		e.HEAD("/piece/metadata/:id", s.GetMetadataHandler)
		e.GET("/piece/:id", s.handleGetPiece)
		e.HEAD("/piece/:id", s.handleGetPiece)
		e.GET("/ipfs/:cid", s.handleGetCid)

		go func() {
			<-ctx.Done()
			logger.Warnw("shutting down the server")
			shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer shutdownCancel()
			//nolint:contextcheck
			if err := e.Shutdown(shutdownCtx); err != nil {
				fmt.Printf("Error shutting down the server: %v\n", err)
			}
			httpDone <- struct{}{}
		}()

		defer func() {
			<-httpDone
		}()

		err := e.Start(s.bind)
		if err != nil {
			return err
		}
	}

	<-ctx.Done()
	return nil
}

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

	metadata, err := GetPieceMetadata(ctx, db, car)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
	}

	// Remove all relevant credentials
	metadata.Source.Metadata = nil

	acceptHeader := c.Request().Header.Get("Accept")
	switch acceptHeader {
	case "application/cbor":
		c.Response().WriteHeader(http.StatusOK)
		c.Response().Header().Set("Content-Type", "application/cbor")
		encoder := cbor.NewEncoder(c.Response().Writer)
		return encoder.Encode(metadata)
	default:
		return c.JSON(http.StatusOK, metadata)
	}
}

func (s *ContentProviderService) GetMetadataHandler(c echo.Context) error {
	return GetMetadataHandler(c, s.DB)
}

type PieceMetadata struct {
	Car       model.Car        `json:"car"`
	Source    model.Source     `json:"source"`
	CarBlocks []model.CarBlock `json:"carBlocks"`
	Items     []model.Item     `json:"items"`
}

func GetPieceMetadata(ctx context.Context, db *gorm.DB, car model.Car) (*PieceMetadata, error) {
	var source model.Source
	err := db.WithContext(ctx).Where("id = ?", car.SourceID).Find(&source).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query for source: %w", err)
	}
	var carBlocks []model.CarBlock
	err = db.WithContext(ctx).Where("car_id = ?", car.ID).
		Find(&carBlocks).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query for CAR blocks: %w", err)
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
	err = db.WithContext(ctx).Where("id IN ?", itemIDs).Find(&items).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query for items: %w", err)
	}
	return &PieceMetadata{
		Car:       car,
		Source:    source,
		CarBlocks: carBlocks,
		Items:     items,
	}, nil
}

func (s *ContentProviderService) GetPieceMetadata(ctx context.Context, car model.Car) (*PieceMetadata, error) {
	return GetPieceMetadata(ctx, s.DB, car)
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
	metadata, err := s.GetPieceMetadata(ctx, car)
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("failed to get piece metadata: %w", err)
	}
	reader, err := store.NewPieceReader(ctx, metadata.Car, metadata.Source, metadata.CarBlocks, metadata.Items, s.Resolver)
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
