package run

import (
	"github.com/data-preservation-programs/go-singularity/api"
	"github.com/urfave/cli/v2"
)

var (
	ApiCmd = &cli.Command{
		Name:  "api",
		Usage: "Run the singularity API",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "bind",
				Usage: "Bind address for the API server",
				Value: "127.0.0.1",
			},
			&cli.IntFlag{
				Name:  "port",
				Usage: "Port for the API server",
				Value: 9090,
			},
		},
		Action: api.Run,
	}
)
