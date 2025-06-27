package admin

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/admin"
	"github.com/urfave/cli/v2"
)

var InitCmd = &cli.Command{
	Name:  "init",
	Usage: "Initialize or upgrade the database",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "identity",
			Usage: "Name of the user or service that is running the Singularity for tracking and logging purpose",
		},
	},
	Description: "This command needs to be run before running any singularity daemon or after any version upgrade",
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()
		err = admin.Default.InitHandler(c.Context, db)
		if err != nil {
			return errors.WithStack(err)
		}
		if c.IsSet("identity") {
			err = admin.Default.SetIdentityHandler(c.Context, db, admin.SetIdentityRequest{
				Identity: c.String("identity"),
			})
			if err != nil {
				return errors.WithStack(err)
			}
		}
		return nil
	},
}
