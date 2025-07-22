package proof

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/proof"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/urfave/cli/v2"
)

var SyncCmd = &cli.Command{
	Name:  "sync",
	Usage: "Sync proofs from Filecoin blockchain",
	Flags: []cli.Flag{
		&cli.Uint64Flag{
			Name:  "deal-id",
			Usage: "Sync proofs for specific deal ID",
		},
		&cli.StringFlag{
			Name:  "provider",
			Usage: "Sync proofs for specific storage provider",
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()

		lotusClient := util.NewLotusClient(c.String("lotus-api"), c.String("lotus-token"))

		request := proof.SyncProofRequest{}

		if c.IsSet("deal-id") {
			dealID := c.Uint64("deal-id")
			request.DealID = &dealID
		}

		if c.IsSet("provider") {
			provider := c.String("provider")
			request.Provider = &provider
		}

		err = proof.Default.SyncHandler(c.Context, db, lotusClient, request)
		if err != nil {
			return errors.WithStack(err)
		}

		cliutil.Print(c, map[string]string{"status": "success"})
		return nil
	},
}
