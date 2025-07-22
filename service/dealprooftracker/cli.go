package dealprooftracker

import (
    "encoding/json"
    "fmt"
    "os"
    "strconv"
    "time"

    "github.com/urfave/cli/v2"
    "github.com/data-preservation-programs/singularity/database"
)

const defaultPollingInterval = 5 * time.Minute

// CLICommands returns the CLI commands for dealprooftracker
func CLICommands() *cli.Command {
    return &cli.Command{
        Name:  "dealproof",
        Usage: "Deal Proof Tracker commands",
        Subcommands: []*cli.Command{
            {
                Name:   "run",
                Usage:  "Start the deal proof tracker service loop",
                Action: runTracker,
                Flags:  commonFlags(),
            },
            {
                Name:      "live",
                Usage:     "Fetch proof info in real time from Lotus",
                ArgsUsage: "<dealID>",
                Action:    liveProof,
                Flags:     commonFlags(),
            },
            {
                Name:      "db",
                Usage:     "Fetch proof info from the local DB",
                ArgsUsage: "<dealID>",
                Action:    dbProof,
                Flags:     commonFlags(),
            },
            {
                Name:   "health",
                Usage:  "Check DB and Lotus connectivity",
                Action: healthCheck,
                Flags:  commonFlags(),
            },
        },
    }
}

func commonFlags() []cli.Flag {
    return []cli.Flag{
        &cli.StringFlag{
            Name:    "database-connection-string",
            EnvVars: []string{"DATABASE_CONNECTION_STRING"},
            Usage:   "Database connection string",
        },
        &cli.StringFlag{
            Name:    "lotus-api",
            EnvVars: []string{"LOTUS_API"},
            Usage:   "Lotus API endpoint",
        },
        &cli.StringFlag{
            Name:    "lotus-token",
            EnvVars: []string{"LOTUS_TOKEN"},
            Usage:   "Lotus API token",
        },
        &cli.StringFlag{
            Name:  "output",
            Usage: "Output format: plain|json",
            Value: "plain",
        },
        &cli.DurationFlag{
            Name:  "interval",
            Usage: "Polling interval for the tracker service",
            Value: defaultPollingInterval,
        },
    }
}

// Helper to open DB and initialize tracker
func resolveDeps(c *cli.Context) (*ProofTracker, error) {
    connStr := c.String("database-connection-string")
    if connStr == "" {
        return nil, fmt.Errorf("database connection string is required")
    }

    db, _, err := database.OpenWithLogger(connStr)
    if err != nil {
        return nil, fmt.Errorf("failed to open DB: %w", err)
    }

    lotusAPI := c.String("lotus-api")
    if lotusAPI == "" {
        return nil, fmt.Errorf("lotus-api is required")
    }

    lotusToken := c.String("lotus-token")
    interval := c.Duration("interval")

    // Create and return the tracker
    return NewProofTracker(db, lotusAPI, lotusToken, interval), nil
}

func runTracker(c *cli.Context) error {
    tracker, err := resolveDeps(c)
    if err != nil {
        return err
    }
    tracker.Start(c.Context)
    select {} // Block forever
}

func liveProof(c *cli.Context) error {
    tracker, err := resolveDeps(c)
    if err != nil {
        return err
    }

    if c.NArg() < 1 {
        return fmt.Errorf("dealID required")
    }
    
    dealID, err := strconv.ParseUint(c.Args().Get(0), 10, 64)
    if err != nil {
        return fmt.Errorf("invalid dealID: %w", err)
    }

    info, err := tracker.GetLiveProofInfo(c.Context, dealID)
    if err != nil {
        return fmt.Errorf("failed to get live proof info: %w", err)
    }

    return printOutput(c, info)
}

func dbProof(c *cli.Context) error {
    tracker, err := resolveDeps(c)
    if err != nil {
        return err
    }

    if c.NArg() < 1 {
        return fmt.Errorf("dealID required")
    }
    
    dealID, err := strconv.ParseUint(c.Args().Get(0), 10, 64)
    if err != nil {
        return fmt.Errorf("invalid dealID: %w", err)
    }

    info, err := tracker.GetDBProofInfo(c.Context, dealID)
    if err != nil {
        return fmt.Errorf("failed to get DB proof info: %w", err)
    }

    return printOutput(c, info)
}

func healthCheck(c *cli.Context) error {
    tracker, err := resolveDeps(c)
    if err != nil {
        return err
    }

    // Check DB connection
    if err := tracker.db.Exec("SELECT 1").Error; err != nil {
        return fmt.Errorf("DB health check failed: %w", err)
    }

    // Check Lotus connection by attempting to list deals (simplified health check)
    info, checkErr := tracker.GetLiveProofInfo(c.Context, 1) // Use deal ID 1 as a test
    if checkErr != nil {
        return fmt.Errorf("Lotus health check failed: %w", checkErr)
    }
    _ = info // Ignore the result, we just care about connectivity

    fmt.Println("Health check: OK")
    return nil
}

func printOutput(c *cli.Context, v interface{}) error {
    switch c.String("output") {
    case "json":
        return printJSON(v)
    default:
        fmt.Printf("%+v\n", v)
        return nil
    }
}

func printJSON(v interface{}) error {
    enc := json.NewEncoder(os.Stdout)
    enc.SetIndent("", "  ")
    return enc.Encode(v)
}
