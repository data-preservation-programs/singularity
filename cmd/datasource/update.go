package datasource

import (
	"context"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/rjNemo/underscore"
	"github.com/urfave/cli/v2"
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
				return stringFlag
			})
			flags = append(flags, cmdFlags...)
		}
		return flags
	}(),
	Action: func(c *cli.Context) error {
		db := database.MustOpenFromCLI(c)
		config := map[string]string{}
		for _, name := range c.LocalFlagNames() {
			if c.IsSet(name) {
				value := c.String(name)
				config[name] = value
			}
		}

		source, err := datasource.UpdateSourceHandler(
			db,
			context.Background(),
			c.Args().Get(0),
			config,
		)
		if err != nil {
			return err.CliError()
		}

		cliutil.PrintToConsole(source, c.Bool("json"))
		return nil
	},
}
