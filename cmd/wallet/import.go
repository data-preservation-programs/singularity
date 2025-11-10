package wallet

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/util/keystore"
	"github.com/urfave/cli/v2"
)

var ImportCmd = &cli.Command{
	Name:      "import",
	Usage:     "Import a wallet from a private key file into the keystore",
	ArgsUsage: "[path, or stdin if omitted]",
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

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

		ks, err := keystore.NewLocalKeyStore(wallet.GetKeystoreDir())
		if err != nil {
			return errors.Wrap(err, "failed to init keystore")
		}

		w, err := wallet.Default.ImportKeystoreHandler(
			c.Context,
			db,
			ks,
			wallet.ImportKeystoreRequest{
				PrivateKey: privateKey,
			})
		if err != nil {
			return errors.WithStack(err)
		}

		cliutil.Print(c, w)
		return nil
	},
}
