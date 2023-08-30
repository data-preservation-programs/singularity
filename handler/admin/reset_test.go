package admin

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestResetHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		require.NoError(t, Default.ResetHandler(ctx, db))
	})
}
