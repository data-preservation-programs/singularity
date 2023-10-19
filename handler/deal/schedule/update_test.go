package schedule

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/gotidy/ptr"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

var updateRequest = UpdateRequest{
	HTTPHeaders:           []string{"a=b"},
	URLTemplate:           ptr.Of("http://127.0.0.1"),
	PricePerGBEpoch:       ptr.Of(0.0),
	PricePerGB:            ptr.Of(0.0),
	PricePerDeal:          ptr.Of(0.0),
	Verified:              ptr.Of(true),
	IPNI:                  ptr.Of(true),
	KeepUnsealed:          ptr.Of(true),
	StartDelay:            ptr.Of("24h"),
	Duration:              ptr.Of("2400h"),
	ScheduleCron:          ptr.Of(""),
	ScheduleDealNumber:    ptr.Of(100),
	TotalDealNumber:       ptr.Of(100),
	ScheduleDealSize:      ptr.Of("1TiB"),
	TotalDealSize:         ptr.Of("1PiB"),
	Notes:                 ptr.Of("notes"),
	MaxPendingDealSize:    ptr.Of("10TiB"),
	MaxPendingDealNumber:  ptr.Of(100),
	AllowedPieceCIDs:      []string{"baga6ea4seaqao7s73y24kcutaosvacpdjgfe5pw76ooefnyqw4ynr3d2y6x2mpq"},
	ScheduleCronPerpetual: ptr.Of(true),
	Force:                 ptr.Of(true),
}

func TestUpdateHandler_DatasetNotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Schedule{
			Preparation: &model.Preparation{},
		}).Error
		require.NoError(t, err)
		_, err = Default.UpdateHandler(ctx, db, 100, updateRequest)
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
}

func TestUpdateHandler_InvalidStartDelay(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Schedule{
			Preparation: &model.Preparation{},
		}).Error
		require.NoError(t, err)
		badRequest := updateRequest
		badRequest.StartDelay = ptr.Of("1year")
		_, err = Default.UpdateHandler(ctx, db, 1, badRequest)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "invalid start delay")
	})
}

func TestUpdateHandler_InvalidDuration(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Schedule{
			Preparation: &model.Preparation{},
		}).Error
		require.NoError(t, err)
		badRequest := updateRequest
		badRequest.Duration = ptr.Of("1year")
		_, err = Default.UpdateHandler(ctx, db, 1, badRequest)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "invalid duration")
	})
}

func TestUpdateHandler_ChangeFromNonCronToCron(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Schedule{
			Preparation: &model.Preparation{},
		}).Error
		require.NoError(t, err)
		badRequest := updateRequest
		badRequest.ScheduleCron = ptr.Of("@daily")
		_, err = Default.UpdateHandler(ctx, db, 1, badRequest)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "Cannot switch ")
	})
}

func TestUpdateHandler_ChangeFromCronToNonCron(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Schedule{
			Preparation:  &model.Preparation{},
			ScheduleCron: "@daily",
		}).Error
		require.NoError(t, err)
		badRequest := updateRequest
		badRequest.ScheduleCron = ptr.Of("")
		_, err = Default.UpdateHandler(ctx, db, 1, badRequest)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "Cannot switch ")
	})
}

func TestUpdateHandler_InvalidScheduleInterval(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Schedule{
			Preparation:  &model.Preparation{},
			ScheduleCron: "@daily",
		}).Error
		require.NoError(t, err)
		badRequest := updateRequest
		badRequest.ScheduleCron = ptr.Of("1year")
		_, err = Default.UpdateHandler(ctx, db, 1, badRequest)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "invalid schedule cron")
	})
}

func TestUpdateHandler_InvalidScheduleDealSize(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Schedule{
			Preparation: &model.Preparation{},
		}).Error
		require.NoError(t, err)
		require.NoError(t, err)
		badRequest := updateRequest
		badRequest.ScheduleDealSize = ptr.Of("One PB")
		_, err = Default.UpdateHandler(ctx, db, 1, badRequest)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "invalid schedule deal size")
	})
}

func TestUpdateHandler_InvalidTotalDealSize(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Schedule{
			Preparation: &model.Preparation{},
		}).Error
		require.NoError(t, err)
		require.NoError(t, err)
		badRequest := updateRequest
		badRequest.TotalDealSize = ptr.Of("One PB")
		_, err = Default.UpdateHandler(ctx, db, 1, badRequest)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "invalid total deal size")
	})
}

func TestUpdateHandler_InvalidPendingDealSize(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Schedule{
			Preparation: &model.Preparation{},
		}).Error
		require.NoError(t, err)
		require.NoError(t, err)
		badRequest := updateRequest
		badRequest.MaxPendingDealSize = ptr.Of("One PB")
		_, err = Default.UpdateHandler(ctx, db, 1, badRequest)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "invalid max pending deal size")
	})
}

func TestUpdateHandler_InvalidAllowedPieceCID_NotCID(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Schedule{
			Preparation: &model.Preparation{},
		}).Error
		require.NoError(t, err)
		require.NoError(t, err)
		badRequest := updateRequest
		badRequest.AllowedPieceCIDs = []string{"not a cid"}
		_, err = Default.UpdateHandler(ctx, db, 1, badRequest)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "invalid allowed piece CID")
	})
}

func TestUpdateHandler_InvalidAllowedPieceCID_NotCommp(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Schedule{
			Preparation: &model.Preparation{},
		}).Error
		require.NoError(t, err)
		require.NoError(t, err)
		badRequest := updateRequest
		badRequest.AllowedPieceCIDs = []string{"bafybeiejlvvmfokp5c6q2eqgbfjeaokz3nqho5c7yy3ov527vsatgsqfma"}
		_, err = Default.UpdateHandler(ctx, db, 1, badRequest)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "not commp")
	})
}

func TestUpdateHandler_Success(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Schedule{
			Preparation: &model.Preparation{},
		}).Error
		require.NoError(t, err)
		schedule, err := Default.UpdateHandler(ctx, db, 1, updateRequest)
		require.NoError(t, err)
		require.NotNil(t, schedule)
		require.True(t, schedule.Force)
	})
}

func TestUpdateHandler_OverrideHeader(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Schedule{
			Preparation: &model.Preparation{},
			HTTPHeaders: map[string]string{"a": "b"},
		}).Error
		require.NoError(t, err)
		updateRequest := updateRequest
		updateRequest.HTTPHeaders = []string{"a=c"}
		schedule, err := Default.UpdateHandler(ctx, db, 1, updateRequest)
		require.NoError(t, err)
		require.NotNil(t, schedule)
		require.Equal(t, "c", schedule.HTTPHeaders["a"])
	})
}

func TestUpdateHandler_ClearAllHeaders(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Schedule{
			Preparation: &model.Preparation{},
			HTTPHeaders: map[string]string{"a": "b"},
		}).Error
		require.NoError(t, err)
		updateRequest := updateRequest
		updateRequest.HTTPHeaders = []string{""}
		schedule, err := Default.UpdateHandler(ctx, db, 1, updateRequest)
		require.NoError(t, err)
		require.NotNil(t, schedule)
		require.Len(t, schedule.HTTPHeaders, 0)
	})
}

func TestUpdateHandler_ClearSingleHeader(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Schedule{
			Preparation: &model.Preparation{},
			HTTPHeaders: map[string]string{"a": "b", "c": "d"},
		}).Error
		require.NoError(t, err)
		updateRequest := updateRequest
		updateRequest.HTTPHeaders = []string{"c="}
		schedule, err := Default.UpdateHandler(ctx, db, 1, updateRequest)
		require.NoError(t, err)
		require.NotNil(t, schedule)
		require.Len(t, schedule.HTTPHeaders, 1)
	})
}

func TestUpdateHandler_InvalidHeader(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Schedule{
			Preparation: &model.Preparation{},
		}).Error
		require.NoError(t, err)
		updateRequest := updateRequest
		updateRequest.HTTPHeaders = []string{"abcd"}
		_, err = Default.UpdateHandler(ctx, db, 1, updateRequest)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	})
}
