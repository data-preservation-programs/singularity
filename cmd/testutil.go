//nolint:forcetypeassert
package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"slices"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/pack"
	"github.com/fatih/color"
	commp "github.com/filecoin-project/go-fil-commp-hashhash"
	"github.com/mattn/go-shellwords"
	"github.com/parnurzeal/gorequest"
	"github.com/rjNemo/underscore"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
)

type RunnerMode string

const (
	Normal  RunnerMode = "normal"
	Verbose RunnerMode = "verbose"
	JSON    RunnerMode = "json"
)

type Runner struct {
	sb   strings.Builder
	mode RunnerMode
}

var colorMutex = sync.Mutex{}

// NewRunner creates a new Runner to capture CLI args
//
// Note: tests invoking this function should stay in cmd.Test package
// because this function relies on environment variables to set database connection string
// so it won't work with parallel execution of different test packages.
func NewRunner() *Runner {
	colorMutex.Lock()
	defer colorMutex.Unlock()
	if color.NoColor {
		color.NoColor = false
	}
	return &Runner{}
}

func (r *Runner) WithMode(mode RunnerMode) *Runner {
	r.mode = mode
	return r
}

func (r *Runner) Run(ctx context.Context, args string) (string, string, error) {
	if strings.HasPrefix(args, "singularity ") {
		switch r.mode {
		case Verbose:
			args = "singularity --verbose " + args[len("singularity "):]
		case JSON:
			args = "singularity --json " + args[len("singularity "):]
		}
	}
	out, stderr, err := runWithCapture(ctx, args)
	green := color.New(color.FgGreen).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	r.sb.WriteString(green("user@localhost") + ":" + blue("~/test") + "$ " + args + "\n")
	r.sb.WriteString(out)
	r.sb.WriteString(stderr)
	r.sb.WriteString("\n")
	return out, stderr, err
}

var removeANSI = regexp.MustCompile(`\x1B\[[0-?]*[ -/]*[@-~]`)
var timeRegex = regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`)

func (r *Runner) Save(t *testing.T, tempDirs ...string) {
	t.Helper()
	t.Helper()
	ansi := r.sb.String()
	if ansi == "" {
		return
	}

	for i, tempDir := range tempDirs {
		ansi = strings.ReplaceAll(ansi, tempDir, "/tempDir/"+strconv.Itoa(i))
	}

	ansi = timeRegex.ReplaceAllString(ansi, "2023-04-05 06:07:08")

	ansiPath := filepath.Join("testdata", t.Name()+".ansi")
	err := os.MkdirAll(filepath.Dir(ansiPath), 0700)
	require.NoError(t, err)
	err = os.WriteFile(ansiPath, []byte(ansi), 0600)
	require.NoError(t, err)

	plain := removeANSI.ReplaceAllString(ansi, "")
	plainPath := filepath.Join("testdata", t.Name()+".txt")
	err = os.WriteFile(plainPath, []byte(plain), 0600)
	require.NoError(t, err)
}

func runWithCapture(ctx context.Context, args string) (string, string, error) {
	// Create a clone of the app so that we can runWithCapture from different tests concurrently
	app := *App
	for i, flag := range app.Flags {
		if flag.Names()[0] == "database-connection-string" {
			app.Flags[i] = &cli.StringFlag{
				Name:        "database-connection-string",
				Usage:       "Connection string to the database",
				DefaultText: "sqlite:" + "./singularity.db",
				Value:       "sqlite:" + "./singularity.db",
				EnvVars:     []string{"DATABASE_CONNECTION_STRING"},
			}
		}
	}
	app.ExitErrHandler = func(c *cli.Context, err error) {}
	parser := shellwords.NewParser()
	parser.ParseEnv = true // Enable environment variable parsing
	parsedArgs, err := parser.Parse(args)
	if err != nil {
		return "", "", errors.WithStack(err)
	}

	outWriter := bytes.NewBuffer(nil)
	errWriter := bytes.NewBuffer(nil)

	// Overwrite the stdout and stderr
	app.Writer = outWriter
	app.ErrWriter = errWriter

	err = app.RunContext(ctx, parsedArgs)
	return outWriter.String(), errWriter.String(), err
}

var pieceCIDRegex = regexp.MustCompile("baga6ea[0-9a-z]+")

func GetAllPieceCIDs(content string) []string {
	found := pieceCIDRegex.FindAllString(content, -1)
	found = underscore.Unique(found)
	slices.Sort(found)
	return found
}

var cidRegex = regexp.MustCompile("bafy[0-9a-z]+")

func GetFirstCID(content string) string {
	return cidRegex.FindString(content)
}

func CalculateCommp(t *testing.T, content []byte, targetPieceSize uint64) string {
	t.Helper()
	calc := &commp.Calc{}
	_, err := bytes.NewBuffer(content).WriteTo(calc)
	require.NoError(t, err)
	c, _, err := pack.GetCommp(calc, targetPieceSize)
	require.NoError(t, err)
	return c.String()
}

func WaitForServerReady(ctx context.Context, url string) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	var timer *time.Timer
	for {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		resp, _, err := gorequest.New().Timeout(time.Second).Get(url).End()
		if err == nil && resp != nil && resp.StatusCode == http.StatusOK {
			return nil
		}
		if timer == nil {
			timer = time.NewTimer(100 * time.Millisecond)
			defer timer.Stop()
		} else {
			timer.Reset(100 * time.Millisecond)
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-timer.C:
		}
	}
}

func Download(ctx context.Context, url string, nThreads int) ([]byte, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// Make a HEAD request to get the size of the file
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close()

	// Get the Content-Length header
	contentLength := resp.ContentLength
	if contentLength < 0 {
		return nil, errors.New("Content-Length header not found")
	}

	// Calculate size of each part
	partSize := contentLength / int64(nThreads)
	var extraSize int64 = 0
	if contentLength%int64(nThreads) != 0 {
		extraSize = contentLength % int64(nThreads)
	}

	// Download each part concurrently
	var wg sync.WaitGroup
	parts := make([][]byte, nThreads)
	errChan := make(chan error, nThreads)
	for i := range nThreads {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			start := int64(i) * partSize
			end := start + partSize - 1
			if i == nThreads-1 {
				end += extraSize // add the remainder to the last part
			}

			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				errChan <- errors.WithStack(err)
				return
			}

			// Set the Range header to Download a pack job
			req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end))
			resp, err := client.Do(req)
			if err != nil {
				errChan <- errors.WithStack(err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode < 200 || resp.StatusCode >= 300 {
				errChan <- errors.Newf("unexpected status code %d", resp.StatusCode)
				return
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				errChan <- errors.WithStack(err)
				return
			}

			// Save the part to the slice
			parts[i] = body
		}(i)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case err = <-errChan:
		return nil, errors.WithStack(err)
	case <-done:
	}

	// Combine the parts
	var result bytes.Buffer
	for _, part := range parts {
		result.Write(part)
	}

	return result.Bytes(), nil
}

func CompareDirectories(t *testing.T, dir1, dir2 string) {
	t.Helper()
	filesInDir2 := make(map[string]struct{})

	err := filepath.Walk(dir1, func(path1 string, info1 os.FileInfo, err error) error {
		// Propagate any error
		require.NoError(t, err)

		// Construct the path to the corresponding file or directory in dir2
		relPath := strings.TrimPrefix(path1, dir1)
		path2 := filepath.Join(dir2, relPath)

		// Get info about the file or directory in dir2
		info2, err := os.Stat(path2)
		if os.IsNotExist(err) {
			require.Failf(t, "Missing file or directory in dir2", "File: %s", relPath)
			return nil
		}
		require.NoError(t, err)

		if !info1.IsDir() {
			filesInDir2[relPath] = struct{}{}
		}

		// If both are directories, no need to compare content
		if info1.IsDir() && info2.IsDir() {
			return nil
		}

		// Compare file sizes
		require.Equal(t, info1.Size(), info2.Size(), "Size mismatch for %s", relPath)

		// Compare file content
		content1, err := os.ReadFile(path1)
		require.NoError(t, err)

		content2, err := os.ReadFile(path2)
		require.NoError(t, err)

		require.True(t, bytes.Equal(content1, content2), "Content mismatch for %s", relPath)

		return nil
	})

	require.NoError(t, err)

	err = filepath.Walk(dir2, func(path2 string, info2 os.FileInfo, err error) error {
		// Propagate any error
		require.NoError(t, err)

		relPath := strings.TrimPrefix(path2, dir2)

		// If we've already checked this file (because it exists in dir1), then skip it
		if _, ok := filesInDir2[relPath]; ok || info2.IsDir() {
			return nil
		}

		// If we get here, it means this file/dir exists in dir2 but not in dir1
		require.Failf(t, "Extra file or directory in dir2", "File: %s", relPath)
		return nil
	})

	require.NoError(t, err)
}
