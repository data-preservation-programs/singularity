package dealpusher

import (
	"testing"

	commcid "github.com/filecoin-project/go-fil-commcid"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
)

func TestBuildPullInputs_PayloadSizeEncodedNotPaddedSize(t *testing.T) {
	t.Parallel()

	v1 := cid.Cid(calculateCommp(t, generateRandomBytes(1000), 4194304))
	const carSize = int64(2_097_436)
	const pieceSize = int64(4_194_304)

	cidsV2, inputs, err := buildPullInputs(
		[]PDPPieceInput{{PieceCID: v1, PieceSize: pieceSize, PayloadSize: carSize}},
		"https://example.test",
	)
	require.NoError(t, err)
	require.Len(t, cidsV2, 1)
	require.Len(t, inputs, 1)

	_, decoded, err := commcid.PieceCidV2ToDataCommitment(cidsV2[0])
	require.NoError(t, err)
	require.Equal(t, uint64(carSize), decoded, "V2 CID must encode PayloadSize, not PieceSize*127/128")
	require.NotEqual(t, uint64(pieceSize)*127/128, decoded)

	require.Equal(t, "https://example.test/piece/"+cidsV2[0].String(), inputs[0].SourceURL)
	require.Equal(t, cidsV2[0].String(), inputs[0].PieceCID)
}

func TestBuildPullInputs_MultiPiecePerIndexEncoding(t *testing.T) {
	t.Parallel()

	v1a := cid.Cid(calculateCommp(t, generateRandomBytes(1000), 4194304))
	v1b := cid.Cid(calculateCommp(t, generateRandomBytes(2000), 4194304))
	v1c := cid.Cid(calculateCommp(t, generateRandomBytes(3000), 8388608))
	require.NotEqual(t, v1a, v1b)
	require.NotEqual(t, v1b, v1c)

	in := []PDPPieceInput{
		{PieceCID: v1a, PieceSize: 4194304, PayloadSize: 1_000_000},
		{PieceCID: v1b, PieceSize: 4194304, PayloadSize: 2_000_000},
		{PieceCID: v1c, PieceSize: 8388608, PayloadSize: 5_000_000},
	}
	cidsV2, inputs, err := buildPullInputs(in, "https://example.test")
	require.NoError(t, err)
	require.Len(t, cidsV2, len(in))
	require.Len(t, inputs, len(in))

	for i, p := range in {
		v1Decoded, payload, err := commcid.PieceCidV1FromV2(cidsV2[i])
		require.NoError(t, err)
		require.Equal(t, uint64(p.PayloadSize), payload, "piece %d payloadSize", i)
		require.Equal(t, p.PieceCID, v1Decoded, "piece %d v1 CID", i)
		require.Equal(t, cidsV2[i].String(), inputs[i].PieceCID)
		require.Equal(t, "https://example.test/piece/"+cidsV2[i].String(), inputs[i].SourceURL)
	}

	require.NotEqual(t, cidsV2[0], cidsV2[1])
	require.NotEqual(t, cidsV2[1], cidsV2[2])
}

func TestBuildPullInputs_RejectsMissingPayloadSize(t *testing.T) {
	t.Parallel()

	v1 := cid.Cid(calculateCommp(t, generateRandomBytes(1000), 4194304))

	for _, payload := range []int64{0, -1} {
		_, _, err := buildPullInputs(
			[]PDPPieceInput{{PieceCID: v1, PieceSize: 4194304, PayloadSize: payload}},
			"https://example.test",
		)
		require.Error(t, err)
		require.Contains(t, err.Error(), "missing PayloadSize")
	}
}
