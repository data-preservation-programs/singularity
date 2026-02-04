# Google Photos

{% code fullWidth="true" %}
```
NAME:
   singularity storage update gphotos - Google Photos

USAGE:
   singularity storage update gphotos [command options] <name|id>

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

   --read-only
      Set to make the Google Photos backend read only.
      
      If you choose read only then rclone will only request read only access
      to your photos, otherwise rclone will request full access.

   --read-size
      Set to read the size of media items.
      
      Normally rclone does not read the size of media items since this takes
      another transaction.  This isn't necessary for syncing.  However
      rclone mount needs to know the size of files in advance of reading
      them, so setting this flag when using rclone mount is recommended if
      you want to read the media.

   --start-year
      Year limits the photos to be downloaded to those which are uploaded after the given year.

   --include-archived
      Also view and download archived media.
      
      By default, rclone does not request archived media. Thus, when syncing,
      archived media is not visible in directory listings or transferred.
      
      Note that media in albums is always visible and synced, no matter
      their archive status.
      
      With this flag, archived media are always visible in directory
      listings and transferred.
      
      Without this flag, archived media will not be visible in directory
      listings and won't be transferred.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --batch-mode
      Upload file batching sync|async|off.
      
      This sets the batch mode used by rclone.
      
      This has 3 possible values
      
      - off - no batching
      - sync - batch uploads and check completion (default)
      - async - batch upload and don't check completion
      
      Rclone will close any outstanding batches when it exits which may make
      a delay on quit.
      

   --batch-size
      Max number of files in upload batch.
      
      This sets the batch size of files to upload. It has to be less than 50.
      
      By default this is 0 which means rclone will calculate the batch size
      depending on the setting of batch_mode.
      
      - batch_mode: async - default batch_size is 50
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
      
      - batch_mode: async - default batch_timeout is 10s
      - batch_mode: sync - default batch_timeout is 1s
      - batch_mode: off - not in use
      

   --batch-commit-timeout
      Max time to wait for a batch to finish committing

   --description
      Description of the remote.


OPTIONS:
   --client-id value      OAuth Client Id. [$CLIENT_ID]
   --client-secret value  OAuth Client Secret. [$CLIENT_SECRET]
   --help, -h             show help
   --read-only            Set to make the Google Photos backend read only. (default: false) [$READ_ONLY]

   Advanced

   --auth-url value              Auth server URL. [$AUTH_URL]
   --batch-commit-timeout value  Max time to wait for a batch to finish committing (default: "10m0s") [$BATCH_COMMIT_TIMEOUT]
   --batch-mode value            Upload file batching sync|async|off. (default: "sync") [$BATCH_MODE]
   --batch-size value            Max number of files in upload batch. (default: 0) [$BATCH_SIZE]
   --batch-timeout value         Max time to allow an idle upload batch before uploading. (default: "0s") [$BATCH_TIMEOUT]
   --description value           Description of the remote. [$DESCRIPTION]
   --encoding value              The encoding for the backend. (default: "Slash,CrLf,InvalidUtf8,Dot") [$ENCODING]
   --include-archived            Also view and download archived media. (default: false) [$INCLUDE_ARCHIVED]
   --read-size                   Set to read the size of media items. (default: false) [$READ_SIZE]
   --start-year value            Year limits the photos to be downloaded to those which are uploaded after the given year. (default: 2000) [$START_YEAR]
   --token value                 OAuth Access Token as a JSON blob. [$TOKEN]
   --token-url value             Token server url. [$TOKEN_URL]

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
   --client-user-agent value                        Set the user-agent to a specified string. To remove, use empty string. (default: rclone default)

   Retry Strategy

   --client-low-level-retries value  Maximum number of retries for low-level client errors (default: 10)
   --client-retry-backoff value      The constant delay backoff for retrying IO read errors (default: 1s)
   --client-retry-backoff-exp value  The exponential delay backoff for retrying IO read errors (default: 1.0)
   --client-retry-delay value        The initial delay before retrying IO read errors (default: 1s)
   --client-retry-max value          Max number of retries for IO read errors (default: 10)
   --client-skip-inaccessible        Skip inaccessible files when opening (default: false)

```
{% endcode %}
