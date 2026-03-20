package wallet

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/util/keystore"
	"github.com/urfave/cli/v2"
)

var ExportKeysCmd = &cli.Command{
	Name:  "export-keys",
	Usage: "Migrate private keys from database (legacy Actor.PrivateKey) to the filesystem keystore",
	Description: `Reads private keys stored in the legacy actors table and saves them to
the filesystem keystore (~/.singularity/keystore or SINGULARITY_KEYSTORE).
Creates Wallet records for each exported key and links them to the
corresponding Actor.

This command is idempotent — actors whose address already has a Wallet
record are skipped. Keys that fail to parse are reported but do not
abort the migration.

After exporting, prompts to drop the orphaned private_key column from
the actors table. This is irreversible — verify keys are in the keystore
before confirming. For scripted use, pass --drop-db-keys --i-am-really-sure
to skip the prompt.`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "drop-db-keys",
			Usage: "drop the private_key column from the actors table after export",
		},
		&cli.BoolFlag{
			Name:  "i-am-really-sure",
			Usage: "confirm column drop (required with --drop-db-keys)",
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		// check if column still exists -- if not, nothing to do
		if !wallet.HasPrivateKeyColumn(db) {
			fmt.Println("nothing to do -- private_key column already dropped")
			return nil
		}

		keystoreDir := wallet.GetKeystoreDir()
		ks, err := keystore.NewLocalKeyStore(keystoreDir)
		if err != nil {
			return errors.Wrap(err, "failed to init keystore")
		}

		result, err := wallet.ExportKeysHandler(c.Context, db, ks)
		if err != nil {
			return errors.WithStack(err)
		}

		fmt.Printf("exported: %d\n", result.Exported)
		if result.Skipped > 0 {
			fmt.Printf("skipped:  %d (wallet already exists)\n", result.Skipped)
		}
		if len(result.Errors) > 0 {
			fmt.Printf("errors:   %d\n", len(result.Errors))
			for _, e := range result.Errors {
				fmt.Printf("  - %s\n", e)
			}
			fmt.Println("\nfix the errors above and re-run before dropping the column")
			return nil
		}

		// list keys in keystore so user can verify
		keys, err := ks.List()
		if err != nil {
			return errors.Wrap(err, "failed to list keystore")
		}
		fmt.Printf("\nkeystore: %s (%d keys)\n", keystoreDir, len(keys))

		// determine whether to drop the column
		dropFlag := c.Bool("drop-db-keys")
		sureFlag := c.Bool("i-am-really-sure")

		if dropFlag && !sureFlag {
			return errors.New("--drop-db-keys requires --i-am-really-sure")
		}

		shouldDrop := dropFlag && sureFlag
		if !shouldDrop {
			// interactive prompt
			fmt.Printf("\n" +
				"WARNING: the next step will DROP the private_key column from the\n" +
				"actors table. This is IRREVERSIBLE. All key material in the database\n" +
				"will be permanently deleted.\n" +
				"\n" +
				"Verify that your keys are present in the keystore directory above\n" +
				"before continuing. For scripted use, pass:\n" +
				"  --drop-db-keys --i-am-really-sure\n\n")
			fmt.Printf("Drop private_key column? [y/N] ")

			scanner := bufio.NewScanner(os.Stdin)
			if !scanner.Scan() {
				return nil
			}
			answer := strings.TrimSpace(strings.ToLower(scanner.Text()))
			if answer != "y" && answer != "yes" {
				fmt.Println("aborted -- keys exported but column retained")
				return nil
			}
		}

		if err := wallet.DropPrivateKeyColumn(db); err != nil {
			return errors.Wrap(err, "failed to drop private_key column")
		}
		fmt.Println("dropped private_key column from actors table")

		return nil
	},
}
