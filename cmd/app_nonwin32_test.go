//go:build !(windows && 386)

package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/stretchr/testify/require"
)

func TestEzPrepBenchmark(t *testing.T) {
	temp := t.TempDir()
	err := os.WriteFile(filepath.Join(temp, "test.img"), []byte("hello world"), 0777)
	require.NoError(t, err)
	ctx := context.Background()
	out, _, err := RunArgsInTest(ctx, fmt.Sprintf("singularity --database-connection-string '%s' ez-prep --output-dir '' --database-file '' -j 1 %s", database.TestConnectionString, escapePath(temp)))
	require.NoError(t, err)
	// contains two CARs, one for the file and another one for the dag
	require.Contains(t, out, "107")
	require.Contains(t, out, "152")
}
