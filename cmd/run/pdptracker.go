package run

import (
	"fmt"
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
- Tracks challenge epochs and live status`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "eth-rpc",
			Usage:    "Ethereum RPC endpoint for FEVM (e.g., https://api.node.glif.io)",
			EnvVars:  []string{"ETH_RPC_URL"},
			Required: true,
		},
		&cli.DurationFlag{
			Name:  "pdp-poll-interval",
			Usage: "Polling interval for PDP transaction status",
			Value: 30 * time.Second,
		},
	},
	Action: func(c *cli.Context) error {
		rpcURL := c.String("eth-rpc")
		if rpcURL == "" {
			return fmt.Errorf("eth-rpc is required")
		}

		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		defer closer.Close()

		pdpClient, err := pdptracker.NewPDPClient(c.Context, rpcURL)
		if err != nil {
			return err
		}
		defer pdpClient.Close()

		cfg := pdptracker.PDPConfig{
			PollingInterval: c.Duration("pdp-poll-interval"),
		}
		if err := cfg.Validate(); err != nil {
			return err
		}

		tracker := pdptracker.NewPDPTracker(
			db,
			cfg,
			pdpClient,
			false,
		)

		return service.StartServers(c.Context, pdptracker.Logger, &tracker)
	},
}
