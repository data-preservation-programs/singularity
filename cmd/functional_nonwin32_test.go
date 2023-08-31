//go:build !(windows && 386)

package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
)

// SQLite is not supported on Windows 32-bit
func TestEzPrep(t *testing.T) {
	source := t.TempDir()
	sizes := []int{0, 1, 1 << 20, 10 << 20, 30 << 20}
	for _, size := range sizes {
		err := os.WriteFile(filepath.Join(source, fmt.Sprintf("size-%d.txt", size)), testutil.GenerateFixedBytes(size), 0777)
		require.NoError(t, err)
	}

	for _, maxSize := range []string{"3MB", "30GB"} {
		t.Run("maxSize_"+maxSize, func(t *testing.T) {
			for _, inmemory := range []bool{true, false} {
				t.Run(fmt.Sprintf("inmemory_%t", inmemory), func(t *testing.T) {
					for _, concurrency := range []int{1, 4} {
						if inmemory && concurrency > 1 {
							continue
						}
						t.Run(fmt.Sprintf("concurrency_%d", concurrency), func(t *testing.T) {
							for _, inline := range []bool{true, false} {
								t.Run(fmt.Sprintf("inline_%t", inline), func(t *testing.T) {
									output := t.TempDir()
									runner := Runner{mode: Normal}
									defer runner.Save(t, source, output)

									databaseFile := ""
									if !inmemory {
										databaseFile = filepath.Join(output, "test.db")
									}

									outFlag := ""
									if !inline {
										outFlag = fmt.Sprintf("-o %s", testutil.EscapePath(output))
									}

									_, _, err := runner.Run(context.Background(),
										fmt.Sprintf("singularity ez-prep %s -M %s -j %d -f %s %s",
											outFlag, maxSize, concurrency, testutil.EscapePath(databaseFile), testutil.EscapePath(source)))
									require.NoError(t, err)
								})
							}
						})
					}
				})
			}
		})
	}
}
