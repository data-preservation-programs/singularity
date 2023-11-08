package endpointfinder

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/data-preservation-programs/singularity/replication"
	"github.com/filecoin-shipyard/boostly"
	lru "github.com/hashicorp/golang-lru/v2"
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
	httpEndpoints    *lru.Cache[string, []peer.AddrInfo]
	endpointErrors   *expirable.LRU[string, error]
}

const (
	defaultLruSize         = 128
	defaultErrorLruSize    = 128
	defaultErrorLruTimeout = 5 * time.Minute
)

type EndpointFinderConfig struct {
	LruSize         int
	ErrorLruSize    int
	ErrorLruTimeout time.Duration
}

func (cfg EndpointFinderConfig) lruSize() int {
	if cfg.LruSize == 0 {
		return defaultLruSize
	}
	return cfg.LruSize
}

func (cfg EndpointFinderConfig) errorLruSize() int {
	if cfg.ErrorLruSize == 0 {
		return defaultErrorLruSize
	}
	return cfg.ErrorLruSize
}

func (cfg EndpointFinderConfig) errorLruTimeout() time.Duration {
	if cfg.ErrorLruTimeout == 0 {
		return defaultErrorLruTimeout
	}
	return cfg.ErrorLruTimeout
}

// NewEndpointFinder returns a new instance of an EndpointFinder
func NewEndpointFinder(cfg EndpointFinderConfig, minerInfoFetcher MinerInfoFetcher, h host.Host) (*EndpointFinder, error) {
	httpEndpoints, err := lru.New[string, []peer.AddrInfo](cfg.lruSize())
	if err != nil {
		return nil, err
	}
	endpointErrors := expirable.NewLRU[string, error](cfg.errorLruSize(), func(key string, value error) {}, cfg.errorLruTimeout())
	return &EndpointFinder{
		minerInfoFetcher: minerInfoFetcher,
		h:                h,
		httpEndpoints:    httpEndpoints,
		endpointErrors:   endpointErrors,
	}, nil
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

	for _, provider := range sps {
		go func(provider string) {
			var result findResult
			// first check our caches
			if providerAddrs, has := ef.httpEndpoints.Get(provider); has {
				result.addrs = providerAddrs
			} else if err, has := ef.endpointErrors.Get(provider); has {
				logger.Errorf("error looking up http endpoint for %s (cached): %s", provider, err)
				result.err = err
			} else {
				// not in caches, perform full lookup of provider
				providerAddrs, err := ef.findHTTPEndpointsForProvider(ctx, provider)
				result = findResult{addrs: providerAddrs, err: err}
				if err != nil {
					ef.endpointErrors.Add(provider, err)
					logger.Errorf("error looking up http endpoint for %s: %s", provider, err)
				} else {
					ef.httpEndpoints.Add(provider, providerAddrs)
				}
			}
			select {
			case addrChan <- result:
			case <-ctx.Done():
			}
		}(provider)
	}

	var err error
	for range sps {
		select {
		case providerAddrs := <-addrChan:
			if providerAddrs.addrs != nil {
				addrInfos = append(addrInfos, providerAddrs.addrs...)
			} else if providerAddrs.err != nil {
				err = multierr.Append(err, providerAddrs.err)
			}
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	if len(addrInfos) == 0 {
		return nil, fmt.Errorf("no http endpoints found for providers %v: %w", sps, err)
	}
	return addrInfos, nil
}
