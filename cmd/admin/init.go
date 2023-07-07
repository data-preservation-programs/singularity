package admin

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/admin"
	"github.com/urfave/cli/v2"
)

var InitCmd = &cli.Command{
	Name:  "init",
	Usage: "Initialize the database",
	Action: func(context *cli.Context) error {
		db, err := database.OpenFromCLI(context)
		if err != nil {
			return err
		}
		return admin.InitHandler(db)
	},
}
