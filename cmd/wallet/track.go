package wallet

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/urfave/cli/v2"
)

var TrackCmd = &cli.Command{
	Name:      "track",
	Usage:     "Track a wallet without importing private key",
	ArgsUsage: "<actor_id>",
	Before:    cliutil.CheckNArgs,
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()

		actorID := c.Args().Get(0)
		lotusClient := util.NewLotusClient(c.String("lotus-api"), c.String("lotus-token"))

		request := wallet.CreateRequest{
			ActorID:   actorID,
			TrackOnly: true,
		}

		w, err := wallet.Default.CreateHandler(c.Context, db, lotusClient, request)
		if err != nil {
			return errors.WithStack(err)
		}

		cliutil.Print(c, w)
		return nil
	},
}
