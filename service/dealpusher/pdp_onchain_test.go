package dealpusher

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type fakeConfirmationClient struct {
	txReceiptFn func(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	blockNumFn  func(ctx context.Context) (uint64, error)
}

func (f *fakeConfirmationClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return f.txReceiptFn(ctx, txHash)
}

func (f *fakeConfirmationClient) BlockNumber(ctx context.Context) (uint64, error) {
	return f.blockNumFn(ctx)
}

func TestOnChainPDP_WaitForConfirmations(t *testing.T) {
	var receiptCalls int
	var blockCalls int

	c := &fakeConfirmationClient{
		txReceiptFn: func(_ context.Context, _ common.Hash) (*types.Receipt, error) {
			receiptCalls++
			if receiptCalls == 1 {
				return nil, ethereum.NotFound
			}
			return &types.Receipt{
				Status:            types.ReceiptStatusSuccessful,
				GasUsed:           10,
				EffectiveGasPrice: big.NewInt(3),
				BlockNumber:       big.NewInt(100),
			}, nil
		},
		blockNumFn: func(_ context.Context) (uint64, error) {
			blockCalls++
			if blockCalls == 1 {
				return 101, nil
			}
			return 102, nil
		},
	}

	adapter := &OnChainPDP{confirmClient: c}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	receipt, err := adapter.WaitForConfirmations(ctx, "0x1111111111111111111111111111111111111111111111111111111111111111", 2, time.Millisecond)
	require.NoError(t, err)
	require.NotNil(t, receipt)
	require.Equal(t, uint64(100), receipt.BlockNumber)
	require.Equal(t, uint64(10), receipt.GasUsed)
	require.Equal(t, uint64(types.ReceiptStatusSuccessful), receipt.Status)
	require.Equal(t, "30", receipt.CostAttoFIL.String())
}

func TestOnChainPDP_WaitForConfirmationsFailedTx(t *testing.T) {
	c := &fakeConfirmationClient{
		txReceiptFn: func(_ context.Context, _ common.Hash) (*types.Receipt, error) {
			return &types.Receipt{
				Status:            types.ReceiptStatusFailed,
				GasUsed:           5,
				EffectiveGasPrice: big.NewInt(2),
				BlockNumber:       big.NewInt(10),
			}, nil
		},
		blockNumFn: func(_ context.Context) (uint64, error) {
			return 10, nil
		},
	}

	adapter := &OnChainPDP{confirmClient: c}
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	receipt, err := adapter.WaitForConfirmations(ctx, "0x2222222222222222222222222222222222222222222222222222222222222222", 1, time.Millisecond)
	require.Error(t, err)
	require.NotNil(t, receipt)
	require.Equal(t, uint64(types.ReceiptStatusFailed), receipt.Status)
}

func TestOnChainPDP_WaitForConfirmationsInvalidHash(t *testing.T) {
	adapter := &OnChainPDP{}
	_, err := adapter.WaitForConfirmations(context.Background(), "not-a-hash", 1, time.Millisecond)
	require.Error(t, err)
}

func TestOnChainPDP_FindProofSetWithRoom_NoProofSets(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		adapter := &OnChainPDP{dbNoContext: db}
		setID, found, err := adapter.findProofSetWithRoom(ctx, "f410foo", "f410provider", 2)
		require.NoError(t, err)
		require.False(t, found)
		require.Zero(t, setID)
	})
}

func TestOnChainPDP_FindProofSetWithRoom_PicksFirstWithCapacity(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		client := "f410foo"
		provider := "f410provider"
		require.NoError(t, db.Create(&model.PDPProofSet{
			SetID:         1,
			ClientAddress: client,
			Provider:      provider,
			IsLive:        true,
		}).Error)
		require.NoError(t, db.Create(&model.PDPProofSet{
			SetID:         2,
			ClientAddress: client,
			Provider:      provider,
			IsLive:        true,
		}).Error)

		set1 := uint64(1)
		require.NoError(t, db.Create(&model.Deal{
			State:      model.DealProposed,
			DealType:   model.DealTypePDP,
			Provider:   provider,
			PieceSize:  1024,
			ProofSetID: &set1,
		}).Error)

		adapter := &OnChainPDP{dbNoContext: db}
		setID, found, err := adapter.findProofSetWithRoom(ctx, client, provider, 1)
		require.NoError(t, err)
		require.True(t, found)
		require.EqualValues(t, 2, setID)
	})
}
