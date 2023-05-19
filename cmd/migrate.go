package cmd

import (
	"github.com/urfave/cli/v2"
)

var MigrateCmd = &cli.Command{
	Name:  "migrate",
	Usage: "Migrate data from old singularity mongodb",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "mongo-connection-string",
			Usage:   "MongoDB connection string",
			EnvVars: []string{"MONGO_CONNECTION_STRING"},
			Value:   "mongodb://localhost:27017",
		},
	},
	// TODO Action: migrate.Migrate,
}
