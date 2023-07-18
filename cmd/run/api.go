package run

import (
	"github.com/data-preservation-programs/singularity/api"
	"github.com/urfave/cli/v2"
)

var (
	APICmd = &cli.Command{
		Name:  "api",
		Usage: "Run the singularity API, including the admin API and the WebDAV API",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:     "enable-admin",
				Category: "Admin API",
				Usage:    "Enable the Admin API",
				Value:    true,
			},
			&cli.StringFlag{
				Name:     "admin-bind",
				Category: "Admin API",
				Usage:    "Bind address for the Admin API server",
				Value:    "127.0.0.1:9090",
			},
			&cli.BoolFlag{
				Name:     "enable-webdav",
				Category: "WebDAV API",
				Usage:    "Enable the WebDAV API",
				Value:    false,
			},
			&cli.StringFlag{
				Name:     "webdav-bind",
				Category: "WebDAV API",
				Usage:    "Webdav API Bind address",
				Value:    "127.0.0.1:9091",
			},
			&cli.UintFlag{
				Name:     "webdav-source-id",
				Category: "WebDAV API",
				Usage:    "ID of Data source used to back the webdav API. The data source must be a local directory",
				Required: true,
			},
		},
		Action: api.Run,
	}
)
