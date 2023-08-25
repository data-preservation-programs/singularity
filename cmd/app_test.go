package cmd

import (
	"bufio"
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/deal/schedule"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"filippo.io/age"
	"github.com/data-preservation-programs/singularity/pack"
	commp "github.com/filecoin-project/go-fil-commp-hashhash"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-blockservice"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	format "github.com/ipfs/go-ipld-format"
	"github.com/ipfs/go-merkledag"
	uio "github.com/ipfs/go-unixfs/io"
	"github.com/ipld/go-car"
	"github.com/ipld/go-car/util"
	"github.com/joho/godotenv"
	"github.com/parnurzeal/gorequest"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
)

type testDB struct {
	db      *gorm.DB
	closer  io.Closer
	connStr string
}

var testDBMap map[string]testDB = make(map[string]testDB)
var mu sync.Mutex

func getTestDB(t *testing.T, dialect string) (db *gorm.DB, closer io.Closer, connStr string, err error) {
	if dialect == "sqlite" {
		connStr = "sqlite:" + t.TempDir() + "/singularity.db"
		db, closer, err = database.Open(connStr, &gorm.Config{})
		return
	}
	mu.Lock()
	defer mu.Unlock()
	found, ok := testDBMap[dialect]
	if ok {
		return found.db, found.closer, found.connStr, nil
	}
	var opError *net.OpError
	switch dialect {
	case "mysql":
		connStr = "mysql://singularity:singularity@tcp(localhost:3306)/singularity?parseTime=true"
		db, closer, err = database.Open(connStr, &gorm.Config{})
		if errors.As(err, &opError) {
			db = nil
			err = nil
		}
	case "postgres":
		connStr = "postgres://singularity:singularity@localhost:5432/singularity?sslmode=disable"
		db, closer, err = database.Open(connStr, &gorm.Config{})
		if errors.As(err, &opError) {
			db = nil
			err = nil
		}
	default:
		err = errors.New("Unsupported dialect")
	}
	if err != nil {
		return
	}
	testDBMap[dialect] = testDB{
		db, closer, connStr,
	}
	return
}

func generateRandomBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}

func decrypt(t *testing.T, key string, encrypted []byte) []byte {
	recipient, err := age.ParseX25519Identity(key)
	require.NoError(t, err)
	decrypted, err := age.Decrypt(bytes.NewReader(encrypted), recipient)
	require.NoError(t, err)
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, decrypted)
	require.NoError(t, err)
	return buf.Bytes()
}

func loadCars(t *testing.T, path string) blockstore.Blockstore {
	files, err := os.ReadDir(path)
	require.NoError(t, err)
	bs := blockstore.NewBlockstore(datastore.NewMapDatastore())
	for _, file := range files {
		f, err := os.Open(filepath.Join(path, file.Name()))
		require.NoError(t, err)
		defer f.Close()
		reader := bufio.NewReader(f)
		_, err = car.ReadHeader(reader)
		require.NoError(t, err)
		for {
			c, data, err := util.ReadNode(reader)
			if err == io.EOF {
				break
			}
			blk, _ := blocks.NewBlockWithCid(data, c)
			err = bs.Put(context.TODO(), blk)
			require.NoError(t, err)
		}
	}
	return bs
}

func getFileFromRootNode(t *testing.T, dagServ format.DAGService, path string, rootCID cid.Cid) []byte {
	ctx := context.TODO()
	segments := strings.Split(path, "/")
	for _, segment := range segments {
		rootNode, err := dagServ.Get(context.Background(), rootCID)
		require.NoError(t, err)
		rootDir, err := uio.NewDirectoryFromNode(dagServ, rootNode)
		require.NoError(t, err)
		links, err := rootDir.Links(ctx)
		require.NoError(t, err)
		link, err := underscore.Find(links, func(link *format.Link) bool {
			return link.Name == segment
		})
		require.NoError(t, err)
		rootCID = link.Cid
	}
	fileNode, err := dagServ.Get(ctx, rootCID)
	require.NoError(t, err)
	dagReader, err := uio.NewDagReader(ctx, fileNode, dagServ)
	require.NoError(t, err)
	content, err := io.ReadAll(dagReader)
	require.NoError(t, err)
	return content
}

func escapePath(p string) string {
	return "'" + strings.ReplaceAll(p, `\`, `\\`) + "'"
}

func TestExtractCar(t *testing.T) {
	testWithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tempDir := t.TempDir()
		carDir := t.TempDir()
		extractDir := t.TempDir()
		err := os.WriteFile(filepath.Join(tempDir, "test1.txt"), []byte("hello"), 0644)
		require.NoError(t, err)
		err = os.MkdirAll(filepath.Join(tempDir, "test"), 0755)
		require.NoError(t, err)
		err = os.WriteFile(filepath.Join(tempDir, "test", "test2.txt"), []byte("world"), 0644)
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity dataset create -o "+escapePath(carDir)+" test")
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity datasource add local test "+escapePath(tempDir))
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity datasource daggen 1")
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		out, _, err := RunArgsInTest(ctx, "singularity --json datasource inspect path 1")
		require.NoError(t, err)
		root := strings.Split(strings.Split(out, "\n")[4], "\"")[3]
		_, _, err = RunArgsInTest(ctx, "singularity tool extract-car -i "+escapePath(carDir)+" -o "+escapePath(extractDir)+" -c "+root)
		require.NoError(t, err)
		content, err := os.ReadFile(filepath.Join(extractDir, "test1.txt"))
		require.NoError(t, err)
		require.Equal(t, "hello", string(content))
		content, err = os.ReadFile(filepath.Join(extractDir, "test", "test2.txt"))
		require.NoError(t, err)
		require.Equal(t, "world", string(content))
	})
}

func testWithAllBackendWithoutReset(t *testing.T, testFunc func(ctx context.Context, t *testing.T, db *gorm.DB)) {
	testWithAllBackendWithResetArg(t, testFunc, false)
}

func testWithAllBackendWithResetArg(t *testing.T, testFunc func(ctx context.Context, t *testing.T, db *gorm.DB), reset bool) {
	for _, backend := range database.SupportedTestDialects {
		db, closer, connStr, err := getTestDB(t, backend)
		if err != nil && strings.Contains(err.Error(), "Unsupported dialect") {
			t.Log("Skip " + backend)
			continue
		}
		require.NoError(t, err)
		if db == nil {
			t.Log("Skip " + backend)
			continue
		}
		defer closer.Close()
		os.Setenv("DATABASE_CONNECTION_STRING", connStr)
		if reset {
			_, _, err := RunArgsInTest(context.Background(), "singularity admin reset --really-do-it")
			require.NoError(t, err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
		defer cancel()
		t.Run(backend, func(t *testing.T) {
			testFunc(ctx, t, db)
		})
	}
}

func testWithAllBackend(t *testing.T, testFunc func(ctx context.Context, t *testing.T, db *gorm.DB)) {
	testWithAllBackendWithResetArg(t, testFunc, true)
}

func TestHelpPage(t *testing.T) {
	testWithAllBackendWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, _, err := RunArgsInTest(ctx, "singularity help")
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity download -h")
		require.NoError(t, err)
	})
}

func TestDealTracker(t *testing.T) {
	testWithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ctx2, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		_, _, err := RunArgsInTest(ctx2, "singularity run deal-tracker")
		require.ErrorContains(t, err, "context deadline exceeded")
	})
}

func TestRunAPI(t *testing.T) {
	testWithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ctx2, cancel := context.WithCancel(ctx)
		serverClosed := make(chan struct{})
		defer func() {
			<-serverClosed
		}()
		defer cancel()
		go func() {
			err := RunArgsInTestNoCapture(ctx2, "singularity run api")
			require.ErrorContains(t, err, "Server closed")
			close(serverClosed)
		}()
		var resp *http.Response
		var body string
		var errs []error
		// try every 100ms for up to 5 seconds for server to come up
		for i := 0; i < 50; i++ {
			time.Sleep(100 * time.Millisecond)
			resp, body, errs = gorequest.New().
				Get("http://127.0.0.1:9090/robots.txt").End()
			if resp != nil && resp.StatusCode == http.StatusOK {
				break
			}
		}
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, "robotstxt.org")
		resp, body, errs = gorequest.New().
			Post("http://127.0.0.1:9090/api/admin/init").End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusNoContent, resp.StatusCode)
		require.Equal(t, "", body)

		resp, body, errs = gorequest.New().
			Post("http://127.0.0.1:9090/api/admin/reset").End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusNoContent, resp.StatusCode)
		require.Equal(t, "", body)

		resp, body, errs = gorequest.New().Post("http://127.0.0.1:9090/api/dataset").
			Send(`{"name":"test","maxSize":"31.5GiB"}`).End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, `"name": "test"`)

		defer func() {
			resp, body, errs = gorequest.New().Delete("http://127.0.0.1:9090/api/dataset/test").End()
			require.Len(t, errs, 0)
			require.Equal(t, http.StatusNoContent, resp.StatusCode)
			require.Equal(t, "", body)
		}()

		resp, body, errs = gorequest.New().Post("http://127.0.0.1:9090/api/dataset").
			Send(`{"name":"test"}`).End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusBadRequest, resp.StatusCode)
		require.Contains(t, body, `"err":`)

		resp, body, errs = gorequest.New().Patch("http://127.0.0.1:9090/api/dataset/test").
			Send(`{"maxSize":"30.5GiB"}`).End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, `"name": "test"`)

		resp, body, errs = gorequest.New().Get("http://127.0.0.1:9090/api/dataset").End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, `"name": "test"`)

		resp, body, errs = gorequest.New().Post("http://127.0.0.1:9090/api/dataset/test/piece").
			Send(`{"pieceCid":"baga6ea4seaqdyupo27fj2fk2mtefzlxvrbf6kdi4twdpccdzbyqrbpsvfsh5ula","pieceSize":"1024","rootCid":"bafy2bzacecq55ww767qv2r3cvlorjxhcvn3dglccajiicsqgctpm7qfrtncw4"}`).End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, `"pieceSize": 1024`)

		resp, body, errs = gorequest.New().Get("http://127.0.0.1:9090/api/dataset/test/piece").End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, `"pieceSize": 1024`)

		key := "7b2254797065223a22736563703235366b31222c22507269766174654b6579223a226b35507976337148327349586343595a58594f5775453149326e32554539436861556b6c4e36695a5763453d227d"
		resp, body, errs = gorequest.New().Post("http://127.0.0.1:9090/api/wallet").
			Send(`{"privateKey":"` + key + `"}`).End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, key)

		resp, body, errs = gorequest.New().Post("http://127.0.0.1:9090/api/wallet/remote").
			Send(`{"remotePeer":"12D3KooWD3eckifWpRn9wQpMG9R9hX3sD158z7EqHWmweQAJU5SA","address":"f1ys5qqiciehcml3sp764ymbbytfn3qoar5fo3iwy"}`).End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, "12D3KooWD3eckifWpRn9wQpMG9R9hX3sD158z7EqHWmweQAJU5SA")

		resp, body, errs = gorequest.New().Delete("http://127.0.0.1:9090/api/wallet/f1ys5qqiciehcml3sp764ymbbytfn3qoar5fo3iwy").End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusNoContent, resp.StatusCode)
		require.Equal(t, ``, body)

		resp, body, errs = gorequest.New().Get("http://127.0.0.1:9090/api/wallet").End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, key)

		resp, body, errs = gorequest.New().Post("http://127.0.0.1:9090/api/dataset/test/wallet/f0808055").End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, `"datasetId": 1`)

		resp, body, errs = gorequest.New().Get("http://127.0.0.1:9090/api/dataset/test/wallet").End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, `"address": "`)

		defer func() {
			resp, body, errs = gorequest.New().Delete("http://127.0.0.1:9090/api/dataset/test/wallet/f0808055").End()
			require.Len(t, errs, 0)
			require.Equal(t, http.StatusNoContent, resp.StatusCode)
			require.Equal(t, ``, body)
		}()

		createSchedule := schedule.CreateRequest{
			DatasetName:        "test",
			Provider:           "f022352",
			StartDelay:         "24h",
			Duration:           "2400h",
			ScheduleCron:       "",
			ScheduleDealSize:   "1P",
			TotalDealSize:      "1P",
			MaxPendingDealSize: "1P",
		}
		createScheduleBody, err := json.Marshal(createSchedule)
		require.NoError(t, err)
		resp, body, errs = gorequest.New().Post("http://127.0.0.1:9090/api/schedule").Send(string(createScheduleBody)).End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, `"id": 1`)

		resp, body, errs = gorequest.New().Get("http://127.0.0.1:9090/api/schedule").End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, `"id": 1`)

		resp, body, errs = gorequest.New().Post("http://127.0.0.1:9090/api/schedule/1/pause").End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, `paused`)

		resp, body, errs = gorequest.New().Post("http://127.0.0.1:9090/api/schedule/1/resume").End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, `active`)

		tmp := t.TempDir()
		payload := map[string]any{
			"sourcePath":        tmp,
			"caseInsensitive":   "false",
			"deleteAfterExport": false,
			"rescanInterval":    "1h",
			"scanningState":     "ready",
		}
		requestBody, err := json.Marshal(payload)
		require.NoError(t, err)
		resp, body, errs = gorequest.New().Post("http://127.0.0.1:9090/api/source/local/dataset/test").
			Send(string(requestBody)).End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, `"id": 1`)

		resp, body, errs = gorequest.New().Get("http://127.0.0.1:9090/api/source").End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, `"id": 1`)

		resp, body, errs = gorequest.New().Patch("http://127.0.0.1:9090/api/source/1").
			Send(`{"deleteAfterExport": true, "rescanInterval": "12h", "localCaseInsensitive":"true"}`).End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, `"id": 1`)

		defer func() {
			resp, body, errs = gorequest.New().Delete("http://127.0.0.1:9090/api/source/1").End()
			require.Len(t, errs, 0)
			require.Equal(t, http.StatusNoContent, resp.StatusCode)
			require.Equal(t, ``, body)
		}()

		resp, body, errs = gorequest.New().Post("http://127.0.0.1:9090/api/source/1/rescan").End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, `"id": 1`)

		resp, body, errs = gorequest.New().Post("http://127.0.0.1:9090/api/source/1/check").
			Send(`{"path":""}`).End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, `[`)

		resp, body, errs = gorequest.New().Get("http://127.0.0.1:9090/api/source/1/summary").End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, `{`)

		resp, body, errs = gorequest.New().Get("http://127.0.0.1:9090/api/source/1/packjobs").End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, `[`)

		resp, body, errs = gorequest.New().Get("http://127.0.0.1:9090/api/source/1/files").End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, `[`)

		resp, body, errs = gorequest.New().Get("http://127.0.0.1:9090/api/source/1/dags").End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, `[`)

		resp, body, errs = gorequest.New().Get("http://127.0.0.1:9090/api/source/1/path").
			Send(`{"path":""}`).End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, `[`)

		resp, body, errs = gorequest.New().Get("http://127.0.0.1:9090/api/file/1").End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusBadRequest, resp.StatusCode)

		resp, body, errs = gorequest.New().Get("http://127.0.0.1:9090/api/packjob/1").End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusBadRequest, resp.StatusCode)

		resp, body, errs = gorequest.New().Post("http://127.0.0.1:9090/api/send_deal").Send(`{}`).End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusBadRequest, resp.StatusCode)
		require.Contains(t, body, `client address not found`)

		resp, body, errs = gorequest.New().Post("http://127.0.0.1:9090/api/deal").Send(`{}`).End()
		require.Len(t, errs, 0)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, body, `[`)
	})
}

func TestListDeals(t *testing.T) {
	testWithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, _, err := RunArgsInTest(ctx, "singularity deal list --dataset test --schedule 1 --provider f01 --state active")
		require.NoError(t, err)
	})
}

func TestResetDatabaseReallyDoIt(t *testing.T) {
	testWithAllBackendWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, _, err := RunArgsInTest(ctx, "singularity admin reset --really-do-it")
		require.NoError(t, err)
	})
}

func TestResetDatabase(t *testing.T) {
	testWithAllBackendWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, _, err := RunArgsInTest(ctx, "singularity admin reset")
		require.ErrorContains(t, err, "really-do-it")
	})
}

func TestInitDatabase(t *testing.T) {
	testWithAllBackendWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, _, err := RunArgsInTest(ctx, "singularity admin init")
		require.NoError(t, err)
	})
}

func TestDealScheduleCrud(t *testing.T) {
	testWithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, _, err := RunArgsInTest(ctx, "singularity dataset create test")
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity wallet add-remote f1l2cc5vuw5moppwsjd3b7cjtwa2exowqo36esklq 12D3KooWD3eckifWpRn9wQpMG9R9hX3sD158z7EqHWmweQAJU5SA")
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity dataset add-wallet test f02170643")
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity deal schedule create test f022352")
		require.NoError(t, err)
		body, _, err := RunArgsInTest(ctx, "singularity deal schedule list")
		require.NoError(t, err)
		require.Contains(t, body, "72h")
		body, _, err = RunArgsInTest(ctx, "singularity deal schedule pause 1")
		require.NoError(t, err)
		require.Contains(t, body, "paused")
		body, _, err = RunArgsInTest(ctx, "singularity deal schedule resume 1")
		require.NoError(t, err)
		require.Contains(t, body, "active")
	})
}

func TestWalletCrud(t *testing.T) {
	godotenv.Load("../.env", ".env")
	testWithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		key := "7b2254797065223a22736563703235366b31222c22507269766174654b6579223a226b35507976337148327349586343595a58594f5775453149326e32554539436861556b6c4e36695a5763453d227d"
		_, _, err := RunArgsInTest(ctx, "singularity dataset create test")
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity wallet import "+key)
		require.NoError(t, err)
		out, _, err := RunArgsInTest(ctx, "singularity wallet list ")
		require.NoError(t, err)
		require.Contains(t, out, "f0808055")
		_, _, err = RunArgsInTest(ctx, "singularity wallet add-remote f1l2cc5vuw5moppwsjd3b7cjtwa2exowqo36esklq 12D3KooWD3eckifWpRn9wQpMG9R9hX3sD158z7EqHWmweQAJU5SA")
		require.NoError(t, err)
		out, _, err = RunArgsInTest(ctx, "singularity wallet list ")
		require.NoError(t, err)
		require.Contains(t, out, "f02170643")
		out, _, err = RunArgsInTest(ctx, "singularity dataset add-wallet test f0808055")
		require.NoError(t, err)
		require.Contains(t, out, "f0808055")
		out, _, err = RunArgsInTest(ctx, "singularity dataset list-wallet test")
		require.NoError(t, err)
		require.Contains(t, out, "f0808055")
		_, _, err = RunArgsInTest(ctx, "singularity dataset remove-wallet test f0808055")
		require.NoError(t, err)
		out, _, err = RunArgsInTest(ctx, "singularity dataset list-wallet test")
		require.NoError(t, err)
		require.NotContains(t, out, "f0808055")
		_, _, err = RunArgsInTest(ctx, "singularity wallet remove --really-do-it f0808055")
		require.NoError(t, err)
		out, _, err = RunArgsInTest(ctx, "singularity wallet list ")
		require.NoError(t, err)
		require.NotContains(t, out, "f0808055")
	})
}

func TestDatasetAddPiece(t *testing.T) {
	testWithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, _, err := RunArgsInTest(ctx, "singularity dataset create test")
		require.NoError(t, err)
		temp := t.TempDir()
		newFile, err := os.Create(filepath.Join(temp, "test.car"))
		require.NoError(t, err)
		blk := blocks.NewBlock([]byte("test"))
		root := blk.Cid()
		_, err = pack.WriteCarHeader(newFile, root)
		require.NoError(t, err)
		_, err = pack.WriteCarBlock(newFile, blk)
		require.NoError(t, err)
		newFile.Close()
		content, err := os.ReadFile(filepath.Join(temp, "test.car"))
		require.NoError(t, err)
		commp := calculateCommp(t, content, 1048576)
		// Add as path
		_, _, err = RunArgsInTest(ctx, fmt.Sprintf("singularity dataset add-piece -p %s test %s %d",
			escapePath(filepath.Join(temp, "test.car")), commp, 1048576))
		require.NoError(t, err)
		out, _, err := RunArgsInTest(ctx, "singularity dataset list-pieces test")
		require.NoError(t, err)
		require.Contains(t, out, commp)
		// Add as known root
		_, _, err = RunArgsInTest(ctx, fmt.Sprintf("singularity dataset add-piece -r %s test %s %d",
			root.String(), commp, 1048576))
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity dataset list-pieces test")
		require.NoError(t, err)
		// Add as unknown root
		_, _, err = RunArgsInTest(ctx, fmt.Sprintf("singularity dataset add-piece test %s %d",
			commp, 1048576))
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity dataset list-pieces test")
		require.NoError(t, err)
	})
}

func TestDatasetCrud(t *testing.T) {
	testWithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tmp := t.TempDir()
		out, _, err := RunArgsInTest(ctx, "singularity dataset create test")
		require.NoError(t, err)
		require.Contains(t, out, "test")
		out, _, err = RunArgsInTest(ctx, "singularity dataset list")
		require.NoError(t, err)
		require.Contains(t, out, "test")
		out, _, err = RunArgsInTest(ctx, "singularity dataset update --output-dir "+escapePath(tmp)+" --max-size 1000 test")
		require.NoError(t, err)
		require.Contains(t, out, tmp)
		require.Contains(t, out, "1000")
		_, _, err = RunArgsInTest(ctx, "singularity dataset remove --really-do-it test")
		require.NoError(t, err)
		out, _, err = RunArgsInTest(ctx, "singularity dataset list")
		require.NoError(t, err)
		require.NotContains(t, out, "test")
	})
}

func TestDatasourceCrud(t *testing.T) {
	testWithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		temp := t.TempDir()
		err := os.Mkdir(filepath.Join(temp, "sub"), 0777)
		require.NoError(t, err)
		err = os.WriteFile(filepath.Join(temp, "sub", "test.txt"), []byte("hello world"), 0777)
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity dataset create test")
		require.NoError(t, err)
		out, _, err := RunArgsInTest(ctx, "singularity datasource add local test "+escapePath(temp))
		require.NoError(t, err)
		require.Contains(t, out, temp)
		out, _, err = RunArgsInTest(ctx, "singularity datasource list")
		require.NoError(t, err)
		require.Contains(t, out, temp)
		out, _, err = RunArgsInTest(ctx, "singularity datasource list --dataset test")
		require.NoError(t, err)
		require.Contains(t, out, temp)
		_, _, err = RunArgsInTest(ctx, "singularity datasource list --dataset notexist")
		require.Error(t, err)
		out, _, err = RunArgsInTest(ctx, "singularity datasource check 1")
		require.NoError(t, err)
		require.Contains(t, out, "sub")
		out, _, err = RunArgsInTest(ctx, "singularity datasource check 1 sub")
		require.NoError(t, err)
		require.Contains(t, out, "sub/test.txt")
		out, _, err = RunArgsInTest(ctx, "singularity datasource update --local-case-sensitive=true --rescan-interval 1h 1")
		require.NoError(t, err)
		require.Contains(t, out, "case_sensitive:true")
		require.Contains(t, out, "3600")
		_, _, err = RunArgsInTest(ctx, "singularity datasource remove --really-do-it 1")
		require.NoError(t, err)
		out, _, err = RunArgsInTest(ctx, "singularity datasource list")
		require.NoError(t, err)
		require.NotContains(t, out, temp)
	})
}

func TestEncryption(t *testing.T) {
	testWithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		temp := t.TempDir()
		carDir := t.TempDir()
		content1 := generateRandomBytes(10)
		content2 := generateRandomBytes(10_000_000)
		os.WriteFile(filepath.Join(temp, "test1.txt"), content1, 0777)
		os.WriteFile(filepath.Join(temp, "test2.txt"), content2, 0777)
		public := "age1th55qj77d32vhumd72de2m3y0nzsxyeahuddz770s8qadz3h6v8quedwf3"
		private := "AGE-SECRET-KEY-1HZG3ESWDVPE3S4AM8WWCZG3H66A6RVJPXPZZEAC04FWZVT6RJ7XQAUV49J"
		_, _, err := RunArgsInTest(ctx, "singularity dataset create --max-size 1500000 -o "+escapePath(carDir)+" --encryption-recipient "+public+" test")
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity datasource add local test "+escapePath(temp))
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		// Run the daggen
		_, _, err = RunArgsInTest(ctx, "singularity datasource daggen 1")
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		// Get the root CID
		out, _, err := RunArgsInTest(ctx, "singularity --json datasource inspect path 1")
		require.NoError(t, err)
		root := strings.Split(strings.Split(out, "\n")[4], "\"")[3]
		bs := loadCars(t, carDir)
		dagServ := merkledag.NewDAGService(blockservice.New(bs, nil))
		rootCID, err := cid.Decode(root)
		require.NoError(t, err)
		content1enc := getFileFromRootNode(t, dagServ, "test1.txt", rootCID)
		content2enc := getFileFromRootNode(t, dagServ, "test2.txt", rootCID)
		require.Equal(t, content1, decrypt(t, private, content1enc))
		require.Equal(t, content2, decrypt(t, private, content2enc))
	})
}

func TestDatasourcePacking(t *testing.T) {
	testWithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		original := uio.HAMTShardingSize
		uio.HAMTShardingSize = 1024
		defer func() { uio.HAMTShardingSize = original }()
		c := 100
		temp := t.TempDir()
		carDir := t.TempDir()
		// multiple nested folder
		err := os.MkdirAll(filepath.Join(temp, "sub1", "sub2", "sub3", "sub4"), 0777)
		require.NoError(t, err)
		// dynamic directory with 10k files
		for i := 0; i < c; i++ {
			err = os.WriteFile(filepath.Join(temp, "sub1", "sub2", "sub3", "sub4", "test"+strconv.Itoa(i)+".txt"), generateRandomBytes(10), 0777)
			require.NoError(t, err)
		}
		// dynamic directory with 10k folders
		for i := 0; i < c; i++ {
			err = os.MkdirAll(filepath.Join(temp, strconv.Itoa(i)), 0777)
			require.NoError(t, err)
			err = os.WriteFile(filepath.Join(temp, strconv.Itoa(i), "test"+strconv.Itoa(i)+".txt"), generateRandomBytes(10), 0777)
			require.NoError(t, err)
		}
		// file of large size
		err = os.WriteFile(filepath.Join(temp, "test1.txt"), generateRandomBytes(10000), 0777)
		require.NoError(t, err)
		// file of empty size
		err = os.WriteFile(filepath.Join(temp, "test2.txt"), []byte{}, 0777)
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity dataset create --max-size 1000 -o "+escapePath(carDir)+" test")
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity datasource add local test "+escapePath(temp))
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		// Check the car folder
		files, err := os.ReadDir(carDir)
		require.NoError(t, err)
		require.Equal(t, 26, len(files))
		// Run the daggen
		_, _, err = RunArgsInTest(ctx, "singularity datasource daggen 1")
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		files, err = os.ReadDir(carDir)
		require.NoError(t, err)
		require.Equal(t, 27, len(files))
		// Get the root CID
		out, _, err := RunArgsInTest(ctx, "singularity --json datasource inspect path 1")
		require.NoError(t, err)
		root := strings.Split(strings.Split(out, "\n")[4], "\"")[3]
		_, _, err = RunArgsInTest(ctx, "singularity tool extract-car -i "+escapePath(carDir)+" -o "+escapePath(carDir)+" -c "+root)
		require.NoError(t, err)
	})
}

func TestDatasourceRescan(t *testing.T) {
	testWithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		temp := t.TempDir()
		err := os.Mkdir(filepath.Join(temp, "sub"), 0777)
		require.NoError(t, err)
		err = os.WriteFile(filepath.Join(temp, "sub", "test1.txt"), generateRandomBytes(10), 0777)
		require.NoError(t, err)
		err = os.WriteFile(filepath.Join(temp, "sub", "test2.txt"), generateRandomBytes(100), 0777)
		require.NoError(t, err)
		err = os.WriteFile(filepath.Join(temp, "sub", "test3.txt"), generateRandomBytes(1000), 0777)
		require.NoError(t, err)
		err = os.WriteFile(filepath.Join(temp, "sub", "test4.txt"), generateRandomBytes(10000), 0777)
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity dataset create --max-size 1000 test")
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity datasource add local test "+escapePath(temp))
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --enable-pack=false --enable-dag=false --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		out, _, err := RunArgsInTest(ctx, "singularity datasource inspect packjobs 1")
		require.NoError(t, err)
		require.Contains(t, out, "ready")
		// We should get 15 pack jobs
		require.Contains(t, out, "15")
		err = os.WriteFile(filepath.Join(temp, "sub", "test5.txt"), generateRandomBytes(10000), 0777)
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity datasource rescan 1")
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --enable-pack=false --enable-dag=false --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		out, _, err = RunArgsInTest(ctx, "singularity datasource inspect packjobs 1")
		require.NoError(t, err)
		// We should get 29 packjobs
		require.Contains(t, out, "29")
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --enable-pack=true --enable-dag=false --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		out, _, err = RunArgsInTest(ctx, "singularity datasource inspect packjobs 1")
		require.NoError(t, err)
		require.NotContains(t, out, "ready")
		require.Contains(t, out, "complete")
		out, _, err = RunArgsInTest(ctx, "singularity datasource inspect files 1")
		require.NoError(t, err)
		require.Contains(t, out, "baf")
		require.Contains(t, out, "test5.txt")
		out, _, err = RunArgsInTest(ctx, "singularity datasource inspect packjobdetail 1")
		require.NoError(t, err)
		require.Contains(t, out, "sub/test1.txt")
		require.Contains(t, out, "sub/test3.txt")
		out, _, err = RunArgsInTest(ctx, "singularity datasource inspect path 1")
		require.NoError(t, err)
		require.Contains(t, out, "sub")
		out, _, err = RunArgsInTest(ctx, "singularity datasource inspect path 1 sub/")
		require.NoError(t, err)
		require.Contains(t, out, "sub/test1.txt")
		require.Contains(t, out, "sub/test3.txt")
		out, _, err = RunArgsInTest(ctx, "singularity datasource daggen 1")
		require.NoError(t, err)
		require.Contains(t, out, "ready")
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --enable-pack=false --enable-dag=true --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		out2, _, err := RunArgsInTest(ctx, "singularity datasource inspect dags 1")
		require.NoError(t, err)
		require.Contains(t, out2, "baf")
		out, _, err = RunArgsInTest(ctx, "singularity datasource daggen 1")
		require.NoError(t, err)
		require.Contains(t, out, "ready")
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --enable-pack=false --enable-dag=true --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		out3, _, err := RunArgsInTest(ctx, "singularity datasource inspect dags 1")
		require.NoError(t, err)
		require.Equal(t, out3, out2)
	})
}

func TestPieceDownload(t *testing.T) {
	testWithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		temp1 := t.TempDir()
		temp2 := t.TempDir()
		err := os.WriteFile(filepath.Join(temp1, "test1.txt"), generateRandomBytes(10), 0777)
		require.NoError(t, err)
		err = os.WriteFile(filepath.Join(temp1, "test2.txt"), generateRandomBytes(10), 0777)
		require.NoError(t, err)
		err = os.WriteFile(filepath.Join(temp1, "test3.txt"), generateRandomBytes(2_000_000), 0777)
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity dataset create --max-size 1MB --output-dir "+escapePath(temp2)+" test")
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity datasource add local test "+escapePath(temp1))
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity datasource daggen 1")
		require.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		out, _, err := RunArgsInTest(ctx, "singularity dataset list-pieces test")
		require.NoError(t, err)
		pieceCIDs := regexp.MustCompile("baga6ea4sea[0-9a-z]+").FindAllString(out, -1)
		require.Len(t, pieceCIDs, 8)
		pieceCIDs = underscore.Unique(pieceCIDs)
		require.Len(t, pieceCIDs, 4)
		ctx2, cancel := context.WithCancel(ctx)
		serverClosed := make(chan struct{})
		defer func() {
			<-serverClosed
		}()
		defer cancel()
		go func() {
			RunArgsInTestNoCapture(ctx2, "singularity run content-provider")
			close(serverClosed)
		}()
		// Wait for HTTP service to be ready
		time.Sleep(2 * time.Second)
		// Wait for HTTP service to shutdown
		defer func() {
			time.Sleep(2 * time.Second)
		}()
		for _, pieceCID := range pieceCIDs {
			content := downloadPiece(t, ctx, pieceCID)
			commp := calculateCommp(t, content, 1024*1024)
			require.Equal(t, pieceCID, commp)
		}
		// Clean up temp2 and try again
		os.RemoveAll(temp2)
		for _, pieceCID := range pieceCIDs {
			content := downloadPiece(t, ctx, pieceCID)
			commp := calculateCommp(t, content, 1024*1024)
			require.Equal(t, pieceCID, commp)
		}
		// multithread download
		for _, pieceCID := range pieceCIDs {
			content := downloadPieceWithThreads(t, ctx, pieceCID, 10)
			commp := calculateCommp(t, content, 1024*1024)
			require.Equal(t, pieceCID, commp)
		}
		// download util
		temp3 := t.TempDir()
		for _, pieceCID := range pieceCIDs {
			_, _, err = RunArgsInTest(ctx, "singularity download --local-links true -o "+escapePath(temp3)+" "+pieceCID)
			require.NoError(t, err)
			content, err := os.ReadFile(filepath.Join(temp3, pieceCID+".car"))
			require.NoError(t, err)
			commp := calculateCommp(t, content, 1024*1024)
			require.Equal(t, pieceCID, commp)
		}
	})
}

func TestGetItemDeals(t *testing.T) {
	testWithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, _, err := RunArgsInTest(ctx, "singularity datasource inspect filedeals 0")
		require.NoError(t, err)
	})
}

func downloadPiece(t *testing.T, ctx context.Context, pieceCID string) []byte {
	t.Log("Downloading piece", pieceCID)
	url := "http://127.0.0.1:7777/piece/" + pieceCID
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	require.NoError(t, err)
	client := &http.Client{}
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		t.Error("Download failed", string(body))
	}
	require.Less(t, resp.StatusCode, 300)
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	return body
}

func downloadPieceWithThreads(t *testing.T, ctx context.Context, pieceCID string, nThreads int) []byte {
	url := "http://127.0.0.1:7777/piece/" + pieceCID

	// Make a HEAD request to get the size of the file
	req, err := http.NewRequestWithContext(ctx, "HEAD", url, nil)
	require.NoError(t, err)
	client := &http.Client{}
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Get the Content-Length header
	contentLength := resp.ContentLength
	if contentLength < 0 {
		t.Error("Content-Length header not found")
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
			require.NoError(t, err)

			// Set the Range header to download a pack job
			req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end))
			// t.Log("Downloading piece", pieceCID, "part", i, "bytes", start, "-", end)
			resp, err := client.Do(req)
			require.NoError(t, err)
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)

			// Save the part to the slice
			parts[i] = body
		}(i)
	}

	// Wait for all the downloads to finish
	wg.Wait()

	// Combine the parts
	var result bytes.Buffer
	for _, part := range parts {
		result.Write(part)
	}

	return result.Bytes()
}

func calculateCommp(t *testing.T, content []byte, targetPieceSize uint64) string {
	calc := &commp.Calc{}
	_, err := bytes.NewBuffer(content).WriteTo(calc)
	require.NoError(t, err)
	c, _, err := pack.GetCommp(calc, targetPieceSize)
	require.NoError(t, err)
	return c.String()
}
