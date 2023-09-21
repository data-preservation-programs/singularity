# Amazon Drive

{% code fullWidth="true" %}
```
NAME:
   singularity storage update acd - Amazon Drive

USAGE:
   singularity storage update acd [command options] <name|id>

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

   --checkpoint
      Checkpoint for internal polling (debug).

   --upload-wait-per-gb
      Additional time per GiB to wait after a failed complete upload to see if it appears.
      
      Sometimes Amazon Drive gives an error when a file has been fully
      uploaded but the file appears anyway after a little while.  This
      happens sometimes for files over 1 GiB in size and nearly every time for
      files bigger than 10 GiB. This parameter controls the time rclone waits
      for the file to appear.
      
      The default value for this parameter is 3 minutes per GiB, so by
      default it will wait 3 minutes for every GiB uploaded to see if the
      file appears.
      
      You can disable this feature by setting it to 0. This may cause
      conflict errors as rclone retries the failed upload but the file will
      most likely appear correctly eventually.
      
      These values were determined empirically by observing lots of uploads
      of big files for a range of file sizes.
      
      Upload with the "-v" flag to see more info about what rclone is doing
      in this situation.

   --templink-threshold
      Files >= this size will be downloaded via their tempLink.
      
      Files this size or more will be downloaded via their "tempLink". This
      is to work around a problem with Amazon Drive which blocks downloads
      of files bigger than about 10 GiB. The default for this is 9 GiB which
      shouldn't need to be changed.
      
      To download files above this threshold, rclone requests a "tempLink"
      which downloads the file through a temporary URL directly from the
      underlying S3 storage.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --client-id value      OAuth Client Id. [$CLIENT_ID]
   --client-secret value  OAuth Client Secret. [$CLIENT_SECRET]
   --help, -h             show help

   Advanced

   --auth-url value            Auth server URL. [$AUTH_URL]
   --checkpoint value          Checkpoint for internal polling (debug). [$CHECKPOINT]
   --encoding value            The encoding for the backend. (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --templink-threshold value  Files >= this size will be downloaded via their tempLink. (default: "9Gi") [$TEMPLINK_THRESHOLD]
   --token value               OAuth Access Token as a JSON blob. [$TOKEN]
   --token-url value           Token server url. [$TOKEN_URL]
   --upload-wait-per-gb value  Additional time per GiB to wait after a failed complete upload to see if it appears. (default: "3m0s") [$UPLOAD_WAIT_PER_GB]

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
