package datasource_test

import (
	"context"
	"crypto/rand"
	"io"
	"os"
	"path"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	resolvers "github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
)

type packRange struct {
	file string
	min  uint64
	max  uint64
}

type file struct {
	file string
	size uint64
}

type expectedPack struct {
	state model.WorkState
	files []packRange
}

func TestPrepareToPackFile(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		testName      string
		files         []file
		datasetSize   string
		expectedPacks []expectedPack
	}{
		{
			testName: "single pack job",
			files: []file{
				{"a", 1000},
				{"b", 1000},
			},
			datasetSize: "31GiB",
			expectedPacks: []expectedPack{
				{
					state: model.Created,
					files: []packRange{
						{"a", 0, 1000},
						{"b", 0, 1000},
					},
				},
			},
		},
		{
			testName: "two pack jobs",
			files: []file{
				{"a", 1000},
				{"b", 1 << 21},
			},
			datasetSize: "1.1MiB",
			expectedPacks: []expectedPack{
				{
					state: model.Ready,
					files: []packRange{
						{"a", 0, 1000},
						{"b", 0, 1 << 20},
					},
				},
				{
					state: model.Created,
					files: []packRange{
						{"b", 1 << 20, 1 << 21},
					},
				},
			},
		},
		{
			testName: "three pack jobs",
			files: []file{
				{"a", 1000},
				{"b", 1 << 21},
				{"c", 1000},
				{"d", 1 << 19},
			},
			datasetSize: "1.1MiB",
			expectedPacks: []expectedPack{
				{
					state: model.Ready,
					files: []packRange{
						{"a", 0, 1000},
						{"b", 0, 1 << 20},
					},
				},
				{
					state: model.Ready,
					files: []packRange{
						{"b", 1 << 20, 1 << 21},
						{"c", 0, 1000},
					},
				},
				{
					state: model.Created,
					files: []packRange{
						{"d", 0, 1 << 19},
					},
				},
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			db, closer, err := database.OpenInMemory()
			require.NoError(t, err)
			defer closer.Close()

			err = model.AutoMigrate(db)
			require.NoError(t, err)

			datasourceHandlerResolver := &resolvers.DefaultHandlerResolver{}

			datasetName := "test"

			_, err = dataset.CreateHandler(ctx, db.WithContext(ctx), dataset.CreateRequest{
				Name: datasetName,

				MaxSizeStr: testCase.datasetSize,
			})
			require.NoError(t, err)

			testSourcePath := path.Join(t.TempDir(), "singularity-test-source")

			require.NoError(t, os.RemoveAll(testSourcePath))
			require.NoError(t, os.Mkdir(testSourcePath, 0744))

			// Create source
			source, err := datasource.CreateDatasourceHandler(ctx, db.WithContext(ctx), "local", datasetName, map[string]any{
				"sourcePath":        testSourcePath,
				"rescanInterval":    "10s",
				"scanningState":     string(model.Created),
				"deleteAfterExport": false,
			})
			require.NoError(t, err)

			// Create test files to add
			files := make(map[string]*model.File, len(testCase.files))
			for _, testFile := range testCase.files {
				data := io.LimitReader(rand.Reader, int64(testFile.size))
				f, err := os.Create(path.Join(testSourcePath, testFile.file))
				require.NoError(t, err)
				buffer := make([]byte, 1000)
				for {
					read, err := data.Read(buffer)
					if read > 0 {
						writeBuf := buffer[:read]
						f.Write(writeBuf)
					}
					if err != nil {
						require.EqualError(t, err, io.EOF.Error())
						break
					}
				}
				err = f.Close()
				require.NoError(t, err)
				// Push files
				file, err := datasource.PushFileHandler(ctx, db.WithContext(ctx), datasourceHandlerResolver, source.ID, datasource.FileInfo{Path: testFile.file})
				require.NoError(t, err)

				files[testFile.file] = file

				// PrepareToPackFile
				incomplete, err := datasource.PrepareToPackFileHandler(ctx, db.WithContext(ctx), file.ID)
				require.NoError(t, err)
				require.Greater(t, incomplete, int64(0))
			}

			// Check that pack job exists
			packJobs, err := inspect.GetSourcePackJobsHandler(ctx, db.WithContext(ctx), source.ID, inspect.GetSourcePackJobsRequest{})
			require.NoError(t, err)

			require.Len(t, packJobs, len(testCase.expectedPacks))
			for i, expectedPack := range testCase.expectedPacks {
				require.Equal(t, expectedPack.state, packJobs[i].PackingState)

				var expectedFileRanges []model.FileRange
				for _, packRange := range expectedPack.files {
					file := files[packRange.file]
					for _, fileRange := range file.FileRanges {
						if fileRange.Offset >= int64(packRange.min) &&
							fileRange.Offset+fileRange.Length <= int64(packRange.max) {
							expectedFileRanges = append(expectedFileRanges, fileRange)
						}
					}
				}
				expectedFileRangeIDs := make([]uint64, 0, len(expectedFileRanges))
				for _, fileRange := range expectedFileRanges {
					expectedFileRangeIDs = append(expectedFileRangeIDs, fileRange.ID)
				}
				var fileRanges []model.FileRange
				err = db.Where("pack_job_id = ?", packJobs[i].ID).Find(&fileRanges).Error
				require.NoError(t, err)
				fileRangeIDs := make([]uint64, 0, len(fileRanges))
				for _, fileRange := range fileRanges {
					fileRangeIDs = append(fileRangeIDs, fileRange.ID)
				}
				require.ElementsMatch(t, expectedFileRangeIDs, fileRangeIDs)
			}
		})
	}
}
