package datasetworker

import (
	"context"
	"os"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack/daggen"
	"github.com/google/uuid"
	"github.com/ipfs/boxo/util"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
)

func TestExportDag(t *testing.T) {
	ctx := context.Background()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	thread := &Thread{
		dbNoContext: db,
		logger:      logger.With("test", true),
		id:          uuid.New(),
	}

	source := model.Source{
		Dataset: &model.Dataset{},
	}
	err = db.Create(&source).Error
	require.NoError(t, err)

	rootData := daggen.NewDirectoryData()
	err = rootData.AddItem(ctx, "test1.txt", cid.NewCidV1(cid.Raw, util.Hash([]byte("test1"))), 5)
	require.NoError(t, err)
	rootDataBytes, err := rootData.MarshalBinary(ctx)
	require.NoError(t, err)
	node, err := rootData.Node()
	require.NoError(t, err)
	rootDir := model.Directory{
		SourceID: source.ID,
		CID:      model.CID(node.Cid()),
		Data:     rootDataBytes,
	}
	err = db.Create(&rootDir).Error
	require.NoError(t, err)

	err = thread.ExportDag(ctx, source)
	require.NoError(t, err)

	tmp := t.TempDir()
	source.Dataset.OutputDirs = []string{tmp}
	err = thread.ExportDag(ctx, source)
	require.NoError(t, err)
	entries, err := os.ReadDir(tmp)
	require.NoError(t, err)
	require.Len(t, entries, 0)

	err = db.Model(&model.Directory{}).Where("id = ?", rootDir.ID).Update("exported", false).Error
	require.NoError(t, err)
	err = thread.ExportDag(ctx, source)
	require.NoError(t, err)
	entries, err = os.ReadDir(tmp)
	require.NoError(t, err)
	require.Len(t, entries, 1)
	require.Equal(t, "baga6ea4seaqfilevbbbguevziwwafurgltselpzbzpsu5lrlyqq7rbmcgeyqemy.car", entries[0].Name())
}
