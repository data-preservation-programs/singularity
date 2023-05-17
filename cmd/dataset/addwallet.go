package dataset

import (
	"fmt"
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/handler/dataset"
	"github.com/urfave/cli/v2"
)

var AddWalletCmd = &cli.Command{
	Name:  "add-wallet",
	Usage: "Associate a wallet with the dataset",
	ArgsUsage: "DATASET_NAME WALLET_ADDRESS",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		wallet, err := dataset.AddWalletHandler(db, c.Args().Get(0), c.Args().Get(1))
		if err != nil {
			return err.CliError()
		}
		fmt.Println(wallet)
		return nil
	},
}
