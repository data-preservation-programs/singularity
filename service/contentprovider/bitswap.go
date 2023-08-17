package contentprovider

import (
	"context"

	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/data-preservation-programs/singularity/store"
	nilrouting "github.com/ipfs/go-ipfs-routing/none"
	bsnetwork "github.com/ipfs/go-libipfs/bitswap/network"
	"github.com/ipfs/go-libipfs/bitswap/server"
	"github.com/libp2p/go-libp2p/core/host"
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

func (BitswapServer) Name() string {
	return "Bitswap"
}

// Start initializes the Bitswap server with the provided context.
// It sets up the necessary routing and networking components,
// and starts serving Bitswap requests.
// It returns channels that signal when the service has stopped or encountered an error.
func (s BitswapServer) Start(ctx context.Context) ([]service.Done, service.Fail, error) {
	nilRouter, err := nilrouting.ConstructNilRouting(ctx, nil, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	net := bsnetwork.NewFromIpfsHost(s.host, nilRouter)
	bs := &store.FileReferenceBlockStore{DBNoContext: s.dbNoContext, HandlerResolver: datasource.DefaultHandlerResolver{}}
	bsserver := server.New(ctx, net, bs)
	net.Start(bsserver)
	done := make(chan struct{})
	fail := make(chan error)

	go func() {
		<-ctx.Done()
		net.Stop()
		bsserver.Close()
		close(done)
	}()
	return []service.Done{done}, fail, nil
}
