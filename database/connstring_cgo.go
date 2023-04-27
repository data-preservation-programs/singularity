//go:build cgo

package database

import (
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"strings"
)

var DatabaseNotSupportedError = errors.New("database not supported")

var logger = log.Logger("database")

func Open(connString string, config *gorm.Config) (*gorm.DB, error) {
	if strings.HasPrefix(connString, "sqlite:") {
		return gorm.Open(sqlite.Open(connString[7:]), config)
	}

	if strings.HasPrefix(connString, "postgres:") {
		return gorm.Open(postgres.Open(connString), config)
	}

	if strings.HasPrefix(connString, "mysql:") {
		return gorm.Open(mysql.Open(connString[6:]), config)
	}

	if strings.HasPrefix(connString, "sqlserver:") {
		return gorm.Open(sqlserver.Open(connString[10:]), config)
	}

	return nil, DatabaseNotSupportedError
}

func MustOpenFromCLI(c *cli.Context) *gorm.DB {
	connString := c.String("database-connection-string")
	db, err := Open(connString, &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
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
