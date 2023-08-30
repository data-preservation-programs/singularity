package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/util/testutil"
	uio "github.com/ipfs/go-unixfs/io"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// TestPrepCreateWithLocalSource tests the following scenario:
// 1. Create a local source with a few files
//   - file of different sizes
//   - nested folders
//   - folder containing lots of files
//
// 2. Create a local output
// 3. Create a preparation with the local source and output
// 4. Start scanning, packing and daggen
// 5. Download the pieces using the piece API
// 6. Download the pieces using the metadata API with download utility
// 7. Extract into folder and compare with the original source
// 8. Repeat above with different maxSize and inline
func TestDataPrepWithLocalSource(t *testing.T) {
	// Prepare local source
	tmp := t.TempDir()
	originalShardingSize := uio.HAMTShardingSize
	uio.HAMTShardingSize = 1024
	defer func() { uio.HAMTShardingSize = originalShardingSize }()

	// create 100 random files
	for i := 0; i < 100; i++ {
		file := filepath.Join(tmp, fmt.Sprintf("file-%d.txt", i))
		content := testutil.GenerateFixedBytes(i)
		err := os.WriteFile(file, content, 0777)
		require.NoError(t, err)
	}

	// create 10 nested folders
	folderPath := tmp
	for i := 0; i < 10; i++ {
		folderPath = filepath.Join(folderPath, fmt.Sprintf("folder-%d", i))
	}
	err := os.MkdirAll(folderPath, 0777)
	require.NoError(t, err)
	err = os.WriteFile(filepath.Join(folderPath, "file.txt"), testutil.GenerateFixedBytes(1000), 0777)

	// create file of different sizes
	sizes := []int{0, 1, 1 << 20, 1<<20 + 1, 20 << 20}
	for _, size := range sizes {
		err = os.WriteFile(filepath.Join(tmp, fmt.Sprintf("size-%d.txt", size)), testutil.GenerateFixedBytes(size), 0777)
	}

	tests := []struct {
		maxSize int
		inline  bool
	}{
		{maxSize: 60 << 20, inline: false},
		{maxSize: 60 << 20, inline: true},
		{maxSize: 15 << 20, inline: false},
		{maxSize: 15 << 20, inline: true},
		{maxSize: 7 << 20, inline: false},
		{maxSize: 7 << 20, inline: true},
		{maxSize: 3 << 20, inline: false},
		{maxSize: 3 << 20, inline: true},
	}
	for _, test := range tests {
		name := fmt.Sprintf("maxSize-%d-inline-%v", test.maxSize, test.inline)
		t.Run(name, func(t *testing.T) {
			testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
				runner := Runner{}
				defer runner.Save(t)
				_, _, err := runner.Run(ctx, fmt.Sprintf("singularity storage create local source '%s'", testutil.EscapePath(tmp)))
				require.NoError(t, err)
				if !test.inline {
					outdir := t.TempDir()
					_, _, err = runner.Run(ctx, fmt.Sprintf("singularity storage create local output '%s'", testutil.EscapePath(outdir)))
					require.NoError(t, err)
					_, _, err = runner.Run(ctx, "singularity prep create --source source --output output")
					require.NoError(t, err)
				} else {
					_, _, err = runner.Run(ctx, "singularity prep create --source source")
					require.NoError(t, err)
				}
			})
		})
	}
}

/*
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

func escape(p string) string {
	return "'" + strings.ReplaceAll(p, `\`, `\\`) + "'"
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



func TestInitDatabase(t *testing.T) {
	WithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, _, err := Run(ctx, "singularity admin init")
		require.NoError(t, err)
	})
}

func TestListDeals(t *testing.T) {
	WithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, _, err := Run(ctx, "singularity deal list --dataset test --schedule 1 --provider f01 --state active")
		require.NoError(t, err)
	})
}

func TestResetDatabaseReallyDoIt(t *testing.T) {
	WithAllBackendWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, _, err := Run(ctx, "singularity admin reset --really-do-it")
		require.NoError(t, err)
	})
}

func TestResetDatabase(t *testing.T) {
	WithAllBackendWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, _, err := Run(ctx, "singularity admin reset")
		require.ErrorContains(t, err, "really-do-it")
	})
}

func TestDealScheduleCrud(t *testing.T) {
	WithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, _, err := Run(ctx, "singularity dataset create test")
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity wallet add-remote f1l2cc5vuw5moppwsjd3b7cjtwa2exowqo36esklq 12D3KooWD3eckifWpRn9wQpMG9R9hX3sD158z7EqHWmweQAJU5SA")
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity dataset add-wallet test f02170643")
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity deal schedule create test f022352")
		require.NoError(t, err)
		body, _, err := Run(ctx, "singularity deal schedule list")
		require.NoError(t, err)
		require.Contains(t, body, "72h")
		body, _, err = Run(ctx, "singularity deal schedule pause 1")
		require.NoError(t, err)
		require.Contains(t, body, "paused")
		body, _, err = Run(ctx, "singularity deal schedule resume 1")
		require.NoError(t, err)
		require.Contains(t, body, "active")
	})
}

func TestWalletCrud(t *testing.T) {
	godotenv.Load("../.env", ".env")
	WithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		key := "7b2254797065223a22736563703235366b31222c22507269766174654b6579223a226b35507976337148327349586343595a58594f5775453149326e32554539436861556b6c4e36695a5763453d227d"
		_, _, err := Run(ctx, "singularity dataset create test")
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity wallet import "+key)
		require.NoError(t, err)
		_, _, err := Run(ctx, "singularity wallet list ")
		require.NoError(t, err)
		require.Contains(t, out, "f0808055")
		_, _, err = Run(ctx, "singularity wallet add-remote f1l2cc5vuw5moppwsjd3b7cjtwa2exowqo36esklq 12D3KooWD3eckifWpRn9wQpMG9R9hX3sD158z7EqHWmweQAJU5SA")
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity wallet list ")
		require.NoError(t, err)
		require.Contains(t, out, "f02170643")
		_, _, err = Run(ctx, "singularity dataset add-wallet test f0808055")
		require.NoError(t, err)
		require.Contains(t, out, "f0808055")
		_, _, err = Run(ctx, "singularity dataset list-wallet test")
		require.NoError(t, err)
		require.Contains(t, out, "f0808055")
		_, _, err = Run(ctx, "singularity dataset remove-wallet test f0808055")
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity dataset list-wallet test")
		require.NoError(t, err)
		require.NotContains(t, out, "f0808055")
		_, _, err = Run(ctx, "singularity wallet remove --really-do-it f0808055")
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity wallet list ")
		require.NoError(t, err)
		require.NotContains(t, out, "f0808055")
	})
}

func TestDatasetAddPiece(t *testing.T) {
	WithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		_, _, err := Run(ctx, "singularity dataset create test")
		require.NoError(t, err)
		temp := t.TempDir()
		newFile, err := os.Create(filepath.Join(temp, "test.car"))
		require.NoError(t, err)
		blk := blocks.NewBlock([]byte("test"))
		root := blk.Cid()
		_, err = util2.WriteCarHeader(newFile, root)
		require.NoError(t, err)
		_, err = util2.WriteCarBlock(newFile, blk)
		require.NoError(t, err)
		newFile.Close()
		content, err := os.ReadFile(filepath.Join(temp, "test.car"))
		require.NoError(t, err)
		commp := calculateCommp(t, content, 1048576)
		// Add as path
		_, _, err = Run(ctx, fmt.Sprintf("singularity dataset add-piece -p %s test %s %d",
			escape(filepath.Join(temp, "test.car")), commp, 1048576))
		require.NoError(t, err)
		_, _, err := Run(ctx, "singularity dataset list-pieces test")
		require.NoError(t, err)
		require.Contains(t, out, commp)
		// Add as known root
		_, _, err = Run(ctx, fmt.Sprintf("singularity dataset add-piece -r %s test %s %d",
			root.String(), commp, 1048576))
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity dataset list-pieces test")
		require.NoError(t, err)
		// Add as unknown root
		_, _, err = Run(ctx, fmt.Sprintf("singularity dataset add-piece test %s %d",
			commp, 1048576))
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity dataset list-pieces test")
		require.NoError(t, err)
	})
}

func TestDatasetCrud(t *testing.T) {
	WithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tmp := t.TempDir()
		_, _, err := Run(ctx, "singularity dataset create test")
		require.NoError(t, err)
		require.Contains(t, out, "test")
		_, _, err = Run(ctx, "singularity dataset list")
		require.NoError(t, err)
		require.Contains(t, out, "test")
		_, _, err = Run(ctx, "singularity dataset update --output-dir "+escape(tmp)+" --max-size 1000 test")
		require.NoError(t, err)
		require.Contains(t, out, tmp)
		require.Contains(t, out, "1000")
		_, _, err = Run(ctx, "singularity dataset remove --really-do-it test")
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity dataset list")
		require.NoError(t, err)
		require.NotContains(t, out, "test")
	})
}

func TestDatasourceCrud(t *testing.T) {
	WithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		temp := t.TempDir()
		err := os.Mkdir(filepath.Join(temp, "sub"), 0777)
		require.NoError(t, err)
		err = os.WriteFile(filepath.Join(temp, "sub", "test.txt"), []byte("hello world"), 0777)
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity dataset create test")
		require.NoError(t, err)
		_, _, err := Run(ctx, "singularity datasource add local test "+escape(temp))
		require.NoError(t, err)
		require.Contains(t, out, temp)
		_, _, err = Run(ctx, "singularity datasource list")
		require.NoError(t, err)
		require.Contains(t, out, temp)
		_, _, err = Run(ctx, "singularity datasource list --dataset test")
		require.NoError(t, err)
		require.Contains(t, out, temp)
		_, _, err = Run(ctx, "singularity datasource list --dataset notexist")
		require.Error(t, err)
		_, _, err = Run(ctx, "singularity datasource check 1")
		require.NoError(t, err)
		require.Contains(t, out, "sub")
		_, _, err = Run(ctx, "singularity datasource check 1 sub")
		require.NoError(t, err)
		require.Contains(t, out, "sub/test.txt")
		_, _, err = Run(ctx, "singularity datasource update --local-case-sensitive=true --rescan-interval 1h 1")
		require.NoError(t, err)
		require.Contains(t, out, "case_sensitive:true")
		require.Contains(t, out, "3600")
		_, _, err = Run(ctx, "singularity datasource remove --really-do-it 1")
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity datasource list")
		require.NoError(t, err)
		require.NotContains(t, out, temp)
	})
}

func TestEncryption(t *testing.T) {
	WithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		temp := t.TempDir()
		carDir := t.TempDir()
		content1 := generateRandomBytes(10)
		content2 := generateRandomBytes(10_000_000)
		os.WriteFile(filepath.Join(temp, "test1.txt"), content1, 0777)
		os.WriteFile(filepath.Join(temp, "test2.txt"), content2, 0777)
		public := "age1th55qj77d32vhumd72de2m3y0nzsxyeahuddz770s8qadz3h6v8quedwf3"
		private := "AGE-SECRET-KEY-1HZG3ESWDVPE3S4AM8WWCZG3H66A6RVJPXPZZEAC04FWZVT6RJ7XQAUV49J"
		_, _, err := Run(ctx, "singularity dataset create --max-size 1500000 -o "+escape(carDir)+" --encryption-recipient "+public+" test")
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity datasource add local test "+escape(temp))
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		// Run the daggen
		_, _, err = Run(ctx, "singularity datasource daggen 1")
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		// Get the root CID
		_, _, err := Run(ctx, "singularity --json datasource inspect path 1")
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
	WithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
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
		_, _, err = Run(ctx, "singularity dataset create --max-size 1000 -o "+escape(carDir)+" test")
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity datasource add local test "+escape(temp))
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		// Check the car folder
		files, err := os.ReadDir(carDir)
		require.NoError(t, err)
		require.Equal(t, 26, len(files))
		// Run the daggen
		_, _, err = Run(ctx, "singularity datasource daggen 1")
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		files, err = os.ReadDir(carDir)
		require.NoError(t, err)
		require.Equal(t, 27, len(files))
		// Get the root CID
		_, _, err := Run(ctx, "singularity --json datasource inspect path 1")
		require.NoError(t, err)
		root := strings.Split(strings.Split(out, "\n")[4], "\"")[3]
		_, _, err = Run(ctx, "singularity tool extract-car -i "+escape(carDir)+" -o "+escape(carDir)+" -c "+root)
		require.NoError(t, err)
	})
}

func TestDatasourceRescan(t *testing.T) {
	WithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
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
		_, _, err = Run(ctx, "singularity dataset create --max-size 1000 test")
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity datasource add local test "+escape(temp))
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity run dataset-worker --enable-pack=false --enable-dag=false --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		_, _, err := Run(ctx, "singularity datasource inspect packjobs 1")
		require.NoError(t, err)
		require.Contains(t, out, "ready")
		// We should get 15 pack jobs
		require.Contains(t, out, "15")
		err = os.WriteFile(filepath.Join(temp, "sub", "test5.txt"), generateRandomBytes(10000), 0777)
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity datasource rescan 1")
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity run dataset-worker --enable-pack=false --enable-dag=false --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity datasource inspect packjobs 1")
		require.NoError(t, err)
		// We should get 29 packjobs
		require.Contains(t, out, "29")
		_, _, err = Run(ctx, "singularity run dataset-worker --enable-pack=true --enable-dag=false --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity datasource inspect packjobs 1")
		require.NoError(t, err)
		require.NotContains(t, out, "ready")
		require.Contains(t, out, "complete")
		_, _, err = Run(ctx, "singularity datasource inspect files 1")
		require.NoError(t, err)
		require.Contains(t, out, "baf")
		require.Contains(t, out, "test5.txt")
		_, _, err = Run(ctx, "singularity datasource inspect packjobdetail 1")
		require.NoError(t, err)
		require.Contains(t, out, "sub/test1.txt")
		require.Contains(t, out, "sub/test3.txt")
		_, _, err = Run(ctx, "singularity datasource inspect path 1")
		require.NoError(t, err)
		require.Contains(t, out, "sub")
		_, _, err = Run(ctx, "singularity datasource inspect path 1 sub/")
		require.NoError(t, err)
		require.Contains(t, out, "sub/test1.txt")
		require.Contains(t, out, "sub/test3.txt")
		_, _, err = Run(ctx, "singularity datasource daggen 1")
		require.NoError(t, err)
		require.Contains(t, out, "ready")
		_, _, err = Run(ctx, "singularity run dataset-worker --enable-pack=false --enable-dag=true --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		out2, _, err := Run(ctx, "singularity datasource inspect dags 1")
		require.NoError(t, err)
		require.Contains(t, out2, "baf")
		_, _, err = Run(ctx, "singularity datasource daggen 1")
		require.NoError(t, err)
		require.Contains(t, out, "ready")
		_, _, err = Run(ctx, "singularity run dataset-worker --enable-pack=false --enable-dag=true --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		out3, _, err := Run(ctx, "singularity datasource inspect dags 1")
		require.NoError(t, err)
		require.Equal(t, out3, out2)
	})
}

func TestPieceDownload(t *testing.T) {
	WithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		temp1 := t.TempDir()
		temp2 := t.TempDir()
		err := os.WriteFile(filepath.Join(temp1, "test1.txt"), generateRandomBytes(10), 0777)
		require.NoError(t, err)
		err = os.WriteFile(filepath.Join(temp1, "test2.txt"), generateRandomBytes(10), 0777)
		require.NoError(t, err)
		err = os.WriteFile(filepath.Join(temp1, "test3.txt"), generateRandomBytes(2_000_000), 0777)
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity dataset create --max-size 1MB --output-dir "+escape(temp2)+" test")
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity datasource add local test "+escape(temp1))
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity datasource daggen 1")
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		_, _, err := Run(ctx, "singularity dataset list-pieces test")
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
			RunNoCapture(ctx2, "singularity run content-provider")
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
			_, _, err = Run(ctx, "singularity download --local-links true -o "+escape(temp3)+" "+pieceCID)
			require.NoError(t, err)
			content, err := os.ReadFile(filepath.Join(temp3, pieceCID+".car"))
			require.NoError(t, err)
			commp := calculateCommp(t, content, 1024*1024)
			require.Equal(t, pieceCID, commp)
		}
	})
}

func TestExtractCar(t *testing.T) {
	WithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tempDir := t.TempDir()
		carDir := t.TempDir()
		extractDir := t.TempDir()
		err := os.WriteFile(filepath.Join(tempDir, "test1.txt"), []byte("hello"), 0644)
		require.NoError(t, err)
		err = os.MkdirAll(filepath.Join(tempDir, "test"), 0755)
		require.NoError(t, err)
		err = os.WriteFile(filepath.Join(tempDir, "test", "test2.txt"), []byte("world"), 0644)
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity dataset create -o "+escape(carDir)+" test")
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity datasource add local test "+escape(tempDir))
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity datasource daggen 1")
		require.NoError(t, err)
		_, _, err = Run(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)
		_, _, err := Run(ctx, "singularity --json datasource inspect path 1")
		require.NoError(t, err)
		root := strings.Split(strings.Split(out, "\n")[4], "\"")[3]
		_, _, err = Run(ctx, "singularity tool extract-car -i "+escape(carDir)+" -o "+escape(extractDir)+" -c "+root)
		require.NoError(t, err)
		content, err := os.ReadFile(filepath.Join(extractDir, "test1.txt"))
		require.NoError(t, err)
		require.Equal(t, "hello", string(content))
		content, err = os.ReadFile(filepath.Join(extractDir, "test", "test2.txt"))
		require.NoError(t, err)
		require.Equal(t, "world", string(content))
	})
}

*/
