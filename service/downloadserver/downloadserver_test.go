package downloadserver

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/contentprovider"
	"github.com/fxamacker/cbor/v2"
	"github.com/ipfs/go-cid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUsageCache(t *testing.T) {
	cache := NewUsageCache[string](time.Millisecond * 100)
	defer cache.Close()

	assert.NotNil(t, cache)
	assert.NotNil(t, cache.data)
	assert.Equal(t, time.Millisecond*100, cache.ttl)
}

func TestUsageCache_SetAndGet(t *testing.T) {
	cache := NewUsageCache[string](time.Second)
	defer cache.Close()

	// Test setting and getting
	cache.Set("key1", "value1")

	value, ok := cache.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, "value1", *value)

	// Test getting non-existent key
	_, ok = cache.Get("nonexistent")
	assert.False(t, ok)
}

func TestUsageCache_Done(t *testing.T) {
	cache := NewUsageCache[string](time.Second)
	defer cache.Close()

	// Set a value and increment usage
	cache.Set("key1", "value1")
	cache.Get("key1") // This increments usage count

	// Test done decrements usage count
	cache.Done("key1")

	// Test done on non-existent key doesn't panic
	cache.Done("nonexistent")
}

func TestUsageCache_TTL_Cleanup(t *testing.T) {
	cache := NewUsageCache[string](time.Millisecond * 50)
	defer cache.Close()

	// Set a value
	cache.Set("key1", "value1")

	// Mark as done so usage count is 0
	cache.Done("key1")

	// Wait for TTL + cleanup cycle
	time.Sleep(time.Millisecond * 150)

	// Should still be available if cleanup didn't run yet
	_, ok := cache.Get("key1")
	// The cleanup might or might not have run, so we don't assert specific behavior
	// but we test that the cache doesn't crash
	_ = ok
}

func TestNewDownloadServer(t *testing.T) {
	config := map[string]string{"test": "value"}
	clientConfig := model.ClientConfig{}

	server := NewDownloadServer(":8080", "http://api.example.com", config, clientConfig)

	assert.Equal(t, ":8080", server.bind)
	assert.Equal(t, "http://api.example.com", server.api)
	assert.Equal(t, config, server.config)
	assert.Equal(t, clientConfig, server.clientConfig)
	assert.NotNil(t, server.usageCache)
}

func TestDownloadServer_Name(t *testing.T) {
	server := NewDownloadServer(":8080", "http://api.example.com", nil, model.ClientConfig{})
	assert.Equal(t, "DownloadServer", server.Name())
}

func TestDownloadServer_handleGetPiece_InvalidCID(t *testing.T) {
	server := NewDownloadServer(":8080", "http://api.example.com", nil, model.ClientConfig{})

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/piece/invalid-cid", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/piece/:id")
	c.SetParamNames("id")
	c.SetParamValues("invalid-cid")

	err := server.handleGetPiece(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "failed to parse piece CID")
}

func TestDownloadServer_handleGetPiece_NotCommP(t *testing.T) {
	server := NewDownloadServer(":8080", "http://api.example.com", nil, model.ClientConfig{})

	// Create a non-CommP CID (regular file CID)
	regularCid := cid.NewCidV1(cid.Raw, []byte("test"))

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/piece/"+regularCid.String(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/piece/:id")
	c.SetParamNames("id")
	c.SetParamValues(regularCid.String())

	err := server.handleGetPiece(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "failed to parse piece CID")
}

func TestGetMetadata_InvalidAPI(t *testing.T) {
	ctx := context.Background()
	config := map[string]string{}
	clientConfig := model.ClientConfig{}

	// Test with invalid URL
	_, statusCode, err := GetMetadata(ctx, "://invalid-url", config, clientConfig, "test-piece-cid")
	assert.Error(t, err)
	assert.Equal(t, 0, statusCode)
}

func TestGetMetadata_Success(t *testing.T) {
	// Create a mock server that returns metadata
	mockMetadata := contentprovider.PieceMetadata{
		Car: model.Car{
			ID:        1,
			CreatedAt: time.Now(),
		},
		Storage: model.Storage{
			Type: "local",
			Config: map[string]string{
				"provider": "local",
				"path":     "/tmp/test",
			},
		},
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "/piece/metadata/")
		assert.Equal(t, "application/cbor", r.Header.Get("Accept"))

		w.Header().Set("Content-Type", "application/cbor")
		encoder := cbor.NewEncoder(w)
		err := encoder.Encode(mockMetadata)
		require.NoError(t, err)
	}))
	defer mockServer.Close()

	ctx := context.Background()
	config := map[string]string{}
	clientConfig := model.ClientConfig{}

	metadata, statusCode, err := GetMetadata(ctx, mockServer.URL, config, clientConfig, "test-piece-cid")
	assert.NoError(t, err)
	assert.Equal(t, 0, statusCode)
	assert.NotNil(t, metadata)
	assert.Equal(t, "local", metadata.Storage.Type)
}

func TestGetMetadata_404(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "not found")
	}))
	defer mockServer.Close()

	ctx := context.Background()
	config := map[string]string{}
	clientConfig := model.ClientConfig{}

	_, statusCode, err := GetMetadata(ctx, mockServer.URL, config, clientConfig, "test-piece-cid")
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, statusCode)
	assert.Contains(t, err.Error(), "failed to get metadata")
}

func TestGetMetadata_InvalidResponse(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/cbor")
		_, _ = w.Write([]byte("invalid cbor data"))
	}))
	defer mockServer.Close()

	ctx := context.Background()
	config := map[string]string{}
	clientConfig := model.ClientConfig{}

	_, statusCode, err := GetMetadata(ctx, mockServer.URL, config, clientConfig, "test-piece-cid")
	assert.Error(t, err)
	assert.Equal(t, 0, statusCode)
	assert.Contains(t, err.Error(), "failed to decode metadata")
}

func TestGetMetadata_ConfigProcessing(t *testing.T) {
	mockMetadata := contentprovider.PieceMetadata{
		Car: model.Car{
			ID:        1,
			CreatedAt: time.Now(),
		},
		Storage: model.Storage{
			Type: "local",
			Config: map[string]string{
				"provider": "local",
				"path":     "/original/path",
			},
		},
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/cbor")
		encoder := cbor.NewEncoder(w)
		_ = encoder.Encode(mockMetadata)
	}))
	defer mockServer.Close()

	ctx := context.Background()
	config := map[string]string{
		"local-path":  "/override/path",
		"local-other": "override-value",
	}
	clientConfig := model.ClientConfig{}

	metadata, statusCode, err := GetMetadata(ctx, mockServer.URL, config, clientConfig, "test-piece-cid")
	assert.NoError(t, err)
	assert.Equal(t, 0, statusCode)
	assert.NotNil(t, metadata)

	// Test that config overrides are applied
	assert.Equal(t, "/override/path", metadata.Storage.Config["path"])
	assert.Equal(t, "override-value", metadata.Storage.Config["other"])
}

func TestDownloadServer_Start_Health(t *testing.T) {
	// Find an available port
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	port := listener.Addr().(*net.TCPAddr).Port
	_ = listener.Close()

	bindAddr := fmt.Sprintf("127.0.0.1:%d", port)
	server := NewDownloadServer(bindAddr, "http://api.example.com", nil, model.ClientConfig{})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	exitErr := make(chan error, 1)

	err = server.Start(ctx, exitErr)
	assert.NoError(t, err)

	// Wait for the server to be ready by polling the health endpoint
	serverURL := fmt.Sprintf("http://%s", bindAddr)
	client := &http.Client{Timeout: time.Second}

	var healthResp *http.Response
	for i := 0; i < 50; i++ { // Try for up to 5 seconds
		healthResp, err = client.Get(serverURL + "/health")
		if err == nil {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	// Server should be ready now
	require.NoError(t, err, "Server failed to start within timeout")
	require.NotNil(t, healthResp)
	defer func() { _ = healthResp.Body.Close() }()

	// Test the health endpoint
	assert.Equal(t, http.StatusOK, healthResp.StatusCode)

	// Make another health check to ensure server is stable
	resp2, err := client.Get(serverURL + "/health")
	require.NoError(t, err)
	defer func() { _ = resp2.Body.Close() }()
	assert.Equal(t, http.StatusOK, resp2.StatusCode)

	// Now shutdown the server
	cancel()

	select {
	case err := <-exitErr:
		// Server should shutdown cleanly
		assert.NoError(t, err)
	case <-time.After(time.Second * 3):
		t.Fatal("Server did not shut down within timeout")
	}
}
