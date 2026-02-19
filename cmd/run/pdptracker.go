package run

import (
	"fmt"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/go-synapse"
	"github.com/data-preservation-programs/go-synapse/constants"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/data-preservation-programs/singularity/service/pdptracker"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v2"
)

var PDPTrackerCmd = &cli.Command{
	Name:  "pdp-tracker",
	Usage: "Start a PDP deal tracker that tracks f41 PDP deals using Shovel event indexing",
	Description: `The PDP tracker uses an embedded Shovel instance to index PDPVerifier contract
events into Postgres, then processes the events to maintain deal state.

This replaces linear chain state scanning with cursor-based event indexing,
eliminating the O(N) RPC overhead that scales with the number of proof sets.

Requires PostgreSQL (Shovel is Postgres-only).`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "eth-rpc",
			Usage:    "Ethereum RPC endpoint for FEVM (e.g., https://api.node.glif.io)",
			EnvVars:  []string{"ETH_RPC_URL"},
			Required: true,
		},
		&cli.DurationFlag{
			Name:  "pdp-poll-interval",
			Usage: "How often to check for new events in Shovel tables",
			Value: 30 * time.Second,
		},
	},
	Action: func(c *cli.Context) error {
		rpcURL := c.String("eth-rpc")
		if rpcURL == "" {
			return fmt.Errorf("eth-rpc is required")
		}

		connStr := c.String("database-connection-string")
		if !strings.HasPrefix(connStr, "postgres:") {
			return errors.New("PDP tracking requires PostgreSQL (Shovel is Postgres-only)")
		}

		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		// detect network and contract address once, shared by indexer and rpc client
		ethClient, err := ethclient.DialContext(c.Context, rpcURL)
		if err != nil {
			return errors.Wrap(err, "failed to connect to RPC")
		}
		network, chainID, err := synapse.DetectNetwork(c.Context, ethClient)
		ethClient.Close()
		if err != nil {
			return errors.Wrap(err, "failed to detect network")
		}

		contractAddr := constants.GetPDPVerifierAddress(network)
		if contractAddr == (common.Address{}) {
			return fmt.Errorf("no PDPVerifier contract for network %s", network)
		}

		pdptracker.Logger.Infow("detected PDP network",
			"network", network,
			"chainId", chainID,
			"contract", contractAddr.Hex(),
		)

		// start shovel indexer
		indexer, err := pdptracker.NewPDPIndexer(c.Context, connStr, rpcURL, uint64(chainID), contractAddr)
		if err != nil {
			return errors.Wrap(err, "failed to create PDP indexer")
		}

		// create rpc client for remaining on-chain calls
		rpcClient, err := pdptracker.NewPDPClient(c.Context, rpcURL, contractAddr)
		if err != nil {
			return errors.Wrap(err, "failed to create PDP RPC client")
		}
		defer rpcClient.Close()

		cfg := pdptracker.PDPConfig{
			PollingInterval: c.Duration("pdp-poll-interval"),
		}
		if err := cfg.Validate(); err != nil {
			return err
		}

		tracker := pdptracker.NewPDPTracker(db, cfg, rpcClient, false)

		return service.StartServers(c.Context, pdptracker.Logger, indexer, &tracker)
	},
}
