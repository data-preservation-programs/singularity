# HiDrive

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add hidrive - HiDrive

USAGE:
   singularity datasource add hidrive [command options] <dataset_name> <source_path>

DESCRIPTION:
   --hidrive-auth-url
      Auth server URL.
      
      Leave blank to use the provider defaults.

   --hidrive-chunk-size
      Chunksize for chunked uploads.
      
      Any files larger than the configured cutoff (or files of unknown size) will be uploaded in chunks of this size.
      
      The upper limit for this is 2147483647 bytes (about 2.000Gi).
      That is the maximum amount of bytes a single upload-operation will support.
      Setting this above the upper limit or to a negative value will cause uploads to fail.
      
      Setting this to larger values may increase the upload speed at the cost of using more memory.
      It can be set to smaller values smaller to save on memory.

   --hidrive-client-id
      OAuth Client Id.
      
      Leave blank normally.

   --hidrive-client-secret
      OAuth Client Secret.
      
      Leave blank normally.

   --hidrive-disable-fetching-member-count
      Do not fetch number of objects in directories unless it is absolutely necessary.
      
      Requests may be faster if the number of objects in subdirectories is not fetched.

   --hidrive-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --hidrive-endpoint
      Endpoint for the service.
      
      This is the URL that API-calls will be made to.

   --hidrive-root-prefix
      The root/parent folder for all paths.
      
      Fill in to use the specified folder as the parent for all paths given to the remote.
      This way rclone can use any folder as its starting point.

      Examples:
         | /       | The topmost directory accessible by rclone.
                   | This will be equivalent with "root" if rclone uses a regular HiDrive user account.
         | root    | The topmost directory of the HiDrive user account
         | <unset> | This specifies that there is no root-prefix for your paths.
                   | When using this you will always need to specify paths to this remote with a valid parent e.g. "remote:/path/to/dir" or "remote:root/path/to/dir".

   --hidrive-scope-access
      Access permissions that rclone should use when requesting access from HiDrive.

      Examples:
         | rw | Read and write access to resources.
         | ro | Read-only access to resources.

   --hidrive-scope-role
      User-level that rclone should use when requesting access from HiDrive.

      Examples:
         | user  | User-level access to management permissions.
                 | This will be sufficient in most cases.
         | admin | Extensive access to management permissions.
         | owner | Full access to management permissions.

   --hidrive-token
      OAuth Access Token as a JSON blob.

   --hidrive-token-url
      Token server url.
      
      Leave blank to use the provider defaults.

   --hidrive-upload-concurrency
      Concurrency for chunked uploads.
      
      This is the upper limit for how many transfers for the same file are running concurrently.
      Setting this above to a value smaller than 1 will cause uploads to deadlock.
      
      If you are uploading small numbers of large files over high-speed links
      and these uploads do not fully utilize your bandwidth, then increasing
      this may help to speed up the transfers.

   --hidrive-upload-cutoff
      Cutoff/Threshold for chunked uploads.
      
      Any files larger than this will be uploaded in chunks of the configured chunksize.
      
      The upper limit for this is 2147483647 bytes (about 2.000Gi).
      That is the maximum amount of bytes a single upload-operation will support.
      Setting this above the upper limit will cause uploads to fail.


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)

   Options for hidrive

   --hidrive-auth-url value                       Auth server URL. [$HIDRIVE_AUTH_URL]
   --hidrive-chunk-size value                     Chunksize for chunked uploads. (default: "48Mi") [$HIDRIVE_CHUNK_SIZE]
   --hidrive-client-id value                      OAuth Client Id. [$HIDRIVE_CLIENT_ID]
   --hidrive-client-secret value                  OAuth Client Secret. [$HIDRIVE_CLIENT_SECRET]
   --hidrive-disable-fetching-member-count value  Do not fetch number of objects in directories unless it is absolutely necessary. (default: "false") [$HIDRIVE_DISABLE_FETCHING_MEMBER_COUNT]
   --hidrive-encoding value                       The encoding for the backend. (default: "Slash,Dot") [$HIDRIVE_ENCODING]
   --hidrive-endpoint value                       Endpoint for the service. (default: "https://api.hidrive.strato.com/2.1") [$HIDRIVE_ENDPOINT]
   --hidrive-root-prefix value                    The root/parent folder for all paths. (default: "/") [$HIDRIVE_ROOT_PREFIX]
   --hidrive-scope-access value                   Access permissions that rclone should use when requesting access from HiDrive. (default: "rw") [$HIDRIVE_SCOPE_ACCESS]
   --hidrive-scope-role value                     User-level that rclone should use when requesting access from HiDrive. (default: "user") [$HIDRIVE_SCOPE_ROLE]
   --hidrive-token value                          OAuth Access Token as a JSON blob. [$HIDRIVE_TOKEN]
   --hidrive-token-url value                      Token server url. [$HIDRIVE_TOKEN_URL]
   --hidrive-upload-concurrency value             Concurrency for chunked uploads. (default: "4") [$HIDRIVE_UPLOAD_CONCURRENCY]
   --hidrive-upload-cutoff value                  Cutoff/Threshold for chunked uploads. (default: "96Mi") [$HIDRIVE_UPLOAD_CUTOFF]

```
{% endcode %}
