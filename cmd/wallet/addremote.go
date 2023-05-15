package wallet

import (
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/handler/wallet"
	"github.com/urfave/cli/v2"
)

var AddRemoteCmd = &cli.Command{
	Name:      "add-remote",
	Usage:     "Add remote wallet",
	ArgsUsage: "WALLET_ADDRESS REMOTE_PEER",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "lotus-api",
			Category: "Lotus",
			Usage:    "Lotus RPC API endpoint",
			Value:    "https://api.node.glif.io/rpc/v1",
		},
		&cli.StringFlag{
			Name:     "lotus-token",
			Category: "Lotus",
			Usage:    "Lotus RPC API token",
			Value:    "",
		},
	},
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)

		err2 := wallet.AddRemoteHandler(db, wallet.AddRemoteRequest{
			WalletAddress: c.Args().Get(0),
			RemotePeer: c.Args().Get(1),
			LotusAPI:   c.String("lotus-api"),
			LotusToken: c.String("lotus-token"),
		})
		if err2 != nil {
			return err2.CliError()
		}

		return nil
	},
}
