package main

import (
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/urfave/cli/v2"
)

var InitCmd = &cli.Command{
	Name:  "init",
	Usage: "Initialize the database",
	Action: func(context *cli.Context) error {
		db := database.MustOpenFromCLI(context)
		err := model.AutoMigrate(db)
		if err != nil {
			return cli.Exit(err.Error(), 1)
		}

		return nil
	},
}
