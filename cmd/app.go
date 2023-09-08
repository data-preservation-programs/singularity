package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/admin"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/cmd/dataprep"
	"github.com/data-preservation-programs/singularity/cmd/deal"
	"github.com/data-preservation-programs/singularity/cmd/deal/schedule"
	"github.com/data-preservation-programs/singularity/cmd/ez"
	"github.com/data-preservation-programs/singularity/cmd/run"
	"github.com/data-preservation-programs/singularity/cmd/storage"
	"github.com/data-preservation-programs/singularity/cmd/tool"
	"github.com/data-preservation-programs/singularity/cmd/wallet"
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-log/v2"
	"github.com/rclone/rclone/lib/terminal"
	"github.com/urfave/cli/v2"
)

var logger = log.Logger("singularity/cmd")

var App = &cli.App{
	Name:  "singularity",
	Usage: "A tool for large-scale clients with PB-scale data onboarding to Filecoin network",
	Description: `Database Backend Support:
  Singularity supports multiple database backend: sqlite3, postgres, mysql5.7+
  Use '--database-connection-string' or $DATABASE_CONNECTION_STRING to specify the database connection string.
    Example for postgres  - postgres://user:pass@example.com:5432/dbname
    Example for mysql     - mysql://user:pass@tcp(localhost:3306)/dbname?parseTime=true
    Example for sqlite3   - sqlite:/absolute/path/to/database.db
                or        - sqlite:relative/path/to/database.db

Network Support:
  Default settings in Singularity are for Mainnet. You may set below environment values for other network:
    For Calibration network:
      * Set LOTUS_API to https://api.calibration.node.glif.io/rpc/v1
      * Set MARKET_DEAL_URL to https://marketdeals-calibration.s3.amazonaws.com/StateMarketDeals.json.zst
      * Set LOTUS_TEST to 1
    For all other networks:
      * Set LOTUS_API to your network's Lotus API endpoint
      * Set MARKET_DEAL_URL to empty string
      * Set LOTUS_TEST to 0 or 1 based on whether the network address starts with 'f' or 't'
    Switching between different networks with the same database instance is not recommended.`,
	EnableBashCompletion: true,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "database-connection-string",
			Usage:       "Connection string to the database",
			DefaultText: "sqlite:" + "./singularity.db",
			Value:       "sqlite:" + "./singularity.db",
			EnvVars:     []string{"DATABASE_CONNECTION_STRING"},
		},
		&cli.BoolFlag{
			Name:  "json",
			Usage: "Enable JSON output",
			Value: false,
		},
		&cli.BoolFlag{
			Name:  "verbose",
			Usage: "Enable verbose output. This will print more columns for the result as well as full error trace",
			Value: false,
		},
		&cli.StringFlag{
			Name:     "lotus-api",
			Category: "Lotus",
			Usage:    "Lotus RPC API endpoint",
			Value:    "https://api.node.glif.io/rpc/v1",
			EnvVars:  []string{"LOTUS_API"},
		},
		&cli.StringFlag{
			Name:     "lotus-token",
			Category: "Lotus",
			Usage:    "Lotus RPC API token",
			Value:    "",
			EnvVars:  []string{"LOTUS_TOKEN"},
		},
		&cli.BoolFlag{
			Name:     "lotus-test",
			Category: "Lotus",
			Usage:    "Whether the runtime environment is using Testnet.",
			EnvVars:  []string{"LOTUS_TEST"},
			Action: func(c *cli.Context, testnet bool) error {
				if testnet {
					address.CurrentNetwork = address.Testnet
					logger.Infow("Current network is set to Testnet")
				}
				return nil
			},
		},
	},
	Commands: []*cli.Command{
		ez.PrepCmd,
		VersionCmd,
		{
			Name:     "admin",
			Usage:    "Admin commands",
			Category: "Operations",
			Subcommands: []*cli.Command{
				admin.InitCmd,
				admin.ResetCmd,
				admin.MigrateDatasetCmd,
				admin.MigrateScheduleCmd,
			},
		},
		DownloadCmd,
		tool.ExtractCarCmd,
		{
			Name:     "deal",
			Usage:    "Replication / Deal making management",
			Category: "Operations",
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
				deal.SendManualCmd,
				deal.ListCmd,
			},
		},
		{
			Name:     "run",
			Category: "Daemons",
			Usage:    "run different singularity components",
			Subcommands: []*cli.Command{
				run.APICmd,
				run.DatasetWorkerCmd,
				run.ContentProviderCmd,
				run.DealTrackerCmd,
				run.DealPusherCmd,
			},
		},
		{
			Name:     "wallet",
			Category: "Operations",
			Usage:    "Wallet management",
			Subcommands: []*cli.Command{
				wallet.ImportCmd,
				wallet.ListCmd,
				wallet.RemoveCmd,
			},
		},
		{
			Name:     "storage",
			Category: "Operations",
			Usage:    "Create and manage storage system connections",
			Subcommands: []*cli.Command{
				storage.CreateCmd,
				storage.ExploreCmd,
				storage.ListCmd,
				storage.RemoveCmd,
				storage.UpdateCmd,
			},
		},
		{
			Name:     "prep",
			Category: "Operations",
			Usage:    "Create and manage dataset preparations",
			Subcommands: []*cli.Command{
				dataprep.CreateCmd,
				dataprep.ListCmd,
				dataprep.StatusCmd,
				dataprep.AttachSourceCmd,
				dataprep.AttachOutputCmd,
				dataprep.DetachOutputCmd,
				dataprep.StartScanCmd,
				dataprep.PauseScanCmd,
				dataprep.StartPackCmd,
				dataprep.PausePackCmd,
				dataprep.StartDagGenCmd,
				dataprep.PauseDagGenCmd,
				dataprep.ListPiecesCmd,
				dataprep.AddPieceCmd,
				dataprep.ExploreCmd,
				dataprep.AttachWalletCmd,
				dataprep.ListWalletsCmd,
				dataprep.DetachWalletCmd,
			},
		},
	},
}

var originalHelpPrinter = cli.HelpPrinter

var Version string

func SetVersionJSON(versionJSON []byte) error {
	var v struct {
		Version string `json:"version"`
	}
	err := json.Unmarshal(versionJSON, &v)
	if err != nil {
		return errors.Wrap(err, "cannot unmarshal version")
	}

	Version = v.Version
	return nil
}

func SetupErrorHandler() {
	App.ExitErrHandler = func(c *cli.Context, err error) {
		if err == nil {
			return
		}
		if c.Bool("verbose") {
			report := fmt.Sprintf("%+v\n\n", err)
			_, _ = App.ErrWriter.Write([]byte(report))
		}
		concise := cliutil.Failure(err.Error()) + "\n"
		_, _ = App.ErrWriter.Write([]byte(concise))
	}
}

const subCommandHelpTemplate = `NAME:
   {{template "helpNameTemplate" .}}

USAGE:
   {{if .UsageText}}{{wrap .UsageText 3}}{{else}}{{.HelpName}} {{if .VisibleFlags}}command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Description}}

DESCRIPTION:
   {{template "descriptionTemplate" .}}{{end}}{{if .VisibleCommands}}

COMMANDS:{{template "visibleCommandCategoryTemplate" .}}{{end}}{{if .VisibleFlagCategories}}

OPTIONS:{{template "visibleFlagCategoryTemplate" .}}{{else if .VisibleFlags}}

OPTIONS:{{template "visibleFlagTemplate" .}}{{end}}
`

func SetupHelpPager() {
	cli.SubcommandHelpTemplate = subCommandHelpTemplate
	//nolint:errcheck
	cli.HelpPrinter = func(w io.Writer, templ string, data any) {
		var helpText bytes.Buffer
		originalHelpPrinter(&helpText, templ, data)
		numLines := strings.Count(helpText.String(), "\n")
		_, maxLinesWithoutPager := terminal.GetSize()
		if numLines < maxLinesWithoutPager-1 {
			w.Write(helpText.Bytes())
			return
		}
		pager := os.Getenv("PAGER")
		if pager == "" {
			pager = "less"
		}

		pagerPath, err := exec.LookPath(pager)
		if err != nil {
			w.Write(helpText.Bytes())
			return
		}
		cmd := exec.Command(pagerPath)
		pagerIn, err := cmd.StdinPipe()
		cmd.Stdout = w
		if err != nil {
			w.Write(helpText.Bytes())
			return
		}

		if err := cmd.Start(); err != nil {
			w.Write(helpText.Bytes())
			return
		}

		if _, err := io.Copy(pagerIn, &helpText); err != nil {
			w.Write(helpText.Bytes())
			return
		}
		pagerIn.Close()
		cmd.Wait()
	}
}
