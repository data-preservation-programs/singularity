# SSH/SFTP

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add sftp - SSH/SFTP

USAGE:
   singularity datasource add sftp [command options] <dataset_name> <source_path>

DESCRIPTION:
   --sftp-ask-password
      Allow asking for SFTP password when needed.
      
      If this is set and no password is supplied then rclone will:
      - ask for a password
      - not contact the ssh agent
      

   --sftp-chunk-size
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
      

   --sftp-ciphers
      Space separated list of ciphers to be used for session encryption, ordered by preference.
      
      At least one must match with server configuration. This can be checked for example using ssh -Q cipher.
      
      This must not be set if use_insecure_cipher is true.
      
      Example:
      
          aes128-ctr aes192-ctr aes256-ctr aes128-gcm@openssh.com aes256-gcm@openssh.com
      

   --sftp-concurrency
      The maximum number of outstanding requests for one file
      
      This controls the maximum number of outstanding requests for one file.
      Increasing it will increase throughput on high latency links at the
      cost of using more memory.
      

   --sftp-disable-concurrent-reads
      If set don't use concurrent reads.
      
      Normally concurrent reads are safe to use and not using them will
      degrade performance, so this option is disabled by default.
      
      Some servers limit the amount number of times a file can be
      downloaded. Using concurrent reads can trigger this limit, so if you
      have a server which returns
      
          Failed to copy: file does not exist
      
      Then you may need to enable this flag.
      
      If concurrent reads are disabled, the use_fstat option is ignored.
      

   --sftp-disable-concurrent-writes
      If set don't use concurrent writes.
      
      Normally rclone uses concurrent writes to upload files. This improves
      the performance greatly, especially for distant servers.
      
      This option disables concurrent writes should that be necessary.
      

   --sftp-disable-hashcheck
      Disable the execution of SSH commands to determine if remote file hashing is available.
      
      Leave blank or set to false to enable hashing (recommended), set to true to disable hashing.

   --sftp-host
      SSH host to connect to.
      
      E.g. "example.com".

   --sftp-idle-timeout
      Max time before closing idle connections.
      
      If no connections have been returned to the connection pool in the time
      given, rclone will empty the connection pool.
      
      Set to 0 to keep connections indefinitely.
      

   --sftp-key-exchange
      Space separated list of key exchange algorithms, ordered by preference.
      
      At least one must match with server configuration. This can be checked for example using ssh -Q kex.
      
      This must not be set if use_insecure_cipher is true.
      
      Example:
      
          sntrup761x25519-sha512@openssh.com curve25519-sha256 curve25519-sha256@libssh.org ecdh-sha2-nistp256
      

   --sftp-key-file
      Path to PEM-encoded private key file.
      
      Leave blank or set key-use-agent to use ssh-agent.
      
      Leading `~` will be expanded in the file name as will environment variables such as `${RCLONE_CONFIG_DIR}`.

   --sftp-key-file-pass
      The passphrase to decrypt the PEM-encoded private key file.
      
      Only PEM encrypted key files (old OpenSSH format) are supported. Encrypted keys
      in the new OpenSSH format can't be used.

   --sftp-key-pem
      Raw PEM-encoded private key.
      
      If specified, will override key_file parameter.

   --sftp-key-use-agent
      When set forces the usage of the ssh-agent.
      
      When key-file is also set, the ".pub" file of the specified key-file is read and only the associated key is
      requested from the ssh-agent. This allows to avoid `Too many authentication failures for *username*` errors
      when the ssh-agent contains many keys.

   --sftp-known-hosts-file
      Optional path to known_hosts file.
      
      Set this value to enable server host key validation.
      
      Leading `~` will be expanded in the file name as will environment variables such as `${RCLONE_CONFIG_DIR}`.

      Examples:
         | ~/.ssh/known_hosts | Use OpenSSH's known_hosts file.

   --sftp-macs
      Space separated list of MACs (message authentication code) algorithms, ordered by preference.
      
      At least one must match with server configuration. This can be checked for example using ssh -Q mac.
      
      Example:
      
          umac-64-etm@openssh.com umac-128-etm@openssh.com hmac-sha2-256-etm@openssh.com
      

   --sftp-md5sum-command
      The command used to read md5 hashes.
      
      Leave blank for autodetect.

   --sftp-pass
      SSH password, leave blank to use ssh-agent.

   --sftp-path-override
      Override path used by SSH shell commands.
      
      This allows checksum calculation when SFTP and SSH paths are
      different. This issue affects among others Synology NAS boxes.
      
      E.g. if shared folders can be found in directories representing volumes:
      
          rclone sync /home/local/directory remote:/directory --sftp-path-override /volume2/directory
      
      E.g. if home directory can be found in a shared folder called "home":
      
          rclone sync /home/local/directory remote:/home/directory --sftp-path-override /volume1/homes/USER/directory

   --sftp-port
      SSH port number.

   --sftp-pubkey-file
      Optional path to public key file.
      
      Set this if you have a signed certificate you want to use for authentication.
      
      Leading `~` will be expanded in the file name as will environment variables such as `${RCLONE_CONFIG_DIR}`.

   --sftp-server-command
      Specifies the path or command to run a sftp server on the remote host.
      
      The subsystem option is ignored when server_command is defined.

   --sftp-set-env
      Environment variables to pass to sftp and commands
      
      Set environment variables in the form:
      
          VAR=value
      
      to be passed to the sftp client and to any commands run (eg md5sum).
      
      Pass multiple variables space separated, eg
      
          VAR1=value VAR2=value
      
      and pass variables with spaces in in quotes, eg
      
          "VAR3=value with space" "VAR4=value with space" VAR5=nospacehere
      
      

   --sftp-set-modtime
      Set the modified time on the remote if set.

   --sftp-sha1sum-command
      The command used to read sha1 hashes.
      
      Leave blank for autodetect.

   --sftp-shell-type
      The type of SSH shell on remote server, if any.
      
      Leave blank for autodetect.

      Examples:
         | none       | No shell access
         | unix       | Unix shell
         | powershell | PowerShell
         | cmd        | Windows Command Prompt

   --sftp-skip-links
      Set to skip any symlinks and any other non regular files.

   --sftp-subsystem
      Specifies the SSH2 subsystem on the remote host.

   --sftp-use-fstat
      If set use fstat instead of stat.
      
      Some servers limit the amount of open files and calling Stat after opening
      the file will throw an error from the server. Setting this flag will call
      Fstat instead of Stat which is called on an already open file handle.
      
      It has been found that this helps with IBM Sterling SFTP servers which have
      "extractability" level set to 1 which means only 1 file can be opened at
      any given time.
      

   --sftp-use-insecure-cipher
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

   --sftp-user
      SSH username.


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)

   Options for sftp

   --sftp-ask-password value               Allow asking for SFTP password when needed. (default: "false") [$SFTP_ASK_PASSWORD]
   --sftp-chunk-size value                 Upload and download chunk size. (default: "32Ki") [$SFTP_CHUNK_SIZE]
   --sftp-ciphers value                    Space separated list of ciphers to be used for session encryption, ordered by preference. [$SFTP_CIPHERS]
   --sftp-concurrency value                The maximum number of outstanding requests for one file (default: "64") [$SFTP_CONCURRENCY]
   --sftp-disable-concurrent-reads value   If set don't use concurrent reads. (default: "false") [$SFTP_DISABLE_CONCURRENT_READS]
   --sftp-disable-concurrent-writes value  If set don't use concurrent writes. (default: "false") [$SFTP_DISABLE_CONCURRENT_WRITES]
   --sftp-disable-hashcheck value          Disable the execution of SSH commands to determine if remote file hashing is available. (default: "false") [$SFTP_DISABLE_HASHCHECK]
   --sftp-host value                       SSH host to connect to. [$SFTP_HOST]
   --sftp-idle-timeout value               Max time before closing idle connections. (default: "1m0s") [$SFTP_IDLE_TIMEOUT]
   --sftp-key-exchange value               Space separated list of key exchange algorithms, ordered by preference. [$SFTP_KEY_EXCHANGE]
   --sftp-key-file value                   Path to PEM-encoded private key file. [$SFTP_KEY_FILE]
   --sftp-key-file-pass value              The passphrase to decrypt the PEM-encoded private key file. [$SFTP_KEY_FILE_PASS]
   --sftp-key-pem value                    Raw PEM-encoded private key. [$SFTP_KEY_PEM]
   --sftp-key-use-agent value              When set forces the usage of the ssh-agent. (default: "false") [$SFTP_KEY_USE_AGENT]
   --sftp-known-hosts-file value           Optional path to known_hosts file. [$SFTP_KNOWN_HOSTS_FILE]
   --sftp-macs value                       Space separated list of MACs (message authentication code) algorithms, ordered by preference. [$SFTP_MACS]
   --sftp-md5sum-command value             The command used to read md5 hashes. [$SFTP_MD5SUM_COMMAND]
   --sftp-pass value                       SSH password, leave blank to use ssh-agent. [$SFTP_PASS]
   --sftp-path-override value              Override path used by SSH shell commands. [$SFTP_PATH_OVERRIDE]
   --sftp-port value                       SSH port number. (default: "22") [$SFTP_PORT]
   --sftp-pubkey-file value                Optional path to public key file. [$SFTP_PUBKEY_FILE]
   --sftp-server-command value             Specifies the path or command to run a sftp server on the remote host. [$SFTP_SERVER_COMMAND]
   --sftp-set-env value                    Environment variables to pass to sftp and commands [$SFTP_SET_ENV]
   --sftp-set-modtime value                Set the modified time on the remote if set. (default: "true") [$SFTP_SET_MODTIME]
   --sftp-sha1sum-command value            The command used to read sha1 hashes. [$SFTP_SHA1SUM_COMMAND]
   --sftp-shell-type value                 The type of SSH shell on remote server, if any. [$SFTP_SHELL_TYPE]
   --sftp-skip-links value                 Set to skip any symlinks and any other non regular files. (default: "false") [$SFTP_SKIP_LINKS]
   --sftp-subsystem value                  Specifies the SSH2 subsystem on the remote host. (default: "sftp") [$SFTP_SUBSYSTEM]
   --sftp-use-fstat value                  If set use fstat instead of stat. (default: "false") [$SFTP_USE_FSTAT]
   --sftp-use-insecure-cipher value        Enable the use of insecure ciphers and key exchange methods. (default: "false") [$SFTP_USE_INSECURE_CIPHER]
   --sftp-user value                       SSH username. (default: "$USER") [$SFTP_USER]

```
{% endcode %}
