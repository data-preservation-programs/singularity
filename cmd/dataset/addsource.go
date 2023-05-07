package dataset

import (
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/handler/dataset"
	"github.com/urfave/cli/v2"
)

var AddSourceCmd = &cli.Command{
	Name:      "add-source",
	Usage:     "Add a new data source to the dataset",
	ArgsUsage: "DATASET_NAME SOURCE_PATH",
	Description: "There are three supported source types:\n" +
		"To add a local file system source, use the path to the directory, i.e. /mnt/dataset\n" +
		"To add a HTTP download sites, i.e. nginx directory explorer, use the URL, i.e. http://download.org\n" +
		"To add an S3 path, use the s3:// prefix, i.e. s3://bucket-name/path/to/dataset",
	Flags: []cli.Flag{
		&cli.DurationFlag{
			Name:        "scan-interval",
			Usage:       "Interval to rescan the data source, or to handle the pushed data source",
			Category:    "Scanning",
			Value:       0,
			DefaultText: "disabled",
		},
		&cli.BoolFlag{
			Name:     "push-only",
			Usage:    "If set to true, the data source will not be scanned, only pushed items will be handled",
			Category: "Scanning",
			Value:    false,
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
		db := database.MustOpenFromCLI(c)
		_, err := dataset.AddSourceHandler(
			db,
			dataset.AddSourceRequest{
				DatasetName:       c.Args().Get(0),
				SourcePath:        c.Args().Get(1),
				ScanInterval:      c.Duration("scan-interval"),
				HTTPHeaders:       c.StringSlice("http-header"),
				S3Endpoint:        c.String("s3-endpoint"),
				S3AccessKeyID:     c.String("s3-access-key-id"),
				S3SecretAccessKey: c.String("s3-secret-access-key"),
				S3Region:          c.String("s3-region"),
				PushOnly:          c.Bool("push-only"),
			},
		)
		return err
	},
}
