package admin

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/admin"
	"github.com/urfave/cli/v2"
)

var ResetCmd = &cli.Command{
	Name:  "reset",
	Usage: "[Dangerous] Reset the database",
	Action: func(context *cli.Context) error {
		db := database.MustOpenFromCLI(context)
		return admin.ResetHandler(db).CliError()
	},
}
