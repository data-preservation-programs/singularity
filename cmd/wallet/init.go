package wallet

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/urfave/cli/v2"
)

var InitCmd = &cli.Command{
	Name:      "init",
	Usage:     "Initialize a wallet",
	ArgsUsage: "<address>",
	Before:    cliutil.CheckNArgs,
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()

		lotusClient := util.NewLotusClient(c.String("lotus-api"), c.String("lotus-token"))
		w, err := wallet.Default.InitHandler(c.Context, db, lotusClient, c.Args().Get(0))
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, w)
		return nil
	},
}
