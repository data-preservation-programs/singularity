# Jottacloud

{% code fullWidth="true" %}
```
NAME:
   singularity storage update jottacloud - Jottacloud

USAGE:
   singularity storage update jottacloud [command options] <name|id>

DESCRIPTION:
   --md5-memory-limit
      Files bigger than this will be cached on disk to calculate the MD5 if required.

   --trashed-only
      Only show files that are in the trash.
      
      This will show trashed files in their original directory structure.

   --hard-delete
      Delete files permanently rather than putting them into the trash.

   --upload-resume-limit
      Files bigger than this can be resumed if the upload fail's.

   --no-versions
      Avoid server side versioning by deleting files and recreating files instead of overwriting them.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --help, -h  show help

   Advanced

   --encoding value             The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --hard-delete                Delete files permanently rather than putting them into the trash. (default: false) [$HARD_DELETE]
   --md5-memory-limit value     Files bigger than this will be cached on disk to calculate the MD5 if required. (default: "10Mi") [$MD5_MEMORY_LIMIT]
   --no-versions                Avoid server side versioning by deleting files and recreating files instead of overwriting them. (default: false) [$NO_VERSIONS]
   --trashed-only               Only show files that are in the trash. (default: false) [$TRASHED_ONLY]
   --upload-resume-limit value  Files bigger than this can be resumed if the upload fail's. (default: "10Mi") [$UPLOAD_RESUME_LIMIT]

   Client Config

   --client-ca-cert value                           Path to CA certificate used to verify servers. To remove, use empty string.
   --client-cert value                              Path to Client SSL certificate (PEM) for mutual TLS auth. To remove, use empty string.
   --client-connect-timeout value                   HTTP Client Connect timeout (default: 1m0s)
   --client-expect-continue-timeout value           Timeout when using expect / 100-continue in HTTP (default: 1s)
   --client-header value [ --client-header value ]  Set HTTP header for all transactions (i.e. key=value). This will replace the existing header values. To remove a header, use --http-header "key="". To remove all headers, use --http-header ""
   --client-insecure-skip-verify                    Do not verify the server SSL certificate (insecure) (default: false)
   --client-key value                               Path to Client SSL private key (PEM) for mutual TLS auth. To remove, use empty string.
   --client-no-gzip                                 Don't set Accept-Encoding: gzip (default: false)
   --client-scan-concurrency value                  Max number of concurrent listing requests when scanning data source (default: 1)
   --client-timeout value                           IO idle timeout (default: 5m0s)
   --client-use-server-mod-time                     Use server modified time if possible (default: false)
   --client-user-agent value                        Set the user-agent to a specified string. To remove, use empty string. (default: rclone/v1.62.2-DEV)

   Retry Strategy

   --client-low-level-retries value  Maximum number of retries for low-level client errors (default: 10)
   --client-retry-backoff value      The constant delay backoff for retrying IO read errors (default: 1s)
   --client-retry-backoff-exp value  The exponential delay backoff for retrying IO read errors (default: 1.0)
   --client-retry-delay value        The initial delay before retrying IO read errors (default: 1s)
   --client-retry-max value          Max number of retries for IO read errors (default: 10)
   --client-skip-inaccessible        Skip inaccessible files when opening (default: false)

```
{% endcode %}
