package cmd

import (
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
	"os"
)

func ReduceCLILogLevel(context *cli.Context) {
	current := logging.GetConfig().Level
	target := logging.LevelError
	if os.Getenv("GOLOG_LOG_LEVEL") == "" {
		target = logging.LevelInfo
	}

	if context.Bool("verbose") {
		target = logging.LevelDebug
	}

	if target < current {
		logging.SetAllLoggers(target)
	}
}
