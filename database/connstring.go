package database

import (
	"io"
	"strings"

	"github.com/cockroachdb/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDatabase(connString string, config *gorm.Config) (*gorm.DB, io.Closer, error) {
	if strings.HasPrefix(connString, "sqlite:") {
		connString = strings.TrimPrefix(connString, "sqlite:")
		newConnString, err := AddPragmaToSQLite(connString)
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}

		// Use our pure Go ModernC SQLite driver
		db, err := gorm.Open(OpenModernC(newConnString), config)
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}
		return db, io.NopCloser(strings.NewReader("")), nil
	}

	if strings.HasPrefix(connString, "postgres:") {
		logger.Info("Opening postgres database")
		db, err := gorm.Open(postgres.Open(connString), config)
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}
		closer, err := db.DB()
		if err != nil {
			return db, nil, errors.WithStack(err)
		}
		return db, closer, nil
	}

	if strings.HasPrefix(connString, "mysql://") {
		logger.Info("Opening mysql database")
		db, err := gorm.Open(mysql.Open(connString[8:]), config)
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}
		closer, err := db.DB()
		if err != nil {
			return db, nil, errors.WithStack(err)
		}
		return db, closer, nil
	}

	return nil, nil, ErrDatabaseNotSupported
}
