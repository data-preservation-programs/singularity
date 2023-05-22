package main

import (
	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/rclone/rclone/fs"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	var commands []*cli.Command
	for _, r := range fs.Registry {
		cmd := datasource.OptionsToCLIFlags(r)
		commands = append(commands, cmd)
	}

	app := &cli.App{
		Name:                 "singularity",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name: "dataset",
				Subcommands: []*cli.Command{
					{
						Name:        "add-source",
						Subcommands: commands,
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
