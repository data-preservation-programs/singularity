package wallet

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/urfave/cli/v2"
)

var CreateCmd = &cli.Command{
	Name:      "create",
	Usage:     "Create a new wallet",
	ArgsUsage: "[type]",
	Description: `Create a new Filecoin wallet using offline keypair generation.

The wallet will be stored locally in the Singularity database and can be used for making deals and other operations. The private key is generated securely and stored encrypted.

SUPPORTED KEY TYPES:
  secp256k1    ECDSA using the secp256k1 curve (default, most common)
  bls          BLS signature scheme (Boneh-Lynn-Shacham)

EXAMPLES:
  # Create a secp256k1 wallet (default)
  singularity wallet create

  # Create a secp256k1 wallet explicitly
  singularity wallet create secp256k1

  # Create a BLS wallet
  singularity wallet create bls

The newly created wallet address and other details will be displayed upon successful creation.`,
	Before: cliutil.CheckNArgs,
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		// Default to secp256k1 if no type is provided
		keyType := c.Args().Get(0)
		if keyType == "" {
			keyType = wallet.KTSecp256k1.String()
		}

		w, err := wallet.Default.CreateHandler(
			c.Context,
			db,
			wallet.CreateRequest{
				KeyType: keyType,
			})
		if err != nil {
			return errors.WithStack(err)
		}

		cliutil.Print(c, w)
		return nil
	},
}
