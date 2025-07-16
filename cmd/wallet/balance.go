package wallet

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/urfave/cli/v2"
)

var BalanceCmd = &cli.Command{
	Name:      "balance",
	Usage:     "Get wallet balance information",
	ArgsUsage: "<wallet_address>",
Description: `Get FIL balance and FIL+ datacap balance for a specific wallet address.
This command queries the Lotus network to retrieve current balance information.

Examples:
  singularity wallet balance f1abcde7zd3lfsv43aj2kb454ymaqw7debhumjxyz
  singularity wallet balance --json f1abcde7zd3lfsv43aj2kb454ymaqw7debhumjxyz

The command returns:
- FIL balance in human-readable format (e.g., "1.000000 FIL")
- Raw balance in attoFIL for precise calculations
- FIL+ datacap balance in GiB format (e.g., "1024.50 GiB") 
- Raw datacap in bytes

If there are issues retrieving either balance, partial results will be shown with error details.`,
	Before: cliutil.CheckNArgs,
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()

		lotusClient := util.NewLotusClient(c.String("lotus-api"), c.String("lotus-token"))
		
		result, err := wallet.Default.GetBalanceHandler(c.Context, db, lotusClient, c.Args().Get(0))
		if err != nil {
			return errors.WithStack(err)
		}

		cliutil.Print(c, result)
		return nil
	},
}
