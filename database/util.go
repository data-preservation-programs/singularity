package database

import (
	"context"
	"github.com/data-preservation-programs/go-singularity/model"
	logging "github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type databaseLogger struct {
	level  logger2.LogLevel
	logger *logging.ZapEventLogger
}

func (d *databaseLogger) LogMode(level logger2.LogLevel) logger2.Interface {
	d.level = level
	return d
}

func (d *databaseLogger) Info(ctx context.Context, s string, i ...interface{}) {
	d.logger.Infof(s, i...)
}

func (d *databaseLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	d.logger.Warnf(s, i...)
}

func (d *databaseLogger) Error(ctx context.Context, s string, i ...interface{}) {
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
		d.logger.Debugw(sql, "rowsAffected", rowsAffected, "elapsed", elapsed, "err", err)
	case logging.LevelWarn:
		d.logger.Warnw(sql, "rowsAffected", rowsAffected, "elapsed", elapsed, "err", err)
	case logging.LevelError:
		d.logger.Errorw(sql, "rowsAffected", rowsAffected, "elapsed", elapsed, "err", err)
	}
}

func MustOpenFromCLI(c *cli.Context) *gorm.DB {
	connString := c.String("database-connection-string")
	logger.Debug("Opening database: ", connString)
	gormLogger := databaseLogger{
		level:  logger2.Warn,
		logger: logging.Logger("database"),
	}
	if c.Bool("verbose") {
		gormLogger.level = logger2.Info
	}
	db, err := Open(connString, &gorm.Config{
		Logger: &gormLogger,
	})
	if err != nil {
		logger.Panic(err)
	}

	if err != nil {
		logger.Panic(err)
	}

	return db
}

func OpenInMemory() *gorm.DB {
	gormLogger := logger2.New(log.New(os.Stderr, "\r\n", log.LstdFlags), logger2.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  logger2.Error,
		IgnoreRecordNotFoundError: false,
		Colorful:                  true,
	})
	db, err := Open("sqlite:file::memory:?cache=shared", &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		logger.Panic(err)
	}

	err = model.AutoMigrate(db)
	if err != nil {
		logger.Panic(err)
	}

	return db
}

func DropAll(db *gorm.DB) error {
	return model.DropAll(db)
}

func FindDatasetByName(db *gorm.DB, name string) (model.Dataset, error) {
	var dataset model.Dataset
	err := db.Where(&model.Dataset{Name: name}).First(&dataset).Error
	return dataset, err
}
