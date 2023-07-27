package run

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/dealmaker"
	"github.com/data-preservation-programs/singularity/service/epochutil"
	"github.com/urfave/cli/v2"
)

var DealMakerCmd = &cli.Command{
	Name:  "dealmaker",
	Usage: "Start a deal making/tracking worker to process deal making",
	Action: func(c *cli.Context) error {
		db, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		if err := model.AutoMigrate(db); err != nil {
			return err
		}

		lotusAPI := c.String("lotus-api")
		lotusToken := c.String("lotus-token")
		err = epochutil.Initialize(c.Context, lotusAPI, lotusToken)
		if err != nil {
			return err
		}

		dealMaker, err := dealmaker.NewDealMakerService(db, c.String("lotus-api"), c.String("lotus-token"))
		if err != nil {
			return cli.Exit(err.Error(), 1)
		}
		return dealMaker.Run(c.Context)
	},
}
