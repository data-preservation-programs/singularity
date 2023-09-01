package datasetworker

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack/packutil"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/google/uuid"
	blocks "github.com/ipfs/go-block-format"
	format "github.com/ipfs/go-ipld-format"
	"github.com/ipfs/go-log/v2"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestScan(t *testing.T) {
	tmp := t.TempDir()
	files := map[string]int{
		"empty.bin":    0,
		"1.bin":        1,
		"1k.bin":       1 << 10,
		"1m.bin":       1 << 20,
		"16m.bin":      16 << 20,
		"1/2/3/10.bin": 10,
		"1/2/3/11.bin": 11,
		"1/2/3 1.bin":  31,
		"1/2/32.bin":   32,
	}
	for path, size := range files {
		err := os.MkdirAll(filepath.Join(tmp, filepath.Dir(path)), 0755)
		require.NoError(t, err)
		err = os.WriteFile(filepath.Join(tmp, path), testutil.GenerateRandomBytes(size), 0644)
		require.NoError(t, err)
	}

	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		thread := &Thread{
			id:          uuid.New(),
			dbNoContext: db,
			logger:      log.Logger("test").With("test", true),
		}
		job := model.Job{
			Type:  model.Scan,
			State: model.Ready,
			Attachment: &model.SourceAttachment{
				Preparation: &model.Preparation{
					MaxSize: 2_000_000,
				},
				Storage: &model.Storage{
					Type: "local",
					Path: tmp,
				},
			},
		}
		err := db.Create(&job).Error
		require.NoError(t, err)
		dir := model.Directory{
			AttachmentID: 1,
		}
		err = db.Create(&dir).Error
		require.NoError(t, err)
		err = thread.scan(ctx, job)
		require.NoError(t, err)

		var dirs []model.Directory
		err = db.Find(&dirs).Error
		require.NoError(t, err)
		require.Len(t, dirs, 4)
		var jobs []model.Job
		err = db.Preload("FileRanges").Find(&jobs).Error
		require.NoError(t, err)
		require.Len(t, jobs, 13)
	})
}

func TestToCarSize(t *testing.T) {
	jobs := []struct {
		origin   int64
		estimate int64
		actual   int64
	}{
		{0, 37, 37},
		{1, 40, 38},
		{1001, 1040, 1039},
		{1048576, 1048615, 1048615},
		{1048577, 1048803, 1048795},
		{2097152, 2097378, 2097377},
		{10485760, 10486714, 10486697},
		{104857600, 104866744, 104866548},
	}
	for _, job := range jobs {
		t.Run(fmt.Sprintf("%d", job.origin), func(t *testing.T) {
			require.Equal(t, job.estimate, toCarSize(job.origin))
			require.Equal(t, job.actual, toCarSizeExpensive(t, job.origin))
			require.GreaterOrEqual(t, job.estimate, job.actual)
		})
	}
}

func toCarSizeExpensive(t *testing.T, size int64) int64 {
	written := int64(0)
	writer := bytes.NewBuffer(nil)
	if size == 0 {
		blk, err := blocks.NewBlockWithCid([]byte{}, packutil.EmptyFileCid)
		require.NoError(t, err)
		n, err := packutil.WriteCarBlock(writer, blk)
		require.NoError(t, err)
		written += n
		return written
	}

	var links []format.Link
	for size > 0 {
		blkSize := packutil.ChunkSize
		if size < packutil.ChunkSize {
			blkSize = size
		}
		blk, err := blocks.NewBlockWithCid(make([]byte, blkSize), packutil.EmptyFileCid)
		require.NoError(t, err)
		n, err := packutil.WriteCarBlock(writer, blk)
		require.NoError(t, err)
		written += n
		links = append(links, format.Link{
			Size: uint64(blkSize),
			Cid:  blk.Cid(),
		})
		size -= blkSize
	}

	if len(links) > 1 {
		blks, _, err := packutil.AssembleFileFromLinks(links)
		require.NoError(t, err)
		for _, blk := range blks {
			n, err := packutil.WriteCarBlock(writer, blk)
			require.NoError(t, err)
			written += n
		}
	}
	return written
}
