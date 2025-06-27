package contentprovider

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/ipfs/boxo/util"
	"github.com/ipfs/go-cid"
	"github.com/labstack/echo/v4"
	"github.com/multiformats/go-varint"
	"github.com/parnurzeal/gorequest"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestHTTPServerStart(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		s := HTTPServer{
			dbNoContext:         db,
			bind:                "127.0.0.1:65432",
			enablePiece:         true,
			enablePieceMetadata: true,
		}
		require.Equal(t, "HTTPServer", s.Name())
		exitErr := make(chan error, 1)
		ctx, cancel := context.WithCancel(ctx)
		err := s.Start(ctx, exitErr)
		require.NoError(t, err)
		time.Sleep(200 * time.Millisecond)
		gorequest.New().Get("http://127.0.0.1:65432/health").End()

		cancel()
		select {
		case <-time.After(1 * time.Second):
			t.Fatal("http server did not stop")
		case err = <-exitErr:
			require.NoError(t, err)
		}
	})
}

func TestHTTPServerHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		e := echo.New()
		s := HTTPServer{
			dbNoContext:         db,
			bind:                ":0",
			enablePiece:         true,
			enablePieceMetadata: true,
		}

		pieceCID := cid.NewCidV1(cid.FilCommitmentUnsealed, util.Hash([]byte("test")))
		err := db.Create(&model.Car{
			PieceCID:      model.CID(pieceCID),
			PieceSize:     128,
			FileSize:      59 + 1 + 36 + 5,
			StoragePath:   "",
			PreparationID: 1,
			PieceType:     model.DataPiece,
			Attachment: &model.SourceAttachment{
				Preparation: &model.Preparation{},
				Storage: &model.Storage{
					Type: "local",
				},
			},
			RootCID: model.CID(testutil.TestCid),
		}).Error
		require.NoError(t, err)
		err = db.Create(&model.CarBlock{
			CarID:          1,
			CID:            model.CID(testutil.TestCid),
			CarOffset:      59,
			CarBlockLength: 1 + 36 + 5,
			Varint:         varint.ToUvarint(36 + 5),
			RawBlock:       []byte("hello"),
		}).Error
		require.NoError(t, err)

		tests := []struct {
			name string
			cid  string
			code int
			body string
			cbor bool
		}{
			{
				name: "not_found",
				cid:  cid.NewCidV1(cid.FilCommitmentUnsealed, util.Hash([]byte("not_exist"))).String(),
				code: http.StatusNotFound,
				body: "piece not found",
			},
			{
				name: "invalid_cid",
				cid:  "invalid",
				code: http.StatusBadRequest,
				body: "failed to parse",
			},
			{
				name: "success",
				cid:  pieceCID.String(),
				code: http.StatusOK,
				body: "",
			},
			{
				name: "success with cbor",
				cid:  pieceCID.String(),
				code: http.StatusOK,
				body: "",
				cbor: true,
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				req := httptest.NewRequest(http.MethodGet, "/piece/metadata/:id", nil)
				if test.cbor {
					req.Header.Set("Accept", "application/cbor")
				}
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/piece/metadata/:id")
				c.SetParamNames("id")
				c.SetParamValues(test.cid)
				err = s.getMetadataHandler(c)
				require.NoError(t, err)
				require.Equal(t, test.code, rec.Code)
				require.Contains(t, rec.Body.String(), test.body)
				if test.cbor {
					require.Equal(t, "application/cbor", rec.Header().Get(echo.HeaderContentType))
				}

				// For successful responses, validate the piece_type field
				if test.code == http.StatusOK && !test.cbor {
					var metadata PieceMetadata
					err = json.Unmarshal(rec.Body.Bytes(), &metadata)
					require.NoError(t, err)
					require.Equal(t, model.DataPiece, metadata.Car.PieceType)
				}
			})

			t.Run(test.name, func(t *testing.T) {
				req := httptest.NewRequest(http.MethodGet, "/piece/:id", nil)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/piece/:id")
				c.SetParamNames("id")
				c.SetParamValues(test.cid)
				err = s.handleGetPiece(c)
				require.NoError(t, err)
				require.Equal(t, test.code, rec.Code)
			})
		}

		// Test DAG piece type
		t.Run("dag_piece_metadata", func(t *testing.T) {
			preparation := &model.Preparation{Name: "test_prep_dag"}
			err := db.Create(preparation).Error
			require.NoError(t, err)

			storage := &model.Storage{Name: "test_storage_dag", Type: "local"}
			err = db.Create(storage).Error
			require.NoError(t, err)

			attachment := &model.SourceAttachment{
				PreparationID: preparation.ID,
				StorageID:     storage.ID,
			}
			err = db.Create(attachment).Error
			require.NoError(t, err)

			dagPieceCID := cid.NewCidV1(cid.FilCommitmentUnsealed, util.Hash([]byte("dag_test")))
			err = db.Create(&model.Car{
				PieceCID:      model.CID(dagPieceCID),
				PieceSize:     256,
				PreparationID: preparation.ID,
				PieceType:     model.DagPiece,
				AttachmentID:  &attachment.ID,
				RootCID:       model.CID(testutil.TestCid),
			}).Error
			require.NoError(t, err)

			req := httptest.NewRequest(http.MethodGet, "/piece/metadata/:id", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/piece/metadata/:id")
			c.SetParamNames("id")
			c.SetParamValues(dagPieceCID.String())
			err = s.getMetadataHandler(c)
			require.NoError(t, err)
			require.Equal(t, http.StatusOK, rec.Code)

			var metadata PieceMetadata
			err = json.Unmarshal(rec.Body.Bytes(), &metadata)
			require.NoError(t, err)
			require.Equal(t, model.DagPiece, metadata.Car.PieceType)
		})

		// Add car file
		tmp := t.TempDir()
		err = db.Model(&model.Car{}).Where("id = ?", 1).Update("storage_path", filepath.Join(tmp, "test.car")).Error
		testfunc := func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/piece/:id", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/piece/:id")
			c.SetParamNames("id")
			c.SetParamValues(pieceCID.String())
			err = s.handleGetPiece(c)
			require.NoError(t, err)
			require.Equal(t, http.StatusOK, rec.Code)
			require.EqualValues(t, 101, rec.Body.Len())
		}
		t.Run("car file deleted, fail back to inline", testfunc)

		err = os.WriteFile(filepath.Join(tmp, "test.car"), []byte("test"), 0644)
		require.NoError(t, err)
		t.Run("car file changed, fail back to inline", testfunc)

		err = os.WriteFile(filepath.Join(tmp, "test.car"), testutil.GenerateRandomBytes(48), 0644)
		require.NoError(t, err)
		t.Run("car file exists", testfunc)
	})
}
