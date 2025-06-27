package wallet

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/gotidy/ptr"
	"github.com/urfave/cli/v2"
)

var UpdateCmd = &cli.Command{
	Name:      "update",
	Usage:     "Update wallet details",
	ArgsUsage: "<address>",
	Description: `Update non-essential details of an existing wallet.

This command allows you to update the following wallet properties:
- Name (optional wallet label)
- Contact information (email for SP)
- Location (region, country for SP)

Essential properties like the wallet address, private key, and balance cannot be modified.

EXAMPLES:
		# Update the actor name
		singularity wallet update f1abc123... --name "My Main Wallet"

		# Update multiple fields at once
		singularity wallet update f1xyz789... --name "Storage Provider" --location "US-East"`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "name",
			Usage: "Set the readable label for the wallet",
		},
		&cli.StringFlag{
			Name:  "contact",
			Usage: "Set the contact information (email) for the wallet",
		},
		&cli.StringFlag{
			Name:  "location",
			Usage: "Set the location (region, country) for the wallet",
		},
	},
	Before: cliutil.CheckNArgs,
	Action: func(c *cli.Context) error {
		address := c.Args().Get(0)

		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()

		// Build the update request
		request := wallet.UpdateRequest{}

		if c.IsSet("name") {
			request.Name = ptr.Of(c.String("name"))
		}

		if c.IsSet("contact") {
			request.Contact = ptr.Of(c.String("contact"))
		}

		if c.IsSet("location") {
			request.Location = ptr.Of(c.String("location"))
		}

		// Check if at least one field is provided for update
		if request.Name == nil && request.Contact == nil && request.Location == nil {
			return errors.New("at least one field must be provided for update")
		}

		w, err := wallet.Default.UpdateHandler(
			c.Context,
			db,
			address,
			request,
		)
		if err != nil {
			return errors.WithStack(err)
		}

		cliutil.Print(c, w)
		return nil
	},
}
