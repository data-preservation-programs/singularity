package wallet

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/urfave/cli/v2"
)

var RemoveCmd = &cli.Command{
	Name:      "remove",
	Usage:     "Remove a wallet",
	ArgsUsage: "<address>",
	Before:    cliutil.CheckNArgs,
	Flags: []cli.Flag{
		cliutil.ReallyDotItFlag,
	},
	Action: func(c *cli.Context) error {
		if err := cliutil.HandleReallyDoIt(c); err != nil {
			return errors.WithStack(err)
		}
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()
		return wallet.Default.RemoveHandler(c.Context, db, c.Args().Get(0))
	},
}
