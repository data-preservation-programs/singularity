# HiDrive

{% code fullWidth="true" %}
```
NAME:
   singularity storage create hidrive - HiDrive

USAGE:
   singularity storage create hidrive [command options]

DESCRIPTION:
   --client-id
      OAuth Client Id.
      
      Leave blank normally.

   --client-secret
      OAuth Client Secret.
      
      Leave blank normally.

   --token
      OAuth Access Token as a JSON blob.

   --auth-url
      Auth server URL.
      
      Leave blank to use the provider defaults.

   --token-url
      Token server url.
      
      Leave blank to use the provider defaults.

   --scope-access
      Access permissions that rclone should use when requesting access from HiDrive.

      Examples:
         | rw | Read and write access to resources.
         | ro | Read-only access to resources.

   --scope-role
      User-level that rclone should use when requesting access from HiDrive.

      Examples:
         | user  | User-level access to management permissions.
         |       | This will be sufficient in most cases.
         | admin | Extensive access to management permissions.
         | owner | Full access to management permissions.

   --root-prefix
      The root/parent folder for all paths.
      
      Fill in to use the specified folder as the parent for all paths given to the remote.
      This way rclone can use any folder as its starting point.

      Examples:
         | /       | The topmost directory accessible by rclone.
         |         | This will be equivalent with "root" if rclone uses a regular HiDrive user account.
         | root    | The topmost directory of the HiDrive user account
         | <unset> | This specifies that there is no root-prefix for your paths.
         |         | When using this you will always need to specify paths to this remote with a valid parent e.g. "remote:/path/to/dir" or "remote:root/path/to/dir".

   --endpoint
      Endpoint for the service.
      
      This is the URL that API-calls will be made to.

   --disable-fetching-member-count
      Do not fetch number of objects in directories unless it is absolutely necessary.
      
      Requests may be faster if the number of objects in subdirectories is not fetched.

   --chunk-size
      Chunksize for chunked uploads.
      
      Any files larger than the configured cutoff (or files of unknown size) will be uploaded in chunks of this size.
      
      The upper limit for this is 2147483647 bytes (about 2.000Gi).
      That is the maximum amount of bytes a single upload-operation will support.
      Setting this above the upper limit or to a negative value will cause uploads to fail.
      
      Setting this to larger values may increase the upload speed at the cost of using more memory.
      It can be set to smaller values smaller to save on memory.

   --upload-cutoff
      Cutoff/Threshold for chunked uploads.
      
      Any files larger than this will be uploaded in chunks of the configured chunksize.
      
      The upper limit for this is 2147483647 bytes (about 2.000Gi).
      That is the maximum amount of bytes a single upload-operation will support.
      Setting this above the upper limit will cause uploads to fail.

   --upload-concurrency
      Concurrency for chunked uploads.
      
      This is the upper limit for how many transfers for the same file are running concurrently.
      Setting this above to a value smaller than 1 will cause uploads to deadlock.
      
      If you are uploading small numbers of large files over high-speed links
      and these uploads do not fully utilize your bandwidth, then increasing
      this may help to speed up the transfers.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --description
      Description of the remote.


OPTIONS:
   --client-id value      OAuth Client Id. [$CLIENT_ID]
   --client-secret value  OAuth Client Secret. [$CLIENT_SECRET]
   --help, -h             show help
   --scope-access value   Access permissions that rclone should use when requesting access from HiDrive. (default: "rw") [$SCOPE_ACCESS]

   Advanced

   --auth-url value                 Auth server URL. [$AUTH_URL]
   --chunk-size value               Chunksize for chunked uploads. (default: "48Mi") [$CHUNK_SIZE]
   --description value              Description of the remote. [$DESCRIPTION]
   --disable-fetching-member-count  Do not fetch number of objects in directories unless it is absolutely necessary. (default: false) [$DISABLE_FETCHING_MEMBER_COUNT]
   --encoding value                 The encoding for the backend. (default: "Slash,Dot") [$ENCODING]
   --endpoint value                 Endpoint for the service. (default: "https://api.hidrive.strato.com/2.1") [$ENDPOINT]
   --root-prefix value              The root/parent folder for all paths. (default: "/") [$ROOT_PREFIX]
   --scope-role value               User-level that rclone should use when requesting access from HiDrive. (default: "user") [$SCOPE_ROLE]
   --token value                    OAuth Access Token as a JSON blob. [$TOKEN]
   --token-url value                Token server url. [$TOKEN_URL]
   --upload-concurrency value       Concurrency for chunked uploads. (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            Cutoff/Threshold for chunked uploads. (default: "96Mi") [$UPLOAD_CUTOFF]

   Client Config

   --client-ca-cert value                           Path to CA certificate used to verify servers
   --client-cert value                              Path to Client SSL certificate (PEM) for mutual TLS auth
   --client-connect-timeout value                   HTTP Client Connect timeout (default: 1m0s)
   --client-expect-continue-timeout value           Timeout when using expect / 100-continue in HTTP (default: 1s)
   --client-header value [ --client-header value ]  Set HTTP header for all transactions (i.e. key=value)
   --client-insecure-skip-verify                    Do not verify the server SSL certificate (insecure) (default: false)
   --client-key value                               Path to Client SSL private key (PEM) for mutual TLS auth
   --client-no-gzip                                 Don't set Accept-Encoding: gzip (default: false)
   --client-scan-concurrency value                  Max number of concurrent listing requests when scanning data source (default: 1)
   --client-timeout value                           IO idle timeout (default: 5m0s)
   --client-use-server-mod-time                     Use server modified time if possible (default: false)
   --client-user-agent value                        Set the user-agent to a specified string (default: rclone default)

   General

   --name value  Name of the storage (default: Auto generated)
   --path value  Path of the storage

   Retry Strategy

   --client-low-level-retries value  Maximum number of retries for low-level client errors (default: 10)
   --client-retry-backoff value      The constant delay backoff for retrying IO read errors (default: 1s)
   --client-retry-backoff-exp value  The exponential delay backoff for retrying IO read errors (default: 1.0)
   --client-retry-delay value        The initial delay before retrying IO read errors (default: 1s)
   --client-retry-max value          Max number of retries for IO read errors (default: 10)
   --client-skip-inaccessible        Skip inaccessible files when opening (default: false)

```
{% endcode %}
