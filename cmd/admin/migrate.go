package admin

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/urfave/cli/v2"
)

var MigrateCmd = &cli.Command{
	Name:  "migrate",
	Usage: "Migrate database up, down, or to a certain version",
	Subcommands: []*cli.Command{
		{
			Name:  "up",
			Usage: "Execute any unrun migrations",
			Action: func(c *cli.Context) error {
				db, closer, err := database.OpenFromCLI(c)
				if err != nil {
					return errors.WithStack(err)
				}
				defer closer.Close()
				return model.Migrator(db).Migrate()
			},
		},
		{
			Name:  "down",
			Usage: "Rollback to previous migration",
			Action: func(c *cli.Context) error {
				db, closer, err := database.OpenFromCLI(c)
				if err != nil {
					return errors.WithStack(err)
				}
				defer closer.Close()
				return model.Migrator(db).RollbackLast()
			},
		},
		{
			Name:      "to",
			Usage:     "Migrate to specified version",
			ArgsUsage: "<migration id>",
			Before:    cliutil.CheckNArgs,
			Action: func(c *cli.Context) error {
				db, closer, err := database.OpenFromCLI(c)
				if err != nil {
					return errors.WithStack(err)
				}
				defer closer.Close()
				return model.Migrator(db).MigrateTo(c.Args().Get(0))
			},
		},
	},
}
