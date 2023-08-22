package admin

import (
	"github.com/data-preservation-programs/singularity/migrate"
	"github.com/urfave/cli/v2"
)

var MigrateDatasetCmd = &cli.Command{
	Name:  "migrate-dataset",
	Usage: "Migrate dataset from old singularity mongodb",
	Description: "Migrate datasets from singularity V1 to V2. Those steps include\n" +
		"  1. Create a new dataset in V2.\n" +
		"  2. Create a new datasource which is either an S3 path or local path.\n" +
		"  3. Create all folder structures and files in the new dataset.\n" +
		"Caveats:\n" +
		"  1. The created dataset won't be compatible with the new dataset worker.\n" +
		"     So do not attempt to resume a data preparation or push new files onto migrated dataset.\n" +
		"     You can make deals or browse the dataset without issues.\n" +
		"  2. The folder CID won't be generated or migrated due to the complexity",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "mongo-connection-string",
			Usage:   "MongoDB connection string",
			EnvVars: []string{"MONGO_CONNECTION_STRING"},
			Value:   "mongodb://localhost:27017",
		},
		&cli.BoolFlag{
			Name:  "skip-files",
			Usage: "Skip migrating details about files and folders. This will make the migration much faster. Useful if you only want to make deals.",
			Value: false,
		},
	},
	Action: migrate.MigrateDataset,
}
