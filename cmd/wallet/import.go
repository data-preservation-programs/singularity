package wallet

import (
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/handler/wallet"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/urfave/cli/v2"
)

var ImportCmd = &cli.Command{
	Name:      "import",
	Usage:     "Import a wallet from exported private key",
	ArgsUsage: "PRIVATE_KEY",
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		err := model.InitializeEncryption(c.String("password"), db)
		if err != nil {
			return cli.Exit(err, 1)
		}

		err2 := wallet.ImportHandler(db, wallet.ImportRequest{
			PrivateKey: c.Args().Get(0),
		})
		if err2 != nil {
			return err2.CliError()
		}

		return nil
	},
}
