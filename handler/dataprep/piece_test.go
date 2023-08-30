package dataprep

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack/packutil"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/gotidy/ptr"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestListPiecesHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			SourceStorages: []model.Storage{{}},
		}).Error
		require.NoError(t, err)
		err = db.Create([]model.Car{{
			AttachmentID:  ptr.Of(uint32(1)),
			PreparationID: 1,
		}, {
			PreparationID: 1,
		}}).Error
		require.NoError(t, err)
		result, err := Default.ListPiecesHandler(ctx, db, 1)
		require.NoError(t, err)
		require.Len(t, result, 2)
	})
}

func TestListPiecesHandler_NotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			SourceStorages: []model.Storage{{}},
		}).Error
		require.NoError(t, err)
		_, err = Default.ListPiecesHandler(ctx, db, 2)
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
}

func TestAddPieceHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{}).Error
		require.NoError(t, err)
		t.Run("not_found", func(t *testing.T) {
			_, err := Default.AddPieceHandler(ctx, db, 100, AddPieceRequest{
				PieceCID:  "",
				PieceSize: "",
				FilePath:  "",
				RootCID:   "",
			})
			require.ErrorIs(t, err, handlererror.ErrNotFound)
		})
		t.Run("pieceCID invalid", func(t *testing.T) {
			_, err := Default.AddPieceHandler(ctx, db, 1, AddPieceRequest{
				PieceCID:  "invalid",
				PieceSize: "32",
				FilePath:  "",
				RootCID:   "",
			})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
			require.ErrorContains(t, err, "invalid piece CID")
		})
		t.Run("pieceCID invalid", func(t *testing.T) {
			_, err := Default.AddPieceHandler(ctx, db, 1, AddPieceRequest{
				PieceCID:  packutil.EmptyFileCid.String(),
				PieceSize: "32",
				FilePath:  "",
				RootCID:   "",
			})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
			require.ErrorContains(t, err, "piece CID must be commp")
		})
		t.Run("pieceCID invalid", func(t *testing.T) {
			_, err := Default.AddPieceHandler(ctx, db, 1, AddPieceRequest{
				PieceCID:  "baga6ea4seaqchxeb6cwpiephnus27kplk7lku225rdhrsgb3ej4smaqwgop6wkq",
				PieceSize: "axxx",
				FilePath:  "",
				RootCID:   "",
			})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
			require.ErrorContains(t, err, "invalid piece size")
		})
		t.Run("pieceCID not power of 2", func(t *testing.T) {
			_, err := Default.AddPieceHandler(ctx, db, 1, AddPieceRequest{
				PieceCID:  "baga6ea4seaqchxeb6cwpiephnus27kplk7lku225rdhrsgb3ej4smaqwgop6wkq",
				PieceSize: "3000",
				FilePath:  "",
				RootCID:   "",
			})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
			require.ErrorContains(t, err, "piece size must be a power of 2")
		})
		t.Run("invalid root cid", func(t *testing.T) {
			_, err := Default.AddPieceHandler(ctx, db, 1, AddPieceRequest{
				PieceCID:  "baga6ea4seaqchxeb6cwpiephnus27kplk7lku225rdhrsgb3ej4smaqwgop6wkq",
				PieceSize: "65536",
				FilePath:  "",
				RootCID:   "xxxx",
			})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
			require.ErrorContains(t, err, "invalid root CID")
		})
		t.Run("invalid file path", func(t *testing.T) {
			_, err := Default.AddPieceHandler(ctx, db, 1, AddPieceRequest{
				PieceCID:  "baga6ea4seaqchxeb6cwpiephnus27kplk7lku225rdhrsgb3ej4smaqwgop6wkq",
				PieceSize: "65536",
				FilePath:  "invalid",
				RootCID:   "",
			})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
			require.ErrorContains(t, err, "failed to open file")
		})
		t.Run("invalid car header", func(t *testing.T) {
			tmp := t.TempDir()
			err := os.WriteFile(filepath.Join(tmp, "a.car"), []byte("invalid"), 0644)
			require.NoError(t, err)
			_, err = Default.AddPieceHandler(ctx, db, 1, AddPieceRequest{
				PieceCID:  "baga6ea4seaqchxeb6cwpiephnus27kplk7lku225rdhrsgb3ej4smaqwgop6wkq",
				PieceSize: "65536",
				FilePath:  filepath.Join(tmp, "a.car"),
				RootCID:   "",
			})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
			require.ErrorContains(t, err, "failed to read CAR header")
		})
		t.Run("success", func(t *testing.T) {
			tmp := t.TempDir()
			f, err := os.Create(filepath.Join(tmp, "a.car"))
			require.NoError(t, err)
			_, err = packutil.WriteCarHeader(f, packutil.EmptyFileCid)
			require.NoError(t, err)
			f.Close()
			c, err := Default.AddPieceHandler(ctx, db, 1, AddPieceRequest{
				PieceCID:  "baga6ea4seaqchxeb6cwpiephnus27kplk7lku225rdhrsgb3ej4smaqwgop6wkq",
				PieceSize: "65536",
				FilePath:  filepath.Join(tmp, "a.car"),
				RootCID:   "",
			})
			require.NoError(t, err)
			require.NotNil(t, c)
		})
	})
}
