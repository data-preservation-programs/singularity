package dataset

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/urfave/cli/v2"
)

var AddWalletCmd = &cli.Command{
	Name:      "add-wallet",
	Usage:     "Associate a wallet with the dataset. The wallet needs to be imported first using the `singularity wallet import` command.",
	ArgsUsage: "DATASET_NAME WALLET_ADDRESS",
	Action: func(c *cli.Context) error {
		db, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		wallet, err := wallet.AddWalletHandler(db, c.Args().Get(0), c.Args().Get(1))
		if err != nil {
			return err
		}
		cliutil.PrintToConsole(wallet, c.Bool("json"), nil)
		return nil
	},
}
