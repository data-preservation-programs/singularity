package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/data-preservation-programs/singularity/cmd/admin"
	"github.com/data-preservation-programs/singularity/cmd/dataset"
	"github.com/data-preservation-programs/singularity/cmd/datasource"
	"github.com/data-preservation-programs/singularity/cmd/datasource/inspect"
	"github.com/data-preservation-programs/singularity/cmd/deal"
	"github.com/data-preservation-programs/singularity/cmd/deal/schedule"
	"github.com/data-preservation-programs/singularity/cmd/deal/spadepolicy"
	"github.com/data-preservation-programs/singularity/cmd/ez"
	"github.com/data-preservation-programs/singularity/cmd/run"
	"github.com/data-preservation-programs/singularity/cmd/tool"
	"github.com/data-preservation-programs/singularity/cmd/wallet"
	"github.com/mattn/go-shellwords"
	"github.com/pkg/errors"
	"github.com/rclone/rclone/lib/terminal"
	"github.com/urfave/cli/v2"
)

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
    For all other networks:
      * Set LOTUS_API to your network's Lotus API endpoint
      * Set MARKET_DEAL_URL to empty string
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
	},
	Commands: []*cli.Command{
		ez.PrepCmd,
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
		VersionCmd,
		DownloadCmd,
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
			Name:     "run",
			Category: "Daemons",
			Usage:    "Run different singularity components",
			Subcommands: []*cli.Command{
				run.APICmd,
				run.DatasetWorkerCmd,
				run.ContentProviderCmd,
				run.DealTrackerCmd,
				run.DealMakerCmd,
			},
		},
		{
			Name:     "dataset",
			Category: "Operations",
			Usage:    "Dataset management",
			Subcommands: []*cli.Command{
				dataset.CreateCmd,
				dataset.ListDatasetCmd,
				dataset.UpdateCmd,
				dataset.RemoveDatasetCmd,
				dataset.AddWalletCmd,
				dataset.ListWalletCmd,
				dataset.RemoveWalletCmd,
				dataset.AddPieceCmd,
				dataset.ListPiecesCmd,
			},
		},
		{
			Name:     "datasource",
			Category: "Operations",
			Usage:    "Data source management",
			Subcommands: []*cli.Command{
				datasource.AddCmd,
				datasource.ListCmd,
				datasource.StatusCmd,
				datasource.RemoveCmd,
				datasource.CheckCmd,
				datasource.UpdateCmd,
				datasource.RescanCmd,
				datasource.DagGenCmd,
				datasource.RepackCmd,
				{
					Name:  "inspect",
					Usage: "Get preparation status of a data source",
					Subcommands: []*cli.Command{
						inspect.ChunksCmd,
						inspect.ItemsCmd,
						inspect.DagsCmd,
						inspect.ChunkDetailCmd,
						inspect.ItemDetailCmd,
						inspect.PathCmd,
					},
				},
			},
		},
		{
			Name:     "wallet",
			Category: "Operations",
			Usage:    "Wallet management",
			Subcommands: []*cli.Command{
				wallet.ImportCmd,
				wallet.ListCmd,
				wallet.AddRemoteCmd,
				wallet.RemoveCmd,
			},
		},
		{
			Name:     "tool",
			Category: "Tooling",
			Usage:    "Tools used for development and debugging",
			Subcommands: []*cli.Command{
				tool.ExtractCarCmd,
			},
		},
	},
}

var originalHelpPrinter = cli.HelpPrinter

var Version string

func SetVersion(versionJSON []byte) error {
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

func SetupHelpPager() {
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

func RunArgsInTest(ctx context.Context, args string) (string, string, error) {
	App.ExitErrHandler = func(c *cli.Context, err error) {
	}
	parser := shellwords.NewParser()
	parser.ParseEnv = true // Enable environment variable parsing
	parsedArgs, err := parser.Parse(args)
	if err != nil {
		return "", "", err
	}

	// Create pipes
	rOut, wOut, _ := os.Pipe()
	rErr, wErr, _ := os.Pipe()

	// Save current stdout and stderr
	oldOut := os.Stdout
	oldErr := os.Stderr
	oldAppWriterOut := App.Writer
	oldAppWriterErr := App.ErrWriter
	defer func() {
		os.Stdout = oldOut
		os.Stderr = oldErr
		App.Writer = oldAppWriterOut
		App.ErrWriter = oldAppWriterErr
	}()

	// Overwrite the stdout and stderr
	os.Stdout = wOut
	os.Stderr = wErr
	App.Writer = wOut
	App.ErrWriter = wErr

	outC := make(chan string) // Buffered to prevent goroutine leak
	errC := make(chan string)
	go func() {
		out, _ := io.ReadAll(rOut)
		outC <- string(out)
	}()
	go func() {
		out, _ := io.ReadAll(rErr)
		errC <- string(out)
	}()

	err = App.RunContext(ctx, parsedArgs)

	// Close the pipes
	wOut.Close()
	wErr.Close()

	// Wait for the output from the goroutines
	outputOut := <-outC
	outputErr := <-errC

	return outputOut, outputErr, err
}
