# Akamai NetStorage

{% code fullWidth="true" %}
```
NAME:
   singularity storage update netstorage - Akamai NetStorage

USAGE:
   singularity storage update netstorage [command options] <name|id>

DESCRIPTION:
   --protocol
      Select between HTTP or HTTPS protocol.
      
      Most users should choose HTTPS, which is the default.
      HTTP is provided primarily for debugging purposes.

      Examples:
         | http  | HTTP protocol
         | https | HTTPS protocol

   --host
      Domain+path of NetStorage host to connect to.
      
      Format should be `<domain>/<internal folders>`

   --account
      Set the NetStorage account name

   --secret
      Set the NetStorage account secret/G2O key for authentication.
      
      Please choose the 'y' option to set your own password then enter your secret.

   --description
      Description of the remote.


OPTIONS:
   --account value  Set the NetStorage account name [$ACCOUNT]
   --help, -h       show help
   --host value     Domain+path of NetStorage host to connect to. [$HOST]
   --secret value   Set the NetStorage account secret/G2O key for authentication. [$SECRET]

   Advanced

   --description value  Description of the remote. [$DESCRIPTION]
   --protocol value     Select between HTTP or HTTPS protocol. (default: "https") [$PROTOCOL]

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
