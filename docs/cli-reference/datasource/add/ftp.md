# FTP

```
NAME:
   singularity datasource add ftp - FTP

USAGE:
   singularity datasource add ftp [command options] <dataset_name> <source_path>

DESCRIPTION:
   --ftp-pass
      FTP password.

   --ftp-tls
      Use Implicit FTPS (FTP over TLS).
      
      When using implicit FTP over TLS the client connects using TLS
      right from the start which breaks compatibility with
      non-TLS-aware servers. This is usually served over port 990 rather
      than port 21. Cannot be used in combination with explicit FTPS.

   --ftp-concurrency
      Maximum number of FTP simultaneous connections, 0 for unlimited.
      
      Note that setting this is very likely to cause deadlocks so it should
      be used with care.
      
      If you are doing a sync or copy then make sure concurrency is one more
      than the sum of `--transfers` and `--checkers`.
      
      If you use `--check-first` then it just needs to be one more than the
      maximum of `--checkers` and `--transfers`.
      
      So for `concurrency 3` you'd use `--checkers 2 --transfers 2
      --check-first` or `--checkers 1 --transfers 1`.
      
      

   --ftp-no-check-certificate
      Do not verify the TLS certificate of the server.

   --ftp-disable-epsv
      Disable using EPSV even if server advertises support.

   --ftp-shut-timeout
      Maximum time to wait for data connection closing status.

   --ftp-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

      Examples:
         | Asterisk,Ctl,Dot,Slash                               | ProFTPd can't handle '*' in file names
         | BackSlash,Ctl,Del,Dot,RightSpace,Slash,SquareBracket | PureFTPd can't handle '[]' or '*' in file names
         | Ctl,LeftPeriod,Slash                                 | VsFTPd can't handle file names starting with dot

   --ftp-disable-mlsd
      Disable using MLSD even if server advertises support.

   --ftp-disable-utf8
      Disable using UTF-8 even if server advertises support.

   --ftp-close-timeout
      Maximum time to wait for a response to close.

   --ftp-tls-cache-size
      Size of TLS session cache for all control and data connections.
      
      TLS cache allows to resume TLS sessions and reuse PSK between connections.
      Increase if default size is not enough resulting in TLS resumption errors.
      Enabled by default. Use 0 to disable.

   --ftp-ask-password
      Allow asking for FTP password when needed.
      
      If this is set and no password is supplied then rclone will ask for a password
      

   --ftp-explicit-tls
      Use Explicit FTPS (FTP over TLS).
      
      When using explicit FTP over TLS the client explicitly requests
      security from the server in order to upgrade a plain text connection
      to an encrypted one. Cannot be used in combination with implicit FTPS.

   --ftp-force-list-hidden
      Use LIST -a to force listing of hidden files and folders. This will disable the use of MLSD.

   --ftp-idle-timeout
      Max time before closing idle connections.
      
      If no connections have been returned to the connection pool in the time
      given, rclone will empty the connection pool.
      
      Set to 0 to keep connections indefinitely.
      

   --ftp-disable-tls13
      Disable TLS 1.3 (workaround for FTP servers with buggy TLS)

   --ftp-host
      FTP host to connect to.
      
      E.g. "ftp.example.com".

   --ftp-user
      FTP username.

   --ftp-port
      FTP port number.

   --ftp-writing-mdtm
      Use MDTM to set modification time (VsFtpd quirk)


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)

   Options for ftp

   --ftp-ask-password value          Allow asking for FTP password when needed. (default: "false") [$FTP_ASK_PASSWORD]
   --ftp-close-timeout value         Maximum time to wait for a response to close. (default: "1m0s") [$FTP_CLOSE_TIMEOUT]
   --ftp-concurrency value           Maximum number of FTP simultaneous connections, 0 for unlimited. (default: "0") [$FTP_CONCURRENCY]
   --ftp-disable-epsv value          Disable using EPSV even if server advertises support. (default: "false") [$FTP_DISABLE_EPSV]
   --ftp-disable-mlsd value          Disable using MLSD even if server advertises support. (default: "false") [$FTP_DISABLE_MLSD]
   --ftp-disable-tls13 value         Disable TLS 1.3 (workaround for FTP servers with buggy TLS) (default: "false") [$FTP_DISABLE_TLS13]
   --ftp-disable-utf8 value          Disable using UTF-8 even if server advertises support. (default: "false") [$FTP_DISABLE_UTF8]
   --ftp-encoding value              The encoding for the backend. (default: "Slash,Del,Ctl,RightSpace,Dot") [$FTP_ENCODING]
   --ftp-explicit-tls value          Use Explicit FTPS (FTP over TLS). (default: "false") [$FTP_EXPLICIT_TLS]
   --ftp-force-list-hidden value     Use LIST -a to force listing of hidden files and folders. This will disable the use of MLSD. (default: "false") [$FTP_FORCE_LIST_HIDDEN]
   --ftp-host value                  FTP host to connect to. [$FTP_HOST]
   --ftp-idle-timeout value          Max time before closing idle connections. (default: "1m0s") [$FTP_IDLE_TIMEOUT]
   --ftp-no-check-certificate value  Do not verify the TLS certificate of the server. (default: "false") [$FTP_NO_CHECK_CERTIFICATE]
   --ftp-pass value                  FTP password. [$FTP_PASS]
   --ftp-port value                  FTP port number. (default: "21") [$FTP_PORT]
   --ftp-shut-timeout value          Maximum time to wait for data connection closing status. (default: "1m0s") [$FTP_SHUT_TIMEOUT]
   --ftp-tls value                   Use Implicit FTPS (FTP over TLS). (default: "false") [$FTP_TLS]
   --ftp-tls-cache-size value        Size of TLS session cache for all control and data connections. (default: "32") [$FTP_TLS_CACHE_SIZE]
   --ftp-user value                  FTP username. (default: "shane") [$FTP_USER]
   --ftp-writing-mdtm value          Use MDTM to set modification time (VsFtpd quirk) (default: "false") [$FTP_WRITING_MDTM]

```
