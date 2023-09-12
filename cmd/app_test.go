package cmd

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/analytics"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

func init() {
	analytics.Enabled = false
}

func TestHelpPage(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		SetupHelpPager()
		var testCommand func(ctx context.Context, t *testing.T, db *gorm.DB, cmd string, cliCommand *cli.Command)
		testCommand = func(ctx context.Context, t *testing.T, db *gorm.DB, cmd string, cliCommand *cli.Command) {
			cmd = cmd + " " + cliCommand.Name
			t.Run(cmd, func(t *testing.T) {
				_, _, err := NewRunner().Run(ctx, cmd+" -h")
				require.NoError(t, err)
			})
			for _, subcommand := range cliCommand.Subcommands {
				if subcommand.Name == "help" {
					continue
				}
				testCommand(ctx, t, db, cmd, subcommand)
			}
		}

		cmd := "singularity"
		t.Run(cmd, func(t *testing.T) {
			_, _, err := NewRunner().Run(ctx, cmd+" -h")
			require.NoError(t, err)
		})
		for _, subcommand := range App.Commands {
			testCommand(ctx, t, db, cmd, subcommand)
		}
	})
}

func TestVersion(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := SetVersionJSON([]byte(`{"version":"v0.0.1"}`))
		require.NoError(t, err)
		out, _, err := NewRunner().Run(ctx, "singularity version")
		require.NoError(t, err)
		require.Contains(t, out, "singularity v0.0.1-unknown-unknown")
	})
}
