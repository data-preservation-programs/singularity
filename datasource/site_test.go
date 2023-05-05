package datasource

import (
	"context"
	"fmt"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSite_Open(t *testing.T) {
	// Test data
	testContent := "This is a test content for the Open function."
	offset := uint64(5)
	length := uint64(10)

	// Create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the "Range" header is present
		rangeHeader := r.Header.Get("Range")
		if rangeHeader == "" {
			http.Error(w, "Missing Range header", http.StatusBadRequest)
			return
		}

		// Parse the "Range" header
		var start, end int
		_, err := fmt.Sscanf(rangeHeader, "bytes=%d-%d", &start, &end)
		if err != nil {
			http.Error(w, "Invalid Range header", http.StatusBadRequest)
			return
		}

		// Validate the range
		if start < 0 || end >= len(testContent) || start > end {
			http.Error(w, "Invalid Range values", http.StatusRequestedRangeNotSatisfiable)
			return
		}

		// Serve the requested range
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusPartialContent)
		w.Write([]byte(testContent)[start : end+1])
	}))
	defer ts.Close()

	// Run the Open function
	site := Site{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	rc, err := site.Open(ctx, ts.URL, offset, length)
	assert.NoError(t, err, "Open function should not return an error")

	// Read the response body
	body, err := ioutil.ReadAll(rc)
	assert.NoError(t, err, "reading response body should not return an error")
	rc.Close()

	// Check the response content
	expectedContent := testContent[offset : offset+length]
	assert.Equal(t, expectedContent, string(body), "response content should match the expected content")
}

func TestSite_Scan(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", "nginx/1.18.0 (Ubuntu)")
		if r.Method == http.MethodHead {
			w.Header().Set("Content-Length", "100")
			w.Header().Set("Last-Modified", "Wed, 03 May 2023 20:36:24 GMT")
			return
		}
		switch r.URL.Path {
		case "/":
			io.WriteString(w, `<html>
<head><title>Index of /</title></head>
<body>
<h1>Index of /</h1><hr><pre><a href="../">../</a>
<a href="dir/">dir/</a>                                               03-May-2023 20:33                   -
<a href="dir2/">dir2/</a>                                              03-May-2023 20:33                   -
<a href="index.nginx-debian.html">index.nginx-debian.html</a>                            03-May-2023 18:04                 612
</pre><hr></body>
</html>
`)
		case "/dir/":
			io.WriteString(w, `<html>
<head><title>Index of /dir/</title></head>
<body>
<h1>Index of /dir/</h1><hr><pre><a href="../">../</a>
<a href="test.txt">test.txt</a>                                           03-May-2023 18:11                   0
<a href="test2.txt">test2.txt</a>                                          03-May-2023 20:33                   0
<a href="test3.txt">test3.txt</a>                                          03-May-2023 20:33                   0
</pre><hr></body>
</html>
`)
		case "/dir2/": io.WriteString(w, `<html>
<head><title>Index of /dir2/</title></head>
<body>
<h1>Index of /dir2/</h1><hr><pre><a href="../">../</a>
<a href="test.txt">test.txt</a>                                           03-May-2023 20:33                   0
<a href="test2.txt">test2.txt</a>                                          03-May-2023 20:36                   5
<a href="test3.txt">test3.txt</a>                                          03-May-2023 20:36                   8
</pre><hr></body>
</html>
`)
		default:
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	}))
	defer ts.Close()

	site := Site{}
	ctx := context.Background()
	entryChan := site.Scan(ctx, ts.URL, ts.URL + "/dir2/test.txt")

	expectedEntries := []Entry{
		{Type: model.URL, Path: ts.URL + "/dir2/test2.txt", Size: 100, LastModified: timePtr(time.Date(2023, 5, 3, 20, 36, 24, 0, time.UTC))},
		{Type: model.URL, Path: ts.URL + "/dir2/test3.txt", Size: 100, LastModified: timePtr(time.Date(2023, 5, 3, 20, 36, 24, 0, time.UTC))},
		{Type: model.URL, Path: ts.URL + "/index.nginx-debian.html", Size: 100, LastModified: timePtr(time.Date(2023, 5, 3, 20, 36, 24, 0, time.UTC))},
	}

	for i, expectedEntry := range expectedEntries {
		entry, ok := <-entryChan
		assert.True(t, ok, "Should receive entry %d", i)
		assert.Equal(t, expectedEntry.Type, entry.Type)
		assert.Equal(t, expectedEntry.Path, entry.Path)
		assert.Equal(t, expectedEntry.Size, entry.Size)
		assert.Equal(t, expectedEntry.LastModified, entry.LastModified)
	}

	moreEntry, ok := <-entryChan
	assert.Falsef(t, ok, "Should not receive more entries, got %v", moreEntry)
}

func timePtr(t time.Time) *time.Time {
	return &t
}

func TestSite_CheckItem(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/exists":
			w.Header().Set("Content-Length", "100")
			w.Header().Set("Last-Modified", "Wed, 03 May 2023 20:36:24 GMT")
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, "Test content")
		case "/not-found":
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}))
	defer ts.Close()

	site := Site{}
	ctx := context.Background()

	// Test with an existing path
	size, lastModified, err := site.CheckItem(ctx, ts.URL+"/exists")
	assert.NoError(t, err)
	assert.Equal(t, uint64(100), size)
	assert.Equal(t, time.Date(2023, 5, 3, 20, 36, 24, 0, time.UTC), *lastModified)

	// Test with a non-existent path
	size, lastModified, err = site.CheckItem(ctx, ts.URL+"/not-found")
	assert.Error(t, err)
	assert.Equal(t, uint64(0), size)
	assert.Nil(t, lastModified)
}
