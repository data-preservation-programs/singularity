package wallet

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/urfave/cli/v2"
)

var ImportCmd = &cli.Command{
	Name:      "import",
	Usage:     "Import a wallet from exported private key",
	ArgsUsage: "[path, or stdin if omitted]",
	Flags: []cli.Flag{
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
		defer func() { _ = closer.Close() }()

		var privateKey string
		if c.Args().Len() > 0 {
			privateKeyBytes, err := os.ReadFile(c.Args().Get(0))
			if err != nil {
				return errors.WithStack(err)
			}
			privateKey = string(privateKeyBytes)
		} else {
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("Enter the private key: ")
			if scanner.Scan() {
				privateKey = scanner.Text()
			} else {
				return errors.Wrap(scanner.Err(), "failed to read from stdin")
			}
		}

		lotusClient := util.NewLotusClient(c.String("lotus-api"), c.String("lotus-token"))
		w, err := wallet.Default.ImportHandler(
			c.Context,
			db,
			lotusClient,
			wallet.ImportRequest{
				PrivateKey: privateKey,
				Name:       c.String("name"),
				Contact:    c.String("contact"),
				Location:   c.String("location"),
			})
		if err != nil {
			return errors.WithStack(err)
		}

		cliutil.Print(c, w)
		return nil
	},
}
