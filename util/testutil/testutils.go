package testutil

import (
	"context"
	"crypto/rand"
	"io"
	"net"
	"os"
	"strings"
	"testing"
	"time"

	rand2 "math/rand"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/boxo/util"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

const pattern = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateFixedBytes(length int) []byte {
	patternLen := len(pattern)
	result := make([]byte, length)
	for i := 0; i < length; i++ {
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

func RandomLetterString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"

	b := make([]byte, length)
	for i := range b {
		//nolint:gosec
		b[i] = charset[rand2.Intn(len(charset))]
	}
	return string(b)
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
	dbName := RandomLetterString(6)
	var opError *net.OpError
	switch dialect {
	case "mysql":
		connStr = "mysql://singularity:singularity@tcp(localhost:3306)/singularity?parseTime=true"
	case "postgres":
		connStr = "postgres://singularity:singularity@localhost:5432/singularity?sslmode=disable"
	default:
		require.Fail(t, "Unsupported dialect: "+dialect)
	}
	var db1 *gorm.DB
	var closer1 io.Closer
	db1, closer1, err = database.OpenWithLogger(connStr)
	if errors.As(err, &opError) {
		return
	}
	err = db1.Exec("CREATE DATABASE " + dbName + "").Error
	require.NoError(t, err)
	connStr = strings.ReplaceAll(connStr, "singularity?", dbName+"?")
	var closer2 io.Closer
	db, closer2, err = database.OpenWithLogger(connStr)
	require.NoError(t, err)
	closer = CloserFunc(func() error {
		require.NoError(t, closer2.Close())
		require.NoError(t, db1.Exec("DROP DATABASE "+dbName+"").Error)
		return closer1.Close()
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
		t.Log("Skip " + backend)
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
		t.Log("Skip " + backend)
		return
	}
	defer closer.Close()
	t.Setenv("DATABASE_CONNECTION_STRING", connStr)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	db = db.WithContext(ctx)

	err := model.AutoMigrate(db)
	require.NoError(t, err)

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
