package run

import (
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/data-preservation-programs/go-singularity/service"
	"github.com/urfave/cli/v2"
)

var ContentProviderCmd = &cli.Command{
	Name:  "content-provider",
	Usage: "Start a content provider that serves retrieval requests",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "http-bind",
			Usage: "Address to bind the HTTP server to",
			Value: "127.0.0.1:8088",
		},
	},
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		err := model.InitializeEncryption(c.String("password"), db)
		if err != nil {
			return cli.Exit(err.Error(), 1)
		}
		service.NewContentProviderService(db, c.String("http-bind")).Start()
		return nil
	},
}
