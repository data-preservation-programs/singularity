package dealtemplate

import (
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dealtemplate"
	"github.com/urfave/cli/v2"
)

var ListCmd = &cli.Command{
	Name:     "list",
	Usage:    "List all deal templates",
	Category: "Deal Template Management",
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()
		db = db.WithContext(c.Context)

		templates, err := dealtemplate.Default.ListHandler(c.Context, db)
		if err != nil {
			return errors.WithStack(err)
		}

		// Handle empty results
		if len(templates) == 0 {
			if !c.Bool("json") {
				fmt.Println("No deal templates found.")
				return nil
			}
		} else {
			// Print summary for non-JSON output
			if !c.Bool("json") {
				fmt.Printf("✓ %d deal template(s) found.\n\n", len(templates))
			}
		}

		cliutil.Print(c, templates)
		return nil
	},
}
