package wallet

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestRemoveHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		t.Run("success", func(t *testing.T) {
			err := db.Create(&model.Actor{
				ID: "test",
			}).Error
			require.NoError(t, err)
			err = Default.RemoveHandler(ctx, db, "test")
			require.NoError(t, err)
			err = Default.RemoveHandler(ctx, db, "test")
			require.ErrorIs(t, err, handlererror.ErrNotFound)
		})
	})
}
