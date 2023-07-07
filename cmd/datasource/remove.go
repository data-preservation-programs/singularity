package datasource

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/urfave/cli/v2"
)

var RemoveCmd = &cli.Command{
	Name:      "remove",
	Usage:     "Remove a data source",
	ArgsUsage: "<source_id>",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		err := datasource.RemoveSourceHandler(
			db,
			c.Args().Get(0),
		)
		return err
	},
}
