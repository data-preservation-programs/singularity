package dealtemplate

import (
	"encoding/json"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/dealtemplate"
	"github.com/urfave/cli/v2"
)

var ListCmd = &cli.Command{
	Name:     "list",
	Usage:    "List all deal templates as pretty-printed JSON",
	Category: "Deal Template Management",
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()
		db = db.WithContext(c.Context)

		templates, err := dealtemplate.Default.ListHandler(c.Context, db)
		if err != nil {
			return errors.WithStack(err)
		}

		if len(templates) == 0 {
			fmt.Println("[]")
			return nil
		}

		jsonBytes, err := json.MarshalIndent(templates, "", "  ")
		if err != nil {
			return errors.Wrap(err, "failed to marshal templates as JSON")
		}
		fmt.Println(string(jsonBytes))
		return nil
	},
}
