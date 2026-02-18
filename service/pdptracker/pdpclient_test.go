package pdptracker

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/data-preservation-programs/go-synapse/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
)

type mockDataSet struct {
	listener      common.Address
	provider      common.Address
	live          bool
	nextChallenge uint64
	pieces        []cid.Cid
}

type mockPDPVerifier struct {
	dataSets map[uint64]*mockDataSet
}

func (m *mockPDPVerifier) GetNextDataSetId(_ *bind.CallOpts) (uint64, error) {
	var max uint64
	for id := range m.dataSets {
		if id > max {
			max = id
		}
	}
	return max + 1, nil
}

func (m *mockPDPVerifier) GetDataSetListener(_ *bind.CallOpts, setId *big.Int) (common.Address, error) {
	data, ok := m.dataSets[setId.Uint64()]
	if !ok {
		return common.Address{}, errors.New("not found")
	}
	return data.listener, nil
}

func (m *mockPDPVerifier) GetDataSetStorageProvider(_ *bind.CallOpts, setId *big.Int) (common.Address, common.Address, error) {
	data, ok := m.dataSets[setId.Uint64()]
	if !ok {
		return common.Address{}, common.Address{}, errors.New("not found")
	}
	return data.provider, common.Address{}, nil
}

func (m *mockPDPVerifier) DataSetLive(_ *bind.CallOpts, setId *big.Int) (bool, error) {
	data, ok := m.dataSets[setId.Uint64()]
	if !ok {
		return false, errors.New("not found")
	}
	return data.live, nil
}

func (m *mockPDPVerifier) GetNextChallengeEpoch(_ *bind.CallOpts, setId *big.Int) (*big.Int, error) {
	data, ok := m.dataSets[setId.Uint64()]
	if !ok {
		return nil, errors.New("not found")
	}
	return new(big.Int).SetUint64(data.nextChallenge), nil
}

func (m *mockPDPVerifier) GetActivePieces(_ *bind.CallOpts, setId *big.Int, offset *big.Int, limit *big.Int) (activePiecesResult, error) {
	data, ok := m.dataSets[setId.Uint64()]
	if !ok {
		return activePiecesResult{}, errors.New("not found")
	}

	start := int(offset.Uint64())
	if start >= len(data.pieces) {
		return activePiecesResult{Pieces: nil, HasMore: false}, nil
	}
	end := start + int(limit.Uint64())
	if end > len(data.pieces) {
		end = len(data.pieces)
	}

	out := make([]contracts.CidsCid, 0, end-start)
	for _, piece := range data.pieces[start:end] {
		out = append(out, contracts.CidsCid{Data: piece.Bytes()})
	}

	return activePiecesResult{
		Pieces:  out,
		HasMore: end < len(data.pieces),
	}, nil
}

func TestChainPDPClient_GetProofSetsForClient(t *testing.T) {
	originalNetwork := address.CurrentNetwork
	t.Cleanup(func() {
		address.CurrentNetwork = originalNetwork
	})
	address.CurrentNetwork = address.Mainnet

	listener := common.HexToAddress("0x1111111111111111111111111111111111111111")
	provider := common.HexToAddress("0x2222222222222222222222222222222222222222")

	listenerAddr, err := address.NewDelegatedAddress(10, listener.Bytes())
	require.NoError(t, err)
	providerAddr, err := address.NewDelegatedAddress(10, provider.Bytes())
	require.NoError(t, err)

	piece1, err := cid.Decode("baga6ea4seaqao7s73y24kcutaosvacpdjgfe5pw76ooefnyqw4ynr3d2y6x2mpq")
	require.NoError(t, err)
	piece2, err := cid.Decode("baga6ea4seaqgwm2a6rfh53y5a4qbm5zhqyixwut3wst6dfrlghm2f5l6t4o2mry")
	require.NoError(t, err)

	mock := &mockPDPVerifier{
		dataSets: map[uint64]*mockDataSet{
			1: {
				listener:      listener,
				provider:      provider,
				live:          true,
				nextChallenge: 42,
				pieces:        []cid.Cid{piece1, piece2},
			},
			2: {
				listener: common.HexToAddress("0x3333333333333333333333333333333333333333"),
				provider: provider,
				live:     false,
				pieces:   []cid.Cid{piece1},
			},
		},
	}

	client := &ChainPDPClient{
		contract: mock,
		pageSize: 1,
	}

	proofSets, err := client.GetProofSetsForClient(context.Background(), listenerAddr)
	require.NoError(t, err)
	require.Len(t, proofSets, 1)

	proofSet := proofSets[0]
	require.EqualValues(t, 1, proofSet.ProofSetID)
	require.Equal(t, listenerAddr, proofSet.ClientAddress)
	require.Equal(t, providerAddr, proofSet.ProviderAddress)
	require.True(t, proofSet.IsLive)
	require.EqualValues(t, 42, proofSet.NextChallengeEpoch)
	require.Len(t, proofSet.PieceCIDs, 2)
	require.True(t, piece1.Equals(proofSet.PieceCIDs[0]))
	require.True(t, piece2.Equals(proofSet.PieceCIDs[1]))
}

func TestChainPDPClient_GetProofSetsForClient_InvalidAddress(t *testing.T) {
	client := &ChainPDPClient{
		contract: &mockPDPVerifier{dataSets: map[uint64]*mockDataSet{}},
		pageSize: 1,
	}

	addr, err := address.NewFromString("f0100")
	require.NoError(t, err)

	_, err = client.GetProofSetsForClient(context.Background(), addr)
	require.Error(t, err)
}
