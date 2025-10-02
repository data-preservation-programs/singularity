package admin

import (
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/urfave/cli/v2"
)

var MigrateCmd = &cli.Command{
	Name:  "migrate",
	Usage: "Migrate database up, down, or to a certain version",

	Before: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()

		// Check if migrations table exists (indicates versioned migration strategy is in place)
		if !db.Migrator().HasTable("migrations") {
			return errors.New("Database has not been initialized with versioned migration strategy. Please run 'singularity admin init' first to migrate your database to the new migration system.")
		}

		return nil
	},
	Subcommands: []*cli.Command{
		{
			Name:  "up",
			Usage: "Execute any unrun migrations",
			Action: func(c *cli.Context) error {
				db, closer, err := database.OpenFromCLI(c)
				if err != nil {
					return errors.WithStack(err)
				}
				defer func() { _ = closer.Close() }()
				return model.GetMigrator(db).Migrate()
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
				defer func() { _ = closer.Close() }()
				return model.GetMigrator(db).RollbackLast()
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
				defer func() { _ = closer.Close() }()

				id := c.Args().Get(0)

				migrator := model.GetMigrator(db)
				last, err := migrator.GetLastMigration()
				if err != nil {
					return errors.WithStack(err)
				}
				if last == id {
					fmt.Println("Already at requested migration")
					return nil
				}

				alreadyRan, err := migrator.HasRunMigration(id)
				if err != nil {
					return errors.WithStack(err)
				} else if alreadyRan {
					return migrator.RollbackTo(id)
				} else {
					return migrator.MigrateTo(id)
				}
			},
		},
		{
			Name:  "which",
			Usage: "Print current migration ID",
			Action: func(c *cli.Context) error {
				db, closer, err := database.OpenFromCLI(c)
				if err != nil {
					return errors.WithStack(err)
				}
				defer func() { _ = closer.Close() }()

				last, err := model.GetMigrator(db).GetLastMigration()
				if err != nil {
					return errors.WithStack(err)
				}
                fmt.Printf("Current migration: %s\n", last)
				return nil
			},
		},
	},
}
