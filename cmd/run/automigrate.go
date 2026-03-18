package run

import (
	"io"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

var NoAutoMigrateFlag = &cli.BoolFlag{
	Name:  "no-automigrate",
	Usage: "skip automatic database migration and correctness checks on startup; only use if you run 'admin init' on every upgrade or manually before starting daemons",
}

// opens the database, runs AutoMigrate (unless --no-automigrate), and checks
// for legacy keys that need export. returns db and closer as usual.
func openAndMigrate(c *cli.Context) (*gorm.DB, io.Closer, error) {
	db, closer, err := database.OpenFromCLI(c)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	if !c.Bool("no-automigrate") {
		if err := model.AutoMigrate(db); err != nil {
			closer.Close()
			return nil, nil, errors.Wrap(err, "automigrate failed")
		}

		if n := wallet.CountLegacyKeys(db); n > 0 {
			closer.Close()
			return nil, nil, errors.Errorf(
				"%d actor(s) have private keys in the database that are not usable by current code.\n"+
					"Run 'singularity wallet export-keys' to migrate them to the filesystem keystore",
				n)
		}
	}

	return db, closer, nil
}
