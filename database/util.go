package database

import (
	"context"
	"io"
	"strings"
	"time"

	"github.com/avast/retry-go"
	"github.com/cockroachdb/errors"
	logging "github.com/ipfs/go-log/v2"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
)

// SQL state value that appears in serialization errors from all DBs.
const sqlSerializationFailure = "40001"

var logger = logging.Logger("database")

var (
	ErrInmemoryWithHighConcurrency = errors.New("cannot use in-memory database with concurrency > 1")
	ErrDatabaseNotSupported        = errors.New("database not supported")
)

func DoRetry(ctx context.Context, f func() error) error {
	attempt := 0
	return retry.Do(func() error {
		attempt++
		err := f()
		if err != nil {
			logger.Debugw("db op failed", "attempt", attempt, "err", err)
			logPgError(err)
		}
		return err
	}, retry.RetryIf(retryOn), retry.LastErrorOnly(true), retry.Context(ctx))
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
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) && !strings.Contains(err.Error(), sqlSerializationFailure) {
		// Demote noisy missing-table errors during test setup/teardown
		emsg := err.Error()
		if strings.Contains(emsg, "no such table") || strings.Contains(emsg, "does not exist") || strings.Contains(emsg, "doesn't exist") {
			lvl = logging.LevelDebug
		} else {
			lvl = logging.LevelError
		}
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

func logPgError(err error) {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		logger.Warnw("pg error detail",
			"code", pgErr.Code,
			"message", pgErr.Message,
			"detail", pgErr.Detail,
			"hint", pgErr.Hint,
			"constraint", pgErr.ConstraintName,
			"table", pgErr.TableName)
	}
}

func retryOn(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		// only retry on serialization failure or deadlock
		return pgErr.Code == "40001" || pgErr.Code == "40P01"
	}

	// sqlite/mysql fallback
	emsg := err.Error()
	return strings.Contains(emsg, "database is locked") ||
		strings.Contains(emsg, "database table is locked") ||
		strings.Contains(emsg, "Record has changed since last read") ||
		strings.Contains(emsg, "Error 1020 (HY000)")
}
