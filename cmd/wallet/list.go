package wallet

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/urfave/cli/v2"
)

var ListCmd = &cli.Command{
	Name:  "list",
	Usage: "List all imported wallets",
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()
		wallets, err := wallet.ListHandler(c.Context, db)
		if err != nil {
			return errors.WithStack(err)
		}

		cliutil.PrintToConsole(wallets, c.Bool("json"), nil)
		return nil
	},
}
