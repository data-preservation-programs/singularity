package endpointfinder_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/data-preservation-programs/singularity/replication"
	"github.com/data-preservation-programs/singularity/retriever/endpointfinder"
	"github.com/filecoin-shipyard/boostly"
	"github.com/ipfs/go-log/v2"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	"github.com/ipld/go-ipld-prime/node/bindnode/registry"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
	"github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/require"
)

func TestEndpointFetcher(t *testing.T) {
	// Suppress error logs during testing to avoid confusing output.
	// These tests intentionally trigger error conditions that generate error logs,
	// but the errors are expected and tested for, so we suppress them to keep
	// test output clean and avoid confusion in CI environments.
	log.SetLogLevel("singularity/retriever/endpointfinder", "fatal")
	defer func() {
		log.SetLogLevel("singularity/retriever/endpointfinder", "info")
	}()

	testCases := []struct {
		testName                 string
		providers                int
		minerInfoNotFindable     bool
		notDialable              bool
		notListeningOnTransports bool
		noHTTP                   bool
		expectedErrString        string
	}{
		{testName: "success path"},
		{
			testName:             "unable to find miner on chain",
			minerInfoNotFindable: true,
			expectedErrString:    "no http endpoints found for providers [%s]: looking up provider info: miner not found",
		},
		{
			testName:          "unable to dial provider",
			notDialable:       true,
			expectedErrString: "no http endpoints found for providers [%s]: querying transports: failed to dial: %s cannot connect to %s",
		},
		{
			testName:                 "provider not listening on protocol",
			notListeningOnTransports: true,
			expectedErrString:        "no http endpoints found for providers [%s]: querying transports: failed to negotiate protocol: protocols not supported: [/fil/retrieval/transports/1.0.0]",
		},
		{
			testName:          "provider not serving http",
			noHTTP:            true,
			expectedErrString: "no http endpoints found for providers [%s]: provider does not support http",
		},
	}
	for i, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			testProvider := fmt.Sprintf("t01000%d", i)
			mn := mocknet.New()
			source, err := mn.GenPeer()
			require.NoError(t, err)
			other, err := mn.GenPeer()
			require.NoError(t, err)
			if !testCase.notDialable {
				mn.LinkAll()
			}
			minerInfoFetcher := &minerInfoFetcher{
				miners: make(map[string]*replication.MinerInfo),
			}
			if !testCase.minerInfoNotFindable {
				minerInfoFetcher.miners[testProvider] = &replication.MinerInfo{
					PeerID:     other.ID(),
					Multiaddrs: other.Addrs(),
				}
			}
			m, err := multiaddr.NewMultiaddr("/dns4/elastic.dag.house/tcp/443/wss")
			require.NoError(t, err)
			response := boostly.TransportsQueryResponse{}
			if !testCase.noHTTP {
				response.Protocols = append(response.Protocols, struct {
					Name      string                `json:"name,omitempty"`
					Addresses []multiaddr.Multiaddr `json:"addresses,omitempty"`
				}{
					Name:      "http",
					Addresses: []multiaddr.Multiaddr{m},
				})
			}

			if !testCase.notListeningOnTransports {
				handler := transportsListener{t, response}.HandleQueries
				other.SetStreamHandler(boostly.FilRetrievalTransportsProtocol_1_0_0, handler)
			}

			endpointFinder := endpointfinder.NewEndpointFinder(minerInfoFetcher, source, endpointfinder.WithErrorLruSize(3))

			addrInfos, err := endpointFinder.FindHTTPEndpoints(context.Background(), []string{testProvider})
			if testCase.expectedErrString == "" {
				require.NoError(t, err)
				require.Len(t, addrInfos, 1)
				require.Equal(t, addrInfos[0], peer.AddrInfo{
					ID:    other.ID(),
					Addrs: []multiaddr.Multiaddr{m},
				})
				// second call should cache
				addrInfos, err := endpointFinder.FindHTTPEndpoints(context.Background(), []string{testProvider})
				require.NoError(t, err)
				require.Len(t, addrInfos, 1)
				require.Equal(t, addrInfos[0], peer.AddrInfo{
					ID:    other.ID(),
					Addrs: []multiaddr.Multiaddr{m},
				})
				require.Equal(t, minerInfoFetcher.callCount, 1)
			} else {
				var errMessage string
				if testCase.testName == "unable to dial provider" {
					errMessage = fmt.Sprintf(testCase.expectedErrString, testProvider, source.ID(), other.ID())
				} else {
					errMessage = fmt.Sprintf(testCase.expectedErrString, testProvider)
				}
				require.EqualError(t, err, errMessage)
				require.Nil(t, addrInfos)
				// second call should cache error
				addrInfos, err := endpointFinder.FindHTTPEndpoints(context.Background(), []string{testProvider})
				require.EqualError(t, err, errMessage)
				require.Nil(t, addrInfos)
				require.Equal(t, minerInfoFetcher.callCount, 1)
			}
		})
	}
}

// TODO these should probably be public in boostly
var (
	transporsIpldSchema = `
type Multiaddr bytes
type Protocol struct {
  Name String
  Addresses [Multiaddr]
}
type TransportsQueryResponse struct {
  Protocols [Protocol]
}`
)

var reg = registry.NewRegistry()

func init() {
	if err := reg.RegisterType(
		(*boostly.TransportsQueryResponse)(nil),
		transporsIpldSchema,
		"TransportsQueryResponse",
		bindnode.TypedBytesConverter((*multiaddr.Multiaddr)(nil), func(b []byte) (any, error) {
			switch ma, err := multiaddr.NewMultiaddrBytes(b); {
			case err != nil:
				return nil, err
			default:
				return &ma, err
			}
		}, func(v any) ([]byte, error) {
			switch ma, ok := v.(*multiaddr.Multiaddr); {
			case !ok:
				return nil, fmt.Errorf("expected *Multiaddr value")
			default:
				return (*ma).Bytes(), nil
			}
		}),
	); err != nil {
		panic(err)
	}
}

type transportsListener struct {
	t        *testing.T
	response boostly.TransportsQueryResponse
}

// Called when the client opens a libp2p stream
func (l transportsListener) HandleQueries(s network.Stream) {
	defer func() { _ = s.Close() }()

	// Write the response to the client
	err := reg.TypeToWriter(&l.response, s, dagcbor.Encode)
	require.NoError(l.t, err)
}

type minerInfoFetcher struct {
	callCount int
	miners    map[string]*replication.MinerInfo
}

var errMinerNotFound = errors.New("miner not found")

func (mif *minerInfoFetcher) GetProviderInfo(ctx context.Context, provider string) (*replication.MinerInfo, error) {
	mif.callCount++
	mi, ok := mif.miners[provider]
	if !ok {
		return nil, errMinerNotFound
	}
	return mi, nil
}
