package admin

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/rjNemo/underscore"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestSetIdentityHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		require.NoError(t, Default.SetIdentityHandler(ctx, db, SetIdentityRequest{
			Identity: "test1",
		}))
		require.NoError(t, Default.SetIdentityHandler(ctx, db, SetIdentityRequest{
			Identity: "test2",
		}))
		var globals []model.Global
		require.NoError(t, db.WithContext(ctx).Find(&globals).Error)
		found, err := underscore.Find(globals, func(global model.Global) bool {
			return global.Key == "identity"
		})
		require.NoError(t, err)
		require.Equal(t, "test2", found.Value)
	})
}
