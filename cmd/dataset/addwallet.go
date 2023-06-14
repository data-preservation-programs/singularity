package dataset

import (
	"fmt"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/urfave/cli/v2"
)

var AddWalletCmd = &cli.Command{
	Name:      "add-wallet",
	Usage:     "[alpha] Associate a wallet with the dataset. The wallet needs to be imported first using the `singularity wallet import` command.",
	ArgsUsage: "DATASET_NAME WALLET_ADDRESS",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		wallet, err := wallet.AddWalletHandler(db, c.Args().Get(0), c.Args().Get(1))
		if err != nil {
			return err.CliError()
		}
		fmt.Println(wallet)
		return nil
	},
}
