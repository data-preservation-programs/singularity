//go:build windows && 386

package database

import (
	"context"
	"io"

	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

var TestConnectionString = "mysql://singularity:singularity@tcp(localhost:3306)/singularity?parseTime=true"

func OpenInMemory() (*gorm.DB, io.Closer, error) {
	db, closer, err := OpenWithLogger(TestConnectionString)
	if err != nil {
		logger.Error(err)
		return nil, nil, errors.WithStack(err)
	}

	ctx := context.Background()
	err = DoRetry(ctx, func() error { return model.DropAll(db) })
	if err != nil {
		logger.Error(err)
		closer.Close()
		return nil, nil, errors.WithStack(err)
	}
	err = DoRetry(ctx, func() error { return model.AutoMigrate(db) })
	if err != nil {
		logger.Error(err)
		closer.Close()
		return nil, nil, errors.WithStack(err)
	}

	return db, closer, nil
}

var SupportedTestDialects = []string{"postgres", "mysql"}
