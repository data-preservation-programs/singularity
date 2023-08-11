package contentprovider

import (
	"context"
	"encoding/base64"

	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/multiformats/go-multiaddr"

	logging "github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

var logger = logging.Logger("contentprovider")

type Service struct {
	servers []service.Server
}

type Config struct {
	HTTP    HTTPConfig
	Bitswap BitswapConfig
}

type HTTPConfig struct {
	Enable bool
	Bind   string
}

type BitswapConfig struct {
	Enable           bool
	IdentityKey      string
	ListenMultiAddrs []string
}

func NewService(db *gorm.DB, config Config) (*Service, error) {
	s := &Service{}

	if config.HTTP.Enable {
		s.servers = append(s.servers, &HTTPServer{
			bind:     config.HTTP.Bind,
			db:       db,
			resolver: &datasource.DefaultHandlerResolver{},
		})
	}

	if config.Bitswap.Enable {
		var private []byte
		if config.Bitswap.IdentityKey == "" {
			var err error
			private, _, _, err = util.GenerateNewPeer()
			if err != nil {
				return nil, err
			}
		} else {
			var err error
			private, err = base64.StdEncoding.DecodeString(config.Bitswap.IdentityKey)
			if err != nil {
				return nil, err
			}
		}
		identityKey, err := crypto.UnmarshalPrivateKey(private)
		if err != nil {
			return nil, err
		}
		if len(config.Bitswap.ListenMultiAddrs) == 0 {
			config.Bitswap.ListenMultiAddrs = []string{"/ip4/0.0.0.0/tcp/0"}
		}
		var listenAddrs []multiaddr.Multiaddr
		for _, addr := range config.Bitswap.ListenMultiAddrs {
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
			logger.Info("libp2p listening on " + m.String())
		}
		logger.Info("peerID: " + h.ID().String())
		s.servers = append(s.servers, &BitswapServer{
			host: h,
			db:   db,
		})
	}
	return s, nil
}

func (s *Service) Start(ctx context.Context) error {
	return service.StartServers(ctx, logger, s.servers...)
}
