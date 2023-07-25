package run

import (
	"time"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/replication"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/urfave/cli/v2"
)

var SpadeAPICmd = &cli.Command{
	Name:  "spade-api",
	Usage: "Start a Spade compatible API for storage provider deal proposal self service",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "bind",
			Usage: "Bind address for the API server",
			Value: "127.0.0.1:9091",
		},
	},
	Action: func(c *cli.Context) error {
		db, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		if err := model.AutoMigrate(db); err != nil {
			return err
		}

		h, err := util.InitHost(nil)
		if err != nil {
			return err
		}

		dealMaker := replication.NewDealMaker(
			util.NewLotusClient(c.String("lotus-api"), c.String("lotus-token")),
			h, time.Hour, time.Minute)
		if err != nil {
			return err
		}

		return service.NewSpadeAPIService(db, dealMaker, &replication.DefaultWalletChooser{}, c.String("bind")).Start()
	},
}
