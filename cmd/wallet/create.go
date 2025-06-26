package wallet

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/urfave/cli/v2"
)

var CreateCmd = &cli.Command{
	Name:      "create",
	Usage:     "Create a new wallet",
	ArgsUsage: "[type]",
	Description: `Create a new Filecoin wallet or storage provider contact entry.

The command automatically detects the wallet type based on provided arguments:
- For UserWallet: Creates a wallet with offline keypair generation
- For SPWallet: Creates a contact entry for a storage provider

SUPPORTED KEY TYPES (for UserWallet):
  secp256k1    ECDSA using the secp256k1 curve (default, most common)
  bls          BLS signature scheme (Boneh-Lynn-Shacham)

EXAMPLES:
  # Create a secp256k1 UserWallet (default)
  singularity wallet create

  # Create a secp256k1 UserWallet explicitly
  singularity wallet create secp256k1

  # Create a BLS UserWallet
  singularity wallet create bls

  # Create an SPWallet contact entry
  singularity wallet create --address f3abc123... --actor-id f01234 --name "Example SP"

The newly created wallet address and other details will be displayed upon successful creation.`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "address",
			Usage: "Storage provider wallet address (creates SPWallet contact)",
		},
		&cli.StringFlag{
			Name:  "actor-id",
			Usage: "Storage provider actor ID (e.g., f01234)",
		},
		&cli.StringFlag{
			Name:  "name",
			Usage: "Optional display name",
		},
		&cli.StringFlag{
			Name:  "contact",
			Usage: "Optional contact information",
		},
		&cli.StringFlag{
			Name:  "location",
			Usage: "Optional provider location",
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		request := wallet.CreateRequest{
			Name:     c.String("name"),
			Contact:  c.String("contact"),
			Location: c.String("location"),
		}

		// Check if this is an SPWallet creation (has address or actor-id)
		address := c.String("address")
		actorID := c.String("actor-id")

		if address != "" || actorID != "" {
			// Create SPWallet contact entry
			request.Address = address
			request.ActorID = actorID
		} else {
			// Create UserWallet with keypair generation
			keyType := c.Args().Get(0)
			if keyType == "" {
				keyType = wallet.KTSecp256k1.String()
			}
			request.KeyType = keyType
		}

		lotusClient := util.NewLotusClient(c.String("lotus-api"), c.String("lotus-token"))

		w, err := wallet.Default.CreateHandler(c.Context, db, lotusClient, request)
		if err != nil {
			return errors.WithStack(err)
		}

		cliutil.Print(c, w)
		return nil
	},
}
