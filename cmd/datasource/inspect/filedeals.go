package inspect

import (
	"strconv"

	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/urfave/cli/v2"
)

var FileDealsCmd = &cli.Command{
	Name:      "filedeals",
	Usage:     "Get all that have been created for a file",
	ArgsUsage: "<file_id>",
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		defer closer.Close()

		fileID, err := strconv.ParseUint(c.Args().Get(0), 10, 64)
		if err != nil {
			return err
		}

		deals, err := inspect.GetFileDealsHandler(db, uint64(fileID))
		if err != nil {
			return err
		}

		cliutil.PrintToConsole(deals, c.Bool("json"), nil)

		return nil
	},
}
