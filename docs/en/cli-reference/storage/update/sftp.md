# SSH/SFTP

{% code fullWidth="true" %}
```
NAME:
   singularity storage update sftp - SSH/SFTP

USAGE:
   singularity storage update sftp [command options] <name|id>

DESCRIPTION:
   --host
      SSH host to connect to.
      
      E.g. "example.com".

   --user
      SSH username.

   --port
      SSH port number.

   --pass
      SSH password, leave blank to use ssh-agent.

   --key-pem
      Raw PEM-encoded private key.
      
      If specified, will override key_file parameter.

   --key-file
      Path to PEM-encoded private key file.
      
      Leave blank or set key-use-agent to use ssh-agent.
      
      Leading `~` will be expanded in the file name as will environment variables such as `${RCLONE_CONFIG_DIR}`.

   --key-file-pass
      The passphrase to decrypt the PEM-encoded private key file.
      
      Only PEM encrypted key files (old OpenSSH format) are supported. Encrypted keys
      in the new OpenSSH format can't be used.

   --pubkey-file
      Optional path to public key file.
      
      Set this if you have a signed certificate you want to use for authentication.
      
      Leading `~` will be expanded in the file name as will environment variables such as `${RCLONE_CONFIG_DIR}`.

   --known-hosts-file
      Optional path to known_hosts file.
      
      Set this value to enable server host key validation.
      
      Leading `~` will be expanded in the file name as will environment variables such as `${RCLONE_CONFIG_DIR}`.

      Examples:
         | ~/.ssh/known_hosts | Use OpenSSH's known_hosts file.

   --key-use-agent
      When set forces the usage of the ssh-agent.
      
      When key-file is also set, the ".pub" file of the specified key-file is read and only the associated key is
      requested from the ssh-agent. This allows to avoid `Too many authentication failures for *username*` errors
      when the ssh-agent contains many keys.

   --use-insecure-cipher
      Enable the use of insecure ciphers and key exchange methods.
      
      This enables the use of the following insecure ciphers and key exchange methods:
      
      - aes128-cbc
      - aes192-cbc
      - aes256-cbc
      - 3des-cbc
      - diffie-hellman-group-exchange-sha256
      - diffie-hellman-group-exchange-sha1
      
      Those algorithms are insecure and may allow plaintext data to be recovered by an attacker.
      
      This must be false if you use either ciphers or key_exchange advanced options.
      

      Examples:
         | false | Use default Cipher list.
         | true  | Enables the use of the aes128-cbc cipher and diffie-hellman-group-exchange-sha256, diffie-hellman-group-exchange-sha1 key exchange.

   --disable-hashcheck
      Disable the execution of SSH commands to determine if remote file hashing is available.
      
      Leave blank or set to false to enable hashing (recommended), set to true to disable hashing.

   --ask-password
      Allow asking for SFTP password when needed.
      
      If this is set and no password is supplied then rclone will:
      - ask for a password
      - not contact the ssh agent
      

   --path-override
      Override path used by SSH shell commands.
      
      This allows checksum calculation when SFTP and SSH paths are
      different. This issue affects among others Synology NAS boxes.
      
      E.g. if shared folders can be found in directories representing volumes:
      
          rclone sync /home/local/directory remote:/directory --sftp-path-override /volume2/directory
      
      E.g. if home directory can be found in a shared folder called "home":
      
          rclone sync /home/local/directory remote:/home/directory --sftp-path-override /volume1/homes/USER/directory

   --set-modtime
      Set the modified time on the remote if set.

   --shell-type
      The type of SSH shell on remote server, if any.
      
      Leave blank for autodetect.

      Examples:
         | none       | No shell access
         | unix       | Unix shell
         | powershell | PowerShell
         | cmd        | Windows Command Prompt

   --md5sum-command
      The command used to read md5 hashes.
      
      Leave blank for autodetect.

   --sha1sum-command
      The command used to read sha1 hashes.
      
      Leave blank for autodetect.

   --skip-links
      Set to skip any symlinks and any other non regular files.

   --subsystem
      Specifies the SSH2 subsystem on the remote host.

   --server-command
      Specifies the path or command to run a sftp server on the remote host.
      
      The subsystem option is ignored when server_command is defined.

   --use-fstat
      If set use fstat instead of stat.
      
      Some servers limit the amount of open files and calling Stat after opening
      the file will throw an error from the server. Setting this flag will call
      Fstat instead of Stat which is called on an already open file handle.
      
      It has been found that this helps with IBM Sterling SFTP servers which have
      "extractability" level set to 1 which means only 1 file can be opened at
      any given time.
      

   --disable-concurrent-reads
      If set don't use concurrent reads.
      
      Normally concurrent reads are safe to use and not using them will
      degrade performance, so this option is disabled by default.
      
      Some servers limit the amount number of times a file can be
      downloaded. Using concurrent reads can trigger this limit, so if you
      have a server which returns
      
          Failed to copy: file does not exist
      
      Then you may need to enable this flag.
      
      If concurrent reads are disabled, the use_fstat option is ignored.
      

   --disable-concurrent-writes
      If set don't use concurrent writes.
      
      Normally rclone uses concurrent writes to upload files. This improves
      the performance greatly, especially for distant servers.
      
      This option disables concurrent writes should that be necessary.
      

   --idle-timeout
      Max time before closing idle connections.
      
      If no connections have been returned to the connection pool in the time
      given, rclone will empty the connection pool.
      
      Set to 0 to keep connections indefinitely.
      

   --chunk-size
      Upload and download chunk size.
      
      This controls the maximum size of payload in SFTP protocol packets.
      The RFC limits this to 32768 bytes (32k), which is the default. However,
      a lot of servers support larger sizes, typically limited to a maximum
      total package size of 256k, and setting it larger will increase transfer
      speed dramatically on high latency links. This includes OpenSSH, and,
      for example, using the value of 255k works well, leaving plenty of room
      for overhead while still being within a total packet size of 256k.
      
      Make sure to test thoroughly before using a value higher than 32k,
      and only use it if you always connect to the same server or after
      sufficiently broad testing. If you get errors such as
      "failed to send packet payload: EOF", lots of "connection lost",
      or "corrupted on transfer", when copying a larger file, try lowering
      the value. The server run by [rclone serve sftp](/commands/rclone_serve_sftp)
      sends packets with standard 32k maximum payload so you must not
      set a different chunk_size when downloading files, but it accepts
      packets up to the 256k total size, so for uploads the chunk_size
      can be set as for the OpenSSH example above.
      

   --concurrency
      The maximum number of outstanding requests for one file
      
      This controls the maximum number of outstanding requests for one file.
      Increasing it will increase throughput on high latency links at the
      cost of using more memory.
      

   --set-env
      Environment variables to pass to sftp and commands
      
      Set environment variables in the form:
      
          VAR=value
      
      to be passed to the sftp client and to any commands run (eg md5sum).
      
      Pass multiple variables space separated, eg
      
          VAR1=value VAR2=value
      
      and pass variables with spaces in in quotes, eg
      
          "VAR3=value with space" "VAR4=value with space" VAR5=nospacehere
      
      

   --ciphers
      Space separated list of ciphers to be used for session encryption, ordered by preference.
      
      At least one must match with server configuration. This can be checked for example using ssh -Q cipher.
      
      This must not be set if use_insecure_cipher is true.
      
      Example:
      
          aes128-ctr aes192-ctr aes256-ctr aes128-gcm@openssh.com aes256-gcm@openssh.com
      

   --key-exchange
      Space separated list of key exchange algorithms, ordered by preference.
      
      At least one must match with server configuration. This can be checked for example using ssh -Q kex.
      
      This must not be set if use_insecure_cipher is true.
      
      Example:
      
          sntrup761x25519-sha512@openssh.com curve25519-sha256 curve25519-sha256@libssh.org ecdh-sha2-nistp256
      

   --macs
      Space separated list of MACs (message authentication code) algorithms, ordered by preference.
      
      At least one must match with server configuration. This can be checked for example using ssh -Q mac.
      
      Example:
      
          umac-64-etm@openssh.com umac-128-etm@openssh.com hmac-sha2-256-etm@openssh.com
      


OPTIONS:
   --disable-hashcheck    Disable the execution of SSH commands to determine if remote file hashing is available. (default: false) [$DISABLE_HASHCHECK]
   --help, -h             show help
   --host value           SSH host to connect to. [$HOST]
   --key-file value       Path to PEM-encoded private key file. [$KEY_FILE]
   --key-file-pass value  The passphrase to decrypt the PEM-encoded private key file. [$KEY_FILE_PASS]
   --key-pem value        Raw PEM-encoded private key. [$KEY_PEM]
   --key-use-agent        When set forces the usage of the ssh-agent. (default: false) [$KEY_USE_AGENT]
   --pass value           SSH password, leave blank to use ssh-agent. [$PASS]
   --port value           SSH port number. (default: 22) [$PORT]
   --pubkey-file value    Optional path to public key file. [$PUBKEY_FILE]
   --use-insecure-cipher  Enable the use of insecure ciphers and key exchange methods. (default: false) [$USE_INSECURE_CIPHER]
   --user value           SSH username. (default: "$USER") [$USER]

   Advanced

   --ask-password               Allow asking for SFTP password when needed. (default: false) [$ASK_PASSWORD]
   --chunk-size value           Upload and download chunk size. (default: "32Ki") [$CHUNK_SIZE]
   --ciphers value              Space separated list of ciphers to be used for session encryption, ordered by preference. [$CIPHERS]
   --concurrency value          The maximum number of outstanding requests for one file (default: 64) [$CONCURRENCY]
   --disable-concurrent-reads   If set don't use concurrent reads. (default: false) [$DISABLE_CONCURRENT_READS]
   --disable-concurrent-writes  If set don't use concurrent writes. (default: false) [$DISABLE_CONCURRENT_WRITES]
   --idle-timeout value         Max time before closing idle connections. (default: "1m0s") [$IDLE_TIMEOUT]
   --key-exchange value         Space separated list of key exchange algorithms, ordered by preference. [$KEY_EXCHANGE]
   --known-hosts-file value     Optional path to known_hosts file. [$KNOWN_HOSTS_FILE]
   --macs value                 Space separated list of MACs (message authentication code) algorithms, ordered by preference. [$MACS]
   --md5sum-command value       The command used to read md5 hashes. [$MD5SUM_COMMAND]
   --path-override value        Override path used by SSH shell commands. [$PATH_OVERRIDE]
   --server-command value       Specifies the path or command to run a sftp server on the remote host. [$SERVER_COMMAND]
   --set-env value              Environment variables to pass to sftp and commands [$SET_ENV]
   --set-modtime                Set the modified time on the remote if set. (default: true) [$SET_MODTIME]
   --sha1sum-command value      The command used to read sha1 hashes. [$SHA1SUM_COMMAND]
   --shell-type value           The type of SSH shell on remote server, if any. [$SHELL_TYPE]
   --skip-links                 Set to skip any symlinks and any other non regular files. (default: false) [$SKIP_LINKS]
   --subsystem value            Specifies the SSH2 subsystem on the remote host. (default: "sftp") [$SUBSYSTEM]
   --use-fstat                  If set use fstat instead of stat. (default: false) [$USE_FSTAT]

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
   --client-use-server-mod-time                     Use server modified time if possible (default: false)
   --client-user-agent value                        Set the user-agent to a specified string. To remove, use empty string. (default: rclone/v1.62.2-DEV)

   Retry Strategy

   --client-low-level-retries value  Maximum number of retries for low-level client errors (default: 10)
   --client-retry-backoff value      The constant delay backoff for retrying IO read errors (default: 1s)
   --client-retry-backoff-exp value  The exponential delay backoff for retrying IO read errors (default: 1.0)
   --client-retry-delay value        The initial delay before retrying IO read errors (default: 1s)
   --client-retry-max value          Max number of retries for IO read errors (default: 10)
   --client-skip-inaccessible        Skip inaccessible files when opening (default: false)

```
{% endcode %}
