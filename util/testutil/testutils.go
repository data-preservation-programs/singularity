package testutil

import (
	"context"
	"crypto/rand"
	"io"
	rand2 "math/rand"
	"net"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/boxo/util"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

const pattern = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateFixedBytes(length int) []byte {
	patternLen := len(pattern)
	result := make([]byte, length)
	for i := range length {
		result[i] = pattern[i%patternLen]
	}
	return result
}

func GenerateRandomBytes(n int) []byte {
	b := make([]byte, n)
	//nolint:errcheck
	_, _ = rand.Read(b)
	return b
}

func RandomLetterString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"

	b := make([]byte, length)
	for i := range b {
		//nolint:gosec
		b[i] = charset[rand2.Intn(len(charset))]
	}
	return string(b)
}

// GenerateUniqueName creates a unique name for testing by combining a prefix with a random suffix
func GenerateUniqueName(prefix string) string {
	return prefix + "-" + RandomLetterString(8) + "-" + RandomLetterString(4)
}

func GetFileTimestamp(t *testing.T, path string) int64 {
	t.Helper()
	info, err := os.Stat(path)
	require.NoError(t, err)
	return info.ModTime().UnixNano()
}

var TestCid = cid.NewCidV1(cid.Raw, util.Hash([]byte("test")))

const TestWalletActorID = "f0808055"

const TestWalletAddr = "f1fib3pv7jua2ockdugtz7viz3cyy6lkhh7rfx3sa"

const TestPrivateKeyHex = "7b2254797065223a22736563703235366b31222c22507269766174654b6579223a226b35507976337148327349586343595a58594f5775453149326e32554539436861556b6c4e36695a5763453d227d"

func EscapePath(p string) string {
	return "'" + strings.ReplaceAll(p, `\`, `\\`) + "'"
}

func getTestDB(t *testing.T, dialect string) (db *gorm.DB, closer io.Closer, connStr string) {
	t.Helper()
	var err error
	if dialect == "sqlite" {
		connStr = "sqlite:" + t.TempDir() + "/singularity.db"
		db, closer, err = database.OpenWithLogger(connStr)
		require.NoError(t, err)
		return
	}
	dbName := RandomLetterString(6)
	var opError *net.OpError
	switch dialect {
	case "mysql":
		connStr = "mysql://singularity:singularity@tcp(localhost:3306)/singularity?parseTime=true"
	case "postgres":
		connStr = "postgres://singularity:singularity@localhost:5432/singularity?sslmode=disable"
	default:
		require.Fail(t, "Unsupported dialect: "+dialect)
	}
	var db1 *gorm.DB
	var closer1 io.Closer
	db1, closer1, err = database.OpenWithLogger(connStr)
	if errors.As(err, &opError) {
		t.Logf("Database %s not available: %v", dialect, err)
		return nil, nil, ""
	}
	if err != nil {
		t.Logf("Failed to connect to %s database: %v", dialect, err)
		return nil, nil, ""
	}
	err = db1.Exec("CREATE DATABASE " + dbName + "").Error
	if err != nil {
		t.Logf("Failed to create test database %s: %v", dbName, err)
		_ = closer1.Close()
		return nil, nil, ""
	}
	connStr = strings.ReplaceAll(connStr, "singularity?", dbName+"?")
	var closer2 io.Closer
	db, closer2, err = database.OpenWithLogger(connStr)
	if err != nil {
		t.Logf("Failed to connect to test database %s: %v", dbName, err)
		db1.Exec("DROP DATABASE " + dbName + "")
		_ = closer1.Close()
		return nil, nil, ""
	}
	closer = CloserFunc(func() error {
		if closer2 != nil {
			_ = closer2.Close()
		}
		if db1 != nil {
			db1.Exec("DROP DATABASE " + dbName + "")
		}
		if closer1 != nil {
			return closer1.Close()
		}
		return nil
	})
	return
}

func One(t *testing.T, testFunc func(ctx context.Context, t *testing.T, db *gorm.DB)) {
	t.Helper()
	backend := SupportedTestDialects[0]
	doOne(t, backend, testFunc)
}

func OneWithoutReset(t *testing.T, testFunc func(ctx context.Context, t *testing.T, db *gorm.DB)) {
	t.Helper()
	backend := SupportedTestDialects[0]
	db, closer, connStr := getTestDB(t, backend)
	if db == nil {
		t.Skip("Skip " + backend + " - database not available")
		return
	}
	defer func() { _ = closer.Close() }()
	t.Setenv("DATABASE_CONNECTION_STRING", connStr)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	db = db.WithContext(ctx)

	t.Run(backend, func(t *testing.T) {
		testFunc(ctx, t, db)
	})
}

func doOne(t *testing.T, backend string, testFunc func(ctx context.Context, t *testing.T, db *gorm.DB)) {
	t.Helper()
	db, closer, connStr := getTestDB(t, backend)
	if db == nil {
		t.Skip("Skip " + backend + " - database not available")
		return
	}
	defer func() { _ = closer.Close() }()
	t.Setenv("DATABASE_CONNECTION_STRING", connStr)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	db = db.WithContext(ctx)

	err := model.GetMigrator(db).Migrate()
	require.NoError(t, err)

	t.Run(backend, func(t *testing.T) {
		testFunc(ctx, t, db)
	})
}

func All(t *testing.T, testFunc func(ctx context.Context, t *testing.T, db *gorm.DB)) {
	t.Helper()
	for _, backend := range SupportedTestDialects {
		doOne(t, backend, testFunc)
	}
}

type CloserFunc func() error

func (f CloserFunc) Close() error {
	return f()
}

// MockLotusClient provides a mock implementation of jsonrpc.RPCClient for testing
type MockLotusClient struct {
	responses map[string]interface{}
	errors    map[string]error
}

// NewMockLotusClient creates a new mock Lotus client for testing
func NewMockLotusClient() *MockLotusClient {
	return &MockLotusClient{
		responses: make(map[string]interface{}),
		errors:    make(map[string]error),
	}
}

// SetResponse sets a mock response for a specific method
func (m *MockLotusClient) SetResponse(method string, response interface{}) {
	m.responses[method] = response
}

// SetError sets a mock error for a specific method
func (m *MockLotusClient) SetError(method string, err error) {
	m.errors[method] = err
}

// CallFor implements jsonrpc.RPCClient interface
func (m *MockLotusClient) CallFor(ctx context.Context, out interface{}, method string, params ...interface{}) error {
	if err, exists := m.errors[method]; exists {
		return err
	}
	if response, exists := m.responses[method]; exists {
		// Simple type assertion for common response types
		switch v := out.(type) {
		case *string:
			if str, ok := response.(string); ok {
				*v = str
			}
		}
		return nil
	}
	return errors.New("mock method not configured: " + method)
}

// Call implements jsonrpc.RPCClient interface
func (m *MockLotusClient) Call(ctx context.Context, method string, params ...interface{}) (*jsonrpc.RPCResponse, error) {
	if err, exists := m.errors[method]; exists {
		return nil, err
	}
	if response, exists := m.responses[method]; exists {
		return &jsonrpc.RPCResponse{Result: response}, nil
	}
	return nil, errors.New("mock method not configured: " + method)
}

// CallBatch implements jsonrpc.RPCClient interface
func (m *MockLotusClient) CallBatch(ctx context.Context, requests jsonrpc.RPCRequests) (jsonrpc.RPCResponses, error) {
	return nil, errors.New("CallBatch not implemented in mock")
}

// CallRaw implements jsonrpc.RPCClient interface
func (m *MockLotusClient) CallRaw(ctx context.Context, request *jsonrpc.RPCRequest) (*jsonrpc.RPCResponse, error) {
	return m.Call(ctx, request.Method, request.Params)
}

// CallBatchRaw implements jsonrpc.RPCClient interface
func (m *MockLotusClient) CallBatchRaw(ctx context.Context, requests jsonrpc.RPCRequests) (jsonrpc.RPCResponses, error) {
	return nil, errors.New("CallBatchRaw not implemented in mock")
}
