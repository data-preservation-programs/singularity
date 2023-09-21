# Akamai NetStorage

{% code fullWidth="true" %}
```
NAME:
   singularity storage create netstorage - Akamai NetStorage

USAGE:
   singularity storage create netstorage [command options] [arguments...]

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


OPTIONS:
   --account value  Set the NetStorage account name [$ACCOUNT]
   --help, -h       show help
   --host value     Domain+path of NetStorage host to connect to. [$HOST]
   --secret value   Set the NetStorage account secret/G2O key for authentication. [$SECRET]

   Advanced

   --protocol value  Select between HTTP or HTTPS protocol. (default: "https") [$PROTOCOL]

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
