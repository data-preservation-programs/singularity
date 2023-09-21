# FTP

{% code fullWidth="true" %}
```
NAME:
   singularity storage create ftp - FTP

USAGE:
   singularity storage create ftp [command options] [arguments...]

DESCRIPTION:
   --host
      FTP host to connect to.
      
      E.g. "ftp.example.com".

   --user
      FTP username.

   --port
      FTP port number.

   --pass
      FTP password.

   --tls
      Use Implicit FTPS (FTP over TLS).
      
      When using implicit FTP over TLS the client connects using TLS
      right from the start which breaks compatibility with
      non-TLS-aware servers. This is usually served over port 990 rather
      than port 21. Cannot be used in combination with explicit FTPS.

   --explicit-tls
      Use Explicit FTPS (FTP over TLS).
      
      When using explicit FTP over TLS the client explicitly requests
      security from the server in order to upgrade a plain text connection
      to an encrypted one. Cannot be used in combination with implicit FTPS.

   --concurrency
      Maximum number of FTP simultaneous connections, 0 for unlimited.
      
      Note that setting this is very likely to cause deadlocks so it should
      be used with care.
      
      If you are doing a sync or copy then make sure concurrency is one more
      than the sum of `--transfers` and `--checkers`.
      
      If you use `--check-first` then it just needs to be one more than the
      maximum of `--checkers` and `--transfers`.
      
      So for `concurrency 3` you'd use `--checkers 2 --transfers 2
      --check-first` or `--checkers 1 --transfers 1`.
      
      

   --no-check-certificate
      Do not verify the TLS certificate of the server.

   --disable-epsv
      Disable using EPSV even if server advertises support.

   --disable-mlsd
      Disable using MLSD even if server advertises support.

   --disable-utf8
      Disable using UTF-8 even if server advertises support.

   --writing-mdtm
      Use MDTM to set modification time (VsFtpd quirk)

   --force-list-hidden
      Use LIST -a to force listing of hidden files and folders. This will disable the use of MLSD.

   --idle-timeout
      Max time before closing idle connections.
      
      If no connections have been returned to the connection pool in the time
      given, rclone will empty the connection pool.
      
      Set to 0 to keep connections indefinitely.
      

   --close-timeout
      Maximum time to wait for a response to close.

   --tls-cache-size
      Size of TLS session cache for all control and data connections.
      
      TLS cache allows to resume TLS sessions and reuse PSK between connections.
      Increase if default size is not enough resulting in TLS resumption errors.
      Enabled by default. Use 0 to disable.

   --disable-tls13
      Disable TLS 1.3 (workaround for FTP servers with buggy TLS)

   --shut-timeout
      Maximum time to wait for data connection closing status.

   --ask-password
      Allow asking for FTP password when needed.
      
      If this is set and no password is supplied then rclone will ask for a password
      

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

      Examples:
         | Asterisk,Ctl,Dot,Slash                               | ProFTPd can't handle '*' in file names
         | BackSlash,Ctl,Del,Dot,RightSpace,Slash,SquareBracket | PureFTPd can't handle '[]' or '*' in file names
         | Ctl,LeftPeriod,Slash                                 | VsFTPd can't handle file names starting with dot


OPTIONS:
   --explicit-tls  Use Explicit FTPS (FTP over TLS). (default: false) [$EXPLICIT_TLS]
   --help, -h      show help
   --host value    FTP host to connect to. [$HOST]
   --pass value    FTP password. [$PASS]
   --port value    FTP port number. (default: 21) [$PORT]
   --tls           Use Implicit FTPS (FTP over TLS). (default: false) [$TLS]
   --user value    FTP username. (default: "$USER") [$USER]

   Advanced

   --ask-password          Allow asking for FTP password when needed. (default: false) [$ASK_PASSWORD]
   --close-timeout value   Maximum time to wait for a response to close. (default: "1m0s") [$CLOSE_TIMEOUT]
   --concurrency value     Maximum number of FTP simultaneous connections, 0 for unlimited. (default: 0) [$CONCURRENCY]
   --disable-epsv          Disable using EPSV even if server advertises support. (default: false) [$DISABLE_EPSV]
   --disable-mlsd          Disable using MLSD even if server advertises support. (default: false) [$DISABLE_MLSD]
   --disable-tls13         Disable TLS 1.3 (workaround for FTP servers with buggy TLS) (default: false) [$DISABLE_TLS13]
   --disable-utf8          Disable using UTF-8 even if server advertises support. (default: false) [$DISABLE_UTF8]
   --encoding value        The encoding for the backend. (default: "Slash,Del,Ctl,RightSpace,Dot") [$ENCODING]
   --force-list-hidden     Use LIST -a to force listing of hidden files and folders. This will disable the use of MLSD. (default: false) [$FORCE_LIST_HIDDEN]
   --idle-timeout value    Max time before closing idle connections. (default: "1m0s") [$IDLE_TIMEOUT]
   --no-check-certificate  Do not verify the TLS certificate of the server. (default: false) [$NO_CHECK_CERTIFICATE]
   --shut-timeout value    Maximum time to wait for data connection closing status. (default: "1m0s") [$SHUT_TIMEOUT]
   --tls-cache-size value  Size of TLS session cache for all control and data connections. (default: 32) [$TLS_CACHE_SIZE]
   --writing-mdtm          Use MDTM to set modification time (VsFtpd quirk) (default: false) [$WRITING_MDTM]

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
