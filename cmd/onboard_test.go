package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestValidateStorageType(t *testing.T) {
	tests := []struct {
		name        string
		storageType string
		prefix      string
		wantErr     bool
	}{
		{
			name:        "empty storage type should pass",
			storageType: "",
			prefix:      "test",
			wantErr:     false,
		},
		{
			name:        "valid storage type should pass",
			storageType: "local",
			prefix:      "test",
			wantErr:     false,
		},
		{
			name:        "invalid storage type should fail",
			storageType: "invalid-type",
			prefix:      "test",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateStorageType(tt.storageType, tt.prefix)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateStorageConfig(t *testing.T) {
	tests := []struct {
		name      string
		configStr string
		prefix    string
		wantErr   bool
	}{
		{
			name:      "valid JSON config should pass",
			configStr: `{"key1": "value1", "key2": "value2"}`,
			prefix:    "test",
			wantErr:   false,
		},
		{
			name:      "invalid JSON should fail",
			configStr: `{"key1": "value1", "key2":}`,
			prefix:    "test",
			wantErr:   true,
		},
		{
			name:      "empty key should fail",
			configStr: `{"": "value1", "key2": "value2"}`,
			prefix:    "test",
			wantErr:   true,
		},
		{
			name:      "empty value should fail",
			configStr: `{"key1": "", "key2": "value2"}`,
			prefix:    "test",
			wantErr:   true,
		},
		{
			name:      "empty config should pass",
			configStr: `{}`,
			prefix:    "test",
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateStorageConfig(tt.configStr, tt.prefix)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetCustomStorageConfig(t *testing.T) {
	tests := []struct {
		name           string
		sourceConfig   string
		outputConfig   string
		storageContext string
		expected       map[string]string
	}{
		{
			name:           "source config should be returned for source context",
			sourceConfig:   `{"key1": "value1"}`,
			outputConfig:   `{"key2": "value2"}`,
			storageContext: "source",
			expected:       map[string]string{"key1": "value1"},
		},
		{
			name:           "output config should be returned for output context",
			sourceConfig:   `{"key1": "value1"}`,
			outputConfig:   `{"key2": "value2"}`,
			storageContext: "output",
			expected:       map[string]string{"key2": "value2"},
		},
		{
			name:           "nil should be returned for unknown context",
			sourceConfig:   `{"key1": "value1"}`,
			outputConfig:   `{"key2": "value2"}`,
			storageContext: "unknown",
			expected:       nil,
		},
		{
			name:           "nil should be returned for empty config",
			sourceConfig:   "",
			outputConfig:   "",
			storageContext: "source",
			expected:       nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &cli.App{
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "source-config"},
					&cli.StringFlag{Name: "output-config"},
				},
			}

			args := []string{"app"}
			if tt.sourceConfig != "" {
				args = append(args, "--source-config", tt.sourceConfig)
			}
			if tt.outputConfig != "" {
				args = append(args, "--output-config", tt.outputConfig)
			}

			var c *cli.Context
			app.Action = func(ctx *cli.Context) error {
				c = ctx
				return nil
			}

			err := app.Run(args)
			assert.NoError(t, err)

			result := getCustomStorageConfig(c, tt.storageContext)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestParseS3Config(t *testing.T) {
	tests := []struct {
		name     string
		flags    map[string]interface{}
		expected map[string]string
	}{
		{
			name: "all S3 flags should be parsed",
			flags: map[string]interface{}{
				"s3-access-key-id":     "test-key",
				"s3-secret-access-key": "test-secret",
				"s3-region":            "us-east-1",
				"s3-endpoint":          "https://s3.amazonaws.com",
				"s3-env-auth":          true,
			},
			expected: map[string]string{
				"access_key_id":     "test-key",
				"secret_access_key": "test-secret",
				"region":            "us-east-1",
				"endpoint":          "https://s3.amazonaws.com",
				"env_auth":          "true",
			},
		},
		{
			name:     "empty flags should return empty config",
			flags:    map[string]interface{}{},
			expected: map[string]string{},
		},
		{
			name: "env-auth false should not set env_auth",
			flags: map[string]interface{}{
				"s3-access-key-id": "test-key",
				"s3-env-auth":      false,
			},
			expected: map[string]string{
				"access_key_id": "test-key",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &cli.App{
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "s3-access-key-id"},
					&cli.StringFlag{Name: "s3-secret-access-key"},
					&cli.StringFlag{Name: "s3-region"},
					&cli.StringFlag{Name: "s3-endpoint"},
					&cli.BoolFlag{Name: "s3-env-auth"},
				},
			}

			args := []string{"app"}
			for flag, value := range tt.flags {
				switch v := value.(type) {
				case string:
					args = append(args, "--"+flag, v)
				case bool:
					if v {
						args = append(args, "--"+flag)
					}
				}
			}

			var c *cli.Context
			app.Action = func(ctx *cli.Context) error {
				c = ctx
				return nil
			}

			err := app.Run(args)
			assert.NoError(t, err)

			result := parseS3Config(c)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestParseGCSConfig(t *testing.T) {
	tests := []struct {
		name     string
		flags    map[string]interface{}
		expected map[string]string
	}{
		{
			name: "all GCS flags should be parsed",
			flags: map[string]interface{}{
				"gcs-service-account-file":        "/path/to/service-account.json",
				"gcs-service-account-credentials": `{"type": "service_account"}`,
				"gcs-project-id":                  "test-project",
				"gcs-env-auth":                    true,
			},
			expected: map[string]string{
				"service_account_file":        "/path/to/service-account.json",
				"service_account_credentials": `{"type": "service_account"}`,
				"project_number":              "test-project",
				"env_auth":                    "true",
			},
		},
		{
			name:     "empty flags should return empty config",
			flags:    map[string]interface{}{},
			expected: map[string]string{},
		},
		{
			name: "env-auth false should not set env_auth",
			flags: map[string]interface{}{
				"gcs-project-id": "test-project",
				"gcs-env-auth":   false,
			},
			expected: map[string]string{
				"project_number": "test-project",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &cli.App{
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "gcs-service-account-file"},
					&cli.StringFlag{Name: "gcs-service-account-credentials"},
					&cli.StringFlag{Name: "gcs-project-id"},
					&cli.BoolFlag{Name: "gcs-env-auth"},
				},
			}

			args := []string{"app"}
			for flag, value := range tt.flags {
				switch v := value.(type) {
				case string:
					args = append(args, "--"+flag, v)
				case bool:
					if v {
						args = append(args, "--"+flag)
					}
				}
			}

			var c *cli.Context
			app.Action = func(ctx *cli.Context) error {
				c = ctx
				return nil
			}

			err := app.Run(args)
			assert.NoError(t, err)

			result := parseGCSConfig(c)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestValidateOnboardInputs(t *testing.T) {
	tests := []struct {
		name    string
		flags   map[string]interface{}
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid minimal input should pass",
			flags: map[string]interface{}{
				"name":              "test-prep",
				"source":            []string{"/tmp/test"},
				"auto-create-deals": false,
				"source-type":       "local",
				"output-type":       "local",
			},
			wantErr: false,
		},
		{
			name: "missing name should fail",
			flags: map[string]interface{}{
				"source":            []string{"/tmp/test"},
				"auto-create-deals": false,
			},
			wantErr: true,
			errMsg:  "preparation name is required",
		},
		{
			name: "missing source should fail",
			flags: map[string]interface{}{
				"name":              "test-prep",
				"auto-create-deals": false,
			},
			wantErr: true,
			errMsg:  "at least one source path is required",
		},
		{
			name: "invalid source storage type should fail",
			flags: map[string]interface{}{
				"name":              "test-prep",
				"source":            []string{"/tmp/test"},
				"source-type":       "invalid-type",
				"auto-create-deals": false,
			},
			wantErr: true,
			errMsg:  "source storage type 'invalid-type' is not supported",
		},
		{
			name: "invalid source config should fail",
			flags: map[string]interface{}{
				"name":              "test-prep",
				"source":            []string{"/tmp/test"},
				"source-type":       "local",
				"source-config":     `{"key": }`,
				"auto-create-deals": false,
			},
			wantErr: true,
			errMsg:  "invalid JSON format for source-config",
		},
		{
			name: "auto-create-deals without provider should fail",
			flags: map[string]interface{}{
				"name":              "test-prep",
				"source":            []string{"/tmp/test"},
				"source-type":       "local",
				"auto-create-deals": true,
			},
			wantErr: true,
			errMsg:  "deal provider is required when auto-create-deals is enabled",
		},
		{
			name: "auto-create-deals with valid config should pass",
			flags: map[string]interface{}{
				"name":              "test-prep",
				"source":            []string{"/tmp/test"},
				"source-type":       "local",
				"auto-create-deals": true,
				"deal-provider":     "f01234",
				"deal-duration":     "535h",
				"deal-price-per-gb": 0.1,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock CLI context with the test flags
			app := &cli.App{
				Flags: append(OnboardCmd.Flags, []cli.Flag{
					&cli.StringFlag{Name: "name"},
					&cli.StringSliceFlag{Name: "source"},
					&cli.BoolFlag{Name: "auto-create-deals"},
					&cli.StringFlag{Name: "source-type"},
					&cli.StringFlag{Name: "output-type"},
					&cli.StringFlag{Name: "source-config"},
					&cli.StringFlag{Name: "deal-provider"},
					&cli.DurationFlag{Name: "deal-duration"},
					&cli.Float64Flag{Name: "deal-price-per-gb"},
				}...),
			}

			args := []string{"test"}
			for flag, value := range tt.flags {
				switch v := value.(type) {
				case string:
					args = append(args, "--"+flag, v)
				case []string:
					for _, item := range v {
						args = append(args, "--"+flag, item)
					}
				case bool:
					if v {
						args = append(args, "--"+flag)
					}
				case float64:
					args = append(args, "--"+flag, fmt.Sprintf("%f", v))
				}
			}

			var c *cli.Context
			app.Action = func(ctx *cli.Context) error {
				c = ctx
				return nil
			}

			err := app.Run(append([]string{"app"}, args...))
			assert.NoError(t, err)

			// Test the validation function
			err = validateOnboardInputs(c)
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestLogInsecureClientConfigWarning(t *testing.T) {
	tests := []struct {
		name            string
		insecureSkipTLS bool
		expectWarning   bool
	}{
		{
			name:            "insecure skip verify should trigger warning",
			insecureSkipTLS: true,
			expectWarning:   true,
		},
		{
			name:            "secure config should not trigger warning",
			insecureSkipTLS: false,
			expectWarning:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &cli.App{
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "client-insecure-skip-verify"},
				},
			}

			args := []string{"test"}
			if tt.insecureSkipTLS {
				args = append(args, "--client-insecure-skip-verify")
			}

			var c *cli.Context
			app.Action = func(ctx *cli.Context) error {
				c = ctx
				return nil
			}

			err := app.Run(append([]string{"app"}, args...))
			assert.NoError(t, err)

			// This test just ensures the function runs without error
			// In a real test, you'd want to capture stdout to verify the warning message
			assert.NotPanics(t, func() {
				logInsecureClientConfigWarning(c)
			})
		})
	}
}

func TestGetProviderDefaults(t *testing.T) {
	tests := []struct {
		name        string
		storageType string
		provider    string
		wantEmpty   bool
	}{
		{
			name:        "invalid storage type should return empty",
			storageType: "invalid-type",
			provider:    "aws",
			wantEmpty:   true,
		},
		{
			name:        "valid storage type should return defaults",
			storageType: "local",
			provider:    "",
			wantEmpty:   false,
		},
		{
			name:        "s3 with aws provider should return defaults",
			storageType: "s3",
			provider:    "aws",
			wantEmpty:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defaults := getProviderDefaults(tt.storageType, tt.provider)
			if tt.wantEmpty {
				assert.Empty(t, defaults)
			} else {
				// We expect some defaults to be set for valid storage types
				assert.NotNil(t, defaults)
			}
		})
	}
}

func TestMergeStorageConfigWithDefaults(t *testing.T) {
	tests := []struct {
		name         string
		storageType  string
		provider     string
		customConfig map[string]string
		expectMerged bool
	}{
		{
			name:        "custom config should override defaults",
			storageType: "local",
			provider:    "",
			customConfig: map[string]string{
				"encoding": "custom-encoding",
			},
			expectMerged: true,
		},
		{
			name:         "empty custom config should return defaults",
			storageType:  "local",
			provider:     "",
			customConfig: map[string]string{},
			expectMerged: true,
		},
		{
			name:         "invalid storage type should return custom config only",
			storageType:  "invalid-type",
			provider:     "",
			customConfig: map[string]string{"key": "value"},
			expectMerged: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merged := mergeStorageConfigWithDefaults(tt.storageType, tt.provider, tt.customConfig)
			assert.NotNil(t, merged)

			// Check that custom config values are present
			for key, value := range tt.customConfig {
				assert.Equal(t, value, merged[key])
			}

			if tt.expectMerged {
				// For valid storage types, we expect the merged config to potentially have more keys
				// than just the custom config
				assert.GreaterOrEqual(t, len(merged), len(tt.customConfig))
			}
		})
	}
}

func TestParseGenericStorageConfig(t *testing.T) {
	tests := []struct {
		name        string
		storageType string
		flags       map[string]interface{}
		expectKeys  []string
	}{
		{
			name:        "invalid storage type should return empty config",
			storageType: "invalid-type",
			flags:       map[string]interface{}{},
			expectKeys:  []string{},
		},
		{
			name:        "local storage should parse encoding flag",
			storageType: "local",
			flags: map[string]interface{}{
				"local-encoding": "base64",
			},
			expectKeys: []string{"encoding"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock CLI context with the specified flags
			app := &cli.App{
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "local-encoding"},
				},
			}

			args := []string{"app"}
			for flag, value := range tt.flags {
				switch v := value.(type) {
				case string:
					args = append(args, "--"+flag, v)
				case bool:
					if v {
						args = append(args, "--"+flag)
					}
				}
			}

			var c *cli.Context
			app.Action = func(ctx *cli.Context) error {
				c = ctx
				return nil
			}

			err := app.Run(args)
			assert.NoError(t, err)

			result := parseGenericStorageConfig(c, tt.storageType)
			assert.NotNil(t, result)

			// Check that expected keys are present
			for _, expectedKey := range tt.expectKeys {
				assert.Contains(t, result, expectedKey)
			}
		})
	}
}

func TestAdvancedS3Config(t *testing.T) {
	tests := []struct {
		name     string
		flags    map[string]interface{}
		expected map[string]string
	}{
		{
			name: "advanced S3 flags should be parsed",
			flags: map[string]interface{}{
				"s3-access-key-id":          "test-key",
				"s3-session-token":          "test-token",
				"s3-storage-class":          "STANDARD_IA",
				"s3-server-side-encryption": "AES256",
				"s3-chunk-size":             "5Mi",
				"s3-force-path-style":       true,
				"s3-requester-pays":         true,
			},
			expected: map[string]string{
				"access_key_id":          "test-key",
				"session_token":          "test-token",
				"storage_class":          "STANDARD_IA",
				"server_side_encryption": "AES256",
				"chunk_size":             "5Mi",
				"force_path_style":       "true",
				"requester_pays":         "true",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &cli.App{
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "s3-access-key-id"},
					&cli.StringFlag{Name: "s3-session-token"},
					&cli.StringFlag{Name: "s3-storage-class"},
					&cli.StringFlag{Name: "s3-server-side-encryption"},
					&cli.StringFlag{Name: "s3-chunk-size"},
					&cli.BoolFlag{Name: "s3-force-path-style"},
					&cli.BoolFlag{Name: "s3-requester-pays"},
				},
			}

			args := []string{"app"}
			for flag, value := range tt.flags {
				switch v := value.(type) {
				case string:
					args = append(args, "--"+flag, v)
				case bool:
					if v {
						args = append(args, "--"+flag)
					}
				}
			}

			var c *cli.Context
			app.Action = func(ctx *cli.Context) error {
				c = ctx
				return nil
			}

			err := app.Run(args)
			assert.NoError(t, err)

			result := parseS3Config(c)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestAdvancedGCSConfig(t *testing.T) {
	tests := []struct {
		name     string
		flags    map[string]interface{}
		expected map[string]string
	}{
		{
			name: "advanced GCS flags should be parsed",
			flags: map[string]interface{}{
				"gcs-project-id":         "test-project",
				"gcs-object-acl":         "private",
				"gcs-storage-class":      "COLDLINE",
				"gcs-location":           "us-central1",
				"gcs-chunk-size":         "8Mi",
				"gcs-bucket-policy-only": true,
				"gcs-anonymous":          true,
			},
			expected: map[string]string{
				"project_number":     "test-project",
				"object_acl":         "private",
				"storage_class":      "COLDLINE",
				"location":           "us-central1",
				"chunk_size":         "8Mi",
				"bucket_policy_only": "true",
				"anonymous":          "true",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &cli.App{
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "gcs-project-id"},
					&cli.StringFlag{Name: "gcs-object-acl"},
					&cli.StringFlag{Name: "gcs-storage-class"},
					&cli.StringFlag{Name: "gcs-location"},
					&cli.StringFlag{Name: "gcs-chunk-size"},
					&cli.BoolFlag{Name: "gcs-bucket-policy-only"},
					&cli.BoolFlag{Name: "gcs-anonymous"},
				},
			}

			args := []string{"app"}
			for flag, value := range tt.flags {
				switch v := value.(type) {
				case string:
					args = append(args, "--"+flag, v)
				case bool:
					if v {
						args = append(args, "--"+flag)
					}
				}
			}

			var c *cli.Context
			app.Action = func(ctx *cli.Context) error {
				c = ctx
				return nil
			}

			err := app.Run(args)
			assert.NoError(t, err)

			result := parseGCSConfig(c)
			assert.Equal(t, tt.expected, result)
		})
	}
}
