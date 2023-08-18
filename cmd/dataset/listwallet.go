package dataset

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/urfave/cli/v2"
)

var ListWalletCmd = &cli.Command{
	Name:      "list-wallet",
	Usage:     "List all associated wallets with the dataset",
	ArgsUsage: "DATASET_NAME",
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()
		wallets, err := wallet.ListWalletHandler(c.Context, db, c.Args().Get(0))
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.PrintToConsole(wallets, c.Bool("json"), nil)
		return nil
	},
}
