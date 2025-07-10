package cliutil

import (
	"time"

	"github.com/urfave/cli/v2"
)

// CommonDealFlags contains reusable deal configuration flags for CLI commands
var CommonDealFlags = []cli.Flag{
	&cli.Float64Flag{
		Name:     "deal-price-per-gb",
		Usage:    "Price in FIL per GiB for storage deals",
		Value:    0.0,
		Category: "Deal Settings",
	},
	&cli.Float64Flag{
		Name:     "deal-price-per-gb-epoch",
		Usage:    "Price in FIL per GiB per epoch for storage deals",
		Value:    0.0,
		Category: "Deal Settings",
	},
	&cli.Float64Flag{
		Name:     "deal-price-per-deal",
		Usage:    "Price in FIL per deal for storage deals",
		Value:    0.0,
		Category: "Deal Settings",
	},
	&cli.DurationFlag{
		Name:     "deal-duration",
		Usage:    "Duration for storage deals (e.g., 535 days)",
		Value:    12840 * time.Hour, // ~535 days
		Category: "Deal Settings",
	},
	&cli.DurationFlag{
		Name:     "deal-start-delay",
		Usage:    "Start delay for storage deals (e.g., 72h)",
		Value:    72 * time.Hour,
		Category: "Deal Settings",
	},
	&cli.BoolFlag{
		Name:     "deal-verified",
		Usage:    "Whether deals should be verified",
		Category: "Deal Settings",
	},
	&cli.BoolFlag{
		Name:     "deal-keep-unsealed",
		Usage:    "Whether to keep unsealed copy of deals",
		Category: "Deal Settings",
	},
	&cli.BoolFlag{
		Name:     "deal-announce-to-ipni",
		Usage:    "Whether to announce deals to IPNI",
		Value:    true,
		Category: "Deal Settings",
	},
	&cli.StringFlag{
		Name:     "deal-provider",
		Usage:    "Storage Provider ID for deals (e.g., f01000)",
		Category: "Deal Settings",
	},
	&cli.StringFlag{
		Name:     "deal-url-template",
		Usage:    "URL template for deals",
		Category: "Deal Settings",
	},
	&cli.StringFlag{
		Name:     "deal-http-headers",
		Usage:    "HTTP headers for deals in JSON format",
		Category: "Deal Settings",
	},
	&cli.StringFlag{
		Name:     "deal-template",
		Usage:    "Name or ID of deal template to use for defaults",
		Category: "Deal Settings",
	},
}

// CommonStorageClientFlags contains reusable storage client configuration flags
var CommonStorageClientFlags = []cli.Flag{
	&cli.IntFlag{
		Name:     "client-retry-max",
		Usage:    "Max number of retries for IO read errors",
		Value:    10,
		Category: "Storage Client Config",
	},
	&cli.DurationFlag{
		Name:     "client-retry-delay",
		Usage:    "The initial delay before retrying IO read errors",
		Value:    time.Second,
		Category: "Storage Client Config",
	},
	&cli.DurationFlag{
		Name:     "client-retry-backoff",
		Usage:    "The constant delay backoff for retrying IO read errors",
		Value:    time.Second,
		Category: "Storage Client Config",
	},
	&cli.Float64Flag{
		Name:     "client-retry-backoff-exp",
		Usage:    "The exponential delay backoff for retrying IO read errors",
		Value:    1.0,
		Category: "Storage Client Config",
	},
	&cli.BoolFlag{
		Name:     "client-skip-inaccessible",
		Usage:    "Skip inaccessible files when opening",
		Category: "Storage Client Config",
	},
	&cli.IntFlag{
		Name:     "client-low-level-retries",
		Usage:    "Maximum number of retries for low-level client errors",
		Value:    10,
		Category: "Storage Client Config",
	},
	&cli.IntFlag{
		Name:     "client-scan-concurrency",
		Usage:    "Max number of concurrent listing requests when scanning data source",
		Value:    1,
		Category: "Storage Client Config",
	},
	&cli.DurationFlag{
		Name:     "client-connect-timeout",
		Usage:    "HTTP Client Connect timeout",
		Category: "Storage Client Config",
	},
	&cli.DurationFlag{
		Name:     "client-timeout",
		Usage:    "IO idle timeout",
		Category: "Storage Client Config",
	},
	&cli.DurationFlag{
		Name:     "client-expect-continue-timeout",
		Usage:    "Timeout when using expect / 100-continue in HTTP",
		Category: "Storage Client Config",
	},
	&cli.BoolFlag{
		Name:     "client-insecure-skip-verify",
		Usage:    "Do not verify the server SSL certificate (insecure)",
		Category: "Storage Client Config",
	},
	&cli.BoolFlag{
		Name:     "client-no-gzip",
		Usage:    "Don't set Accept-Encoding: gzip",
		Category: "Storage Client Config",
	},
	&cli.StringFlag{
		Name:     "client-user-agent",
		Usage:    "Set the user-agent to a specified string",
		Category: "Storage Client Config",
	},
	&cli.PathFlag{
		Name:     "client-ca-cert",
		Usage:    "Path to CA certificate used to verify servers",
		Category: "Storage Client Config",
	},
	&cli.PathFlag{
		Name:     "client-cert",
		Usage:    "Path to Client SSL certificate (PEM) for mutual TLS auth",
		Category: "Storage Client Config",
	},
	&cli.PathFlag{
		Name:     "client-key",
		Usage:    "Path to Client SSL private key (PEM) for mutual TLS auth",
		Category: "Storage Client Config",
	},
	&cli.StringSliceFlag{
		Name:     "client-header",
		Usage:    "Set HTTP header for all transactions (i.e. key=value)",
		Category: "Storage Client Config",
	},
	&cli.BoolFlag{
		Name:     "client-use-server-mod-time",
		Usage:    "Use server modified time if possible",
		Category: "Storage Client Config",
	},
}

// CommonS3Flags contains reusable S3 configuration flags
var CommonS3Flags = []cli.Flag{
	&cli.StringFlag{
		Name:     "s3-access-key-id",
		Usage:    "S3 Access Key ID",
		Category: "S3 Configuration",
	},
	&cli.StringFlag{
		Name:     "s3-secret-access-key",
		Usage:    "S3 Secret Access Key",
		Category: "S3 Configuration",
	},
	&cli.StringFlag{
		Name:     "s3-region",
		Usage:    "S3 Region (e.g., us-east-1)",
		Category: "S3 Configuration",
	},
	&cli.StringFlag{
		Name:     "s3-endpoint",
		Usage:    "Custom S3 endpoint URL",
		Category: "S3 Configuration",
	},
	&cli.BoolFlag{
		Name:     "s3-env-auth",
		Usage:    "Use environment variables for S3 authentication",
		Category: "S3 Configuration",
	},
	&cli.StringFlag{
		Name:     "s3-profile",
		Usage:    "AWS profile to use from shared credentials file",
		Category: "S3 Configuration",
	},
	&cli.StringFlag{
		Name:     "s3-shared-credentials-file",
		Usage:    "Path to AWS shared credentials file",
		Category: "S3 Configuration",
	},
	&cli.StringFlag{
		Name:     "s3-session-token",
		Usage:    "AWS session token",
		Category: "S3 Configuration",
	},
	&cli.StringFlag{
		Name:     "s3-storage-class",
		Usage:    "S3 storage class (e.g., STANDARD, REDUCED_REDUNDANCY, STANDARD_IA)",
		Category: "S3 Configuration",
	},
	&cli.StringFlag{
		Name:     "s3-server-side-encryption",
		Usage:    "Server-side encryption algorithm (e.g., AES256, aws:kms)",
		Category: "S3 Configuration",
	},
	&cli.StringFlag{
		Name:     "s3-sse-kms-key-id",
		Usage:    "KMS key ID for server-side encryption",
		Category: "S3 Configuration",
	},
	&cli.StringFlag{
		Name:     "s3-chunk-size",
		Usage:    "Upload chunk size (default: 5Mi)",
		Category: "S3 Configuration",
	},
	&cli.StringFlag{
		Name:     "s3-upload-concurrency",
		Usage:    "Number of concurrent uploads (default: 4)",
		Category: "S3 Configuration",
	},
	&cli.StringFlag{
		Name:     "s3-copy-cutoff",
		Usage:    "Cutoff for switching to multipart copy (default: 4.656Gi)",
		Category: "S3 Configuration",
	},
	&cli.StringFlag{
		Name:     "s3-upload-cutoff",
		Usage:    "Cutoff for switching to chunked upload (default: 200Mi)",
		Category: "S3 Configuration",
	},
	&cli.StringFlag{
		Name:     "s3-acl",
		Usage:    "Canned ACL for objects (e.g., private, public-read)",
		Category: "S3 Configuration",
	},
	&cli.BoolFlag{
		Name:     "s3-requester-pays",
		Usage:    "Enable requester pays for S3 bucket",
		Category: "S3 Configuration",
	},
	&cli.BoolFlag{
		Name:     "s3-force-path-style",
		Usage:    "Force path-style access instead of virtual-hosted-style",
		Category: "S3 Configuration",
	},
	&cli.BoolFlag{
		Name:     "s3-v2-auth",
		Usage:    "Use AWS signature version 2 authentication",
		Category: "S3 Configuration",
	},
	&cli.BoolFlag{
		Name:     "s3-use-accelerate-endpoint",
		Usage:    "Use AWS S3 accelerated endpoint",
		Category: "S3 Configuration",
	},
	&cli.BoolFlag{
		Name:     "s3-leave-parts-on-error",
		Usage:    "Leave successfully uploaded parts on S3 for manual recovery",
		Category: "S3 Configuration",
	},
}

// CommonGCSFlags contains reusable GCS configuration flags
var CommonGCSFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     "gcs-service-account-file",
		Usage:    "Path to GCS service account JSON file",
		Category: "GCS Configuration",
	},
	&cli.StringFlag{
		Name:     "gcs-service-account-credentials",
		Usage:    "GCS service account JSON credentials (inline)",
		Category: "GCS Configuration",
	},
	&cli.StringFlag{
		Name:     "gcs-project-id",
		Usage:    "GCS Project ID",
		Category: "GCS Configuration",
	},
	&cli.BoolFlag{
		Name:     "gcs-env-auth",
		Usage:    "Use environment variables for GCS authentication",
		Category: "GCS Configuration",
	},
	&cli.StringFlag{
		Name:     "gcs-object-acl",
		Usage:    "Access control list for objects (e.g., private, public-read)",
		Category: "GCS Configuration",
	},
	&cli.StringFlag{
		Name:     "gcs-bucket-acl",
		Usage:    "Access control list for buckets (e.g., private, public-read)",
		Category: "GCS Configuration",
	},
	&cli.BoolFlag{
		Name:     "gcs-bucket-policy-only",
		Usage:    "Use bucket policy only for access control",
		Category: "GCS Configuration",
	},
	&cli.StringFlag{
		Name:     "gcs-location",
		Usage:    "Location for new buckets (e.g., us-central1, europe-west1)",
		Category: "GCS Configuration",
	},
	&cli.StringFlag{
		Name:     "gcs-storage-class",
		Usage:    "Storage class for objects (e.g., STANDARD, NEARLINE, COLDLINE)",
		Category: "GCS Configuration",
	},
	&cli.StringFlag{
		Name:     "gcs-token",
		Usage:    "OAuth access token as JSON blob",
		Category: "GCS Configuration",
	},
	&cli.StringFlag{
		Name:     "gcs-auth-url",
		Usage:    "Auth server URL for OAuth",
		Category: "GCS Configuration",
	},
	&cli.StringFlag{
		Name:     "gcs-token-url",
		Usage:    "Token server URL for OAuth",
		Category: "GCS Configuration",
	},
	&cli.BoolFlag{
		Name:     "gcs-anonymous",
		Usage:    "Access public buckets anonymously",
		Category: "GCS Configuration",
	},
	&cli.StringFlag{
		Name:     "gcs-chunk-size",
		Usage:    "Upload chunk size (default: 8Mi)",
		Category: "GCS Configuration",
	},
	&cli.StringFlag{
		Name:     "gcs-upload-cutoff",
		Usage:    "Cutoff for switching to chunked upload (default: 8Mi)",
		Category: "GCS Configuration",
	},
	&cli.StringFlag{
		Name:     "gcs-copy-cutoff",
		Usage:    "Cutoff for switching to multipart copy (default: 8Mi)",
		Category: "GCS Configuration",
	},
	&cli.BoolFlag{
		Name:     "gcs-decompress",
		Usage:    "Decompress gzip-encoded files",
		Category: "GCS Configuration",
	},
	&cli.StringFlag{
		Name:     "gcs-endpoint",
		Usage:    "Custom endpoint for GCS API",
		Category: "GCS Configuration",
	},
	&cli.StringFlag{
		Name:     "gcs-encoding",
		Usage:    "The encoding for the backend",
		Category: "GCS Configuration",
	},
}
