package wallet

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/urfave/cli/v2"
)

var AddRemoteCmd = &cli.Command{
	Name:      "add-remote",
	Usage:     "Add remote wallet",
	ArgsUsage: "<address> <remote_peer>",
	Flags:     []cli.Flag{},
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)

		w, err2 := wallet.AddRemoteHandler(db, wallet.AddRemoteRequest{
			Address:    c.Args().Get(0),
			RemotePeer: c.Args().Get(1),
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
