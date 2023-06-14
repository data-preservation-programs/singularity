# Box

```
NAME:
   singularity datasource add box - Box

USAGE:
   singularity datasource add box [command options] <dataset_name> <source_path>

DESCRIPTION:
   --box-token-url
      Token server url.
      
      Leave blank to use the provider defaults.

   --box-box-config-file
      Box App config.json location
      
      Leave blank normally.
      
      Leading `~` will be expanded in the file name as will environment variables such as `${RCLONE_CONFIG_DIR}`.

   --box-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --box-client-id
      OAuth Client Id.
      
      Leave blank normally.

   --box-auth-url
      Auth server URL.
      
      Leave blank to use the provider defaults.

   --box-owned-by
      Only show items owned by the login (email address) passed in.

   --box-access-token
      Box App Primary Access Token
      
      Leave blank normally.

   --box-list-chunk
      Size of listing chunk 1-1000.

   --box-box-sub-type
      

      Examples:
         | user       | Rclone should act on behalf of a user.
         | enterprise | Rclone should act on behalf of a service account.

   --box-upload-cutoff
      Cutoff for switching to multipart upload (>= 50 MiB).

   --box-client-secret
      OAuth Client Secret.
      
      Leave blank normally.

   --box-root-folder-id
      Fill in for rclone to use a non root folder as its starting point.

   --box-token
      OAuth Access Token as a JSON blob.

   --box-commit-retries
      Max number of times to try committing a multipart file.


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)

   Options for box

   --box-access-token value     Box App Primary Access Token [$BOX_ACCESS_TOKEN]
   --box-auth-url value         Auth server URL. [$BOX_AUTH_URL]
   --box-box-config-file value  Box App config.json location [$BOX_BOX_CONFIG_FILE]
   --box-box-sub-type value     (default: "user") [$BOX_BOX_SUB_TYPE]
   --box-client-id value        OAuth Client Id. [$BOX_CLIENT_ID]
   --box-client-secret value    OAuth Client Secret. [$BOX_CLIENT_SECRET]
   --box-commit-retries value   Max number of times to try committing a multipart file. (default: "100") [$BOX_COMMIT_RETRIES]
   --box-encoding value         The encoding for the backend. (default: "Slash,BackSlash,Del,Ctl,RightSpace,InvalidUtf8,Dot") [$BOX_ENCODING]
   --box-list-chunk value       Size of listing chunk 1-1000. (default: "1000") [$BOX_LIST_CHUNK]
   --box-owned-by value         Only show items owned by the login (email address) passed in. [$BOX_OWNED_BY]
   --box-root-folder-id value   Fill in for rclone to use a non root folder as its starting point. (default: "0") [$BOX_ROOT_FOLDER_ID]
   --box-token value            OAuth Access Token as a JSON blob. [$BOX_TOKEN]
   --box-token-url value        Token server url. [$BOX_TOKEN_URL]
   --box-upload-cutoff value    Cutoff for switching to multipart upload (>= 50 MiB). (default: "50Mi") [$BOX_UPLOAD_CUTOFF]

```
