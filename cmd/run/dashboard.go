package run

import (
	"github.com/data-preservation-programs/go-singularity/dashboardapi"
	"github.com/urfave/cli/v2"
)

var (
	DashboardCmd = &cli.Command{
		Name:  "dashboard",
		Usage: "Run the singularity dashboard",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "bind",
				Usage: "Bind address for the dashboard server",
				Value: "127.0.0.1",
			},
			&cli.IntFlag{
				Name:  "port",
				Usage: "Port for the dashboard server",
				Value: 9090,
			},
		},
		Action: dashboardapi.Run,
	}
)
