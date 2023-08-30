package cmd

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/handler/deal"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/gotidy/ptr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func swapDealHandler(mockHandler deal.Handler) func() {
	actual := deal.Default
	deal.Default = mockHandler
	return func() {
		deal.Default = actual
	}
}

func TestSendDealHandler(t *testing.T) {
	mockHandler := new(MockDeal)
	defer swapDealHandler(mockHandler)()
	mockHandler.On("SendManualHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&model.Deal{
		ID:               1,
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
		State:            "proposed",
		Provider:         "f01",
		ProposalID:       "proposal_id",
		Label:            "label",
		PieceCID:         model.CID(testutil.TestCid),
		PieceSize:        1024,
		StartEpoch:       1001,
		EndEpoch:         1999,
		SectorStartEpoch: 1500,
		Price:            "0",
		Verified:         true,
		ClientID:         "client_id",
	}, nil)
	out, _, err := Run(context.Background(), "singularity deal send-manual client provider piece_cid 1024")
	require.NoError(t, err)
	Save(t, out, "deal_send_manual.txt")
	out, _, err = Run(context.Background(), "singularity --verbose deal send-manual client provider piece_cid 1024")
	require.NoError(t, err)
	Save(t, out, "deal_send_manual_verbose.txt")
}

func TestListDealHandler(t *testing.T) {
	mockHandler := new(MockDeal)
	defer swapDealHandler(mockHandler)()
	mockHandler.On("ListHandler", mock.Anything, mock.Anything, mock.Anything).Return([]model.Deal{
		{
			ID:               1,
			CreatedAt:        time.Time{},
			UpdatedAt:        time.Time{},
			DealID:           ptr.Of(uint64(100)),
			State:            "active",
			Provider:         "f01",
			ProposalID:       "proposal_id",
			Label:            "label",
			PieceCID:         model.CID(testutil.TestCid),
			PieceSize:        1024,
			StartEpoch:       1001,
			EndEpoch:         1999,
			SectorStartEpoch: 1500,
			Price:            "0",
			Verified:         true,
			ScheduleID:       ptr.Of(uint32(5)),
			ClientID:         "client_id",
		},
		{
			ID:               2,
			CreatedAt:        time.Time{},
			UpdatedAt:        time.Time{},
			State:            "proposed",
			Provider:         "f01",
			ProposalID:       "proposal_id_2",
			Label:            "label_2",
			PieceCID:         model.CID(testutil.TestCid),
			PieceSize:        1024,
			StartEpoch:       1011,
			EndEpoch:         2011,
			SectorStartEpoch: 1600,
			Price:            "0",
			Verified:         false,
			ScheduleID:       ptr.Of(uint32(5)),
			ClientID:         "client_id",
		},
	}, nil)
	out, _, err := Run(context.Background(), "singularity deal list")
	require.NoError(t, err)
	Save(t, out, "deal_list.txt")
	out, _, err = Run(context.Background(), "singularity --verbose deal list")
	require.NoError(t, err)
	Save(t, out, "deal_list_verbose.txt")
}
