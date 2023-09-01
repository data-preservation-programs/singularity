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
	"gorm.io/gorm"
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
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(MockSchedule)
		defer swapScheduleHandler(mockHandler)()
		mockHandler.On("PauseHandler", mock.Anything, mock.Anything, mock.Anything).Return(&testSchedule, nil)
		_, _, err := runner.Run(ctx, "singularity deal schedule pause 1")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose deal schedule pause 1")
		require.NoError(t, err)
	})
}

func TestScheduleResumeHandler(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(MockSchedule)
		defer swapScheduleHandler(mockHandler)()
		mockHandler.On("ResumeHandler", mock.Anything, mock.Anything, mock.Anything).Return(&testSchedule, nil)
		_, _, err := runner.Run(ctx, "singularity deal schedule resume 1")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose deal schedule resume 1")
		require.NoError(t, err)
	})
}

func TestScheduleListHandler(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(MockSchedule)
		defer swapScheduleHandler(mockHandler)()
		mockHandler.On("ListHandler", mock.Anything, mock.Anything).Return([]model.Schedule{testSchedule}, nil)
		_, _, err := runner.Run(ctx, "singularity deal schedule list")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose deal schedule list")
		require.NoError(t, err)
	})
}

func TestScheduleCreateHandler(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(MockSchedule)
		defer swapScheduleHandler(mockHandler)()
		tmp := t.TempDir()
		err := os.WriteFile(filepath.Join(tmp, "cid.txt"), []byte(testutil.TestCid.String()), 0644)
		require.NoError(t, err)
		mockHandler.On("CreateHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&testSchedule, nil)
		_, _, err = runner.Run(ctx, fmt.Sprintf("singularity deal schedule create --allowed-piece-cid-file %s --allowed-piece-cid %s 5 provider",
			testutil.EscapePath(filepath.Join(tmp, "cid.txt")),
			testutil.TestCid.String()))
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, fmt.Sprintf("singularity --verbose deal schedule create --allowed-piece-cid-file %s --allowed-piece-cid %s 5 provider",
			testutil.EscapePath(filepath.Join(tmp, "cid.txt")),
			testutil.TestCid.String()))
		require.NoError(t, err)
	})
}
