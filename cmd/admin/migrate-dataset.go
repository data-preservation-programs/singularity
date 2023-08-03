package admin

import (
	"github.com/data-preservation-programs/singularity/migrate"
	"github.com/urfave/cli/v2"
)

var MigrateDatasetCmd = &cli.Command{
	Name:  "migrate-dataset",
	Usage: "Migrate dataset from old singularity mongodb",
	Description: "Migrate datasets from singularity V1 to V2. The steps include\n" +
		"  1. Create a new dataset in V2.\n" +
		"  2. Create a new datasource which is either an S3 path or local path.\n" +
		"  3. Create all folder structures and files in the new dataset.\n" +
		"Caveats:\n" +
		"  1. The created dataset won't be compatible with the new dataset worker.\n" +
		"     So do not attempt to trigger a data preparation on migrated dataset.\n" +
		"  2. The folder CID won't be generated or migrated",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "mongo-connection-string",
			Usage:   "MongoDB connection string",
			EnvVars: []string{"MONGO_CONNECTION_STRING"},
			Value:   "mongodb://localhost:27017",
		},
	},
	Action: migrate.MigrateDataset,
}
