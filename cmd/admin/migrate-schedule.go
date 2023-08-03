package admin

import (
	"github.com/data-preservation-programs/singularity/migrate"
	"github.com/urfave/cli/v2"
)

var MigrateScheduleCmd = &cli.Command{
	Name:  "migrate-schedule",
	Usage: "Migrate schedule from old singularity mongodb",
	Description: "Migrate schedules from singularity V1 to V2. Note that\n" +
		"  1. You must complete dataset migration first\n" +
		"  2. You will need to import all relevant private keys to the database\n" +
		"  3. All new schedules will be created with status 'paused'\n" +
		"  4. The deal states will not be migrated over as it will be populated with deal tracker automatically.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "mongo-connection-string",
			Usage:   "MongoDB connection string",
			EnvVars: []string{"MONGO_CONNECTION_STRING"},
			Value:   "mongodb://localhost:27017",
		},
	},
	Action: migrate.MigrateSchedule,
}
