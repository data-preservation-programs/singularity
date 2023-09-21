# Citrix Sharefile

{% code fullWidth="true" %}
```
NAME:
   singularity storage create sharefile - Citrix Sharefile

USAGE:
   singularity storage create sharefile [command options] [arguments...]

DESCRIPTION:
   --upload-cutoff
      Cutoff for switching to multipart upload.

   --root-folder-id
      ID of the root folder.
      
      Leave blank to access "Personal Folders".  You can use one of the
      standard values here or any folder ID (long hex number ID).

      Examples:
         | <unset>    | Access the Personal Folders (default).
         | favorites  | Access the Favorites folder.
         | allshared  | Access all the shared folders.
         | connectors | Access all the individual connectors.
         | top        | Access the home, favorites, and shared folders as well as the connectors.

   --chunk-size
      Upload chunk size.
      
      Must a power of 2 >= 256k.
      
      Making this larger will improve performance, but note that each chunk
      is buffered in memory one per transfer.
      
      Reducing this will reduce memory usage but decrease performance.

   --endpoint
      Endpoint for API calls.
      
      This is usually auto discovered as part of the oauth process, but can
      be set manually to something like: https://XXX.sharefile.com
      

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --help, -h              show help
   --root-folder-id value  ID of the root folder. [$ROOT_FOLDER_ID]

   Advanced

   --chunk-size value     Upload chunk size. (default: "64Mi") [$CHUNK_SIZE]
   --encoding value       The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,LeftSpace,LeftPeriod,RightSpace,RightPeriod,InvalidUtf8,Dot") [$ENCODING]
   --endpoint value       Endpoint for API calls. [$ENDPOINT]
   --upload-cutoff value  Cutoff for switching to multipart upload. (default: "128Mi") [$UPLOAD_CUTOFF]

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
   --client-use-server-mod-time                     Use server modified time if possible (default: false)
   --client-user-agent value                        Set the user-agent to a specified string (default: rclone/v1.62.2-DEV)

   Retry Strategy

   --client-low-level-retries value  Maximum number of retries for low-level client errors (default: 10)
   --client-retry-backoff value      The constant delay backoff for retrying IO read errors (default: 1s)
   --client-retry-backoff-exp value  The exponential delay backoff for retrying IO read errors (default: 1.0)
   --client-retry-delay value        The initial delay before retrying IO read errors (default: 1s)
   --client-retry-max value          Max number of retries for IO read errors (default: 10)
   --client-skip-inaccessible        Skip inaccessible files when opening (default: false)

```
{% endcode %}
