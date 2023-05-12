package dataset

import (
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/handler/dataset"
	"github.com/urfave/cli/v2"
)

var RemoveWalletCmd = &cli.Command{
	Name:      "remove-wallet",
	Usage:     "Remove an associated wallet from the dataset",
	ArgsUsage: "DATASET_NAME WALLET_ADDRESS",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		err := dataset.RemoveWalletHandler(db, c.Args().Get(0), c.Args().Get(1))
		if err != nil {
			return err.CliError()
		}
		return nil
	},
}
