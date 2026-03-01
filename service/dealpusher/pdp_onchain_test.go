package dealpusher

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
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

func TestCreateProofSetOptions(t *testing.T) {
	opts := createProofSetOptions()
	require.Equal(t, common.Address{}, opts.Listener)
	require.NotNil(t, opts.Value)
	require.Equal(t, "1000000000000000000", opts.Value.String())
}
