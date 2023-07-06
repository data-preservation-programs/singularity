package run

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/replication"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/urfave/cli/v2"
	"time"
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
		&cli.StringFlag{
			Name:     "lotus-api",
			Category: "Lotus",
			Usage:    "Lotus RPC API endpoint, only used to get miner info",
			Value:    "https://api.node.glif.io/rpc/v1",
		},
		&cli.StringFlag{
			Name:     "lotus-token",
			Category: "Lotus",
			Usage:    "Lotus RPC API token, only used to get miner info",
			Value:    "",
		},
	},
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		err := model.AutoMigrate(db)
		if err != nil {
			return handler.NewHandlerError(err)
		}

		h, err := util.InitHost(nil)
		if err != nil {
			return err
		}

		dealMaker := replication.NewDealMaker(c.String("lotus-api"), c.String("lotus-token"), h, time.Hour, time.Minute)
		if err != nil {
			return err
		}

		return service.NewSpadeAPIService(db, dealMaker, &replication.DefaultWalletChooser{}, c.String("bind")).Start()
	},
}
