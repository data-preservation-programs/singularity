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
	Action: func(c *cli.Context) error {
		if c.NArg() != 1 {
			return errors.New("template ID or name is required")
		}

		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()
		db = db.WithContext(c.Context)

		err = dealtemplate.Default.DeleteHandler(c.Context, db, c.Args().First())
		if err != nil {
			return errors.WithStack(err)
		}

		return nil
	},
}
