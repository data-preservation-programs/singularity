package sppool

import (
	"encoding/json"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/deal/sppool"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/urfave/cli/v2"
)

var AddProviderCmd = &cli.Command{
	Name:      "add-provider",
	Usage:     "Add a storage provider to an SP Pool",
	Before:    cliutil.CheckNArgs,
	ArgsUsage: "<pool_id>",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "provider",
			Usage:    "Storage Provider ID",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "policy",
			Usage:    `Replication policy as JSON, e.g. '{"market": 1, "pdp": 1}'`,
			Required: true,
		},
		&cli.StringFlag{
			Name:  "url-template",
			Usage: "Optional per-provider URL template override",
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		poolID, err := strconv.ParseUint(c.Args().Get(0), 10, 32)
		if err != nil {
			return errors.Wrapf(err, "failed to parse pool ID %s", c.Args().Get(0))
		}

		var policy model.ReplicationPolicy
		if err := json.Unmarshal([]byte(c.String("policy")), &policy); err != nil {
			return errors.Wrap(err, "failed to parse policy JSON")
		}

		request := sppool.AddProviderRequest{
			Provider: c.String("provider"),
			Policy:   policy,
		}
		if c.IsSet("url-template") {
			v := c.String("url-template")
			request.URLTemplate = &v
		}

		lotusClient := util.NewLotusClient(c.String("lotus-api"), c.String("lotus-token"))
		provider, err := sppool.Default.AddProviderHandler(c.Context, db, lotusClient, uint32(poolID), request)
		if err != nil {
			return errors.WithStack(err)
		}
		cliutil.Print(c, provider)
		return nil
	},
}
