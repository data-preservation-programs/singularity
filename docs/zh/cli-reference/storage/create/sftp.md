# SSH/SFTP

```
名称:
   singularity storage create sftp - SSH/SFTP

用法:
   singularity storage create sftp [命令选项] [参数...]

描述:
   --host
      要连接的SSH主机。
      
      例如 "example.com"。

   --user
      SSH用户名。

   --port
      SSH端口号。

   --pass
      SSH密码，留空以使用ssh-agent。

   --key-pem
      原始的PEM编码的私钥。
      
      如果指定，将覆盖key_file参数。

   --key-file
      PEM编码的私钥文件的路径。
      
      留空或将key-use-agent设置为"use ssh-agent"。
      
      文件名中的开始`~`会被扩展，环境变量例如`${RCLONE_CONFIG_DIR}`也会被扩展。

   --key-file-pass
      解密PEM编码的私钥文件用的口令。
      
      仅支持PEM加密的密钥文件（旧的OpenSSH格式）。不能使用新的OpenSSH格式的加密的密钥。

   --pubkey-file
      可选的公钥文件的路径。
      
      如果有已签名的证书要用于身份验证，请设置此选项。
      
      文件名中的开始`~`会被扩展，环境变量例如`${RCLONE_CONFIG_DIR}`也会被扩展。

   --known-hosts-file
      可选的known_hosts文件的路径。
      
      设置此值以启用服务器主机密钥验证。
      
      文件名中的开始`~`会被扩展，环境变量例如`${RCLONE_CONFIG_DIR}`也会被扩展。

      示例:
         | ~/.ssh/known_hosts | 使用OpenSSH的known_hosts文件。

   --key-use-agent
      当设置时，强制使用ssh-agent。
      
      当还设置了key-file时，将读取指定密钥文件的".pub"文件，并且只会从ssh-agent请求相关的密钥。这可以避免在ssh-agent包含许多密钥时出现“Too many authentication failures for *username*”错误。

   --use-insecure-cipher
      启用不安全的密码和密钥交换方法的使用。
      
      这会启用以下不安全的密码和密钥交换方法：
      
      - aes128-cbc
      - aes192-cbc
      - aes256-cbc
      - 3des-cbc
      - diffie-hellman-group-exchange-sha256
      - diffie-hellman-group-exchange-sha1
      
      这些算法是不安全的，攻击者可能会通过它们来恢复明文数据。
      
      如果使用密码或key_exchange高级选项，请将该选项设置为false。
      

      示例:
         | false | 使用默认密码列表。
         | true  | 启用aes128-cbc密码和diffie-hellman-group-exchange-sha256、diffie-hellman-group-exchange-sha1密钥交换。

   --disable-hashcheck
      禁用执行SSH命令以确定远程文件哈希是否可用。
      
      留空或将其设置为false以启用哈希计算（推荐），设置为true以禁用哈希计算。

   --ask-password
      允许在需要时询问SFTP密码。
      
      如果设置了此选项且未提供密码，则rclone将：
      - 请求密码
      - 不会询问ssh代理
      
   --path-override
      重写SSH shell命令使用的路径。
      
      这样可以在SFTP和SSH路径不同的情况下进行校验和计算。这个问题会影响到 Synology NAS 等设备。
      
      如果共享文件夹可以在表示卷的目录中找到：
      
          rclone sync /home/local/directory remote:/directory --sftp-path-override /volume2/directory
      
      如果主目录可以在名为“home”的共享文件夹中找到：
      
          rclone sync /home/local/directory remote:/home/directory --sftp-path-override /volume1/homes/USER/directory

   --set-modtime
      如果设置了，就将远程文件的修改时间设置为与本地文件一致。

   --shell-type
      远程服务器上的SSH shell类型，如果有的话。
      
      留空以进行自动检测。

      示例:
         | none       | 没有shell访问权限
         | unix       | Unix shell
         | powershell | PowerShell
         | cmd        | Windows命令提示符

   --md5sum-command
      读取md5哈希的命令。
      
      留空以进行自动检测。

   --sha1sum-command
      读取sha1哈希的命令。
      
      留空以进行自动检测。

   --skip-links
      设置为跳过任何符号链接和任何其他非常规文件。

   --subsystem
      指定远程主机上的SSH2子系统。

   --server-command
      指定在远程主机上运行SFTP服务器的路径或命令。
      
      如果定义了server_command，将忽略subsystem选项。

   --use-fstat
      如果设置，则使用fstat而不是stat。
      
      一些服务器限制打开的文件数量，并且在打开文件之后调用Stat将引发服务器错误。设置此标志将在已打开的文件句柄上调用Fstat而不是Stat。
      
      已发现，这对于IBM Sterling SFTP服务器有帮助，该服务器的“extractability”级别设置为1，这意味着任何给定时间只能打开一个文件。

   --disable-concurrent-reads
      如果设置，则不使用并发读取。
      
      通常，并发读取是安全的，并且不使用它们会降低性能，因此此选项默认禁用。
      
      一些服务器限制文件的下载次数。使用并发读取可能触发此限制，因此如果您的服务器返回
      
          Failed to copy: file does not exist
      
      那么您可能需要启用此标志。
      
      如果禁用了并发读取，则忽略use_fstat选项。

   --disable-concurrent-writes
      如果设置，则不使用并发写入。
      
      通常，rclone使用并发写入来上传文件。这大大提高了性能，尤其对于远程服务器而言。
      
      如有必要，此选项禁用并发写入。

   --idle-timeout
      空闲连接关闭的最大时间。
      
      如果在给定的时间内没有将连接返回到连接池中，则rclone将清空连接池。
      
      将其设置为0以无限期保持连接。

   --chunk-size
      上传和下载块的大小。
      
      这控制SFTP协议数据包中有效负载的最大大小。
      RFC将其限制为32768字节（32k），这是默认值。然而很多服务器支持更大的大小，通常限制为最大总包大小为256k，将其设置为更大将大大增加高延迟链接上的传输速度。这个包括OpenSSH，例如，使用255k的值效果很好，留有足够的空间用于开销，同时仍然在256k的总包大小范围内。
      
      在使用大于32k的值之前，请务必进行全面测试，并且仅在始终连接到相同的服务器或经过充分广泛的测试后使用。如果在复制较大的文件时出现“failed to send packet payload: EOF”，大量的“connection lost”或“corrupted on transfer”等错误，请尝试降低该值。由[rclone serve sftp](/commands/rclone_serve_sftp)运行的服务器将使用标准的32k最大有效负载发送数据包，因此当下载文件时不能设置不同的chunk_size，但它接受的数据包的总大小可达256k，因此对于上传，chunk_size可以像上面的OpenSSH示例那样设置。

   --concurrency
      一个文件的同时待处理请求的最大数量。
      
      这控制一个文件的同时待处理请求的最大数量。增加该数量可以提高高延迟链接上的吞吐量，但会使用更多的内存。

   --set-env
      要传递给sftp和命令的环境变量。
      
      以以下形式设置环境变量：
      
          VAR=value
      
      以及以空格分隔的多个变量，例如
      
          VAR1=value VAR2=value
      
      并且以引号括起来的带有空格的变量，例如
      
          "VAR3=value with space" "VAR4=value with space" VAR5=nospacehere
      
      

   --ciphers
      以空格分隔的加密密码列表，按偏好顺序排列。
      
      至少有一个密码必须与服务器配置匹配。可以使用ssh -Q cipher检查配置。
      
      如果use_insecure_cipher为true，则不能设置此选项。
      
      示例:
      
          aes128-ctr aes192-ctr aes256-ctr aes128-gcm@openssh.com aes256-gcm@openssh.com
      

   --key-exchange
      以空格分隔的密钥交换算法列表，按偏好顺序排列。
      
      至少有一个算法必须与服务器配置匹配。可以使用ssh -Q kex检查配置。
      
      如果use_insecure_cipher为true，则不能设置此选项。
      
      示例:
      
          sntrup761x25519-sha512@openssh.com curve25519-sha256 curve25519-sha256@libssh.org ecdh-sha2-nistp256
      

   --macs
      以空格分隔的MAC（消息认证码）算法列表，按偏好顺序排列。
      
      至少有一种算法必须与服务器配置匹配。可以使用ssh -Q mac检查配置。
      
      示例:
      
          umac-64-etm@openssh.com umac-128-etm@openssh.com hmac-sha2-256-etm@openssh.com
      


选项:
   --disable-hashcheck    禁用执行SSH命令以确定远程文件哈希是否可用。 (默认值: false) [$DISABLE_HASHCHECK]
   --help, -h             显示帮助
   --host value           要连接的SSH主机。 [$HOST]
   --key-file value       PEM编码的私钥文件的路径。 [$KEY_FILE]
   --key-file-pass value  解密PEM编码的私钥文件用的口令。 [$KEY_FILE_PASS]
   --key-pem value        原始的PEM编码的私钥。 [$KEY_PEM]
   --key-use-agent        当设置时，强制使用ssh-agent。 (默认值: false) [$KEY_USE_AGENT]
   --pass value           SSH密码，留空以使用ssh-agent。 [$PASS]
   --port value           SSH端口号。 (默认值: 22) [$PORT]
   --pubkey-file value    可选的公钥文件的路径。 [$PUBKEY_FILE]
   --use-insecure-cipher  启用不安全的密码和密钥交换方法的使用。 (默认值: false) [$USE_INSECURE_CIPHER]
   --user value           SSH用户名。 (默认值: "$USER") [$USER]

   高级

   --ask-password               允许在需要时询问SFTP密码。 (默认值: false) [$ASK_PASSWORD]
   --chunk-size value           上传和下载块的大小。 (默认值: "32Ki") [$CHUNK_SIZE]
   --ciphers value              以空格分隔的加密密码列表，按偏好顺序排列。 [$CIPHERS]
   --concurrency value          一个文件的同时待处理请求的最大数量。 (默认值: 64) [$CONCURRENCY]
   --disable-concurrent-reads   如果设置，则不使用并发读取。 (默认值: false) [$DISABLE_CONCURRENT_READS]
   --disable-concurrent-writes  如果设置，则不使用并发写入。 (默认值: false) [$DISABLE_CONCURRENT_WRITES]
   --idle-timeout value         空闲连接关闭的最大时间。 (默认值: "1m0s") [$IDLE_TIMEOUT]
   --key-exchange value         以空格分隔的密钥交换算法列表，按偏好顺序排列。 [$KEY_EXCHANGE]
   --known-hosts-file value     可选的known_hosts文件的路径。 [$KNOWN_HOSTS_FILE]
   --macs value                 以空格分隔的MAC（消息认证码）算法列表，按偏好顺序排列。 [$MACS]
   --md5sum-command value       读取md5哈希的命令。 [$MD5SUM_COMMAND]
   --path-override value        重写SSH shell命令使用的路径。 [$PATH_OVERRIDE]
   --server-command value       指定在远程主机上运行SFTP服务器的路径或命令。 [$SERVER_COMMAND]
   --set-env value              要传递给sftp和命令的环境变量 [$SET_ENV]
   --set-modtime                如果设置了，就将远程文件的修改时间设置为与本地文件一致。 (默认值: true) [$SET_MODTIME]
   --sha1sum-command value      读取sha1哈希的命令。 [$SHA1SUM_COMMAND]
   --shell-type value           远程服务器上的SSH shell类型，如果有的话。 [$SHELL_TYPE]
   --skip-links                 设置为跳过任何符号链接和任何其他非常规文件。 (默认值: false) [$SKIP_LINKS]
   --subsystem value            指定远程主机上的SSH2子系统。 (默认值: "sftp") [$SUBSYSTEM]
   --use-fstat                  如果设置，则使用fstat而不是stat。 (默认值: false) [$USE_FSTAT]

   常规

   --name value  存储的名称 (默认值: 自动生成)
   --path value  存储的路径
```