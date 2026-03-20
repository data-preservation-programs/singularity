package run

import (
	"github.com/data-preservation-programs/singularity/api"
	"github.com/urfave/cli/v2"
)

var APICmd = &cli.Command{
	Name:  "api",
	Usage: "Run the singularity API",
	Flags: []cli.Flag{
		NoAutoMigrateFlag,
		&cli.StringFlag{
			Name:  "bind",
			Usage: "Bind address for the API server",
			Value: ":9090",
		},
	},
	Action: func(c *cli.Context) error {
		// run automigrate + legacy key check before handing off to api.Run,
		// which opens its own db connection internally
		db, closer, err := openAndMigrate(c)
		if err != nil {
			return err
		}
		closer.Close()
		_ = db

		return api.Run(c)
	},
}
