package dataprep

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/urfave/cli/v2"
)

var AttachWalletCmd = &cli.Command{
	Name:      "attach-wallet",
	Usage:     "Attach a wallet to a preparation",
	ArgsUsage: "<preparation id|name> <wallet_id>",
	Category:  "Wallet Management",
	Before:    cliutil.CheckNArgs,
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()
		prep, err := wallet.Default.AttachHandler(c.Context, db, c.Args().Get(0), c.Args().Get(1))
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, prep)
		return nil
	},
}

var ListWalletsCmd = &cli.Command{
	Name:      "list-wallets",
	Usage:     "List attached wallets with a preparation",
	ArgsUsage: "<preparation id|name>",
	Category:  "Wallet Management",
	Before:    cliutil.CheckNArgs,
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()
		prep, err := wallet.Default.ListAttachedHandler(c.Context, db, c.Args().Get(0))
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, prep)
		return nil
	},
}

var DetachWalletCmd = &cli.Command{
	Name:      "detach-wallet",
	Usage:     "Detach a wallet to a preparation",
	ArgsUsage: "<preparation id|name> <wallet_id>",
	Category:  "Wallet Management",
	Before:    cliutil.CheckNArgs,
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()
		prep, err := wallet.Default.DetachHandler(c.Context, db, c.Args().Get(0), c.Args().Get(1))
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, prep)
		return nil
	},
}
