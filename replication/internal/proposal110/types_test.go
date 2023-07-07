package proposal110

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
)

func TestProposalMarshalCBOR(t *testing.T) {
	pieceCID, err := cid.Decode("baga6ea4seaqdyupo27fj2fk2mtefzlxvrbf6kdi4twdpccdzbyqrbpsvfsh5ula")
	require.NoError(t, err)
	client, err := address.NewFromString("f01000")
	require.NoError(t, err)
	provider, err := address.NewFromString("f01001")
	require.NoError(t, err)
	rootCID, err := cid.Decode("bafy2bzaceczlclcg4notjmrz4ayenf7fi4mngnqbgjs27r3resyhzwxjnviay")
	require.NoError(t, err)
	proposal := Proposal{
		DealProposal: &ClientDealProposal{
			Proposal: DealProposal{
				PieceCID:     pieceCID,
				PieceSize:    1024,
				VerifiedDeal: true,
				Client:       client,
				Provider:     provider,
				Label: DealLabel{
					bs:        []byte("hello"),
					notString: false,
				},
				StartEpoch: 100,
				EndEpoch:   200,
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
		},
		Piece: &DataRef{
			TransferType: "type",
			Root:         rootCID,
			PieceCid:     &pieceCID,
			PieceSize:    1024,
			RawBlockSize: 0,
		},
		FastRetrieval: true,
	}
	buf := bytes.NewBuffer([]byte{})
	err = proposal.MarshalCBOR(buf)
	require.NoError(t, err)
	unmarshalled := Proposal{}
	err = unmarshalled.UnmarshalCBOR(buf)
	require.NoError(t, err)
	require.Equal(t, proposal, unmarshalled)
}
