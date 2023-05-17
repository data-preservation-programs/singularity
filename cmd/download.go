package cmd

import (
	"github.com/data-preservation-programs/go-singularity/handler"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"strings"
)

var DownloadCmd = &cli.Command{
	Name:      "download",
	Usage:     "Download a CAR file from the metadata API",
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
		httpHeaders := c.StringSlice("http-header")
		s3Endpoint := c.String("s3-endpoint")
		s3Region := c.String("s3-region")
		s3AccessKeyID := c.String("s3-access-key-id")
		s3SecretAccessKey := c.String("s3-secret-access-key")
		var meta model.Metadata
		if len(httpHeaders) > 0 {
			headers := map[string]string{}
			for _, header := range httpHeaders {
				parts := strings.SplitN(header, "=", 2)
				if len(parts) != 2 {
					return errors.New("invalid header: " + header)
				}

				headers[parts[0]] = parts[1]
			}

			var err error
			meta, err = model.HTTPMetadata{Headers: headers}.Encode()
			if err != nil {
				return err
			}
		} else {
			var err error
			meta, err = model.S3Metadata{
				Region:          s3Region,
				Endpoint:        s3Endpoint,
				AccessKeyID:     s3AccessKeyID,
				SecretAccessKey: s3SecretAccessKey,
			}.Encode()
			if err != nil {
				return err
			}
		}
		return handler.DownloadHandler(piece, api, meta)
	},
}
