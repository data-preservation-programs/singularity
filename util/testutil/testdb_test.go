package testutil

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestTestDB(t *testing.T) {
	All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Test that database connection works
		assert.NotNil(t, db)

		// Test that context is properly set
		assert.NotNil(t, ctx)

		// Test basic database operation
		var result int
		err := db.Raw("SELECT 1").Scan(&result).Error
		require.NoError(t, err)
		assert.Equal(t, 1, result)
	})
}

func TestOne(t *testing.T) {
	One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Test that we get a valid database connection
		assert.NotNil(t, db)
		assert.NotNil(t, ctx)

		// Test context timeout
		deadline, ok := ctx.Deadline()
		assert.True(t, ok)
		assert.True(t, deadline.After(time.Now()))
	})
}

func TestOneWithoutReset(t *testing.T) {
	OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Test that we get a valid database connection
		assert.NotNil(t, db)
		assert.NotNil(t, ctx)

		// Test that database operations work
		var count int64
		err := db.Raw("SELECT COUNT(*) FROM information_schema.tables").Scan(&count).Error
		if err != nil {
			// Might fail on SQLite, try a different query
			err = db.Raw("SELECT 1").Scan(&count).Error
			require.NoError(t, err)
		}
	})
}

func TestGenerateFixedBytes(t *testing.T) {
	// Test with various lengths
	testCases := []int{0, 1, 10, 26, 62, 100}

	for _, length := range testCases {
		result := GenerateFixedBytes(length)
		assert.Equal(t, length, len(result))

		// Test that result is deterministic
		result2 := GenerateFixedBytes(length)
		assert.Equal(t, result, result2)

		// Test that pattern is followed for non-zero lengths
		if length > 0 {
			assert.True(t, result[0] >= 'a' && result[0] <= 'z' ||
				result[0] >= 'A' && result[0] <= 'Z' ||
				result[0] >= '0' && result[0] <= '9')
		}
	}
}

func TestGenerateRandomBytesVariousLengths(t *testing.T) {
	// Test with various lengths
	testCases := []int{0, 1, 10, 100}

	for _, length := range testCases {
		result := GenerateRandomBytes(length)
		assert.Equal(t, length, len(result))

		// Test that results are different (very high probability)
		if length > 0 {
			result2 := GenerateRandomBytes(length)
			assert.NotEqual(t, result, result2)
		}
	}
}

func TestRandomLetterString(t *testing.T) {
	// Test with various lengths
	testCases := []int{0, 1, 5, 26, 100}

	for _, length := range testCases {
		result := RandomLetterString(length)
		assert.Equal(t, length, len(result))

		// Test that all characters are lowercase letters
		for _, char := range result {
			assert.True(t, char >= 'a' && char <= 'z')
		}

		// Test that results are different (very high probability)
		if length > 0 {
			result2 := RandomLetterString(length)
			// With random generation, there's a tiny chance they're the same
			// but for reasonable lengths it's extremely unlikely
			if length > 3 {
				assert.NotEqual(t, result, result2)
			}
		}
	}
}

func TestEscapePath(t *testing.T) {
	testCases := map[string]string{
		"simple":                  "'simple'",
		"path/with/slashes":       "'path/with/slashes'",
		"path\\with\\backslashes": "'path\\\\with\\\\backslashes'",
		"":                        "''",
		"path with spaces":        "'path with spaces'",
	}

	for input, expected := range testCases {
		result := EscapePath(input)
		assert.Equal(t, expected, result)
	}
}

func TestConstants(t *testing.T) {
	// Test that constants are properly defined
	assert.NotEmpty(t, TestCid.String())
	assert.NotEmpty(t, TestWalletAddr)
	assert.NotEmpty(t, TestPrivateKeyHex)

	// Test wallet address format
	assert.True(t, len(TestWalletAddr) > 0)
	assert.True(t, TestWalletAddr[0] == 'f')

	// Test private key hex format
	assert.True(t, len(TestPrivateKeyHex) > 0)
}
