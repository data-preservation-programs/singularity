package wallet

import (
	"github.com/data-preservation-programs/go-singularity/cmd/cliutil"
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/handler/wallet"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/urfave/cli/v2"
)

var ListCmd = &cli.Command{
	Name:      "list",
	Usage:     "List all imported wallets",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		err := model.InitializeEncryption(c.String("password"), db)
		if err != nil {
			return cli.Exit(err, 1)
		}

		wallets, err2 := wallet.ListHandler(db)
		if err2 != nil {
			return err2.CliError()
		}

		cliutil.PrintToConsole(wallets, c.Bool("json"))
		return nil
	},
}
