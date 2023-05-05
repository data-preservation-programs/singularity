package database

import (
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func MustOpenFromCLI(c *cli.Context) *gorm.DB {
	connString := c.String("database-connection-string")
	gormLogger := logger2.New(log.New(os.Stderr, "\r\n", log.LstdFlags), logger2.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  logger2.Warn,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	})
	db, err := Open(connString, &gorm.Config{
		Logger: gormLogger,
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
