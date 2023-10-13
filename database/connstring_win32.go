//go:build 386

package database

import (
	"io"
	"strings"

	"github.com/cockroachdb/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func open(connString string, config *gorm.Config) (*gorm.DB, io.Closer, error) {
	var db *gorm.DB
	var closer io.Closer
	var err error

	if strings.HasPrefix(connString, "postgres:") {
		logger.Info("Opening postgres database")
		db, err = gorm.Open(postgres.Open(connString), config)
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}
		closer, err = db.DB()
		return db, closer, errors.WithStack(err)
	}

	if strings.HasPrefix(connString, "mysql://") {
		logger.Info("Opening mysql database")
		db, err = gorm.Open(mysql.Open(connString[8:]), config)
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}
		closer, err = db.DB()
		return db, closer, errors.WithStack(err)
	}

	return nil, nil, ErrDatabaseNotSupported
}
