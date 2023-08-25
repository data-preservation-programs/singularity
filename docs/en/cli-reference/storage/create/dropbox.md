# Dropbox

{% code fullWidth="true" %}
```
NAME:
   singularity storage create dropbox - Dropbox

USAGE:
   singularity storage create dropbox [command options] <name> <path>

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

   --chunk_size
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
      

   --shared_files
      Instructs rclone to work on individual shared files.
      
      In this mode rclone's features are extremely limited - only list (ls, lsl, etc.) 
      operations and read operations (e.g. downloading) are supported in this mode.
      All other operations will be disabled.

   --shared_folders
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

   --batch_mode
      Upload file batching sync|async|off.
      
      This sets the batch mode used by rclone.
      
      For full info see [the main docs](https://rclone.org/dropbox/#batch-mode)
      
      This has 3 possible values
      
      - off - no batching
      - sync - batch uploads and check completion (default)
      - async - batch upload and don't check completion
      
      Rclone will close any outstanding batches when it exits which may make
      a delay on quit.
      

   --batch_size
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
      

   --batch_timeout
      Max time to allow an idle upload batch before uploading.
      
      If an upload batch is idle for more than this long then it will be
      uploaded.
      
      The default for this is 0 which means rclone will choose a sensible
      default based on the batch_mode in use.
      
      - batch_mode: async - default batch_timeout is 500ms
      - batch_mode: sync - default batch_timeout is 10s
      - batch_mode: off - not in use
      

   --batch_commit_timeout
      Max time to wait for a batch to finish committing

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --client_id value      OAuth Client Id. [$CLIENT_ID]
   --client_secret value  OAuth Client Secret. [$CLIENT_SECRET]
   --help, -h             show help

   Advanced

   --auth_url value              Auth server URL. [$AUTH_URL]
   --batch_commit_timeout value  Max time to wait for a batch to finish committing (default: "10m0s") [$BATCH_COMMIT_TIMEOUT]
   --batch_mode value            Upload file batching sync|async|off. (default: "sync") [$BATCH_MODE]
   --batch_size value            Max number of files in upload batch. (default: 0) [$BATCH_SIZE]
   --batch_timeout value         Max time to allow an idle upload batch before uploading. (default: "0s") [$BATCH_TIMEOUT]
   --chunk_size value            Upload chunk size (< 150Mi). (default: "48Mi") [$CHUNK_SIZE]
   --encoding value              The encoding for the backend. (default: "Slash,BackSlash,Del,RightSpace,InvalidUtf8,Dot") [$ENCODING]
   --impersonate value           Impersonate this user when using a business account. [$IMPERSONATE]
   --shared_files                Instructs rclone to work on individual shared files. (default: false) [$SHARED_FILES]
   --shared_folders              Instructs rclone to work on shared folders. (default: false) [$SHARED_FOLDERS]
   --token value                 OAuth Access Token as a JSON blob. [$TOKEN]
   --token_url value             Token server url. [$TOKEN_URL]

```
{% endcode %}
