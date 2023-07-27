package database

import (
	"context"
	"strings"
	"time"

	"github.com/avast/retry-go"
	"github.com/data-preservation-programs/singularity/model"
	logging "github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
)

func retryOn(err error) bool {
	return strings.Contains(err.Error(), "database is locked") || strings.Contains(err.Error(), "database table is locked")
}

func DoRetry(f func() error) error {
	return retry.Do(f, retry.RetryIf(retryOn), retry.LastErrorOnly(true))
}

type databaseLogger struct {
	level  logger2.LogLevel
	logger *logging.ZapEventLogger
}

func (d *databaseLogger) LogMode(level logger2.LogLevel) logger2.Interface {
	d.level = level
	return d
}

func (d *databaseLogger) Info(ctx context.Context, s string, i ...any) {
	d.logger.Infof(s, i...)
}

func (d *databaseLogger) Warn(ctx context.Context, s string, i ...any) {
	d.logger.Warnf(s, i...)
}

func (d *databaseLogger) Error(ctx context.Context, s string, i ...any) {
	d.logger.Errorf(s, i...)
}

func (d *databaseLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rowsAffected := fc()
	elapsed := time.Since(begin)
	isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	lvl := logging.LevelDebug
	if elapsed > time.Second {
		lvl = logging.LevelWarn
		sql = "[SLOW!] " + sql
	}
	if err != nil && !isNotFound {
		lvl = logging.LevelError
	}

	switch lvl {
	case logging.LevelDebug:
		if len(sql) > 1000 {
			sql = sql[:1000] + "...(trimmed)"
		}
		d.logger.Debugw(sql, "rowsAffected", rowsAffected, "elapsed", elapsed, "err", err)
	case logging.LevelWarn:
		d.logger.Warnw(sql, "rowsAffected", rowsAffected, "elapsed", elapsed, "err", err)
	case logging.LevelError:
		d.logger.Errorw(sql, "rowsAffected", rowsAffected, "elapsed", elapsed, "err", err)
	}
}

func OpenWithDefaults(connString string) (*gorm.DB, error) {
	logger.Debug("Opening database: ", connString)
	gormLogger := databaseLogger{
		level:  logger2.Info,
		logger: logger,
	}
	return Open(connString, &gorm.Config{
		Logger:         &gormLogger,
		TranslateError: true,
	})
}

func OpenFromCLI(c *cli.Context) (*gorm.DB, error) {
	connString := c.String("database-connection-string")
	return OpenWithDefaults(connString)
}

func OpenInMemory() (*gorm.DB, error) {
	db, err := OpenWithDefaults("sqlite:file::memory:?cache=shared")
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	err = DoRetry(func() error { return model.DropAll(db) })
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	err = DoRetry(func() error { return model.AutoMigrate(db) })
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return db, nil
}

func DropAll(db *gorm.DB) error {
	return model.DropAll(db)
}

func FindDatasetByName(db *gorm.DB, name string) (model.Dataset, error) {
	var dataset model.Dataset
	err := db.Where("name = ?", name).First(&dataset).Error
	return dataset, err
}
