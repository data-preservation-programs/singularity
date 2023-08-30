//go:build windows && 386

package testutil

import (
	"context"
	"io"
	"net"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

var SupportedTestDialects = []string{"mysql", "postgres"}

func One(t *testing.T, testFunc func(ctx context.Context, t *testing.T, db *gorm.DB)) {
	t.Helper()
	backend := SupportedTestDialects[0]
	doOne(t, backend, testFunc)
}

func doOne(t *testing.T, backend string, testFunc func(ctx context.Context, t *testing.T, db *gorm.DB)) {
	t.Helper()
	db, closer, connStr := getTestDB(t, backend)
	if db == nil {
		t.Log("Skip " + backend)
		return
	}
	defer closer.Close()
	os.Setenv("DATABASE_CONNECTION_STRING", connStr)
	defer func() {
		os.Unsetenv("DATABASE_CONNECTION_STRING")
	}()
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

func getTestDB(t *testing.T, dialect string) (db *gorm.DB, closer io.Closer, connStr string) {
	t.Helper()
	var err error
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
	closer = CloserFunc(func() error {
		require.NoError(t, closer2.Close())
		require.NoError(t, db1.Exec("DROP DATABASE "+dbName+"").Error)
		return closer1.Close()
	})
	return
}
