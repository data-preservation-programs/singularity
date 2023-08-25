# FTP

{% code fullWidth="true" %}
```
NAME:
   singularity storage update ftp - FTP

USAGE:
   singularity storage update ftp [command options] <name>

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

   --explicit_tls
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
      
      

   --no_check_certificate
      Do not verify the TLS certificate of the server.

   --disable_epsv
      Disable using EPSV even if server advertises support.

   --disable_mlsd
      Disable using MLSD even if server advertises support.

   --disable_utf8
      Disable using UTF-8 even if server advertises support.

   --writing_mdtm
      Use MDTM to set modification time (VsFtpd quirk)

   --force_list_hidden
      Use LIST -a to force listing of hidden files and folders. This will disable the use of MLSD.

   --idle_timeout
      Max time before closing idle connections.
      
      If no connections have been returned to the connection pool in the time
      given, rclone will empty the connection pool.
      
      Set to 0 to keep connections indefinitely.
      

   --close_timeout
      Maximum time to wait for a response to close.

   --tls_cache_size
      Size of TLS session cache for all control and data connections.
      
      TLS cache allows to resume TLS sessions and reuse PSK between connections.
      Increase if default size is not enough resulting in TLS resumption errors.
      Enabled by default. Use 0 to disable.

   --disable_tls13
      Disable TLS 1.3 (workaround for FTP servers with buggy TLS)

   --shut_timeout
      Maximum time to wait for data connection closing status.

   --ask_password
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
   --explicit_tls  Use Explicit FTPS (FTP over TLS). (default: false) [$EXPLICIT_TLS]
   --help, -h      show help
   --host value    FTP host to connect to. [$HOST]
   --pass value    FTP password. [$PASS]
   --port value    FTP port number. (default: 21) [$PORT]
   --tls           Use Implicit FTPS (FTP over TLS). (default: false) [$TLS]
   --user value    FTP username. (default: "$USER") [$USER]

   Advanced

   --ask_password          Allow asking for FTP password when needed. (default: false) [$ASK_PASSWORD]
   --close_timeout value   Maximum time to wait for a response to close. (default: "1m0s") [$CLOSE_TIMEOUT]
   --concurrency value     Maximum number of FTP simultaneous connections, 0 for unlimited. (default: 0) [$CONCURRENCY]
   --disable_epsv          Disable using EPSV even if server advertises support. (default: false) [$DISABLE_EPSV]
   --disable_mlsd          Disable using MLSD even if server advertises support. (default: false) [$DISABLE_MLSD]
   --disable_tls13         Disable TLS 1.3 (workaround for FTP servers with buggy TLS) (default: false) [$DISABLE_TLS13]
   --disable_utf8          Disable using UTF-8 even if server advertises support. (default: false) [$DISABLE_UTF8]
   --encoding value        The encoding for the backend. (default: "Slash,Del,Ctl,RightSpace,Dot") [$ENCODING]
   --force_list_hidden     Use LIST -a to force listing of hidden files and folders. This will disable the use of MLSD. (default: false) [$FORCE_LIST_HIDDEN]
   --idle_timeout value    Max time before closing idle connections. (default: "1m0s") [$IDLE_TIMEOUT]
   --no_check_certificate  Do not verify the TLS certificate of the server. (default: false) [$NO_CHECK_CERTIFICATE]
   --shut_timeout value    Maximum time to wait for data connection closing status. (default: "1m0s") [$SHUT_TIMEOUT]
   --tls_cache_size value  Size of TLS session cache for all control and data connections. (default: 32) [$TLS_CACHE_SIZE]
   --writing_mdtm          Use MDTM to set modification time (VsFtpd quirk) (default: false) [$WRITING_MDTM]

```
{% endcode %}
