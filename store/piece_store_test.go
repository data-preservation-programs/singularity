package store

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/gotidy/ptr"
	"github.com/ipfs/go-cid"
	util "github.com/ipfs/go-ipfs-util"
	"github.com/multiformats/go-varint"
	fs2 "github.com/rclone/rclone/fs"
	"github.com/stretchr/testify/require"
)

func TestPieceReader_FileChanged(t *testing.T) {
	ctx := context.Background()
	tmp := t.TempDir()
	testFileContent := []byte("12345678901234567890")
	cidValue := cid.NewCidV1(cid.Raw, util.Hash(testFileContent))
	err := os.WriteFile(filepath.Join(tmp, "1.txt"), testFileContent, 0644)
	require.NoError(t, err)

	car := model.Car{
		Header:   []byte("header"),
		FileSize: 63,
	}
	source := model.Source{
		ID:   1,
		Type: "local",
		Path: tmp,
	}
	carBlocks := []model.CarBlock{
		{
			CarOffset:      6,
			CarBlockLength: 57,
			Varint:         []byte{56},
			FileID:         ptr.Of(uint64(1)),
			CID:            model.CID(cidValue),
		},
	}
	files := []model.File{{
		ID:                        1,
		SourceID:                  1,
		Path:                      "1.txt",
		LastModifiedTimestampNano: testutil.GetFileTimestamp(t, filepath.Join(tmp, "1.txt")),
		Size:                      20,
	}}
	reader, err := NewPieceReader(ctx, car, source,
		carBlocks, files,
		datasource.DefaultHandlerResolver{})
	require.NoError(t, err)
	require.NotNil(t, reader)
	defer require.NoError(t, reader.Close())
	// Update the file
	err = os.WriteFile(filepath.Join(tmp, "1.txt"), []byte("changed"), 0644)
	require.NoError(t, err)
	_, err = io.ReadAll(reader)
	require.ErrorIs(t, err, &FileHasChangedError{})
	// Remove the file
	_, err = reader.Seek(0, io.SeekStart)
	require.NoError(t, err)
	err = os.Remove(filepath.Join(tmp, "1.txt"))
	require.NoError(t, err)
	_, err = io.ReadAll(reader)
	require.ErrorIs(t, err, fs2.ErrorObjectNotFound)
}

func TestPieceReader_LargeFile(t *testing.T) {
	tmp := t.TempDir()
	testFileContent := testutil.GenerateRandomBytes(1024 * 1024)
	cidValue := cid.NewCidV1(cid.Raw, util.Hash(testFileContent))
	var err error
	err = os.WriteFile(filepath.Join(tmp, "1.txt"), testFileContent, 0644)
	require.NoError(t, err)
	ctx := context.Background()
	size := int64(6 + 39 + 1024*1024)
	car := model.Car{
		Header:   []byte("header"),
		FileSize: size,
	}
	source := model.Source{
		ID:   1,
		Type: "local",
		Path: tmp,
	}
	carBlocks := []model.CarBlock{
		{
			CarOffset:      6,
			CarBlockLength: int32(size - 6),
			Varint:         varint.ToUvarint(36 + 1024*1024),
			FileID:         ptr.Of(uint64(1)),
			CID:            model.CID(cidValue),
		},
	}
	files := []model.File{{
		ID:                        1,
		SourceID:                  1,
		Path:                      "1.txt",
		LastModifiedTimestampNano: testutil.GetFileTimestamp(t, filepath.Join(tmp, "1.txt")),
		Size:                      1024 * 1024,
	}}
	reader, err := NewPieceReader(ctx, car, source,
		carBlocks, files,
		datasource.DefaultHandlerResolver{})
	require.NoError(t, err)
	require.NotNil(t, reader)
	read, err := io.ReadAll(reader)
	require.NoError(t, err)
	require.EqualValues(t, size, len(read))
}

func TestPieceReader_ReadSeek(t *testing.T) {
	tmp := t.TempDir()
	testFileContent := []byte("12345678901234567890")
	cidValue := cid.NewCidV1(cid.Raw, util.Hash(testFileContent))
	var err error
	err = os.WriteFile(filepath.Join(tmp, "1.txt"), testFileContent, 0644)
	require.NoError(t, err)
	err = os.WriteFile(filepath.Join(tmp, "2.txt"), testFileContent, 0644)
	require.NoError(t, err)
	err = os.WriteFile(filepath.Join(tmp, "3.txt"), testFileContent, 0644)
	require.NoError(t, err)
	ctx := context.Background()
	size := int64(234)

	car := model.Car{
		Header:   []byte("header"),
		FileSize: size,
	}
	source := model.Source{
		ID:   1,
		Type: "local",
		Path: tmp,
	}
	carBlocks := []model.CarBlock{
		{
			CarOffset:      6,
			CarBlockLength: 57,
			Varint:         []byte{56},
			FileID:         ptr.Of(uint64(1)),
			CID:            model.CID(cidValue),
		},
		{
			CarOffset:      63,
			CarBlockLength: 57,
			Varint:         []byte{56},
			FileID:         ptr.Of(uint64(2)),
			CID:            model.CID(cidValue),
		},
		{
			CarOffset:      120,
			CarBlockLength: 57,
			Varint:         []byte{56},
			FileID:         ptr.Of(uint64(3)),
			CID:            model.CID(cidValue),
		},
		{
			CarOffset:      177,
			CarBlockLength: 57,
			Varint:         []byte{56},
			CID:            model.CID(cidValue),
			RawBlock:       testFileContent,
		},
	}
	files := []model.File{{
		ID:                        1,
		SourceID:                  1,
		Path:                      "1.txt",
		LastModifiedTimestampNano: testutil.GetFileTimestamp(t, filepath.Join(tmp, "1.txt")),
		Size:                      20,
	}, {
		ID:                        2,
		SourceID:                  1,
		Path:                      "2.txt",
		LastModifiedTimestampNano: testutil.GetFileTimestamp(t, filepath.Join(tmp, "2.txt")),
		Size:                      20,
	}, {
		ID:                        3,
		SourceID:                  1,
		Path:                      "3.txt",
		LastModifiedTimestampNano: testutil.GetFileTimestamp(t, filepath.Join(tmp, "3.txt")),
		Size:                      20,
	}}
	reader, err := NewPieceReader(ctx, car, source,
		carBlocks, files,
		datasource.DefaultHandlerResolver{})
	require.NoError(t, err)
	require.NotNil(t, reader)
	defer require.NoError(t, reader.Close())
	pos, err := reader.Seek(1, io.SeekStart)
	require.NoError(t, err)
	require.EqualValues(t, 1, pos)
	pos, err = reader.Seek(1, io.SeekCurrent)
	require.NoError(t, err)
	require.EqualValues(t, 2, pos)
	pos, err = reader.Seek(-1, io.SeekEnd)
	require.NoError(t, err)
	require.EqualValues(t, 233, pos)
	_, err = reader.Seek(-1, io.SeekStart)
	require.ErrorIs(t, err, ErrNegativeOffset)
	_, err = reader.Seek(0, 100)
	require.ErrorIs(t, err, ErrInvalidWhence)
	_, err = reader.Seek(10000, io.SeekStart)
	require.ErrorIs(t, err, ErrOffsetOutOfRange)

	posMap := map[int64]int{
		1:    -1,
		5:    -1,
		6:    0,
		62:   0,
		63:   1,
		119:  1,
		120:  2,
		176:  2,
		177:  3,
		233:  3,
		size: 3,
	}

	for pos, blockIndex := range posMap {
		_, err = reader.Seek(pos, io.SeekStart)
		require.NoError(t, err)
		require.EqualValues(t, blockIndex, reader.blockIndex)
		cloned := reader.Clone(context.Background())
		defer require.NoError(t, cloned.Close())
		require.EqualValues(t, -1, cloned.blockIndex)
		require.EqualValues(t, 0, cloned.pos)
	}

	for _, pos := range []int64{0, 20, 50, 100, 150, 170, 230, size} {
		_, err = reader.Seek(pos, io.SeekStart)
		require.NoError(t, err)
		read, err := io.ReadAll(reader)
		require.NoError(t, err)
		require.EqualValues(t, size-pos, len(read))
	}

	for _, length := range []int64{0, 20, 50, 100, 150, 170, 230, size} {
		_, err = reader.Seek(0, io.SeekStart)
		require.NoError(t, err)
		n, err := io.ReadFull(reader, make([]byte, length))
		require.NoError(t, err)
		require.EqualValues(t, length, n)
	}

	// context cancelled
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	reader.ctx = ctx
	_, err = io.ReadAll(reader)
	require.ErrorIs(t, err, context.Canceled)
}

func TestNewPieceReader_InvalidConstruction(t *testing.T) {
	tmp := t.TempDir()
	testFilename := "test.txt"
	ctx := context.Background()
	car := model.Car{
		Header:   []byte("header"),
		FileSize: 63,
	}
	source := model.Source{
		ID:   1,
		Type: "local",
		Path: tmp,
	}

	tests := []struct {
		carBlocks []model.CarBlock
		files     []model.File
		err       error
	}{
		{
			carBlocks: []model.CarBlock{},
			files: []model.File{{
				ID:       1,
				SourceID: 999,
				Path:     testFilename,
				Size:     20,
			}},
			err: ErrSourceMismatch,
		},
		{
			carBlocks: []model.CarBlock{},
			files: []model.File{{
				ID:       1,
				SourceID: 1,
				Path:     testFilename,
				Size:     20,
			}},
			err: ErrNoCarBlocks,
		},
		{
			carBlocks: []model.CarBlock{
				{
					CarOffset: 11,
				},
			},
			files: []model.File{{
				ID:       1,
				SourceID: 1,
				Path:     testFilename,
				Size:     20,
			}},
			err: ErrInvalidStartOffset,
		},
		{
			carBlocks: []model.CarBlock{
				{
					CarOffset:      6,
					CarBlockLength: 1000,
				},
			},
			files: []model.File{{
				ID:       1,
				SourceID: 1,
				Path:     testFilename,
				Size:     20,
			}},
			err: ErrInvalidEndOffset,
		},
		{
			carBlocks: []model.CarBlock{
				{
					CarOffset:      6,
					CarBlockLength: 30,
				},
				{
					CarOffset:      60,
					CarBlockLength: 3,
				},
			},
			files: []model.File{{
				ID:       1,
				SourceID: 1,
				Path:     testFilename,
				Size:     20,
			}},
			err: ErrIncontiguousBlocks,
		},
		{
			carBlocks: []model.CarBlock{
				{
					CarOffset:      6,
					CarBlockLength: 57,
				},
			},
			files: []model.File{{
				ID:       1,
				SourceID: 1,
				Path:     testFilename,
				Size:     20,
			}},
			err: varint.ErrUnderflow,
		},
		{
			carBlocks: []model.CarBlock{
				{
					CarOffset:      6,
					CarBlockLength: 57,
					Varint:         []byte{100},
				},
			},
			files: []model.File{{
				ID:       1,
				SourceID: 1,
				Path:     testFilename,
				Size:     20,
			}},
			err: ErrVarintDoesNotMatchBlockLength,
		},
		{
			carBlocks: []model.CarBlock{
				{
					CarOffset:      6,
					CarBlockLength: 57,
					Varint:         []byte{56},
					FileID:         ptr.Of(uint64(100)),
				},
			},
			files: []model.File{{
				ID:       1,
				SourceID: 1,
				Path:     testFilename,
				Size:     20,
			}},
			err: ErrFileNotProvided,
		},
	}

	for _, test := range tests {
		t.Run(test.err.Error(), func(t *testing.T) {
			_, err := NewPieceReader(ctx, car, source,
				test.carBlocks, test.files,
				datasource.DefaultHandlerResolver{})
			require.ErrorIs(t, err, test.err)
		})
	}
}
