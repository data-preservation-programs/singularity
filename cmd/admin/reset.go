package admin

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/admin"
	"github.com/urfave/cli/v2"
)

var ResetCmd = &cli.Command{
	Name:  "reset",
	Usage: "Reset the database",
	Flags: []cli.Flag{cliutil.ReallyDotItFlag},
	Action: func(c *cli.Context) error {
		if err := cliutil.HandleReallyDoIt(c); err != nil {
			return errors.WithStack(err)
		}
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()
		return admin.Default.ResetHandler(c.Context, db)
	},
}
