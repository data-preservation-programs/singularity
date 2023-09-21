# Digi Storage, https://storage.rcs-rds.ro/

{% code fullWidth="true" %}
```
NAME:
   singularity storage create koofr digistorage - Digi Storage, https://storage.rcs-rds.ro/

USAGE:
   singularity storage create koofr digistorage [command options] [arguments...]

DESCRIPTION:
   --mountid
      Mount ID of the mount to use.
      
      If omitted, the primary mount is used.

   --setmtime
      Does the backend support setting modification time.
      
      Set this to false if you use a mount ID that points to a Dropbox or Amazon Drive backend.

   --user
      Your user name.

   --password
      Your password for rclone (generate one at https://storage.rcs-rds.ro/app/admin/preferences/password).

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --help, -h        show help
   --password value  Your password for rclone (generate one at https://storage.rcs-rds.ro/app/admin/preferences/password). [$PASSWORD]
   --user value      Your user name. [$USER]

   Advanced

   --encoding value  The encoding for the backend. (default: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --mountid value   Mount ID of the mount to use. [$MOUNTID]
   --setmtime        Does the backend support setting modification time. (default: true) [$SETMTIME]

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
