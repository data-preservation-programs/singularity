# Mail.ru Cloud

{% code fullWidth="true" %}
```
NAME:
   singularity storage update mailru - Mail.ru Cloud

USAGE:
   singularity storage update mailru [command options] <name|id>

DESCRIPTION:
   --user
      User name (usually email).

   --pass
      Password.
      
      This must be an app password - rclone will not work with your normal
      password. See the Configuration section in the docs for how to make an
      app password.
      

   --speedup-enable
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

   --speedup-file-patterns
      Comma separated list of file name patterns eligible for speedup (put by hash).
      
      Patterns are case insensitive and can contain '*' or '?' meta characters.

      Examples:
         | <unset>                 | Empty list completely disables speedup (put by hash).
         | *                       | All files will be attempted for speedup.
         | *.mkv,*.avi,*.mp4,*.mp3 | Only common audio/video files will be tried for put by hash.
         | *.zip,*.gz,*.rar,*.pdf  | Only common archives or PDF books will be tried for speedup.

   --speedup-max-disk
      This option allows you to disable speedup (put by hash) for large files.
      
      Reason is that preliminary hashing can exhaust your RAM or disk space.

      Examples:
         | 0  | Completely disable speedup (put by hash).
         | 1G | Files larger than 1Gb will be uploaded directly.
         | 3G | Choose this option if you have less than 3Gb free on local disk.

   --speedup-max-memory
      Files larger than the size given below will always be hashed on disk.

      Examples:
         | 0    | Preliminary hashing will always be done in a temporary disk location.
         | 32M  | Do not dedicate more than 32Mb RAM for preliminary hashing.
         | 256M | You have at most 256Mb RAM free for hash calculations.

   --check-hash
      What should copy do if file checksum is mismatched or invalid.

      Examples:
         | true  | Fail with error.
         | false | Ignore and continue.

   --user-agent
      HTTP user agent used internally by client.
      
      Defaults to "rclone/VERSION" or "--user-agent" provided on command line.

   --quirks
      Comma separated list of internal maintenance flags.
      
      This option must not be used by an ordinary user. It is intended only to
      facilitate remote troubleshooting of backend issues. Strict meaning of
      flags is not documented and not guaranteed to persist between releases.
      Quirks will be removed when the backend grows stable.
      Supported quirks: atomicmkdir binlist unknowndirs

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --help, -h        show help
   --pass value      Password. [$PASS]
   --speedup-enable  Skip full upload if there is another file with same data hash. (default: true) [$SPEEDUP_ENABLE]
   --user value      User name (usually email). [$USER]

   Advanced

   --check-hash                   What should copy do if file checksum is mismatched or invalid. (default: true) [$CHECK_HASH]
   --encoding value               The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --quirks value                 Comma separated list of internal maintenance flags. [$QUIRKS]
   --speedup-file-patterns value  Comma separated list of file name patterns eligible for speedup (put by hash). (default: "*.mkv,*.avi,*.mp4,*.mp3,*.zip,*.gz,*.rar,*.pdf") [$SPEEDUP_FILE_PATTERNS]
   --speedup-max-disk value       This option allows you to disable speedup (put by hash) for large files. (default: "3Gi") [$SPEEDUP_MAX_DISK]
   --speedup-max-memory value     Files larger than the size given below will always be hashed on disk. (default: "32Mi") [$SPEEDUP_MAX_MEMORY]
   --user-agent value             HTTP user agent used internally by client. [$USER_AGENT]

   HTTP Client Config

   --client-ca-cert value                           Path to CA certificate used to verify servers. To remove, use empty string.
   --client-cert value                              Path to Client SSL certificate (PEM) for mutual TLS auth. To remove, use empty string.
   --client-connect-timeout value                   HTTP Client Connect timeout (default: 1m0s)
   --client-expect-continue-timeout value           Timeout when using expect / 100-continue in HTTP (default: 1s)
   --client-header value [ --client-header value ]  Set HTTP header for all transactions (i.e. key=value). This will replace the existing header values. To remove a header, use --http-header "key="". To remove all headers, use --http-header ""
   --client-insecure-skip-verify                    Do not verify the server SSL certificate (insecure) (default: false)
   --client-key value                               Path to Client SSL private key (PEM) for mutual TLS auth. To remove, use empty string.
   --client-no-gzip                                 Don't set Accept-Encoding: gzip (default: false)
   --client-timeout value                           IO idle timeout (default: 5m0s)
   --client-user-agent value                        Set the user-agent to a specified string. To remove, use empty string. (default: rclone/v1.62.2-DEV)

   Retry Strategy

   --client-retry-backoff value      The constant delay backoff for retrying IO read errors (default: 1s)
   --client-retry-backoff-exp value  The exponential delay backoff for retrying IO read errors (default: 1.0)
   --client-retry-delay value        The initial delay before retrying IO read errors (default: 1s)
   --client-retry-max value          Max number of retries for IO read errors (default: 10)

```
{% endcode %}
