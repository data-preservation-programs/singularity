package util

import (
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/p2p/muxer/yamux"
	"github.com/libp2p/go-libp2p/p2p/security/noise"
	tls "github.com/libp2p/go-libp2p/p2p/security/tls"
	quic "github.com/libp2p/go-libp2p/p2p/transport/quic"
	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
	"github.com/libp2p/go-libp2p/p2p/transport/websocket"
	webtransport "github.com/libp2p/go-libp2p/p2p/transport/webtransport"
	"github.com/multiformats/go-multiaddr"
)

const yamuxID = "/yamux/1.0.0"

// InitHost initializes a new libp2p host with the provided options and listen addresses.
//
// Parameters:
//   - opts: A slice of libp2p options to configure the host.
//   - listenAddrs: A variadic list of multiaddresses for the host to listen on.
//
// Returns:
//  1. A libp2p host instance that represents the initialized host.
//  2. An error, which will be non-nil if any error occurred during host initialization.
func InitHost(opts []libp2p.Option, listenAddrs ...multiaddr.Multiaddr) (host.Host, error) {
	opts = append([]libp2p.Option{
		libp2p.Identity(nil),
		libp2p.ResourceManager(&network.NullResourceManager{}),
	},
		opts...)
	if len(listenAddrs) > 0 {
		opts = append([]libp2p.Option{libp2p.ListenAddrs(listenAddrs...)}, opts...)
	}
	// add transports
	opts = append([]libp2p.Option{
		libp2p.Transport(tcp.NewTCPTransport, tcp.WithMetrics()),
		libp2p.Transport(websocket.New),
		libp2p.Transport(quic.NewTransport),
		libp2p.Transport(webtransport.New),
	},
		opts...)
	// add security
	opts = append([]libp2p.Option{
		libp2p.Security(tls.ID, tls.New),
		libp2p.Security(noise.ID, noise.New),
	},
		opts...)

	// add muxers
	opts = append([]libp2p.Option{
		libp2p.Muxer(yamuxID, yamuxTransport()),
	},
		opts...)

	//nolint:wrapcheck
	return libp2p.New(opts...)
}

func yamuxTransport() network.Multiplexer {
	tpt := *yamux.DefaultTransport
	tpt.AcceptBacklog = 512
	return &tpt
}
