package cmd

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/gotidy/ptr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

var testPreparation = model.Preparation{
	ID:                1,
	CreatedAt:         time.Time{},
	UpdatedAt:         time.Time{},
	DeleteAfterExport: false,
	MaxSize:           100,
	PieceSize:         200,
	Wallets: []model.Wallet{{
		ID:         "client_id",
		Address:    "client_address",
		PrivateKey: "private_key",
	}},
	SourceStorages: []model.Storage{{
		ID:        1,
		Name:      "source",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Type:      "local",
		Path:      "/tmp/source",
	}},
	OutputStorages: []model.Storage{{
		ID:        2,
		Name:      "output1",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Type:      "local",
		Path:      "/tmp/output1",
	}, {
		ID:        3,
		Name:      "output2",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Type:      "local",
		Path:      "/tmp/output2",
	}},
}

func swapDataPrepHandler(mockHandler dataprep.Handler) func() {
	actual := dataprep.Default
	dataprep.Default = mockHandler
	return func() {
		dataprep.Default = actual
	}
}

func TestDataPrepCreateHandler(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(dataprep.MockDataPrep)
		defer swapDataPrepHandler(mockHandler)()

		mockHandler.On("CreatePreparationHandler", mock.Anything, mock.Anything, mock.Anything).Return(&testPreparation, nil)
		_, _, err := runner.Run(ctx, "singularity prep create --source source --output output")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose prep create --source source --output output")
		require.NoError(t, err)
	})

}
func TestDataPrepCreateHandler_WithStorage(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(dataprep.MockDataPrep)
		defer swapDataPrepHandler(mockHandler)()

		mockHandler.On("CreatePreparationHandler", mock.Anything, mock.Anything, mock.Anything).Return(&testPreparation, nil)
		output := t.TempDir()
		_, _, err := runner.Run(ctx, fmt.Sprintf("singularity prep create --source source --local-output %s",
			testutil.EscapePath(output)))
		require.NoError(t, err)
	})
}

func TestDataPrepListHandler(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(dataprep.MockDataPrep)
		defer swapDataPrepHandler(mockHandler)()

		mockHandler.On("ListHandler", mock.Anything, mock.Anything).Return([]model.Preparation{testPreparation}, nil)
		_, _, err := runner.Run(ctx, "singularity prep list")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose prep list")
		require.NoError(t, err)
	})
}

func TestDataPrepAttachSourceHandler(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(dataprep.MockDataPrep)
		defer swapDataPrepHandler(mockHandler)()

		mockHandler.On("AddSourceStorageHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&testPreparation, nil)
		_, _, err := runner.Run(ctx, "singularity prep attach-source 1 source")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose prep attach-source 1 source")
		require.NoError(t, err)
	})
}

func TestDataPrepAttachOutputHandler(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(dataprep.MockDataPrep)
		defer swapDataPrepHandler(mockHandler)()

		mockHandler.On("AddOutputStorageHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&testPreparation, nil)
		_, _, err := runner.Run(ctx, "singularity prep attach-output 1 source")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose prep attach-output 1 source")
		require.NoError(t, err)
	})
}
func TestDataPrepDetachOutputHandler(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(dataprep.MockDataPrep)
		defer swapDataPrepHandler(mockHandler)()

		mockHandler.On("RemoveOutputStorageHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&testPreparation, nil)
		_, _, err := runner.Run(ctx, "singularity prep detach-output 1 source")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose prep detach-output 1 source")
		require.NoError(t, err)
	})
}

func TestDataPreparationAttachWalletHandler(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(wallet.MockWallet)
		defer swapWalletHandler(mockHandler)()

		mockHandler.On("AttachHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&testPreparation, nil)
		_, _, err := runner.Run(ctx, "singularity prep attach-wallet 1 test")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose prep attach-wallet 1 test")
		require.NoError(t, err)
	})
}

func TestDataPreparationDetachWalletHandler(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(wallet.MockWallet)
		defer swapWalletHandler(mockHandler)()

		mockHandler.On("DetachHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&testPreparation, nil)
		_, _, err := runner.Run(ctx, "singularity prep detach-wallet 1 test")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose prep detach-wallet 1 test")
		require.NoError(t, err)
	})
}

func TestDataPreparationListAttachedWalletHandler(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(wallet.MockWallet)
		defer swapWalletHandler(mockHandler)()

		mockHandler.On("ListAttachedHandler", mock.Anything, mock.Anything, mock.Anything).Return(testPreparation.Wallets, nil)
		_, _, err := runner.Run(ctx, "singularity prep list-wallets 1")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose prep list-wallets 1")
		require.NoError(t, err)
	})
}

func TestDataPreparationExploreHandler(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(dataprep.MockDataPrep)
		defer swapDataPrepHandler(mockHandler)()

		mockHandler.On("ExploreHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&dataprep.ExploreResult{SubEntries: []dataprep.DirEntry{
			{
				Path:  "file1",
				IsDir: false,
				CID:   "cid1",
				FileVersions: []dataprep.Version{{
					ID:           1,
					CID:          "cid1",
					Hash:         "hash1",
					Size:         100,
					LastModified: time.Time{},
				}, {
					ID:           2,
					CID:          "cid2",
					Hash:         "hash2",
					Size:         200,
					LastModified: time.Time{},
				}},
			},
			{
				Path:  "dir",
				IsDir: true,
				CID:   "cid4",
			},
			{
				Path:  "file2",
				IsDir: false,
				CID:   "cid3",
				FileVersions: []dataprep.Version{{
					ID:           1,
					CID:          "cid3",
					Hash:         "hash3",
					Size:         300,
					LastModified: time.Time{},
				}},
			},
		},
			Path: "/",
			CID:  "root_cid",
		}, nil)
		_, _, err := runner.Run(ctx, "singularity prep explore 1 storage path")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose prep explore 1 storage path")
		require.NoError(t, err)
	})
}

func TestDataPreparationAddPieceHandler(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(dataprep.MockDataPrep)
		defer swapDataPrepHandler(mockHandler)()

		mockHandler.On("AddPieceHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&model.Car{
			ID:            1,
			CreatedAt:     time.Time{},
			PieceCID:      model.CID(testutil.TestCid),
			PieceSize:     100,
			RootCID:       model.CID{},
			FileSize:      100,
			StorageID:     ptr.Of(model.StorageID(1)),
			StoragePath:   "test1.car",
			PreparationID: 1,
		}, nil)
		_, _, err := runner.Run(ctx, "singularity prep add-piece --piece-cid xxx --piece-size 100 --file-size 100 1")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose prep add-piece --piece-cid xxx --piece-size 100 --file-size 100 1")
		require.NoError(t, err)
	})
}

func TestDataPreparationListPiecesHandler(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(dataprep.MockDataPrep)
		defer swapDataPrepHandler(mockHandler)()

		mockHandler.On("ListPiecesHandler", mock.Anything, mock.Anything, mock.Anything).Return([]dataprep.PieceList{
			{
				SourceStorageID: ptr.Of(model.StorageID(1)),
				AttachmentID:    ptr.Of(model.SourceAttachmentID(1)),
				SourceStorage: &model.Storage{
					ID:        1,
					Name:      "local",
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
					Type:      "local",
					Path:      "/tmp",
				},
				Pieces: []model.Car{{
					ID:            1,
					CreatedAt:     time.Time{},
					PieceCID:      model.CID(testutil.TestCid),
					PieceSize:     100,
					FileSize:      200,
					StorageID:     ptr.Of(model.StorageID(1)),
					StoragePath:   "test1.car",
					PreparationID: 1,
				}, {
					ID:            2,
					CreatedAt:     time.Time{},
					PieceCID:      model.CID(testutil.TestCid),
					PieceSize:     300,
					FileSize:      400,
					StorageID:     ptr.Of(model.StorageID(1)),
					StoragePath:   "test2.car",
					PreparationID: 1,
				}},
			},
			{
				Pieces: []model.Car{{
					ID:            3,
					CreatedAt:     time.Time{},
					PieceCID:      model.CID(testutil.TestCid),
					PieceSize:     500,
					FileSize:      600,
					StorageID:     ptr.Of(model.StorageID(1)),
					StoragePath:   "test3.car",
					PreparationID: 1,
				}, {
					ID:            4,
					CreatedAt:     time.Time{},
					PieceCID:      model.CID(testutil.TestCid),
					PieceSize:     700,
					FileSize:      800,
					StorageID:     ptr.Of(model.StorageID(1)),
					StoragePath:   "test4.car",
					PreparationID: 1,
				}},
			},
		}, nil)
		_, _, err := runner.Run(ctx, "singularity prep list-pieces 1")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose prep list-pieces 1")
		require.NoError(t, err)
	})
}
