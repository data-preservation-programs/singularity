package dealpusher

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
)

type noopPDPProofSetManager struct{}

func (noopPDPProofSetManager) EnsureProofSet(_ context.Context, _ model.Wallet, _ string) (uint64, error) {
	return 1, nil
}

func (noopPDPProofSetManager) QueueAddRoots(_ context.Context, _ uint64, _ []cid.Cid, _ PDPSchedulingConfig) (*PDPQueuedTx, error) {
	return &PDPQueuedTx{Hash: "0x1"}, nil
}

type noopPDPTransactionConfirmer struct{}

func (noopPDPTransactionConfirmer) WaitForConfirmations(_ context.Context, txHash string, _ uint64, _ time.Duration) (*PDPTransactionReceipt, error) {
	return &PDPTransactionReceipt{Hash: txHash}, nil
}

func TestDealPusher_ResolveScheduleDealType_DefaultsToMarket(t *testing.T) {
	d := &DealPusher{}
	require.Equal(t, model.DealTypeMarket, d.resolveScheduleDealType(&model.Schedule{}))
}

func TestDealPusher_RunSchedule_PDPWithoutDependenciesReturnsConfiguredError(t *testing.T) {
	d := &DealPusher{
		scheduleDealTypeResolver: func(_ *model.Schedule) model.DealType {
			return model.DealTypePDP
		},
	}

	state, err := d.runSchedule(context.Background(), &model.Schedule{})
	require.Error(t, err)
	require.Equal(t, model.ScheduleError, state)
	require.Contains(t, err.Error(), "pdp scheduling dependencies are not configured")
}

func TestDealPusher_RunSchedule_PDPWithDependenciesReturnsNotImplemented(t *testing.T) {
	d := &DealPusher{
		pdpProofSetManager: noopPDPProofSetManager{},
		pdpTxConfirmer:     noopPDPTransactionConfirmer{},
		scheduleDealTypeResolver: func(_ *model.Schedule) model.DealType {
			return model.DealTypePDP
		},
	}

	state, err := d.runSchedule(context.Background(), &model.Schedule{})
	require.Error(t, err)
	require.Equal(t, model.ScheduleError, state)
	require.Contains(t, err.Error(), "pdp scheduling path is not implemented")
}
