package admin

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/admin"
	"github.com/urfave/cli/v2"
)

var InitCmd = &cli.Command{
	Name:        "init",
	Usage:       "Initialize or upgrade the database",
	Description: "This commands need to be run before running any singularity daemon or after any version upgrade",
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()
		return admin.Default.InitHandler(c.Context, db)
	},
}
