package run

import (
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/service/autodeal"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/urfave/cli/v2"
)

var AutoDealCmd = &cli.Command{
	Name:  "autodeal",
	Usage: "Start the auto-deal daemon to automatically create deal schedules when preparations complete",
	Flags: []cli.Flag{
		&cli.DurationFlag{
			Name:  "check-interval",
			Usage: "How often to check for ready preparations",
			Value: 30 * time.Second,
		},
		&cli.BoolFlag{
			Name:  "enable-batch-mode",
			Usage: "Enable batch processing mode to scan for ready preparations",
			Value: true,
		},
		&cli.BoolFlag{
			Name:  "exit-on-complete",
			Usage: "Exit when there are no more preparations to process",
			Value: false,
		},
		&cli.BoolFlag{
			Name:  "exit-on-error",
			Usage: "Exit when any error occurs",
			Value: false,
		},
		&cli.IntFlag{
			Name:  "max-retries",
			Usage: "Maximum number of retries before backing off (0 = unlimited)",
			Value: 3,
		},
		&cli.DurationFlag{
			Name:  "retry-interval",
			Usage: "Base interval for retry backoff",
			Value: 5 * time.Minute,
		},
		&cli.BoolFlag{
			Name:  "enable-job-hooks",
			Usage: "Enable automatic triggering on job completion (requires worker integration)",
			Value: true,
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		lotusClient := util.NewLotusClient(c.String("lotus-api"), c.String("lotus-token"))

		// Configure the trigger service
		autodeal.DefaultTriggerService.SetEnabled(c.Bool("enable-job-hooks"))

		// Configure the monitor service
		config := autodeal.MonitorConfig{
			CheckInterval:   c.Duration("check-interval"),
			EnableBatchMode: c.Bool("enable-batch-mode"),
			ExitOnComplete:  c.Bool("exit-on-complete"),
			ExitOnError:     c.Bool("exit-on-error"),
			MaxRetries:      c.Int("max-retries"),
			RetryInterval:   c.Duration("retry-interval"),
		}

		monitor := autodeal.NewMonitorService(db, lotusClient, config)
		err = monitor.Run(c.Context)
		if err != nil {
			return errors.WithStack(err)
		}

		return nil
	},
}