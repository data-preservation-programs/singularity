package endpointfinder

import (
	"context"
	"fmt"

	"github.com/data-preservation-programs/singularity/replication"
	"github.com/filecoin-shipyard/boostly"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
)

type MinerInfoFetcher interface {
	GetProviderInfo(ctx context.Context, provider string) (*replication.MinerInfo, error)
}

type EndpointFinder struct {
	minerInfoFetcher MinerInfoFetcher
	h                host.Host
	httpEndpoints    *lru.Cache[string, peer.AddrInfo]
}

func NewEndpointFinder(minerInfoFetcher MinerInfoFetcher, h host.Host) (*EndpointFinder, error) {
	httpEndpoints, err := lru.New[string, peer.AddrInfo](128)
	if err != nil {
		return nil, err
	}
	return &EndpointFinder{
		minerInfoFetcher: minerInfoFetcher,
		h:                h,
		httpEndpoints:    httpEndpoints,
	}, nil
}

func (ef *EndpointFinder) fetchHTTPEndpoint(ctx context.Context, provider string) (peer.AddrInfo, error) {
	minerInfo, err := ef.minerInfoFetcher.GetProviderInfo(ctx, provider)
	if err != nil {
		return peer.AddrInfo{}, fmt.Errorf("looking up provider info: %w", err)
	}
	ef.h.Peerstore().AddAddrs(minerInfo.PeerID, minerInfo.Multiaddrs, peerstore.TempAddrTTL)
	response, err := boostly.QueryTransports(ctx, ef.h, minerInfo.PeerID)
	if err != nil {
		return peer.AddrInfo{}, fmt.Errorf("querying transports: %w", err)
	}
	for _, protocol := range response.Protocols {
		if protocol.Name == "http" {
			return peer.AddrInfo{
				ID:    minerInfo.PeerID,
				Addrs: protocol.Addresses,
			}, nil
		}
	}
	return peer.AddrInfo{}, fmt.Errorf("provider does not support http: %w", err)
}

func (ef *EndpointFinder) findOrFetchHTTPEndpoint(ctx context.Context, provider string) (peer.AddrInfo, error) {
	addrInfo, has := ef.httpEndpoints.Get(provider)
	if has {
		return addrInfo, nil
	}
	addrInfo, err := ef.fetchHTTPEndpoint(ctx, provider)
	if err != nil {
		return peer.AddrInfo{}, err
	}
	ef.httpEndpoints.Add(provider, addrInfo)
	return addrInfo, nil
}

func (ef *EndpointFinder) FindHTTPEndpoints(ctx context.Context, sps []string) ([]peer.AddrInfo, error) {
	addrInfos := make([]peer.AddrInfo, 0, len(sps))
	for _, sp := range sps {
		// TODO: should we ignore if some but not all providers are configured correctly?
		addrInfo, err := ef.findOrFetchHTTPEndpoint(ctx, sp)
		if err != nil {
			return nil, err
		}
		addrInfos = append(addrInfos, addrInfo)
	}
	return addrInfos, nil
}
