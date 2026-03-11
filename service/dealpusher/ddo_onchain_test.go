package dealpusher

import (
	"context"
	"math/big"
	"strings"
	"testing"
	"time"

	ddocontract "github.com/Eastore-project/ddo-client/pkg/contract/ddo"
	"github.com/data-preservation-programs/go-synapse/signer"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
)

type ddoConfirmClientStub struct {
	receiptCalls int
	receipts     []*ethtypes.Receipt
	receiptErrs  []error
	blockNumber  uint64
	blockErr     error
}

func (s *ddoConfirmClientStub) TransactionReceipt(_ context.Context, _ common.Hash) (*ethtypes.Receipt, error) {
	idx := s.receiptCalls
	s.receiptCalls++
	if idx < len(s.receiptErrs) && s.receiptErrs[idx] != nil {
		return nil, s.receiptErrs[idx]
	}
	if idx < len(s.receipts) {
		return s.receipts[idx], nil
	}
	if len(s.receipts) > 0 {
		return s.receipts[len(s.receipts)-1], nil
	}
	return nil, ethereum.NotFound
}

func (s *ddoConfirmClientStub) BlockNumber(context.Context) (uint64, error) {
	if s.blockErr != nil {
		return 0, s.blockErr
	}
	return s.blockNumber, nil
}

func TestDDOSchedulingConfigValidate(t *testing.T) {
	cfg := DDOSchedulingConfig{
		BatchSize:         10,
		ConfirmationDepth: 5,
		PollingInterval:   30 * time.Second,
		TermMin:           518400,
		TermMax:           5256000,
		ExpirationOffset:  172800,
	}
	require.NoError(t, cfg.Validate())

	cfg.BatchSize = 0
	require.ErrorContains(t, cfg.Validate(), "batch size")
}

func TestOnChainDDOBuildPieceInfos(t *testing.T) {
	parsed, err := cid.Parse("bafkqaaa")
	require.NoError(t, err)

	adapter := &OnChainDDO{
		paymentToken: common.HexToAddress("0xb3042734b608a1B16e9e86B374A3f3e389B4cDf0"),
	}
	cfg := DDOSchedulingConfig{
		BatchSize:         10,
		ConfirmationDepth: 5,
		PollingInterval:   30 * time.Second,
		TermMin:           518400,
		TermMax:           5256000,
		ExpirationOffset:  172800,
	}

	pieceInfos, err := adapter.buildPieceInfos([]DDOPieceSubmission{{
		PieceCID:    parsed,
		PieceSize:   2048,
		ProviderID:  1234,
		DownloadURL: "https://example.test/piece",
	}}, cfg)
	require.NoError(t, err)
	require.Len(t, pieceInfos, 1)
	require.Equal(t, parsed.Bytes(), pieceInfos[0].PieceCid)
	require.EqualValues(t, 2048, pieceInfos[0].Size)
	require.EqualValues(t, 1234, pieceInfos[0].Provider)
	require.Equal(t, cfg.TermMin, pieceInfos[0].TermMin)
	require.Equal(t, cfg.TermMax, pieceInfos[0].TermMax)
	require.Equal(t, cfg.ExpirationOffset, pieceInfos[0].ExpirationOffset)
	require.Equal(t, "https://example.test/piece", pieceInfos[0].DownloadURL)
	require.Equal(t, adapter.paymentToken, pieceInfos[0].PaymentTokenAddress)
}

func TestOnChainDDOValidateSignerMismatch(t *testing.T) {
	baseSigner, err := signer.FromLotusExport(testutil.TestPrivateKeyHex)
	require.NoError(t, err)
	evmSigner, ok := signer.AsEVM(baseSigner)
	require.True(t, ok)

	otherSigner, err := signer.NewSecp256k1Signer([]byte{2})
	require.NoError(t, err)

	adapter := &OnChainDDO{signerAddr: otherSigner.EVMAddress()}
	err = adapter.validateSigner(evmSigner)
	require.ErrorContains(t, err, "does not match requested signer")
}

func TestOnChainDDOWaitForConfirmations(t *testing.T) {
	adapter := &OnChainDDO{
		confirmClient: &ddoConfirmClientStub{
			receiptErrs: []error{ethereum.NotFound, nil},
			receipts: []*ethtypes.Receipt{{
				Status:      ethtypes.ReceiptStatusSuccessful,
				GasUsed:     12345,
				BlockNumber: big.NewInt(100),
			}},
			blockNumber: 106,
		},
	}

	receipt, err := adapter.WaitForConfirmations(context.Background(), "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", 5, time.Millisecond)
	require.NoError(t, err)
	require.EqualValues(t, 100, receipt.BlockNumber)
	require.EqualValues(t, 12345, receipt.GasUsed)
	require.EqualValues(t, ethtypes.ReceiptStatusSuccessful, receipt.Status)
}

func TestOnChainDDOParseAllocationIDs(t *testing.T) {
	parsedABI, err := abi.JSON(strings.NewReader(ddocontract.DDOClientABI))
	require.NoError(t, err)
	event := parsedABI.Events["AllocationCreated"]

	receipt := &ethtypes.Receipt{
		Logs: []*ethtypes.Log{{
			Topics: []common.Hash{
				event.ID,
				common.HexToHash("0x0000000000000000000000001111111111111111111111111111111111111111"),
				common.BigToHash(big.NewInt(42)),
			},
		}},
	}

	adapter := &OnChainDDO{
		confirmClient: &ddoConfirmClientStub{
			receipts: []*ethtypes.Receipt{receipt},
		},
	}

	allocationIDs, err := adapter.ParseAllocationIDs(context.Background(), "0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
	require.NoError(t, err)
	require.Equal(t, []uint64{42}, allocationIDs)
}
