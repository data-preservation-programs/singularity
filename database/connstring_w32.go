//go:build windows && 386

package database

import (
	"strings"

	"github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var ErrDatabaseNotSupported = errors.New("database not supported")

var logger = log.Logger("database")

func Open(connString string, config *gorm.Config) (*gorm.DB, error) {
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
