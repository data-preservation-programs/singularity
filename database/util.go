package database

import (
	"context"
	"io"
	"strings"
	"time"

	"github.com/avast/retry-go"
	"github.com/cockroachdb/errors"
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
)

func retryOn(err error) bool {
	emsg := err.Error()
	return strings.Contains(emsg, "could not serialize access") || strings.Contains(emsg, "Deadlock found when trying to get lock") || strings.Contains(emsg, "database is locked") || strings.Contains(emsg, "database table is locked")
}

func DoRetry(ctx context.Context, f func() error) error {
	return retry.Do(f, retry.RetryIf(retryOn), retry.LastErrorOnly(true), retry.Context(ctx))
}

type databaseLogger struct {
	level logger2.LogLevel
}

func (d *databaseLogger) LogMode(level logger2.LogLevel) logger2.Interface {
	d.level = level
	return d
}

func (d *databaseLogger) Info(ctx context.Context, s string, i ...any) {
	logger.Infof(s, i...)
}

func (d *databaseLogger) Warn(ctx context.Context, s string, i ...any) {
	logger.Warnf(s, i...)
}

func (d *databaseLogger) Error(ctx context.Context, s string, i ...any) {
	logger.Errorf(s, i...)
}

func (d *databaseLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rowsAffected := fc()
	elapsed := time.Since(begin)
	lvl := logging.LevelDebug
	if len(sql) > 1000 {
		sql = sql[:1000] + "...(trimmed)"
	}
	if elapsed > time.Second {
		lvl = logging.LevelWarn
		sql = "[SLOW!] " + sql
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) && !strings.Contains(err.Error(), "could not serialize access") {
		lvl = logging.LevelError
	}

	// Uncomment for logging everything in testing
	// lvl = logging.LevelError
	switch lvl {
	case logging.LevelDebug:
		logger.Debugw(sql, "rowsAffected", rowsAffected, "elapsed", elapsed, "err", err)
	case logging.LevelWarn:
		logger.Warnw(sql, "rowsAffected", rowsAffected, "elapsed", elapsed, "err", err)
	case logging.LevelError:
		logger.Errorw(sql, "rowsAffected", rowsAffected, "elapsed", elapsed, "err", err)
	}
}

func OpenWithLogger(connString string) (*gorm.DB, io.Closer, error) {
	gormLogger := databaseLogger{
		level: logger2.Info,
	}
	return open(connString, &gorm.Config{
		Logger:         &gormLogger,
		TranslateError: true,
	})
}

func OpenFromCLI(c *cli.Context) (*gorm.DB, io.Closer, error) {
	connString := c.String("database-connection-string")
	return OpenWithLogger(connString)
}

var ErrInmemoryWithHighConcurrency = errors.New("cannot use in-memory database with concurrency > 1")

var ErrDatabaseNotSupported = errors.New("database not supported")

var logger = logging.Logger("database")
