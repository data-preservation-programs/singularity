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

// NewService creates a new Service instance with the provided database and configuration.
//
// The NewService function takes the following parameters:
// - db: The gorm.DB instance for database operations.
// - config: The Config struct containing the service configuration.
//
// The function performs the following steps:
// 1. Creates an empty Service instance.
// 2. If the HTTP server is enabled in the configuration, creates an HTTPServer instance and adds it to the servers slice.
//   - The HTTPServer is configured with the bind address, database without context, and a DefaultHandlerResolver.
//
// 3. If the Bitswap server is enabled in the configuration, initializes the identity key based on the configuration.
//   - If the identity key is not provided, generates a new peer identity key.
//   - If the identity key is provided, decodes it from base64.
//   - Unmarshals the private key from the identity key bytes.
//   - If no listen multiaddresses are provided, sets a default listen multiaddress.
//   - Converts each listen multiaddress string to a Multiaddr instance.
//   - Initializes a libp2p host with the identity key and listen multiaddresses.
//   - Logs the libp2p listening addresses and peer ID.
//   - Creates a BitswapServer instance with the libp2p host and database without context, and adds it to the servers slice.
//
// 4. Returns the created Service instance and nil for the error if all steps are executed successfully.
func NewService(db *gorm.DB, config Config) (*Service, error) {
	s := &Service{}

	if config.HTTP.Enable {
		s.servers = append(s.servers, &HTTPServer{
			bind:        config.HTTP.Bind,
			dbNoContext: db,
			resolver:    &storagesystem.DefaultHandlerResolver{},
		})
	}

	if config.Bitswap.Enable {
		var private []byte
		if config.Bitswap.IdentityKey == "" {
			var err error
			private, _, _, err = util.GenerateNewPeer()
			if err != nil {
				return nil, errors.WithStack(err)
			}
		} else {
			var err error
			private, err = base64.StdEncoding.DecodeString(config.Bitswap.IdentityKey)
			if err != nil {
				return nil, errors.WithStack(err)
			}
		}
		identityKey, err := crypto.UnmarshalPrivateKey(private)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if len(config.Bitswap.ListenMultiAddrs) == 0 {
			config.Bitswap.ListenMultiAddrs = []string{"/ip4/0.0.0.0/tcp/0"}
		}
		var listenAddrs []multiaddr.Multiaddr
		for _, addr := range config.Bitswap.ListenMultiAddrs {
			ma, err := multiaddr.NewMultiaddr(addr)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			listenAddrs = append(listenAddrs, ma)
		}
		h, err := util.InitHost([]libp2p.Option{libp2p.Identity(identityKey)}, listenAddrs...)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		for _, m := range h.Addrs() {
			logger.Info("libp2p listening on " + m.String())
		}
		logger.Info("peerID: " + h.ID().String())
		s.servers = append(s.servers, &BitswapServer{
			host:        h,
			dbNoContext: db,
		})
	}
	return s, nil
}

func (s *Service) Start(ctx context.Context) error {
	return service.StartServers(ctx, logger, s.servers...)
}
