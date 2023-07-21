//go:build !(windows && 386)

package replication

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/replication/internal/proposal110"
	"github.com/data-preservation-programs/singularity/replication/internal/proposal120"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/filecoin-project/go-address"
	cborutil "github.com/filecoin-project/go-cbor-util"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	"github.com/jellydator/ttlcache/v3"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func testProposal(t *testing.T) proposal110.ClientDealProposal {
	pieceCID, err := cid.Decode("baga6ea4seaqdyupo27fj2fk2mtefzlxvrbf6kdi4twdpccdzbyqrbpsvfsh5ula")
	require.NoError(t, err)
	clientAddr, err := address.NewFromString("f01000")
	require.NoError(t, err)
	provider, err := address.NewFromString("f01001")
	require.NoError(t, err)
	return proposal110.ClientDealProposal{
		Proposal: proposal110.DealProposal{
			PieceCID:     pieceCID,
			PieceSize:    1024,
			VerifiedDeal: true,
			Client:       clientAddr,
			Provider:     provider,
			Label:        proposal110.DealLabel{},
			StartEpoch:   100,
			EndEpoch:     200,
			StoragePricePerEpoch: abi.TokenAmount{
				Int: big.NewInt(101),
			},
			ProviderCollateral: abi.TokenAmount{
				Int: big.NewInt(102),
			},
			ClientCollateral: abi.TokenAmount{
				Int: big.NewInt(103),
			},
		},
		ClientSignature: crypto.Signature{
			Type: 1,
			Data: []byte("signature"),
		},
	}
}

func setupBasicHost(t *testing.T, ctx context.Context, port string) host.Host {
	m, err := multiaddr.NewMultiaddr("/ip4/127.0.0.1/tcp/" + port)
	require.NoError(t, err)
	h, err := util.InitHost(nil, m)
	require.NoError(t, err)
	h.SetStreamHandler(StorageProposalV120, func(s network.Stream) {
		var deal proposal120.DealParams
		err := cborutil.ReadCborRPC(s, &deal)
		require.NoError(t, err)
		resp := &proposal120.DealResponse{
			Accepted: true,
			Message:  "accepted",
		}
		err = cborutil.WriteCborRPC(s, resp)
		require.NoError(t, err)
	})
	h.SetStreamHandler(StorageProposalV111, func(s network.Stream) {
		var deal proposal110.Proposal
		err := cborutil.ReadCborRPC(s, &deal)
		require.NoError(t, err)
		c, err := cid.Decode("bafy2bzaceczlclcg4notjmrz4ayenf7fi4mngnqbgjs27r3resyhzwxjnviay")
		require.NoError(t, err)
		resp := &proposal110.SignedResponse{
			Response: proposal110.Response{
				State:          1,
				Message:        "accepted",
				Proposal:       c,
				PublishMessage: nil,
			},
			Signature: &crypto.Signature{
				Type: 1,
				Data: []byte("signature"),
			},
		}
		err = cborutil.WriteCborRPC(s, resp)
		require.NoError(t, err)
	})
	go func() {
		<-ctx.Done()
		h.Close()
	}()
	return h
}

func TestDealMaker_MakeDeal(t *testing.T) {
	addr := "f1fib3pv7jua2ockdugtz7viz3cyy6lkhh7rfx3sa"
	key := "7b2254797065223a22736563703235366b31222c22507269766174654b6579223a226b35507976337148327349586343595a58594f5775453149326e32554539436861556b6c4e36695a5763453d227d"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	server := setupBasicHost(t, ctx, "10001")
	client := setupBasicHost(t, ctx, "10002")
	defer server.Close()
	defer client.Close()
	maker := NewDealMaker(nil, client, time.Hour, time.Second)
	defer maker.Close()
	wallet := model.Wallet{
		ID:         "f047684",
		Address:    addr,
		PrivateKey: key,
	}
	rootCID, err := cid.Decode("bafy2bzaceczlclcg4notjmrz4ayenf7fi4mngnqbgjs27r3resyhzwxjnviay")
	require.NoError(t, err)
	c, err := cid.Decode("baga6ea4seaqdyupo27fj2fk2mtefzlxvrbf6kdi4twdpccdzbyqrbpsvfsh5ula")
	require.NoError(t, err)
	car := model.Car{
		RootCID:   model.CID(rootCID),
		PieceCID:  model.CID(c),
		PieceSize: 1024,
		FileSize:  1000,
	}
	dealConfig := DealConfig{
		Provider:        "f01000",
		StartDelay:      time.Minute,
		Duration:        time.Hour,
		Verified:        true,
		KeepUnsealed:    true,
		AnnounceToIPNI:  true,
		PricePerDeal:    0,
		PricePerGB:      0,
		PricePerGBEpoch: 0,
	}
	maker.minerInfoCache.Set("f01000", &MinerInfo{
		PeerID:     server.ID(),
		Multiaddrs: server.Addrs(),
	}, ttlcache.DefaultTTL)
	maker.collateralCache.Set("1024-true", abi.NewTokenAmount(1000000000000000000), ttlcache.DefaultTTL)

	maker.protocolsCache.Set(server.ID(), []protocol.ID{
		StorageProposalV120,
	}, ttlcache.DefaultTTL)

	_, err = maker.MakeDeal(ctx, wallet, car, dealConfig)
	require.NoError(t, err)

	maker.protocolsCache.Set(server.ID(), []protocol.ID{
		StorageProposalV111,
	}, ttlcache.DefaultTTL)

	_, err = maker.MakeDeal(ctx, wallet, car, dealConfig)
	require.NoError(t, err)

}

func TestDealMaker_MakeDeal111(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	server := setupBasicHost(t, ctx, "10001")
	client := setupBasicHost(t, ctx, "10002")
	defer server.Close()
	defer client.Close()
	maker := NewDealMaker(nil, client, time.Hour, time.Second)
	defer maker.Close()
	rootCID, err := cid.Decode("bafy2bzaceczlclcg4notjmrz4ayenf7fi4mngnqbgjs27r3resyhzwxjnviay")
	require.NoError(t, err)
	proposal := testProposal(t)
	resp, err := maker.makeDeal111(
		ctx,
		proposal,
		DealConfig{
			Provider:       "f01001",
			StartDelay:     time.Minute,
			Duration:       time.Hour,
			Verified:       true,
			KeepUnsealed:   true,
			AnnounceToIPNI: true,
		},
		rootCID,
		peer.AddrInfo{
			ID:    server.ID(),
			Addrs: server.Addrs(),
		})
	require.NoError(t, err)
	require.Equal(t, uint64(1), resp.Response.State)
	require.Equal(t, "accepted", resp.Response.Message)
}

func TestDealConfig_GetPrice(t *testing.T) {
	config := DealConfig{
		PricePerDeal:    1.0,
		PricePerGB:      0,
		PricePerGBEpoch: 0,
	}
	require.Equal(t, *big.NewInt(1e18), *config.GetPrice(1000, time.Minute).Int)

	config = DealConfig{
		PricePerDeal:    0,
		PricePerGB:      0.1,
		PricePerGBEpoch: 0,
	}

	require.Equal(t, *big.NewInt(1e11), *config.GetPrice(1000, time.Minute).Int)

	config = DealConfig{
		PricePerDeal:    0,
		PricePerGB:      0,
		PricePerGBEpoch: 0.1,
	}

	require.Equal(t, *big.NewInt(2e11), *config.GetPrice(1000, time.Minute).Int)
}

func TestDealMaker_MakeDeal120(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	server := setupBasicHost(t, ctx, "10001")
	client := setupBasicHost(t, ctx, "10002")
	defer server.Close()
	defer client.Close()
	maker := NewDealMaker(nil, client, time.Hour, time.Second)
	defer maker.Close()
	rootCID, err := cid.Decode("bafy2bzaceczlclcg4notjmrz4ayenf7fi4mngnqbgjs27r3resyhzwxjnviay")
	require.NoError(t, err)
	proposal := testProposal(t)
	resp, err := maker.makeDeal120(
		ctx,
		proposal,
		uuid.New(),
		DealConfig{
			Provider:       "f01001",
			StartDelay:     time.Minute,
			Duration:       time.Hour,
			Verified:       true,
			HTTPHeaders:    []string{"key=value"},
			URLTemplate:    "http://localhost:8080/piece/{PIECE_CID}",
			KeepUnsealed:   true,
			AnnounceToIPNI: true,
		},
		0,
		rootCID,
		peer.AddrInfo{
			ID:    server.ID(),
			Addrs: server.Addrs(),
		})
	require.NoError(t, err)
	require.True(t, resp.Accepted)
	require.Equal(t, "accepted", resp.Message)
}

func TestDealMaker_GetCollateral(t *testing.T) {
	lotusClient := new(MockRPCClient)
	maker := NewDealMaker(nil, nil, time.Hour, time.Second)
	maker.lotusClient = lotusClient
	lotusClient.On("CallFor", mock.Anything, mock.Anything, "Filecoin.StateDealProviderCollateralBounds", mock.Anything).
		Return(nil).Run(func(args mock.Arguments) {
		resultPtr := args.Get(1).(*DealProviderCollateralBound)
		*resultPtr = DealProviderCollateralBound{
			Min: "8649874114492479",
			Max: "2000000000000000000000",
		}
	})
	result, err := maker.getMinCollateral(context.Background(), 34359738368, false)
	require.NoError(t, err)
	require.Equal(t, "8649874114492479", result.String())
}

func TestDealMaker_GetProtocols(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	server := setupBasicHost(t, ctx, "10001")
	client := setupBasicHost(t, ctx, "10002")
	defer server.Close()
	defer client.Close()
	maker := NewDealMaker(nil, client, time.Hour, time.Second)
	defer maker.Close()
	time.Sleep(100 * time.Millisecond)
	protocols, err := maker.getProtocols(ctx, peer.AddrInfo{
		ID:    server.ID(),
		Addrs: server.Addrs(),
	})
	require.NoError(t, err)
	require.Contains(t, protocols, protocol.ID("/ipfs/ping/1.0.0"))
	require.Contains(t, protocols, protocol.ID(StorageProposalV120))
	require.Contains(t, protocols, protocol.ID(StorageProposalV111))
}

func TestDealMaker_GetProviderInfo(t *testing.T) {
	lotusClient := new(MockRPCClient)
	maker := NewDealMaker(nil, nil, time.Hour, time.Second)
	maker.lotusClient = lotusClient

	lotusClient.On("CallFor", mock.Anything, mock.Anything, "Filecoin.StateMinerInfo", []interface{}{"address1", nil}).
		Return(nil).Run(func(args mock.Arguments) {
		resultPtr := args.Get(1).(*MinerInfo)
		*resultPtr = MinerInfo{
			PeerIDEncoded:           "12D3KooWRTsCNvyZr6zWvN2YtKuygfTyG5TqZfZ464472D4ZCqYd",
			MultiaddrsBase64Encoded: []string{"BGvR+oMGXcE="},
		}
	})

	info, err := maker.getProviderInfo(context.Background(), "address1")
	require.NoError(t, err)
	require.Len(t, info.Multiaddrs, 1)
	require.Contains(t, info.Multiaddrs[0].String(), "/tcp/24001")
	require.Equal(t, "12D3KooWRTsCNvyZr6zWvN2YtKuygfTyG5TqZfZ464472D4ZCqYd", info.PeerID.String())
}
