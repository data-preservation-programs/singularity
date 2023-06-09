package cmd

import (
	"context"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

func generateRandomBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}

func testWithAllBackendWithoutReset(t *testing.T, testFunc func(ctx context.Context, t *testing.T, db *gorm.DB)) {
	testWithAllBackendWithResetArg(t, testFunc, false)
}

func testWithAllBackendWithResetArg(t *testing.T, testFunc func(ctx context.Context, t *testing.T, db *gorm.DB), reset bool) {
	temp := t.TempDir()
	backends := [][2]string {
		{"sqlite", "sqlite:" + temp + "/singularity.db"},
		//{"mysql", "mysql://root:password@tcp(localhost:3306)/singularity?parseTime=true"},
		//{"postgres" , "postgres://postgres:password@localhost:5432/singularity?sslmode=disable"},
	}
	for _, backend := range backends {
		os.Setenv("DATABASE_CONNECTION_STRING", backend[1])
		if reset {
			_, _, err := RunArgsInTest(context.Background(), "singularity admin reset")
			assert.NoError(t, err)
		}
		db, err := database.Open(backend[1], &gorm.Config{})
		assert.NoError(t, err)
		ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Minute)
		defer cancel()
		t.Run(backend[0], func (t *testing.T) {
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

func TestDatasourceCrud(t *testing.T) {
	testWithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		temp := t.TempDir()
		os.Mkdir(temp + "/sub", 0777)
		os.WriteFile(temp + "/sub/test.txt", []byte("hello world"), 0777)
		_, _, err := RunArgsInTest(ctx, "singularity dataset create test")
		assert.NoError(t, err)
		out, _, err := RunArgsInTest(ctx, "singularity datasource add local test " + temp)
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

func TestDatasourcePackingCorrectness(t *testing.T) {
	testWithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		temp := t.TempDir()
		carDir := t.TempDir()
		// multiple nested folder
		os.MkdirAll(temp + "/sub1/sub2/sub3/sub4", 0777)
		// dynamic directory with 10k files
		for i := 0; i < 10_000; i++ {
			os.WriteFile(temp + "/sub1/sub2/sub3/sub4/test" + strconv.Itoa(i) + ".txt", generateRandomBytes(10), 0777)
		}
		// dynamic directory with 10k folders
		for i := 0; i < 10_000; i++ {
			os.MkdirAll(temp + "/" + strconv.Itoa(i), 0777)
			os.WriteFile(temp + "/" + strconv.Itoa(i) + "/test" + strconv.Itoa(i) + ".txt", generateRandomBytes(10), 0777)
		}
		// file of large size
		os.WriteFile(temp + "/test1.txt", generateRandomBytes(10000), 0777)
		// file of empty size
		os.WriteFile(temp + "/test2.txt", []byte{}, 0777)
		_, _, err := RunArgsInTest(ctx, "singularity dataset create --max-size 1000 -o " + carDir + " test")
		assert.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity datasource add local --delete-after-export=true test " + temp)
		assert.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --exit-on-complete=true")
		assert.NoError(t, err)
	})
}


func TestDatasourceRescan(t *testing.T) {
	testWithAllBackend(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		temp := t.TempDir()
		os.Mkdir(temp + "/sub", 0777)
		os.WriteFile(temp + "/sub/test1.txt", generateRandomBytes(10), 0777)
		os.WriteFile(temp + "/sub/test2.txt", generateRandomBytes(100), 0777)
		os.WriteFile(temp + "/sub/test3.txt", generateRandomBytes(1000), 0777)
		os.WriteFile(temp + "/sub/test4.txt", generateRandomBytes(10000), 0777)
		_, _, err := RunArgsInTest(ctx, "singularity dataset create --max-size 1000 test")
		assert.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity datasource add local test " + temp)
		assert.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --enable-pack=false --enable-dag=false --exit-on-complete=true")
		assert.NoError(t, err)
		out, _, err := RunArgsInTest(ctx, "singularity datasource inspect chunks 1")
		assert.NoError(t, err)
		assert.Contains(t, out, "ready")
		// We should get 15 chunks
		assert.Contains(t, out, "15")
		os.WriteFile(temp + "/sub/test5.txt", generateRandomBytes(10000), 0777)
		_, _, err = RunArgsInTest(ctx, "singularity datasource rescan 1")
		assert.NoError(t, err)
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --enable-pack=false --enable-dag=false --exit-on-complete=true")
		assert.NoError(t, err)
		out, _, err = RunArgsInTest(ctx, "singularity datasource inspect chunks 1")
		assert.NoError(t, err)
		// We should get 29 chunks
		assert.Contains(t, out, "29")
		assert.NotContains(t, out, "30")
		_, _, err = RunArgsInTest(ctx, "singularity run dataset-worker --enable-pack=true --enable-dag=false --exit-on-complete=true")
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
		out, _, err = RunArgsInTest(ctx, "singularity datasource inspect dir 1")
		assert.NoError(t, err)
		assert.Contains(t, out, "sub")
		out, _, err = RunArgsInTest(ctx, "singularity datasource inspect dir 1 sub/")
		assert.NoError(t, err)
		assert.Contains(t, out, "sub/test1.txt")
		assert.Contains(t, out, "sub/test3.txt")
	})
}
