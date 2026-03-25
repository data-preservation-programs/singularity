package sppool

import (
	"context"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestCreateHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		t.Run("success", func(t *testing.T) {
			pool, err := Default.CreateHandler(ctx, db, CreateRequest{
				Name:       "test-pool",
				StartDelay: "72h",
				Duration:   "12840h",
			})
			require.NoError(t, err)
			require.Equal(t, "test-pool", pool.Name)
		})

		t.Run("duplicate name", func(t *testing.T) {
			_, err := Default.CreateHandler(ctx, db, CreateRequest{
				Name:       "test-pool",
				StartDelay: "72h",
				Duration:   "12840h",
			})
			require.True(t, errors.Is(err, handlererror.ErrDuplicateRecord))
		})

		t.Run("empty name", func(t *testing.T) {
			_, err := Default.CreateHandler(ctx, db, CreateRequest{
				Name:       "",
				StartDelay: "72h",
				Duration:   "12840h",
			})
			require.True(t, errors.Is(err, handlererror.ErrInvalidParameter))
		})

		t.Run("cron without deal number", func(t *testing.T) {
			_, err := Default.CreateHandler(ctx, db, CreateRequest{
				Name:         "cron-pool",
				StartDelay:   "72h",
				Duration:     "12840h",
				ScheduleCron: "0 * * * *",
			})
			require.True(t, errors.Is(err, handlererror.ErrInvalidParameter))
		})
	})
}
