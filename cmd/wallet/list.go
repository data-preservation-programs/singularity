package wallet

import (
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/urfave/cli/v2"
)

var ListCmd = &cli.Command{
	Name:      "list",
	Usage:     "List all imported wallets or a specific wallet if address is provided",
	ArgsUsage: "[wallet_address]",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "with-balance",
			Usage: "Fetch and display live wallet balances from Lotus",
			Value: false,
		},
		&cli.StringFlag{
			Name:    "lotus-api",
			Usage:   "Lotus JSON-RPC API endpoint for fetching live balances",
			EnvVars: []string{"LOTUS_API"},
		},
		&cli.StringFlag{
			Name:    "lotus-token",
			Usage:   "Lotus API authorization token",
			EnvVars: []string{"LOTUS_TOKEN"},
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()

		var address string
		if c.NArg() > 0 {
			address = c.Args().Get(0)
		}

		if c.Bool("with-balance") {
			// If no Lotus API is provided, use the default from app config
			lotusAPI := c.String("lotus-api")
			if lotusAPI == "" {
				lotusAPI = "https://api.node.glif.io/rpc/v1"
			}
			lotusClient := util.NewLotusClient(lotusAPI, c.String("lotus-token"))

			if address != "" {
				// For a single address, call GetBalanceHandler directly like balance command
				resp, err := wallet.Default.GetBalanceHandler(c.Context, db, lotusClient, address)
				if err != nil {
					return errors.WithStack(err)
				}
				cliutil.Print(c, resp)
				return nil
			}

			// For all wallets, convert WalletWithBalance to BalanceResponse for consistent output
			walletsWithBalance, err := wallet.ListWithBalanceHandler(c.Context, db, lotusClient)
			if err != nil {
				return errors.WithStack(err)
			}
			responses := make([]*wallet.BalanceResponse, len(walletsWithBalance))
			for i, w := range walletsWithBalance {
				responses[i] = &wallet.BalanceResponse{
					Address:        w.Address,
					Balance:        w.Balance,
					BalanceAttoFIL: w.BalanceAttoFIL,
					DataCap:        w.DataCap,
					DataCapBytes:   w.DataCapBytes,
					Error:          w.Error,
				}
			}
			cliutil.Print(c, responses)
			return nil
		}

		wallets, err := wallet.Default.ListHandler(c.Context, db)
		if err != nil {
			return errors.WithStack(err)
		}
		if address != "" {
			found := false
			for _, w := range wallets {
				if w.Address == address {
					cliutil.Print(c, []model.Wallet{w})
					found = true
					break
				}
			}
			if !found {
				fmt.Fprintln(c.App.Writer, "No wallet found with the specified address.")
			}
		} else {
			cliutil.Print(c, wallets)
		}
		return nil
	},
}


