package dealtemplate

import (
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dealtemplate"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

var GetCmd = &cli.Command{
	Name:      "get",
	Usage:     "Get a deal template by ID or name",
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

		templateIdentifier := c.Args().First()
		template, err := dealtemplate.Default.GetHandler(c.Context, db, templateIdentifier)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.Errorf("Template \"%s\" not found", templateIdentifier)
			}
			return errors.WithStack(err)
		}

		// Print context before template data
		if !c.Bool("json") {
			fmt.Printf("→ Deal Template: %s (ID: %d)\n", template.Name, template.ID)
		}

		cliutil.Print(c, *template)
		return nil
	},
}
