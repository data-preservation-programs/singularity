package datasource

import (
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResolveSourceType(t *testing.T) {
	tests := []struct {
		name          string
		path          string
		expectedType  model.SourceType
		expectedPath  string
		expectedError bool
	}{
		{
			name:          "S3 path",
			path:          "s3://bucket-name/folder",
			expectedType:  model.S3Path,
			expectedPath:  "s3://bucket-name/folder",
			expectedError: false,
		},
		{
			name:          "Website path",
			path:          "http://example.com",
			expectedType:  model.Website,
			expectedPath:  "http://example.com",
			expectedError: false,
		},
		{
			name:          "Local directory path",
			path:          "/",
			expectedType:  model.Dir,
			expectedPath:  "/",
			expectedError: false,
		},
		{
			name:          "Non-existent path",
			path:          "/path/does/not/exist",
			expectedType:  "",
			expectedPath:  "",
			expectedError: true,
		},
		{
			name:          "Invalid path",
			path:          "::://invalid/path",
			expectedType:  "",
			expectedPath:  "",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualType, actualPath, actualError := ResolveSourceType(tt.path)

			assert.Equal(t, tt.expectedType, actualType)
			assert.Equal(t, tt.expectedPath, actualPath)
			if tt.expectedError {
				assert.Error(t, actualError)
			} else {
				assert.NoError(t, actualError)
			}
		})
	}
}
