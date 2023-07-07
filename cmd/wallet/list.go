package wallet

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/urfave/cli/v2"
)

var ListCmd = &cli.Command{
	Name:  "list",
	Usage: "List all imported wallets",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		wallets, err2 := wallet.ListHandler(db)
		if err2 != nil {
			return err2
		}

		cliutil.PrintToConsole(wallets, c.Bool("json"), nil)
		return nil
	},
}
