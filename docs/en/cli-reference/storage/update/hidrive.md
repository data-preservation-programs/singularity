# HiDrive

{% code fullWidth="true" %}
```
NAME:
   singularity storage update hidrive - HiDrive

USAGE:
   singularity storage update hidrive [command options] <name>

DESCRIPTION:
   --client_id
      OAuth Client Id.
      
      Leave blank normally.

   --client_secret
      OAuth Client Secret.
      
      Leave blank normally.

   --token
      OAuth Access Token as a JSON blob.

   --auth_url
      Auth server URL.
      
      Leave blank to use the provider defaults.

   --token_url
      Token server url.
      
      Leave blank to use the provider defaults.

   --scope_access
      Access permissions that rclone should use when requesting access from HiDrive.

      Examples:
         | rw | Read and write access to resources.
         | ro | Read-only access to resources.

   --scope_role
      User-level that rclone should use when requesting access from HiDrive.

      Examples:
         | user  | User-level access to management permissions.
         |       | This will be sufficient in most cases.
         | admin | Extensive access to management permissions.
         | owner | Full access to management permissions.

   --root_prefix
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

   --disable_fetching_member_count
      Do not fetch number of objects in directories unless it is absolutely necessary.
      
      Requests may be faster if the number of objects in subdirectories is not fetched.

   --chunk_size
      Chunksize for chunked uploads.
      
      Any files larger than the configured cutoff (or files of unknown size) will be uploaded in chunks of this size.
      
      The upper limit for this is 2147483647 bytes (about 2.000Gi).
      That is the maximum amount of bytes a single upload-operation will support.
      Setting this above the upper limit or to a negative value will cause uploads to fail.
      
      Setting this to larger values may increase the upload speed at the cost of using more memory.
      It can be set to smaller values smaller to save on memory.

   --upload_cutoff
      Cutoff/Threshold for chunked uploads.
      
      Any files larger than this will be uploaded in chunks of the configured chunksize.
      
      The upper limit for this is 2147483647 bytes (about 2.000Gi).
      That is the maximum amount of bytes a single upload-operation will support.
      Setting this above the upper limit will cause uploads to fail.

   --upload_concurrency
      Concurrency for chunked uploads.
      
      This is the upper limit for how many transfers for the same file are running concurrently.
      Setting this above to a value smaller than 1 will cause uploads to deadlock.
      
      If you are uploading small numbers of large files over high-speed links
      and these uploads do not fully utilize your bandwidth, then increasing
      this may help to speed up the transfers.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --client_id value      OAuth Client Id. [$CLIENT_ID]
   --client_secret value  OAuth Client Secret. [$CLIENT_SECRET]
   --help, -h             show help
   --scope_access value   Access permissions that rclone should use when requesting access from HiDrive. (default: "rw") [$SCOPE_ACCESS]

   Advanced

   --auth_url value                 Auth server URL. [$AUTH_URL]
   --chunk_size value               Chunksize for chunked uploads. (default: "48Mi") [$CHUNK_SIZE]
   --disable_fetching_member_count  Do not fetch number of objects in directories unless it is absolutely necessary. (default: false) [$DISABLE_FETCHING_MEMBER_COUNT]
   --encoding value                 The encoding for the backend. (default: "Slash,Dot") [$ENCODING]
   --endpoint value                 Endpoint for the service. (default: "https://api.hidrive.strato.com/2.1") [$ENDPOINT]
   --root_prefix value              The root/parent folder for all paths. (default: "/") [$ROOT_PREFIX]
   --scope_role value               User-level that rclone should use when requesting access from HiDrive. (default: "user") [$SCOPE_ROLE]
   --token value                    OAuth Access Token as a JSON blob. [$TOKEN]
   --token_url value                Token server url. [$TOKEN_URL]
   --upload_concurrency value       Concurrency for chunked uploads. (default: 4) [$UPLOAD_CONCURRENCY]
   --upload_cutoff value            Cutoff/Threshold for chunked uploads. (default: "96Mi") [$UPLOAD_CUTOFF]

```
{% endcode %}
