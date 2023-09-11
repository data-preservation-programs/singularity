# SSH/SFTP

{% code fullWidth="true" %}
```
名称:
   singularity 存储 更新 sftp - SSH/SFTP

用法:
   singularity 存储 更新 sftp [命令选项] <名称|ID>

说明:
   --host
      SSH 主机名称。
      
      例如："example.com"。

   --user
      SSH 用户名。

   --port
      SSH 端口号。

   --pass
      SSH 密码，留空以使用 SSH 代理。

   --key-pem
      原始的 PEM 编码私钥。
      
      如果指定了此参数，则会覆盖 key_file 参数。

   --key-file
      PEM 编码私钥文件的路径。
      
      留空或将 key_use_agent 设置为使用 SSH 代理。
      
      文件名中的 `~` 字符会扩展为绝对路径，环境变量如 `${RCLONE_CONFIG_DIR}` 也会被扩展。

   --key-file-pass
      解密 PEM 编码私钥文件所需的密码。
      
      仅支持 PEM 密码保护的私钥文件（旧的 OpenSSH 格式）。新的 OpenSSH 格式的加密密钥无法使用。

   --pubkey-file
      可选的公钥文件路径。
      
      如果要使用已签名的证书进行身份验证，则设置此参数。
      
      文件名中的 `~` 字符会扩展为绝对路径，环境变量如 `${RCLONE_CONFIG_DIR}` 也会被扩展。

   --known-hosts-file
      可选的已知主机文件路径。
      
      设置该值以启用服务器主机密钥验证。
      
      文件名中的 `~` 字符会扩展为绝对路径，环境变量如 `${RCLONE_CONFIG_DIR}` 也会被扩展。

      示例:
         | ~/.ssh/known_hosts | 使用 OpenSSH 的已知主机文件。

   --key-use-agent
      当设置时，强制使用 ssh 代理。
      
      当同时设置了 key-file 时，将读取指定密钥文件的 ".pub" 文件，并且只请求 ssh 代理中关联的密钥。这有助于避免 ssh 代理中包含多个密钥时出现“Too many authentication failures for *username*”错误。

   --use-insecure-cipher
      启用使用不安全的加密算法和密钥交换方法。
      
      这将启用以下不安全的加密算法和密钥交换方法：
      
      - aes128-cbc
      - aes192-cbc
      - aes256-cbc
      - 3des-cbc
      - diffie-hellman-group-exchange-sha256
      - diffie-hellman-group-exchange-sha1
      
      这些算法是不安全的，可能允许攻击者恢复明文数据。
      
      如果使用了 ciphers 或 key_exchange 高级选项，则必须将此选项设置为 false。
      

      示例:
         | false | 使用默认的加密算法列表。
         | true  | 启用 aes128-cbc 加密算法和 diffie-hellman-group-exchange-sha256、diffie-hellman-group-exchange-sha1 密钥交换。

   --disable-hashcheck
      禁用执行 SSH 命令来确定是否支持远程文件哈希计算。
      
      留空或设置为 false 以启用哈希计算（建议），设置为 true 以禁用哈希计算。

   --ask-password
      允许在需要时询问 SFTP 密码。
      
      如果设置了此参数并且未提供密码，则 rclone 将：
      - 询问密码
      - 不使用 ssh 代理进行验证
      

   --path-override
      覆盖 SSH shell 命令要使用的路径。
      
      这允许在 SFTP 和 SSH 路径不同的情况下进行校验和计算。这个问题会影响 Synology NAS 等设备。
      
      例如，如果共享文件夹位于表示卷的目录中：
      
          rclone sync /home/local/directory remote:/directory --sftp-path-override /volume2/directory
      
      例如，如果主目录位于名为 "home" 的共享文件夹中：
      
          rclone sync /home/local/directory remote:/home/directory --sftp-path-override /volume1/homes/USER/directory

   --set-modtime
      如果设置了，则在远程服务器上设置修改时间。

   --shell-type
      远程服务器上的 SSH shell 类型，如果有的话。
      
      留空以自动检测。

      示例:
         | none       | 不允许 shell 访问
         | unix       | Unix shell
         | powershell | PowerShell
         | cmd        | Windows 命令提示符

   --md5sum-command
      用于读取 md5 哈希值的命令。
      
      留空以自动检测。

   --sha1sum-command
      用于读取 sha1 哈希值的命令。
      
      留空以自动检测。

   --skip-links
      设置为跳过任何符号链接和其他非常规文件。

   --subsystem
      指定远程主机上的 SSH2 子系统。

   --server-command
      指定在远程主机上运行 sftp 服务器的路径或命令。
      
      定义了 server_command 参数后，子系统选项将被忽略。

   --use-fstat
      如果设置了，则使用 fstat 替代 stat。
      
      有些服务器限制了打开的文件数量，在打开文件后调用 Stat 时会从服务器抛出错误。设置此标志将在已打开的文件句柄上调用 Fstat，而不是调用 Stat。
      
      其实际上有助于解决 IBM Sterling SFTP 服务器的问题，该服务器将“可提取性”级别设置为 1，这意味着任意给定时间只能打开 1 个文件。
      

   --disable-concurrent-reads
      如果设置了，则不使用并发读取。
      
      通常，并发读取是安全的并且不使用它们会降低性能，因此此选项默认禁用。
      
      有些服务器限制了文件的下载次数。使用并发读取可能会触发此限制，所以如果您的服务器返回
      
          Failed to copy: file does not exist
      
      那么您可能需要启用此标志。
      
      如果禁用了并发读取，则 use_fstat 选项将被忽略。
      

   --disable-concurrent-writes
      如果设置了，则不使用并发写入。
      
      通常，rclone 使用并发写入来上传文件。这将显著提高性能，特别是对于远程服务器。
      
      如果需要，此选项会禁用并发写入。
      

   --idle-timeout
      空闲连接关闭之前的最大时间。
      
      如果在给定的时间内没有返回连接到连接池，则 rclone 将清空连接池。
      
      设置为 0 以无限期保持连接。
      

   --chunk-size
      上传和下载的块大小。
      
      这控制 SFTP 协议数据包中有效载荷的最大大小。RFC 将其限制为 32768 字节（32k），这是默认值。
      然而，很多服务器支持更大的大小，通常限制为最大总包大小为 256k，并且将其设置得更大将极大地增加高延迟链路上的传输速度。
      这包括 OpenSSH，在这种情况下，使用值 255k 效果很好，留出了足够的空间用于开销，同时仍然在 256k 的总包大小内。
      
      在使用大于 32k 的值之前，请务必进行彻底测试，并且只在始终连接到相同服务器或经过足够广泛的测试后使用。如果出现“failed to send packet payload: EOF”、“connection lost”等错误，则尝试降低此值。由[rclone serve sftp](/commands/rclone_serve_sftp)运行的服务器将发送标准的 32k 最大有效载荷大小的数据包，因此在下载文件时不能设置不同的 chunk_size，但它接受达到 256k 总大小的数据包，因此对于上传，chunk_size 可以设置为上面的 OpenSSH 示例中的值。
      

   --concurrency
      一个文件的最大未完成请求数量
      
      这控制一个文件的最大未完成请求数量。将其增加可以提高高延迟链路上的吞吐量，但会占用更多的内存。
      

   --set-env
      要传递给 sftp 和命令的环境变量
      
      以以下形式设置环境变量：
      
          VAR=value
      
      以传递给 sftp 客户端和运行的任何命令（例如 md5sum）。
      
      通过以空格分隔多个变量，例如
      
          VAR1=value VAR2=value
      
      以及以引号中带有空格的方式传递变量，例如
      
          "VAR3=value with space" "VAR4=value with space" VAR5=nospacehere
      
      

   --ciphers
      用空格分隔的密码列表，按优先级排序用于会话加密。
      
      至少有一个密码必须与服务器配置匹配。可以使用 ssh -Q cipher 等命令来检查。
      
      如果 use_insecure_cipher 为 true，则不能设置此选项。
      
      示例:
      
          aes128-ctr aes192-ctr aes256-ctr aes128-gcm@openssh.com aes256-gcm@openssh.com
      

   --key-exchange
      用空格分隔的密钥交换算法列表，按优先级排序。
      
      至少有一个算法必须与服务器配置匹配。可以使用 ssh -Q kex 等命令来检查。
      
      如果 use_insecure_cipher 为 true，则不能设置此选项。
      
      示例:
      
          sntrup761x25519-sha512@openssh.com curve25519-sha256 curve25519-sha256@libssh.org ecdh-sha2-nistp256
      

   --macs
      用空格分隔的 MAC（消息认证码）算法列表，按优先级排序。
      
      至少有一个算法必须与服务器配置匹配。可以使用 ssh -Q mac 等命令来检查。
      
      示例:
      
          umac-64-etm@openssh.com umac-128-etm@openssh.com hmac-sha2-256-etm@openssh.com
      


选项:
   --disable-hashcheck    禁用执行 SSH 命令来确定是否支持远程文件哈希计算。 (默认值: false) [$DISABLE_HASHCHECK]
   --help, -h             显示帮助
   --host value           SSH 主机名称。 [$HOST]
   --key-file value       PEM 编码私钥文件的路径。 [$KEY_FILE]
   --key-file-pass value  解密 PEM 编码私钥文件所需的密码。 [$KEY_FILE_PASS]
   --key-pem value        原始的 PEM 编码私钥。 [$KEY_PEM]
   --key-use-agent        当设置时，强制使用 ssh 代理。 (默认值: false) [$KEY_USE_AGENT]
   --pass value           SSH 密码，留空以使用 ssh 代理。 [$PASS]
   --port value           SSH 端口号。 (默认值: 22) [$PORT]
   --pubkey-file value    可选的公钥文件路径。 [$PUBKEY_FILE]
   --use-insecure-cipher  启用使用不安全的加密算法和密钥交换方法。 (默认值: false) [$USE_INSECURE_CIPHER]
   --user value           SSH 用户名。 (默认值: "$USER") [$USER]

   高级选项

   --ask-password               允许在需要时询问 SFTP 密码。 (默认值: false) [$ASK_PASSWORD]
   --chunk-size value           上传和下载的块大小。 (默认值: "32Ki") [$CHUNK_SIZE]
   --ciphers value              用空格分隔的密码列表，按优先级排序用于会话加密。 [$CIPHERS]
   --concurrency value          一个文件的最大未完成请求数量 (默认值: 64) [$CONCURRENCY]
   --disable-concurrent-reads   如果设置了，则不使用并发读取。 (默认值: false) [$DISABLE_CONCURRENT_READS]
   --disable-concurrent-writes  如果设置了，则不使用并发写入。 (默认值: false) [$DISABLE_CONCURRENT_WRITES]
   --idle-timeout value         空闲连接关闭之前的最大时间。 (默认值: "1m0s") [$IDLE_TIMEOUT]
   --key-exchange value         用空格分隔的密钥交换算法列表，按优先级排序。 [$KEY_EXCHANGE]
   --known-hosts-file value     可选的已知主机文件路径。 [$KNOWN_HOSTS_FILE]
   --macs value                 用空格分隔的 MAC（消息认证码）算法列表，按优先级排序。 [$MACS]
   --md5sum-command value       用于读取 md5 哈希值的命令。 [$MD5SUM_COMMAND]
   --path-override value        覆盖 SSH shell 命令要使用的路径。 [$PATH_OVERRIDE]
   --server-command value       指定在远程主机上运行 sftp 服务器的路径或命令。 [$SERVER_COMMAND]
   --set-env value              要传递给 sftp 和命令的环境变量 [$SET_ENV]
   --set-modtime                如果设置了，则在远程服务器上设置修改时间。 (默认值: true) [$SET_MODTIME]
   --sha1sum-command value      用于读取 sha1 哈希值的命令。 [$SHA1SUM_COMMAND]
   --shell-type value           远程服务器上的 SSH shell 类型，如果有的话。 [$SHELL_TYPE]
   --skip-links                 设置为跳过任何符号链接和其他非常规文件。 (默认值: false) [$SKIP_LINKS]
   --subsystem value            指定远程主机上的 SSH2 子系统。 (默认值: "sftp") [$SUBSYSTEM]
   --use-fstat                  如果设置了，则使用 fstat 替代 stat。 (默认值: false) [$USE_FSTAT]

```
{% endcode %}