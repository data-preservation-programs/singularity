package main

import (
	"github.com/data-preservation-programs/go-singularity/cmd/dataset"
	"github.com/data-preservation-programs/go-singularity/cmd/repl"
	"github.com/data-preservation-programs/go-singularity/cmd/run"
	"github.com/data-preservation-programs/go-singularity/util/must"
	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

// @title Singularity API
// @version beta
// @description This is the API for Singularity, a tool for large-scale clients with PB-scale data onboarding to Filecoin network.
// @host localhost:9090
// @BasePath /admin/api
// @securityDefinitions none
func main() {
	app := &cli.App{
		Name:  "singularity",
		Usage: "A tool for large-scale clients with PB-scale data onboarding to Filecoin network",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "database-connection-string",
				Usage: "Connection string to the database.\n" +
					"Supported database: sqlite3, postgres, mysql\n" +
					"Example for postgres - postgres://user:pass@example.com:5432/dbname\n" +
					"Example for mysql    - mysql://user:pass@example.com:5432/dbname\n" +
					"Example for sqlite3  - sqlite:////absolute/path/to/database.db\n" +
					"            or       - sqlite:///relative/path/to/database.db\n",
				DefaultText: "sqlite:" + must.String(os.UserHomeDir()) + "/.singularity/singularity.db",
				Value:       "sqlite:" + must.String(os.UserHomeDir()) + "/.singularity/singularity.db",
				EnvVars:     []string{"DATABASE_CONNECTION_STRING"},
			},
			&cli.StringFlag{
				Name:    "password",
				Usage:   "Password used to derive encryption key for credentials encryption",
				EnvVars: []string{"PASSWORD"},
				Value:   "1234",
			},
			&cli.BoolFlag{
				Name:  "verbose",
				Usage: "Enable verbose logging",
				Value: false,
			},
			&cli.BoolFlag{
				Name:  "json",
				Usage: "Enable JSON output",
				Value: false,
			},
		},
		Commands: []*cli.Command{
			MigrateCmd,
			DownloadCmd,
			InitCmd,
			{
				Name:    "replication",
				Aliases: []string{"repl"},
				Usage:   "Replication / deal making management",
				Subcommands: []*cli.Command{
					repl.ScheduleCmd,
				},
			},
			{
				Name:  "run",
				Usage: "Run different singularity components",
				Subcommands: []*cli.Command{
					run.DatasetWorkerCmd,
					run.DataListenerCmd,
					run.ContentProviderCmd,
					run.ReplicationCmd,
					run.SpadeAPICmd,
					run.ApiCmd,
				},
			},
			{
				Name:  "dataset",
				Usage: "Dataset preparation management",
				Subcommands: []*cli.Command{
					dataset.CreateCmd,
					dataset.ListDatasetCmd,
					dataset.RemoveDatasetCmd,
					dataset.AddSourceCmd,
					dataset.ListSourceCmd,
					dataset.RemoveSourceCmd,
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
