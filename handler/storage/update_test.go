package storage

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/gotidy/ptr"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestUpdateStorageHandler(t *testing.T) {
	t.Run("storage not found", func(t *testing.T) {
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			_, err := Default.UpdateStorageHandler(ctx, db, "test", UpdateRequest{Config: nil})
			require.ErrorIs(t, err, handlererror.ErrNotFound)
		})
	})
	t.Run("change local path config", func(t *testing.T) {
		for _, name := range []string{"1", "name"} {
			t.Run(name, func(t *testing.T) {
				testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
					tmp := t.TempDir()
					_, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", "name", tmp, nil, model.ClientConfig{}})
					require.NoError(t, err)
					storage, err := Default.UpdateStorageHandler(ctx, db, name, UpdateRequest{Config: map[string]string{
						"copy_links": "true",
					}})
					require.NoError(t, err)
					require.Equal(t, "true", storage.Config["copy_links"])
				})
			})
		}
	})
	t.Run("change client config", func(t *testing.T) {
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			tmp := t.TempDir()
			_, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", "name", tmp, nil, model.ClientConfig{
				ConnectTimeout:        ptr.Of(int64(time.Minute)),
				Timeout:               ptr.Of(int64(time.Minute)),
				ExpectContinueTimeout: ptr.Of(int64(time.Minute)),
				InsecureSkipVerify:    ptr.Of(true),
				NoGzip:                ptr.Of(true),
				UserAgent:             ptr.Of("1"),
				CaCert:                []string{"1"},
				ClientCert:            ptr.Of("1"),
				ClientKey:             ptr.Of("1"),
				Headers:               map[string]string{"a": "b"},
				DisableHTTP2:          ptr.Of(true),
				DisableHTTPKeepAlives: ptr.Of(true),
				LowLevelRetries:       ptr.Of(20),
				UseServerModTime:      ptr.Of(true),
				ScanConcurrency:       ptr.Of(10),
			}})
			require.NoError(t, err)
			newConfig := model.ClientConfig{
				ConnectTimeout:        ptr.Of(int64(time.Hour)),
				Timeout:               ptr.Of(int64(time.Hour)),
				ExpectContinueTimeout: ptr.Of(int64(time.Hour)),
				InsecureSkipVerify:    ptr.Of(false),
				NoGzip:                ptr.Of(false),
				UserAgent:             ptr.Of("0"),
				CaCert:                []string{"0"},
				ClientCert:            ptr.Of("0"),
				ClientKey:             ptr.Of("0"),
				Headers:               map[string]string{"a": "c"},
				DisableHTTP2:          ptr.Of(false),
				DisableHTTPKeepAlives: ptr.Of(false),
				LowLevelRetries:       ptr.Of(10),
				UseServerModTime:      ptr.Of(false),
				ScanConcurrency:       ptr.Of(1),
			}
			storage, err := Default.UpdateStorageHandler(ctx, db, "name", UpdateRequest{ClientConfig: newConfig})
			require.NoError(t, err)
			require.EqualValues(t, newConfig, storage.ClientConfig)
		})
	})
	t.Run("clear client config", func(t *testing.T) {
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			tmp := t.TempDir()
			_, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", "name", tmp, nil, model.ClientConfig{
				UserAgent:  ptr.Of("1"),
				CaCert:     []string{"1"},
				ClientCert: ptr.Of("1"),
				ClientKey:  ptr.Of("1"),
				Headers:    map[string]string{"a": "b"},
			}})
			require.NoError(t, err)
			newConfig := model.ClientConfig{
				UserAgent:  ptr.Of(""),
				CaCert:     []string{""},
				ClientCert: ptr.Of(""),
				ClientKey:  ptr.Of(""),
				Headers:    map[string]string{"a": ""},
			}
			storage, err := Default.UpdateStorageHandler(ctx, db, "name", UpdateRequest{ClientConfig: newConfig})
			require.NoError(t, err)
			require.EqualValues(t, model.ClientConfig{Headers: map[string]string{}}, storage.ClientConfig)
		})
	})
	t.Run("invalid config", func(t *testing.T) {
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			tmp := t.TempDir()
			_, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", "name", tmp, nil, model.ClientConfig{}})
			require.NoError(t, err)
			_, err = Default.UpdateStorageHandler(ctx, db, "name", UpdateRequest{Config: map[string]string{
				"copy_links": "invalid",
			}})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		})
	})
}
