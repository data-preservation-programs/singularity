# Mega

{% code fullWidth="true" %}
```
NAME:
   singularity storage create mega - Mega

USAGE:
   singularity storage create mega [command options] [arguments...]

DESCRIPTION:
   --user
      User name.

   --pass
      Password.

   --debug
      Output more debug from Mega.
      
      If this flag is set (along with -vv) it will print further debugging
      information from the mega backend.

   --hard-delete
      Delete files permanently rather than putting them into the trash.
      
      Normally the mega backend will put all deletions into the trash rather
      than permanently deleting them.  If you specify this then rclone will
      permanently delete objects instead.

   --use-https
      Use HTTPS for transfers.
      
      MEGA uses plain text HTTP connections by default.
      Some ISPs throttle HTTP connections, this causes transfers to become very slow.
      Enabling this will force MEGA to use HTTPS for all transfers.
      HTTPS is normally not necesary since all data is already encrypted anyway.
      Enabling it will increase CPU usage and add network overhead.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --help, -h    show help
   --pass value  Password. [$PASS]
   --user value  User name. [$USER]

   Advanced

   --debug           Output more debug from Mega. (default: false) [$DEBUG]
   --encoding value  The encoding for the backend. (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --hard-delete     Delete files permanently rather than putting them into the trash. (default: false) [$HARD_DELETE]
   --use-https       Use HTTPS for transfers. (default: false) [$USE_HTTPS]

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
