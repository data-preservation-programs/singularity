# 1Fichier

{% code fullWidth="true" %}
```
NAME:
   singularity storage create fichier - 1Fichier

USAGE:
   singularity storage create fichier [command options]

DESCRIPTION:
   --api-key
      Your API Key, get it from https://1fichier.com/console/params.pl.

   --shared-folder
      If you want to download a shared folder, add this parameter.

   --file-password
      If you want to download a shared file that is password protected, add this parameter.

   --folder-password
      If you want to list the files in a shared folder that is password protected, add this parameter.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --api-key value  Your API Key, get it from https://1fichier.com/console/params.pl. [$API_KEY]
   --help, -h       show help

   Advanced

   --encoding value         The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot") [$ENCODING]
   --file-password value    If you want to download a shared file that is password protected, add this parameter. [$FILE_PASSWORD]
   --folder-password value  If you want to list the files in a shared folder that is password protected, add this parameter. [$FOLDER_PASSWORD]
   --shared-folder value    If you want to download a shared folder, add this parameter. [$SHARED_FOLDER]

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
