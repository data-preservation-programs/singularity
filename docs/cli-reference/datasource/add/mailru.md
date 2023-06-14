# Mail.ru Cloud

```
NAME:
   singularity datasource add mailru - Mail.ru Cloud

USAGE:
   singularity datasource add mailru [command options] <dataset_name> <source_path>

DESCRIPTION:
   --mailru-user
      User name (usually email).

   --mailru-pass
      Password.
      
      This must be an app password - rclone will not work with your normal
      password. See the Configuration section in the docs for how to make an
      app password.
      

   --mailru-speedup-enable
      Skip full upload if there is another file with same data hash.
      
      This feature is called "speedup" or "put by hash". It is especially efficient
      in case of generally available files like popular books, video or audio clips,
      because files are searched by hash in all accounts of all mailru users.
      It is meaningless and ineffective if source file is unique or encrypted.
      Please note that rclone may need local memory and disk space to calculate
      content hash in advance and decide whether full upload is required.
      Also, if rclone does not know file size in advance (e.g. in case of
      streaming or partial uploads), it will not even try this optimization.

      Examples:
         | true  | Enable
         | false | Disable

   --mailru-speedup-file-patterns
      Comma separated list of file name patterns eligible for speedup (put by hash).
      
      Patterns are case insensitive and can contain '*' or '?' meta characters.

      Examples:
         | <unset>                 | Empty list completely disables speedup (put by hash).
         | *                       | All files will be attempted for speedup.
         | *.mkv,*.avi,*.mp4,*.mp3 | Only common audio/video files will be tried for put by hash.
         | *.zip,*.gz,*.rar,*.pdf  | Only common archives or PDF books will be tried for speedup.

   --mailru-speedup-max-disk
      This option allows you to disable speedup (put by hash) for large files.
      
      Reason is that preliminary hashing can exhaust your RAM or disk space.

      Examples:
         | 0  | Completely disable speedup (put by hash).
         | 1G | Files larger than 1Gb will be uploaded directly.
         | 3G | Choose this option if you have less than 3Gb free on local disk.

   --mailru-speedup-max-memory
      Files larger than the size given below will always be hashed on disk.

      Examples:
         | 0    | Preliminary hashing will always be done in a temporary disk location.
         | 32M  | Do not dedicate more than 32Mb RAM for preliminary hashing.
         | 256M | You have at most 256Mb RAM free for hash calculations.

   --mailru-check-hash
      What should copy do if file checksum is mismatched or invalid.

      Examples:
         | true  | Fail with error.
         | false | Ignore and continue.

   --mailru-user-agent
      HTTP user agent used internally by client.
      
      Defaults to "rclone/VERSION" or "--user-agent" provided on command line.

   --mailru-quirks
      Comma separated list of internal maintenance flags.
      
      This option must not be used by an ordinary user. It is intended only to
      facilitate remote troubleshooting of backend issues. Strict meaning of
      flags is not documented and not guaranteed to persist between releases.
      Quirks will be removed when the backend grows stable.
      Supported quirks: atomicmkdir binlist unknowndirs

   --mailru-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)

   Options for mailru

   --mailru-check-hash value             What should copy do if file checksum is mismatched or invalid. (default: "true") [$MAILRU_CHECK_HASH]
   --mailru-encoding value               The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$MAILRU_ENCODING]
   --mailru-pass value                   Password. [$MAILRU_PASS]
   --mailru-speedup-enable value         Skip full upload if there is another file with same data hash. (default: "true") [$MAILRU_SPEEDUP_ENABLE]
   --mailru-speedup-file-patterns value  Comma separated list of file name patterns eligible for speedup (put by hash). (default: "*.mkv,*.avi,*.mp4,*.mp3,*.zip,*.gz,*.rar,*.pdf") [$MAILRU_SPEEDUP_FILE_PATTERNS]
   --mailru-speedup-max-disk value       This option allows you to disable speedup (put by hash) for large files. (default: "3Gi") [$MAILRU_SPEEDUP_MAX_DISK]
   --mailru-speedup-max-memory value     Files larger than the size given below will always be hashed on disk. (default: "32Mi") [$MAILRU_SPEEDUP_MAX_MEMORY]
   --mailru-user value                   User name (usually email). [$MAILRU_USER]

```
