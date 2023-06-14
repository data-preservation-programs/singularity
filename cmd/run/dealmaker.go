package run

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/urfave/cli/v2"
)

var DealMakerCmd = &cli.Command{
	Name:  "dealmaker",
	Usage: "[Alpha] Start a deal making/tracking worker to process deal making",
	Flags: []cli.Flag{
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
		dealMaker, err := service.NewDealMakerService(db, c.String("lotus-api"), c.String("lotus-token"))
		if err != nil {
			return cli.Exit(err.Error(), 1)
		}
		return dealMaker.Run(c.Context)
	},
}
