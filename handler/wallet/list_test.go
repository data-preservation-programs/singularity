package wallet

import (
	"context"
	"github.com/cockroachdb/errors"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestListHandler(t *testing.T) {
// Skip this test if MySQL is not available
if testing.Short() {
	t.Skip("Skipping due to missing MySQL (integration test)")
}
testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
	t.Run("success", func(t *testing.T) {
		err := db.Create(&model.Wallet{}).Error
		require.NoError(t, err)
		wallets, err := Default.ListHandler(ctx, db)
		require.NoError(t, err)
		require.Len(t, wallets, 1)
	})
})
}


func TestListWithBalanceHandler(t *testing.T) {
   // Skip this test if MySQL is not available
   if testing.Short() {
	   t.Skip("Skipping due to missing MySQL (integration test)")
   }
   t.Run("all Lotus errors", func(t *testing.T) {
		   testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
					   w := model.Wallet{Address: "f1abcde7zd3lfsv43aj2kb454ymaqw7debhumjxyz"}
					   err := db.Create(&w).Error
					   require.NoError(t, err)
					   mockLotus := testutil.NewMockLotusClient()
					   mockLotus.SetError("Filecoin.WalletBalance", errors.New("lotus unreachable"))
					   mockLotus.SetError("Filecoin.StateVerifiedClientStatus", errors.New("lotus unreachable"))
					   wallets, err := ListWithBalanceHandler(ctx, db, mockLotus)
					   require.Error(t, err)
					   require.Len(t, wallets, 1)
					   require.NotNil(t, wallets[0].Error)
					   require.Contains(t, *wallets[0].Error, "lotus unreachable")
			   })
	   })

	   t.Run("partial Lotus errors", func(t *testing.T) {
			   testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
					   w1 := model.Wallet{Address: "f1abcde7zd3lfsv43aj2kb454ymaqw7debhumjxyz"}
					   w2 := model.Wallet{Address: "f1fib3pv7jua2ockdugtz7viz3cyy6lkhh7rfx3sa"}
					   err := db.Create(&w1).Error
					   require.NoError(t, err)
					   err = db.Create(&w2).Error
					   require.NoError(t, err)
					   mockLotus := testutil.NewMockLotusClient()
					   mockLotus.SetResponse("Filecoin.WalletBalance", "1000000000000000000")
					   mockLotus.SetResponse("Filecoin.StateVerifiedClientStatus", "0")
					   mockLotus.SetError("Filecoin.WalletBalance", errors.New("lotus error"))
					   wallets, err := ListWithBalanceHandler(ctx, db, mockLotus)
					   require.NoError(t, err)
					   require.Len(t, wallets, 2)
					   foundError := false
					   for _, w := range wallets {
							   if w.Error != nil {
									   foundError = true
							   }
					   }
					   require.True(t, foundError)
			   })
	   })

	   t.Run("Lotus timeout", func(t *testing.T) {
			   testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
					   w := model.Wallet{Address: "f1abcde7zd3lfsv43aj2kb454ymaqw7debhumjxyz"}
					   err := db.Create(&w).Error
					   require.NoError(t, err)
					   mockLotus := testutil.NewMockLotusClient()
					   mockLotus.SetError("Filecoin.WalletBalance", context.DeadlineExceeded)
					   wallets, err := ListWithBalanceHandler(ctx, db, mockLotus)
					   require.Error(t, err)
					   require.Len(t, wallets, 1)
					   require.NotNil(t, wallets[0].Error)
					   require.Contains(t, *wallets[0].Error, "context deadline exceeded")
			   })
	   })
	   t.Run("empty db", func(t *testing.T) {
			   testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
					   mockLotus := testutil.NewMockLotusClient()
					   wallets, err := ListWithBalanceHandler(ctx, db, mockLotus)
					   require.NoError(t, err)
					   require.Len(t, wallets, 0)
			   })
	   })

	   t.Run("db error", func(t *testing.T) {
			   ctx := context.Background()
			   badDB, _ := gorm.Open(nil, &gorm.Config{})
			   mockLotus := testutil.NewMockLotusClient()
			   wallets, err := ListWithBalanceHandler(ctx, badDB, mockLotus)
			   require.Error(t, err)
			   require.Nil(t, wallets)
	   })
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Insert a wallet with a different valid address
		w := model.Wallet{Address: "f1abcde7zd3lfsv43aj2kb454ymaqw7debhumjxyz"}
		err := db.Create(&w).Error
		require.NoError(t, err)

		// Set up mock Lotus client
		mockLotus := testutil.NewMockLotusClient()
		mockLotus.SetResponse("Filecoin.WalletBalance", "1000000000000000000") // 1 FIL
		mockLotus.SetResponse("Filecoin.StateVerifiedClientStatus", "0")

		wallets, err := ListWithBalanceHandler(ctx, db, mockLotus)
		require.NoError(t, err)
		require.Len(t, wallets, 1)
		require.Equal(t, "1.000000 FIL", wallets[0].Balance)
		require.Equal(t, "1000000000000000000", wallets[0].BalanceAttoFIL)
		require.Equal(t, "0", wallets[0].DataCap)
	})
}
