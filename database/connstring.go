//go:build !cgo

package database

import (
	"io"
	"strings"

	"github.com/glebarez/sqlite"
	"github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ErrDatabaseNotSupported = errors.New("database not supported")

var logger = log.Logger("database")

func Open(connString string, config *gorm.Config) (*gorm.DB, io.Closer, error) {
	var db *gorm.DB
	var closer io.Closer
	var err error
	if strings.HasPrefix(connString, "sqlite:") {
		logger.Info("Opening sqlite database (non-cgo version)")
		db, err = gorm.Open(sqlite.Open(connString[7:]), config)
		if err != nil {
			return nil, nil, err
		}

		closer, err = db.DB()
		if err != nil {
			return nil, nil, err
		}

		err = db.Exec("PRAGMA foreign_keys = ON").Error
		if err != nil {
			closer.Close()
			return nil, nil, err
		}

		err = db.Exec("PRAGMA busy_timeout = 50000").Error
		if err != nil {
			closer.Close()
			return nil, nil, err
		}

		err = db.Exec("PRAGMA journal_mode = WAL").Error
		if err != nil {
			closer.Close()
			return nil, nil, err
		}

		return db, closer, nil
	}

	if strings.HasPrefix(connString, "postgres:") {
		logger.Info("Opening postgres database")
		db, err = gorm.Open(postgres.Open(connString), config)
		if err != nil {
			return nil, nil, err
		}
		closer, err = db.DB()
		return db, closer, err
	}

	if strings.HasPrefix(connString, "mysql://") {
		logger.Info("Opening mysql database")
		db, err = gorm.Open(mysql.Open(connString[8:]), config)
		if err != nil {
			return nil, nil, err
		}
		closer, err = db.DB()
		return db, closer, err
	}

	return nil, nil, ErrDatabaseNotSupported
}
