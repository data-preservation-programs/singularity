package run

import (
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/data-preservation-programs/singularity/service/dealpusher"
	"github.com/data-preservation-programs/singularity/service/epochutil"
	"github.com/urfave/cli/v2"
)

var DealPusherCmd = &cli.Command{
	Name:  "deal-pusher",
	Usage: "Start a deal pusher that monitors deal schedules and pushes deals to storage providers",
	Flags: []cli.Flag{
		&cli.UintFlag{
			Name:    "deal-attempts",
			Usage:   "Number of times to attempt a deal before giving up",
			Aliases: []string{"d"},
			Value:   3,
		},
		&cli.UintFlag{
			Name:        "max-replication-factor",
			Usage:       "Max number of replicas for each individual PieceCID across all clients and providers",
			Aliases:     []string{"M"},
			DefaultText: "Unlimited",
		},
		&cli.IntFlag{
			Name:  "pdp-batch-size",
			Usage: "Number of roots to include in each PDP add-roots transaction",
			Value: 128,
		},
		&cli.IntFlag{
			Name:  "pdp-max-pieces-per-proofset",
			Usage: "Maximum pieces per proof set before handing off to the storage provider",
			Value: 1024,
		},
		&cli.Uint64Flag{
			Name:  "pdp-confirmation-depth",
			Usage: "Number of block confirmations required for PDP transactions",
			Value: 5,
		},
		&cli.DurationFlag{
			Name:  "pdp-poll-interval",
			Usage: "Polling interval for PDP transaction confirmation checks",
			Value: 30 * time.Second,
		},
		&cli.StringFlag{
			Name:    "eth-rpc",
			Usage:   "Ethereum RPC endpoint for FEVM (required to execute PDP and DDO schedules on-chain)",
			EnvVars: []string{"ETH_RPC_URL"},
		},
		&cli.StringFlag{
			Name:    "ddo-contract",
			Usage:   "DDO Diamond proxy contract address",
			EnvVars: []string{"DDO_CONTRACT_ADDRESS"},
		},
		&cli.StringFlag{
			Name:    "ddo-payments-contract",
			Usage:   "DDO Payments proxy contract address",
			EnvVars: []string{"DDO_PAYMENTS_CONTRACT_ADDRESS"},
		},
		&cli.StringFlag{
			Name:    "ddo-payment-token",
			Usage:   "ERC20 payment token address (e.g. USDFC)",
			EnvVars: []string{"DDO_PAYMENT_TOKEN"},
		},
		&cli.IntFlag{
			Name:  "ddo-batch-size",
			Usage: "Number of pieces per DDO allocation transaction",
			Value: 10,
		},
		&cli.Uint64Flag{
			Name:  "ddo-confirmation-depth",
			Usage: "Number of block confirmations required for DDO transactions",
			Value: 5,
		},
		&cli.DurationFlag{
			Name:  "ddo-poll-interval",
			Usage: "Polling interval for DDO transaction confirmation checks",
			Value: 30 * time.Second,
		},
		&cli.Int64Flag{
			Name:  "ddo-term-min",
			Usage: "Minimum term in epochs (~6 months default)",
			Value: 518400,
		},
		&cli.Int64Flag{
			Name:  "ddo-term-max",
			Usage: "Maximum term in epochs (~5 years default)",
			Value: 5256000,
		},
		&cli.Int64Flag{
			Name:  "ddo-expiration-offset",
			Usage: "Expiration offset in epochs",
			Value: 172800,
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()
		lotusAPI := c.String("lotus-api")
		lotusToken := c.String("lotus-token")
		err = epochutil.Initialize(c.Context, lotusAPI, lotusToken)
		if err != nil {
			return errors.WithStack(err)
		}

		pdpCfg := dealpusher.PDPSchedulingConfig{
			BatchSize:            c.Int("pdp-batch-size"),
			MaxPiecesPerProofSet: c.Int("pdp-max-pieces-per-proofset"),
			ConfirmationDepth:    c.Uint64("pdp-confirmation-depth"),
			PollingInterval:      c.Duration("pdp-poll-interval"),
		}
		if err := pdpCfg.Validate(); err != nil {
			return errors.WithStack(err)
		}

		opts := []dealpusher.Option{
			dealpusher.WithPDPSchedulingConfig(pdpCfg),
		}
		if rpcURL := c.String("eth-rpc"); rpcURL != "" {
			pdpAdapter, err := dealpusher.NewOnChainPDP(c.Context, db, rpcURL)
			if err != nil {
				return errors.Wrap(err, "failed to initialize PDP on-chain adapter")
			}
			defer pdpAdapter.Close()

			opts = append(opts,
				dealpusher.WithPDPProofSetManager(pdpAdapter),
				dealpusher.WithPDPTransactionConfirmer(pdpAdapter),
			)
		}

		if ddoContract := c.String("ddo-contract"); ddoContract != "" {
			ddoCfg := dealpusher.DDOSchedulingConfig{
				BatchSize:         c.Int("ddo-batch-size"),
				ConfirmationDepth: c.Uint64("ddo-confirmation-depth"),
				PollingInterval:   c.Duration("ddo-poll-interval"),
				TermMin:           c.Int64("ddo-term-min"),
				TermMax:           c.Int64("ddo-term-max"),
				ExpirationOffset:  c.Int64("ddo-expiration-offset"),
			}
			opts = append(opts, dealpusher.WithDDOSchedulingConfig(ddoCfg))
			// Path B merge point: instantiate OnChainDDO here and pass via:
			// ddoAdapter, err := dealpusher.NewOnChainDDO(c.Context, db, rpcURL,
			//     ddoContract,
			//     c.String("ddo-payments-contract"),
			//     c.String("ddo-payment-token"),
			//     privateKey,
			// )
			// opts = append(opts,
			//     dealpusher.WithDDODealManager(ddoAdapter),
			//     dealpusher.WithDDOAllocationTracker(ddoAdapter),
			// )
		}

		dm, err := dealpusher.NewDealPusher(
			db,
			c.String("lotus-api"),
			c.String("lotus-token"),
			c.Uint("deal-attempts"),
			c.Uint("max-replication-factor"),
			opts...,
		)
		if err != nil {
			return errors.WithStack(err)
		}
		return service.StartServers(c.Context, dealpusher.Logger, dm)
	},
}
