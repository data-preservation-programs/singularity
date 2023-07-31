package cmd

import (
	"fmt"
	"runtime/debug"

	"github.com/urfave/cli/v2"
)

var VersionCmd = &cli.Command{
	Name:    "version",
	Usage:   "Print version information",
	Aliases: []string{"v"},
	Action: func(context *cli.Context) error {
		buildInfo, ok := debug.ReadBuildInfo()
		if !ok {
			fmt.Println("unknown version")
		}

		version := buildInfo.Main.Version
		if version == "(devel)" {
			version = Version
		}
		var revision string
		var modified string
		for _, setting := range buildInfo.Settings {
			switch setting.Key {
			case "vcs.revision":
				revision = setting.Value[:7]
			case "vcs.modified":
				modified = setting.Value
			}
		}
		if revision == "" {
			revision = "-unknown"
		} else {
			revision = "-" + revision
		}
		switch modified {
		case "true":
			modified = "-dirty"
		case "false":
			modified = ""
		case "":
			modified = "-unknown"
		default:
			modified = "-" + modified
		}
		fmt.Printf("singularity %s%s%s\n", version, revision, modified)
		return nil
	},
}
