package pack

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/gotidy/ptr"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestAssembleCar(t *testing.T) {
	tmp := t.TempDir()
	out := t.TempDir()
	err := os.WriteFile(filepath.Join(tmp, "test.txt"), testutil.GenerateRandomBytes(5), 0644)
	require.NoError(t, err)
	stat, err := os.Stat(filepath.Join(tmp, "test.txt"))
	require.NoError(t, err)
	err = os.WriteFile(filepath.Join(tmp, "large.txt"), testutil.GenerateRandomBytes(5_000_000), 0644)
	require.NoError(t, err)
	_, err = os.Stat(filepath.Join(tmp, "large.txt"))
	require.NoError(t, err)
	jobs := []struct {
		name     string
		job      model.Job
		fileSize int64
		one      bool
	}{
		{
			name:     "single file",
			fileSize: 101,
			job: model.Job{
				Type:  model.Pack,
				State: model.Processing,
				Attachment: &model.SourceAttachment{
					Preparation: &model.Preparation{
						MaxSize:   2000000,
						PieceSize: 1 << 21,
					},
					Storage: &model.Storage{
						Type: "local",
						Path: tmp,
					},
				},
				FileRanges: []model.FileRange{
					{
						Offset: 0,
						Length: 5,
						File: &model.File{
							Path:             "test.txt",
							Size:             stat.Size(),
							LastModifiedNano: stat.ModTime().UnixNano(),
							AttachmentID:     1,
							Directory: &model.Directory{
								AttachmentID: 1,
							},
						},
					},
				},
			},
		},
		{
			name:     "splitted file",
			fileSize: 138,
			job: model.Job{
				Type:  model.Pack,
				State: model.Processing,
				Attachment: &model.SourceAttachment{
					Preparation: &model.Preparation{
						MaxSize:   2000000,
						PieceSize: 1 << 21,
					},
					Storage: &model.Storage{
						Type: "local",
						Path: tmp,
					},
				},
				FileRanges: []model.FileRange{
					{
						Offset: 0,
						Length: 2,
						File: &model.File{
							Path:             "test.txt",
							Size:             stat.Size(),
							LastModifiedNano: stat.ModTime().UnixNano(),
							AttachmentID:     1,
							Directory: &model.Directory{
								AttachmentID: 1,
							},
						},
					},
					{
						Offset: 2,
						Length: 3,
						FileID: 1,
						File: &model.File{
							ID:               1,
							Path:             "test.txt",
							Size:             stat.Size(),
							LastModifiedNano: stat.ModTime().UnixNano(),
							AttachmentID:     1,
							DirectoryID:      ptr.Of(uint64(1)),
						},
					},
				},
			},
		},
		{
			name:     "single file non-inline with deletion",
			fileSize: 101,
			one:      true,
			job: model.Job{
				Type:  model.Pack,
				State: model.Processing,
				Attachment: &model.SourceAttachment{
					Preparation: &model.Preparation{
						MaxSize:   2000000,
						PieceSize: 1 << 21,
						OutputStorages: []model.Storage{
							{
								Name: "out",
								Type: "local",
								Path: out,
							},
						},
						DeleteAfterExport: true,
					},
					Storage: &model.Storage{
						Name: "tmp",
						Type: "local",
						Path: tmp,
					},
				},
				FileRanges: []model.FileRange{
					{
						Offset: 0,
						Length: 5,
						File: &model.File{
							Path:             "test.txt",
							Size:             stat.Size(),
							LastModifiedNano: stat.ModTime().UnixNano(),
							AttachmentID:     1,
							Directory: &model.Directory{
								AttachmentID: 1,
							},
						},
					},
				},
			},
		},
	}

	for _, job := range jobs {
		t.Run(job.name, func(t *testing.T) {
			testFunc := func(ctx context.Context, t *testing.T, db *gorm.DB) {
				err := db.Create(&job.job).Error
				require.NoError(t, err)
				car, err := Pack(ctx, db, job.job)
				require.NoError(t, err)
				require.Equal(t, job.fileSize, car.FileSize)
			}
			if job.one {
				testutil.One(t, testFunc)
			} else {
				testutil.All(t, testFunc)
			}
		})
	}
}
