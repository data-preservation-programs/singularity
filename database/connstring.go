//go:build !cgo && !386

package database

import (
	"io"
	"net/url"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func open(connString string, config *gorm.Config) (*gorm.DB, io.Closer, error) {
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
		return db, closer, errors.WithStack(err)
	}

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

func AddPragmaToSQLite(connString string) (string, error) {
	u, err := url.Parse(connString)
	if err != nil {
		return "", errors.WithStack(err)
	}

	qs := u.Query()
	qs.Add("_pragma", "busy_timeout(50000)")
	qs.Set("_pragma", "foreign_keys(1)")
	if strings.HasPrefix(connString, "file::memory:") {
		qs.Set("_pragma", "journal_mode(MEMORY)")
		qs.Set("mode", "memory")
		qs.Set("cache", "shared")
	} else {
		qs.Set("_pragma", "journal_mode(WAL)")
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}
