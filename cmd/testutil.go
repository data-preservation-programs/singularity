//nolint:forcetypeassert
package cmd

import (
	"bytes"
	"context"
	"fmt"
	color2 "image/color"
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

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/data-preservation-programs/singularity/handler/deal"
	"github.com/data-preservation-programs/singularity/handler/deal/schedule"
	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack"
	"github.com/data-preservation-programs/singularity/replication"
	"github.com/fatih/color"
	commp "github.com/filecoin-project/go-fil-commp-hashhash"
	"github.com/jiro4989/textimg/v3/config"
	"github.com/jiro4989/textimg/v3/image"
	"github.com/jiro4989/textimg/v3/parser"
	"github.com/mattn/go-shellwords"
	"github.com/parnurzeal/gorequest"
	"github.com/rjNemo/underscore"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
	"github.com/ybbus/jsonrpc/v3"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
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

func (r *Runner) Run(ctx context.Context, args string) (string, string, error) {
	color.NoColor = false
	if strings.HasPrefix(args, "singularity ") {
		switch r.mode {
		case Verbose:
			args = "singularity --verbose " + args[len("singularity "):]
		case JSON:
			args = "singularity --json " + args[len("singularity "):]
		}
	}
	out, stderr, err := Run(ctx, args)
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
	ansi := r.sb.String()

	for i, tempDir := range tempDirs {
		ansi = strings.ReplaceAll(ansi, tempDir, "/tempDir/"+strconv.Itoa(i))
	}

	ansi = timeRegex.ReplaceAllString(ansi, "2023-04-05 06:07:08")

	plain := removeANSI.ReplaceAllString(ansi, "")
	plainPath := filepath.Join("testdata", t.Name()+".txt")
	err := os.MkdirAll(filepath.Dir(plainPath), 0700)
	require.NoError(t, err)
	err = os.WriteFile(plainPath, []byte(plain), 0600)
	require.NoError(t, err)

	imagePath := filepath.Join("testdata", t.Name()+".png")

	c := config.Config{
		Foreground:    "255,255,255,255",
		Background:    "40,20,20,255",
		Outpath:       imagePath,
		FontSize:      20,
		LineCount:     1,
		SlideWidth:    1,
		FileExtension: ".png",
	}

	err = c.Adjust([]string{ansi}, config.EnvVars{})
	require.NoError(t, err)

	tokens, err := parser.Parse(strings.Join(c.Texts, "\n"))
	require.NoError(t, err)

	bw := tokens.MaxStringWidth()
	bh := len(tokens.StringLines())
	param := &image.ImageParam{
		BaseWidth:          bw,
		BaseHeight:         bh,
		ForegroundColor:    color2.RGBA(c.ForegroundColor),
		BackgroundColor:    color2.RGBA(c.BackgroundColor),
		FontFace:           c.FontFace,
		EmojiFontFace:      c.EmojiFontFace,
		EmojiDir:           c.EmojiDir,
		FontSize:           c.FontSize,
		Delay:              c.Delay,
		UseAnimation:       c.UseAnimation,
		AnimationLineCount: c.LineCount,
		ResizeWidth:        c.ResizeWidth,
		ResizeHeight:       c.ResizeHeight,
		UseEmoji:           c.UseEmojiFont,
	}
	img := image.NewImage(param)
	err = img.Draw(tokens)
	require.NoError(t, err)
	err = img.Draw(tokens)
	require.NoError(t, err)
	err = img.Encode(c.Writer, c.FileExtension)
	require.NoError(t, err)
}

func Run(ctx context.Context, args string) (string, string, error) {
	// Create a clone of the app so that we can run from different tests concurrently
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

type MockAdmin struct {
	mock.Mock
}

func (m *MockAdmin) InitHandler(ctx context.Context, db *gorm.DB) error {
	args := m.Called(ctx, db)
	return args.Error(0)
}

func (m *MockAdmin) ResetHandler(ctx context.Context, db *gorm.DB) error {
	args := m.Called(ctx, db)
	return args.Error(0)
}

type MockDataPrep struct {
	mock.Mock
}

func (m *MockDataPrep) CreatePreparationHandler(ctx context.Context, db *gorm.DB, request dataprep.CreateRequest) (*model.Preparation, error) {
	args := m.Called(ctx, db, request)
	return args.Get(0).(*model.Preparation), args.Error(1)
}

func (m *MockDataPrep) StartDagGenHandler(ctx context.Context, db *gorm.DB, id uint32, name string) (*model.Job, error) {
	args := m.Called(ctx, db, id, name)
	return args.Get(0).(*model.Job), args.Error(1)
}

func (m *MockDataPrep) PauseDagGenHandler(ctx context.Context, db *gorm.DB, id uint32, name string) (*model.Job, error) {
	args := m.Called(ctx, db, id, name)
	return args.Get(0).(*model.Job), args.Error(1)
}

func (m *MockDataPrep) ExploreHandler(ctx context.Context, db *gorm.DB, id uint32, name string, path string) (*dataprep.ExploreResult, error) {
	args := m.Called(ctx, db, id, name, path)
	return args.Get(0).(*dataprep.ExploreResult), args.Error(1)
}

func (m *MockDataPrep) ListHandler(ctx context.Context, db *gorm.DB) ([]model.Preparation, error) {
	args := m.Called(ctx, db)
	return args.Get(0).([]model.Preparation), args.Error(1)
}

func (m *MockDataPrep) AddOutputStorageHandler(ctx context.Context, db *gorm.DB, id uint32, output string) (*model.Preparation, error) {
	args := m.Called(ctx, db, id, output)
	return args.Get(0).(*model.Preparation), args.Error(1)
}

func (m *MockDataPrep) RemoveOutputStorageHandler(ctx context.Context, db *gorm.DB, id uint32, output string) (*model.Preparation, error) {
	args := m.Called(ctx, db, id, output)
	return args.Get(0).(*model.Preparation), args.Error(1)
}

func (m *MockDataPrep) StartPackHandler(ctx context.Context, db *gorm.DB, id uint32, name string, jobID int64) ([]model.Job, error) {
	args := m.Called(ctx, db, id, name, jobID)
	return args.Get(0).([]model.Job), args.Error(1)
}

func (m *MockDataPrep) PausePackHandler(ctx context.Context, db *gorm.DB, id uint32, name string, jobID int64) ([]model.Job, error) {
	args := m.Called(ctx, db, id, name, jobID)
	return args.Get(0).([]model.Job), args.Error(1)
}

func (m *MockDataPrep) ListPiecesHandler(ctx context.Context, db *gorm.DB, id uint32) ([]dataprep.PieceList, error) {
	args := m.Called(ctx, db, id)
	return args.Get(0).([]dataprep.PieceList), args.Error(1)
}

func (m *MockDataPrep) AddPieceHandler(ctx context.Context, db *gorm.DB, id uint32, request dataprep.AddPieceRequest) (*model.Car, error) {
	args := m.Called(ctx, db, id, request)
	return args.Get(0).(*model.Car), args.Error(1)
}

func (m *MockDataPrep) StartScanHandler(ctx context.Context, db *gorm.DB, id uint32, name string) (*model.Job, error) {
	args := m.Called(ctx, db, id, name)
	return args.Get(0).(*model.Job), args.Error(1)
}

func (m *MockDataPrep) PauseScanHandler(ctx context.Context, db *gorm.DB, id uint32, name string) (*model.Job, error) {
	args := m.Called(ctx, db, id, name)
	return args.Get(0).(*model.Job), args.Error(1)
}

func (m *MockDataPrep) AddSourceStorageHandler(ctx context.Context, db *gorm.DB, id uint32, source string) (*model.Preparation, error) {
	args := m.Called(ctx, db, id, source)
	return args.Get(0).(*model.Preparation), args.Error(1)
}

func (m *MockDataPrep) GetStatusHandler(ctx context.Context, db *gorm.DB, id uint32) ([]dataprep.SourceStatus, error) {
	args := m.Called(ctx, db, id)
	return args.Get(0).([]dataprep.SourceStatus), args.Error(1)
}

type MockSchedule struct {
	mock.Mock
}

func (m *MockSchedule) CreateHandler(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient, request schedule.CreateRequest) (*model.Schedule, error) {
	args := m.Called(ctx, db, lotusClient, request)
	return args.Get(0).(*model.Schedule), args.Error(1)
}

func (m *MockSchedule) ListHandler(ctx context.Context, db *gorm.DB) ([]model.Schedule, error) {
	args := m.Called(ctx, db)
	return args.Get(0).([]model.Schedule), args.Error(1)
}

func (m *MockSchedule) PauseHandler(ctx context.Context, db *gorm.DB, scheduleID uint32) (*model.Schedule, error) {
	args := m.Called(ctx, db, scheduleID)
	return args.Get(0).(*model.Schedule), args.Error(1)
}

func (m *MockSchedule) ResumeHandler(ctx context.Context, db *gorm.DB, scheduleID uint32) (*model.Schedule, error) {
	args := m.Called(ctx, db, scheduleID)
	return args.Get(0).(*model.Schedule), args.Error(1)
}

type MockDeal struct {
	mock.Mock
}

func (m *MockDeal) ListHandler(ctx context.Context, db *gorm.DB, request deal.ListDealRequest) ([]model.Deal, error) {
	args := m.Called(ctx, db, request)
	return args.Get(0).([]model.Deal), args.Error(1)
}

func (m *MockDeal) SendManualHandler(ctx context.Context, db *gorm.DB, dealMaker replication.DealMaker, request deal.Proposal) (*model.Deal, error) {
	args := m.Called(ctx, db, dealMaker, request)
	return args.Get(0).(*model.Deal), args.Error(1)
}

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) CreateStorageHandler(ctx context.Context, db *gorm.DB, storageType string, request storage.CreateRequest) (*model.Storage, error) {
	args := m.Called(ctx, db, storageType, request)
	return args.Get(0).(*model.Storage), args.Error(1)
}

func (m *MockStorage) ExploreHandler(ctx context.Context, db *gorm.DB, name string, path string) ([]storage.DirEntry, error) {
	args := m.Called(ctx, db, name, path)
	return args.Get(0).([]storage.DirEntry), args.Error(1)
}

func (m *MockStorage) ListStoragesHandler(ctx context.Context, db *gorm.DB) ([]model.Storage, error) {
	args := m.Called(ctx, db)
	return args.Get(0).([]model.Storage), args.Error(1)
}

func (m *MockStorage) RemoveHandler(ctx context.Context, db *gorm.DB, name string) error {
	args := m.Called(ctx, db, name)
	return args.Error(0)
}

func (m *MockStorage) UpdateStorageHandler(ctx context.Context, db *gorm.DB, name string, config map[string]string) (*model.Storage, error) {
	args := m.Called(ctx, db, name, config)
	return args.Get(0).(*model.Storage), args.Error(1)
}

type MockWallet struct {
	mock.Mock
}

func (m *MockWallet) AttachHandler(ctx context.Context, db *gorm.DB, preparationID uint32, wallet string) (*model.Preparation, error) {
	args := m.Called(ctx, db, preparationID, wallet)
	return args.Get(0).(*model.Preparation), args.Error(1)
}

func (m *MockWallet) DetachHandler(ctx context.Context, db *gorm.DB, preparationID uint32, wallet string) (*model.Preparation, error) {
	args := m.Called(ctx, db, preparationID, wallet)
	return args.Get(0).(*model.Preparation), args.Error(1)
}

func (m *MockWallet) ImportHandler(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient, request wallet.ImportRequest) (*model.Wallet, error) {
	args := m.Called(ctx, db, lotusClient, request)
	return args.Get(0).(*model.Wallet), args.Error(1)
}

func (m *MockWallet) ListHandler(ctx context.Context, db *gorm.DB) ([]model.Wallet, error) {
	args := m.Called(ctx, db)
	return args.Get(0).([]model.Wallet), args.Error(1)
}

func (m *MockWallet) ListAttachedHandler(ctx context.Context, db *gorm.DB, preparationID uint32) ([]model.Wallet, error) {
	args := m.Called(ctx, db, preparationID)
	return args.Get(0).([]model.Wallet), args.Error(1)
}

func (m *MockWallet) RemoveHandler(ctx context.Context, db *gorm.DB, address string) error {
	args := m.Called(ctx, db, address)
	return args.Error(0)
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
	for {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		resp, _, err := gorequest.New().Timeout(time.Second).Get(url).End()
		if err == nil && resp != nil && resp.StatusCode == http.StatusOK {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(100 * time.Millisecond):
		}
	}
}

func Download(ctx context.Context, url string, nThreads int) ([]byte, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// Make a HEAD request to get the size of the file
	req, err := http.NewRequestWithContext(ctx, "HEAD", url, nil)
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
	for i := 0; i < nThreads; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			start := int64(i) * partSize
			end := start + partSize - 1
			if i == nThreads-1 {
				end += extraSize // add the remainder to the last part
			}

			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
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
