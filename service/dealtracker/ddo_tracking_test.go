package dealtracker

import (
	"context"
	"fmt"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/ipfs/boxo/util"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type mockDDOAllocTracker struct {
	statuses map[uint64]*model.DDOAllocationStatus
	calls    []uint64
	err      error
}

func (m *mockDDOAllocTracker) GetAllocationInfo(_ context.Context, allocationID uint64) (*model.DDOAllocationStatus, error) {
	m.calls = append(m.calls, allocationID)
	if m.err != nil {
		return nil, m.err
	}
	s, ok := m.statuses[allocationID]
	if !ok {
		return &model.DDOAllocationStatus{}, nil
	}
	return s, nil
}

func TestTrackDDOAllocations_NilTracker(t *testing.T) {
	dt := &DealTracker{}
	require.NoError(t, dt.trackDDOAllocations(context.Background()))
}

func TestTrackDDOAllocations_ActivatesDeals(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		allocActivated := uint64(100)
		allocPending := uint64(200)

		cidA := model.CID(cid.NewCidV1(cid.Raw, util.Hash([]byte("ddo-a"))))
		cidB := model.CID(cid.NewCidV1(cid.Raw, util.Hash([]byte("ddo-b"))))

		require.NoError(t, db.Create([]model.Deal{
			{
				State:           model.DealProposed,
				DealType:        model.DealTypeDDO,
				Provider:        "f01234",
				PieceCID:        cidA,
				PieceSize:       1024,
				DDOAllocationID: &allocActivated,
			},
			{
				State:           model.DealProposed,
				DealType:        model.DealTypeDDO,
				Provider:        "f01234",
				PieceCID:        cidB,
				PieceSize:       1024,
				DDOAllocationID: &allocPending,
			},
		}).Error)

		mock := &mockDDOAllocTracker{
			statuses: map[uint64]*model.DDOAllocationStatus{
				allocActivated: {Activated: true, SectorNumber: 7},
				allocPending:   {Activated: false},
			},
		}

		dt := &DealTracker{dbNoContext: db, ddoAllocTracker: mock}
		require.NoError(t, dt.trackDDOAllocations(ctx))

		require.Len(t, mock.calls, 2)

		var deals []model.Deal
		require.NoError(t, db.Order("id asc").Find(&deals).Error)
		require.Len(t, deals, 2)
		require.Equal(t, model.DealActive, deals[0].State)
		require.Equal(t, model.DealProposed, deals[1].State)
	})
}

func TestTrackDDOAllocations_ContinuesOnError(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		allocBad := uint64(300)
		allocGood := uint64(400)

		cidC := model.CID(cid.NewCidV1(cid.Raw, util.Hash([]byte("ddo-c"))))
		cidD := model.CID(cid.NewCidV1(cid.Raw, util.Hash([]byte("ddo-d"))))

		require.NoError(t, db.Create([]model.Deal{
			{
				State:           model.DealProposed,
				DealType:        model.DealTypeDDO,
				Provider:        "f01234",
				PieceCID:        cidC,
				PieceSize:       1024,
				DDOAllocationID: &allocBad,
			},
			{
				State:           model.DealProposed,
				DealType:        model.DealTypeDDO,
				Provider:        "f01234",
				PieceCID:        cidD,
				PieceSize:       1024,
				DDOAllocationID: &allocGood,
			},
		}).Error)

		mock := &mockDDOAllocTracker{
			statuses: map[uint64]*model.DDOAllocationStatus{
				allocGood: {Activated: true, SectorNumber: 9},
			},
		}

		// per-allocation errors: return error only for allocBad
		errorTracker := &perAllocationErrorTracker{
			inner:   mock,
			errorOn: allocBad,
		}

		dt := &DealTracker{dbNoContext: db, ddoAllocTracker: errorTracker}
		require.NoError(t, dt.trackDDOAllocations(ctx))

		var deals []model.Deal
		require.NoError(t, db.Order("id asc").Find(&deals).Error)
		require.Len(t, deals, 2)
		// first deal stays proposed due to error
		require.Equal(t, model.DealProposed, deals[0].State)
		// second deal activated despite first failing
		require.Equal(t, model.DealActive, deals[1].State)
	})
}

func TestTrackDDOAllocations_SkipsNonPendingDeals(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		allocID := uint64(500)
		cidE := model.CID(cid.NewCidV1(cid.Raw, util.Hash([]byte("ddo-e"))))

		require.NoError(t, db.Create(&model.Deal{
			State:           model.DealActive,
			DealType:        model.DealTypeDDO,
			Provider:        "f01234",
			PieceCID:        cidE,
			PieceSize:       1024,
			DDOAllocationID: &allocID,
		}).Error)

		mock := &mockDDOAllocTracker{
			statuses: map[uint64]*model.DDOAllocationStatus{
				allocID: {Activated: true},
			},
		}

		dt := &DealTracker{dbNoContext: db, ddoAllocTracker: mock}
		require.NoError(t, dt.trackDDOAllocations(ctx))

		// already-active deal should not be queried
		require.Empty(t, mock.calls)
	})
}

func TestRunOnce_DoesNotEpochExpireDDODeals(t *testing.T) {
	testutil.SkipIfNotExternalAPI(t)
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		cidDDOActive := model.CID(cid.NewCidV1(cid.Raw, util.Hash([]byte("ddo-active"))))
		cidDDOProposed := model.CID(cid.NewCidV1(cid.Raw, util.Hash([]byte("ddo-proposed"))))

		allocA := uint64(600)
		allocB := uint64(601)

		require.NoError(t, db.Create([]model.Deal{
			{
				State:           model.DealActive,
				DealType:        model.DealTypeDDO,
				ClientID:        "t0100",
				Provider:        "sp1",
				ProposalID:      "ddo-active",
				PieceCID:        cidDDOActive,
				PieceSize:       100,
				StartEpoch:      100,
				EndEpoch:        200,
				DDOAllocationID: &allocA,
			},
			{
				State:           model.DealProposed,
				DealType:        model.DealTypeDDO,
				ClientID:        "t0100",
				Provider:        "sp1",
				ProposalID:      "ddo-proposed",
				PieceCID:        cidDDOProposed,
				PieceSize:       100,
				StartEpoch:      100,
				EndEpoch:        200,
				DDOAllocationID: &allocB,
			},
		}).Error)

		url, server := setupTestServerWithBody(t, `{}`)
		defer server.Close()

		tracker := NewDealTracker(db, 0, url, testutil.TestLotusAPI, "", true)
		require.NoError(t, tracker.runOnce(ctx))

		var ddoActive model.Deal
		require.NoError(t, db.Where("proposal_id = ?", "ddo-active").First(&ddoActive).Error)
		require.Equal(t, model.DealActive, ddoActive.State)

		var ddoProposed model.Deal
		require.NoError(t, db.Where("proposal_id = ?", "ddo-proposed").First(&ddoProposed).Error)
		require.Equal(t, model.DealProposed, ddoProposed.State)
	})
}

// TODO: DDO allocation expiry/termination tracking
// AllocationInfo has no explicit terminal state field -- just Activated bool.
// unclear what happens on-chain when a DDO allocation's term expires:
// - does Activated flip back to false?
// - does GetAllocationInfo return zeroed struct or error?
// - is expiry purely a payments-rail concern (rail settlement)?
// need to confirm contract semantics before implementing active -> expired.
//
// func TestTrackDDOAllocations_ExpiresTerminatedDeals(t *testing.T) {
// }

type perAllocationErrorTracker struct {
	inner   *mockDDOAllocTracker
	errorOn uint64
}

func (p *perAllocationErrorTracker) GetAllocationInfo(ctx context.Context, allocationID uint64) (*model.DDOAllocationStatus, error) {
	if allocationID == p.errorOn {
		return nil, fmt.Errorf("rpc error for allocation %d", allocationID)
	}
	return p.inner.GetAllocationInfo(ctx, allocationID)
}
