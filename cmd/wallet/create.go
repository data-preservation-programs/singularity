package wallet

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/urfave/cli/v2"
)

var CreateCmd = &cli.Command{
	Name:      "create",
	Usage:     "Create a new wallet",
	ArgsUsage: "[type]",
	Before:    cliutil.CheckNArgs,
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		// Default to secp256k1 if no type is provided
		keyType := c.Args().Get(0)
		if keyType == "" {
			keyType = wallet.KTSecp256k1.String()
		}

		w, err := wallet.Default.CreateHandler(
			c.Context,
			db,
			wallet.CreateRequest{
				KeyType: keyType,
			})
		if err != nil {
			return errors.WithStack(err)
		}

		cliutil.Print(c, w)
		return nil
	},
}
