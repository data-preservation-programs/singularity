package wallet

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/urfave/cli/v2"
)

var RemoveCmd = &cli.Command{
	Name:      "remove",
	Usage:     "Remove a wallet",
	ArgsUsage: "<address>",
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		defer closer.Close()
		return wallet.RemoveHandler(db, c.Args().Get(0))
	},
}
