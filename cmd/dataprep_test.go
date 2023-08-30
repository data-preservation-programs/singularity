package cmd

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/data-preservation-programs/singularity/model"
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
	Wallets:           nil,
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
	CompareWith(t, out, "dataprep_create.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep create --source source --output output")
	require.NoError(t, err)
	CompareWith(t, out, "dataprep_create_verbose.txt")
}

func TestDataPrepListHandler(t *testing.T) {
	mockHandler := new(MockDataPrep)
	defer swapDataPrepHandler(mockHandler)()

	mockHandler.On("ListHandler", mock.Anything, mock.Anything).Return([]model.Preparation{testPreparation}, nil)
	out, _, err := Run(context.Background(), "singularity prep list")
	require.NoError(t, err)
	CompareWith(t, out, "dataprep_create.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep list")
	require.NoError(t, err)
	CompareWith(t, out, "dataprep_create_verbose.txt")
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
	CompareWith(t, out, "dataprep_daggen_job.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep start-daggen 1 name")
	require.NoError(t, err)
	CompareWith(t, out, "dataprep_daggen_job_verbose.txt")
}

func TestDataPrepPauseDagGenHandler(t *testing.T) {
	mockHandler := new(MockDataPrep)
	defer swapDataPrepHandler(mockHandler)()

	mockHandler.On("PauseDagGenHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&testDagGenJob, nil)
	out, _, err := Run(context.Background(), "singularity prep pause-daggen 1 name")
	require.NoError(t, err)
	CompareWith(t, out, "dataprep_daggen_job.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep pause-daggen 1 name")
	require.NoError(t, err)
	CompareWith(t, out, "dataprep_daggen_job_verbose.txt")
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
	CompareWith(t, out, "dataprep_scan_job.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep start-scan 1 name")
	require.NoError(t, err)
	CompareWith(t, out, "dataprep_scan_job_verbose.txt")
}

func TestDataPrepPauseScanHandler(t *testing.T) {
	mockHandler := new(MockDataPrep)
	defer swapDataPrepHandler(mockHandler)()

	mockHandler.On("PauseScanHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&testScanob, nil)
	out, _, err := Run(context.Background(), "singularity prep pause-scan 1 name")
	require.NoError(t, err)
	CompareWith(t, out, "dataprep_scan_job.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep pause-scan 1 name")
	require.NoError(t, err)
	CompareWith(t, out, "dataprep_scan_job_verbose.txt")
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
	CompareWith(t, out, "dataprep_pack_job.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep start-pack 1 name")
	require.NoError(t, err)
	CompareWith(t, out, "dataprep_pack_job_verbose.txt")
}

func TestDataPrepPausePackHandler(t *testing.T) {
	mockHandler := new(MockDataPrep)
	defer swapDataPrepHandler(mockHandler)()

	mockHandler.On("PausePackHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]model.Job{testPackJob}, nil)
	out, _, err := Run(context.Background(), "singularity prep pause-pack 1 name")
	require.NoError(t, err)
	CompareWith(t, out, "dataprep_pack_job.txt")
	out, _, err = Run(context.Background(), "singularity --verbose prep pause-pack 1 name")
	require.NoError(t, err)
	CompareWith(t, out, "dataprep_pack_job_verbose.txt")
}

func TestDataPrepAttachSourceHandler(t *testing.T) {

}
