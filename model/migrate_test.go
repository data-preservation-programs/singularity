package model_test

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestFKSetNullOnDelete(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Create preparation with storage
		storage := model.Storage{Name: "test", Type: "local", Path: "/tmp"}
		require.NoError(t, db.Create(&storage).Error)

		prep := model.Preparation{Name: "test", MaxSize: 1024, PieceSize: 1024}
		require.NoError(t, db.Create(&prep).Error)

		attachment := model.SourceAttachment{PreparationID: prep.ID, StorageID: storage.ID}
		require.NoError(t, db.Create(&attachment).Error)

		// Create child records
		job := model.Job{Type: model.Pack, State: model.Complete, AttachmentID: &attachment.ID}
		require.NoError(t, db.Create(&job).Error)

		dir := model.Directory{Name: "root", AttachmentID: &attachment.ID}
		require.NoError(t, db.Create(&dir).Error)

		file := model.File{Path: "test.txt", Size: 100, AttachmentID: &attachment.ID, DirectoryID: &dir.ID}
		require.NoError(t, db.Create(&file).Error)

		car := model.Car{PieceSize: 1024, PreparationID: &prep.ID, AttachmentID: &attachment.ID, JobID: &job.ID}
		require.NoError(t, db.Create(&car).Error)

		carBlock := model.CarBlock{CarOffset: 0, CarID: &car.ID, FileID: &file.ID}
		require.NoError(t, db.Create(&carBlock).Error)

		// Delete preparation (cascades to attachment)
		require.NoError(t, db.Delete(&prep).Error)

		// Verify child records exist with NULL FKs
		var loadedJob model.Job
		require.NoError(t, db.First(&loadedJob, job.ID).Error)
		require.Nil(t, loadedJob.AttachmentID)

		var loadedDir model.Directory
		require.NoError(t, db.First(&loadedDir, dir.ID).Error)
		require.Nil(t, loadedDir.AttachmentID)

		var loadedFile model.File
		require.NoError(t, db.First(&loadedFile, file.ID).Error)
		require.Nil(t, loadedFile.AttachmentID)

		var loadedCar model.Car
		require.NoError(t, db.First(&loadedCar, car.ID).Error)
		require.Nil(t, loadedCar.PreparationID)
		require.Nil(t, loadedCar.AttachmentID)

		var loadedCarBlock model.CarBlock
		require.NoError(t, db.First(&loadedCarBlock, carBlock.ID).Error)
		// CarBlock FKs still set - Car and File exist (orphaned), not deleted yet
		require.NotNil(t, loadedCarBlock.CarID)
		require.NotNil(t, loadedCarBlock.FileID)

		// Delete orphaned car - CarBlock.CarID should become NULL
		require.NoError(t, db.Delete(&loadedCar).Error)
		require.NoError(t, db.First(&loadedCarBlock, carBlock.ID).Error)
		require.Nil(t, loadedCarBlock.CarID)

		// Delete orphaned file - CarBlock.FileID should become NULL
		require.NoError(t, db.Delete(&loadedFile).Error)
		require.NoError(t, db.First(&loadedCarBlock, carBlock.ID).Error)
		require.Nil(t, loadedCarBlock.FileID)
	})
}

func TestInferPieceTypes(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		prep := model.Preparation{Name: "test", MaxSize: 1024, PieceSize: 1024}
		require.NoError(t, db.Create(&prep).Error)

		storage := model.Storage{Name: "test", Type: "local", Path: "/tmp"}
		require.NoError(t, db.Create(&storage).Error)

		attachment := model.SourceAttachment{PreparationID: prep.ID, StorageID: storage.ID}
		require.NoError(t, db.Create(&attachment).Error)

		file := model.File{Path: "test.txt", Size: 100, AttachmentID: &attachment.ID}
		require.NoError(t, db.Create(&file).Error)

		// inline data piece: car_blocks reference files
		inlineCar := model.Car{PieceSize: 1024, PreparationID: &prep.ID, AttachmentID: &attachment.ID}
		require.NoError(t, db.Create(&inlineCar).Error)
		cb := model.CarBlock{CarOffset: 0, CarID: &inlineCar.ID, FileID: &file.ID}
		require.NoError(t, db.Create(&cb).Error)

		// non-inline data piece: no file refs in car_blocks, but num_of_files > 0
		nonInlineCar := model.Car{PieceSize: 1024, NumOfFiles: 5, PreparationID: &prep.ID, AttachmentID: &attachment.ID}
		require.NoError(t, db.Create(&nonInlineCar).Error)

		// dag piece: no file refs, num_of_files == 0
		dagCar := model.Car{PieceSize: 1024, PreparationID: &prep.ID, AttachmentID: &attachment.ID}
		require.NoError(t, db.Create(&dagCar).Error)

		// all should have empty piece_type
		for _, id := range []model.CarID{inlineCar.ID, nonInlineCar.ID, dagCar.ID} {
			var c model.Car
			require.NoError(t, db.First(&c, id).Error)
			require.Empty(t, c.PieceType)
		}

		// run migration
		require.NoError(t, model.AutoMigrate(db))

		var c1, c2, c3 model.Car

		require.NoError(t, db.First(&c1, inlineCar.ID).Error)
		require.Equal(t, model.DataPiece, c1.PieceType, "inline car with file refs should be data")

		require.NoError(t, db.First(&c2, nonInlineCar.ID).Error)
		require.Equal(t, model.DataPiece, c2.PieceType, "non-inline car with num_of_files > 0 should be data")

		require.NoError(t, db.First(&c3, dagCar.ID).Error)
		require.Equal(t, model.DagPiece, c3.PieceType, "car with no file refs and num_of_files == 0 should be dag")

		// idempotent: running again should not change anything
		require.NoError(t, model.AutoMigrate(db))

		var c4 model.Car
		require.NoError(t, db.First(&c4, nonInlineCar.ID).Error)
		require.Equal(t, model.DataPiece, c4.PieceType)
	})
}
