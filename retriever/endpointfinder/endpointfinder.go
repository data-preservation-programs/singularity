package endpointfinder

import (
	"context"
	"errors"
	"fmt"

	"github.com/data-preservation-programs/singularity/replication"
	"github.com/filecoin-shipyard/boostly"
	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"go.uber.org/multierr"
)

var logger = log.Logger("singularity/retriever/endpointfinder")

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
	httpEndpoints    *expirable.LRU[string, []peer.AddrInfo]
	endpointErrors   *expirable.LRU[string, error]
}

// NewEndpointFinder returns a new instance of an EndpointFinder
func NewEndpointFinder(minerInfoFetcher MinerInfoFetcher, h host.Host, opts ...Option) *EndpointFinder {
	cfg := applyOptions(opts...)

	httpEndpoints := expirable.NewLRU[string, []peer.AddrInfo](cfg.LruSize, func(key string, value []peer.AddrInfo) {}, cfg.LruTimeout)
	endpointErrors := expirable.NewLRU[string, error](cfg.ErrorLruSize, func(key string, value error) {}, cfg.ErrorLruTimeout)

	return &EndpointFinder{
		minerInfoFetcher: minerInfoFetcher,
		h:                h,
		httpEndpoints:    httpEndpoints,
		endpointErrors:   endpointErrors,
	}
}

func (ef *EndpointFinder) findHTTPEndpointsForProvider(ctx context.Context, provider string) ([]peer.AddrInfo, error) {
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

// FindHTTPEndpoints finds http endpoints for a given set of providers
func (ef *EndpointFinder) FindHTTPEndpoints(ctx context.Context, sps []string) ([]peer.AddrInfo, error) {
	addrInfos := make([]peer.AddrInfo, 0, len(sps))
	type findResult struct {
		addrs []peer.AddrInfo
		err   error
	}
	addrChan := make(chan findResult)
	var toLookup int
	var errsum error

	for _, provider := range sps {
		// first check our caches
		if providerAddrs, has := ef.httpEndpoints.Get(provider); has {
			addrInfos = append(addrInfos, providerAddrs...)
		} else if err, has := ef.endpointErrors.Get(provider); has {
			logger.Errorf("error looking up http endpoint for %s (cached): %s", provider, err)
			errsum = multierr.Append(errsum, err)
		} else {
			// not in caches, perform full lookup of provider asynchronously
			toLookup++
			go func(provider string) {
				providerAddrs, err := ef.findHTTPEndpointsForProvider(ctx, provider)
				if err != nil {
					ef.endpointErrors.Add(provider, err)
					logger.Errorf("error looking up http endpoint for %s: %s", provider, err)
				} else {
					ef.httpEndpoints.Add(provider, providerAddrs)
				}
				select {
				case addrChan <- findResult{addrs: providerAddrs, err: err}:
				case <-ctx.Done():
				}
			}(provider)
		}
	}

	for i := 0; i < toLookup; i++ {
		select {
		case providerAddrs := <-addrChan:
			if providerAddrs.addrs != nil {
				addrInfos = append(addrInfos, providerAddrs.addrs...)
			} else if providerAddrs.err != nil {
				errsum = multierr.Append(errsum, providerAddrs.err)
			}
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	if len(addrInfos) == 0 {
		return nil, fmt.Errorf("no http endpoints found for providers %v: %w", sps, errsum)
	}
	return addrInfos, nil
}
