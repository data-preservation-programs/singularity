package endpointfinder

import (
	"context"
	"errors"
	"fmt"

	"github.com/data-preservation-programs/singularity/replication"
	"github.com/filecoin-shipyard/boostly"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
)

// ErrHTTPNotSupported indicates we were able to look up the provider and contact them,
// but they reported that they do not serve HTTP retrievals
var ErrHTTPNotSupported = errors.New("provider does not support http")

// MinerInfoFetcher is an interface for looking up chain miner info from
// an SP name
type MinerInfoFetcher interface {
	GetProviderInfo(ctx context.Context, provider string) (*replication.MinerInfo, error)
}

// EndpointFinder handles translating SP miner addresses to HTTP retrieval
// endpoints for those miners, use the chain and the SP retrieval transports
// protocol. It also caches records for performance
type EndpointFinder struct {
	minerInfoFetcher MinerInfoFetcher
	h                host.Host
	httpEndpoints    *lru.Cache[string, []peer.AddrInfo]
}

// NewEndpointFinder returns a new instance of an EndpointFinder
func NewEndpointFinder(minerInfoFetcher MinerInfoFetcher, h host.Host, size int) (*EndpointFinder, error) {
	httpEndpoints, err := lru.New[string, []peer.AddrInfo](size)
	if err != nil {
		return nil, err
	}
	return &EndpointFinder{
		minerInfoFetcher: minerInfoFetcher,
		h:                h,
		httpEndpoints:    httpEndpoints,
	}, nil
}

func (ef *EndpointFinder) fetchHTTPEndpoint(ctx context.Context, provider string) ([]peer.AddrInfo, error) {
	// lookup the provider on chain
	minerInfo, err := ef.minerInfoFetcher.GetProviderInfo(ctx, provider)
	if err != nil {
		return nil, fmt.Errorf("looking up provider info: %w", err)
	}
	// query provider for supported transports
	ef.h.Peerstore().AddAddrs(minerInfo.PeerID, minerInfo.Multiaddrs, peerstore.TempAddrTTL)
	response, err := boostly.QueryTransports(ctx, ef.h, minerInfo.PeerID)
	if err != nil {
		return nil, fmt.Errorf("querying transports: %w", err)
	}
	// filter supported transports to get http endpoints
	for _, protocol := range response.Protocols {
		if protocol.Name == "http" {
			addrs, err := peer.AddrInfosFromP2pAddrs(protocol.Addresses...)
			// if no peer id is present, use provider's id
			if err != nil {
				addrs = []peer.AddrInfo{{
					ID:    minerInfo.PeerID,
					Addrs: protocol.Addresses,
				}}
			}
			return addrs, nil
		}
	}
	return nil, ErrHTTPNotSupported
}

// findOrFetchHTTPEndpoint attempts to load from cache before calling fetchHTTPEndpoint
func (ef *EndpointFinder) findOrFetchHTTPEndpoint(ctx context.Context, provider string) ([]peer.AddrInfo, error) {
	addrInfos, has := ef.httpEndpoints.Get(provider)
	if has {
		return addrInfos, nil
	}
	addrInfos, err := ef.fetchHTTPEndpoint(ctx, provider)
	if err != nil {
		return nil, err
	}
	ef.httpEndpoints.Add(provider, addrInfos)
	return addrInfos, nil
}

// FindHTTPEndpoints finds http endpoints for a given set of providers
func (ef *EndpointFinder) FindHTTPEndpoints(ctx context.Context, sps []string) ([]peer.AddrInfo, error) {
	addrInfos := make([]peer.AddrInfo, 0, len(sps))
	for _, sp := range sps {
		// TODO: should we ignore if some but not all providers are configured correctly?
		nextAddrInfos, err := ef.findOrFetchHTTPEndpoint(ctx, sp)
		if err != nil {
			return nil, err
		}
		addrInfos = append(addrInfos, nextAddrInfos...)
	}
	return addrInfos, nil
}
