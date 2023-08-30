package cmd

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/google/uuid"
	"github.com/gotidy/ptr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
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
	mockHandler := new(MockDataPrep)
	defer swapDataPrepHandler(mockHandler)()

	mockHandler.On("CreatePreparationHandler", mock.Anything, mock.Anything, mock.Anything).Return(&testPreparation, nil)
	out, _, err := Run(context.Background(), "singularity prep create --source source --output output")
	require.NoError(t, err)
	Save(t, out, "dataprep_create.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep create --source source --output output")
	require.NoError(t, err)
	Save(t, out, "dataprep_create_verbose.txt")
}

func TestDataPrepListHandler(t *testing.T) {
	mockHandler := new(MockDataPrep)
	defer swapDataPrepHandler(mockHandler)()

	mockHandler.On("ListHandler", mock.Anything, mock.Anything).Return([]model.Preparation{testPreparation}, nil)
	out, _, err := Run(context.Background(), "singularity prep list")
	require.NoError(t, err)
	Save(t, out, "dataprep_create.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep list")
	require.NoError(t, err)
	Save(t, out, "dataprep_create_verbose.txt")
}

var testDagGenJob = model.Job{
	ID:           1,
	Type:         model.DagGen,
	State:        model.Ready,
	WorkerID:     nil,
	AttachmentID: 1,
}

func TestDataPrepStartDagGenHandler(t *testing.T) {
	mockHandler := new(MockDataPrep)
	defer swapDataPrepHandler(mockHandler)()

	mockHandler.On("StartDagGenHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&testDagGenJob, nil)
	out, _, err := Run(context.Background(), "singularity prep start-daggen 1 name")
	require.NoError(t, err)
	Save(t, out, "dataprep_daggen_job.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep start-daggen 1 name")
	require.NoError(t, err)
	Save(t, out, "dataprep_daggen_job_verbose.txt")
}

func TestDataPrepPauseDagGenHandler(t *testing.T) {
	mockHandler := new(MockDataPrep)
	defer swapDataPrepHandler(mockHandler)()

	mockHandler.On("PauseDagGenHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&testDagGenJob, nil)
	out, _, err := Run(context.Background(), "singularity prep pause-daggen 1 name")
	require.NoError(t, err)
	Save(t, out, "dataprep_daggen_job.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep pause-daggen 1 name")
	require.NoError(t, err)
	Save(t, out, "dataprep_daggen_job_verbose.txt")
}

var testScanob = model.Job{
	ID:           1,
	Type:         model.Scan,
	State:        model.Ready,
	WorkerID:     nil,
	AttachmentID: 1,
}

func TestDataPrepStartScanHandler(t *testing.T) {
	mockHandler := new(MockDataPrep)
	defer swapDataPrepHandler(mockHandler)()

	mockHandler.On("StartScanHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&testScanob, nil)
	out, _, err := Run(context.Background(), "singularity prep start-scan 1 name")
	require.NoError(t, err)
	Save(t, out, "dataprep_scan_job.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep start-scan 1 name")
	require.NoError(t, err)
	Save(t, out, "dataprep_scan_job_verbose.txt")
}

func TestDataPrepPauseScanHandler(t *testing.T) {
	mockHandler := new(MockDataPrep)
	defer swapDataPrepHandler(mockHandler)()

	mockHandler.On("PauseScanHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&testScanob, nil)
	out, _, err := Run(context.Background(), "singularity prep pause-scan 1 name")
	require.NoError(t, err)
	Save(t, out, "dataprep_scan_job.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep pause-scan 1 name")
	require.NoError(t, err)
	Save(t, out, "dataprep_scan_job_verbose.txt")
}

var testPackJob = model.Job{
	ID:           1,
	Type:         model.Pack,
	State:        model.Ready,
	WorkerID:     nil,
	AttachmentID: 1,
}

func TestDataPrepStartPackHandler(t *testing.T) {
	mockHandler := new(MockDataPrep)
	defer swapDataPrepHandler(mockHandler)()

	mockHandler.On("StartPackHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]model.Job{testPackJob}, nil)
	out, _, err := Run(context.Background(), "singularity prep start-pack 1 name")
	require.NoError(t, err)
	Save(t, out, "dataprep_pack_job.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep start-pack 1 name")
	require.NoError(t, err)
	Save(t, out, "dataprep_pack_job_verbose.txt")
}

func TestDataPrepPausePackHandler(t *testing.T) {
	mockHandler := new(MockDataPrep)
	defer swapDataPrepHandler(mockHandler)()

	mockHandler.On("PausePackHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]model.Job{testPackJob}, nil)
	out, _, err := Run(context.Background(), "singularity prep pause-pack 1 name")
	require.NoError(t, err)
	Save(t, out, "dataprep_pack_job.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep pause-pack 1 name")
	require.NoError(t, err)
	Save(t, out, "dataprep_pack_job_verbose.txt")
}

func TestDataPrepAttachSourceHandler(t *testing.T) {
	mockHandler := new(MockDataPrep)
	defer swapDataPrepHandler(mockHandler)()

	mockHandler.On("AddSourceStorageHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&testPreparation, nil)
	out, _, err := Run(context.Background(), "singularity prep attach-source 1 source")
	require.NoError(t, err)
	Save(t, out, "dataprep_create.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep attach-source 1 source")
	require.NoError(t, err)
	Save(t, out, "dataprep_create_verbose.txt")
}

func TestDataPrepAttachOutputHandler(t *testing.T) {
	mockHandler := new(MockDataPrep)
	defer swapDataPrepHandler(mockHandler)()

	mockHandler.On("AddOutputStorageHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&testPreparation, nil)
	out, _, err := Run(context.Background(), "singularity prep attach-output 1 source")
	require.NoError(t, err)
	Save(t, out, "dataprep_create.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep attach-output 1 source")
	require.NoError(t, err)
	Save(t, out, "dataprep_create_verbose.txt")
}
func TestDataPrepDetachOutputHandler(t *testing.T) {
	mockHandler := new(MockDataPrep)
	defer swapDataPrepHandler(mockHandler)()

	mockHandler.On("RemoveOutputStorageHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&testPreparation, nil)
	out, _, err := Run(context.Background(), "singularity prep detach-output 1 source")
	require.NoError(t, err)
	Save(t, out, "dataprep_create.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep detach-output 1 source")
	require.NoError(t, err)
	Save(t, out, "dataprep_create_verbose.txt")
}

func TestDataPreparationAttachWalletHandler(t *testing.T) {
	mockHandler := new(MockWallet)
	defer swapWalletHandler(mockHandler)()

	mockHandler.On("AttachHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&testPreparation, nil)
	out, _, err := Run(context.Background(), "singularity prep attach-wallet 1 test")
	require.NoError(t, err)
	Save(t, out, "dataprep_create.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep attach-wallet 1 test")
	require.NoError(t, err)
	Save(t, out, "dataprep_create_verbose.txt")
}

func TestDataPreparationDetachWalletHandler(t *testing.T) {
	mockHandler := new(MockWallet)
	defer swapWalletHandler(mockHandler)()

	mockHandler.On("DetachHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&testPreparation, nil)
	out, _, err := Run(context.Background(), "singularity prep detach-wallet 1 test")
	require.NoError(t, err)
	Save(t, out, "dataprep_create.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep detach-wallet 1 test")
	require.NoError(t, err)
	Save(t, out, "dataprep_create_verbose.txt")
}

func TestDataPreparationListAttachedWalletHandler(t *testing.T) {
	mockHandler := new(MockWallet)
	defer swapWalletHandler(mockHandler)()

	mockHandler.On("ListAttachedHandler", mock.Anything, mock.Anything, mock.Anything).Return(testPreparation.Wallets, nil)
	out, _, err := Run(context.Background(), "singularity prep list-wallets 1")
	require.NoError(t, err)
	Save(t, out, "dataprep_list_wallets.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep list-wallets 1")
	require.NoError(t, err)
	Save(t, out, "dataprep_list_wallets_verbose.txt")
}

func TestDataPreparationExploreHandler(t *testing.T) {
	mockHandler := new(MockDataPrep)
	defer swapDataPrepHandler(mockHandler)()

	mockHandler.On("ExploreHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]dataprep.DirEntry{
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
	}, nil)
	out, _, err := Run(context.Background(), "singularity prep explore 1 storage path")
	require.NoError(t, err)
	Save(t, out, "dataprep_explore.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep explore 1 storage path")
	require.NoError(t, err)
	Save(t, out, "dataprep_explore_verbose.txt")
}

func TestDataPreparationAddPieceHandler(t *testing.T) {
	mockHandler := new(MockDataPrep)
	defer swapDataPrepHandler(mockHandler)()

	mockHandler.On("AddPieceHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&model.Car{
		ID:            1,
		CreatedAt:     time.Time{},
		PieceCID:      model.CID(testutil.TestCid),
		PieceSize:     100,
		RootCID:       model.CID{},
		FileSize:      100,
		StorageID:     ptr.Of(uint32(1)),
		StoragePath:   "test1.car",
		PreparationID: 1,
	}, nil)
	out, _, err := Run(context.Background(), "singularity prep add-piece --piece-cid xxx --piece-size 100 1")
	require.NoError(t, err)
	Save(t, out, "dataprep_add_piece.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep add-piece --piece-cid xxx --piece-size 100 1")
	require.NoError(t, err)
	Save(t, out, "dataprep_add_piece_verbose.txt")
}

func TestDataPreparationListPiecesHandler(t *testing.T) {
	mockHandler := new(MockDataPrep)
	defer swapDataPrepHandler(mockHandler)()

	mockHandler.On("ListPiecesHandler", mock.Anything, mock.Anything, mock.Anything).Return([]dataprep.PieceList{
		{
			SourceStorageID: ptr.Of(uint32(1)),
			AttachmentID:    ptr.Of(uint32(1)),
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
				StorageID:     ptr.Of(uint32(1)),
				StoragePath:   "test1.car",
				PreparationID: 1,
			}, {
				ID:            2,
				CreatedAt:     time.Time{},
				PieceCID:      model.CID(testutil.TestCid),
				PieceSize:     300,
				FileSize:      400,
				StorageID:     ptr.Of(uint32(1)),
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
				StorageID:     ptr.Of(uint32(1)),
				StoragePath:   "test3.car",
				PreparationID: 1,
			}, {
				ID:            4,
				CreatedAt:     time.Time{},
				PieceCID:      model.CID(testutil.TestCid),
				PieceSize:     700,
				FileSize:      800,
				StorageID:     ptr.Of(uint32(1)),
				StoragePath:   "test4.car",
				PreparationID: 1,
			}},
		},
	}, nil)
	out, _, err := Run(context.Background(), "singularity prep list-pieces 1")
	require.NoError(t, err)
	Save(t, out, "dataprep_list_pieces.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep list-pieces 1")
	require.NoError(t, err)
	Save(t, out, "dataprep_list_pieces_verbose.txt")
}

func TestDataPreparationGetStatusHandler(t *testing.T) {
	mockHandler := new(MockDataPrep)
	defer swapDataPrepHandler(mockHandler)()

	mockHandler.On("GetStatusHandler", mock.Anything, mock.Anything, mock.Anything).Return([]dataprep.SourceStatus{
		{
			AttachmentID:    ptr.Of(uint32(1)),
			SourceStorageID: ptr.Of(uint32(1)),
			SourceStorage: &model.Storage{
				ID:        1,
				Name:      "source",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				Type:      "local",
				Path:      "/tmp",
			},
			Jobs: []model.Job{
				{
					ID:           1,
					Type:         model.Pack,
					State:        model.Processing,
					WorkerID:     ptr.Of(uuid.NewString()),
					AttachmentID: 1,
				},
			},
		},
	}, nil)
	out, _, err := Run(context.Background(), "singularity prep status 1")
	require.NoError(t, err)
	Save(t, out, "dataprep_status.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep status 1")
	require.NoError(t, err)
	Save(t, out, "dataprep_status_verbose.txt")
}
