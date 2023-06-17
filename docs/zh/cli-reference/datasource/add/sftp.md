# SSH/SFTP

{% code fullWidth="true" %}
```
名称:
   singularity datasource add sftp - SSH/SFTP

用法:
   singularity datasource add sftp [命令选项] <数据集名称> <源路径>

描述:
   --sftp-key-pem
      原始PEM编码的私钥。
      
      如果指定，将覆盖key_file参数。

   --sftp-subsystem
      指定远程主机上的SSH2子系统。

   --sftp-chunk-size
      上传和下载块大小。
      
      这控制了SFTP协议数据包中有效负载的最大大小。
      RFC将此限制为32768字节(32k)，这是默认值。然而，
      很多服务器支持更大的尺寸，通常限于最大
      256k的总包大小，并将其设置为更大，将使传输速度
      比高延迟链接更高。这包括OpenSSH和
      例如，在高达256k的总大小时，使用255k的值效果很好，
      为开销留下了充足的空间，同时仍在256k的总数据包大小范围内。
      
      在使用大于32k的值之前，请务必进行充分测试，
      并仅在始终连接到同一服务器或测试足够宽泛后才使用该值。如果你得到这样的错误
      "failed to send packet payload: EOF"、大量的"connection lost"、
      或在拷贝较大文件时出现"corrupted on transfer"，请尝试降低值。由[rclone serve sftp](/commands/rclone_serve_sftp)
      运行的服务器发送具有标准32k最大有效负载的数据包，因此您不能
      在下载文件时设置不同的chunk_size，但它可以接受
      长度达到256k的数据包，因此对于上传，chunk_size是可以设置的
      像上面的OpenSSH示例那样大。
      

   --sftp-macs
      通过首选项排序的以空格分隔的MAC(消息认证代码)算法列表。
      
      其中至少一个必须匹配服务器配置。例如，可以使用ssh -Q mac进行检查。
      
      例:
      
          umac-64-etm@openssh.com umac-128-etm@openssh.com hmac-sha2-256-etm@openssh.com
      

   --sftp-user
      SSH用户名。

   --sftp-port
      SSH端口号。

   --sftp-path-override
      覆盖SSH shell命令使用的路径。
      
      这允许在SFTP和SSH路径不同时进行校验和计算。这会影响到其他Synology NAS箱等。
      
      例如，如果共享文件夹可以在表示卷的目录中找到:
      
          rclone sync /home/local/directory remote:/directory --sftp-path-override /volume2/directory
      
      例如，如果主目录可以在名为"home"的共享文件夹中找到:
      
          rclone sync /home/local/directory remote:/home/directory --sftp-path-override /volume1/homes/USER/directory

   --sftp-use-fstat
      如果设置，使用fstat而不是stat。
      
      一些服务器限制了打开文件的数量，并且在打开文件之后调用Stat
      会从服务器中抛出错误。设置此标志将调用
      在已经打开的文件句柄上调用的Fstat而不是在Stat下调用。
      
      已经发现这对拥有"extractability"级别为1的IBM Sterling SFTP服务器很有帮助，这意味着只能在一个给定的时间内打开一个文件。
      

   --sftp-host
      SSH主机名。
      
      例如: "example.com"。

   --sftp-key-use-agent
      当设置时，强制使用ssh-agent。
      
      在也设置了key-file的情况下，会读取指定key-file的".pub"文件，并且仅请求关联的密钥。
      这样可以避免ssh-agent包含许多密钥时出现"Too many authentication failures for *username*"错误。

   --sftp-skip-links
      设置以跳过任何符号链接和任何其他非常规文件。

   --sftp-disable-concurrent-writes
      如果设置，不使用并发写。
      
      通常，rclone使用并发写来上传文件。这大大提高了
      性能，特别是对于远程服务器。
      
      如果有必要，此选项将禁用并发写。
      

   --sftp-set-env
      要传递给sftp和命令的环境变量
      
      以以下形式设置环境变量:
      
          VAR=value
      
      以传递到sftp客户端和运行的任何命令(例如md5sum)中的一组:
      
      用空格分隔多个变量，例如
      
          VAR1=value VAR2=value
      
      用引号包含带空格的变量，例如
      
          "VAR3=value with space" "VAR4=value with space" VAR5=nospacehere
      
      

   --sftp-ciphers
      用首选顺序排序的以空格分隔的密码列表，用于会话加密。
      
      其中至少一个必须与服务器配置匹配。例如，可以使用ssh-Q cipher进行检查。
      
      如果use_insecure_cipher为true，则不能设置此选项。
      
      例:
      
          aes128-ctr aes192-ctr aes256-ctr aes128-gcm@openssh.com aes256-gcm@openssh.com
      

   --sftp-shell-type
      远程服务器上的SSH shell的类型(如果存在)。
      
      留空以自动检测。

      示例:
         | none       | 无shell访问
         | unix       | Unix shell
         | powershell | PowerShell
         | cmd        | Windows Command Prompt

   --sftp-key-exchange
      用首选顺序排序的一组密钥交换算法。
      
      其中至少一个必须匹配服务器配置。例如，可以使用ssh-Q kex进行检查。
      
      如果use_insecure_cipher为true，则不能设置此选项。
      
      例子:
      
          sntrup761x25519-sha512@openssh.com curve25519-sha256 curve25519-sha256@libssh.org ecdh-sha2-nistp256
      

   --sftp-idle-timeout
      关闭空闲连接之前的最长时间。
      
      如果在给定的时间内没有将任何连接返回到连接池中，则rclone将清空连接池。
      
      设置为0以无限期保持连接。
      

   --sftp-concurrency
      一个文件的最大未完成请求数
      
      这控制了一个文件的最大未完成请求数。
      增加它将提高高延迟链接的吞吐量，
      但使用更多的内存成本。
      

   --sftp-use-insecure-cipher
      启用使用不安全的密码和密钥交换方法。
      
      这启用了以下不安全的密码和密钥交换方法:
      
      - aes128-cbc
      - aes192-cbc
      - aes256-cbc
      - 3des-cbc
      - diffie-hellman-group-exchange-sha256
      - diffie-hellman-group-exchange-sha1
      
      这些算法是不安全的，可能会允许攻击者恢复明文数据。
      
      如果您使用了密码或密钥交换高级选项，则必须为false。
      

      示例:
         | false | 使用默认密码列表。
         | true  | 启用了aes128-cbc
设置此值以启用服务器主机密钥验证。

前缀`〜`将在文件名中扩展，环境变量如`${RCLONE_CONFIG_DIR}`也是如此。

示例：
| ~ / .ssh / known_hosts | 使用OpenSSH的已知主机文件。

-- sftp-ask-password
允许在需要密码时询问SFTP密码。

如果设置了此选项并且没有提供密码，则rclone将：
- 要求密码
- 不联系ssh代理

--sftp-pass
SSH密码，留空以使用ssh-agent。

--sftp-key-file
PEM编码私钥文件的路径。

留空或将key-use-agent设置为使用ssh-agent。

前缀`〜`将在文件名中扩展，环境变量如`${RCLONE_CONFIG_DIR}`也是如此。

--sftp-md5sum-command
用于读取md5哈希的命令。

留空以进行自动检测。


选项：
--help，-h 显示帮助

数据准备选项

--delete-after-export [危险]将数据集导出到CAR文件后删除数据集文件。 （默认值：false）
--rescan-interval value在上次成功扫描后，当此间隔时间过去时自动重新扫描源目录（默认值：禁用）

SFTP选项

--sftp-ask-password value允许在需要密码时询问SFTP密码。 （默认值：“false”）[$SFTP_ASK_PASSWORD]
--sftp-chunk-size value上传和下载块的大小。（默认值：“32Ki”）[$SFTP_CHUNK_SIZE]
--sftp-ciphers value按优先级排序的用于会话加密的加密算法的空格分隔列表。[$SFTP_CIPHERS]
--sftp-concurrency value一个文件的最大未完成请求数（默认值：“64”）[$SFTP_CONCURRENCY]
--sftp-disable-concurrent-reads value如果设置，不要使用并行读取。 （默认值：“false”）[$SFTP_DISABLE_CONCURRENT_READS]
--sftp-disable-concurrent-writes value如果设置，不要使用并发写入。 （默认值：“false”）[$SFTP_DISABLE_CONCURRENT_WRITES]
--sftp-disable-hashcheck value禁用执行SSH命令来确定是否可以远程文件哈希。 （默认值：“false”）[$SFTP_DISABLE_HASHCHECK]
--sftp-host value SSH连接到的主机。[$SFTP_HOST]
--sftp-idle-timeout value关闭空闲连接之前的最长时间。（默认值：“1m0s”）[$SFTP_IDLE_TIMEOUT]
--sftp-key-exchange value按优先级排序的密钥交换算法的空格分隔列表。[$SFTP_KEY_EXCHANGE]
--sftp-key-file valuePEM编码私钥文件的路径。[$SFTP_KEY_FILE]
--sftp-key-file-pass value解密PEM编码私钥文件的密码。[$SFTP_KEY_FILE_PASS]
--sftp-key-pem value原始PEM编码私钥。[$SFTP_KEY_PEM]
--sftp-key-use-agent 值设置时强制使用ssh-agent。 （默认值：“false”）[$SFTP_KEY_USE_AGENT]
--sftp-known-hosts-file value可选的已知主机文件路径。[$SFTP_KNOWN_HOSTS_FILE]
--sftp-macs value按优先级排序的MAC（消息认证码）算法的空格分隔列表。[$SFTP_MACS]
--sftp-md5sum-command value用于读取MD5哈希的命令。[$SFTP_MD5SUM_COMMAND]
--sftp-pass value SSH密码，留空以使用ssh-agent。[$SFTP_PASS]
--sftp-path-override value覆盖SSH shell命令使用的路径。[$SFTP_PATH_OVERRIDE]
--sftp-port value SSH端口号。 （默认值：“22”）[$SFTP_PORT]
--sftp-pubkey-file value可选路径到公钥文件。[$SFTP_PUBKEY_FILE]
--sftp-server-command value指定在远程主机上运行sftp服务器的路径或命令。[$SFTP_SERVER_COMMAND]
--sftp-set-env value传递给sftp和命令的环境变量[$SFTP_SET_ENV]
--sftp-set-modtime value如果设置，则在远程上设置修改时间。 （默认值：“true”）[$SFTP_SET_MODTIME]
--sftp-sha1sum-command value用于读取sha1哈希的命令。[$SFTP_SHA1SUM_COMMAND]
--sftp-shell-type value远程服务器上SSH shell的类型（如果有）。[$SFTP_SHELL_TYPE]
--sftp-skip-links value设置为跳过任何符号链接和任何其他非常规文件。 （默认值：“false”）[$SFTP_SKIP_LINKS]
--sftp-subsystem value指定远程主机上的SSH2子系统。 （默认值：“sftp”）[$SFTP_SUBSYSTEM]
--sftp-use-fstat value如果设置，请使用fstat而不是stat。 （默认值：“false”）[$SFTP_USE_FSTAT]
--sftp-use-insecure-cipher value启用使用不安全的密码和密钥交换方法。 （默认值：“false”）[$SFTP_USE_INSECURE_CIPHER]
--sftp-user 值SSH用户名。 （默认值：“shane”）[$SFTP_USER]