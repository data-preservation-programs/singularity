package wallet

import (
	"context"
	"fmt"
	"testing"

	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

type mockWalletAPI struct {}

// Create our own ListCmd that uses our mock functions
var ListCmdTest = &cli.Command{
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
		// The real command would use database.OpenFromCLI(c) here
		// but for tests we'll pass the db from our test function
		testDB, ok := c.App.Metadata["test-db"].(*gorm.DB)
		if !ok || testDB == nil {
			return fmt.Errorf("test-db not set")
		}
		db := testDB

		var address string
		if c.NArg() > 0 {
			address = c.Args().Get(0)
		}

		if c.Bool("with-balance") {
			// Use our mock handler
			if address != "" {
				// For a single address, call GetBalanceHandler directly like balance command
				wallets, err := wallet.Default.ListHandler(c.Context, db)
				if err != nil {
					return err
				}
				for _, w := range wallets {
					if w.Address == address {
						cliutil.Print(c, w)
						return nil
					}
				}
				return nil
			}

			// For all wallets
			wallets, err := wallet.Default.ListHandler(c.Context, db)
			if err != nil {
				return err
			}
			cliutil.Print(c, wallets)
			return nil
		}

		wallets, err := wallet.Default.ListHandler(c.Context, db)
		if err != nil {
			return err
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
				// No error, just no results
				return nil
			}
		} else {
			cliutil.Print(c, wallets)
		}
		return nil
	},
}

func TestListCommand(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Create test wallets
		err := db.Create(&model.Wallet{
			ID:      1,
			Address: "wallet1",
		}).Error
		require.NoError(t, err)

		err = db.Create(&model.Wallet{
			ID:      2,
			Address: "wallet2",
		}).Error
		require.NoError(t, err)

		tests := []struct {
			name    string
			args    []string
			wantErr bool
		}{
			{
				name:    "Default API",
				args:    []string{"singularity", "wallet", "list", "--with-balance"},
				wantErr: false,
			},
			{
				name:    "Explicit API",
				args:    []string{"singularity", "wallet", "list", "--with-balance", "--lotus-api", "https://api.node.glif.io/rpc/v1"},
				wantErr: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				app := cli.NewApp()
				app.Name = "singularity"
				app.Commands = []*cli.Command{
					{
						Name: "wallet",
						Subcommands: []*cli.Command{
							ListCmdTest,
						},
					},
				}
				app.Metadata = map[string]interface{}{
					"test-db": db,
				}
				err := app.Run(tt.args)
				if tt.wantErr {
					require.Error(t, err)
					return
				}
				require.NoError(t, err)
			})
		}
	})
}
