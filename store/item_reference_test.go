package store

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/ipfs/boxo/util"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/rclone/rclone/fs"
	"github.com/stretchr/testify/require"
)

func TestFileReferenceBlockStore_Has(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	store := FileReferenceBlockStore{
		DBNoContext: db,
	}

	ctx := context.Background()
	has, err := store.Has(ctx, testutil.TestCid)
	require.NoError(t, err)
	require.False(t, has)

	err = db.Create(&model.CarBlock{
		Car: &model.Car{
			Preparation: &model.Preparation{},
		},
		CID: model.CID(testutil.TestCid),
	}).Error
	require.NoError(t, err)

	has, err = store.Has(ctx, testutil.TestCid)
	require.NoError(t, err)
	require.True(t, has)
}

func TestFileReferenceBlockStore_NotImplemented(t *testing.T) {
	store := &FileReferenceBlockStore{}
	require.ErrorIs(t, store.Put(context.Background(), nil), ErrNotImplemented)
	require.ErrorIs(t, store.PutMany(context.Background(), nil), ErrNotImplemented)
	c, err := store.AllKeysChan(context.Background())
	require.ErrorIs(t, err, ErrNotImplemented)
	require.Nil(t, c)
	require.ErrorIs(t, store.DeleteBlock(context.Background(), cid.Undef), ErrNotImplemented)
	store.HashOnRead(true)
}

func TestFileReferenceBlockStore_GetSize(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	store := FileReferenceBlockStore{
		DBNoContext: db,
	}
	ctx := context.Background()
	cidValue := testutil.TestCid
	_, err = store.GetSize(ctx, cidValue)
	require.ErrorIs(t, err, format.ErrNotFound{})
	err = db.Create(&model.CarBlock{
		Car: &model.Car{
			Preparation: &model.Preparation{},
		},
		CID:      model.CID(cidValue),
		RawBlock: []byte("test"),
	}).Error
	require.NoError(t, err)
	size, err := store.GetSize(ctx, cidValue)
	require.NoError(t, err)
	require.EqualValues(t, 4, size)
}

func TestFileReferenceBlockStore_Get_RawBlock(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	store := FileReferenceBlockStore{
		DBNoContext: db,
	}

	ctx := context.Background()
	cidValue := testutil.TestCid
	_, err = store.Get(ctx, cidValue)
	require.ErrorIs(t, err, format.ErrNotFound{})

	err = db.Create(&model.CarBlock{
		Car: &model.Car{
			Preparation: &model.Preparation{},
		},
		CID:      model.CID(cidValue),
		RawBlock: []byte("test"),
	}).Error
	require.NoError(t, err)
	blk, err := store.Get(ctx, cidValue)
	require.NoError(t, err)
	require.Equal(t, []byte("test"), blk.RawData())
}

func TestFileReferenceBlockStore_Get_FileBlock(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	store := FileReferenceBlockStore{
		DBNoContext: db,
	}

	ctx := context.Background()
	tmp := t.TempDir()
	err = os.WriteFile(filepath.Join(tmp, "1.txt"), []byte("test"), 0644)
	require.NoError(t, err)
	cidValue := cid.NewCidV1(cid.Raw, util.Hash([]byte("test")))
	_, err = store.Get(ctx, cidValue)
	require.ErrorIs(t, err, format.ErrNotFound{})

	err = db.Create(&model.CarBlock{
		Car: &model.Car{
			Preparation: &model.Preparation{},
		},
		CID: model.CID(cidValue),
		File: &model.File{
			SourceStorage: &model.Storage{
				Path: tmp,
				Type: "local",
			},
			Path:             "1.txt",
			Size:             4,
			LastModifiedNano: testutil.GetFileTimestamp(t, filepath.Join(tmp, "1.txt")),
		},
		FileOffset:     0,
		CarBlockLength: 36 + 1 + 4,
	}).Error
	require.NoError(t, err)
	blk, err := store.Get(ctx, cidValue)
	require.NoError(t, err)
	require.Equal(t, []byte("test"), blk.RawData())

	// Change the file
	err = os.WriteFile(filepath.Join(tmp, "1.txt"), []byte("test2"), 0644)
	require.NoError(t, err)
	_, err = store.Get(ctx, cidValue)
	require.ErrorIs(t, err, &FileHasChangedError{})

	// File removed
	err = os.Remove(filepath.Join(tmp, "1.txt"))
	require.NoError(t, err)
	_, err = store.Get(ctx, cidValue)
	require.ErrorIs(t, err, fs.ErrorObjectNotFound)
}
