package cmd

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func swapWalletHandler(mockHandler wallet.Handler) func() {
	actual := wallet.Default
	wallet.Default = mockHandler
	return func() {
		wallet.Default = actual
	}
}

func TestWalletCreate(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		var err error
		err = model.GetMigrator(db).Migrate()
		require.NoError(t, err)

		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(wallet.MockWallet)
		defer swapWalletHandler(mockHandler)()
		mockHandler.On("CreateHandler", mock.Anything, mock.Anything, mock.Anything).Return(&model.Wallet{
			ActorID:    "id",
			Address:    "address",
			PrivateKey: "private",
		}, nil)

		_, _, err = runner.Run(ctx, "singularity wallet create")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose wallet create")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity wallet create secp256k1")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity wallet create bls")
		require.NoError(t, err)
	})
}

func TestWalletCreate_BadType(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		var err error
		err = model.GetMigrator(db).Migrate()
		require.NoError(t, err)

		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(wallet.MockWallet)
		defer swapWalletHandler(mockHandler)()
		mockHandler.On("CreateHandler", mock.Anything, mock.Anything, mock.Anything).Return((*model.Wallet)(nil), errors.New("unsupported key type: not-a-real-type"))

		_, _, err = runner.Run(ctx, "singularity wallet create not-a-real-type")
		require.Error(t, err)
	})
}

func TestWalletImport(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		var err error
		err = model.GetMigrator(db).Migrate()
		require.NoError(t, err)

		tmp := t.TempDir()
		err = os.WriteFile(filepath.Join(tmp, "private"), []byte("private"), 0644)
		require.NoError(t, err)

		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(wallet.MockWallet)
		defer swapWalletHandler(mockHandler)()
		mockHandler.On("ImportHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&model.Wallet{
			ActorID:     "id",
			Address:     "address",
			PrivateKey:  "private",
			ActorName:   "Test Actor",
			ContactInfo: "test@example.com",
			Location:    "Test Location",
		}, nil)

		// Test basic import
		_, _, err = runner.Run(ctx, "singularity wallet import "+testutil.EscapePath(filepath.Join(tmp, "private")))
		require.NoError(t, err)

		// Test import with metadata flags
		_, _, err = runner.Run(ctx, "singularity wallet import --name 'Test Actor' --contact 'test@example.com' --location 'Test Location' "+testutil.EscapePath(filepath.Join(tmp, "private")))
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose wallet import "+testutil.EscapePath(filepath.Join(tmp, "private")))
		require.NoError(t, err)
	})
}

func TestWalletInit(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		var err error
		err = model.GetMigrator(db).Migrate()
		require.NoError(t, err)

		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(wallet.MockWallet)
		defer swapWalletHandler(mockHandler)()
		mockHandler.On("InitHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&model.Wallet{
			ActorID:    "id",
			Address:    "address",
			PrivateKey: "private",
		}, nil)

		_, _, err = runner.Run(ctx, "singularity wallet init xxx")
		require.NoError(t, err)
	})
}

func TestWalletListWithBalance(t *testing.T) {
	t.Run("Lotus error", func(t *testing.T) {
		testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			var err error
			err = model.GetMigrator(db).Migrate()
			require.NoError(t, err)

			runner := NewRunner()
			defer runner.Save(t)
			mockHandler := new(wallet.MockWallet)
			defer swapWalletHandler(mockHandler)()
			mockHandler.On("ListWithBalanceHandler", mock.Anything, mock.Anything, mock.Anything).Return([]wallet.WalletWithBalance{
				{
					Wallet: model.Wallet{
						ActorID:    "id1",
						Address:    "f1abcde7zd3lfsv43aj2kb454ymaqw7debhumjxyz",
						PrivateKey: "private1",
					},
					Balance:        "1 FIL",
					BalanceAttoFIL: "1000000000000000000",
					DataCap:        "1 TiB",
					DataCapBytes:   1099511627776,
				},
			}, nil)
			_, _, err = runner.Run(context.Background(), "singularity wallet list --with-balance --lotus-api http://mock --lotus-token mock")
			require.NoError(t, err)
		})
	})

	t.Run("Partial Lotus error", func(t *testing.T) {
		testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			var err error
			err = model.GetMigrator(db).Migrate()
			require.NoError(t, err)

			runner := NewRunner()
			defer runner.Save(t)
			mockHandler := new(wallet.MockWallet)
			defer swapWalletHandler(mockHandler)()
			errMsg := "lotus error"
			mockHandler.On("ListWithBalanceHandler", mock.Anything, mock.Anything, mock.Anything).Return([]wallet.WalletWithBalance{
				{
					Wallet: model.Wallet{
						ActorID:    "id1",
						Address:    "f1abcde7zd3lfsv43aj2kb454ymaqw7debhumjxyz",
						PrivateKey: "private1",
					},
					Error: &errMsg,
				}, 
				{
					Wallet: model.Wallet{
						ActorID:    "id2",
						Address:    "f1fib3pv7jua2ockdugtz7viz3cyy6lkhh7rfx3sa",
						PrivateKey: "private2",
					},
					Balance:        "1 FIL",
					BalanceAttoFIL: "1000000000000000000",
					DataCap:        "1 TiB",
					DataCapBytes:   1099511627776,
				}}, nil)
			_, _, err = runner.Run(context.Background(), "singularity wallet list --with-balance --lotus-api http://mock --lotus-token mock")
			require.NoError(t, err)
		})
	})

	t.Run("Missing Lotus info", func(t *testing.T) {
		testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			var err error
			err = model.GetMigrator(db).Migrate()
			require.NoError(t, err)

			runner := NewRunner()
			defer runner.Save(t)
			mockHandler := new(wallet.MockWallet)
			defer swapWalletHandler(mockHandler)()
			mockHandler.On("ListWithBalanceHandler", mock.Anything, mock.Anything, mock.Anything).Return([]wallet.WalletWithBalance{
				{
					Wallet: model.Wallet{
						ActorID:    "id1",
						Address:    "f1abcde7zd3lfsv43aj2kb454ymaqw7debhumjxyz",
						PrivateKey: "private1",
					},
					Balance:        "1 FIL",
					BalanceAttoFIL: "1000000000000000000",
					DataCap:        "1 TiB",
					DataCapBytes:   1099511627776,
				}}, nil)
			_, _, err = runner.Run(context.Background(), "singularity wallet list --with-balance")
			require.Error(t, err)
			require.Contains(t, err.Error(), "Both --lotus-api and --lotus-token must be provided")
		})
	})
}

func TestWalletList(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		var err error
		err = model.GetMigrator(db).Migrate()
		require.NoError(t, err)

		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(wallet.MockWallet)
		defer swapWalletHandler(mockHandler)()
		mockHandler.On("ListHandler", mock.Anything, mock.Anything).Return([]model.Wallet{{
			ActorID:    "id1",
			Address:    "address1",
			PrivateKey: "private1",
		}, {
			ActorID:    "id2",
			Address:    "address2",
			PrivateKey: "private2",
		}}, nil)

		_, _, err = runner.Run(ctx, "singularity wallet list")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose wallet list")
		require.NoError(t, err)
	})
}

func TestWalletRemove(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		var err error
		err = model.GetMigrator(db).Migrate()
		require.NoError(t, err)

		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(wallet.MockWallet)
		defer swapWalletHandler(mockHandler)()
		mockHandler.On("RemoveHandler", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		_, _, err = runner.Run(ctx, "singularity wallet remove --really-do-it xxx")
		require.NoError(t, err)
	})
}

func TestWalletRemove_NoReallyDoIt(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		var err error
		err = model.GetMigrator(db).Migrate()
		require.NoError(t, err)

		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(wallet.MockWallet)
		defer swapWalletHandler(mockHandler)()
		mockHandler.On("RemoveHandler", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		_, _, err = runner.Run(ctx, "singularity wallet remove xxx")
		require.ErrorIs(t, err, cliutil.ErrReallyDoIt)
	})
}

func TestWalletUpdate(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		var err error
		err = model.GetMigrator(db).Migrate()
		require.NoError(t, err)

		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(wallet.MockWallet)
		defer swapWalletHandler(mockHandler)()
		mockHandler.On("UpdateHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&model.Wallet{
			ActorID:     "id",
			ActorName:   "Updated Name",
			Address:     "address",
			ContactInfo: "test@example.com",
			Location:    "US-East",
			WalletType:  model.SPWallet,
		}, nil)

		_, _, err = runner.Run(ctx, "singularity wallet update --name Updated --contact test@example.com --location US-East address")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose wallet update --name Updated address")
		require.NoError(t, err)
	})
}

func TestWalletUpdate_NoAddress(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		var err error
		err = model.GetMigrator(db).Migrate()
		require.NoError(t, err)

		runner := NewRunner()
		defer runner.Save(t)

		_, _, err = runner.Run(ctx, "singularity wallet update --name Test")
		require.Error(t, err)
		require.Contains(t, err.Error(), "incorrect number of arguments")
	})
}

func TestWalletUpdate_NoFields(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		var err error
		err = model.GetMigrator(db).Migrate()
		require.NoError(t, err)

		runner := NewRunner()
		defer runner.Save(t)

		_, _, err = runner.Run(ctx, "singularity wallet update address")
		require.Error(t, err)
		require.Contains(t, err.Error(), "at least one field must be provided for update")
	})
}
