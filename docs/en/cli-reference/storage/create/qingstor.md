# QingCloud Object Storage

{% code fullWidth="true" %}
```
NAME:
   singularity storage create qingstor - QingCloud Object Storage

USAGE:
   singularity storage create qingstor [command options] [arguments...]

DESCRIPTION:
   --env-auth
      Get QingStor credentials from runtime.
      
      Only applies if access_key_id and secret_access_key is blank.

      Examples:
         | false | Enter QingStor credentials in the next step.
         | true  | Get QingStor credentials from the environment (env vars or IAM).

   --access-key-id
      QingStor Access Key ID.
      
      Leave blank for anonymous access or runtime credentials.

   --secret-access-key
      QingStor Secret Access Key (password).
      
      Leave blank for anonymous access or runtime credentials.

   --endpoint
      Enter an endpoint URL to connection QingStor API.
      
      Leave blank will use the default value "https://qingstor.com:443".

   --zone
      Zone to connect to.
      
      Default is "pek3a".

      Examples:
         | pek3a | The Beijing (China) Three Zone.
         |       | Needs location constraint pek3a.
         | sh1a  | The Shanghai (China) First Zone.
         |       | Needs location constraint sh1a.
         | gd2a  | The Guangdong (China) Second Zone.
         |       | Needs location constraint gd2a.

   --connection-retries
      Number of connection retries.

   --upload-cutoff
      Cutoff for switching to chunked upload.
      
      Any files larger than this will be uploaded in chunks of chunk_size.
      The minimum is 0 and the maximum is 5 GiB.

   --chunk-size
      Chunk size to use for uploading.
      
      When uploading files larger than upload_cutoff they will be uploaded
      as multipart uploads using this chunk size.
      
      Note that "--qingstor-upload-concurrency" chunks of this size are buffered
      in memory per transfer.
      
      If you are transferring large files over high-speed links and you have
      enough memory, then increasing this will speed up the transfers.

   --upload-concurrency
      Concurrency for multipart uploads.
      
      This is the number of chunks of the same file that are uploaded
      concurrently.
      
      NB if you set this to > 1 then the checksums of multipart uploads
      become corrupted (the uploads themselves are not corrupted though).
      
      If you are uploading small numbers of large files over high-speed links
      and these uploads do not fully utilize your bandwidth, then increasing
      this may help to speed up the transfers.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --access-key-id value      QingStor Access Key ID. [$ACCESS_KEY_ID]
   --endpoint value           Enter an endpoint URL to connection QingStor API. [$ENDPOINT]
   --env-auth                 Get QingStor credentials from runtime. (default: false) [$ENV_AUTH]
   --help, -h                 show help
   --secret-access-key value  QingStor Secret Access Key (password). [$SECRET_ACCESS_KEY]
   --zone value               Zone to connect to. [$ZONE]

   Advanced

   --chunk-size value          Chunk size to use for uploading. (default: "4Mi") [$CHUNK_SIZE]
   --connection-retries value  Number of connection retries. (default: 3) [$CONNECTION_RETRIES]
   --encoding value            The encoding for the backend. (default: "Slash,Ctl,InvalidUtf8") [$ENCODING]
   --upload-concurrency value  Concurrency for multipart uploads. (default: 1) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value       Cutoff for switching to chunked upload. (default: "200Mi") [$UPLOAD_CUTOFF]

   General

   --name value  Name of the storage (default: Auto generated)
   --path value  Path of the storage

   HTTP Client Config

   --client-ca-cert value                           Path to CA certificate used to verify servers
   --client-cert value                              Path to Client SSL certificate (PEM) for mutual TLS auth
   --client-connect-timeout value                   HTTP Client Connect timeout (default: 1m0s)
   --client-expect-continue-timeout value           Timeout when using expect / 100-continue in HTTP (default: 1s)
   --client-header value [ --client-header value ]  Set HTTP header for all transactions (i.e. key=value)
   --client-insecure-skip-verify                    Do not verify the server SSL certificate (insecure) (default: false)
   --client-key value                               Path to Client SSL private key (PEM) for mutual TLS auth
   --client-no-gzip                                 Don't set Accept-Encoding: gzip (default: false)
   --client-timeout value                           IO idle timeout (default: 5m0s)
   --client-user-agent value                        Set the user-agent to a specified string (default: rclone/v1.62.2-DEV)

   Retry Strategy

   --client-retry-backoff value      The constant delay backoff for retrying IO read errors (default: 1s)
   --client-retry-backoff-exp value  The exponential delay backoff for retrying IO read errors (default: 1.0)
   --client-retry-delay value        The initial delay before retrying IO read errors (default: 1s)
   --client-retry-max value          Max number of retries for IO read errors (default: 10)

```
{% endcode %}
