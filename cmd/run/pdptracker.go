package run

import (
	"time"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/data-preservation-programs/singularity/service/pdptracker"
	"github.com/urfave/cli/v2"
)

var PDPTrackerCmd = &cli.Command{
	Name:  "pdp-tracker",
	Usage: "Start a PDP deal tracker that tracks f41 PDP deals for all relevant wallets",
	Description: `The PDP tracker monitors Proof of Data Possession (PDP) deals on the Filecoin network.
Unlike legacy f05 market deals, PDP deals use proof sets managed through the PDPVerifier contract
where data is verified through cryptographic challenges.

This tracker:
- Monitors proof sets for tracked wallets
- Updates deal status based on on-chain proof set state
- Tracks challenge epochs and live status

Note: Full functionality requires the go-synapse library integration.
See: https://github.com/data-preservation-programs/go-synapse`,
	Flags: []cli.Flag{
		&cli.DurationFlag{
			Name:  "interval",
			Usage: "How often to check for PDP deal updates",
			Value: 10 * time.Minute,
		},
		&cli.StringFlag{
			Name:    "lotus-api",
			Usage:   "Lotus RPC API endpoint",
			EnvVars: []string{"LOTUS_API"},
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		defer closer.Close()

		pdpClient, err := pdptracker.NewPDPClient(c.Context, c.String("lotus-api"))
		if err != nil {
			return err
		}
		defer pdpClient.Close()

		tracker := pdptracker.NewPDPTracker(
			db,
			c.Duration("interval"),
			c.String("lotus-api"),
			pdpClient,
			false,
		)

		return service.StartServers(c.Context, pdptracker.Logger, &tracker)
	},
}
