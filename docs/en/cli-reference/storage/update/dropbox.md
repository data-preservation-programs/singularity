# Dropbox

{% code fullWidth="true" %}
```
NAME:
   singularity storage update dropbox - Dropbox

USAGE:
   singularity storage update dropbox [command options] <name|id>

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

   --chunk-size
      Upload chunk size (< 150Mi).
      
      Any files larger than this will be uploaded in chunks of this size.
      
      Note that chunks are buffered in memory (one at a time) so rclone can
      deal with retries.  Setting this larger will increase the speed
      slightly (at most 10% for 128 MiB in tests) at the cost of using more
      memory.  It can be set smaller if you are tight on memory.

   --impersonate
      Impersonate this user when using a business account.
      
      Note that if you want to use impersonate, you should make sure this
      flag is set when running "rclone config" as this will cause rclone to
      request the "members.read" scope which it won't normally. This is
      needed to lookup a members email address into the internal ID that
      dropbox uses in the API.
      
      Using the "members.read" scope will require a Dropbox Team Admin
      to approve during the OAuth flow.
      
      You will have to use your own App (setting your own client_id and
      client_secret) to use this option as currently rclone's default set of
      permissions doesn't include "members.read". This can be added once
      v1.55 or later is in use everywhere.
      

   --shared-files
      Instructs rclone to work on individual shared files.
      
      In this mode rclone's features are extremely limited - only list (ls, lsl, etc.) 
      operations and read operations (e.g. downloading) are supported in this mode.
      All other operations will be disabled.

   --shared-folders
      Instructs rclone to work on shared folders.
            
      When this flag is used with no path only the List operation is supported and 
      all available shared folders will be listed. If you specify a path the first part 
      will be interpreted as the name of shared folder. Rclone will then try to mount this 
      shared to the root namespace. On success shared folder rclone proceeds normally. 
      The shared folder is now pretty much a normal folder and all normal operations 
      are supported. 
      
      Note that we don't unmount the shared folder afterwards so the 
      --dropbox-shared-folders can be omitted after the first use of a particular 
      shared folder.

   --batch-mode
      Upload file batching sync|async|off.
      
      This sets the batch mode used by rclone.
      
      For full info see [the main docs](https://rclone.org/dropbox/#batch-mode)
      
      This has 3 possible values
      
      - off - no batching
      - sync - batch uploads and check completion (default)
      - async - batch upload and don't check completion
      
      Rclone will close any outstanding batches when it exits which may make
      a delay on quit.
      

   --batch-size
      Max number of files in upload batch.
      
      This sets the batch size of files to upload. It has to be less than 1000.
      
      By default this is 0 which means rclone which calculate the batch size
      depending on the setting of batch_mode.
      
      - batch_mode: async - default batch_size is 100
      - batch_mode: sync - default batch_size is the same as --transfers
      - batch_mode: off - not in use
      
      Rclone will close any outstanding batches when it exits which may make
      a delay on quit.
      
      Setting this is a great idea if you are uploading lots of small files
      as it will make them a lot quicker. You can use --transfers 32 to
      maximise throughput.
      

   --batch-timeout
      Max time to allow an idle upload batch before uploading.
      
      If an upload batch is idle for more than this long then it will be
      uploaded.
      
      The default for this is 0 which means rclone will choose a sensible
      default based on the batch_mode in use.
      
      - batch_mode: async - default batch_timeout is 500ms
      - batch_mode: sync - default batch_timeout is 10s
      - batch_mode: off - not in use
      

   --batch-commit-timeout
      Max time to wait for a batch to finish committing

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --client-id value      OAuth Client Id. [$CLIENT_ID]
   --client-secret value  OAuth Client Secret. [$CLIENT_SECRET]
   --help, -h             show help

   Advanced

   --auth-url value              Auth server URL. [$AUTH_URL]
   --batch-commit-timeout value  Max time to wait for a batch to finish committing (default: "10m0s") [$BATCH_COMMIT_TIMEOUT]
   --batch-mode value            Upload file batching sync|async|off. (default: "sync") [$BATCH_MODE]
   --batch-size value            Max number of files in upload batch. (default: 0) [$BATCH_SIZE]
   --batch-timeout value         Max time to allow an idle upload batch before uploading. (default: "0s") [$BATCH_TIMEOUT]
   --chunk-size value            Upload chunk size (< 150Mi). (default: "48Mi") [$CHUNK_SIZE]
   --encoding value              The encoding for the backend. (default: "Slash,BackSlash,Del,RightSpace,InvalidUtf8,Dot") [$ENCODING]
   --impersonate value           Impersonate this user when using a business account. [$IMPERSONATE]
   --shared-files                Instructs rclone to work on individual shared files. (default: false) [$SHARED_FILES]
   --shared-folders              Instructs rclone to work on shared folders. (default: false) [$SHARED_FOLDERS]
   --token value                 OAuth Access Token as a JSON blob. [$TOKEN]
   --token-url value             Token server url. [$TOKEN_URL]

   HTTP Client Config

   --client-ca-cert value                           Path to CA certificate used to verify servers. To remove, use empty string.
   --client-cert value                              Path to Client SSL certificate (PEM) for mutual TLS auth. To remove, use empty string.
   --client-connect-timeout value                   HTTP Client Connect timeout (default: 1m0s)
   --client-expect-continue-timeout value           Timeout when using expect / 100-continue in HTTP (default: 1s)
   --client-header value [ --client-header value ]  Set HTTP header for all transactions (i.e. key=value). This will replace the existing header values. To remove a header, use --http-header "key="". To remove all headers, use --http-header ""
   --client-insecure-skip-verify                    Do not verify the server SSL certificate (insecure) (default: false)
   --client-key value                               Path to Client SSL private key (PEM) for mutual TLS auth. To remove, use empty string.
   --client-no-gzip                                 Don't set Accept-Encoding: gzip (default: false)
   --client-timeout value                           IO idle timeout (default: 5m0s)
   --client-user-agent value                        Set the user-agent to a specified string. To remove, use empty string. (default: rclone/v1.62.2-DEV)

   Retry Strategy

   --client-retry-backoff value      The constant delay backoff for retrying IO read errors (default: 1s)
   --client-retry-backoff-exp value  The exponential delay backoff for retrying IO read errors (default: 1.0)
   --client-retry-delay value        The initial delay before retrying IO read errors (default: 1s)
   --client-retry-max value          Max number of retries for IO read errors (default: 10)
   --client-skip-inaccessible        Skip inaccessible files when opening (default: false)

```
{% endcode %}
