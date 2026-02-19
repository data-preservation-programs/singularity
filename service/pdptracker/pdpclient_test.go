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

type mockContractCaller struct {
	listeners map[uint64]common.Address
	pieces    map[uint64][]cid.Cid
}

func (m *mockContractCaller) GetDataSetListener(_ *bind.CallOpts, setId *big.Int) (common.Address, error) {
	addr, ok := m.listeners[setId.Uint64()]
	if !ok {
		return common.Address{}, errors.New("not found")
	}
	return addr, nil
}

func (m *mockContractCaller) GetActivePieces(_ *bind.CallOpts, setId *big.Int, offset *big.Int, limit *big.Int) (activePiecesResult, error) {
	all, ok := m.pieces[setId.Uint64()]
	if !ok {
		return activePiecesResult{}, errors.New("not found")
	}

	start := int(offset.Uint64())
	if start >= len(all) {
		return activePiecesResult{Pieces: nil, HasMore: false}, nil
	}
	end := start + int(limit.Uint64())
	if end > len(all) {
		end = len(all)
	}

	out := make([]contracts.CidsCid, 0, end-start)
	for _, piece := range all[start:end] {
		out = append(out, contracts.CidsCid{Data: piece.Bytes()})
	}

	return activePiecesResult{
		Pieces:  out,
		HasMore: end < len(all),
	}, nil
}

func TestChainPDPClient_GetDataSetListener(t *testing.T) {
	listener := common.HexToAddress("0x1111111111111111111111111111111111111111")
	mock := &mockContractCaller{
		listeners: map[uint64]common.Address{
			1: listener,
		},
	}
	client := &ChainPDPClient{contract: mock, pageSize: 100}

	addr, err := client.GetDataSetListener(context.Background(), 1)
	require.NoError(t, err)
	require.Equal(t, listener, addr)

	_, err = client.GetDataSetListener(context.Background(), 99)
	require.Error(t, err)
}

func TestChainPDPClient_GetActivePieces_Pagination(t *testing.T) {
	piece1, err := cid.Decode("baga6ea4seaqao7s73y24kcutaosvacpdjgfe5pw76ooefnyqw4ynr3d2y6x2mpq")
	require.NoError(t, err)
	piece2, err := cid.Decode("baga6ea4seaqgwm2a6rfh53y5a4qbm5zhqyixwut3wst6dfrlghm2f5l6t4o2mry")
	require.NoError(t, err)

	mock := &mockContractCaller{
		pieces: map[uint64][]cid.Cid{
			1: {piece1, piece2},
		},
	}
	// page size 1 to force multiple pages
	client := &ChainPDPClient{contract: mock, pageSize: 1}

	result, err := client.GetActivePieces(context.Background(), 1)
	require.NoError(t, err)
	require.Len(t, result, 2)
	require.True(t, piece1.Equals(result[0]))
	require.True(t, piece2.Equals(result[1]))
}

func TestDelegatedAddressRoundtrip(t *testing.T) {
	originalNetwork := address.CurrentNetwork
	t.Cleanup(func() { address.CurrentNetwork = originalNetwork })
	address.CurrentNetwork = address.Mainnet

	ethAddr := common.HexToAddress("0x1111111111111111111111111111111111111111")
	filAddr, err := commonToDelegatedAddress(ethAddr)
	require.NoError(t, err)
	require.Equal(t, address.Delegated, filAddr.Protocol())

	roundtrip, err := delegatedAddressToCommon(filAddr)
	require.NoError(t, err)
	require.Equal(t, ethAddr, roundtrip)
}

func TestDelegatedAddressToCommon_InvalidProtocol(t *testing.T) {
	addr, err := address.NewFromString("f0100")
	require.NoError(t, err)

	_, err = delegatedAddressToCommon(addr)
	require.Error(t, err)
}
