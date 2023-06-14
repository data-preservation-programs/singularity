package run

import (
	"context"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/replication"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/urfave/cli/v2"
)

var SpadeAPICmd = &cli.Command{
	Name:  "spade-api",
	Usage: "[Alpha] Start a Spade compatible API for storage provider deal proposal self service",
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

		h, err := util.InitHost(context.Background(), nil)
		if err != nil {
			return err
		}

		dealMaker, err := replication.NewDealMaker(c.String("lotus-api"), c.String("lotus-token"), h)
		if err != nil {
			return err
		}

		return service.NewSpadeAPIService(db, dealMaker, &replication.WalletChooser{}, c.String("bind")).Start()
	},
}
