# Complete data onboarding workflow (storage → preparation → scanning → deal creation)

{% code fullWidth="true" %}
```
NAME:
   singularity onboard - Complete data onboarding workflow (storage → preparation → scanning → deal creation)

USAGE:
   singularity onboard [command options]

DESCRIPTION:
   The onboard command provides a unified workflow for complete data onboarding.

   It performs the following steps automatically:
   1. Creates storage connections (if paths provided)
   2. Creates data preparation with deal template configuration
   3. Starts scanning immediately
   4. Enables automatic job progression (scan → pack → daggen → deals)
   5. Optionally starts managed workers to process jobs

   This is the simplest way to onboard data from source to storage deals.
   Use deal templates to configure deal parameters - individual deal flags are not supported.

OPTIONS:
   --auto-create-deals                Enable automatic deal creation after preparation completion (default: true)
   --json                             Output result in JSON format for automation (default: false)
   --max-size value                   Maximum size of a single CAR file (default: "31.5GiB")
   --max-workers value                Maximum number of workers to run (default: 3)
   --name value                       Name for the preparation
   --no-dag                           Disable maintaining folder DAG structure (default: false)
   --output value [ --output value ]  Output path(s) for CAR files (local paths or remote URLs like s3://bucket/path)
   --source value [ --source value ]  Source path(s) to onboard (local paths or remote URLs like s3://bucket/path)
   --sp-validation                    Enable storage provider validation (default: false)
   --start-workers                    Start managed workers to process jobs automatically (default: true)
   --timeout value                    Timeout for waiting for completion (0 = no timeout) (default: 0s)
   --wait-for-completion              Wait and monitor until all jobs complete (default: false)
   --wallet-validation                Enable wallet balance validation (default: false)

   Remote Storage Configuration

   --source-type value      Source storage type (local, s3, gcs, azure, etc.) (default: "local")
   --source-provider value  Source storage provider (for s3: aws, minio, wasabi, etc.)
   --output-type value      Output storage type (local, s3, gcs, azure, etc.) (default: "local")
   --output-provider value  Output storage provider

   Storage Configuration

   --output-config value  Output storage configuration in JSON format (key-value pairs)
   --output-name value    Custom name for output storage (auto-generated if not provided)
   --source-config value  Source storage configuration in JSON format (key-value pairs)
   --source-name value    Custom name for source storage (auto-generated if not provided)

   Deal Settings

   --deal-template-id value   Deal template ID to use for deal configuration (required when auto-create-deals is enabled)

   Storage Client Config

   --client-ca-cert value                           Path to CA certificate used to verify servers
   --client-cert value                              Path to Client SSL certificate (PEM) for mutual TLS auth
   --client-connect-timeout value                   HTTP Client Connect timeout (default: 0s)
   --client-expect-continue-timeout value           Timeout when using expect / 100-continue in HTTP (default: 0s)
   --client-header value [ --client-header value ]  Set HTTP header for all transactions (i.e. key=value)
   --client-insecure-skip-verify                    Do not verify the server SSL certificate (insecure) (default: false)
   --client-key value                               Path to Client SSL private key (PEM) for mutual TLS auth
   --client-low-level-retries value                 Maximum number of retries for low-level client errors (default: 10)
   --client-no-gzip                                 Don't set Accept-Encoding: gzip (default: false)
   --client-retry-backoff value                     The constant delay backoff for retrying IO read errors (default: 1s)
   --client-retry-backoff-exp value                 The exponential delay backoff for retrying IO read errors (default: 1)
   --client-retry-delay value                       The initial delay before retrying IO read errors (default: 1s)
   --client-retry-max value                         Max number of retries for IO read errors (default: 10)
   --client-scan-concurrency value                  Max number of concurrent listing requests when scanning data source (default: 1)
   --client-skip-inaccessible                       Skip inaccessible files when opening (default: false)
   --client-timeout value                           IO idle timeout (default: 0s)
   --client-use-server-mod-time                     Use server modified time if possible (default: false)
   --client-user-agent value                        Set the user-agent to a specified string

   Dynamic Storage Backend Flags

   The onboard command dynamically generates flags for all supported storage backends.
   For each backend (s3, gcs, azure, etc.), both source and output flags are available:
   
   Source Storage: --source-{backend}-{option}
   Output Storage: --output-{backend}-{option}
   
   Examples:
   - S3 access key: --source-s3-access-key-id, --output-s3-access-key-id
   - GCS project: --source-gcs-project-number, --output-gcs-project-number
   - Azure account: --source-azureblob-account, --output-azureblob-account
   
   Run 'singularity onboard --help' to see all available backend-specific flags.

```
{% endcode %}

## Remote Storage Support

The onboard command now supports remote storage backends for both source and output locations. This provides full parity with the `singularity storage create` command.

### Supported Storage Types

All rclone-supported storage backends are available:

- **Cloud Storage**: S3 (AWS, Wasabi, MinIO, etc.), Google Cloud Storage, Azure Blob Storage
- **File Storage**: Local filesystem, SFTP, FTP, WebDAV
- **Cloud Providers**: Google Drive, Dropbox, OneDrive, Box
- **Specialized**: HDFS, Swift, B2, and many more

### Configuration Methods

1. **CLI Flags**: Use dynamic backend-specific flags (e.g., `--source-s3-access-key-id`)
2. **JSON Config**: Use `--source-config` and `--output-config` with JSON key-value pairs
3. **Environment Variables**: Standard rclone environment variables are supported
4. **Provider Defaults**: Automatic application of provider-specific defaults

### Examples

#### S3 with AWS Provider
```bash
singularity onboard \
  --name "my-dataset" \
  --source "s3://my-bucket/data/" \
  --source-type "s3" \
  --source-provider "aws" \
  --source-s3-access-key-id "AKIAIOSFODNN7EXAMPLE" \
  --source-s3-secret-access-key "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY" \
  --source-s3-region "us-west-2" \
  --output "/local/output/path" \
  --deal-template-id "1"
```

#### Google Cloud Storage
```bash
singularity onboard \
  --name "gcs-dataset" \
  --source "gcs://my-bucket/data/" \
  --source-type "gcs" \
  --source-gcs-project-number "my-project-123" \
  --source-gcs-service-account-file "/path/to/service-account.json" \
  --output "gcs://output-bucket/cars/" \
  --output-type "gcs" \
  --output-gcs-project-number "my-project-123" \
  --output-gcs-service-account-file "/path/to/service-account.json" \
  --deal-template-id "1"
```

#### Using JSON Configuration
```bash
singularity onboard \
  --name "json-config-dataset" \
  --source "s3://my-bucket/data/" \
  --source-type "s3" \
  --source-config '{"access_key_id": "AKIAIOSFODNN7EXAMPLE", "secret_access_key": "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY", "region": "us-west-2"}' \
  --output "/local/output/path" \
  --deal-template-id "1"
```

### Error Handling

The onboard command includes comprehensive validation:

- **Storage Type Validation**: Ensures supported backend types
- **Configuration Validation**: Validates JSON format and required fields
- **Connectivity Testing**: Tests storage connectivity before processing
- **Provider Validation**: Validates provider-specific configurations

### Security Considerations

- **Insecure Configurations**: Warnings for insecure client configurations
- **Credential Handling**: Secure handling of authentication credentials
- **SSL/TLS**: Full support for certificate-based authentication
- **Headers**: Support for custom HTTP headers for authentication

For the complete list of available flags and options, run:
```bash
singularity onboard --help
```