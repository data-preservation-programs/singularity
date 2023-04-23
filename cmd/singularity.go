package main

import (
	"github.com/data-preservation-programs/go-singularity/cmd/prep"
	"github.com/data-preservation-programs/go-singularity/cmd/repl"
	"github.com/data-preservation-programs/go-singularity/cmd/run"
	"github.com/data-preservation-programs/go-singularity/util/must"
	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

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
		},
		Commands: []*cli.Command{
			DownloadCmd,
			{
				Name:    "replication",
				Aliases: []string{"repl"},
				Usage:   "Replication / deal making management",
				Subcommands: []*cli.Command{
					repl.ScheduleCmd,
					repl.ListCmd,
					repl.PauseCmd,
					repl.ResumeCmd,
					repl.StatusCmd,
				},
			},
			{
				Name:  "run",
				Usage: "Run different singularity components",
				Subcommands: []*cli.Command{
					run.PrepWorkerCmd,
					run.DataListenerCmd,
					run.ContentProviderCmd,
					run.ReplicationCmd,
					run.SpadeAPICmd,
				},
			},
			{
				Name:    "preparation",
				Aliases: []string{"prep"},
				Usage:   "Dataset preparation management",
				Subcommands: []*cli.Command{
					prep.CreateCmd,
					prep.ListCmd,
					prep.PauseCmd,
					prep.RemoveCmd,
					prep.ResumeCmd,
					prep.StatusCmd,
					prep.UpdateCmd,
					prep.QueueCmd,
					prep.ExportCmd,
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
