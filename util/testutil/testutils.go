package testutil

import (
	"context"
	"crypto/rand"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/google/uuid"
	"github.com/ipfs/boxo/util"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

const pattern = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateFixedBytes(length int) []byte {
	patternLen := len(pattern)
	result := make([]byte, length)
	for i := range length {
		result[i] = pattern[i%patternLen]
	}
	return result
}

func GenerateRandomBytes(n int) []byte {
	b := make([]byte, n)
	//nolint:errcheck
	rand.Read(b)
	return b
}

// GenerateUniqueName creates a unique name for testing by combining a prefix with a UUID suffix
func GenerateUniqueName(prefix string) string {
	return prefix + "-" + strings.ReplaceAll(uuid.New().String(), "-", "")
}

func GetFileTimestamp(t *testing.T, path string) int64 {
	t.Helper()
	info, err := os.Stat(path)
	require.NoError(t, err)
	return info.ModTime().UnixNano()
}

var TestCid = cid.NewCidV1(cid.Raw, util.Hash([]byte("test")))

const TestWalletAddr = "f1fib3pv7jua2ockdugtz7viz3cyy6lkhh7rfx3sa"

const TestPrivateKeyHex = "7b2254797065223a22736563703235366b31222c22507269766174654b6579223a226b35507976337148327349586343595a58594f5775453149326e32554539436861556b6c4e36695a5763453d227d"

// TestLotusAPI is the Lotus API endpoint to use for tests
// Using /rpc/v1 (stable) instead of deprecated /rpc/v0
const TestLotusAPI = "https://api.node.glif.io/rpc/v1"

// SkipIfNotExternalAPI skips the test if SINGULARITY_TEST_EXTERNAL_API is not set
// Use this for tests that make external API calls (e.g., Lotus/Filecoin APIs)
// These are skipped by default because external APIs may be unreliable or rate-limited
func SkipIfNotExternalAPI(t *testing.T) {
	t.Helper()
	if os.Getenv("SINGULARITY_TEST_EXTERNAL_API") == "" {
		t.Skip("Skipping external API test. Set SINGULARITY_TEST_EXTERNAL_API=true to run")
	}
}

func EscapePath(p string) string {
	return "'" + strings.ReplaceAll(p, `\`, `\\`) + "'"
}

func getTestDB(t *testing.T, dialect string) (db *gorm.DB, closer io.Closer, connStr string) {
	t.Helper()
	var err error
	if dialect == "sqlite" {
		connStr = "sqlite:" + t.TempDir() + "/singularity.db"
		db, closer, err = database.OpenWithLogger(connStr)
		require.NoError(t, err)
		return
	}
	dbName := "test_" + strings.ReplaceAll(uuid.New().String(), "-", "")
	if dialect != "postgres" {
		require.Fail(t, "Unsupported dialect: "+dialect)
	}
	pgPort := os.Getenv("PGPORT")
	connStr = "postgres://singularity@localhost:" + pgPort + "/postgres?sslmode=disable"
	cmd := exec.Command("createdb", "-h", "localhost", "-p", pgPort, "-U", "singularity", "-E", "UTF8", dbName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Logf("Failed to create PostgreSQL database %s: %v, output: %s", dbName, err, string(output))
		return nil, nil, ""
	}
	t.Logf("Created PostgreSQL database %s", dbName)
	connStr = strings.ReplaceAll(connStr, "postgres?", dbName+"?")
	var closer2 io.Closer
	db, closer2, err = database.OpenWithLogger(connStr)
	if err != nil {
		t.Logf("Failed to connect to test database %s: %v", dbName, err)
		exec.Command("dropdb", "-h", "localhost", "-p", pgPort, "-U", "singularity", dbName).Run()
		return nil, nil, ""
	}
	closer = CloserFunc(func() error {
		if closer2 != nil {
			_ = closer2.Close()
		}
		exec.Command("dropdb", "-h", "localhost", "-p", pgPort, "-U", "singularity", dbName).Run()
		return nil
	})
	return
}

func One(t *testing.T, testFunc func(ctx context.Context, t *testing.T, db *gorm.DB)) {
	t.Helper()
	backend := SupportedTestDialects[0]
	doOne(t, backend, testFunc)
}

func OneWithoutReset(t *testing.T, testFunc func(ctx context.Context, t *testing.T, db *gorm.DB)) {
	t.Helper()
	backend := SupportedTestDialects[0]
	db, closer, connStr := getTestDB(t, backend)
	if db == nil {
		t.Skip("Skip " + backend + " - database not available")
		return
	}
	defer closer.Close()
	t.Setenv("DATABASE_CONNECTION_STRING", connStr)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	db = db.WithContext(ctx)

	t.Run(backend, func(t *testing.T) {
		testFunc(ctx, t, db)
	})
}

func doOne(t *testing.T, backend string, testFunc func(ctx context.Context, t *testing.T, db *gorm.DB)) {
	t.Helper()
	db, closer, connStr := getTestDB(t, backend)
	if db == nil {
		t.Skip("Skip " + backend + " - database not available")
		return
	}
	defer closer.Close()
	t.Setenv("DATABASE_CONNECTION_STRING", connStr)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	db = db.WithContext(ctx)

	err := model.AutoMigrate(db)
	require.NoError(t, err)

	// Clear any existing data from tables with unique constraints
	tables := []string{
		"output_attachments",
		"source_attachments",
		"storages",
		"wallets",
		"deal_schedules",
		"preparations",
	}

	// Get DB type from connection string
	isPostgres := strings.HasPrefix(connStr, "postgres:")
	for _, table := range tables {
		var err error
		if isPostgres {
			err = db.Exec("TRUNCATE TABLE " + table + " CASCADE").Error
		} else {
			err = db.Exec("DELETE FROM " + table).Error
		}
		if err != nil {
			emsg := err.Error()
			// Suppress noisy logs when tables don't exist yet across backends
			if strings.Contains(emsg, "no such table") || strings.Contains(emsg, "does not exist") || strings.Contains(emsg, "doesn't exist") {
				continue
			}
			t.Logf("Warning: Failed to clear table %s: %v", table, err)
			// Don't fail the test for other errors either
		}
	}

	t.Run(backend, func(t *testing.T) {
		testFunc(ctx, t, db)
	})
}

func All(t *testing.T, testFunc func(ctx context.Context, t *testing.T, db *gorm.DB)) {
	t.Helper()
	for _, backend := range SupportedTestDialects {
		doOne(t, backend, testFunc)
	}
}

type CloserFunc func() error

func (f CloserFunc) Close() error {
	return f()
}
