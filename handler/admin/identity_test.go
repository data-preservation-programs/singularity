package admin

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/util/testutil"
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
	})
}
