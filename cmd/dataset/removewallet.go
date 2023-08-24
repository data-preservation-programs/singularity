package dataset

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/urfave/cli/v2"
)

var RemoveWalletCmd = &cli.Command{
	Name:      "remove-wallet",
	Usage:     "Remove an associated wallet from the dataset",
	ArgsUsage: "DATASET_NAME WALLET_ADDRESS",
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()
		return wallet.RemoveWalletHandler(c.Context, db, c.Args().Get(0), c.Args().Get(1))
	},
}
