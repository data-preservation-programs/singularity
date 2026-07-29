package store

import (
	"context"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/ipfs/boxo/util"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/multiformats/go-varint"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// makeSpanFixture writes a file to a local storage and creates car_blocks
// leaf rows slicing it into blockSize chunks, mirroring what pack produces.
func makeSpanFixture(t *testing.T, db *gorm.DB, content []byte, blockSize int) []cid.Cid {
	t.Helper()
	tmp := t.TempDir()
	name := "data.bin"
	require.NoError(t, os.WriteFile(filepath.Join(tmp, name), content, 0644))

	prep := &model.Preparation{Name: t.Name() + filepath.Base(tmp)}
	require.NoError(t, db.Create(prep).Error)
	storage := &model.Storage{Name: prep.Name, Type: "local", Path: tmp}
	require.NoError(t, db.Create(storage).Error)
	attachment := &model.SourceAttachment{PreparationID: prep.ID, StorageID: storage.ID}
	require.NoError(t, db.Create(attachment).Error)
	dir := &model.Directory{AttachmentID: &attachment.ID}
	require.NoError(t, db.Create(dir).Error)
	file := &model.File{
		Path:             name,
		Size:             int64(len(content)),
		LastModifiedNano: testutil.GetFileTimestamp(t, filepath.Join(tmp, name)),
		AttachmentID:     &attachment.ID,
		DirectoryID:      &dir.ID,
	}
	require.NoError(t, db.Create(file).Error)
	car := &model.Car{
		PieceCID:      model.CID(cid.NewCidV1(cid.FilCommitmentUnsealed, util.Hash([]byte(t.Name())))),
		PieceSize:     1 << 20,
		RootCID:       model.CID(testutil.TestCid),
		PreparationID: &prep.ID,
		AttachmentID:  &attachment.ID,
		PieceType:     model.DataPiece,
	}
	require.NoError(t, db.Create(car).Error)

	var cids []cid.Cid
	carOffset := int64(59)
	for off := 0; off < len(content); off += blockSize {
		chunk := content[off:min(off+blockSize, len(content))]
		c := cid.NewCidV1(cid.Raw, util.Hash(chunk))
		v := varint.ToUvarint(uint64(c.ByteLen() + len(chunk)))
		require.NoError(t, db.Create(&model.CarBlock{
			CarID:          &car.ID,
			CID:            model.CID(c),
			CarOffset:      carOffset,
			CarBlockLength: int32(len(v) + c.ByteLen() + len(chunk)),
			Varint:         v,
			FileID:         &file.ID,
			FileOffset:     int64(off),
		}).Error)
		carOffset += int64(len(v) + c.ByteLen() + len(chunk))
		cids = append(cids, c)
	}
	return cids
}

func TestStorageBlockStore_SequentialRead(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		content := testutil.GenerateRandomBytes(100)
		blockSize := 8
		cids := makeSpanFixture(t, db, content, blockSize)

		s := NewStorageBlockStore(db, SpanConfig{SpanBlocks: 3, PrefetchSpans: 2, CacheBlocks: 64})
		defer s.Close()

		for i, c := range cids {
			blk, err := s.Get(ctx, c)
			require.NoError(t, err)
			expected := content[i*blockSize : min((i+1)*blockSize, len(content))]
			require.Equal(t, expected, blk.RawData())

			size, err := s.GetSize(ctx, c)
			require.NoError(t, err)
			require.Equal(t, len(expected), size)
		}
	})
}

func TestStorageBlockStore_RandomAccess(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		content := testutil.GenerateRandomBytes(96)
		blockSize := 8
		cids := makeSpanFixture(t, db, content, blockSize)

		s := NewStorageBlockStore(db, SpanConfig{SpanBlocks: 2, PrefetchSpans: 1, CacheBlocks: 64})
		defer s.Close()

		for i := len(cids) - 1; i >= 0; i-- {
			blk, err := s.Get(ctx, cids[i])
			require.NoError(t, err)
			require.Equal(t, content[i*blockSize:(i+1)*blockSize], blk.RawData())
		}
	})
}

func TestStorageBlockStore_CacheServesWithoutDB(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		content := testutil.GenerateRandomBytes(32)
		cids := makeSpanFixture(t, db, content, 8)

		s := NewStorageBlockStore(db, SpanConfig{SpanBlocks: 4, CacheBlocks: 64})
		defer s.Close()

		// first Get fetches the whole span into cache
		_, err := s.Get(ctx, cids[0])
		require.NoError(t, err)

		require.NoError(t, db.Where("1 = 1").Delete(&model.CarBlock{}).Error)

		// all blocks of the span are still served from cache
		for i, c := range cids {
			blk, err := s.Get(ctx, c)
			require.NoError(t, err)
			require.Equal(t, content[i*8:(i+1)*8], blk.RawData())

			has, err := s.Has(ctx, c)
			require.NoError(t, err)
			require.True(t, has)
		}
	})
}

func TestStorageBlockStore_ConcurrentReads(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		contentA := testutil.GenerateRandomBytes(128)
		contentB := testutil.GenerateRandomBytes(128)
		cidsA := makeSpanFixture(t, db, contentA, 8)
		cidsB := makeSpanFixture(t, db, contentB, 8)

		s := NewStorageBlockStore(db, SpanConfig{SpanBlocks: 2, PrefetchSpans: 2, MaxBackendReads: 4, CacheBlocks: 8})
		defer s.Close()

		var wg sync.WaitGroup
		for g := 0; g < 8; g++ {
			wg.Add(1)
			go func(seed int64) {
				defer wg.Done()
				r := rand.New(rand.NewSource(seed))
				for range 50 {
					content, cids := contentA, cidsA
					if r.Intn(2) == 1 {
						content, cids = contentB, cidsB
					}
					i := r.Intn(len(cids))
					blk, err := s.Get(ctx, cids[i])
					require.NoError(t, err)
					require.Equal(t, content[i*8:(i+1)*8], blk.RawData())
				}
			}(int64(g))
		}
		wg.Wait()
	})
}

func TestStorageBlockStore_FileChanged(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		content := testutil.GenerateRandomBytes(32)
		cids := makeSpanFixture(t, db, content, 8)

		var storage model.Storage
		require.NoError(t, db.First(&storage).Error)
		require.NoError(t, os.WriteFile(filepath.Join(storage.Path, "data.bin"), []byte("changed"), 0644))

		s := NewStorageBlockStore(db, SpanConfig{})
		defer s.Close()

		_, err := s.Get(ctx, cids[0])
		require.ErrorIs(t, err, ErrFileHasChanged)
	})
}

func TestStorageBlockStore_NotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		s := NewStorageBlockStore(db, SpanConfig{})
		defer s.Close()

		_, err := s.Get(ctx, testutil.TestCid)
		require.ErrorAs(t, err, &format.ErrNotFound{})
		_, err = s.GetSize(ctx, testutil.TestCid)
		require.ErrorAs(t, err, &format.ErrNotFound{})
		has, err := s.Has(ctx, testutil.TestCid)
		require.NoError(t, err)
		require.False(t, has)
	})
}

func TestStorageBlockStore_CancelledWaiterDetaches(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		content := testutil.GenerateRandomBytes(32)
		cids := makeSpanFixture(t, db, content, 8)

		s := NewStorageBlockStore(db, SpanConfig{SpanBlocks: 4, MaxBackendReads: 1})
		defer s.Close()

		// occupy the only connection slot so the fetch parks on the semaphore
		s.sem <- struct{}{}

		waitCtx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
		defer cancel()
		_, err := s.Get(waitCtx, cids[0])
		require.ErrorIs(t, err, context.DeadlineExceeded)

		// the fetch survived the caller: free the slot and it completes into
		// the cache
		<-s.sem
		require.Eventually(t, func() bool {
			return s.cache.Contains(cids[0].KeyString())
		}, 5*time.Second, 10*time.Millisecond)

		blk, err := s.Get(ctx, cids[0])
		require.NoError(t, err)
		require.Equal(t, content[:8], blk.RawData())
	})
}
