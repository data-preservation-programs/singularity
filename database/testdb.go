//go:build !(windows && 386)

package database

import (
	"io"

	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func OpenInMemory() (*gorm.DB, io.Closer, error) {
	db, closer, err := OpenWithLogger("sqlite:file::memory:?cache=shared")
	if err != nil {
		logger.Error(err)
		return nil, nil, err
	}

	err = DoRetry(func() error { return model.DropAll(db) })
	if err != nil {
		logger.Error(err)
		closer.Close()
		return nil, nil, err
	}
	err = DoRetry(func() error { return model.AutoMigrate(db) })
	if err != nil {
		logger.Error(err)
		closer.Close()
		return nil, nil, err
	}

	return db, closer, nil
}
