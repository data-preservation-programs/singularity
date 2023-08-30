package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/handler/deal/schedule"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var testSchedule = model.Schedule{
	ID:                   1,
	CreatedAt:            time.Time{},
	UpdatedAt:            time.Time{},
	URLTemplate:          "https://127.0.0.1/{PIECE_CID}",
	HTTPHeaders:          []string{"a=b"},
	Provider:             "provider",
	PricePerGBEpoch:      0,
	PricePerGB:           0,
	PricePerDeal:         0,
	TotalDealNumber:      100,
	TotalDealSize:        200,
	Verified:             true,
	KeepUnsealed:         true,
	AnnounceToIPNI:       true,
	StartDelay:           300,
	Duration:             400,
	State:                model.ScheduleActive,
	ScheduleCron:         "* * * * *",
	ScheduleDealNumber:   500,
	ScheduleDealSize:     600,
	MaxPendingDealNumber: 700,
	MaxPendingDealSize:   800,
	Notes:                "my note",
	PreparationID:        5,
}

func swapScheduleHandler(mockHandler schedule.Handler) func() {
	actual := schedule.Default
	schedule.Default = mockHandler
	return func() {
		schedule.Default = actual
	}
}

func TestSchedulePauseHandler(t *testing.T) {
	mockHandler := new(MockSchedule)
	defer swapScheduleHandler(mockHandler)()
	mockHandler.On("PauseHandler", mock.Anything, mock.Anything, mock.Anything).Return(&testSchedule, nil)
	out, _, err := Run(context.Background(), "singularity deal schedule pause 1")
	require.NoError(t, err)
	Save(t, out, "schedule_pause.txt")
	out, _, err = Run(context.Background(), "singularity --verbose deal schedule pause 1")
	require.NoError(t, err)
	Save(t, out, "schedule_pause_verbose.txt")
}

func TestScheduleResumeHandler(t *testing.T) {
	mockHandler := new(MockSchedule)
	defer swapScheduleHandler(mockHandler)()
	mockHandler.On("ResumeHandler", mock.Anything, mock.Anything, mock.Anything).Return(&testSchedule, nil)
	out, _, err := Run(context.Background(), "singularity deal schedule resume 1")
	require.NoError(t, err)
	Save(t, out, "schedule_resume.txt")
	out, _, err = Run(context.Background(), "singularity --verbose deal schedule resume 1")
	require.NoError(t, err)
	Save(t, out, "schedule_resume_verbose.txt")
}

func TestScheduleListHandler(t *testing.T) {
	mockHandler := new(MockSchedule)
	defer swapScheduleHandler(mockHandler)()
	mockHandler.On("ListHandler", mock.Anything, mock.Anything).Return([]model.Schedule{testSchedule}, nil)
	out, _, err := Run(context.Background(), "singularity deal schedule list")
	require.NoError(t, err)
	Save(t, out, "schedule_list.txt")
	out, _, err = Run(context.Background(), "singularity --verbose deal schedule list")
	require.NoError(t, err)
	Save(t, out, "schedule_list_verbose.txt")
}

func TestScheduleCreateHandler(t *testing.T) {
	mockHandler := new(MockSchedule)
	defer swapScheduleHandler(mockHandler)()
	tmp := t.TempDir()
	err := os.WriteFile(filepath.Join(tmp, "cid.txt"), []byte(testutil.TestCid.String()), 0644)
	require.NoError(t, err)
	mockHandler.On("CreateHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&testSchedule, nil)
	out, _, err := Run(context.Background(), fmt.Sprintf("singularity deal schedule create --allowed-piece-cid-file '%s' --allowed-piece-cid '%s' 5 provider",
		testutil.EscapePath(filepath.Join(tmp, "cid.txt")),
		testutil.TestCid.String()))
	require.NoError(t, err)
	Save(t, out, "schedule_create.txt")
	out, _, err = Run(context.Background(), fmt.Sprintf("singularity --verbose deal schedule create --allowed-piece-cid-file '%s' --allowed-piece-cid '%s' 5 provider",
		testutil.EscapePath(filepath.Join(tmp, "cid.txt")),
		testutil.TestCid.String()))
	require.NoError(t, err)
	Save(t, out, "schedule_create_verbose.txt")
}
