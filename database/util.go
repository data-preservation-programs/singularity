package database

import (
	"context"
	"io"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	logging "github.com/ipfs/go-log/v2"
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

func retryOn(err error) bool {
	emsg := err.Error()
	return strings.Contains(emsg, sqlSerializationFailure) || strings.Contains(emsg, "database is locked") || strings.Contains(emsg, "database table is locked")
}

func DoRetry(ctx context.Context, f func() error) error {
	var lastErr error
	retryLimit := 3

	for i := 0; i < retryLimit; i++ {
		err := f()
		if err == nil {
			return nil
		}
		lastErr = err

		if IsUniqueConstraintError(err) {
			return errors.Wrap(handlererror.ErrDuplicateKey, err.Error())
		}

		if !retryOn(err) {
			break
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(50 * time.Millisecond):
			continue
		}
	}

	return lastErr
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
		lvl = logging.LevelError
	}

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

// IsUniqueConstraintError checks if an error is due to a duplicate key constraint violation.
func IsUniqueConstraintError(err error) bool {
	if err == nil {
		return false
	}
	errMsg := err.Error()
	return strings.Contains(errMsg, "duplicate key") ||
		strings.Contains(errMsg, "duplicated key") ||
		strings.Contains(errMsg, "UNIQUE constraint failed") ||
		strings.Contains(errMsg, "violates unique constraint")
}
