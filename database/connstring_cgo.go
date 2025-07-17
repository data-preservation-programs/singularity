//go:build cgo && !386 && !darwin

package database

import (
	"io"
	"strings"

	"github.com/cockroachdb/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)



func OpenDatabase(connString string, config *gorm.Config) (*gorm.DB, io.Closer, error) {
	var db *gorm.DB
	var closer io.Closer
	var err error
	if strings.HasPrefix(connString, "sqlite:") {
		connString, err = AddPragmaToSQLite(connString[7:])
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}
		db, err = gorm.Open(sqlite.Open(connString), config)
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}
		closer, err = db.DB()
		if err != nil {
			return db, nil, errors.WithStack(err)
		}
		return db, closer, nil
	}

	if strings.HasPrefix(connString, "postgres:") {
		logger.Info("Opening postgres database")
		db, err = gorm.Open(postgres.Open(connString), config)
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}
		closer, err = db.DB()
		if err != nil {
			return db, nil, errors.WithStack(err)
		}
		return db, closer, nil
	}

	if strings.HasPrefix(connString, "mysql://") {
		logger.Info("Opening mysql database")
		db, err = gorm.Open(mysql.Open(connString[8:]), config)
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}
		closer, err = db.DB()
		if err != nil {
			return db, nil, errors.WithStack(err)
		}
		return db, closer, nil
	}

	return nil, nil, ErrDatabaseNotSupported
}
