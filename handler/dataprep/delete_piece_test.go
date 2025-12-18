package dataprep

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/gotidy/ptr"
	"github.com/ipfs/boxo/util"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// Helper to create a valid piece CID for testing
func testPieceCID(seed string) cid.Cid {
	return cid.NewCidV1(cid.FilCommitmentUnsealed, util.Hash([]byte(seed)))
}

func TestDeletePieceHandler_PrepNotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		pieceCID := testPieceCID("test1")
		err := Default.DeletePieceHandler(ctx, db, "nonexistent", pieceCID.String(), DeletePieceRequest{})
		require.ErrorIs(t, err, handlererror.ErrNotFound)
		require.ErrorContains(t, err, "preparation")
	})
}

func TestDeletePieceHandler_InvalidPieceCID(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		prep := model.Preparation{Name: "test-prep"}
		require.NoError(t, db.Create(&prep).Error)

		err := Default.DeletePieceHandler(ctx, db, "test-prep", "invalid-cid", DeletePieceRequest{})
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	})
}

func TestDeletePieceHandler_PieceNotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		prep := model.Preparation{Name: "test-prep"}
		require.NoError(t, db.Create(&prep).Error)

		pieceCID := testPieceCID("nonexistent")
		err := Default.DeletePieceHandler(ctx, db, "test-prep", pieceCID.String(), DeletePieceRequest{})
		require.ErrorIs(t, err, handlererror.ErrNotFound)
		require.ErrorContains(t, err, "piece")
	})
}

func TestDeletePieceHandler_DealExistsWithoutForce(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		prep := model.Preparation{Name: "test-prep"}
		require.NoError(t, db.Create(&prep).Error)

		pieceCID := testPieceCID("test-piece")
		car := model.Car{
			PieceCID:      model.CID(pieceCID),
			PreparationID: &prep.ID,
			PieceType:     model.DataPiece,
		}
		require.NoError(t, db.Create(&car).Error)

		// Create a wallet first to satisfy FK constraint
		wallet := model.Wallet{ID: "f01234", Address: "f01234"}
		require.NoError(t, db.Create(&wallet).Error)

		deal := model.Deal{
			PieceCID: model.CID(pieceCID),
			Provider: "f05678",
			ClientID: "f01234",
		}
		require.NoError(t, db.Create(&deal).Error)

		err := Default.DeletePieceHandler(ctx, db, "test-prep", pieceCID.String(), DeletePieceRequest{Force: false})
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "deals")
	})
}

func TestDeletePieceHandler_DealExistsWithForce(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		prep := model.Preparation{Name: "test-prep"}
		require.NoError(t, db.Create(&prep).Error)

		pieceCID := testPieceCID("test-piece")
		car := model.Car{
			PieceCID:      model.CID(pieceCID),
			PreparationID: &prep.ID,
			PieceType:     model.DataPiece,
		}
		require.NoError(t, db.Create(&car).Error)

		// Create a wallet first to satisfy FK constraint
		wallet := model.Wallet{ID: "f01234", Address: "f01234"}
		require.NoError(t, db.Create(&wallet).Error)

		deal := model.Deal{
			PieceCID: model.CID(pieceCID),
			Provider: "f05678",
			ClientID: "f01234",
		}
		require.NoError(t, db.Create(&deal).Error)

		err := Default.DeletePieceHandler(ctx, db, "test-prep", pieceCID.String(), DeletePieceRequest{Force: true})
		require.NoError(t, err)

		// Verify car is deleted
		var count int64
		db.Model(&model.Car{}).Count(&count)
		require.Zero(t, count)
	})
}

func TestDeletePieceHandler_DataPiece_ResetsFileRanges(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		prep := model.Preparation{
			Name:           "test-prep",
			SourceStorages: []model.Storage{{}},
		}
		require.NoError(t, db.Create(&prep).Error)

		job := model.Job{
			AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
			State:        model.Complete,
		}
		require.NoError(t, db.Create(&job).Error)

		pieceCID := testPieceCID("data-piece")
		car := model.Car{
			PieceCID:      model.CID(pieceCID),
			PreparationID: &prep.ID,
			JobID:         &job.ID,
			PieceType:     model.DataPiece,
		}
		require.NoError(t, db.Create(&car).Error)

		file := model.File{
			AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
			Path:         "test.txt",
			Size:         100,
		}
		require.NoError(t, db.Create(&file).Error)

		fileRange := model.FileRange{
			FileID: file.ID,
			JobID:  &job.ID,
			Offset: 0,
			Length: 100,
		}
		require.NoError(t, db.Create(&fileRange).Error)

		err := Default.DeletePieceHandler(ctx, db, "test-prep", pieceCID.String(), DeletePieceRequest{DeleteCar: false})
		require.NoError(t, err)

		// Verify file_range.job_id is reset to NULL
		var fr model.FileRange
		require.NoError(t, db.First(&fr, fileRange.ID).Error)
		require.Nil(t, fr.JobID, "file_range.job_id should be reset to NULL")

		// Verify car is deleted
		var count int64
		db.Model(&model.Car{}).Count(&count)
		require.Zero(t, count)
	})
}

func TestDeletePieceHandler_DagPiece_ResetsDirectories(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		prep := model.Preparation{
			Name:           "test-prep",
			SourceStorages: []model.Storage{{}},
		}
		require.NoError(t, db.Create(&prep).Error)

		pieceCID := testPieceCID("dag-piece")
		car := model.Car{
			PieceCID:      model.CID(pieceCID),
			PreparationID: &prep.ID,
			AttachmentID:  ptr.Of(model.SourceAttachmentID(1)),
			PieceType:     model.DagPiece,
		}
		require.NoError(t, db.Create(&car).Error)

		dir := model.Directory{
			AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
			Name:         "test-dir",
			Exported:     true,
		}
		require.NoError(t, db.Create(&dir).Error)

		err := Default.DeletePieceHandler(ctx, db, "test-prep", pieceCID.String(), DeletePieceRequest{DeleteCar: false})
		require.NoError(t, err)

		// Verify directory.exported is reset to false
		var d model.Directory
		require.NoError(t, db.First(&d, dir.ID).Error)
		require.False(t, d.Exported, "directory.exported should be reset to false")

		// Verify car is deleted
		var count int64
		db.Model(&model.Car{}).Count(&count)
		require.Zero(t, count)
	})
}

func TestDeletePieceHandler_NonInline_DeletesCarFile(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tmp := t.TempDir()
		carFile := filepath.Join(tmp, "test.car")
		require.NoError(t, os.WriteFile(carFile, []byte("car data"), 0o644))

		storage := model.Storage{
			Name: "output",
			Type: "local",
			Path: tmp,
		}
		require.NoError(t, db.Create(&storage).Error)

		prep := model.Preparation{
			Name:           "test-prep",
			OutputStorages: []model.Storage{storage},
		}
		require.NoError(t, db.Create(&prep).Error)

		pieceCID := testPieceCID("noninline-piece")
		car := model.Car{
			PieceCID:      model.CID(pieceCID),
			PreparationID: &prep.ID,
			StorageID:     &storage.ID,
			StoragePath:   "test.car",
			PieceType:     model.DataPiece,
		}
		require.NoError(t, db.Create(&car).Error)

		err := Default.DeletePieceHandler(ctx, db, "test-prep", pieceCID.String(), DeletePieceRequest{DeleteCar: true})
		require.NoError(t, err)

		// Verify file is deleted
		_, err = os.Stat(carFile)
		require.True(t, os.IsNotExist(err), "CAR file should be deleted")

		// Verify car record is deleted
		var count int64
		db.Model(&model.Car{}).Count(&count)
		require.Zero(t, count)
	})
}

func TestDeletePieceHandler_Inline_NoStorageDeletion(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		prep := model.Preparation{Name: "test-prep"}
		require.NoError(t, db.Create(&prep).Error)

		pieceCID := testPieceCID("inline-piece")
		car := model.Car{
			PieceCID:            model.CID(pieceCID),
			PreparationID:       &prep.ID,
			StorageID:           nil, // Inline - no storage
			PieceType:           model.DataPiece,
			MinPieceSizePadding: 100, // Indicates inline
		}
		require.NoError(t, db.Create(&car).Error)

		carBlock := model.CarBlock{
			CarID: &car.ID,
		}
		require.NoError(t, db.Create(&carBlock).Error)

		err := Default.DeletePieceHandler(ctx, db, "test-prep", pieceCID.String(), DeletePieceRequest{DeleteCar: true})
		require.NoError(t, err)

		// Verify car and car_blocks are deleted
		var carCount, blockCount int64
		db.Model(&model.Car{}).Count(&carCount)
		db.Model(&model.CarBlock{}).Count(&blockCount)
		require.Zero(t, carCount)
		require.Zero(t, blockCount)
	})
}

func TestDeletePieceHandler_DeletesCarBlocks(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		prep := model.Preparation{Name: "test-prep"}
		require.NoError(t, db.Create(&prep).Error)

		pieceCID := testPieceCID("blocks-piece")
		car := model.Car{
			PieceCID:      model.CID(pieceCID),
			PreparationID: &prep.ID,
			PieceType:     model.DataPiece,
		}
		require.NoError(t, db.Create(&car).Error)

		// Create multiple car blocks
		for i := 0; i < 5; i++ {
			carBlock := model.CarBlock{
				CarID:     &car.ID,
				CarOffset: int64(i * 1000),
			}
			require.NoError(t, db.Create(&carBlock).Error)
		}

		var blockCount int64
		db.Model(&model.CarBlock{}).Count(&blockCount)
		require.Equal(t, int64(5), blockCount)

		err := Default.DeletePieceHandler(ctx, db, "test-prep", pieceCID.String(), DeletePieceRequest{DeleteCar: false})
		require.NoError(t, err)

		// Verify all car_blocks are deleted
		db.Model(&model.CarBlock{}).Count(&blockCount)
		require.Zero(t, blockCount)
	})
}

func TestDeletePieceHandler_PieceBelongsToDifferentPrep(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		prep1 := model.Preparation{Name: "prep1"}
		prep2 := model.Preparation{Name: "prep2"}
		require.NoError(t, db.Create(&prep1).Error)
		require.NoError(t, db.Create(&prep2).Error)

		pieceCID := testPieceCID("other-prep-piece")
		// Create car belonging to prep1
		car := model.Car{
			PieceCID:      model.CID(pieceCID),
			PreparationID: &prep1.ID,
			PieceType:     model.DataPiece,
		}
		require.NoError(t, db.Create(&car).Error)

		// Try to delete from prep2 - should fail
		err := Default.DeletePieceHandler(ctx, db, "prep2", pieceCID.String(), DeletePieceRequest{})
		require.ErrorIs(t, err, handlererror.ErrNotFound)
		require.ErrorContains(t, err, "piece")
	})
}
