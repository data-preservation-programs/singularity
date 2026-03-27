package dealtracker

import (
	"context"
	"fmt"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/boxo/util"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type mockF05PaymentTracker struct {
	receipts map[string]*F05PaymentReceipt
	calls    []string
	err      error
}

func (m *mockF05PaymentTracker) GetTransactionReceipt(_ context.Context, txHash string) (*F05PaymentReceipt, error) {
	m.calls = append(m.calls, txHash)
	if m.err != nil {
		return nil, m.err
	}
	receipt, ok := m.receipts[txHash]
	if !ok {
		return nil, nil
	}
	return receipt, nil
}

func TestTrackF05Payments_NilTracker(t *testing.T) {
	dt := &DealTracker{}
	require.NoError(t, dt.trackF05Payments(context.Background()))
}

func TestTrackF05Payments_UpdatesConfirmedAndFailedDeals(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		txConfirmed := common.HexToHash("0x100").Hex()
		txFailed := common.HexToHash("0x200").Hex()
		txPending := common.HexToHash("0x300").Hex()
		txAlreadyConfirmed := common.HexToHash("0x400").Hex()

		cidA := model.CID(cid.NewCidV1(cid.Raw, util.Hash([]byte("f05-a"))))
		cidB := model.CID(cid.NewCidV1(cid.Raw, util.Hash([]byte("f05-b"))))
		cidC := model.CID(cid.NewCidV1(cid.Raw, util.Hash([]byte("f05-c"))))
		cidD := model.CID(cid.NewCidV1(cid.Raw, util.Hash([]byte("f05-d"))))

		confirmedStatus := "confirmed"
		require.NoError(t, db.Create([]model.Deal{
			{
				State:            model.DealProposed,
				DealType:         model.DealTypeF05Paid,
				Provider:         "f01234",
				PieceCID:         cidA,
				PieceSize:        1024,
				F05PaymentTxHash: &txConfirmed,
			},
			{
				State:            model.DealProposed,
				DealType:         model.DealTypeF05Paid,
				Provider:         "f01234",
				PieceCID:         cidB,
				PieceSize:        1024,
				F05PaymentTxHash: &txFailed,
			},
			{
				State:            model.DealProposed,
				DealType:         model.DealTypeF05Paid,
				Provider:         "f01234",
				PieceCID:         cidC,
				PieceSize:        1024,
				F05PaymentTxHash: &txPending,
			},
			{
				State:            model.DealProposed,
				DealType:         model.DealTypeF05Paid,
				Provider:         "f01234",
				PieceCID:         cidD,
				PieceSize:        1024,
				F05PaymentTxHash: &txAlreadyConfirmed,
				F05PaymentStatus: &confirmedStatus,
			},
		}).Error)

		mock := &mockF05PaymentTracker{
			receipts: map[string]*F05PaymentReceipt{
				txConfirmed: {Status: 1, BlockNumber: 10, GasUsed: 1000},
				txFailed:    {Status: 0, BlockNumber: 11, GasUsed: 2000},
			},
		}

		dt := &DealTracker{dbNoContext: db, f05PayTracker: mock}
		require.NoError(t, dt.trackF05Payments(ctx))

		require.ElementsMatch(t, []string{txConfirmed, txFailed, txPending}, mock.calls)

		var deals []model.Deal
		require.NoError(t, db.Order("id asc").Find(&deals).Error)
		require.Len(t, deals, 4)

		require.NotNil(t, deals[0].F05PaymentStatus)
		require.Equal(t, "confirmed", *deals[0].F05PaymentStatus)
		require.Equal(t, model.DealProposed, deals[0].State)

		require.NotNil(t, deals[1].F05PaymentStatus)
		require.Equal(t, "failed", *deals[1].F05PaymentStatus)
		require.Equal(t, model.DealErrored, deals[1].State)
		require.Contains(t, deals[1].ErrorMessage, txFailed)

		require.Nil(t, deals[2].F05PaymentStatus)
		require.Equal(t, model.DealProposed, deals[2].State)

		require.NotNil(t, deals[3].F05PaymentStatus)
		require.Equal(t, "confirmed", *deals[3].F05PaymentStatus)
		require.Equal(t, model.DealProposed, deals[3].State)
	})
}

func TestTrackF05Payments_ContinuesOnError(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		txBad := common.HexToHash("0x500").Hex()
		txGood := common.HexToHash("0x600").Hex()

		cidA := model.CID(cid.NewCidV1(cid.Raw, util.Hash([]byte("f05-e"))))
		cidB := model.CID(cid.NewCidV1(cid.Raw, util.Hash([]byte("f05-f"))))

		require.NoError(t, db.Create([]model.Deal{
			{
				State:            model.DealProposed,
				DealType:         model.DealTypeF05Paid,
				Provider:         "f01234",
				PieceCID:         cidA,
				PieceSize:        1024,
				F05PaymentTxHash: &txBad,
			},
			{
				State:            model.DealProposed,
				DealType:         model.DealTypeF05Paid,
				Provider:         "f01234",
				PieceCID:         cidB,
				PieceSize:        1024,
				F05PaymentTxHash: &txGood,
			},
		}).Error)

		dt := &DealTracker{
			dbNoContext: db,
			f05PayTracker: &perTxErrorF05Tracker{
				errorOn: txBad,
				inner: &mockF05PaymentTracker{
					receipts: map[string]*F05PaymentReceipt{
						txGood: {Status: 1},
					},
				},
			},
		}
		require.NoError(t, dt.trackF05Payments(ctx))

		var deals []model.Deal
		require.NoError(t, db.Order("id asc").Find(&deals).Error)
		require.Len(t, deals, 2)
		require.Nil(t, deals[0].F05PaymentStatus)
		require.NotNil(t, deals[1].F05PaymentStatus)
		require.Equal(t, "confirmed", *deals[1].F05PaymentStatus)
	})
}

type perTxErrorF05Tracker struct {
	inner   *mockF05PaymentTracker
	errorOn string
}

func (p *perTxErrorF05Tracker) GetTransactionReceipt(ctx context.Context, txHash string) (*F05PaymentReceipt, error) {
	if txHash == p.errorOn {
		return nil, fmt.Errorf("rpc error for tx %s", txHash)
	}
	return p.inner.GetTransactionReceipt(ctx, txHash)
}
