# SSH/SFTP

{% code fullWidth="true" %}
```
名称:
   singularity 数据源 添加 sftp - SSH/SFTP

用法:
   singularity 数据源 添加 sftp [命令选项] <数据集名称> <源路径>

描述:
   --sftp-ask-password
      当需要时允许询问 SFTP 密码。
      
      如果设置了这个选项并且没有提供密码，那么 rclone 将会：
      - 询问密码
      - 不会联系到 ssh 代理
      

   --sftp-chunk-size
      上传和下载的分块大小。
      
      这个选项控制 SFTP 协议包的有效载荷的最大大小。
      RFC RFC 限制为 32768 字节 (32k)，这是默认值。然而，很多服务器支持更大的大小，
      通常限制为最大 256k 的总包大小，将其设置更大将显著增加高延迟链路上的传输速度。
      这包括 OpenSSH，在这个例子中，使用 255k 的值效果很好，这样还留下了足够的
      空间进行开销，同时仍然位于 256k 的总包大小内。
      
      请确保在使用比 32k 更大的值之前进行彻底的测试，并且只在总是连接到同一个
      服务器或经过足够广泛的测试后才使用。如果在复制较大的文件时遇到如下错误：
      "failed to send packet payload: EOF"、很多 "connection lost" 或 "corrupted on transfer"，
      请尝试降低这个值。由[rclone serve sftp](/commands/rclone_serve_sftp) 运行的服务器
      发送具有标准的 32k 最大有效载荷的数据包，因此当下载文件时，绝不要设置
      不同的 chunk_size，但是它接受多达 256k 总大小的数据包，因此对于上传，
      chunk_size 可以设置为与上面 OpenSSH 示例中一样。
      

   --sftp-ciphers
      以空格分隔的密码列表，按首选顺序排列，用于会话加密。
      
      必须至少有一个与服务器配置匹配。可以使用 ssh -Q cipher 命令来检查。
      
      如果 use_insecure_cipher 为 true，则不得设置此选项。
      
      示例：
      
          aes128-ctr aes192-ctr aes256-ctr aes128-gcm@openssh.com aes256-gcm@openssh.com
      

   --sftp-concurrency
      单个文件的最大未完成请求数
      
      这个选项控制单个文件的最大未完成请求数。
      增加该值将提高高延迟链路上的吞吐量，但代价是使用更多的内存。
      

   --sftp-disable-concurrent-reads
      如果设置了，则不使用并发读取。
      
      通常，并发读取是安全的，并且不使用它会降低性能，因此此选项默认禁用。
      
      一些服务器限制文件可以下载的次数。使用并发读取可能会触发此限制，
      因此，如果您有一个返回
      
          Failed to copy: file does not exist
      
      的服务器，则可能需要启用此标志。
      
      如果禁用了并发读取，则忽略 use_fstat 选项。
      

   --sftp-disable-concurrent-writes
      如果设置了，则不使用并发写入。
      
      通常，rclone 使用并发写入来上传文件。这极大地提高了性能，特别是对于远程服务器而言。
      
      如果必要，此选项将禁用并发写入。
      

   --sftp-disable-hashcheck
      禁用执行 SSH 命令以确定是否可用远程文件散列。
      
      留空或设置为 false 以启用散列 (推荐)，设置为 true 以禁用散列。

   --sftp-host
      要连接的 SSH 主机。
      
      例如 "example.com"。

   --sftp-idle-timeout
      空闲连接关闭之前的最大时间。
      
      如果在给定的时间内没有将连接返回到连接池中，rclone 将清空连接池。
      
      设置为 0 以无限期保留连接。
      

   --sftp-key-exchange
      以空格分隔的密钥交换算法列表，按首选顺序排列。
      
      必须至少有一个与服务器配置匹配。可以使用 ssh -Q kex 命令来检查。
      
      如果 use_insecure_cipher 为 true，则不得设置此选项。
      
      示例：
      
          sntrup761x25519-sha512@openssh.com curve25519-sha256 curve25519-sha256@libssh.org ecdh-sha2-nistp256
      

   --sftp-key-file
      PEM 编码的私钥文件的路径。
      
      留空或将 key_use_agent 设置为使用 ssh-agent。
      
      主文件名中的 `~` 会被展开为文件名，环境变量例如 `${RCLONE_CONFIG_DIR}` 也会被展开。

   --sftp-key-file-pass
      用于解密 PEM 编码的私钥文件的密码。
      
      仅支持 PEM 加密的旧版本 OpenSSH 格式的密钥文件。不能使用新的 OpenSSH 格式的加密密钥。

   --sftp-key-pem
      原始的 PEM 编码的私钥。
      
      如果指定了此选项，将覆盖 key_file 参数。

   --sftp-key-use-agent
      当设置时，强制使用 ssh-agent。
      
      当 key-file 也被设置时，将会读取指定的 key-file 的 ".pub" 文件，并且只请求来自 ssh-agent 的关联密钥。
      这可以避免当 ssh-agent 包含多个密钥时出现  `Too many authentication failures for *username*` 错误。

   --sftp-known-hosts-file
      已知的主机文件的可选路径。
      
      设置这个值以启用服务器主机密钥验证。
      
      主文件名中的 `~` 会被展开为文件名，环境变量例如 `${RCLONE_CONFIG_DIR}` 也会被展开。

      示例：
         | ~/.ssh/known_hosts | 使用 OpenSSH 的 known_hosts 文件。

   --sftp-macs
      以空格分隔的 MAC (message authentication code) 算法列表，按首选顺序排列。
      
      必须至少有一个与服务器配置匹配。可以使用 ssh -Q mac 命令来检查。
      
      示例：
      
          umac-64-etm@openssh.com umac-128-etm@openssh.com hmac-sha2-256-etm@openssh.com
      

   --sftp-md5sum-command
      用于读取 md5 散列的命令。
      
      留空以自动检测。

   --sftp-pass
      SSH 密码，留空以使用 ssh-agent。

   --sftp-path-override
      覆盖由 SSH shell 命令使用的路径。
      
      这允许在 SFTP 和 SSH 路径不同时进行散列计算。这个问题会影响到包括 Synology NAS 箱子在内的部分设备。
      
      例如，如果共享的文件夹可以在表示卷的目录中找到：
      
          rclone sync /home/local/directory remote:/directory --sftp-path-override /volume2/directory
      
      例如，如果个人文件夹可以在一个名为 "home" 的共享文件夹中找到：
      
          rclone sync /home/local/directory remote:/home/directory --sftp-path-override /volume1/homes/USER/directory

   --sftp-port
      SSH 端口号。

   --sftp-pubkey-file
      证书文件的可选路径。
      
      如果您有一个签名证书要用于身份验证，请设置此选项。
      
      主文件名中的 `~` 会被展开为文件名，环境变量例如 `${RCLONE_CONFIG_DIR}` 也会被展开。

   --sftp-server-command
      指定在远程主机上运行 sftp 服务器路径或命令。
      
      当定义了server_command时，subsystem选项将被忽略。

   --sftp-set-env
      要传递给 sftp 和命令的环境变量
      
      以以下形式设置环境变量：
      
          VAR=value
      
      以向 sftp 客户端和运行的任何命令传递。
      
      使用空格分隔多个变量，例如
      
          VAR1=value VAR2=value
      
      并在引号内传递具有空格的变量，例如
      
          "VAR3=value with space" "VAR4=value with space" VAR5=nospacehere
      
      

   --sftp-set-modtime
      如果设置了，设置远程文件的修改时间。

   --sftp-sha1sum-command
      用于读取 sha1 散列的命令。
      
      留空以自动检测。

   --sftp-shell-type
      远程服务器上的 SSH shell 类型，如果有的话。
      
      留空以自动检测。

      示例：
         | none       | 无 Shell 访问
         | unix       | Unix Shell
         | powershell | PowerShell
         | cmd        | Windows 命令提示符

   --sftp-skip-links
      设置以跳过任何符号链接和任何其他非普通文件。

   --sftp-subsystem
      指定远程主机上的 SSH2 子系统。

   --sftp-use-fstat
      如果设置了，则使用 fstat 而不是 stat。
      
      一些服务器限制打开的文件数量，在打开文件后调用 Stat 会导致服务器抛出错误。
      设置此标志将调用 Fstat 而不是已打开的文件句柄上调用的 Stat。
      
      发现它对于 IBM Sterling SFTP 服务器很有帮助，它的 "extractability" 级别设置为 1，
      这意味着任何给定时间只能打开1个文件。
      

   --sftp-use-insecure-cipher
      启用使用不安全的密码和密钥交换方法。
      
      这将启用使用以下不安全的密码和密钥交换方法：
      
      - aes128-cbc
      - aes192-cbc
      - aes256-cbc
      - 3des-cbc
      - diffie-hellman-group-exchange-sha256
      - diffie-hellman-group-exchange-sha1
      
      这些算法是不安全的，可能允许攻击者恢复明文数据。
      
      如果同时使用 ciphers 或 key_exchange 高级选项，则必须将此选项设置为 false。
      

      示例：
         | false | 使用默认密码列表。
         | true  | 启用 aes128-cbc 密码和 diffie-hellman-group-exchange-sha256、diffie-hellman-group-exchange-sha1 密钥交换。

   --sftp-user
      SSH 用户名。


选项:
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险操作] 导出到 CAR 文件后删除数据集的文件。 (默认值: false)
   --rescan-interval value  当距离上次成功扫描的时间间隔超过该间隔时，自动重新扫描源目录 (默认值: 禁用)
   --scanning-state value   设置初始的扫描状态 (默认值: ready)

   sftp 选项

   --sftp-ask-password value               允许需要时询问 SFTP 密码。 (默认值: "false") [$SFTP_ASK_PASSWORD]
   --sftp-chunk-size value                 上传和下载的分块大小。 (默认值: "32Ki") [$SFTP_CHUNK_SIZE]
   --sftp-ciphers value                    以空格分隔的密码列表，按首选顺序排列。 [$SFTP_CIPHERS]
   --sftp-concurrency value                单个文件的最大未完成请求数。 (默认值: "64") [$SFTP_CONCURRENCY]
   --sftp-disable-concurrent-reads value   如果设置了，则不使用并发读取。 (默认值: "false") [$SFTP_DISABLE_CONCURRENT_READS]
   --sftp-disable-concurrent-writes value  如果设置了，则不使用并发写入。 (默认值: "false") [$SFTP_DISABLE_CONCURRENT_WRITES]
   --sftp-disable-hashcheck value          禁用执行 SSH 命令以确定是否可用远程文件散列。 (默认值: "false") [$SFTP_DISABLE_HASHCHECK]
   --sftp-host value                       要连接的 SSH 主机。 [$SFTP_HOST]
   --sftp-idle-timeout value               空闲连接关闭之前的最大时间。 (默认值: "1m0s") [$SFTP_IDLE_TIMEOUT]
   --sftp-key-exchange value               以空格分隔的密钥交换算法列表，按首选顺序排列。 [$SFTP_KEY_EXCHANGE]
   --sftp-key-file value                   PEM 编码的私钥文件的路径。 [$SFTP_KEY_FILE]
   --sftp-key-file-pass value              用于解密 PEM 编码的私钥文件的密码。 [$SFTP_KEY_FILE_PASS]
   --sftp-key-pem value                    原始的 PEM 编码的私钥。 [$SFTP_KEY_PEM]
   --sftp-key-use-agent value              当设置时，强制使用 ssh-agent。 (默认值: "false") [$SFTP_KEY_USE_AGENT]
   --sftp-known-hosts-file value           已知的主机文件的可选路径。 [$SFTP_KNOWN_HOSTS_FILE]
   --sftp-macs value                       以空格分隔的 MAC (message authentication code) 算法列表，按首选顺序排列。 [$SFTP_MACS]
   --sftp-md5sum-command value             用于读取 md5 散列的命令。 [$SFTP_MD5SUM_COMMAND]
   --sftp-pass value                       SSH 密码，留空以使用 ssh-agent。 [$SFTP_PASS]
   --sftp-path-override value              覆盖由 SSH shell 命令使用的路径。 [$SFTP_PATH_OVERRIDE]
   --sftp-port value                       SSH 端口号。 (默认值: "22") [$SFTP_PORT]
   --sftp-pubkey-file value                证书文件的可选路径。 [$SFTP_PUBKEY_FILE]
   --sftp-server-command value             指定在远程主机上运行 sftp 服务器路径或命令。 [$SFTP_SERVER_COMMAND]
   --sftp-set-env value                    要传递给 sftp 和命令的环境变量 [$SFTP_SET_ENV]
   --sftp-set-modtime value                如果设置了，设置远程文件的修改时间。 (默认值: "true") [$SFTP_SET_MODTIME]
   --sftp-sha1sum-command value            用于读取 sha1 散列的命令。 [$SFTP_SHA1SUM_COMMAND]
   --sftp-shell-type value                 远程服务器上的 SSH shell 类型，如果有的话。 [$SFTP_SHELL_TYPE]
   --sftp-skip-links value                 设置以跳过任何符号链接和任何其他非普通文件。 (默认值: "false") [$SFTP_SKIP_LINKS]
   --sftp-subsystem value                  指定远程主机上的 SSH2 子系统。 (默认值: "sftp") [$SFTP_SUBSYSTEM]
   --sftp-use-fstat value                  如果设置了，则使用 fstat 而不是 stat。 (默认值: "false") [$SFTP_USE_FSTAT]
   --sftp-use-insecure-cipher value        启用使用不安全的密码和密钥交换方法。 (默认值: "false") [$SFTP_USE_INSECURE_CIPHER]
   --sftp-user value                       SSH 用户名。 (默认值: "$USER") [$SFTP_USER]

```
{% endcode %}