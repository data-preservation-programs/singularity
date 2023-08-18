package wallet

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/urfave/cli/v2"
)

var ImportCmd = &cli.Command{
	Name:      "import",
	Usage:     "Import a wallet from exported private key",
	ArgsUsage: "PRIVATE_KEY",
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		lotusClient := util.NewLotusClient(c.String("lotus-api"), c.String("lotus-token"))
		w, err := wallet.ImportHandler(
			c.Context,
			db,
			lotusClient,
			wallet.ImportRequest{
				PrivateKey: c.Args().Get(0),
			})
		if err != nil {
			return errors.WithStack(err)
		}

		cliutil.PrintToConsole(w, c.Bool("json"), nil)
		return nil
	},
}
