package cmd

import (
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/urfave/cli/v2"
)

var DownloadCmd = &cli.Command{
	Name:      "download",
	Usage:     "Download a CAR file from the metadata API",
	Category:  "Utility",
	ArgsUsage: "PIECE_CID",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "api",
			Usage: "URL of the metadata API",
			Value: "http://127.0.0.1:9090",
		},
		&cli.StringSliceFlag{
			Name:     "http-header",
			Category: "HTTP data source",
			Aliases:  []string{"H"},
			EnvVars:  []string{"HTTP_HEADER"},
			Usage:    "http headers to be passed with the request (i.e. key=value). The value shoud not be encoded",
		},
		&cli.StringFlag{
			Name:     "s3-endpoint",
			Usage:    "Custom S3 endpoint",
			Category: "S3 data source",
			EnvVars:  []string{"S3_ENDPOINT"},
		},
		&cli.StringFlag{
			Name:     "s3-region",
			Usage:    "S3 region to use with AWS S3",
			Category: "S3 data source",
			EnvVars:  []string{"S3_REGION"},
		},
		&cli.StringFlag{
			Name:     "s3-access-key-id",
			Usage:    "IAM access key ID",
			Category: "S3 data source",
			EnvVars:  []string{"AWS_ACCESS_KEY_ID"},
		},
		&cli.StringFlag{
			Name:     "s3-secret-access-key",
			Usage:    "IAM secret access key",
			Category: "S3 data source",
			EnvVars:  []string{"AWS_SECRET_ACCESS_KEY"},
		},
	},
	Action: func(c *cli.Context) error {
		model.DisableEncryption = true
		piece := c.Args().First()
		api := c.String("api")
		return handler.DownloadHandler(piece, api, nil)
	},
}
