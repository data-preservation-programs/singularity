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
	Group:                 ptr.Of("batch-a"),
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
	DealType:              ptr.Of(string(model.DealTypeMarket)),
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

func TestUpdateHandler_InvalidDealType(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Schedule{
			Preparation: &model.Preparation{},
		}).Error
		require.NoError(t, err)
		badRequest := updateRequest
		badRequest.DealType = ptr.Of("unknown")
		_, err = Default.UpdateHandler(ctx, db, 1, badRequest)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "invalid deal type")
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
		require.Equal(t, "batch-a", schedule.Group)
		require.Equal(t, model.DealTypeMarket, schedule.DealType)
	})
}

func TestUpdateHandler_ClearGroup(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Schedule{
			Group:       "batch-a",
			Preparation: &model.Preparation{},
		}).Error
		require.NoError(t, err)
		req := UpdateRequest{Group: ptr.Of("")}
		schedule, err := Default.UpdateHandler(ctx, db, 1, req)
		require.NoError(t, err)
		require.Empty(t, schedule.Group)
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

func TestUpdateHandler_DDORequiresURLTemplate(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Schedule{
			Preparation: &model.Preparation{},
		}).Error
		require.NoError(t, err)
		req := UpdateRequest{DealType: ptr.Of(string(model.DealTypeDDO))}
		_, err = Default.UpdateHandler(ctx, db, 1, req)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "URL template")
	})
}

func TestUpdateHandler_DDOClearURLTemplateRejected(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Schedule{
			Preparation: &model.Preparation{},
			DealType:    model.DealTypeDDO,
			URLTemplate: "https://example.com/{PIECE_CID}",
		}).Error
		require.NoError(t, err)
		req := UpdateRequest{URLTemplate: ptr.Of("")}
		_, err = Default.UpdateHandler(ctx, db, 1, req)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "URL template")
	})
}

func TestUpdateHandler_DDORejectsPricePerGBEpoch(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Schedule{
			Preparation: &model.Preparation{},
			DealType:    model.DealTypeDDO,
			URLTemplate: "https://example.com/{PIECE_CID}",
		}).Error
		require.NoError(t, err)
		req := UpdateRequest{PricePerGBEpoch: ptr.Of(1e-18)}
		_, err = Default.UpdateHandler(ctx, db, 1, req)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "DDO schedules do not accept")
	})
}

func TestUpdateHandler_DDORejectsPricePerGB(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Schedule{
			Preparation: &model.Preparation{},
			DealType:    model.DealTypeDDO,
			URLTemplate: "https://example.com/{PIECE_CID}",
		}).Error
		require.NoError(t, err)
		req := UpdateRequest{PricePerGB: ptr.Of(0.01)}
		_, err = Default.UpdateHandler(ctx, db, 1, req)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "DDO schedules do not accept")
	})
}

func TestUpdateHandler_DDORejectsPricePerDeal(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Schedule{
			Preparation: &model.Preparation{},
			DealType:    model.DealTypeDDO,
			URLTemplate: "https://example.com/{PIECE_CID}",
		}).Error
		require.NoError(t, err)
		req := UpdateRequest{PricePerDeal: ptr.Of(0.1)}
		_, err = Default.UpdateHandler(ctx, db, 1, req)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "DDO schedules do not accept")
	})
}

// switching a market schedule to DDO with a non-zero price flag in the same
// PATCH must be rejected -- effective deal type after the update is DDO.
func TestUpdateHandler_DDOSwitchWithPriceRejected(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Schedule{
			Preparation: &model.Preparation{},
			DealType:    model.DealTypeMarket,
			URLTemplate: "https://example.com/{PIECE_CID}",
		}).Error
		require.NoError(t, err)
		req := UpdateRequest{
			DealType:        ptr.Of(string(model.DealTypeDDO)),
			PricePerGBEpoch: ptr.Of(1e-18),
		}
		_, err = Default.UpdateHandler(ctx, db, 1, req)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "DDO schedules do not accept")
	})
}

// explicitly setting price flags to zero on a DDO schedule is allowed -- it
// is a no-op write, not an attempt to override the SP price.
func TestUpdateHandler_DDOZeroPriceAllowed(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Schedule{
			Preparation: &model.Preparation{},
			DealType:    model.DealTypeDDO,
			URLTemplate: "https://example.com/{PIECE_CID}",
		}).Error
		require.NoError(t, err)
		req := UpdateRequest{
			PricePerGBEpoch: ptr.Of(0.0),
			PricePerGB:      ptr.Of(0.0),
			PricePerDeal:    ptr.Of(0.0),
		}
		_, err = Default.UpdateHandler(ctx, db, 1, req)
		require.NoError(t, err)
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
