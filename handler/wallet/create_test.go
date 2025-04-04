package wallet

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestCreateHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		lotusClient := util.NewLotusClient("https://api.node.glif.io/rpc/v0", "")

		t.Run("success", func(t *testing.T) {
			w, err := Default.CreateHandler(ctx, db, lotusClient)
			require.NoError(t, err)
			require.NotEmpty(t, w.Address)
			require.NotEmpty(t, w.PrivateKey)

			_, err = Default.ImportHandler(ctx, db, lotusClient, ImportRequest{
				PrivateKey: w.PrivateKey,
			})
			require.ErrorIs(t, err, handlererror.ErrDuplicateRecord)
		})

		t.Run("invalid response", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			lotusClient := util.NewLotusClient("http://127.0.0.1", "")
			_, err := Default.CreateHandler(ctx, db, lotusClient)
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		})
	})
}
