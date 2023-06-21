package datasource

import (
	"context"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/rjNemo/underscore"
	"github.com/urfave/cli/v2"
	"time"
)

var UpdateCmd = &cli.Command{
	Name:      "update",
	Usage:     "Update the config options of a source",
	ArgsUsage: "<source_id>",
	Flags: func() []cli.Flag {
		var flags []cli.Flag
		for _, cmd := range AddCmd.Subcommands {
			cmdFlags := underscore.Map(cmd.Flags, func(flag cli.Flag) cli.Flag {
				stringFlag, ok := flag.(*cli.StringFlag)
				if !ok {
					return flag
				}
				stringFlag.Required = false
				stringFlag.Category = "Options for " + cmd.Name
				return stringFlag
			})
			flags = append(flags, cmdFlags...)
		}
		keys := make(map[string]cli.Flag)
		newFlags := make([]cli.Flag, 0)
		for _, flag := range flags {
			if _, ok := keys[flag.Names()[0]]; !ok {
				keys[flag.Names()[0]] = flag
				newFlags = append(newFlags, flag)
			}
		}
		return newFlags
	}(),
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		config := map[string]string{}
		var deleteAfterExport *bool
		var rescanInterval *time.Duration
		for _, name := range c.LocalFlagNames() {
			if c.IsSet(name) {
				if name == "delete-after-export" {
					b := c.Bool(name)
					deleteAfterExport = &b
					continue
				}
				if name == "rescan-interval" {
					d := c.Duration(name)
					rescanInterval = &d
					continue
				}
				value := c.String(name)
				config[name] = value
			}
		}

		source, err := datasource.UpdateSourceHandler(
			db,
			context.Background(),
			c.Args().Get(0),
			deleteAfterExport,
			rescanInterval,
			config,
		)
		if err != nil {
			return err.CliError()
		}

		cliutil.PrintToConsole(source, c.Bool("json"), nil)
		return nil
	},
}
