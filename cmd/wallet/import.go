package wallet

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/urfave/cli/v2"
)

var ImportCmd = &cli.Command{
	Name:      "import",
	Usage:     "Import a wallet from exported private key",
	ArgsUsage: "PRIVATE_KEY",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		w, err2 := wallet.ImportHandler(db, wallet.ImportRequest{
			PrivateKey: c.Args().Get(0),
			LotusAPI:   c.String("lotus-api"),
			LotusToken: c.String("lotus-token"),
		})
		if err2 != nil {
			return err2.CliError()
		}

		cliutil.PrintToConsole(w, c.Bool("json"), nil)
		return nil
	},
}
