package job

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/gotidy/ptr"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestPrepareToPackSourceHandler_NotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := Default.PrepareToPackSourceHandler(ctx, db, "id", "name")
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
}

func TestPrepareToPackSourceHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tmpdir := t.TempDir()
		preparation := model.Preparation{
			MaxSize:   1 << 20, // 1MiB
			PieceSize: 1 << 21, // 2MiB
			Name:      "prep",
			SourceStorages: []model.Storage{
				{
					Name: "source",
					Type: "local",
					Path: tmpdir,
				},
			},
		}
		err := db.Create(&preparation).Error
		require.NoError(t, err)
		files := []model.File{{
			Path:         "1.bin",
			Hash:         "",
			Size:         500_000,
			AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
			Directory: &model.Directory{
				AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
			},
			FileRanges: []model.FileRange{{
				Offset: 0,
				Length: 500_000,
			}},
		}, {
			Path:         "2.bin",
			Hash:         "",
			Size:         500_000,
			AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
			Directory: &model.Directory{
				AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
			},
			FileRanges: []model.FileRange{{
				Offset: 0,
				Length: 500_000,
			}},
		}, {
			Path:         "3.bin",
			Hash:         "",
			Size:         500_000,
			AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
			Directory: &model.Directory{
				AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
			},
			FileRanges: []model.FileRange{{
				Offset: 0,
				Length: 500_000,
			}},
		}}
		err = db.Create(&files).Error
		require.NoError(t, err)
		err = Default.PrepareToPackSourceHandler(ctx, db, "prep", "source")
		require.NoError(t, err)

		var jobs []model.Job
		err = db.Find(&jobs).Error
		require.NoError(t, err)
		require.Len(t, jobs, 2)
	})
}

func TestPackHandler_JobNotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, err := Default.PackHandler(ctx, db, 1)
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
}

func TestPackHandler_Success(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tmpdir := t.TempDir()
		err := os.WriteFile(filepath.Join(tmpdir, "test.txt"), []byte("test"), 0644)
		require.NoError(t, err)
		stat, err := os.Stat(filepath.Join(tmpdir, "test.txt"))
		require.NoError(t, err)
		job := model.Job{
			Type:  model.Pack,
			State: model.Ready,
			Attachment: &model.SourceAttachment{
				Preparation: &model.Preparation{
					MaxSize:   1 << 34,
					PieceSize: 1 << 35,
					Name:      "prep",
				},
				Storage: &model.Storage{
					Name: "source",
					Type: "local",
					Path: tmpdir,
				},
			},
			FileRanges: []model.FileRange{
				{
					Offset: 0,
					Length: 4,
					File: &model.File{
						Path:             "test.txt",
						Size:             4,
						LastModifiedNano: stat.ModTime().UnixNano(),
						AttachmentID:     ptr.Of(model.SourceAttachmentID(1)),
						Directory: &model.Directory{
							AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
						},
					},
				},
			},
		}
		err = db.Create(&job).Error
		require.NoError(t, err)
		car, err := Default.PackHandler(ctx, db, 1)
		require.NoError(t, err)
		require.NotNil(t, car)
		// With default minPieceSize of 1 MiB, the virtual file size is padded to (1 MiB * 127/128)
		require.EqualValues(t, 1040384, car.FileSize)
		// PieceCID reflects the 1 MiB piece size (padded from natural ~128 byte piece)
		require.EqualValues(t, "baga6ea4seaqpikooah5wmbpjmnvx3ysyf36xagymjtbccnf5twt2cpaqcgcwqha", car.PieceCID.String())
		err = db.Find(&job, 1).Error
		require.NoError(t, err)
		require.Equal(t, model.Complete, job.State)
	})
}

func TestStartPackHandler_StorageNotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{}).Error
		require.NoError(t, err)
		_, err = Default.StartPackHandler(ctx, db, "1", "not found", 1)
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
}

func TestStartPackHandler_PreparationNotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			SourceStorages: []model.Storage{{
				Name: "source",
			}},
		}).Error
		require.NoError(t, err)
		_, err = Default.StartPackHandler(ctx, db, "2", "source", 1)
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
}

func TestStartPackHandler_JobNotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			SourceStorages: []model.Storage{{
				Name: "source",
			}},
		}).Error
		require.NoError(t, err)
		_, err = Default.StartPackHandler(ctx, db, "1", "source", 1)
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
}

func TestStartPackHandler_StartExisting(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			SourceStorages: []model.Storage{{
				Name: "source",
			}},
		}).Error
		require.NoError(t, err)
		err = db.Create(&model.Job{
			AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
			State:        model.Error,
			Type:         model.Pack,
		}).Error
		require.NoError(t, err)
		jobs, err := Default.StartPackHandler(ctx, db, "1", "source", 1)
		require.NoError(t, err)
		require.Len(t, jobs, 1)
		require.Equal(t, model.Ready, jobs[0].State)
		require.Equal(t, model.Pack, jobs[0].Type)
	})
}

func TestStartPackHandler_AlreadyProcessing(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			SourceStorages: []model.Storage{{
				Name: "source",
			}},
		}).Error
		require.NoError(t, err)
		err = db.Create(&model.Job{
			AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
			State:        model.Processing,
			Type:         model.Pack,
		}).Error
		require.NoError(t, err)
		_, err = Default.StartPackHandler(ctx, db, "1", "source", 1)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	})
}

func TestStartPackHandler_All(t *testing.T) {
	for _, name := range []string{"1", "name"} {
		t.Run(name, func(t *testing.T) {
			testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
				err := db.Create(&model.Preparation{
					Name: "name",
					SourceStorages: []model.Storage{{
						Name: "source",
					}},
				}).Error
				require.NoError(t, err)
				err = db.Create(&model.Job{
					AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
					State:        model.Error,
					Type:         model.Pack,
				}).Error
				require.NoError(t, err)
				jobs, err := Default.StartPackHandler(ctx, db, name, "source", 0)
				require.NoError(t, err)
				require.Len(t, jobs, 1)
				require.Equal(t, model.Ready, jobs[0].State)
				require.Equal(t, model.Pack, jobs[0].Type)
			})
		})
	}
}

func TestPausePackHandler_All(t *testing.T) {
	for _, name := range []string{"1", "name"} {
		t.Run(name, func(t *testing.T) {
			testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
				err := db.Create(&model.Preparation{
					Name: "name",
					SourceStorages: []model.Storage{{
						Name: "source",
					}},
				}).Error
				require.NoError(t, err)
				err = db.Create(&model.Job{
					AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
					State:        model.Ready,
					Type:         model.Pack,
				}).Error
				require.NoError(t, err)
				jobs, err := Default.PausePackHandler(ctx, db, name, "source", 0)
				require.NoError(t, err)
				require.Len(t, jobs, 1)
				require.Equal(t, model.Paused, jobs[0].State)
				require.Equal(t, model.Pack, jobs[0].Type)
			})
		})
	}
}

func TestPausePackHandler_Existing(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			SourceStorages: []model.Storage{{
				Name: "source",
			}},
		}).Error
		require.NoError(t, err)
		err = db.Create(&model.Job{
			AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
			State:        model.Ready,
			Type:         model.Pack,
		}).Error
		require.NoError(t, err)
		jobs, err := Default.PausePackHandler(ctx, db, "1", "source", 1)
		require.NoError(t, err)
		require.Len(t, jobs, 1)
		require.Equal(t, model.Paused, jobs[0].State)
		require.Equal(t, model.Pack, jobs[0].Type)
	})
}

func TestPausePackHandler_JobNotExist(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			SourceStorages: []model.Storage{{
				Name: "source",
			}},
		}).Error
		require.NoError(t, err)
		err = db.Create(&model.Job{
			AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
			State:        model.Ready,
			Type:         model.Pack,
		}).Error
		require.NoError(t, err)
		_, err = Default.PausePackHandler(ctx, db, "1", "source", 2)
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
}

func TestPausePackHandler_AlreadyPaused(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			SourceStorages: []model.Storage{{
				Name: "source",
			}},
		}).Error
		require.NoError(t, err)
		err = db.Create(&model.Job{
			AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
			State:        model.Paused,
			Type:         model.Pack,
		}).Error
		require.NoError(t, err)
		_, err = Default.PausePackHandler(ctx, db, "1", "source", 1)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	})
}
