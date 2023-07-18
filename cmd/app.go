package cmd

import (
	"bytes"
	"context"
	"fmt"
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
	"github.com/data-preservation-programs/singularity/cmd/wallet"
	"github.com/data-preservation-programs/singularity/util/must"
	"github.com/mattn/go-shellwords"
	"github.com/urfave/cli/v2"
)

var App = &cli.App{
	Name:                 "singularity",
	Usage:                "A tool for large-scale clients with PB-scale data onboarding to Filecoin network",
	EnableBashCompletion: true,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "database-connection-string",
			Usage: "Connection string to the database.\n" +
				"Supported database: sqlite3, postgres, mysql\n" +
				"Example for postgres  - postgres://user:pass@example.com:5432/dbname\n" +
				"Example for mysql     - mysql://user:pass@tcp(localhost:3306)/dbname?charset=ascii&parseTime=true\n" +
				"                          Note: the database needs to be created using ascii Character Set:" +
				"                                `CREATE DATABASE <dbname> DEFAULT CHARACTER SET ascii`\n" +
				"Example for sqlite3   - sqlite:/absolute/path/to/database.db\n" +
				"            or        - sqlite:relative/path/to/database.db\n",
			DefaultText: "sqlite:" + must.String(os.UserHomeDir()) + "/singularity.db",
			Value:       "sqlite:" + must.String(os.UserHomeDir()) + "/singularity.db",
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
		},
		&cli.StringFlag{
			Name:     "lotus-token",
			Category: "Lotus",
			Usage:    "Lotus RPC API token",
			Value:    "",
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
				admin.MigrateCmd,
			},
		},
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
				run.SpadeAPICmd,
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
	},
}

func RunApp(ctx context.Context, args []string) error {
	printer := cli.HelpPrinter
	//nolint:errcheck
	cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
		var helpText bytes.Buffer
		printer(&helpText, templ, data)
		numLines := strings.Count(helpText.String(), "\n")
		const maxLinesWithoutPager = 40
		if numLines <= maxLinesWithoutPager {
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
	if err := App.RunContext(ctx, args); err != nil {
		return err
	}

	return nil
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

	// Overwrite the stdout and stderr
	os.Stdout = wOut
	os.Stderr = wErr

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

	err = RunApp(ctx, parsedArgs)

	// Close the pipes
	wOut.Close()
	wErr.Close()

	// Restore original stdout and stderr
	os.Stdout = oldOut
	os.Stderr = oldErr

	// Wait for the output from the goroutines
	outputOut := <-outC
	outputErr := <-errC

	// Let's still print it to stdout and stderr
	fmt.Println(outputOut)
	fmt.Fprintln(os.Stderr, outputErr)
	return outputOut, outputErr, err
}
