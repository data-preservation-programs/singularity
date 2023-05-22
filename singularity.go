package main

import (
	"github.com/data-preservation-programs/singularity/cmd"
	"github.com/data-preservation-programs/singularity/cmd/admin"
	"github.com/data-preservation-programs/singularity/cmd/datasource"
	"github.com/data-preservation-programs/singularity/cmd/datasource/status"
	"github.com/data-preservation-programs/singularity/cmd/deal/schedule"
	"github.com/data-preservation-programs/singularity/cmd/deal/spadepolicy"
	"github.com/data-preservation-programs/singularity/cmd/wallet"
	"log"
	"os"

	"github.com/data-preservation-programs/singularity/cmd/dataset"
	"github.com/data-preservation-programs/singularity/cmd/deal"
	"github.com/data-preservation-programs/singularity/cmd/run"
	"github.com/data-preservation-programs/singularity/util/must"
	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli/v2"
)

// @title Singularity API
// @version beta
// @description This is the API for Singularity, a tool for large-scale clients with PB-scale data onboarding to Filecoin network.
// @host localhost:9090
// @BasePath /api
// @securityDefinitions none
func main() {
	app := &cli.App{
		Name:                 "singularity",
		Usage:                "A tool for large-scale clients with PB-scale data onboarding to Filecoin network",
		EnableBashCompletion: true,
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
			{
				Name:  "admin",
				Usage: "Admin commands",
				Subcommands: []*cli.Command{
					admin.InitCmd,
					admin.ResetCmd,
					admin.MigrateCmd,
				},
			},
			cmd.DownloadCmd,
			{
				Name:  "deal",
				Usage: "Replication / Deal making management",
				Subcommands: []*cli.Command{
					{
						Name:  "schedule",
						Usage: "Schedule deals",
						Subcommands: []*cli.Command{
							schedule.CreateCmd,
							schedule.ListCmd,
							schedule.PauseCmd,
							schedule.ResumeCmd,
						},
					},
					{
						Name:  "spade-policy",
						Usage: "Manage SPADE policies",
						Subcommands: []*cli.Command{
							spadepolicy.CreateCmd,
							spadepolicy.ListCmd,
							spadepolicy.RemoveCmd,
						},
					},
					deal.SendManualCmd,
					deal.ListCmd,
				},
			},
			{
				Name:  "run",
				Usage: "Run different singularity components",
				Subcommands: []*cli.Command{
					run.ApiCmd,
					run.DatasetWorkerCmd,
					run.ContentProviderCmd,
					run.DealMakerCmd,
					run.SpadeAPICmd,
				},
			},
			{
				Name:  "dataset",
				Usage: "Dataset management",
				Subcommands: []*cli.Command{
					dataset.CreateCmd,
					dataset.ListDatasetCmd,
					dataset.RemoveDatasetCmd,
					dataset.AddWalletCmd,
					dataset.ListWalletCmd,
					dataset.RemoveWalletCmd,
					dataset.AddPieceCmd,
					dataset.ListPiecesCmd,
				},
			},
			{
				Name:  "datasource",
				Usage: "Data source management",
				Subcommands: []*cli.Command{
					datasource.AddCmd,
					datasource.ListCmd,
					datasource.RemoveCmd,
					datasource.CheckCmd,
					datasource.UpdateCmd,
					{
						Name:  "status",
						Usage: "Get preparation status of a data source",
						Subcommands: []*cli.Command{
							status.SummaryCmd,
							status.ChunksCmd,
							status.ItemsCmd,
						},
					},
				},
			},
			{
				Name:  "wallet",
				Usage: "Wallet management",
				Subcommands: []*cli.Command{
					wallet.ImportCmd,
					wallet.ListCmd,
					wallet.AddRemoteCmd,
					wallet.RemoveCmd,
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
