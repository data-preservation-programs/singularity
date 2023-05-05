package handler

import (
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

func InitHandler(db *gorm.DB) error {
	err := model.AutoMigrate(db)
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	return nil
}
