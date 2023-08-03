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
		"  2. All new schedules will be created with status 'paused'\n" +
		"  3. The deal states will not be migrated over as it will be populated with deal tracker automatically\n" +
		"  4. --output-csv is no longer supported. We will provide a new tool in the future\n" +
		"  5. # of replicas is no longer supported as part of the schedule. We will make this a configurable policy in the future\n" +
		"  6. --force is no longer supported. We may add similar support to ignore all policy restrictions in the future\n" +
		"  7. --offline is no longer supported. It will be always offline deal for legacy market and online deal for boost market if URL template is configured",
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
