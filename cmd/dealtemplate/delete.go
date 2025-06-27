package dealtemplate

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dealtemplate"
	"github.com/urfave/cli/v2"
)

var DeleteCmd = &cli.Command{
	Name:      "delete",
	Usage:     "Delete a deal template by ID or name",
	Category:  "Deal Template Management",
	ArgsUsage: "<template_id_or_name>",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "force",
			Usage: "Force deletion without confirmation",
		},
	},
	Action: func(c *cli.Context) error {
		if c.NArg() != 1 {
			return errors.New("template ID or name is required")
		}

		templateIdentifier := c.Args().First()

		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()
		db = db.WithContext(c.Context)

		err = dealtemplate.Default.DeleteHandler(c.Context, db, templateIdentifier)
		if err != nil {
			return errors.WithStack(err)
		}

		// Print success confirmation
		println("✓ Deal template \"" + templateIdentifier + "\" deleted successfully")
		return nil
	},
}
