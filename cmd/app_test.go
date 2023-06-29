package cmd

import (
	"bufio"
	"bytes"
	"context"
	"filippo.io/age"
	"fmt"
	"github.com/data-preservation-programs/singularity/database"
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
	"github.com/rjNemo/underscore"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

func generateRandomBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}

func decrypt(t *testing.T, key string, encrypted []byte) []byte {
	recipient, err := age.ParseX25519Identity(key)
	assert.NoError(t, err)
	decrypted, err := age.Decrypt(bytes.NewReader(encrypted), recipient)
	assert.NoError(t, err)
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, decrypted)
	assert.NoError(t, err)
	return buf.Bytes()
}

func loadCars(t *testing.T, path string) blockstore.Blockstore {
	files, err := os.ReadDir(path)
	assert.NoError(t, err)
	bs := blockstore.NewBlockstore(datastore.NewMapDatastore())
	for _, file := range files {
		f, err := os.Open(path + "/" + file.Name())
		assert.NoError(t, err)
		reader := bufio.NewReader(f)
		_, err = car.ReadHeader(reader)
		assert.NoError(t, err)
		for {
			c, data, err := util.ReadNode(reader)
			if err == io.EOF {
				break
			}
			blk, _ := blocks.NewBlockWithCid(data, c)
			err = bs.Put(context.TODO(), blk)
			assert.NoError(t, err)
		}
	}
	return bs
}

func getFileFromRootNode(t *testing.T, dagServ format.DAGService, path string, rootCID cid.Cid) []byte {
	ctx := context.TODO()
	segments := strings.Split(path, "/")
	for _, segment := range segments {
		rootNode, err := dagServ.Get(context.Background(), rootCID)
		assert.NoError(t, err)
		rootDir, err := uio.NewDirectoryFromNode(dagServ, rootNode)
		assert.NoError(t, err)
		links, err := rootDir.Links(ctx)
		assert.NoError(t, err)
		link, err := underscore.Find(links, func(link *format.Link) bool {
			return link.Name == segment
		})
		assert.NoError(t, err)
		rootCID = link.Cid
	}
	fileNode, err := dagServ.Get(ctx, rootCID)
	assert.NoError(t, err)
	dagReader, err := uio.NewDagReader(ctx, fileNode, dagServ)
	assert.NoError(t, err)
	content, err := io.ReadAll(dagReader)
	assert.NoError(t, err)
	return content
}
func listDirsFromRootNode(t *testing.T, dagServ format.DAGService, path string, rootCID cid.Cid) []string {
	ctx := context.TODO()
	segments := strings.Split(path, "/")
	if path == "" {
		segments = []string{}
	}
	for _, segment := range segments {
		rootNode, err := dagServ.Get(context.Background(), rootCID)
		assert.NoError(t, err)
		rootDir, err := uio.NewDirectoryFromNode(dagServ, rootNode)
		assert.NoError(t, err)
		links, err := rootDir.Links(ctx)
		assert.NoError(t, err)
		link, err := underscore.Find(links, func(link *format.Link) bool {
			return link.Name == segment
		})
		assert.NoError(t, err)
		rootCID = link.Cid
	}

	rootNode, err := dagServ.Get(context.Background(), rootCID)
	assert.NoError(t, err)
	rootDir, err := uio.NewDirectoryFromNode(dagServ, rootNode)
	assert.NoError(t, err)
	links, err := rootDir.Links(ctx)
	return underscore.Map(links, func(link *format.Link) string {
		return link.Name
	})
}

func testWithAllBackendWithoutReset(t *testing.T, testFunc func(ctx context.Context, t *testing.T, db *gorm.DB)) {
	testWithAllBackendWithResetArg(t, testFunc, false)
}

func testWithAllBackendWithResetArg(t *testing.T, testFunc func(ctx context.Context, t *testing.T, db *gorm.DB), reset bool) {
	backends := [][2]string{
		{"sqlite", "sqlite:" + t.TempDir() + "/singularity.db"},
		//{"mysql", "mysql://root:password@tcp(localhost:3306)/singularity?parseTime=true"},
		//{"postgres", "postgres://postgres:password@localhost:5432/singularity?sslmode=disable"},
	}
	for _, backend := range backends {
		os.Setenv("DATABASE_CONNECTION_STRING", backend[1])
		if reset {
			_, _, err := RunArgsInTest(context.Background(), "singularity admin reset")
			assert.NoError(t, err)
		}
		db, err := database.Open(backend[1], &gorm.Config{})
		assert.NoError(t, err)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
		defer cancel()
		t.Run(backend[0], func(t *testing.T) {
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
		assert.NoError(t, err)
	})
}

func TestResetDatabase(t *testing.T) {
	testWithAllBackendWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, _, err := RunArgsInTest(ctx, "singularity admin reset")
		assert.NoError(t, err)
	})
}

func TestInitDatabase(t *testing.T) {
	testWithAllBackendWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, _, err := RunArgsInTest(ctx, "singularity admin init")
		assert.NoError(t, err)
	})
}

func TestDatasetCrud(t *testing.T) {
	testWithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		out, _, err := RunArgsInTest(ctx, "singularity dataset create test")
		assert.NoError(t, err)
		assert.Contains(t, out, "test")
		out, _, err = RunArgsInTest(ctx, "singularity dataset list")
		assert.NoError(t, err)
		assert.Contains(t, out, "test")
		out, _, err = RunArgsInTest(ctx, "singularity dataset update --output-dir /tmp --max-size 1000 test")
		assert.NoError(t, err)
		assert.Contains(t, out, "/tmp")
		assert.Contains(t, out, "1000")
		_, _, err = RunArgsInTest(ctx, "singularity dataset remove test")
		assert.NoError(t, err)
		out, _, err = RunArgsInTest(ctx, "singularity dataset list")
		assert.NoError(t, err)
		assert.NotContains(t, out, "test")
	})
}

func TestEzPrepBenchmark(t *testing.T) {
	temp := t.TempDir()
	exec.Command("truncate", "-s", "1G", temp+"/test.img").Run()
	ctx := context.Background()
	out, _, err := RunArgsInTest(ctx, "singularity ez-prep --output-dir '' --database-file '' -j 8 "+temp)
	assert.NoError(t, err)
	// contains two CARs, one for the file and another one for the dag
	assert.Contains(t, out, "1073833069")
	assert.Contains(t, out, "156")
}

func TestDatasourceCrud(t *testing.T) {
	testWithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		temp := t.TempDir()
		os.Mkdir(temp+"/sub", 0777)
		os.WriteFile(temp+"/sub/test.txt", []byte("hello world"), 0777)
		_, _, err := RunArgsInTest(ctx, "singularity dataset create test")
		assert.NoError(t, err)
		out, _, err := RunArgsInTest(ctx, "singularity datasource add local test "+temp)
		assert.NoError(t, err)
		assert.Contains(t, out, temp)
		out, _, err = RunArgsInTest(ctx, "singularity datasource list")
		assert.NoError(t, err)
		assert.Contains(t, out, temp)
		out, _, err = RunArgsInTest(ctx, "singularity datasource list --dataset test")
		assert.NoError(t, err)
		assert.Contains(t, out, temp)
		out, _, err = RunArgsInTest(ctx, "singularity datasource list --dataset notexist")
		assert.Error(t, err)
		out, _, err = RunArgsInTest(ctx, "singularity datasource check 1")
		assert.NoError(t, err)
		assert.Contains(t, out, "sub")
		out, _, err = RunArgsInTest(ctx, "singularity datasource check 1 sub")
		assert.NoError(t, err)
		assert.Contains(t, out, "sub/test.txt")
		out, _, err = RunArgsInTest(ctx, "singularity datasource update --local-case-sensitive=true --rescan-interval 1h 1")
		assert.NoError(t, err)
		assert.Contains(t, out, "case_sensitive:true")
		assert.Contains(t, out, "3600")
		out, _, err = RunArgsInTest(ctx, "singularity datasource remove 1")
		assert.NoError(t, err)
		out, _, err = RunArgsInTest(ctx, "singularity datasource list")
		assert.NoError(t, err)
		assert.NotContains(t, out, temp)
	})
}

func TestEncryption(t *testing.T) {
	testWithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		temp := t.TempDir()
		carDir := t.TempDir()
		content1 := generateRandomBytes(10)
		content2 := generateRandomBytes(10_000_000)
		os.WriteFile(temp+"/test1.txt", content1, 0777)
		os.WriteFile(temp+"/test2.txt", content2, 0777)
		public := "age1th55qj77d32vhumd72de2m3y0nzsxyeahuddz770s8qadz3h6v8quedwf3"
		private := "AGE-SECRET-KEY-1HZG3ESWDVPE3S4AM8WWCZG3H66A6RVJPXPZZEAC04FWZVT6RJ7XQAUV49J"
		_, _, err := RunArgsInTest(ctx, "singularity dataset create --max-size 1500000 -o "+carDir+" --encryption-recipient "+public+" test")
		assert.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity datasource add local test "+temp)
		assert.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		assert.NoError(t, err)
		// Run the daggen
		_, _, err = RunArgsInTest(ctx, "singularity datasource daggen 1")
		assert.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		assert.NoError(t, err)
		// Get the root CID
		out, _, err := RunArgsInTest(ctx, "singularity --json datasource inspect path 1")
		assert.NoError(t, err)
		root := strings.Split(strings.Split(out, "\n")[4], "\"")[3]
		bs := loadCars(t, carDir)
		dagServ := merkledag.NewDAGService(blockservice.New(bs, nil))
		rootCID := cid.MustParse(root)
		content1enc := getFileFromRootNode(t, dagServ, "test1.txt", rootCID)
		content2enc := getFileFromRootNode(t, dagServ, "test2.txt", rootCID)
		assert.Equal(t, content1, decrypt(t, private, content1enc))
		assert.Equal(t, content2, decrypt(t, private, content2enc))
	})
}

func TestDatasourcePacking(t *testing.T) {
	testWithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		original := uio.HAMTShardingSize
		uio.HAMTShardingSize = 1024
		defer func() { uio.HAMTShardingSize = original }()
		c := 1_000
		temp := t.TempDir()
		carDir := t.TempDir()
		// multiple nested folder
		os.MkdirAll(temp+"/sub1/sub2/sub3/sub4", 0777)
		// dynamic directory with 10k files
		for i := 0; i < c; i++ {
			os.WriteFile(temp+"/sub1/sub2/sub3/sub4/test"+strconv.Itoa(i)+".txt", generateRandomBytes(10), 0777)
		}
		// dynamic directory with 10k folders
		for i := 0; i < c; i++ {
			os.MkdirAll(temp+"/"+strconv.Itoa(i), 0777)
			os.WriteFile(temp+"/"+strconv.Itoa(i)+"/test"+strconv.Itoa(i)+".txt", generateRandomBytes(10), 0777)
		}
		// file of large size
		os.WriteFile(temp+"/test1.txt", generateRandomBytes(10000), 0777)
		// file of empty size
		os.WriteFile(temp+"/test2.txt", []byte{}, 0777)
		_, _, err := RunArgsInTest(ctx, "singularity dataset create --max-size 1000 -o "+carDir+" test")
		assert.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity datasource add local test "+temp)
		assert.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		assert.NoError(t, err)
		// Check the car folder
		files, err := os.ReadDir(carDir)
		assert.NoError(t, err)
		assert.Equal(t, 131, len(files))
		// Run the daggen
		_, _, err = RunArgsInTest(ctx, "singularity datasource daggen 1")
		assert.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		assert.NoError(t, err)
		files, err = os.ReadDir(carDir)
		assert.NoError(t, err)
		assert.Equal(t, 132, len(files))
		// Get the root CID
		out, _, err := RunArgsInTest(ctx, "singularity --json datasource inspect path 1")
		assert.NoError(t, err)
		root := strings.Split(strings.Split(out, "\n")[4], "\"")[3]
		// Now load all car files to a block store and check if the resolved directory is same as the original
		t.Log(root)
		bs := loadCars(t, carDir)
		dagServ := merkledag.NewDAGService(blockservice.New(bs, nil))
		rootCID := cid.MustParse(root)
		entries := listDirsFromRootNode(t, dagServ, "", rootCID)
		var content []byte
		assert.Equal(t, 1003, len(entries))
		assert.True(t, slices.Contains(entries, "sub1"))
		assert.True(t, slices.Contains(entries, "test1.txt"))
		assert.True(t, slices.Contains(entries, "test2.txt"))
		assert.True(t, slices.Contains(entries, "0"))
		assert.True(t, slices.Contains(entries, "999"))
		content, err = os.ReadFile(temp + "/2/test2.txt")
		assert.Equal(t, content, getFileFromRootNode(t, dagServ, "2/test2.txt", rootCID))
		content, err = os.ReadFile(temp + "/sub1/sub2/sub3/sub4/test2.txt")
		assert.Equal(t, content, getFileFromRootNode(t, dagServ, "sub1/sub2/sub3/sub4/test2.txt", rootCID))
		content, err = os.ReadFile(temp + "/test1.txt")
		assert.Equal(t, content, getFileFromRootNode(t, dagServ, "test1.txt", rootCID))
		content, err = os.ReadFile(temp + "/test2.txt")
		assert.Equal(t, content, getFileFromRootNode(t, dagServ, "test2.txt", rootCID))
	})
}

func TestDatasourceRescan(t *testing.T) {
	testWithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		temp := t.TempDir()
		os.Mkdir(temp+"/sub", 0777)
		os.WriteFile(temp+"/sub/test1.txt", generateRandomBytes(10), 0777)
		os.WriteFile(temp+"/sub/test2.txt", generateRandomBytes(100), 0777)
		os.WriteFile(temp+"/sub/test3.txt", generateRandomBytes(1000), 0777)
		os.WriteFile(temp+"/sub/test4.txt", generateRandomBytes(10000), 0777)
		_, _, err := RunArgsInTest(ctx, "singularity dataset create --max-size 1000 test")
		assert.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity datasource add local test "+temp)
		assert.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --enable-pack=false --enable-dag=false --exit-on-complete=true --exit-on-error=true")
		assert.NoError(t, err)
		out, _, err := RunArgsInTest(ctx, "singularity datasource inspect chunks 1")
		assert.NoError(t, err)
		assert.Contains(t, out, "ready")
		// We should get 15 chunks
		assert.Contains(t, out, "15")
		os.WriteFile(temp+"/sub/test5.txt", generateRandomBytes(10000), 0777)
		_, _, err = RunArgsInTest(ctx, "singularity datasource rescan 1")
		assert.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --enable-pack=false --enable-dag=false --exit-on-complete=true --exit-on-error=true")
		assert.NoError(t, err)
		out, _, err = RunArgsInTest(ctx, "singularity datasource inspect chunks 1")
		assert.NoError(t, err)
		// We should get 29 chunks
		assert.Contains(t, out, "29")
		assert.NotContains(t, out, "30")
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --enable-pack=true --enable-dag=false --exit-on-complete=true --exit-on-error=true")
		assert.NoError(t, err)
		out, _, err = RunArgsInTest(ctx, "singularity datasource inspect chunks 1")
		assert.NoError(t, err)
		assert.NotContains(t, out, "ready")
		assert.Contains(t, out, "complete")
		assert.Contains(t, out, "baf")
		assert.Contains(t, out, "baga")
		out, _, err = RunArgsInTest(ctx, "singularity datasource inspect items 1")
		assert.Contains(t, out, "baf")
		assert.Contains(t, out, "test5.txt")
		out, _, err = RunArgsInTest(ctx, "singularity datasource inspect chunkdetail 1")
		assert.Contains(t, out, "sub/test1.txt")
		assert.Contains(t, out, "sub/test3.txt")
		out, _, err = RunArgsInTest(ctx, "singularity datasource inspect path 1")
		assert.NoError(t, err)
		assert.Contains(t, out, "sub")
		out, _, err = RunArgsInTest(ctx, "singularity datasource inspect path 1 sub/")
		assert.NoError(t, err)
		assert.Contains(t, out, "sub/test1.txt")
		assert.Contains(t, out, "sub/test3.txt")
		out, _, err = RunArgsInTest(ctx, "singularity datasource daggen 1")
		assert.NoError(t, err)
		assert.Contains(t, out, "ready")
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --enable-pack=false --enable-dag=true --exit-on-complete=true --exit-on-error=true")
		assert.NoError(t, err)
		out2, _, err := RunArgsInTest(ctx, "singularity datasource inspect dags 1")
		assert.NoError(t, err)
		assert.Contains(t, out2, "baf")
		out, _, err = RunArgsInTest(ctx, "singularity datasource daggen 1")
		assert.NoError(t, err)
		assert.Contains(t, out, "ready")
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --enable-pack=false --enable-dag=true --exit-on-complete=true --exit-on-error=true")
		assert.NoError(t, err)
		out3, _, err := RunArgsInTest(ctx, "singularity datasource inspect dags 1")
		assert.NoError(t, err)
		assert.Equal(t, out3, out2)
	})
}

func TestPieceDownload(t *testing.T) {
	testWithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		temp1 := t.TempDir()
		temp2 := t.TempDir()
		os.WriteFile(temp1+"/test1.txt", generateRandomBytes(10), 0777)
		os.WriteFile(temp1+"/test2.txt", generateRandomBytes(10), 0777)
		os.WriteFile(temp1+"/test3.txt", generateRandomBytes(10_000_000), 0777)
		_, _, err := RunArgsInTest(ctx, "singularity dataset create --max-size 4MB --output-dir "+temp2+" test")
		assert.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity datasource add local test "+temp1)
		assert.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		assert.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity datasource daggen 1")
		assert.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		assert.NoError(t, err)
		out, _, err := RunArgsInTest(ctx, "singularity dataset list-pieces test")
		assert.NoError(t, err)
		pieceCIDs := regexp.MustCompile("baga6ea4sea[0-9a-z]+").FindAllString(out, -1)
		assert.Len(t, pieceCIDs, 8)
		pieceCIDs = underscore.Unique(pieceCIDs)
		assert.Len(t, pieceCIDs, 4)
		ctx2, cancel := context.WithCancel(ctx)
		defer cancel()
		go func() {
			_, _, err := RunArgsInTest(ctx2, "singularity run content-provider")
			assert.NoError(t, err)
			<-ctx2.Done()
		}()
		// Wait for HTTP service to be ready
		time.Sleep(1 * time.Second)
		for _, pieceCID := range pieceCIDs {
			content := downloadPiece(t, ctx, pieceCID)
			commp := calculateCommp(t, content, 4*1024*1024)
			assert.Equal(t, pieceCID, commp)
		}
		// Clean up temp2 and try again
		os.RemoveAll(temp2)
		for _, pieceCID := range pieceCIDs {
			content := downloadPiece(t, ctx, pieceCID)
			commp := calculateCommp(t, content, 4*1024*1024)
			assert.Equal(t, pieceCID, commp)
		}
		// multithread download
		for _, pieceCID := range pieceCIDs {
			content := downloadPieceWithThreads(t, ctx, pieceCID, 10)
			commp := calculateCommp(t, content, 4*1024*1024)
			assert.Equal(t, pieceCID, commp)
		}
	})
}

func downloadPiece(t *testing.T, ctx context.Context, pieceCID string) []byte {
	t.Log("Downloading piece", pieceCID)
	url := "http://127.0.0.1:7777/piece/" + pieceCID
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	assert.NoError(t, err)
	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Less(t, resp.StatusCode, 300)
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	return body
}

func downloadPieceWithThreads(t *testing.T, ctx context.Context, pieceCID string, nThreads int) []byte {
	url := "http://127.0.0.1:7777/piece/" + pieceCID

	// Make a HEAD request to get the size of the file
	req, err := http.NewRequestWithContext(ctx, "HEAD", url, nil)
	assert.NoError(t, err)
	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)
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
			assert.NoError(t, err)

			// Set the Range header to download a chunk
			req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end))
			t.Log("Downloading piece", pieceCID, "part", i, "bytes", start, "-", end)
			resp, err := client.Do(req)
			assert.NoError(t, err)
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			assert.NoError(t, err)

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
	assert.NoError(t, err)
	c, _, err := pack.GetCommp(calc, targetPieceSize)
	assert.NoError(t, err)
	return c.String()
}
