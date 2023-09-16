# Sia Decentralized Cloud

{% code fullWidth="true" %}
```
NAME:
   singularity storage create sia - Sia Decentralized Cloud

USAGE:
   singularity storage create sia [command options] [arguments...]

DESCRIPTION:
   --api-url
      Sia daemon API URL, like http://sia.daemon.host:9980.
      
      Note that siad must run with --disable-api-security to open API port for other hosts (not recommended).
      Keep default if Sia daemon runs on localhost.

   --api-password
      Sia Daemon API Password.
      
      Can be found in the apipassword file located in HOME/.sia/ or in the daemon directory.

   --user-agent
      Siad User Agent
      
      Sia daemon requires the 'Sia-Agent' user agent by default for security

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --api-password value  Sia Daemon API Password. [$API_PASSWORD]
   --api-url value       Sia daemon API URL, like http://sia.daemon.host:9980. (default: "http://127.0.0.1:9980") [$API_URL]
   --help, -h            show help

   Advanced

   --encoding value    The encoding for the backend. (default: "Slash,Question,Hash,Percent,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --user-agent value  Siad User Agent (default: "Sia-Agent") [$USER_AGENT]

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
