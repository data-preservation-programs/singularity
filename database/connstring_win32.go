//go:build windows && 386

package database

import (
	"io"
	"strings"

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
