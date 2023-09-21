# Internet Archive

{% code fullWidth="true" %}
```
NAME:
   singularity storage create internetarchive - Internet Archive

USAGE:
   singularity storage create internetarchive [command options] [arguments...]

DESCRIPTION:
   --access-key-id
      IAS3 Access Key.
      
      Leave blank for anonymous access.
      You can find one here: https://archive.org/account/s3.php

   --secret-access-key
      IAS3 Secret Key (password).
      
      Leave blank for anonymous access.

   --endpoint
      IAS3 Endpoint.
      
      Leave blank for default value.

   --front-endpoint
      Host of InternetArchive Frontend.
      
      Leave blank for default value.

   --disable-checksum
      Don't ask the server to test against MD5 checksum calculated by rclone.
      Normally rclone will calculate the MD5 checksum of the input before
      uploading it so it can ask the server to check the object against checksum.
      This is great for data integrity checking but can cause long delays for
      large files to start uploading.

   --wait-archive
      Timeout for waiting the server's processing tasks (specifically archive and book_op) to finish.
      Only enable if you need to be guaranteed to be reflected after write operations.
      0 to disable waiting. No errors to be thrown in case of timeout.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --access-key-id value      IAS3 Access Key. [$ACCESS_KEY_ID]
   --help, -h                 show help
   --secret-access-key value  IAS3 Secret Key (password). [$SECRET_ACCESS_KEY]

   Advanced

   --disable-checksum      Don't ask the server to test against MD5 checksum calculated by rclone. (default: true) [$DISABLE_CHECKSUM]
   --encoding value        The encoding for the backend. (default: "Slash,LtGt,CrLf,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --endpoint value        IAS3 Endpoint. (default: "https://s3.us.archive.org") [$ENDPOINT]
   --front-endpoint value  Host of InternetArchive Frontend. (default: "https://archive.org") [$FRONT_ENDPOINT]
   --wait-archive value    Timeout for waiting the server's processing tasks (specifically archive and book_op) to finish. (default: "0s") [$WAIT_ARCHIVE]

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
   --client-user-agent value                        Set the user-agent to a specified string (default: rclone/v1.62.2-DEV)

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
