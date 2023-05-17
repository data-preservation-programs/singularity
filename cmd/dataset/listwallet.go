package dataset

import (
	"github.com/data-preservation-programs/go-singularity/cmd/cliutil"
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/handler/dataset"
	"github.com/urfave/cli/v2"
)

var ListWalletCmd = &cli.Command{
	Name:  "list-wallet",
	Usage: "List all associated wallets with the dataset",
	ArgsUsage: "DATASET_NAME",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		wallets, err := dataset.ListWalletHandler(db, c.Args().Get(0))
		if err != nil {
			return err.CliError()
		}
		cliutil.PrintToConsole(wallets, c.Bool("json"))
		return nil
	},
}
