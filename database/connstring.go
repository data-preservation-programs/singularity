//go:build !cgo

package database

import (
	"github.com/glebarez/sqlite"
	"github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"strings"
)

var ErrDatabaseNotSupported = errors.New("database not supported")

var logger = log.Logger("database")

func Open(connString string, config *gorm.Config) (*gorm.DB, error) {
	if strings.HasPrefix(connString, "sqlite:") {
		db, err := gorm.Open(sqlite.Open(connString[7:]), config)
		if err != nil {
			return nil, err
		}

		err = db.Exec("PRAGMA foreign_keys = ON").Error
		if err != nil {
			return nil, err
		}

		err = db.Exec("PRAGMA busy_timeout = 5000").Error
		if err != nil {
			return nil, err
		}

		err = db.Exec("PRAGMA journal_mode = WAL").Error
		if err != nil {
			return nil, err
		}

		return db, nil
	}

	if strings.HasPrefix(connString, "postgres:") {
		return gorm.Open(postgres.Open(connString), config)
	}

	if strings.HasPrefix(connString, "mysql://") {
		return gorm.Open(mysql.Open(connString[8:]), config)
	}

	if strings.HasPrefix(connString, "sqlserver://") {
		return gorm.Open(sqlserver.Open(connString), config)
	}

	return nil, ErrDatabaseNotSupported
}
