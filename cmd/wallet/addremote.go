package wallet

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/urfave/cli/v2"
)

var AddRemoteCmd = &cli.Command{
	Name:      "add-remote",
	Usage:     "Add remote wallet",
	ArgsUsage: "<address> <remote_peer>",
	Flags:     []cli.Flag{},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		defer closer.Close()

		lotusClient := util.NewLotusClient(c.String("lotus-api"), c.String("lotus-token"))
		w, err2 := wallet.AddRemoteHandler(
			c.Context,
			db,
			lotusClient,
			wallet.AddRemoteRequest{
				Address:    c.Args().Get(0),
				RemotePeer: c.Args().Get(1),
			})
		if err2 != nil {
			return err2
		}

		cliutil.PrintToConsole(w, c.Bool("json"), nil)
		return nil
	},
}
