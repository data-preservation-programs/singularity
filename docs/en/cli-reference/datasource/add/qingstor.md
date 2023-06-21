# QingCloud Object Storage

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add qingstor - QingCloud Object Storage

USAGE:
   singularity datasource add qingstor [command options] <dataset_name> <source_path>

DESCRIPTION:
   --qingstor-access-key-id
      QingStor Access Key ID.
      
      Leave blank for anonymous access or runtime credentials.

   --qingstor-connection-retries
      Number of connection retries.

   --qingstor-chunk-size
      Chunk size to use for uploading.
      
      When uploading files larger than upload_cutoff they will be uploaded
      as multipart uploads using this chunk size.
      
      Note that "--qingstor-upload-concurrency" chunks of this size are buffered
      in memory per transfer.
      
      If you are transferring large files over high-speed links and you have
      enough memory, then increasing this will speed up the transfers.

   --qingstor-upload-concurrency
      Concurrency for multipart uploads.
      
      This is the number of chunks of the same file that are uploaded
      concurrently.
      
      NB if you set this to > 1 then the checksums of multipart uploads
      become corrupted (the uploads themselves are not corrupted though).
      
      If you are uploading small numbers of large files over high-speed links
      and these uploads do not fully utilize your bandwidth, then increasing
      this may help to speed up the transfers.

   --qingstor-env-auth
      Get QingStor credentials from runtime.
      
      Only applies if access_key_id and secret_access_key is blank.

      Examples:
         | false | Enter QingStor credentials in the next step.
         | true  | Get QingStor credentials from the environment (env vars or IAM).

   --qingstor-endpoint
      Enter an endpoint URL to connection QingStor API.
      
      Leave blank will use the default value "https://qingstor.com:443".

   --qingstor-zone
      Zone to connect to.
      
      Default is "pek3a".

      Examples:
         | pek3a | The Beijing (China) Three Zone.
                 | Needs location constraint pek3a.
         | sh1a  | The Shanghai (China) First Zone.
                 | Needs location constraint sh1a.
         | gd2a  | The Guangdong (China) Second Zone.
                 | Needs location constraint gd2a.

   --qingstor-upload-cutoff
      Cutoff for switching to chunked upload.
      
      Any files larger than this will be uploaded in chunks of chunk_size.
      The minimum is 0 and the maximum is 5 GiB.

   --qingstor-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --qingstor-secret-access-key
      QingStor Secret Access Key (password).
      
      Leave blank for anonymous access or runtime credentials.


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)

   Options for qingstor

   --qingstor-access-key-id value       QingStor Access Key ID. [$QINGSTOR_ACCESS_KEY_ID]
   --qingstor-chunk-size value          Chunk size to use for uploading. (default: "4Mi") [$QINGSTOR_CHUNK_SIZE]
   --qingstor-connection-retries value  Number of connection retries. (default: "3") [$QINGSTOR_CONNECTION_RETRIES]
   --qingstor-encoding value            The encoding for the backend. (default: "Slash,Ctl,InvalidUtf8") [$QINGSTOR_ENCODING]
   --qingstor-endpoint value            Enter an endpoint URL to connection QingStor API. [$QINGSTOR_ENDPOINT]
   --qingstor-env-auth value            Get QingStor credentials from runtime. (default: "false") [$QINGSTOR_ENV_AUTH]
   --qingstor-secret-access-key value   QingStor Secret Access Key (password). [$QINGSTOR_SECRET_ACCESS_KEY]
   --qingstor-upload-concurrency value  Concurrency for multipart uploads. (default: "1") [$QINGSTOR_UPLOAD_CONCURRENCY]
   --qingstor-upload-cutoff value       Cutoff for switching to chunked upload. (default: "200Mi") [$QINGSTOR_UPLOAD_CUTOFF]
   --qingstor-zone value                Zone to connect to. [$QINGSTOR_ZONE]

```
{% endcode %}
