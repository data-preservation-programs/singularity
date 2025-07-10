package contentprovider

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/store"
	"github.com/data-preservation-programs/singularity/util"
	bsnetwork "github.com/ipfs/boxo/bitswap/network"
	"github.com/ipfs/boxo/bitswap/server"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/multiformats/go-multiaddr"
	"gorm.io/gorm"
)

// BitswapServer represents a server instance for handling Bitswap protocol interactions.
// Bitswap is a peer-to-peer data trading protocol in which peers request the data they need,
// and respond to other peers' requests based on certain policies.
type BitswapServer struct {
	// dbNoContext is a GORM database instance that doesn't use context for managing database connections.
	dbNoContext *gorm.DB

	// host is a libp2p host used to build and configure a new Bitswap instance.
	host host.Host
}

func NewBitswapServer(dbNoContext *gorm.DB, private crypto.PrivKey, addrs ...multiaddr.Multiaddr) (*BitswapServer, error) {
	h, err := util.InitHost([]libp2p.Option{libp2p.Identity(private)}, addrs...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for _, m := range h.Addrs() {
		logger.Info("libp2p listening on " + m.String())
	}
	logger.Info("peerID: " + h.ID().String())
	return &BitswapServer{
		dbNoContext: dbNoContext,
		host:        h,
	}, nil
}

func (BitswapServer) Name() string {
	return "Bitswap"
}

// Start initializes the Bitswap server with the provided context.
// It sets up the necessary routing and networking components,
// and starts serving Bitswap requests.
// It returns channels that signal when the service has stopped or encountered an error.
func (s BitswapServer) Start(ctx context.Context, exitErr chan<- error) error {
	// Updated for boxo v0.26.0 - NewFromIpfsHost no longer takes routing as second parameter
	net := bsnetwork.NewFromIpfsHost(s.host)
	bs := &store.FileReferenceBlockStore{DBNoContext: s.dbNoContext}
	bsserver := server.New(ctx, net, bs)
	net.Start(bsserver)

	go func() {
		<-ctx.Done()
		net.Stop()
		// Updated for boxo v0.26.0 - Close() no longer returns an error
		bsserver.Close()
		_ = s.host.Close()
		if exitErr != nil {
			exitErr <- nil
		}
	}()
	return nil
}
