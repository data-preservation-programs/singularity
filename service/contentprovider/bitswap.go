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

type BitswapServer struct {
	db   *gorm.DB
	host host.Host
}

func (BitswapServer) Name() string {
	return "Bitswap"
}

func (s BitswapServer) Start(ctx context.Context) (service.Done, service.Fail, error) {
	nilRouter, err := nilrouting.ConstructNilRouting(ctx, nil, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	net := bsnetwork.NewFromIpfsHost(s.host, nilRouter)
	bs := &store.ItemReferenceBlockStore{DB: s.db, HandlerResolver: datasource.DefaultHandlerResolver{}}
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
	return done, fail, nil
}
